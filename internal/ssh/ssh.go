package ssh

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

var publicKey string

func SetupSSH() error {
	email, err := askForEmail()
	if err != nil {
		return err
	}

	publicKeyPath := os.Getenv("HOME") + "/.ssh/id_ed25519.pub"

	if _, err := os.Stat(publicKeyPath); os.IsNotExist(err) {
		// Generate SSH keys using ssh-keygen command
		fmt.Println("Generating SSH key pair...")
		cmd := exec.Command("ssh-keygen", "-t", "ed25519", "-C", email, "-N", "")
		err := cmd.Run()
		if err != nil {
			return fmt.Errorf("failed to generate SSH key: %v", err)
		}
	}

	publicKeyBytes, err := os.ReadFile(publicKeyPath)
	if err != nil {
		return fmt.Errorf("failed to read public key: %v", err)
	}

	publicKey = string(publicKeyBytes)

	fmt.Println("Public SSH Key:")
	fmt.Println(publicKey)

	return nil
}

func askForEmail() (string, error) {
	fmt.Print("Enter your email for SSH key: ")
	reader := bufio.NewReader(os.Stdin)
	email, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("failed to read email: %v", err)
	}
	return email[:len(email)-1], nil // Remove the newline character
}

func GetPublicKey() string {
	return publicKey
}
