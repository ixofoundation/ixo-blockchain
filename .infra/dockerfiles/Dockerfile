# Dockerfile for docker environment to build and run the ixo blockchain

FROM --platform=linux/amd64 golang:1.22-bookworm AS base

# Install system dependencies
RUN apt-get update && apt-get install -y \
    build-essential \
    ca-certificates \
    curl \
    git \
    jq \
    && rm -rf /var/lib/apt/lists/*

# Install Ignite CLI (successor to Starport) - v28 for Cosmos SDK v0.50 compatibility
RUN curl https://get.ignite.com/cli@v28.11.2! | bash

# Download correct libwasmvm version based on go.mod
WORKDIR /app
COPY go.mod go.sum ./
RUN WASMVM_VERSION=$(go list -m github.com/CosmWasm/wasmvm | sed 's/.* //') && \
    ARCH=$(uname -m) && \
    wget https://github.com/CosmWasm/wasmvm/releases/download/${WASMVM_VERSION}/libwasmvm.${ARCH}.so \
    -O /usr/local/lib/libwasmvm.${ARCH}.so && \
    ldconfig

COPY ./ .
