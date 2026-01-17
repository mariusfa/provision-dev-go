package brewinstaller

import (
	"errors"
	"testing"
)

type fakeRunner struct {
	IsInstalled bool
}

func (f *fakeRunner) Run(name string, arg ...string) error {
	if name == "curl" && arg[0] == "--version" && f.IsInstalled {
		return nil
	}
	if name == "brew" && arg[0] == "install" && arg[1] == "curl" {
		f.IsInstalled = true
		return nil
	}

	return errors.New("command failed")
}

func (f *fakeRunner) RunWithOutput(name string, arg ...string) (string, error) {
	return "", nil
}

func newFakeRunner() *fakeRunner {
	return &fakeRunner{IsInstalled: false}
}

func TestBrewInstaller(t *testing.T) {
	fakeCommandRunner := newFakeRunner()
	runner = fakeCommandRunner
	if err := InstallPackage("curl"); err != nil {
		t.Errorf("InstallPackage() failed: %v", err)
	}

	if !fakeCommandRunner.IsInstalled {
		t.Error("Package 'curl' was not installed as expected")
	}
}
