package actions

import (
	"log"
	"os"

	"github.com/easyasync/easyaysnc-cli/commands"
)

//SourceStruct TODO- to be updated with better naming
type SourceStruct struct {
	ListSource ListSource
}

//ListSource iterface defines operation to return slice of sources
type ListSource interface {
	GetSources() ([]commands.Source, error)
}

//listSources method to retrieve sources
func (source SourceStruct) listSources() ([]commands.Source, error) {
	sources, err := source.ListSource.GetSources()
	if err != nil {
		log.Println("ERROR: could not get source")
		os.Exit(1)
	}
	return sources, err
}

//ListSources method to retrieve sources
func (source SourceStruct) ListSources() error {
	sources, _ := source.listSources()

	for _, s := range sources {
		routesString := ""
		for _, r := range s.Routes {
			routesString = routesString + r.URL + ", "
		}

		log.Println("Source: ", s.Name, " Routes: ", routesString)
	}

	return nil
}

//NewSourceStruct constructor to create SourceStruct
func NewSourceStruct(listSource ListSource) *SourceStruct {
	return &SourceStruct{listSource}
}
