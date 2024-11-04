package xclip

import "testing"

func TestSetupXclip(t *testing.T) {
	xclipInstalled := false

	isXclipInstalled = func() bool {
		return xclipInstalled
	}

	installXclip = func() error {
		xclipInstalled = true
		return nil
	}

	if err := SetupXclip(); err != nil {
		t.Errorf("SetupXclip() = %v; want nil", err)
	}

	if !xclipInstalled {
		t.Error("xclip not installed")
	}
}
