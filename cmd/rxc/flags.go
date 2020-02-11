package main

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/urfave/cli/v2"
)

func userRexOSDir() string {
	if runtime.GOOS == "windows" {
		home := os.Getenv("HOMEDRIVE") + os.Getenv("HOMEPATH")
		if home == "" {
			home = os.Getenv("USERPROFILE")
		}
		return filepath.Join(home, "rexos")
	} else if runtime.GOOS == "linux" {
		home := os.Getenv("XDG_CONFIG_HOME")
		if home != "" {
			return filepath.Join(home, "rexos")
		}
	}
	return filepath.Join(os.Getenv("HOME"), ".config", "rexos")
}

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
		Value: userRexOSDir() + "/config.yml",
		EnvVars: []string{
			"REXCTL_CONFIG",
		},
	},
}
