###############################################################################
###                                Release                                  ###
###############################################################################
release-help:
	@echo "release subcommands"
	@echo ""
	@echo "Usage:"
	@echo "  make release-[command]"
	@echo ""
	@echo "Available Commands:"
	@echo "  dry-run                   Perform a dry run release"
	@echo "  snapshot                  Create a snapshot release"

GORELEASER_IMAGE := ghcr.io/goreleaser/goreleaser-cross:v$(GO_VERSION)
COSMWASM_VERSION := $(shell go list -m github.com/CosmWasm/wasmvm | sed 's/.* //')

release-dry-run:
	docker run \
		--rm \
		-e COSMWASM_VERSION=$(COSMWASM_VERSION) \
		-v /var/run/docker.sock:/var/run/docker.sock \
		-v `pwd`:/go/src/ixod \
		-w /go/src/ixod \
		$(GORELEASER_IMAGE) \
		release \
		--clean \
		--skip-publish

# TODO check why added this before?
# --skip-validate

release-snapshot:
	docker run \
		--rm \
		-e COSMWASM_VERSION=$(COSMWASM_VERSION) \
		-v /var/run/docker.sock:/var/run/docker.sock \
		-v `pwd`:/go/src/ixod \
		-w /go/src/ixod \
		$(GORELEASER_IMAGE) \
		release \
		--clean \
		--snapshot \
		--skip-validate \
		--skip-publish