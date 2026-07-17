#!/bin/bash

echo "========================================="
echo "Starting Golang AWS Secret Manager..."
echo "========================================="

# Load environment variables
if [ -f .env ]; then
    export $(grep -v '^#' .env | xargs)
fi

# Download dependencies
go mod tidy

# Run the application
go run ./cmd/server