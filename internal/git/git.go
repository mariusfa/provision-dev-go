package git

import "provision/internal/utils/commandrunner"

var runner commandrunner.ICommandRunner = commandrunner.NewCommandRunner()

func SetupGit() error {
	// TODO: Implement
	println("Not implemented yet")
	isSet := isGitConfigSet()

	println(isSet)
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
