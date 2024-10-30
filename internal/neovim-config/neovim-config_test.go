package neovimconfig

import "testing"

func TestNeovimConfig(t *testing.T) {
	isNeovimConfigClonedResult := false

	neovimConfigExists = func() bool {
		return isNeovimConfigClonedResult
	}
	cloneNeovimConfig = func() error { return nil }

	if err := SetupNeovimConfig(); err != nil {
		t.Errorf("NeovimConfig() failed: %v", err)
	}

	if !isNeovimConfigClonedResult {
		t.Errorf("NeovimConfig() failed: Neovim config not cloned")
	}
}
