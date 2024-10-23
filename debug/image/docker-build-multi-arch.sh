#!/bin/bash

IMAGE="adgit/debug"
TAG="20241023-v1"

docker buildx build \
  --no-cache \
  --push \
  --platform linux/arm64,linux/amd64 \
  --tag ${IMAGE}:${TAG} .
