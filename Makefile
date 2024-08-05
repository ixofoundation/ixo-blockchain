#!/usr/bin/make -f

# the subcommands are located in the specific makefiles
include scripts/makefiles/build.mk
include scripts/makefiles/deps.mk
include scripts/makefiles/docker.mk
include scripts/makefiles/lint.mk
include scripts/makefiles/proto.mk
include scripts/makefiles/release.mk

.DEFAULT_GOAL := help
help:
	@echo "Available top-level commands:"
	@echo ""
	@echo "Usage:"
	@echo "    make [command]"
	@echo ""
	@echo "  make build                 Build ixod binary"
	@echo "  make build-help            Show available build commands"
	@echo "  make deps                  Show available deps commands"
	@echo "  make docker                Show available docker commands"
	@echo "  make install               Install ixod binary"
	@echo "  make lint                  Show available lint commands"
	@echo "  make proto                 Show available proto commands"
	@echo "  make release               Create a new release"
	@echo "  make release-help          Show available release commands"
	@echo ""
	@echo "Run 'make [subcommand]' to see the available commands for each subcommand."

VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')

LEDGER_ENABLED ?= true
SDK_PACK := $(shell go list -m github.com/cosmos/cosmos-sdk | sed  's/ /\@/g')
BUILDDIR ?= $(CURDIR)/build
DOCKER := $(shell which docker)
# HTTPS_GIT := https://github.com/ixofoundation/ixo-blockchain.git

# Go version to be used in docker images
GO_VERSION := $(shell cat go.mod | grep -E 'go [0-9].[0-9]+' | cut -d ' ' -f 2)
# currently installed Go version
GO_MODULE := $(shell cat go.mod | grep "module " | cut -d ' ' -f 2)
GO_MAJOR_VERSION = $(shell go version | cut -c 14- | cut -d' ' -f1 | cut -d'.' -f1)
GO_MINOR_VERSION = $(shell go version | cut -c 14- | cut -d' ' -f1 | cut -d'.' -f2)
# minimum supported Go version
GO_MINIMUM_MAJOR_VERSION = $(shell cat go.mod | grep -E 'go [0-9].[0-9]+' | cut -d ' ' -f2 | cut -d'.' -f1)
GO_MINIMUM_MINOR_VERSION = $(shell cat go.mod | grep -E 'go [0-9].[0-9]+' | cut -d ' ' -f2 | cut -d'.' -f2)
# message to be printed if Go does not meet the minimum required version
GO_VERSION_ERR_MSG = "ERROR: Go version $(GO_MINIMUM_MAJOR_VERSION).$(GO_MINIMUM_MINOR_VERSION)+ is required"

export GO111MODULE = on
# export COSMOS_SDK_TEST_KEYRING = n

###############################################################################
###                            Build Tags/Flags                             ###
###############################################################################

# process build tags

build_tags = netgo
ifeq ($(LEDGER_ENABLED),true)
	ifeq ($(OS),Windows_NT)
    GCCEXE = $(shell where gcc.exe 2> NUL)
    ifeq ($(GCCEXE),)
      $(error gcc.exe not installed for ledger support, please install or set LEDGER_ENABLED=false)
    else
      build_tags += ledger
    endif
  else
    UNAME_S = $(shell uname -s)
    ifeq ($(UNAME_S),OpenBSD)
      $(warning OpenBSD detected, disabling ledger support (https://github.com/cosmos/cosmos-sdk/issues/1988))
    else
      GCC = $(shell command -v gcc 2> /dev/null)
      ifeq ($(GCC),)
        $(error gcc not installed for ledger support, please install or set LEDGER_ENABLED=false)
      else
        build_tags += ledger
      endif
    endif
  endif
endif

ifeq ($(WITH_CLEVELDB),yes)
  build_tags += gcc
endif

build_tags += $(BUILD_TAGS)
build_tags := $(strip $(build_tags))

whitespace :=
whitespace := $(whitespace) $(whitespace)
comma := ,
build_tags_comma_sep := $(subst $(whitespace),$(comma),$(build_tags))

# process linker flags

ldflags = \
    -X github.com/cosmos/cosmos-sdk/version.Name=ixo \
    -X github.com/cosmos/cosmos-sdk/version.AppName=ixod \
    -X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
    -X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) \
    -X "github.com/cosmos/cosmos-sdk/version.BuildTags=$(build_tags_comma_sep)"

ifeq ($(WITH_CLEVELDB),yes)
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=cleveldb
endif
ifeq ($(LINK_STATICALLY),true)
	ldflags += -linkmode=external -extldflags "-Wl,-z,muldefs -static"
endif
ldflags += $(LDFLAGS)
ldflags := $(strip $(ldflags))

BUILD_FLAGS := -tags "$(build_tags)" -ldflags '$(ldflags)'


###############################################################################
###                             Build / Install                             ###
###############################################################################

build: build-check-version go.sum
	mkdir -p $(BUILDDIR)/
	GOWORK=off go build -mod=readonly $(BUILD_FLAGS) -o $(BUILDDIR)/ ./cmd/ixod

install: build-check-version go.sum
	GOWORK=off go install -mod=readonly $(BUILD_FLAGS) ./cmd/ixod

###############################################################################
###                                 RUN                                     ###
###############################################################################

run:
	./scripts/clean_build.sh
	./scripts/run_only.sh

run_with_all_data:
	./scripts/clean_build.sh
	./scripts/run_with_all_data.sh

run_with_genesis:
	./scripts/run_with_genesis.sh

###############################################################################
###                                Release                                  ###
###############################################################################

GORELEASER_IMAGE := ghcr.io/goreleaser/goreleaser-cross:v$(GO_VERSION)
COSMWASM_VERSION := $(shell go list -m github.com/CosmWasm/wasmvm | sed 's/.* //')

ifdef GITHUB_TOKEN
release:
	docker run \
		--rm \
		-e GITHUB_TOKEN=$(GITHUB_TOKEN) \
		-e COSMWASM_VERSION=$(COSMWASM_VERSION) \
		-v /var/run/docker.sock:/var/run/docker.sock \
		-v `pwd`:/go/src/ixod \
		-w /go/src/ixod \
		$(GORELEASER_IMAGE) \
		release \
		--clean
else
release:
	@echo "Error: GITHUB_TOKEN is not defined. Please define it before running 'make release'."
endif

.PHONY: all build-linux install format lint \
	go-mod-cache draw-deps clean build \
	release release-dry-run release-snapshot update-deps
