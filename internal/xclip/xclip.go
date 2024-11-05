package xclip

import "provision/internal/utils/commandrunner"

var runner commandrunner.ICommandRunner = commandrunner.NewCommandRunner()

func SetupXclip() error {
	// TODO: add xclip install, sudo apt install xclip
	println(isXclipInstalled())
	return nil
}

var isXclipInstalled = func() bool {
	if err := runner.Run("xclip", "-version"); err != nil {
		return false
	}
	return true
}

var installXclip = func() error {
	// TODO: impl
	return nil
}
