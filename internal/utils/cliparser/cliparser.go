package cliparser

import "os"

type SubPackage int

const (
	SSH SubPackage = iota
	ALL
	UNKNOWN
)

var GetCliArgs = func() []string {
	return os.Args[1:]
}

// TODO: use err instead of unknown
func CliParser() SubPackage {
	args := GetCliArgs()
	argsLength := len(args)
	if argsLength > 1 {
		return UNKNOWN
	}

	if argsLength == 0 {
		return ALL
	}

	command := args[0]

	if command == "ssh" {
		return SSH
	}
	return UNKNOWN
}
