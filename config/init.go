package config

import (
	"github.com/easyasync/easyaysnc-cli/commands"
	"github.com/urfave/cli"
)

//CreateApp method to create initial app
func CreateApp() *cli.App {
	app := cli.NewApp()

	app.Name = "easyasync"
	app.Usage = "Command line interface for easyasync cli"
	app.Version = "0.0.1"
	app.Author = "VSS"

	app.Commands = []cli.Command{
		commands.Register(),
	}

	return app
}
