#!/bin/bash

GO_VERSION="go1.23.2.linux-amd64.tar.gz"
GO_INSTALL_PATH="/usr/local/go"

function print_initialization_banner {
  local border="============================"
  local message="Script provision-dev initializing"

  echo "$border"
  echo "$message"
  echo "$border"
}

function is_go_installed {
  if command -v go &> /dev/null; then
    return 0
  else
    return 1
  fi
}

function download_go {
  echo "Downloading Go version $GO_VERSION..."
  wget https://go.dev/dl/$GO_VERSION
}

function install_go {
  echo "Installing Go..."
  sudo rm -rf $GO_INSTALL_PATH
  sudo tar -C /usr/local -xzf $GO_VERSION
  rm $GO_VERSION
  echo "export PATH=\$PATH:/usr/local/go/bin" >> ~/.profile
  source ~/.profile
  echo "Go has been installed."
}

function run_go_project {
  if [ -f "cmd/main.go" ]; then
    echo "Running Go project from cmd/main.go"
    go run cmd/main.go
  else
    echo "cmd/main.go does not exist."
    exit 1
  fi
}

function main {
  print_initialization_banner
  if is_go_installed; then
    echo "Go is already installed."
  else
    download_go
    install_go
  fi

  run_go_project
}

# Entry point
main "$1"
