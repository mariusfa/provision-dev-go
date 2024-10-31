package neovimconfig

import "os"

const githubNeovimConfigURL = "git@github.com:mariusfa/lzy-vim-starter.git"

var neovimConfigPath = os.Getenv("HOME") + "/.config/nvim"

func SetupNeovimConfig() error {
	// TODO: Implement
	return nil
}

var neovimConfigExists = func() bool {
	// TODO: Implement
	return false
}

var cloneNeovimConfig = func() error {
	// TODO: Implement
	return nil
}
