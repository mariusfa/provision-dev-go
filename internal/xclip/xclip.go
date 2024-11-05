package xclip

import "provision/internal/utils/commandrunner"

var runner commandrunner.ICommandRunner = commandrunner.NewCommandRunner()

func SetupXclip() error {
	if isXclipInstalled() {
		println("xclip is already installed")
		return nil
	}

	if err := installXclip(); err != nil {
		return err
	}

	println("xclip installed")
	return nil
}

var isXclipInstalled = func() bool {
	if err := runner.Run("xclip", "-version"); err != nil {
		return false
	}
	return true
}

var installXclip = func() error {
	if err := runner.Run("apt", "install", "xclip"); err != nil {
		return err
	}
	return nil
}
