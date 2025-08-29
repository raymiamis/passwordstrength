#!/bin/bash

set -e

BINARY_NAME="passwordstrength"
INSTALL_PATH="/usr/local/bin"

echo "Checking Go-Installation..."
if ! command -v go &> /dev/null; then
    echo -e "\e[31mGo is not installed.\e[0m Please install Go and try again."
    echo "Check https://go.dev for Go installation guide"
    exit 1
fi

echo "Building $BINARY_NAME..."
go build -o "$BINARY_NAME"

echo "Installing to $INSTALL_PATH (sudo required)..."
sudo cp "$BINARY_NAME" "$INSTALL_PATH"

echo -e "\e[32mInstallation complete!\e[0m"
echo ""
echo "Use the password strength checker wherever you want with:"
echo -e "  \e[36m$BINARY_NAME\e[0m"
