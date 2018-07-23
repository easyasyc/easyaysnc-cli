package commands

import (
	"github.com/urfave/cli"
)

//RegisterAction interface to define register command actions
type RegisterAction interface {
	RegisterService(sourcename string, route string) error
}

//RegisterCommand interface to define method to return cli.Command for a RegisterCommand
type RegisterCommand interface {
	Register(action RegisterAction) cli.Command
}

//Register is a command to register a route to a service
func Register(action RegisterAction) cli.Command {
	return cli.Command{

		Name:    "register",
		Aliases: []string{"r"},
		Usage:   "register an app with async event",
		Action: func(c *cli.Context) error {
			sourcename := c.Args().Get(0)
			route := c.Args().Get(1)
			return action.RegisterService(sourcename, route)
		},
	}
}
