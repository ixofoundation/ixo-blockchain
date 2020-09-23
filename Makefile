#!/usr/bin/make -f

VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')
SDK_PACK := $(shell go list -m github.com/cosmos/cosmos-sdk | sed  's/ /\@/g')

export GO111MODULE = on
export COSMOS_SDK_TEST_KEYRING = n

# process build tags

build_tags =
ifeq ($(WITH_CLEVELDB),yes)
  build_tags += gcc
endif
build_tags += $(BUILD_TAGS)
build_tags := $(strip $(build_tags))

whitespace :=
whitespace += $(whitespace)
comma := ,
build_tags_comma_sep := $(subst $(whitespace),$(comma),$(build_tags))

# process linker flags

ldflags = \
    -X github.com/cosmos/cosmos-sdk/version.Name=ixo \
	-X github.com/cosmos/cosmos-sdk/version.ServerName=ixod \
	-X github.com/cosmos/cosmos-sdk/version.ClientName=ixocli \
	-X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
	-X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) \
	-X "github.com/ixofoundation/ixo-blockchain/version.BuildTags=$(build_tags_comma_sep)"

ifeq ($(WITH_CLEVELDB),yes)
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=cleveldb
endif
ldflags += $(LDFLAGS)
ldflags := $(strip $(ldflags))

BUILD_FLAGS := -tags "$(build_tags)" -ldflags '$(ldflags)'

all: lint install

build: go.sum
ifeq ($(OS),Windows_NT)
	go build -mod=readonly $(BUILD_FLAGS) -o build/ixod.exe ./cmd/ixod
	go build -mod=readonly $(BUILD_FLAGS) -o build/ixocli.exe ./cmd/ixocli
else
	go build -mod=readonly $(BUILD_FLAGS) -o build/ixod ./cmd/ixod
	go build -mod=readonly $(BUILD_FLAGS) -o build/ixocli ./cmd/ixocli
endif

install: go.sum
	go install -mod=readonly $(BUILD_FLAGS) ./cmd/ixod
	go install -mod=readonly $(BUILD_FLAGS) ./cmd/ixocli

########################################
### Tools & dependencies

go-mod-cache: go.sum
	@echo "--> Download go modules to local cache"
	@go mod download

go.sum: go.mod
	@echo "--> Ensure dependencies have not been modified"
	@go mod verify

draw-deps:
	@# requires brew install graphviz or apt-get install graphviz
	go get github.com/RobotsAndPencils/goviz
	@goviz -i github.com/ixofoundation/ixo-blockchain/cmd/ixod -d 2 | dot -Tpng -o dependency-graph.png

.PHONY: all install go-mod-cache draw-deps build

########################################
### Run

run:
	./scripts/clean_build.sh
	./scripts/run_only.sh

run_with_some_data:
	./scripts/clean_build.sh
	./scripts/run_with_some_data.sh

run_with_all_data:
	./scripts/clean_build.sh
	./scripts/run_with_all_data.sh
