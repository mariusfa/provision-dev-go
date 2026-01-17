package brewinstaller

import "provision/internal/utils/commandrunner"

var runner commandrunner.ICommandRunner = commandrunner.NewCommandRunner()

func InstallPackage(packageName string) error {
	if isPackageInstalled(packageName) {
		println(packageName + " is already installed")
		return nil
	}

	println("Installing " + packageName)
	if err := runner.Run("brew", "install", packageName); err != nil {
		return err
	}
	println(packageName + " installed successfully")
	return nil
}

func isPackageInstalled(packageName string) bool {
	command := commandMapper(packageName)
	versionCommand := versionMapper(packageName)
	if err := runner.Run(command, versionCommand); err == nil {
		return true
	}
	return false
}

func commandMapper(packageName string) string {
	if packageName == "ripgrep" {
		return "rg"
	}
	return packageName
}

func versionMapper(packageName string) string {
	switch packageName {
	case "xclip":
		return "-version"
	case "tmux":
		return "-V"
	case "unzip":
		return "-v"
	default:
		return "--version"
	}
}
