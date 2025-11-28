#!/bin/bash

set -e
echo "Installing Splitwise Core package..."
GOOS=$OS GOARCH=$ARCH go build -o splitwise ./cmd/
sudo mv splitwise /usr/local/bin/splitwise
echo "Build completed. Executable 'splitwise' created."