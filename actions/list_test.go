package actions

import (
	"errors"
	"os"
	"os/exec"
	"reflect"
	"testing"

	"github.com/easyasync/easyaysnc-cli/commands"
)

type mockListSource struct {
}

func (mock mockListSource) GetSources() ([]commands.Source, error) {
	return make([]commands.Source, 0), nil
}

type mockBadListSource struct {
}

func (mock mockBadListSource) GetSources() ([]commands.Source, error) {
	return make([]commands.Source, 0), errors.New("")
}
func TestSourceStructConstructorContainsListSource(t *testing.T) {
	actual := NewSourceStruct(mockListSource{}).ListSource
	if actual == nil {
		t.Log("ListSource is nil")
		t.Fail()
	}

}

func TestSourceStructListSourcesReturnsNoError(t *testing.T) {
	sources := mockListSource{}

	err := NewSourceStruct(sources).ListSources()
	if err != nil {
		t.Fatalf("error received when none expected %e", err)
	}

}

func TestSourceStructListSourcesReturnsSource(t *testing.T) {
	sources := mockListSource{}
	expected, err := sources.GetSources()
	if err != nil {
		t.Fatalf("error received when none expected %e", err)
	}
	actual, err := NewSourceStruct(sources).listSources()
	if err != nil {
		t.Fatalf("error received when none expected %e", err)
	}
	if !reflect.DeepEqual(actual, expected) {
		t.Log("ListSource is nil")
		t.Fail()
	}

}

func TestSourceStructListSourcesReturnsErrorProcessExitWithCode1(t *testing.T) {
	source := NewSourceStruct(mockBadListSource{})

	if os.Getenv("BE_CRASHER") == "1" {
		source.ListSources()
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestSourceStructListSourcesReturnsErrorProcessExitWithCode1")
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process ran with err %v, want exit status 1", err)

}
