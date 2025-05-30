package node

import (
	"fmt"
	"os"
	"provision/internal/utils/commandrunner"
)

var runner commandrunner.ICommandRunner = commandrunner.NewCommandRunner()

func SetupNode() error {
	if isNodeInstalled() {
		println("Nodejs is already installed")
		return nil
	}

	if err := installNode(); err != nil {
		return err
	}

	println("Nodejs installed")

	return nil
}

func download() error {
	println("Downloading nodejs")
	if err := runner.Run("wget", "https://nodejs.org/dist/v22.16.0/node-v22.16.0-linux-x64.tar.gz"); err != nil {
		return fmt.Errorf("error downloading nodejs: %w", err)
	}
	return nil
}

func extract() error {
	println("Extracting nodejs")

	appsPath := os.Getenv("HOME") + "/apps"
	fmt.Printf("Extracting to nodejs to: %s\n", appsPath)
	if err := runner.Run("tar", "-xvf", "node-v22.16.0-linux-x64.tar.gz", "-C", appsPath); err != nil {
		return fmt.Errorf("error extracting nodejs: %w", err)
	}
	return nil
}

// TODO: Refactor to remove dups of code.
func symlink() error {
	println("Symlinking node")
	nodeBinary := os.Getenv("HOME") + "/apps/node-v22.16.0-linux-x64/bin/node"
	nodeSymlinkPath := os.Getenv("HOME") + "/apps/bin/node"
	if err := runner.Run("ln", "-s", nodeBinary, nodeSymlinkPath); err != nil {
		return fmt.Errorf("error linking node: %w", err)
	}

	println("Symlinking npm")
	npmBinary := os.Getenv("HOME") + "/apps/node-v22.16.0-linux-x64/bin/npm"
	npmSymlinkPath := os.Getenv("HOME") + "/apps/bin/npm"
	if err := runner.Run("ln", "-s", npmBinary, npmSymlinkPath); err != nil {
		return fmt.Errorf("error linking npm: %w", err)
	}

	println("Symlinking npx")
	npxBinary := os.Getenv("HOME") + "/apps/node-v22.16.0-linux-x64/bin/npx"
	npxSymlinkPath := os.Getenv("HOME") + "/apps/bin/npx"
	if err := runner.Run("ln", "-s", npxBinary, npxSymlinkPath); err != nil {
		return fmt.Errorf("error linking npx: %w", err)
	}
	return nil
}

func cleanup() error {
	println("Cleaning up node tar file")
	if err := runner.Run("rm", "node-v22.16.0-linux-x64.tar.gz"); err != nil {
		return fmt.Errorf("error removing neovim tar: %w", err)
	}
	return nil
}

var isNodeInstalled = func() bool {
	if err := runner.Run("node", "--version"); err != nil {
		return false
	}
	return true
}

var installNode = func() error {
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

// https://nodejs.org/dist/v22.16.0/node-v22.16.0-linux-x64.tar.gz
// check bin folder for node, npm, npx
// npm symlinks to lib/node_modules/npm/bin/npm-cli.js*
// npx symlinks to lib/node_modules/npm/bin/npx-cli.js*
// They work is symlink as with node
