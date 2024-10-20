package git

import "provision/internal/utils/commandrunner"

var runner commandrunner.ICommandRunner = commandrunner.NewCommandRunner()

func SetupGit() error {
	// TODO: Implement
	println("Not implemented yet")
	isSet, err := isGitConfigSet()
	if err != nil {
		return err
	}

	println(isSet)
	return nil
}

func isGitConfigSet() (bool, error) {
	username, err := runner.RunWithOutput("git", "config", "--get", "user.name")
	if err != nil {
		return false, err
	}
	email, err := runner.RunWithOutput("git", "config", "--get", "user.email")
	if err != nil {
		return false, err
	}
	return username != "" && email != "", nil
}
