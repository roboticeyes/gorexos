package main

import (
	"log"
	"os"

	"github.com/roboticeyes/rexos/cmd/rxc/commands"
	"github.com/urfave/cli/v2"
)

const (
	version = "v0.1"
)

func main() {

	// c := config.ReadConfig("/home/breiting/.config/rexos/config.yml")
	// fmt.Println(c)

	app := cli.NewApp()
	app.Name = "rxc"
	app.Usage = "REXos control client"
	app.Version = version
	app.Copyright = "(c) 2020 Robotic Eyes GmbH"
	app.EnableBashCompletion = true
	app.Flags = GlobalFlags

	app.Commands = []*cli.Command{
		commands.LoginCommand,
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
