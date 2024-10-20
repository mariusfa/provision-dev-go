package userinput

// Tests not needed because it is just a wrapper

import (
	"bufio"
	"fmt"
	"os"
)

func AskForInput(prompt string) (string, error) {
	fmt.Print(prompt)
	return readInput()
}

func readInput() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", fmt.Errorf("failed to read input : %w", err)
	}
	return input[:len(input)-1], nil // Remove the newline character
}
