package main

import (
	"os"

	"github.com/SierraSoftworks/stablehand/commands"
	"github.com/SierraSoftworks/stablehand/config"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "Stablehand"
	app.Usage = "A tool to help you keep your Rancher server clean in production"

	app.Author = "Benjamin Pannell"
	app.Email = "admin@sierrasoftworks.com"
	app.Copyright = "Sierra Softworks © 2016"
	app.Version = version

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "server",
			EnvVar:      "CATTLE_URL",
			Destination: &config.Rancher.Server,
			Usage:       "The URL of your Rancher server",
		},
		cli.StringFlag{
			Name:        "key",
			EnvVar:      "CATTLE_ACCESS_KEY",
			Destination: &config.Rancher.AccessKey,
			Usage:       "The access key used to sign into your Rancher server",
		},
		cli.StringFlag{
			Name:        "secret",
			EnvVar:      "CATTLE_SECRET_KEY",
			Destination: &config.Rancher.SecretKey,
			Usage:       "The secret key used to sign into your Rancher server",
		},
	}

	app.Commands = cli.Commands{
		commands.List,
		commands.Deactivate,
		commands.Remove,
		commands.Purge,
	}

	app.Run(os.Args)
}
