#!/bin/bash
# This script builds backend for the wigit web app and deploys to remote servers

# Exit if any command fails
set -e

# Clone the repository
git clone https://github.com/wigit-gh/webapp.git

# Change directory into the repository
cd webapp

# Build the backend
GOOS=linux GOARCH=amd64 go build -o wwapp_be

# Add key to ssh-agent
eval "$(ssh-agent -s)"
ssh-add ~/.ssh/id_rsa

# Copy the binary to the servers and resart the service
scp -i ~/.ssh/id_rsa wwapp_be "$BACKEND01":~/webapp/
ssh -i ~/.ssh/id_rsa "$BACKEND01" "sudo service wwapp_be restart"
scp -i ~/.ssh/id_rsa wwapp_be "$BACKEND02":~/webapp/
ssh -i ~/.ssh/id_rsa "$BACKEND02" "sudo service wwapp_be restart"
