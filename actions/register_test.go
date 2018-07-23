package actions

import (
	"errors"
	"os"
	"os/exec"
	"testing"

	"github.com/aws/aws-sdk-go/service/cloudformation"
)

const (
	sourcename = "name"
	resource   = "resource"
)

type mockDispatch struct {
}

func (mock mockDispatch) CreateSource(sourcename string, route string) error {
	return nil
}

type mockBadDispatch struct {
}

func (mock mockBadDispatch) CreateSource(sourcename string, route string) error {
	return errors.New("errored")
}

type mockExecutor struct {
}

func (mock mockExecutor) CreateStack(templateBody string, stackName string, parameters []*cloudformation.Parameter) error {
	return nil
}
func (mock mockExecutor) UpdateStack(templateBody string, stackName string, parameters []*cloudformation.Parameter) error {
	return nil
}

func (mock mockExecutor) PauseUntilCreateFinished(stackName string) error {
	return nil
}
func (mock mockExecutor) PauseUntilUpdateFinished(stackName string) error {
	return nil
}

type mockBadExecutor struct {
}

func (mock mockBadExecutor) CreateStack(templateBody string, stackName string, parameters []*cloudformation.Parameter) error {
	return nil
}
func (mock mockBadExecutor) UpdateStack(templateBody string, stackName string, parameters []*cloudformation.Parameter) error {
	return nil
}

func (mock mockBadExecutor) PauseUntilCreateFinished(stackName string) error {
	return errors.New("ERROR")
}
func (mock mockBadExecutor) PauseUntilUpdateFinished(stackName string) error {
	return nil
}
func TestNewRegisterActionDispatcherNotNil(t *testing.T) {
	actual := NewRegisterAction(mockDispatch{}, mockExecutor{})
	if actual.Dispatch == nil {
		t.Log("Dispatcher not set")
		t.Fail()
	}
}

func TestNewRegisterActionExecutorNotNil(t *testing.T) {
	actual := NewRegisterAction(mockDispatch{}, mockExecutor{})
	if actual.Executor == nil {
		t.Log("Executor not set")
		t.Fail()
	}
}

func TestSourceFailsProcessExitsWithStatus1(t *testing.T) {
	actual := NewRegisterAction(mockBadDispatch{}, mockExecutor{})
	if os.Getenv("BE_CRASHER") == "1" {
		actual.RegisterService(sourcename, resource)
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestSourceFailsProcessExitsWithStatus1")
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process ran with err %v, want exit status 1", err)
}

func TestExecutorFailsProcessExitsWithStatus1(t *testing.T) {
	actual := NewRegisterAction(mockDispatch{}, mockBadExecutor{})
	if os.Getenv("BE_CRASHER") == "1" {
		actual.RegisterService(sourcename, resource)
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestExecutorFailsProcessExitsWithStatus1")
	cmd.Env = append(os.Environ(), "BE_CRASHER=1")
	err := cmd.Run()
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("process ran with err %v, want exit status 1", err)
}
