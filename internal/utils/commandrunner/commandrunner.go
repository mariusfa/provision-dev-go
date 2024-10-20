package commandrunner

import "os/exec"

type ICommandRunner interface {
	Run(name string, arg ...string) error
	RunWithOutput(name string, arg ...string) (string, error)
}

type CommandRunner struct{}

func (c *CommandRunner) Run(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	return cmd.Run()
}

func (c *CommandRunner) RunWithOutput(name string, arg ...string) (string, error) {
	cmd := exec.Command(name, arg...)
	byteOutput, err := cmd.Output()
	if err != nil {
		return "", err
	}
	return string(byteOutput), nil
}

func NewCommandRunner() *CommandRunner {
	return &CommandRunner{}
}
