#!/bin/bash

# Copy api to apipublic
rm -rf lib/apipublic

mkdir -p lib/apipublic

cp -r api/v1alpha1 lib/apipublic/v1alpha1

rm -rf lib/apipublic/v1alpha1/validation.go # this file import the internal/api/v1alpha1/validation.go, need to remove it

echo "Successfully copied api to lib/apipublic"
