package alias

import "testing"

func TestSetupAliasNoAliases(t *testing.T) {
	aliasesWritten := []string{}
	writeAliases := func(aliases []string) error {
		aliasesWritten = aliases
		return nil
	}
	getAliasesFromBash := func() ([]string, error) {
		return []string{}, nil
	}
	getAliasesFromAliasFile := func() ([]string, error) {
		return []string{
			"alias ..='cd ..'",
		}, nil
	}
	err := SetupAlias()
	if err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	if len(aliasesWritten) != 1 {
		t.Errorf("Expected 1, got %v", len(aliasesWritten))
	}

	if aliasesWritten[0] != "alias ..='cd ..'" {
		t.Errorf("Expected 'alias ..='cd ..'', got %v", aliasesWritten[0])
	}
}
