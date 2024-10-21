package alias

import (
	"bufio"
	"os"
	"path/filepath"
	"strings"
)

func SetupAlias() error {
	return nil
}

var writeAliases = func(aliases []string) error {
	// TODO: Implement
	return nil
}

var getAliasesFromBash = func() ([]string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return []string{}, err
	}

	bashrcPath := filepath.Join(homeDir, ".bashrc")
	return getAliasFromFile(bashrcPath)
}

var getAliasesFromAliasFile = func() ([]string, error) {
	return getAliasFromFile("./aliases")
}

func getAliasFromFile(filePath string) ([]string, error) {
	aliasesFound := []string{}

	file, err := os.Open(filePath)
	if err != nil {
		return aliasesFound, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "alias") {
			aliasesFound = append(aliasesFound, line)
		}
	}

	return aliasesFound, nil
}
