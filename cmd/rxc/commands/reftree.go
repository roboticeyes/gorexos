package commands

import (
	"fmt"

	"github.com/gookit/color"
	"github.com/roboticeyes/gorexos/pkg/http/rexos"
	"github.com/urfave/cli/v2"
)

// ReferenceTreeCommand performs operations on a REX project
var ReferenceTreeCommand = &cli.Command{
	Name:   "tree",
	Usage:  "Displays the reference tree for a given project",
	Action: referenceTreeAction,
}

func referenceTreeAction(ctx *cli.Context) error {

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

	root, err := rexos.GetReferenceTreeByProjectUrn(handler, "robotic-eyes:project:9232")
	if err != nil {
		color.Red.Println("Error:", err)
		return nil
	}

	fmt.Println(root)
	return nil
}
