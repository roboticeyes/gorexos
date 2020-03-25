package commands

import (
	"os"
	"os/exec"

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
	},
}

func referenceTreeAction(ctx *cli.Context) error {

	urn := ctx.String("urn")

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

	f, _ := os.Create("/tmp/xxx.dot")
	defer f.Close()
	err = refTree.WriteToDot(f)
	// err = refTree.WriteToDot(os.Stdout)
	if err != nil {
		color.Red.Println("Cannot dump as dotfile:", err)
	}

	cmd := exec.Command("dot", "-Tpng", "-o", "/tmp/xxx.png", "/tmp/xxx.dot")
	if err := cmd.Run(); err != nil {
		color.Red.Println("Cannot execute dot command:", err)
	}

	return nil
}
