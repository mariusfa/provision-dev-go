package alias

import "testing"

func TestSetupAlias(t *testing.T) {
	// TODO: setup fakes
	err := SetupAlias()
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}
}
