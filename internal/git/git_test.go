package git

import "testing"

func TestGitSetup(t *testing.T) {
	// TODO: setup fake user input
	// TODO: setup fake cli get git config
	// TODO: setup fake cli setup git config
	err := SetupGit()
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
	t.Skip("Not implemented yet")
}
