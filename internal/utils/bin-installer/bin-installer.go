package bininstaller

import (
	"fmt"
	"os"
	"provision/internal/utils/commandrunner"
	"strings"
)

var runner commandrunner.ICommandRunner = commandrunner.NewCommandRunner()

func InstallPackage(name, downloadURL string) error {
	if isPackageInstalled(name) {
		println(name + " is already installed")
		return nil
	}

	filename, err := getFilenameFromURL(downloadURL)
	if err != nil {
		return err
	}

	folder, err := getFolderFromFilename(filename)
	if err != nil {
		return err
	}

	if err := download(downloadURL); err != nil {
		return err
	}

	if err := extract(filename); err != nil {
		return err
	}

	if err := symlink(folder, name); err != nil {
		return err
	}

	// TODO: refactor to function
	if name == "node" {
		if err := symlink(folder, "npm"); err != nil {
			return err
		}
		if err := symlink(folder, "npx"); err != nil {
			return err
		}
	}

	if err := cleanup(filename); err != nil {
		return err
	}

	println(name + " installed successfully")
	return nil
}

func isPackageInstalled(name string) bool {
	if err := runner.Run(name, "--version"); err != nil {
		return false
	}
	return true
}

// example URL:
// https://github.com/neovim/neovim/releases/download/v0.11.1/nvim-linux-x86_64.tar.gz
func getFilenameFromURL(url string) (string, error) {
	parts := strings.Split(url, "/")
	if len(parts) == 0 {
		return "", fmt.Errorf("invalid URL: %s", url)
	}
	filename := parts[len(parts)-1]
	return filename, nil
}

// example filename: nvim-linux-x86_64.tar.gz
func getFolderFromFilename(filename string) (string, error) {
	parts := strings.Split(filename, ".")
	if len(parts) < 3 {
		return "", fmt.Errorf("invalid filename: %s", filename)
	}
	folder := strings.Join(parts[:len(parts)-2], ".")
	return folder, nil
}

func download(url string) error {
	println("Downloading", url)
	if err := runner.Run("wget", url); err != nil {
		return fmt.Errorf("failed to download from %s: %w", url, err)
	}
	return nil
}

func extract(filename string) error {
	println("Extracting", filename)
	appsPath := os.Getenv("HOME") + "/apps"
	fmt.Printf("Extracting to: %s\n", appsPath)
	if err := runner.Run("tar", "-xvf", filename, "-C", appsPath); err != nil {
		return fmt.Errorf("error extracting %s: %w", filename, err)
	}
	return nil
}

func symlink(folder, name string) error {
	println("Symlinking", name)
	binary := os.Getenv("HOME") + "/apps/" + folder + "/bin/" + name
	symlinkPath := os.Getenv("HOME") + "/apps/bin/" + name
	if err := runner.Run("ln", "-s", binary, symlinkPath); err != nil {
		return fmt.Errorf("error linking %s: %w", name, err)
	}

	return nil
}

func cleanup(filename string) error {
	println("Cleaning up tar file", filename)
	if err := runner.Run("rm", filename); err != nil {
		return fmt.Errorf("error removing tar %s: %w", filename, err)
	}
	return nil
}
