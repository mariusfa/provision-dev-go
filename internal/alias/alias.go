package alias

import (
	"bufio"
	"fmt"
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
	aliasesBashrc, err := getAliasFromFile(bashrcPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read aliases from bashrc: %w", err)
	}

	return aliasesBashrc, nil
}

var getAliasesFromAliasFile = func() ([]string, error) {
	aliasesInAliasFile, err := getAliasFromFile("./aliases")
	if err != nil {
		return nil, fmt.Errorf("failed to read from alias file: %w", err)
	}
	return aliasesInAliasFile, nil
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
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	bashrcPath := filepath.Join(homeDir, ".bashrc")
	return appendToFile(bashrcPath, aliases)
}

func appendToFile(filePath string, lines []string) error {
	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		return fmt.Errorf("failed to open file to append to: %w", err)
	}
	defer file.Close()

	for _, line := range lines {
		_, err = file.WriteString(line + "\n")
		if err != nil {
			return fmt.Errorf("failed to write alias to file: %w", err)
		}
	}

	return nil
}
