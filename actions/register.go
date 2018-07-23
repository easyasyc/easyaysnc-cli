package actions

import (
	"fmt"
	"os"

	"github.com/larse514/amazonian/cf"
)

//HTTPRegisterAction Implementation of RegisterAction
type HTTPRegisterAction struct {
	Dispatch Dispatch
	Executor cf.Executor
}

//Dispatch interface to define methods to enable dispatching
type Dispatch interface {
	CreateSource(sourcename string, route string) error
}

//RegisterService is a method to register a service to easyasync
func (action HTTPRegisterAction) RegisterService(sourcename string, route string) error {

	fmt.Printf("Registering source %q to route %q... \n", sourcename, route)

	//create the source
	err := action.Dispatch.CreateSource(sourcename, route)

	if err != nil {
		fmt.Println("ERROR: could not create source ", err)
		os.Exit(1)
	}

	err = action.Executor.PauseUntilCreateFinished(sourcename)
	if err != nil {
		fmt.Println("ERROR: could not create source ", err)
		os.Exit(1)
	}

	return nil
}

//NewRegisterAction constructor to create HTTPRegisterAction
func NewRegisterAction(dispatch Dispatch, executor cf.Executor) *HTTPRegisterAction {
	return &HTTPRegisterAction{dispatch, executor}
}
