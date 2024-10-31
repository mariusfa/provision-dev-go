package neovimconfig

import (
	"fmt"
	"os"
	"provision/internal/utils/commandrunner"
)

const githubNeovimConfigURL = "git@github.com:mariusfa/lzy-vim-starter.git"

var neovimConfigPath = os.Getenv("HOME") + "/.config/nvim"

var runner commandrunner.ICommandRunner = commandrunner.NewCommandRunner()

func SetupNeovimConfig() error {
	if neovimConfigExists() {
		println("Neovim config already exists")
		return nil
	}

	if err := cloneNeovimConfig(); err != nil {
		return err
	}

	println("Neovim config cloned")
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
	if err := runner.Run("git", "clone", githubNeovimConfigURL, neovimConfigPath); err != nil {
		return fmt.Errorf("error cloning neovim config: %w", err)
	}
	return nil
}
