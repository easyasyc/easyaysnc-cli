package commands

import (
	"testing"
)

type mockListAction struct {
}

func (mock mockListAction) ListSources() error {
	return nil
}
func TestSourceStructListSourcesReturnsSource(t *testing.T) {
	action := mockListAction{}
	command := List(action)

	if command.Name != "list-all" {
		t.Log("Name was incorrect")
		t.Fail()
	}

}
