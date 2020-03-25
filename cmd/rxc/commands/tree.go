package commands

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/gookit/color"
	"github.com/roboticeyes/gorexos/pkg/http/rexos"
	"github.com/urfave/cli/v2"
)

// ReferenceTreeCommand performs operations on a REX project
var ReferenceTreeCommand = &cli.Command{
	Name:   "tree",
	Usage:  "Displays the reference tree for a given project",
	Action: referenceTreeAction,
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:     "urn",
			Usage:    "Project URN",
			Required: true,
		},
		&cli.StringFlag{
			Name:    "output",
			Usage:   "Output dot file",
			Aliases: []string{"o"},
		},
		&cli.StringFlag{
			Name:  "format",
			Usage: "Image format, default is png (e.g. png, svg)",
		},
	},
}

func referenceTreeAction(ctx *cli.Context) error {

	urn := ctx.String("urn")
	output := ctx.String("output")
	format := ctx.String("format")

	if format == "" {
		format = "png"
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

	refTree, err := rexos.GetReferenceTreeByProjectUrn(handler, urn)
	if err != nil {
		color.Red.Println("Error:", err)
		return nil
	}

	refTree.Beautify()

	if output == "" {
		err = refTree.WriteToDot(os.Stdout)
		if err != nil {
			color.Red.Println("Cannot generate dotfile:", err)
		}
		return nil
	}

	f, _ := os.Create(output)
	defer f.Close()
	err = refTree.WriteToDot(f)
	if err != nil {
		color.Red.Println("Cannot write dotfile:", err)
	}

	// check if dot command is available, if so, output an image file as well
	dot, err := exec.LookPath("dot")
	if err != nil {
		return nil
	}

	imageFile := output[:len(output)-len(filepath.Ext(output))] + "." + format
	cmd := exec.Command(dot, "-T"+format, "-o", imageFile, output)
	if err := cmd.Run(); err != nil {
		color.Red.Println("Cannot execute dot command:", err)
	}

	return nil
}
