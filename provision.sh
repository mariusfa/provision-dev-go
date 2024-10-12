#!/bin/bash

GO_VERSION="go1.23.2.linux-amd64.tar.gz"
GO_INSTALL_PATH="/usr/local/go"

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

function run_go_file {
  local go_file="$1"
  if [ -f "$go_file" ]; then
    echo "Running Go file $go_file"
    go run "$go_file"
  else
    echo "Go file $go_file does not exist."
    exit 1
  fi
}

function main {
  if is_go_installed; then
    echo "Go is already installed."
  else
    download_go
    install_go
  fi

  run_go_file "$1"
}

# Entry point
main "$1"
