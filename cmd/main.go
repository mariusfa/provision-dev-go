package main

import (
	"fmt"
	"log"
	"provision/internal/alias"
	"provision/internal/apps"
	"provision/internal/gcc"
	"provision/internal/git"
	"provision/internal/neovim"
	neovimconfig "provision/internal/neovim-config"
	"provision/internal/ripgrep"
	"provision/internal/ssh"
	"provision/internal/utils/cliparser"
	"provision/internal/xclip"
)

// TODO: create a custom print util to print with colors and icons
// TODO: add fzf install, sudo apt install fzf
// TODO: add help command to print all sub commands

func main() {
	printInitializationBanner()

	cliOptions, err := cliparser.CliParser()
	if err != nil {
		fmt.Printf("Error parsing CLI: %v\n", err)
		return
	}

	switch cliOptions {
	case cliparser.ALL:
		runAll()
	case cliparser.XCLIP:
		runXclip()
	case cliparser.RIPGREP:
		runRipgrep()
	case cliparser.GCC:
		runGcc()
	case cliparser.SSH:
		runSsh()
	case cliparser.GIT:
		runGit()
	case cliparser.APPS:
		runApps()
	case cliparser.ALIAS:
		runAlias()
	case cliparser.NEOVIM:
		runNeovim()
	case cliparser.NEOVIM_CONFIG:
		runNeovimConfig()
	default:
		fmt.Println("Invalid option")
	}

	printRememberSourceProfile()
}

func printInitializationBanner() {
	border := "============================"
	message := "Go provision-dev initializing"

	fmt.Println(border)
	fmt.Println(message)
	fmt.Println(border)
}

func printRememberSourceProfile() {
	fmt.Println("Remember to source your profile: source ~/.profile")
	fmt.Println("Remember to source your bashrc: source ~/.bashrc")
}

func runAll() {
	println("Running all")
	runXclip()
	runRipgrep()
	runGcc()
	runSsh()
	runGit()
	runApps()
	runAlias()
	runNeovim()
	runNeovimConfig()
}

func runXclip() {
	println("Running Xclip setup")
	if err := xclip.SetupXclip(); err != nil {
		log.Fatalf("Error setting up Xclip: %v\n", err)
	}
}

func runRipgrep() {
	println("Running ripgrep setup")
	if err := ripgrep.SetupRipgrep(); err != nil {
		log.Fatalf("Error setting up ripgrep: %v\n", err)
	}
}

func runGcc() {
	println("Running Gcc setup")
	if err := gcc.SetupGcc(); err != nil {
		log.Fatalf("Error setting up Gcc: %v\n", err)
	}
}

func runSsh() {
	println("Running SSH setup")
	err := ssh.SetupSSH()
	if err != nil {
		log.Fatalf("Error setting up SSH: %v\n", err)
	}
}

func runGit() {
	println("Running Git setup")
	err := git.SetupGit()
	if err != nil {
		log.Fatalf("Error setting up Git: %v\n", err)
	}
}

func runApps() {
	println("Running Apps setup")
	err := apps.SetupApps()
	if err != nil {
		log.Fatalf("Error setting up Apps: %v\n", err)
	}
}

func runAlias() {
	println("Running Alias setup")
	err := alias.SetupAlias()
	if err != nil {
		log.Fatalf("Error setting up aliases: %v\n", err)
	}
}

func runNeovim() {
	println("Running Neovim setup")
	err := neovim.SetupNeovim()
	if err != nil {
		log.Fatalf("Error setting up Neovim: %v\n", err)
	}
}

func runNeovimConfig() {
	println("Running Neovim config setup")
	err := neovimconfig.SetupNeovimConfig()
	if err != nil {
		log.Fatalf("Error setting up Neovim config: %v\n", err)
	}
}
