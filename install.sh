#!/bin/bash

set -e

BINARY_NAME="passwordstrength"
INSTALL_PATH="/usr/local/bin"

GREEN='\033[0;32m'
RED='\033[0;31m'
NC='\033[0m'

echo "Checking Go-Installation..."
if ! command -v go &> /dev/null; then
    echo -e "${RED}Go is not installed. Please install Go and try again.${NC}"
    exit 1
fi

echo "Building $BINARY_NAME..."
go build -o "$BINARY_NAME"

echo "Installing to $INSTALL_PATH (sudo required)..."
sudo cp "$BINARY_NAME" "$INSTALL_PATH"

echo -e "${GREEN}Installation complete!${NC}"
echo ""
echo "Use the password strength checker wherever you want with:"
echo -e "  ${GREEN}$BINARY_NAME${NC}"
