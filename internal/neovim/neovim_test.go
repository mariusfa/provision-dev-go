package neovim

import "testing"

type fakeRunner struct{}

func (f *fakeRunner) Run(name string, arg ...string) error {
	return nil
}

func (f *fakeRunner) RunWithOutput(name string, arg ...string) (string, error) {
	return "", nil
}

func newFakeRunner() *fakeRunner {
	return &fakeRunner{}
}

func TestSetupNeovim(t *testing.T) {
	neovimInstall := false
	isNeovimInstalled = func() bool {
		return neovimInstall
	}
	installNeovim = func() error {
		neovimInstall = true
		return nil
	}
	if err := SetupNeovim(); err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	if neovimInstall == false {
		t.Errorf("Expected neovimInstall to be true, got false")
	}
}
