package ssh

import (
	"os"
	"testing"
)

type fakeRunner struct{}

func (f *fakeRunner) Run(name string, arg ...string) error {
	return nil
}

func newFakeRunner() *fakeRunner {
	return &fakeRunner{}
}

func TestSsh(t *testing.T) {
	fakeRunner := newFakeRunner()
	runner = fakeRunner

	tmpDir := t.TempDir()
	os.Setenv("HOME", tmpDir)

	if err := generateSshKey("testmail"); err != nil {
		t.Errorf("Error setting up SSH: %v", err)
	}
}
