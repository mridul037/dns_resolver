#!/bin/bash

# Simple DNS Lookup Installer

# Define the repository URL
REPO_URL="https://github.com/mridul037/dns-resolver.git"

# Function to display messages
echo_message() {
    echo -e "\n*** $1 ***\n"
}

# Check for git
if ! command -v git &> /dev/null; then
    echo_message "Git is not installed. Please install Git first."
    exit 1
fi

# Clone the repository
echo_message "Cloning the repository..."
git clone $REPO_URL

# Change to the repository directory
cd simple-dns-lookup || { echo_message "Clone failed! Exiting."; exit 1; }

# Make the script executable
echo_message "Making the script executable..."
chmod +x dns_lookup.sh

# Inform the user of the next steps
echo_message "Installation complete!"
echo "You can now use the DNS lookup tool by running:"
echo "./dns_lookup.sh <domain> <record_type> [dns_server]"
