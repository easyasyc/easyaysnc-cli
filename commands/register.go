package commands

import (
	"github.com/easyasync/easyaysnc-cli/actions"
	"github.com/urfave/cli"
)

//RegisterAction interface to define register command actions
type RegisterAction interface {
	RegisterService(c *cli.Context) error
}

//Register is a command to register a route to a service
func Register() cli.Command {
	action := actions.HTTPRegisterAction{}
	return cli.Command{

		Name:    "register",
		Aliases: []string{"r"},
		Usage:   "register an app with async event",
		Action:  action.RegisterService,
	}
}
