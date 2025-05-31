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

func isPackageInstalled(packageName string) bool {
	if err := runner.Run(packageName, "--version"); err != nil {
		return false
	}
	return true
}
