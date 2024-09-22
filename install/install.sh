#!/bin/bash

command_exists() {
    command -v "$1" >/dev/null 2>&1
}

install_mpv() {
    if command_exists apt-get; then
        echo "Installing mpv using apt-get..."
        sudo apt-get update -y
        sudo apt-get install -y mpv
    elif command_exists dnf; then
        echo "Installing mpv using dnf..."
        sudo dnf install -y mpv
    else
        echo "Unsupported package manager. Please install mpv manually."
        exit 1
    fi
}

if command_exists mpv; then
    echo "mpv is installed."
else
    echo "mpv is not installed. Installing mpv..."
    install_mpv
    if ! command_exists mpv; then
        echo "Failed to install mpv. Please install it manually."
        exit 1
    fi
fi

PROJECT_DIR=$(pwd)
OUTPUT_BINARY="mankka"

cd "$PROJECT_DIR" || { echo "Failed to change directory"; exit 1; }

echo "Building Go project..."
go build -o "$OUTPUT_BINARY"

if [ $? -eq 0 ]; then
    echo "Build succeeded. Binary created: $OUTPUT_BINARY"
else
    echo "Build failed."
    exit 1
fi

echo "Adding created binary $OUTPUT_BINARY to /usr/bin/"
sudo cp $OUTPUT_BINARY /usr/bin/$OUTPUT_BINARY
