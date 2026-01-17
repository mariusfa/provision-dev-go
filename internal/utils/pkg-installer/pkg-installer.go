package pkginstaller

import (
	"fmt"
	"os/exec"

	aptinstaller "provision/internal/utils/apt-installer"
	brewinstaller "provision/internal/utils/brew-installer"
)

type PackageManager int

const (
	APT PackageManager = iota
	BREW
)

var detectedManager PackageManager

func init() {
	detectedManager = detectPackageManager()
}

func detectPackageManager() PackageManager {
	if _, err := exec.LookPath("apt"); err == nil {
		return APT
	}
	if _, err := exec.LookPath("brew"); err == nil {
		return BREW
	}
	return APT // default fallback
}

func InstallPackage(packageName string) error {
	switch detectedManager {
	case APT:
		return aptinstaller.InstallPackage(packageName)
	case BREW:
		return brewinstaller.InstallPackage(packageName)
	default:
		return fmt.Errorf("no supported package manager found")
	}
}
