package aptinstaller

import "testing"

type fakeRunner struct {
	isInstalled bool
}

func (f *fakeRunner) Run(name string, arg ...string) error {
	// check curl --version return error
	// check  apt install -y curl return nil
	return nil
}

func (f *fakeRunner) RunWithOutput(name string, arg ...string) (string, error) {
	return "", nil
}

func newFakeRunner() *fakeRunner {
	return &fakeRunner{}
}

func TestAptInstaller(t *testing.T) {
	fakeCommandRunner := newFakeRunner()
	aptInstaller := NewAptInstaller()
	aptInstaller.commandRunner = fakeCommandRunner
	if err := aptInstaller.InstallPackage("curl"); err != nil {
		t.Errorf("InstallPackage() failed: %v", err)
	}

	expectedCheckInstalledCommand := "curl --version"
	if fakeCommandRunner.checkInstalled != expectedCheckInstalledCommand {
		t.Errorf("Expected command: %s, got: %s", expectedCheckInstalledCommand, fakeCommandRunner.checkInstalled)
	}

	expectedInstallCommand := "sudo apt install -y curl"
	if fakeCommandRunner.installCommand != expectedInstallCommand {
		t.Errorf("Expected command: %s, got: %s", expectedInstallCommand, fakeCommandRunner.installCommand)
	}
}
