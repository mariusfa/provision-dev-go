package apps

import "testing"

func TestSetupAppsFolderNotExits(t *testing.T) {
	// TODO: implement
	appsFolder := false
	folderExists = func() bool {
		return appsFolder
	}
	updateBashRc = func() error {
		appsFolder = true
		return nil
	}
	if err := SetupApps(); err != nil {
		t.Errorf("SetupApps() = %v; want nil", err)
	}

	if !appsFolder {
		t.Errorf("appsFolder = %v; want true", appsFolder)
	}
}
