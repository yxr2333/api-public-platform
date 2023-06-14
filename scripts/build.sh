#!/bin/sh

# Change the directory to the project root
cd "$(dirname "$0")/.."

# Build the Go project
go build -o build/main cmd/main.go