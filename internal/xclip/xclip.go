package xclip

func SetupXclip() error {
	// TODO: add xclip install, sudo apt install xclip
	println(isXclipInstalled())
	return nil
}

var isXclipInstalled = func() bool {
	// TODO: impl
	return false
}

var installXclip = func() error {
	// TODO: impl
	return nil
}
