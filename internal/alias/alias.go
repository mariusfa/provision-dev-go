package alias

import (
	"bufio"
	"os"
	"path/filepath"
	"slices"
	"strings"
)

func SetupAlias() error {
	bashAliases, err := getAliasesFromBash()
	if err != nil {
		return err
	}

	aliasesInProject, err := getAliasesFromAliasFile()
	if err != nil {
		return err
	}

	aliasesToInsert := findMissingAliasesInBashrc(aliasesInProject, bashAliases)
	if len(aliasesToInsert) > 0 {
		err = writeAliases(aliasesToInsert)
		if err != nil {
			return err
		}
		println("Aliases inserted in .bashrc")
	} else {
		println("All aliases are already in .bashrc")
	}

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

func findMissingAliasesInBashrc(aliases []string, bashAliases []string) []string {
	missingAliases := []string{}
	for _, alias := range aliases {
		if !slices.Contains(bashAliases, alias) {
			missingAliases = append(missingAliases, alias)
		}
	}
	return missingAliases
}

var writeAliases = func(aliases []string) error {
	// TODO: Implement this function
	return nil
}
