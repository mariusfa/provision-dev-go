package main

import (
	"fmt"
	"provision/internal/ssh"
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

// TODO: check for command line argument if to run only one package. For instance only run ssh package
func main() {
	printInitializationBanner()

	err := ssh.SetupSSH()
	if err != nil {
		fmt.Printf("Error setting up SSH: %v\n", err)
		return
	}

	printRememberSourceProfile()
}
