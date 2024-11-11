package gcc

import "provision/internal/utils/commandrunner"

var runner commandrunner.ICommandRunner = commandrunner.NewCommandRunner()

func SetupGcc() error {
	if isGccInstalled() {
		println("gcc is already installed")
		return nil
	}

	if err := installGcc(); err != nil {
		return err
	}

	println("gcc installed")
	return nil
}

var isGccInstalled = func() bool {
	if err := runner.Run("gcc", "--version"); err != nil {
		return false
	}
	return true
}

var installGcc = func() error {
	if err := runner.Run("sudo", "apt", "install", "gcc", "-y"); err != nil {
		return err
	}
	return nil
}
