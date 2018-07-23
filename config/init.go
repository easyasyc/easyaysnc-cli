package config

import (
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/easyasync/easyaysnc-cli/actions"
	"github.com/easyasync/easyaysnc-cli/commands"
	"github.com/easyasync/easyaysnc-cli/source"
	"github.com/larse514/amazonian/cf"
	"github.com/urfave/cli"
)

//CreateApp method to create initial app
func CreateApp() *cli.App {
	app := cli.NewApp()
	// Initialize a session in us-west-2 that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)

	if err != nil {
		log.Fatal("error creating session")
		os.Exit(1)
	}
	cloudformation := cloudformation.New(sess)

	executor := cf.CFExecutor{Client: cloudformation}
	httpClient := http.Client{}
	dispatchSource := source.DispatchSource{&httpClient, "https://dze1hwqe83.execute-api.us-east-1.amazonaws.com/ci/sources/:name/routes"}
	listSource := source.NewList(&httpClient, "https://dze1hwqe83.execute-api.us-east-1.amazonaws.com/ci/sources")

	registerAction := actions.NewRegisterAction(dispatchSource, executor)
	listAction := actions.NewSourceStruct(listSource)

	app.Name = "easyasync"
	app.Usage = "Command line interface for easyasync cli"
	app.Version = "0.0.1"
	app.Author = "VSS"
	app.Commands = []cli.Command{
		commands.Register(registerAction),
		commands.List(listAction),
	}

	return app
}
