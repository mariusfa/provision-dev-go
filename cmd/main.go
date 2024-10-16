package main

import (
	"fmt"
	"provision/internal/ssh"
	"provision/internal/utils/cliparser"
)

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
	err := ssh.SetupSSH()
	if err != nil {
		fmt.Printf("Error setting up SSH: %v\n", err)
		return
	}
}

func runSsh() {
	println("Running SSH")
	err := ssh.SetupSSH()
	if err != nil {
		fmt.Printf("Error setting up SSH: %v\n", err)
		return
	}
}

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
	default:
		fmt.Println("Invalid option")
	}

	printRememberSourceProfile()
}
