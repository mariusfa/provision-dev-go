package neovim

import (
	"fmt"
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
		return fmt.Errorf("error downloading neovim: %w", err)
	}
	appsFolder := os.Getenv("HOME") + "/apps"
	if err := runner.Run("tar", "-xvf", "nvim-linux64.tar.gz", "-C", appsFolder); err != nil {
		return fmt.Errorf("error extracting neovim: %w", err)
	}
	if err := runner.Run("ln", "-s", "~/apps/nvim-linux64/bin/nvim", "~/apps/bin/nvim"); err != nil {
		return fmt.Errorf("error linking neovim: %w", err)
	}
	if err := runner.Run("rm", "nvim-linux64.tar.gz"); err != nil {
		return fmt.Errorf("error removing neovim tar: %w", err)
	}
	return nil
}
