#!/usr/bin/env bash

# Exit immediately if a command fails
set -e

# Set environment variables
export RSS_URL="https://www.theguardian.com/world/rss"
export SOURCE_IDS="the_guardian, bbc_news_int"
export SOURCE_TYPE="s3"
export TARGET_TYPE="postgres"

# Optional: print what we're running
echo "Running Go app with:"
echo "  RSS_URL=$RSS_URL"
echo "  SOURCE_ID=$SOURCE_ID"
echo "  SOURCE_TYPE=$SOURCE_TYPE"
echo "  TARGET_TYPE=$TARGET_TYPE"

# Run the Go program
go run ./cmd/main.go
