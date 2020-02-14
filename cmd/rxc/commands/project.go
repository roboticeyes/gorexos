package commands

import (
	"fmt"

	"github.com/gookit/color"
	"github.com/roboticeyes/gorexos/pkg/http/rexos"
	"github.com/urfave/cli/v2"
)

// ProjectCommand performs operations on a REX project
var ProjectCommand = &cli.Command{
	Name:   "projects",
	Usage:  "Project operations in REXos",
	Action: projectAction,
}

func projectAction(ctx *cli.Context) error {

	session, err := rexos.OpenSession(SessionFile)
	if !session.Valid() {
		color.Red.Println("Session is not valid and is expired on ", session.Expires)
		return nil
	}
	handler := rexos.NewRequestHandler()
	err = handler.AuthenticateWithSession(session)
	if err != nil {
		color.Red.Println("Cannot authenticate, please use login")
	}

	p, err := rexos.GetProjectByUrn(handler, "robotic-eyes:project:3147")
	if err != nil {
		color.Red.Println("Error:", err)
		return nil
	}
	fmt.Println(p)
	return nil
}
