package git

import "testing"

func TestGitSetup(t *testing.T) {
	isGitConfigSet = func() bool {
		return false
	}
	askForEmail = func() (string, error) {
		return "testmail", nil
	}
	askForUsername = func() (string, error) {
		return "testuser", nil
	}
	setGitConfig = func(_, _ string) error {
		return nil
	}
	err := SetupGit()
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
}
