package buildessentials

import "testing"

func TestBuildEssentials(t *testing.T) {
	if err := SetupBuildEssentials(); err != nil {
		t.Errorf("SetupBuildEssentials() failed: %v", err)
	}
}
