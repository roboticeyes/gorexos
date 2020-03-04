package commands

import (
	"fmt"

	"github.com/gookit/color"
	"github.com/roboticeyes/gorexos/pkg/http/rexos"
	"github.com/urfave/cli/v2"
)

// StatisticsCommand shows the user statistics
var StatisticsCommand = &cli.Command{
	Name:   "stats",
	Usage:  "Statistics of the current user",
	Action: statsAction,
}

func statsAction(ctx *cli.Context) error {

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

	stats, err := rexos.GetUserStatistics(handler, session.UserID)

	fmt.Println("NumberOfProjects              :", stats.NumberOfProjects)
	fmt.Printf("TotalUsedDiskSpace [MB]       : %.2f\n", float64(stats.TotalUsedDiskSpace)/1000.0/1000.0)
	fmt.Println("NumberOfPublicShareActions    :", stats.NumberOfPublicShareActions)
	fmt.Println("MaxNumberOfPublicShareActions :", stats.MaxNumberOfPublicShareActions)

	return nil
}
