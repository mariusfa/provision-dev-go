package cliparser

import (
	"errors"
	"os"
)

type SubPackage int

const (
	ALL SubPackage = iota
	SSH
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

	if command == "ssh" {
		return SSH, nil
	}
	return UNKNOWN, errors.New("unknown command")
}
