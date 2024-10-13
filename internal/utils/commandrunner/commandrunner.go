package commandrunner

import "os/exec"

type ICommandRunner interface {
	Run(name string, arg ...string) error
}

type CommandRunner struct{}

func (c *CommandRunner) Run(name string, arg ...string) error {
	cmd := exec.Command(name, arg...)
	return cmd.Run()
}

func NewCommandRunner() *CommandRunner {
	return &CommandRunner{}
}
