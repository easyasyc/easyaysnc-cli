package commands

import (
	"github.com/urfave/cli"
)

//Source is the struct representation of a source object
//it contains a name and a route object
type Source struct {
	Name   string  `json:"name"`
	Routes []Route `json:"routes"`
}

//Route structure represeneting JSON object
type Route struct {
	URL string `json:"url"`
}

//ListAction interface to define methos to retrieve sources
type ListAction interface {
	ListSources() error
}

//List is a command to list all routes for each source
func List(action ListAction) cli.Command {
	return cli.Command{

		Name:    "list-all",
		Aliases: []string{"l"},
		Usage:   "list all sources and their subscribed routes",
		Action: func(c *cli.Context) error {
			return action.ListSources()
		},
	}
}
