package commands

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/gookit/color"
	"github.com/roboticeyes/gorexos/pkg/http/rexos"
	"github.com/urfave/cli/v2"
)

type TargetInput struct {
	Target      rexos.Target       `json:"target"`
	TargetFiles []rexos.TargetFile `json:"targetFiles"`
}

// TargetCommand performs operations on a REX target
var TargetCommand = &cli.Command{
	Name:  "targets",
	Usage: "Target operations in REXos",
	Subcommands: []*cli.Command{
		{
			Name:   "create",
			Usage:  "create a new target",
			Action: createTarget,
		},
		{
			Name:  "delete",
			Usage: "deletes a target",
			Action: func(c *cli.Context) error {
				fmt.Println("not implemented")
				return nil
			},
		},
	},
}

func createTarget(ctx *cli.Context) error {

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

	targetInput := TargetInput{
		Target: rexos.DefaultTarget(),
	}

	if ctx.NArg() > 0 {
		fileName := ctx.Args().Get(0)
		f, err := os.Open(fileName)
		if err != nil {
			panic(err)
		}
		defer f.Close()
		buf, err := ioutil.ReadAll(f)
		if err != nil {
			panic(err)
		}
		err = json.Unmarshal(buf, &targetInput)
	} else {
		color.Red.Println("Missing input JSON file")
		return nil
	}

	fmt.Println(targetInput.Target)

	res, err := rexos.CreateTarget(handler, targetInput.Target)
	if err != nil {
		color.Red.Println(err)
	}

	// Upload all target files
	for _, v := range targetInput.TargetFiles {
		fmt.Println("Uploading target file", v.LocalFile)
		targetFileUrn, err := rexos.CreateTargetFile(handler, res.Urn, v)
		if err != nil {
			color.Red.Println("Cannot create target file:", err)
		}
		err = rexos.UploadTargetFile(handler, res.Urn, targetFileUrn, v.LocalFile)
		if err != nil {
			color.Red.Println("Cannot create target file:", err)
		}
	}

	// https://test.rex.codes/v1/73f03414-6694-4173-a5ad-a5c74d5416af

	color.Green.Println("Successfully created target:", res.Urn)

	// fmt.Println("DELETE project")
	// handler.Delete("/inspection/v1/targets/"+res.Urn, "")
	return nil
}
