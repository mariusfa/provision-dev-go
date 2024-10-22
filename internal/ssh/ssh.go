package ssh

import (
	"bufio"
	"fmt"
	"os"
	"provision/internal/utils/commandrunner"
)

var publicKey string

var (
	runner         commandrunner.ICommandRunner = commandrunner.NewCommandRunner()
	publicKeyPath                               = os.Getenv("HOME") + "/.ssh/id_ed25519.pub"
	privateKeyPath                              = os.Getenv("HOME") + "/.ssh/id_ed25519"
)

func SetupSSH() error {
	if _, err := os.Stat(publicKeyPath); os.IsNotExist(err) {
		email, err := askForEmail()
		if err != nil {
			return err
		}
		if err := generateSshKey(email); err != nil {
			return err
		}
	}

	if err := setPublicKey(publicKeyPath); err != nil {
		return err
	}

	fmt.Println("Public SSH Key:")
	fmt.Println(publicKey)
	enterToContinue()

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

func generateSshKey(email string) error {
	fmt.Println("Generating SSH key pair...")
	err := runner.Run("ssh-keygen", "-t", "ed25519", "-C", email, "-f", privateKeyPath, "-N", "")
	if err != nil {
		return fmt.Errorf("failed to generate SSH key: %w", err)
	}

	return nil
}

func enterToContinue() {
	fmt.Print("Copy ssh key and setup in github before continuing. Press Enter to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func setPublicKey(publicKeyPath string) error {
	publicKeyBytes, err := os.ReadFile(publicKeyPath)
	if err != nil {
		return fmt.Errorf("failed to read public key: %w", err)
	}

	publicKey = string(publicKeyBytes)
	return nil
}

func GetPublicKey() string {
	return publicKey
}
