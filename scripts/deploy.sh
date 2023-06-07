#!/usr/bin/env bash
# This script builds and deploys the backend on the servers.

# Exit if any command returns an error code
set -e
cd webapp
git stash && git checkout main && git pull origin
GOOS=linux GOARCH=amd64 /usr/local/go/bin/go build -o "$BE_EXEC"
sudo service "$BE_EXEC" restart
