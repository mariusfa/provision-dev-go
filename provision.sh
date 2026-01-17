#!/bin/bash

GO_VERSION="go1.25.6.linux-amd64.tar.gz"
GO_TEMP_DIR="/tmp/go-bootstrap"

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

function download_and_extract_go {
  echo "Downloading Go version $GO_VERSION..."
  mkdir -p $GO_TEMP_DIR
  wget -q -O "$GO_TEMP_DIR/$GO_VERSION" "https://go.dev/dl/$GO_VERSION"
  tar -C $GO_TEMP_DIR -xzf "$GO_TEMP_DIR/$GO_VERSION"
  rm "$GO_TEMP_DIR/$GO_VERSION"
}

function run_go_project {
  local go_bin="$1"
  if [ -f "main.go" ]; then
    echo "Running Go project from main.go"
    $go_bin run main.go
  else
    echo "main.go does not exist."
    exit 1
  fi
}

function cleanup {
  echo "Cleaning up temporary Go installation..."
  rm -rf $GO_TEMP_DIR
}

function main {
  print_initialization_banner

  if is_go_installed; then
    echo "Go is already installed."
    run_go_project "go"
  else
    download_and_extract_go
    run_go_project "$GO_TEMP_DIR/go/bin/go"
    cleanup
  fi
}

# Entry point
main "$1"
