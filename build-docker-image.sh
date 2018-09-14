#!/bin/sh

IMAGE_TAG=$1
if [ -z "$IMAGE_TAG" ];
then
  IMAGE_TAG=dev
fi
COMMIT_HASH=$(git rev-parse --short HEAD)
docker build -t trustlab/ixo-blockchain:$IMAGE_TAG  --build-arg COMMIT_HASH=$COMMIT_HASH .