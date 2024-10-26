package neovim

import (
	"os"
	"provision/internal/utils/commandrunner"
)

var runner commandrunner.ICommandRunner = commandrunner.NewCommandRunner()

var neovimConfigPath = os.Getenv("HOME") + "/.config/nvim"

func SetupNeovim() error {
	println(isNeovimInstalled())
	return nil
}

var isNeovimInstalled = func() bool {
	info, err := os.Stat(neovimConfigPath)
	if os.IsNotExist(err) || !info.IsDir() {
		return false
	}
	return true
}

// Download link
// https://github.com/neovim/neovim/releases/latest/download/nvim-linux64.tar.gz
var installNeovim = func() error {
	// TODO: Implement
	// Download
	// Extract
	// Put in home/apps/
	// Create symlink from home/apps/nvim-linux64/bin/nvim to home/apps/bin/nvim
	//  Delete the downloaded file

	return nil
}
