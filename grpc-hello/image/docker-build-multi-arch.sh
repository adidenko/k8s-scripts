#!/bin/bash

IMAGE="adgit/grpc-hello"
TAG="20241009-v1"

docker buildx build \
  --no-cache \
  --push \
  --platform linux/arm64,linux/amd64 \
  --tag ${IMAGE}:${TAG} .
