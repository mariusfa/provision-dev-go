package apps

import "os"

var appsPath = os.Getenv("HOME") + "/apps/bin"

func SetupApps() error {
	if !folderExists() {
		if err := createFolder(); err != nil {
			return err
		}
		updateBashRc()
	}
	return nil
}

var folderExists = func() bool {
	info, err := os.Stat(appsPath)
	if os.IsNotExist(err) || !info.IsDir() {
		return false
	}
	return true
}

var updateBashRc = func() error {
	// TODO: Implement
	return nil
}

var createFolder = func() error {
	return os.MkdirAll(appsPath, os.ModePerm)
}
