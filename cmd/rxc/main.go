package main

import (
	"os"

	"github.com/gookit/color"
	"github.com/roboticeyes/gorexos/cmd/rxc/commands"
	"github.com/roboticeyes/gorexos/pkg/config"
	"github.com/urfave/cli/v2"
)

var (
	// Version string from ldflags
	Version string
	// Build string from ldflags
	Build string
)

func main() {

	app := cli.NewApp()
	app.Name = "rxc"
	app.Usage = "REXos control client"
	app.Version = Version + " - build " + Build
	app.Description = `
	This tool allows to communicate with REXos. The first operation is to login by calling rxc login. Please make
	sure that you have a proper config file. Under Linux you can put your default config under ~/.config/rexos/config.yml.
	A sample can be found in the sample folder. The login operation caches the token on your local machine, and other
	commands can pick up the credentials. Simply perform a rxc projects to list your projects.

	Login by: rxc login
	List your projects: rxc projects

	You may also call rxc withou any parameters which starts a permanent session handler in the terminal.
	`
	app.Copyright = "(c) 2020 Robotic Eyes GmbH"
	app.EnableBashCompletion = true
	app.Flags = GlobalFlags
	app.Before = func(c *cli.Context) error {

		configReader, err := os.Open(c.String("config"))
		if err != nil {
			color.Red.Printf("Cannot open config file %s. You can also set a config file with the --config flag\n", c.String("config"))
			return err
		}
		config := config.ReadConfig(configReader)
		c.App.Metadata = make(map[string]interface{})
		c.App.Metadata["config"] = config
		return nil
	}

	app.Action = func(c *cli.Context) error {
		NewApp(c).Run()
		return nil
	}

	app.Commands = []*cli.Command{
		commands.LoginCommand,
		commands.ProjectCommand,
		commands.StatisticsCommand,
		commands.ReferenceTreeCommand,
		commands.TargetCommand,
	}

	err := app.Run(os.Args)
	if err != nil {
		//log.Fatal(err)
	}

}
