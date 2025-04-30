#!/bin/bash

IMAGE="adgit/ioperf"
TAG="20250430-v1"

docker buildx build \
  --no-cache \
  --push \
  --platform linux/arm64,linux/amd64 \
  --tag ${IMAGE}:${TAG} .
