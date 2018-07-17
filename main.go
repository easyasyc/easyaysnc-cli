package main

import (
	"log"
	"os"

	"github.com/easyasync/easyaysnc-cli/commands"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "easyasync"
	app.Usage = "Command line interface for easyasync cli"
	app.Version = "0.0.1"
	app.Author = "VSS"

	app.Commands = []cli.Command{
		commands.Register(),
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
