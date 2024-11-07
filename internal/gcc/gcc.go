package gcc

import "provision/internal/utils/commandrunner"

var runner commandrunner.ICommandRunner = commandrunner.NewCommandRunner()

func SetupGcc() error {
	// TODO: impl
	println(isGccInstalled())
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
