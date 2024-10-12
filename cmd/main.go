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

func main() {
    printInitializationBanner()

	err := ssh.SetupSSH()
	if err != nil {
		fmt.Printf("Error setting up SSH: %v\n", err)
		return
	}
}
