#!/bin/bash

# Clean up
rm -rf lib/api lib/api/client

# Create directories if they don't exist
mkdir -p lib/api lib/api/client

# Copy client directory
cp -r internal/api/client/* lib/api/client/

echo "Successfully copied client and api directories to /lib"

# ---

# Copy api to apipublic
rm -rf apipublic

mkdir -p apipublic

cp -r api/v1alpha1 apipublic/v1alpha1

rm -rf apipublic/v1alpha1/validation.go

echo "Successfully copied api to apipublic"
