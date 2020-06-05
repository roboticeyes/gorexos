package commands

import (
	"fmt"
	"os"
	"time"

	"github.com/gookit/color"
	"github.com/roboticeyes/gorexos/pkg/http/rexos"
	"github.com/roboticeyes/gorexos/pkg/http/rexos/translation"
	"github.com/urfave/cli/v2"
)

// TranslateCommand performs operations for the Translation composite
var TranslateCommand = &cli.Command{
	Name:   "translate",
	Usage:  "Translates an input geometry to a REXfile usine the REX translation composite service",
	Action: translateAction,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "pipeline",
			Usage:    "Specifies the type of pipeline to be used (standard [default], ifc)",
			Required: false,
			Aliases:  []string{"p"},
		},
	},
}

func translateAction(ctx *cli.Context) error {

	if ctx.Args().Len() < 2 {
		color.Red.Println("Please provide at least an input and output file as arguments.")
		return nil
	}

	pipeline := ctx.String("pipeline")
	if pipeline == "" {
		pipeline = "standard"
	}

	session, err := rexos.OpenStoredSession()
	if err != nil {
		return err
	}
	if !session.Valid() {
		color.Red.Println("Session is not valid and is expired on ", session.Expires)
		return nil
	}
	handler := rexos.NewRequestHandler()
	err = handler.AuthenticateWithSession(session)
	if err != nil {
		color.Red.Println("Cannot authenticate, please use login")
	}

	masterInputFile := ctx.Args().First()

	job := translation.Job{
		Name: masterInputFile,
	}
	color.Green.Print("Creating job ... ", job.ID)
	job, err = translation.CreateJob(handler, job)
	if err != nil {
		color.Red.Println("FAILED - ", err)
		return err
	}
	color.Green.Printf("success (id=%s)\n", job.ID)

	// upload master file
	err = uploadFile(handler, job.ID, masterInputFile)
	if err != nil {
		panic(err)
	}

	// upload rest of the files except the last one which should be the output
	for i := 1; i < ctx.Args().Len()-1; i++ {
		err = uploadFile(handler, job.ID, ctx.Args().Get(i))
		if err != nil {
			panic(err)
		}
	}

	// start job
	color.Green.Print("Starting job ... ")
	err = translation.StartJob(handler, job.ID, pipeline)
	if err != nil {
		color.Red.Println("FAILED - ", err)
		return err
	}
	color.Green.Printf("success\n")

	color.Green.Print("Polling job status [")
	for {
		fmt.Print(".")
		j, err := translation.GetJob(handler, job.ID)
		if err != nil {
			panic(err)
		}
		if j.Status == "done" {
			break
		}
		time.Sleep(1 * time.Second)
	}
	color.Green.Println("] done.")

	color.Green.Print("Downloading file ... ")
	outputFileName := ctx.Args().Get(ctx.Args().Len() - 1)
	result, err := os.Create(outputFileName)
	if err != nil {
		panic(err)
	}
	defer result.Close()
	translation.GetJobResult(handler, job.ID, result)
	color.Green.Print("done.")

	return nil
}

func uploadFile(handler rexos.RequestHandler, jobId, file string) error {
	color.Green.Printf("Uploading %s ... ", file)
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = translation.UploadFile(handler, jobId, file, f)
	if err != nil {
		color.Red.Println("FAILED - ", err)
		return err
	}
	color.Green.Printf("success\n")
	return nil
}
