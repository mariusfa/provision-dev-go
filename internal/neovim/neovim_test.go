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
	isNeovimInstalled = func() bool {
		return false
	}
	installNeovim = func() error {
		return nil
	}
}
