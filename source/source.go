package source

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"
)

//DispatchSource struct implementing Dispatch interface
type DispatchSource struct {
	Client *http.Client
	URL    string
}

//RouteDTO struct representing json
type RouteDTO struct {
	Route              Route `json:"route"`
	WithSourceCreation bool  `json:"withSourceCreation"`
}

//Route struct representing route json
type Route struct {
	URL string `json:"url"`
}

//CreateSource method to bind a route to a source
func (dispatch DispatchSource) CreateSource(sourcename string, route string) error {
	url := strings.Replace(dispatch.URL, ":name", sourcename, -1)
	log.Println("DEBUG: created url ", url)
	request := RouteDTO{Route: Route{URL: route}, WithSourceCreation: true}

	out, err := json.Marshal(request)

	if err != nil {
		return err
	}
	jsonString := string(out)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(jsonString)))
	req.Header.Set("Content-Type", "application/json")

	if err != nil {
		return err
	}

	resp, err := dispatch.Client.Do(req)
	if err != nil {
		return err
	}
	if !is2xx(&resp.StatusCode) {
		log.Println("client returned error ", resp)
		os.Exit(127)
	}

	return nil

}

func is2xx(status *int) bool {
	switch *status {
	case 200:
		return true
	case 201:
		return true
	case 202:
		return true
	default:
		return false
	}
}
