package ripgrep

import "provision/internal/utils/commandrunner"

var runner commandrunner.ICommandRunner = commandrunner.NewCommandRunner()

func SetupRipgrep() error {
	if isRipgrepInstalled() {
		println("ripgrep is already installed")
		return nil
	}

	if err := installRipgrep(); err != nil {
		return err
	}

	println("ripgrep installed")
	return nil
}

var isRipgrepInstalled = func() bool {
	if err := runner.Run("rgrep", "--version"); err != nil {
		return false
	}
	return true
}

var installRipgrep = func() error {
	if err := runner.Run("sudo", "apt", "install", "ripgrep", "-y"); err != nil {
		return err
	}
	return nil
}
