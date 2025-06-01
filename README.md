# Prvision Dev Go
Personal provisioning of my ubuntu development environment. Using Go for this becuase why not? :)

## Usage
```bash
./provision.sh
```

This will install a go version in `/usr/local/go`, and then run the `main.go` file in the current directory.

`main.go` will also install go in `~/app/go` and create a symlink in `~/apps/bin/go`. This go version will be your main go version for development.

After the first run, you can use `go run main.go` to run the provisioning script again.

## Running seperate parts of provisioning
```bash
go run main.go <app-name>
```

Example for neovim install:
```bash
go run main.go nvim
```


### Copilot auth
Shutdown WSL and restart ubuntu. Then start neovim and `:Copilot setup`.
```command
wsl.exe --shutdown
```

### Upgrading a bin package
This is useful for packages such as:
- nvim
- go
- node

First uninstall the old version. See the section about "Uninstall bin packages".

Bump filename version for specific package in provision-dev-go prosject.

Then:
```bash
go run main.go <app-name>
```


### Uninstall bin packages
Delete folder `~/app/<app-name>`

Delete bin symlink in `~/apps/bin/<app-name>`

This is useful for updating a bin package such as `nvim` or `node`.
