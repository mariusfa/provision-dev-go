# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

Personal development environment provisioning tool that automates setup of development tools on Ubuntu/Fedora systems. Zero external dependencies - uses only Go standard library.

## Commands

```bash
# Run all provisioning tasks
./provision.sh

# Run specific setup task
go run main.go [command]
# Commands: ssh, git, apps, alias, neovim, neovim-config, node, neovim-req, go, tmux-config

# Run tests
go test ./...
```

## Architecture

### CLI Router Pattern
`main.go` uses a switch statement to route commands to setup functions in `internal/` packages. Each setup function runs independently and is idempotent (safe to run multiple times).

### Command Execution Abstraction
All packages use `ICommandRunner` interface for shell commands:
```go
type ICommandRunner interface {
    Run(name string, arg ...string) error
    RunWithOutput(name string, arg ...string) (string, error)
}
```
This enables testing with mock implementations.

### Package Manager Abstraction
`utils/pkg-installer` auto-detects APT or Brew and routes `InstallPackage()` to the appropriate installer.

### Binary Installation Pattern
Binary installers (neovim, node, golang) follow this flow:
1. Download tar.gz to ~/apps/
2. Extract to ~/apps/{package-name}
3. Create symlink in ~/apps/bin/
4. Cleanup tar file

Versions are hardcoded in each package.

### Testing Pattern
Tests mock package-level functions:
```go
getAliasesFromBash = func() ([]string, error) {
    return []string{}, nil
}
```

## Key Directories

- `internal/` - Feature packages (ssh, git, neovim, node, etc.)
- `internal/utils/` - Shared utilities (installers, CLI parser, command runner)
- `aliases` - Bash aliases configuration file
