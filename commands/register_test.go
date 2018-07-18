package commands

import (
	"testing"

	"github.com/urfave/cli"
)

type mockRegisterAction struct {
}

func (mock mockRegisterAction) RegisterService(c *cli.Context) error {
	return nil
}

func TestRegisterNameIsCorrectValue(t *testing.T) {

	command := Register()

	expected := "register"

	actual := command.Name

	if actual != expected {
		t.Log("expected ", expected, " got ", actual)
		t.Fail()
	}

}
