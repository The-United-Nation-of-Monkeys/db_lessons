#!/bin/bash

# Script to generate RSA key pair for JWT tokens
# Usage: ./scripts/generate_keys.sh

KEYS_DIR="keys"
PRIVATE_KEY="$KEYS_DIR/private.pem"
PUBLIC_KEY="$KEYS_DIR/public.pem"

# Create keys directory if it doesn't exist
mkdir -p "$KEYS_DIR"

# Generate private key
openssl genrsa -out "$PRIVATE_KEY" 2048

# Generate public key from private key
openssl rsa -in "$PRIVATE_KEY" -pubout -out "$PUBLIC_KEY"

echo "Keys generated successfully!"
echo "Private key: $PRIVATE_KEY"
echo "Public key: $PUBLIC_KEY"
echo ""
echo "Make sure to add these files to .gitignore!"

