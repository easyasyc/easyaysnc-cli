package commands

import (
	"testing"
)

type mockRegisterAction struct {
}

func (mock mockRegisterAction) RegisterService(sourcename string, route string) error {
	return nil
}

func TestRegisterNameIsCorrectValue(t *testing.T) {
	action := mockRegisterAction{}
	command := Register(action)

	expected := "register"

	actual := command.Name

	if actual != expected {
		t.Log("expected ", expected, " got ", actual)
		t.Fail()
	}

}
