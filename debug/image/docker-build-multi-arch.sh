#!/bin/bash

IMAGE="adgit/debug"
TAG="20260105-v1"
TTLSH="ttl.sh/adgit-debug-${TAG}"

# Ask which image to build
echo "Do you want to build regular image or ttl image?"
echo
echo "1) regular"
echo "2) ttl"
echo
read -p "Select 1 or 2: " choice

if [[ "$choice" == "2" ]] ; then
  IMAGE="$TTLSH"
  TAG="12h"
fi

echo
echo "Building and pushing multi-arch image: ${IMAGE}:${TAG}"

docker buildx build \
  --no-cache \
  --push \
  --platform linux/arm64,linux/amd64 \
  --tag ${IMAGE}:${TAG} .
