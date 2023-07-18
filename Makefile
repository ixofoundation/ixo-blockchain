#!/usr/bin/make -f

VERSION := $(shell echo $(shell git describe --tags) | sed 's/^v//')
COMMIT := $(shell git log -1 --format='%H')
LEDGER_ENABLED ?= yes
SDK_PACK := $(shell go list -m github.com/cosmos/cosmos-sdk | sed  's/ /\@/g')
TM_VERSION := $(shell go list -m github.com/tendermint/tendermint | sed 's:.* ::') # grab everything after the space in "github.com/tendermint/tendermint v0.34.7"
DOCKER := $(shell which docker)
DOCKER_BUF := $(DOCKER) run --rm -v $(CURDIR):/workspace --workdir /workspace bufbuild/buf
HTTPS_GIT := https://github.com/ixofoundation/ixo-blockchain.git

export GO111MODULE = on
export COSMOS_SDK_TEST_KEYRING = n

# process build tags

build_tags =
ifeq ($(WITH_CLEVELDB),yes)
  build_tags += gcc
endif
ifeq ($(LEDGER_ENABLED),yes)
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

build_tags += $(BUILD_TAGS)
build_tags := $(strip $(build_tags))

whitespace :=
whitespace += $(whitespace)
comma := ,
build_tags_comma_sep := $(subst $(whitespace),$(comma),$(build_tags))

# process linker flags

ldflags = \
    -X github.com/cosmos/cosmos-sdk/version.Name=ixo \
    -X github.com/cosmos/cosmos-sdk/version.AppName=ixod \
    -X github.com/cosmos/cosmos-sdk/version.Version=$(VERSION) \
    -X github.com/cosmos/cosmos-sdk/version.Commit=$(COMMIT) \
    -X "github.com/ixofoundation/ixo-blockchain/version.BuildTags=$(build_tags_comma_sep)" \
    -X github.com/tendermint/tendermint/version.TMCoreSemVer=$(TM_VERSION)

ifeq ($(WITH_CLEVELDB),yes)
  ldflags += -X github.com/cosmos/cosmos-sdk/types.DBBackend=cleveldb
endif
ldflags += $(LDFLAGS)
ldflags := $(strip $(ldflags))

BUILD_FLAGS := -tags "$(build_tags)" -ldflags '$(ldflags)'

###############################################################################
###                             Build / Install                             ###
###############################################################################

all: lint install

build: go.sum
ifeq ($(OS),Windows_NT)
	go build -mod=readonly $(BUILD_FLAGS) -o build/ixod.exe ./cmd/ixod
else
	go build -mod=readonly $(BUILD_FLAGS) -o build/ixod ./cmd/ixod
endif

install: go.sum
	go install -mod=readonly $(BUILD_FLAGS) ./cmd/ixod

###############################################################################
###                               Go Modules                                ###
###############################################################################

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
###                                Protobuf                                 ###
###############################################################################

protoVer=0.11.2
protoImageName=ghcr.io/cosmos/proto-builder:$(protoVer)
protoImage=$(DOCKER) run --rm -v $(CURDIR):/workspace --workdir /workspace $(protoImageName)

proto-all: proto-format proto-gen proto-docs

new-proto-gen:
	@echo "Generating Protobuf files"
	@$(protoImage) sh ./scripts/protoc-gen.sh

proto-swagger-gen:
	@echo "Generating Protobuf Swagger"
	@$(protoImage) sh ./scripts/protoc-swagger-gen.sh

proto-format:
	@$(protoImage) find ./ -not -path "./third_party/*" -name *.proto -exec clang-format -i {} \;

proto-lint:
	@$(protoImage) buf lint --error-format=json

proto-check-breaking:
	@$(protoImage) buf breaking --against $(HTTPS_GIT)#branch=main

proto-update-deps:
	@echo "Updating Protobuf dependencies"
	$(DOCKER) run --rm -v $(CURDIR)/proto:/workspace --workdir /workspace $(protoImageName) buf mod update

containerProtoVer=v0.2
containerProtoImage=tendermintdev/sdk-proto-gen:$(containerProtoVer)
containerProtoGen=cosmos-sdk-proto-gen-$(containerProtoVer)

proto-gen:
	@echo "Generating Protobuf files"
	docker rm $(containerProtoGen) || true
	docker run --name $(containerProtoGen) -v $(CURDIR):/workspace --workdir /workspace $(containerProtoImage) sh ./scripts/protocgen.sh

proto-docs:
	@echo "Generating Protobuf docs"
	docker rm $(containerProtoGen) || true
	docker run --name $(containerProtoGen) -v $(CURDIR):/workspace --workdir /workspace $(containerProtoImage) sh ./scripts/protoc-docs-gen.sh
