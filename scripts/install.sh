#!/bin/bash

# Download sv binary
echo "Downloading sv-cli binary..."
SV_URL="https://github.com/rohsyl/sv-cli/releases/download/latest/sv"
INSTALL_DIR="/usr/local/bin"
SV_PATH="$INSTALL_DIR/sv"

# Download the binary
sudo curl -Lo "$SV_PATH" "$SV_URL"
sudo chmod +x "$SV_PATH"

echo "Installation complete. You can now use sv-cli."
