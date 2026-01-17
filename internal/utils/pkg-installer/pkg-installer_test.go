package pkginstaller

import (
	"testing"
)

func TestDetectPackageManager(t *testing.T) {
	// This tests the actual system detection
	// On a system with apt, it should return APT
	// On a system with brew (but no apt), it should return BREW
	manager := detectPackageManager()
	if manager != APT && manager != BREW {
		t.Errorf("detectPackageManager() returned unexpected value: %v", manager)
	}
}
