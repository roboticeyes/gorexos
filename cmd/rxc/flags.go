package main

import (
	"github.com/roboticeyes/gorexos/pkg/config"
	"github.com/urfave/cli/v2"
)

// GlobalFlags are global CLI flags
var GlobalFlags = []cli.Flag{
	&cli.BoolFlag{
		Name:  "debug",
		Usage: "run in debug mode",
		EnvVars: []string{
			"REXCTL_DEBUG",
		},
	},
	&cli.IntFlag{
		Name:  "workers, w",
		Usage: "number of workers for indexing",
	},
	&cli.StringFlag{
		Name:  "config",
		Usage: "Specify config file",
		Value: config.UserRexOSDir() + "/config.yml",
		EnvVars: []string{
			"REXCTL_CONFIG",
		},
	},
}
