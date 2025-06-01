package bininstaller

import (
	"fmt"
	"testing"
)

type fakeRunner struct {
	IsInstalled bool
}

func (f *fakeRunner) Run(name string, arg ...string) error {
	if name == "nvim" && arg[0] == "--version" && f.IsInstalled {
		return nil
	}

	if name == "wget" && arg[0] == "https://fake-download-url.com/file.tar.gz" {
		return nil
	}

	if name == "tar" && arg[0] == "-xvf" && arg[1] == "file.tar.gz" {
		return nil
	}

	if name == "ln" && arg[0] == "-s" && arg[1] == "/home/mariusfa/apps/nvim/bin/nvim" && arg[2] == "/home/mariusfa/apps/bin/nvim" {
		f.IsInstalled = true
		return nil
	}

	if name == "rm" && arg[0] == "file.tar.gz" {
		return nil
	}

	return fmt.Errorf("unexpected command: %s %v", name, arg)
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

	if err := InstallPackage("nvim", "https://fake-download-url.com/file.tar.gz"); err != nil {
		t.Errorf("InstallPackage() failed: %v", err)
	}
	if !fakeCommandRunner.IsInstalled {
		t.Error("Package 'nvim' was not installed as expected")
	}

}
