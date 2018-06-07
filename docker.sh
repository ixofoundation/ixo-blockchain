#!/bin/sh
docker build -t trustlab/ixo-blockchain -f - . < ./docker/blockchain/Dockerfile
docker build -t trustlab/ixo-blockchain-rest -f - . < ./docker/rest/Dockerfile
docker-compose up -d