package neovimreq

import aptinstaller "provision/internal/utils/apt-installer"

var packagesToInstall = []string{
	"unzip",
	"ripgrep",
	"xclip",
	"gcc",
	"fzf",
}

func SetupNeovimRequirements() error {
	for _, packageName := range packagesToInstall {
		if err := aptinstaller.InstallPackage(packageName); err != nil {
			return err
		}
	}
	return nil
}
