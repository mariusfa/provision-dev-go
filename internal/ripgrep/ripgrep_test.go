package ripgrep

import "testing"

func TestSetupRipgrep(t *testing.T) {
	ripgrepInstalled := false

	isRipgrepInstalled = func() bool {
		return ripgrepInstalled
	}

	installRipgrep = func() error {
		ripgrepInstalled = true
		return nil
	}

	if err := SetupRipgrep(); err != nil {
		t.Errorf("SetupRipgrep() = %v; want nil", err)
	}

	if !ripgrepInstalled {
		t.Error("ripgrepInstalled not installed")
	}
}
