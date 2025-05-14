#!/bin/bash

# Get the version from git tags
VERSION=$(git describe --tags --always --dirty 2>/dev/null || echo "dev")

# Build the project
echo "Building mood-tracker version: $VERSION"
go build -ldflags="-X 'main.Version=$VERSION'" -o mood-tracker


echo "Build complete! Version: $VERSION" 