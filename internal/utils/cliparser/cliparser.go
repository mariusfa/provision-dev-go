package cliparser

import (
	"errors"
	"os"
)

type SubPackage int

const (
	ALL SubPackage = iota
	XCLIP
	SSH
	GIT
	APPS
	ALIAS
	NEOVIM
	NEOVIM_CONFIG
	UNKNOWN
)

var GetCliArgs = func() []string {
	return os.Args[1:]
}

func CliParser() (SubPackage, error) {
	args := GetCliArgs()
	argsLength := len(args)
	if argsLength > 1 {
		return UNKNOWN, errors.New("too many args")
	}

	if argsLength == 0 {
		return ALL, nil
	}

	command := args[0]

	switch command {
	case "xclip":
		return XCLIP, nil
	case "ssh":
		return SSH, nil
	case "git":
		return GIT, nil
	case "apps":
		return APPS, nil
	case "alias":
		return ALIAS, nil
	case "neovim":
		return NEOVIM, nil
	case "neovim-config":
		return NEOVIM_CONFIG, nil
	}

	return UNKNOWN, errors.New("unknown command")
}
