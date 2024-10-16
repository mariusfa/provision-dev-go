package cliparser

import "testing"

func TestCliParser(t *testing.T) {
	GetCliArgs = func() []string {
		return []string{"ssh"}
	}

	result, error := CliParser()
	if error != nil {
		t.Errorf("Expected nil, got %v", error)
	}
	if result != SSH {
		t.Errorf("Expected SSH, got %v", result)
	}
}
