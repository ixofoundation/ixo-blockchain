###############################################################################
###                                  Build                                  ###
###############################################################################

build-help:
	@echo "build subcommands"
	@echo ""
	@echo "Usage:"
	@echo "  make build-[command]"
	@echo ""
	@echo "Available Commands:"
	@echo "  all                              Build all targets"
	@echo "  check-version                    Check Go version"
	@echo "  dev-build                        Build development version"
	@echo "  dev-install                      Install development build"
	@echo "  install-with-autocomplete        Install with autocomplete support"
	@echo "  linux                            Build for Linux"
	@echo "  reproducible                     Build reproducible binaries"
	@echo "  reproducible-amd64               Build reproducible amd64 binary"
	@echo "  reproducible-arm64               Build reproducible arm64 binary"

build-check-version:
	@echo "Go version: $(GO_MAJOR_VERSION).$(GO_MINOR_VERSION)"
	@if [ $(GO_MAJOR_VERSION) -gt $(GO_MINIMUM_MAJOR_VERSION) ]; then \
		echo "Go version is sufficient"; \
		exit 0; \
	elif [ $(GO_MAJOR_VERSION) -lt $(GO_MINIMUM_MAJOR_VERSION) ]; then \
		echo '$(GO_VERSION_ERR_MSG)'; \
		exit 1; \
	elif [ $(GO_MINOR_VERSION) -lt $(GO_MINIMUM_MINOR_VERSION) ]; then \
		echo '$(GO_VERSION_ERR_MSG)'; \
		exit 1; \
	fi

build-all: build-check-version go.sum
	mkdir -p $(BUILDDIR)/
	GOWORK=off go build -mod=readonly $(BUILD_FLAGS) -o $(BUILDDIR)/ ./...

# disables optimization, inlining and symbol removal
GC_FLAGS := -gcflags="all=-N -l"
REMOVE_STRING := -w -s
DEBUG_BUILD_FLAGS:= $(subst $(REMOVE_STRING),,$(BUILD_FLAGS))
DEBUG_LDFLAGS = $(subst $(REMOVE_STRING),,$(ldflags))

build-dev-install: go.sum
	GOWORK=off go install $(DEBUG_BUILD_FLAGS) $(GC_FLAGS) ./cmd/ixod

build-dev-build:
	mkdir -p $(BUILDDIR)/
	GOWORK=off go build $(GC_FLAGS) -mod=readonly -ldflags '$(DEBUG_LDFLAGS)' -gcflags "all=-N -l" -trimpath -o $(BUILDDIR) ./...;

build-install-with-autocomplete: build-check-version go.sum
	GOWORK=off go install -mod=readonly $(BUILD_FLAGS) ./cmd/ixod
	@PARENT_SHELL=$$(ps -o ppid= -p $$PPID | xargs ps -o comm= -p); \
	if echo "$$PARENT_SHELL" | grep -q "zsh"; then \
		if ! grep -q ". <(ixod enable-cli-autocomplete zsh)" ~/.zshrc; then \
			echo ". <(ixod enable-cli-autocomplete zsh)" >> ~/.zshrc; \
			echo; \
			echo "Autocomplete enabled. Run 'source ~/.zshrc' to complete installation."; \
		else \
			echo; \
			echo "Autocomplete already enabled in ~/.zshrc"; \
		fi \
	elif echo "$$PARENT_SHELL" | grep -q "bash" && [ "$$(uname)" = "Darwin" ]; then \
		if ! grep -q -e "\. <(ixod enable-cli-autocomplete bash)" -e '\[\[ -r "/opt/homebrew/etc/profile.d/bash_completion.sh" \]\] && \. "/opt/homebrew/etc/profile.d/bash_completion.sh"' ~/.bash_profile; then \
			brew install bash-completion; \
			echo '[ -r "/opt/homebrew/etc/profile.d/bash_completion.sh" ] && . "/opt/homebrew/etc/profile.d/bash_completion.sh"' >> ~/.bash_profile; \
			echo ". <(ixod enable-cli-autocomplete bash)" >> ~/.bash_profile; \
			echo; \
			echo; \
			echo "Autocomplete enabled. Run 'source ~/.bash_profile' to complete installation."; \
		else \
			echo "Autocomplete already enabled in ~/.bash_profile"; \
		fi \
	elif echo "$$PARENT_SHELL" | grep -q "bash" && [ "$$(uname)" = "Linux" ]; then \
		if ! grep -q ". <(ixod enable-cli-autocomplete bash)" ~/.bash_profile; then \
			sudo apt-get install -y bash-completion; \
			echo '[ -r "/etc/bash_completion" ] && . "/etc/bash_completion"' >> ~/.bash_profile; \
			echo ". <(ixod enable-cli-autocomplete bash)" >> ~/.bash_profile; \
			echo; \
			echo "Autocomplete enabled. Run 'source ~/.bash_profile' to complete installation."; \
		else \
			echo; \
			echo "Autocomplete already enabled in ~/.bash_profile"; \
		fi \
	else \
		echo "Shell or OS not recognized. Skipping autocomplete setup."; \
	fi


# Cross-building for arm64 from amd64 (or vice-versa) takes
# a lot of time due to QEMU virtualization but it's the only way (afaik)
# to get a statically linked binary with CosmWasm
build-reproducible: build-reproducible-amd64 build-reproducible-arm64

build-reproducible-amd64: go.sum
	mkdir -p $(BUILDDIR)
	$(DOCKER) buildx create --name ixobuilder || true
	$(DOCKER) buildx use ixobuilder
	$(DOCKER) buildx build \
		--build-arg GO_VERSION=$(GO_VERSION) \
		--build-arg GIT_VERSION=$(VERSION) \
		--build-arg GIT_COMMIT=$(COMMIT) \
		--build-arg RUNNER_IMAGE=alpine:3.17 \
		--platform linux/amd64 \
		-t ixo:local-amd64 \
		--load \
		-f Dockerfile .
	$(DOCKER) rm -f ixobinary || true
	$(DOCKER) create -ti --name ixobinary ixo:local-amd64
	$(DOCKER) cp ixobinary:/bin/ixod $(BUILDDIR)/ixod-linux-amd64
	$(DOCKER) rm -f ixobinary

build-reproducible-arm64: go.sum
	mkdir -p $(BUILDDIR)
	$(DOCKER) buildx create --name ixobuilder || true
	$(DOCKER) buildx use ixobuilder
	$(DOCKER) buildx build \
		--build-arg GO_VERSION=$(GO_VERSION) \
		--build-arg GIT_VERSION=$(VERSION) \
		--build-arg GIT_COMMIT=$(COMMIT) \
		--build-arg RUNNER_IMAGE=alpine:3.17 \
		--platform linux/arm64 \
		-t ixo:local-arm64 \
		--load \
		-f Dockerfile .
	$(DOCKER) rm -f ixobinary || true
	$(DOCKER) create -ti --name ixobinary ixo:local-arm64
	$(DOCKER) cp ixobinary:/bin/ixod $(BUILDDIR)/ixod-linux-arm64
	$(DOCKER) rm -f ixobinary

build-linux: go.sum
	LEDGER_ENABLED=false GOOS=linux GOARCH=amd64 $(MAKE) build
