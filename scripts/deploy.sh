#!/usr/bin/env bash
# This script builds and deploys the backend on the servers.

# Exit if any command returns an error code
set -e
GOOS=linux GOARCH=amd64 $(which go) build -o "$BE_EXEC"
sudo service "$BE_EXEC" restart
