PACKAGES=$(shell go list ./... | grep -v '/vendor/')

# COMMIT_HASH is an argument passed from the Dockerfile
ifeq ($(COMMIT_HASH),)
  # in the case of it not being passed (Jenkins) get the commit hash from the current directory (will fail if not a git repo)
  COMMIT_HASH := $(shell git rev-parse --short HEAD)
endif

BUILD_FLAGS = -ldflags "-X github.com/ixofoundation/ixo-cosmos/version.GitCommit=${COMMIT_HASH}"

all: check_tools get_vendor_deps build test

########################################
### CI

ci: get_tools get_vendor_deps build test_cover

########################################
### Build

# This can be unified later, here for easy demos
build:
	rm -rf build
ifeq ($(OS),Windows_NT)
	go build $(BUILD_FLAGS) -o build/ixod.exe ./cmd/ixod
	go build $(BUILD_FLAGS) -o build/ixocli.exe ./cmd/ixocli
else
	go build $(BUILD_FLAGS) -o build/ixod ./cmd/ixod
	go build $(BUILD_FLAGS) -o build/ixocli ./cmd/ixocli
endif

install: 
	go install $(BUILD_FLAGS) ./cmd/ixod
	go install $(BUILD_FLAGS) ./cmd/ixocli

#dist:
#	@bash publish/dist.sh
#	@bash publish/publish.sh

########################################
### Tools & dependencies

check_tools:
	cd tools && $(MAKE) check_tools

update_tools:
	cd tools && $(MAKE) update_tools

get_tools:
	cd tools && $(MAKE) get_tools

get_vendor_deps:
	go mod verify

draw_deps:
	@# requires brew install graphviz or apt-get install graphviz
	go get github.com/RobotsAndPencils/goviz
	@goviz -i github.com/tendermint/tendermint/cmd/tendermint -d 3 | dot -Tpng -o dependency-graph.png


########################################
### Documentation

#godocs:
#	@echo "--> Wait a few seconds and visit http://localhost:6060/pkg/github.com/cosmos/cosmos-sdk/types"
#	godoc -http=:6060


########################################
### Testing

#test: test_unit # test_cli

# Must  be run in each package seperately for the visualization
# Added here for easy reference
# coverage:
#	 go test -coverprofile=c.out && go tool cover -html=c.out

#test_unit:
#	@go test $(PACKAGES)

#test_cover:
#	@bash tests/test_cover.sh

#benchmark:
#	@go test -bench=. $(PACKAGES)


########################################
### Devdoc

#DEVDOC_SAVE = docker commit `docker ps -a -n 1 -q` devdoc:local

#devdoc_init:
#	docker run -it -v "$(CURDIR):/go/src/github.com/cosmos/cosmos-sdk" -w "/go/src/github.com/cosmos/cosmos-sdk" tendermint/devdoc echo
	# TODO make this safer
#	$(call DEVDOC_SAVE)

#devdoc:
#	docker run -it -v "$(CURDIR):/go/src/github.com/cosmos/cosmos-sdk" -w "/go/src/github.com/cosmos/cosmos-sdk" devdoc:local bash

#devdoc_save:
	# TODO make this safer
#	$(call DEVDOC_SAVE)

#devdoc_clean:
#	docker rmi -f $$(docker images -f "dangling=true" -q)

#devdoc_update:
#	docker pull tendermint/devdoc


# To avoid unintended conflicts with file names, always add to .PHONY
# unless there is a reason not to.
# https://www.gnu.org/software/make/manual/html_node/Phony-Targets.html
.PHONY: build build_examples install install_examples dist check_tools get_tools get_vendor_deps draw_deps test test_unit test_tutorial benchmark devdoc_init devdoc devdoc_save devdoc_update
