package apps

import (
	"fmt"
	"os"
)

var appsPath = os.Getenv("HOME") + "/apps/bin"

// Required for neovim install location
func SetupApps() error {
	if !folderExists() {
		if err := createFolder(); err != nil {
			return err
		}
		updateBashRc()
	}
	return nil
}

var folderExists = func() bool {
	info, err := os.Stat(appsPath)
	if os.IsNotExist(err) || !info.IsDir() {
		return false
	}
	return true
}

var updateBashRc = func() error {
	const bashrcString = "export PATH=$PATH:$HOME/apps/bin"
	bashrcPath := os.Getenv("HOME") + "/.bashrc"

	file, err := os.OpenFile(bashrcPath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return fmt.Errorf("failed to open file to append to: %w", err)
	}
	defer file.Close()

	file.WriteString("\n")
	file.WriteString("# Add apps to PATH\n")

	if _, err := file.WriteString(bashrcString); err != nil {
		return fmt.Errorf("failed to write alias to file: %w", err)
	}
	return nil
}

var createFolder = func() error {
	println("Creating apps folder")
	return os.MkdirAll(appsPath, os.ModePerm)
}
