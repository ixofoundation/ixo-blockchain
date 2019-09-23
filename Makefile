# COMMIT_HASH is an argument passed from the Dockerfile
ifeq ($(COMMIT_HASH),)
  # in the case of it not being passed (Jenkins) get the commit hash from the current directory (will fail if not a git repo)
  COMMIT_HASH := $(shell git rev-parse --short HEAD)
endif

export GO111MODULE=on

all:install

build: go.sum
	rm -rf build
ifeq ($(OS),Windows_NT)
	go build  -o build/ixod.exe ./cmd/ixod
	go build  -o build/ixocli.exe ./cmd/ixocli
else
	go build  -o build/ixod ./cmd/ixod
	go build  -o build/ixocli ./cmd/ixocli
endif

install: go.sum
	go install  ./cmd/ixod
	go install  ./cmd/ixocli

go.sum:
	@echo "--> Ensure dependencies have not been modified"
	@go mod verify

.PHONY: all build install  go.sum
