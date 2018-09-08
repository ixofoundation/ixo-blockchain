#!/bin/sh

COMMIT_HASH=$(git rev-parse --short HEAD)
docker build -t trustlab/ixo-blockchain:dev  --build-arg COMMIT_HASH=$COMMIT_HASH .