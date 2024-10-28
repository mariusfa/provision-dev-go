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
	if err := download(); err != nil {
		return err
	}
	if err := extract(); err != nil {
		return err
	}
	if err := symlink(); err != nil {
		return err
	}
	if err := cleanup(); err != nil {
		return err
	}
	return nil
}

func download() error {
	println("Downloading neovim")
	if err := runner.Run("wget", "https://github.com/neovim/neovim/releases/latest/download/nvim-linux64.tar.gz"); err != nil {
		return fmt.Errorf("error downloading neovim: %w", err)
	}
	return nil
}

func extract() error {
	println("Extracting neovim")
	if err := runner.Run("tar", "-xvf", "nvim-linux64.tar.gz", "-C", os.Getenv("HOME")+"/apps"); err != nil {
		return fmt.Errorf("error extracting neovim: %w", err)
	}
	return nil
}

func symlink() error {
	println("Symlinking neovim")
	neovimBinary := os.Getenv("HOME") + "/apps/nvim-linux64/bin/nvim"
	symlinkPath := os.Getenv("HOME") + "/apps/bin/nvim"
	if err := runner.Run("ln", "-s", neovimBinary, symlinkPath); err != nil {
		return fmt.Errorf("error linking neovim: %w", err)
	}
	return nil
}

func cleanup() error {
	println("Cleaning up neovim tar file")
	if err := runner.Run("rm", "nvim-linux64.tar.gz"); err != nil {
		return fmt.Errorf("error removing neovim tar: %w", err)
	}
	return nil
}
