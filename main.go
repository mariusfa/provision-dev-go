package main

import (
	"fmt"
	"log"
	"provision/internal/alias"
	"provision/internal/apps"
	"provision/internal/git"
	"provision/internal/neovim"
	neovimconfig "provision/internal/neovim-config"
	neovimreq "provision/internal/neovim-req"
	"provision/internal/node"
	"provision/internal/ssh"
	"provision/internal/utils/cliparser"
)

// TODO: create a custom print util to print with colors and icons
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
	case cliparser.NODE:
		runNode()
	case cliparser.NEOVIM_REQ:
		runNeovimRequirements()
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
	runNeovimRequirements()
	runNode()
	runSsh()
	runGit()
	runApps()
	runAlias()
	runNeovim()
	runNeovimConfig()
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

func runNode() {
	println("Running Nodejs setup")
	err := node.SetupNode()
	if err != nil {
		log.Fatalf("Error setting up Nodejs: %v\n", err)
	}
}

func runNeovimRequirements() {
	println("Running Neovim requirements and other packages setup")
	err := neovimreq.SetupNeovimRequirements()
	if err != nil {
		log.Fatalf("Error setting up Neovim requirements and other packages: %v\n", err)
	}
}
