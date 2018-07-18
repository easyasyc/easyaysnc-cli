package commands

import (
	"log"
	"net/http"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudformation"
	"github.com/easyasync/easyaysnc-cli/actions"
	"github.com/easyasync/easyaysnc-cli/source"
	"github.com/larse514/amazonian/cf"
	"github.com/urfave/cli"
)

//RegisterAction interface to define register command actions
type RegisterAction interface {
	RegisterService(c *cli.Context) error
}

type RegisterCommand interface {
	Register() cli.Command
}

//todo- NEED A MUCH BETTER NAME
// type RegisterCommandStruct struct {
// }

//Register is a command to register a route to a service
func Register() cli.Command {

	// Initialize a session in us-west-2 that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-east-1")},
	)

	if err != nil {
		log.Fatal("error creating session")
		os.Exit(127)
	}
	cloudformation := cloudformation.New(sess)

	executor := cf.CFExecutor{Client: cloudformation}

	dispatchSource := source.DispatchSource{Client: &http.Client{},
		URL: "https://dze1hwqe83.execute-api.us-east-1.amazonaws.com/ci/sources/:name/routes"}
	action := actions.HTTPRegisterAction{Executor: executor, Dispatch: dispatchSource}
	return cli.Command{

		Name:    "register",
		Aliases: []string{"r"},
		Usage:   "register an app with async event",
		Action:  action.RegisterService,
	}
}
