package buildessentials

import "testing"

func TestBuildEssentials(t *testing.T) {
	beInstalled := false

	isBuildEssentialsInstalled := func() bool {
		return beInstalled
	}

	installBuildEssentials := func() error {
		beInstalled = true
		return nil
	}

	if err := SetupBuildEssentials(); err != nil {
		t.Errorf("SetupBuildEssentials() failed: %v", err)
	}
}
