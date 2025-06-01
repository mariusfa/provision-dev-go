

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
