#!/bin/bash

# Clean up
rm -rf lib/api lib/api/client

# Create directories if they don't exist
mkdir -p lib/api lib/api/client

# Copy client directory
cp -r internal/api/client/* lib/api/client/

echo "Successfully copied client and api directories to /lib"
