package ssh

// TODO: Create test file

import (
	"bufio"
	"fmt"
	"os"
	"provision/internal/utils/commandrunner"
)

var publicKey string

var runner commandrunner.ICommandRunner = commandrunner.NewCommandRunner()

func SetupSSH() error {
	publicKeyPath := os.Getenv("HOME") + "/.ssh/id_ed25519.pub"
	privateKeyPath := os.Getenv("HOME") + "/.ssh/id_ed25519"

	if _, err := os.Stat(publicKeyPath); os.IsNotExist(err) {
		email, err := askForEmail()
		if err != nil {
			return err
		}
		if err := generateSSHKey(email, privateKeyPath); err != nil {
			return err
		}
	}

	if err := setPublicKey(publicKeyPath); err != nil {
		return err
	}

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

func generateSSHKey(email string, privateKeyPath string) error {
	fmt.Println("Generating SSH key pair...")
	err := runner.Run("ssh-keygen", "-t", "ed25519", "-C", email, "-f", privateKeyPath, "-N", "")
	if err != nil {
		return fmt.Errorf("failed to generate SSH key: %w", err)
	}
	
	return nil
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
