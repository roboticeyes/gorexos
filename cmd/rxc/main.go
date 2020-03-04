package main

import (
	"os"

	"github.com/gookit/color"
	"github.com/roboticeyes/gorexos/cmd/rxc/commands"
	"github.com/roboticeyes/gorexos/pkg/config"
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

	app.Commands = []*cli.Command{
		commands.LoginCommand,
		commands.ProjectCommand,
	}

	err := app.Run(os.Args)
	if err != nil {
		//log.Fatal(err)
	}

}
