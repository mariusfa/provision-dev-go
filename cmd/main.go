package main

import (
	"fmt"
	"log"
	"provision/internal/alias"
	"provision/internal/git"
	"provision/internal/ssh"
	"provision/internal/utils/cliparser"
)

// TODO: create a custom print util to print with colors and icons

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
	case cliparser.ALIAS:
		runAlias()
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
}

func runAll() {
	println("Running all")
	runSsh()
	runGit()
	runAlias()
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

func runAlias() {
	println("Running Alias setup")
	err := alias.SetupAlias()
	if err != nil {
		log.Fatalf("Error setting up aliases: %v\n", err)
	}
}
