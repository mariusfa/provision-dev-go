package gcc

import "testing"

func TestBuildEssentials(t *testing.T) {
	gccInstalled := false

	isGccInstalled = func() bool {
		return gccInstalled
	}

	installGcc = func() error {
		gccInstalled = true
		return nil
	}

	if err := SetupGcc(); err != nil {
		t.Errorf("SetupBuildEssentials() failed: %v", err)
	}

	if !gccInstalled {
		t.Errorf("SetupBuildEssentials() failed: build essentials not installed")
	}
}
