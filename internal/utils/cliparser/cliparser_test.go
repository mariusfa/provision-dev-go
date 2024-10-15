package cliparser

import "testing"

func TestCliParser(t *testing.T) {
	GetCliArgs = func() []string {
		return []string{"ssh"}
	}

	result := CliParser()
	if result != SSH {
		t.Errorf("Expected SSH, got %v", result)
	}
}
