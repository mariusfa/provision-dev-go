package neovimconfig

import (
	"os"
	"provision/internal/utils/commandrunner"
)

const githubNeovimConfigURL = "git@github.com:mariusfa/lzy-vim-starter.git"

var neovimConfigPath = os.Getenv("HOME") + "/.config/nvim"

var runner commandrunner.ICommandRunner = commandrunner.NewCommandRunner()

func SetupNeovimConfig() error {
	// TODO: Implement
	return nil
}

var neovimConfigExists = func() bool {
	info, err := os.Stat(neovimConfigPath)
	if os.IsNotExist(err) || !info.IsDir() {
		return false
	}
	return true
}

var cloneNeovimConfig = func() error {
	// TODO: Implement
	return nil
}
