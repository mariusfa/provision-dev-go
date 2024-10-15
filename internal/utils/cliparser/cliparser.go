package cliparser

import "os"

type SubPackage int

const (
	SSH SubPackage = iota
	ALL
	UNKNOWN
)

type CliParserF func() SubPackage

func CliParser() SubPackage {
	args := os.Args[1:]
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
