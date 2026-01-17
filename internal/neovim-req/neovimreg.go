package neovimreq

import pkginstaller "provision/internal/utils/pkg-installer"

var packagesToInstall = []string{
	"unzip",
	"ripgrep",
	"xclip",
	"gcc",
	"fzf",
	"tmux",
}

func SetupNeovimRequirements() error {
	for _, packageName := range packagesToInstall {
		if err := pkginstaller.InstallPackage(packageName); err != nil {
			return err
		}
	}
	return nil
}
