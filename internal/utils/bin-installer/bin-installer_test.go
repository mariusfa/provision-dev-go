package bininstaller

import (
	"errors"
	"testing"
)

type fakeRunner struct {
	IsInstalled bool
}

func (f *fakeRunner) Run(name string, arg ...string) error {
	if name == "nvim" && arg[0] == "--version" && f.IsInstalled {
		return nil
	}

	if name == "wget" && arg[0] == "https://fake-download-url.com" {
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

func TestBinInstaller(t *testing.T) {
	fakeCommandRunner := newFakeRunner()
	runner = fakeCommandRunner

	if err := InstallPackage("nvim", "https://fake-download-url.com"); err != nil {
		t.Errorf("InstallPackage() failed: %v", err)
	}
	if !fakeCommandRunner.IsInstalled {
		t.Error("Package 'nvim' was not installed as expected")
	}

}
