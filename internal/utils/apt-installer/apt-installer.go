package aptinstaller

import "provision/internal/utils/commandrunner"

var runner commandrunner.ICommandRunner = commandrunner.NewCommandRunner()

func InstallPackage(packageName string) error {
	if isPackageInstalled(packageName) {
		println(packageName + " is already installed")
		return nil
	}

	println("Installing " + packageName)
	if err := runner.Run("sudo", "apt", "install", "-y", packageName); err != nil {
		return err
	}
	println(packageName + " installed successfully")
	return nil
}

// TODO: find fix for ripgrep not showing as installed
func isPackageInstalled(packageName string) bool {
	errUsualVersion := runner.Run(packageName, "--version")
	errAlternativeVersion := runner.Run(packageName, "-version")
	if errUsualVersion == nil || errAlternativeVersion == nil {
		return true
	}
	return false
}
