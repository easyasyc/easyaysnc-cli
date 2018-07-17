package commands

import (
	"fmt"

	"github.com/urfave/cli"
)

//Register is a command to register a route to a service
func Register() cli.Command {
	return cli.Command{

		Name:    "register",
		Aliases: []string{"r"},
		Usage:   "register an app with async event",
		Action: func(c *cli.Context) error {
			fmt.Printf("Service is %q\n", c.Args().Get(0))
			return nil
		},
	}
}
