package tmuxconfig

import (
	"fmt"
	"os"
	"provision/internal/utils/commandrunner"
)

const githubTmuxConfigURL = "git@github.com:mariusfa/tmux-config.git"

var tmuxConfigPath = os.Getenv("HOME") + "/.config/tmux"
var tmuxConfigFile = tmuxConfigPath + "/.tmux.conf"
var tmuxSymlinkPath = os.Getenv("HOME") + "/.tmux.conf"

var runner commandrunner.ICommandRunner = commandrunner.NewCommandRunner()

func SetupTmuxConfig() error {
	if tmuxConfigExists() {
		println("Tmux config already exists")
		return nil
	}

	if err := cloneTmuxConfig(); err != nil {
		return err
	}
	println("Tmux config cloned")

	if err := symlinkTmuxConfig(); err != nil {
		return err
	}

	return nil
}

var tmuxConfigExists = func() bool {
	info, err := os.Stat(tmuxConfigPath)
	if os.IsNotExist(err) || !info.IsDir() {
		return false
	}
	return true
}

var cloneTmuxConfig = func() error {
	if err := runner.Run("git", "clone", githubTmuxConfigURL, tmuxConfigPath); err != nil {
		return fmt.Errorf("error cloning tmux config: %w", err)
	}
	return nil
}

var symlinkTmuxConfig = func() error {
	println("Symlinking tmux config")
	if err := runner.Run("ln", "-s", tmuxConfigFile, tmuxSymlinkPath); err != nil {
		return fmt.Errorf("error symlinking tmux config: %w", err)
	}
	return nil
}
