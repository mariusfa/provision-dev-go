package neovimconfig

import "testing"

func TestNeovimConfig(t *testing.T) {
	isNeovimConfigClonedResult := false

	isNeovimConfigSetup = func() bool {
		return isNeovimConfigClonedResult
	}
	cloneNeoviConfig = func() error { return nil }

	if err := SetupNeovimConfig(); err != nil {
		t.Errorf("NeovimConfig() failed: %v", err)
	}

	if !isNeovimConfigClonedResult {
		t.Errorf("NeovimConfig() failed: Neovim config not cloned")
	}
}
