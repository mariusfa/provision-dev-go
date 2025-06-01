package golang

import bininstaller "provision/internal/utils/bin-installer"

func SetupGo() error {
	if err := bininstaller.InstallPackage("go", "https://go.dev/dl/go1.24.3.linux-amd64.tar.gz"); err != nil {
		return err
	}
	return nil
}
