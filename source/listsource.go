package source

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/easyasync/easyaysnc-cli/commands"
)

//TransferObject structure represeneting JSON object
type TransferObject struct {
	Sources []commands.Source `json:"sources"`
}

//List implementation of ListSource
type List struct {
	Client *http.Client
	URL    string
}

//GetSources returns a list of Sources
func (list List) GetSources() ([]commands.Source, error) {

	log.Println("DEBUG: about to send request to url ", list.URL)
	req, err := http.NewRequest(http.MethodGet, list.URL, nil)
	req.Header.Set("Content-Type", "application/json")

	resp, err := list.Client.Do(req)

	if err != nil {
		log.Println("ERROR: error making http call to source with error", err)
		return nil, errors.New("error making http call")
	}

	defer resp.Body.Close()
	s := TransferObject{}
	err = json.NewDecoder(resp.Body).Decode(&s)

	if err != nil {
		log.Println("ERROR: error unmarshalling response ", err)
		return nil, errors.New("Unmarshal error")
	}

	return s.Sources, nil

}

//NewList constructor for List
func NewList(client *http.Client, url string) *List {
	return &List{client, url}
}
