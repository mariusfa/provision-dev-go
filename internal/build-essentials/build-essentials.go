package buildessentials

import "provision/internal/utils/commandrunner"

var runner commandrunner.ICommandRunner = commandrunner.NewCommandRunner()

func SetupBuildEssentials() error {
	// TODO: impl
	println(isBuildEssentialsInstalled())
	return nil
}

var isBuildEssentialsInstalled = func() bool {
	if err := runner.Run("gcc", "--version"); err != nil {
		return false
	}
	return true
}

var installBuildEssentials = func() error {
	// TODO: impl
	return nil
}
