package git

import (
	"provision/internal/utils/commandrunner"
	"provision/internal/utils/userinput"
)

var runner commandrunner.ICommandRunner = commandrunner.NewCommandRunner()

// Assume git is installed. Hence ubuntu
func SetupGit() error {
	isSet := isGitConfigSet()

	if isSet {
		println("Git config already set")
		return nil
	}

	username, err := askForUsername()
	if err != nil {
		return err
	}

	email, err := askForEmail()
	if err != nil {
		return err
	}

	if err := setGitConfig(username, email); err != nil {
		return err
	}

	return nil
}

func isGitConfigSet() bool {
	username, err := runner.RunWithOutput("git", "config", "--get", "user.name")
	if err != nil {
		return false
	}
	email, err := runner.RunWithOutput("git", "config", "--get", "user.email")
	if err != nil {
		return false
	}
	return username != "" && email != ""
}

func askForUsername() (string, error) {
	return userinput.AskForInput("Enter your username for git: ")
}

func askForEmail() (string, error) {
	return userinput.AskForInput("Enter your email for git: ")
}

func setGitConfig(username, email string) error {
	if err := runner.Run("git", "config", "--global", "user.name", username); err != nil {
		return err
	}
	if err := runner.Run("git", "config", "--global", "user.email", email); err != nil {
		return err
	}
	return nil
}
