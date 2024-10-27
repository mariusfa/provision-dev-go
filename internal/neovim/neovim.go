package neovim

import (
	"os"
	"provision/internal/utils/commandrunner"
)

var runner commandrunner.ICommandRunner = commandrunner.NewCommandRunner()

var neovimConfigPath = os.Getenv("HOME") + "/.config/nvim"

func SetupNeovim() error {
	if isNeovimInstalled() {
		println("Neovim is already installed")
	}

	if err := installNeovim(); err != nil {
		return err
	}

	println("Neovim installed")

	return nil
}

// TODO: Change to check for bin file in apps folder
var isNeovimInstalled = func() bool {
	info, err := os.Stat(neovimConfigPath)
	if os.IsNotExist(err) || !info.IsDir() {
		return false
	}
	return true
}

var installNeovim = func() error {
	if err := runner.Run("wget", "https://github.com/neovim/neovim/releases/latest/download/nvim-linux64.tar.gz"); err != nil {
		return err
	}
	if err := runner.Run("tar -xvf nvim-linux64.tar.gz -C ~/apps"); err != nil {
		return err
	}
	if err := runner.Run("ln -s ~/apps/nvim-linux64/bin/nvim ~/apps/bin/nvim"); err != nil {
		return err
	}
	if err := runner.Run("rm nvim-linux64.tar.gz"); err != nil {
		return err
	}
	return nil
}
