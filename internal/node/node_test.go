package node

import "testing"

type fakeRunner struct{}

func (f *fakeRunner) Run(name string, arg ...string) error {
	return nil
}

func (f *fakeRunner) RunWithOutput(name string, arg ...string) (string, error) {
	return "", nil
}

func TestSetupNode(t *testing.T) {
	nodeInstall := false
	isNodeInstalled = func() bool {
		return nodeInstall
	}
	installNode = func() error {
		nodeInstall = true
		return nil
	}
	if err := SetupNode(); err != nil {
		t.Errorf("Expected nil, got %v", err)
	}

	if nodeInstall == false {
		t.Errorf("Expected neovimInstall to be true, got false")
	}
}
