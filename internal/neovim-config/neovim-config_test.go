package neovimconfig

import "testing"

func TestNeovimConfig(t *testing.T) {
	isNeovimConfigClonedResult := false
	isCloneNeovimConfigCalled := false

	neovimConfigExists = func() bool {
		return isNeovimConfigClonedResult
	}
	cloneNeovimConfig = func() error {
		isNeovimConfigClonedResult = true
		isCloneNeovimConfigCalled = true
		return nil
	}

	if err := SetupNeovimConfig(); err != nil {
		t.Errorf("NeovimConfig() failed: %v", err)
	}

	if !isNeovimConfigClonedResult {
		t.Errorf("NeovimConfig() failed: Neovim config not cloned")
	}
	if !isCloneNeovimConfigCalled {
		t.Errorf("NeovimConfig() failed: cloneNeovimConfig() not called")
	}
}
