package commands

import (
	"fmt"
	"os"
	"text/tabwriter"

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

	projects, err := rexos.GetAllProjectsByOwner(handler, session.UserID)

	tw := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintf(tw, "Urn\tName\tType\tScheme\tPublic\t#ProjectFiles\n")
	for _, p := range projects {
		fmt.Fprintf(tw, "%s\t%s\t%s\t%s\t%v\t%d\n", p.Urn, p.Name, p.Type, p.Scheme, p.Public, p.NumberOfProjectFiles)
	}
	tw.Flush()

	// p, err := rexos.GetProjectByUrn(handler, "robotic-eyes:project:3147")
	// if err != nil {
	// 	color.Red.Println("Error:", err)
	// 	return nil
	// }
	// fmt.Println(p)
	return nil
}
