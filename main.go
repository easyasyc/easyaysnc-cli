package main

import (
	"log"
	"os"

	"github.com/easyasync/easyaysnc-cli/config"
)

func main() {
	app := config.CreateApp()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
