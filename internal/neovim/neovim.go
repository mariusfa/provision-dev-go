package neovim

import "provision/internal/utils/commandrunner"

var runner commandrunner.ICommandRunner = commandrunner.NewCommandRunner()

func SetupNeovim() error {
	// TODO: Implement
	return nil
}

var isNeovimInstalled = func() bool {
	// TODO: Implement
	return true
}

var installNeovim = func() error {
	// TODO: Implement
	return nil
}
