package actions

import (
	"fmt"

	"github.com/urfave/cli"
)

//HTTPRegisterAction Implementation of RegisterAction
type HTTPRegisterAction struct {
}

//RegisterService is a method to register a service to easyasync
func (action HTTPRegisterAction) RegisterService(c *cli.Context) error {
	fmt.Printf("Service is %q\n", c.Args().Get(0))
	return nil
}
