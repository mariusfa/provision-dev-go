package neovim

import (
	"os"
	"provision/internal/utils/commandrunner"
)

var runner commandrunner.ICommandRunner = commandrunner.NewCommandRunner()

var neovimPath = os.Getenv("HOME") + "/.config/nvim"

func SetupNeovim() error {
	println("Checking neovim folder")
	println(isNeovimInstalled())
	return nil
}

var isNeovimInstalled = func() bool {
	info, err := os.Stat(neovimPath)
	if os.IsNotExist(err) || !info.IsDir() {
		return false
	}
	return true
}

var installNeovim = func() error {
	// TODO: Implement
	return nil
}
