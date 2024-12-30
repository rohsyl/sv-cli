#!/bin/bash

# Check if Go is installed
if ! command -v go &> /dev/null; then
    echo "Go is not installed. Installing Go..."

    # Determine the latest Go version
    GO_VERSION="1.21.1"
    GO_URL="https://go.dev/dl/go${GO_VERSION}.linux-amd64.tar.gz"
    TEMP_FILE="/tmp/go.tar.gz"

    # Download and install Go
    curl -Lo "$TEMP_FILE" "$GO_URL"
    sudo tar -C /usr/local -xzf "$TEMP_FILE"
    rm -f "$TEMP_FILE"

    # Add Go to PATH
    echo "export PATH=\$PATH:/usr/local/go/bin" >> ~/.bashrc
    source ~/.bashrc

    echo "Go has been installed."
else
    echo "Go is already installed."
fi

# Download sv binary
echo "Downloading sv-cli binary..."
SV_URL="https://github.com/rohsyl/sv-cli/releases/download/latest/sv"
INSTALL_DIR="/usr/local/bin"
SV_PATH="$INSTALL_DIR/sv"

# Download the binary
sudo curl -Lo "$SV_PATH" "$SV_URL"
sudo chmod +x "$SV_PATH"

echo "Installation complete. You can now use sv-cli."
