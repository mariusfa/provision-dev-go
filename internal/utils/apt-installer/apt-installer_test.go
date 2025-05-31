package aptinstaller

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
	if name == "sudo" && arg[0] == "apt" && arg[1] == "install" && arg[2] == "-y" && arg[3] == "curl" {
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

func TestAptInstaller(t *testing.T) {
	fakeCommandRunner := newFakeRunner()
	runner = fakeCommandRunner
	if err := InstallPackage("curl"); err != nil {
		t.Errorf("InstallPackage() failed: %v", err)
	}

	if !fakeCommandRunner.IsInstalled {
		t.Error("Package 'curl' was not installed as expected")
	}
}
