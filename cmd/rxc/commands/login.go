package commands

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// LoginCommand performs a login to the REXos platform with the given credentials
var LoginCommand = &cli.Command{
	Name:   "login",
	Usage:  "Login into REXos",
	Action: loginAction,
}

func loginAction(ctx *cli.Context) error {

	fmt.Println("Logging in ...")

	fmt.Println("config: ", ctx.String("config"))

	return nil
}
