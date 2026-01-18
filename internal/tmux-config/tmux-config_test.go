package tmuxconfig

import "testing"

func TestSetupTmuxConfig(t *testing.T) {
	isTmuxConfigClonedResult := false
	isCloneTmuxConfigCalled := false

	tmuxConfigExists = func() bool {
		return isTmuxConfigClonedResult
	}
	cloneTmuxConfig = func() error {
		isTmuxConfigClonedResult = true
		isCloneTmuxConfigCalled = true
		return nil
	}

	if err := SetupTmuxConfig(); err != nil {
		t.Errorf("SetupTmuxConfig() failed: %v", err)
	}

	if !isTmuxConfigClonedResult {
		t.Errorf("SetupTmuxConfig() failed: Tmux config not cloned")
	}
	if !isCloneTmuxConfigCalled {
		t.Errorf("SetupTmuxConfig() failed: cloneTmuxConfig() not called")
	}
}

func TestSetupTmuxConfigAlreadyCloned(t *testing.T) {
	isTmuxConfigClonedResult := true
	isCloneTmuxConfigCalled := false
	tmuxConfigExists = func() bool {
		return isTmuxConfigClonedResult
	}
	cloneTmuxConfig = func() error {
		isCloneTmuxConfigCalled = true
		return nil
	}
	if err := SetupTmuxConfig(); err != nil {
		t.Errorf("SetupTmuxConfig() failed: %v", err)
	}
	if !isTmuxConfigClonedResult {
		t.Errorf("SetupTmuxConfig() failed: Tmux config not cloned")
	}
	if isCloneTmuxConfigCalled {
		t.Errorf("SetupTmuxConfig() failed: cloneTmuxConfig() called")
	}
}
