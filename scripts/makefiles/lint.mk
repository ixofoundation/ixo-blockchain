###############################################################################
###                                Linting                                  ###
###############################################################################
lint-help:
	@echo "lint subcommands"
	@echo ""
	@echo "Usage:"
	@echo "  make lint-[command]"
	@echo ""
	@echo "Available Commands:"
	@echo "  all                   Run all linters"
	@echo "  format                Run linters with auto-fix"
	@echo "  markdown              Run markdown linter with auto-fix"
	@echo "  mdlint                Run markdown linter"
	@echo "  setup-pre-commit      Set pre-commit git hook"
	@echo "  typo                  Run codespell to check typos (not implemented yet)"
	@echo "  fix-typo              Run codespell to fix typos (not implemented yet)"
lint: lint-help

golangci_lint_cmd=golangci-lint
golangci_version=v1.53.3

lint-all:
	@echo "--> Running linter"
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(golangci_version)
	@$(golangci_lint_cmd) run --timeout=10m

lint-format:
	@echo "--> Running linter"
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(golangci_version)
	@$(golangci_lint_cmd) run --fix --out-format=tab --issues-exit-code=0

# format:
# 	@go install mvdan.cc/gofumpt@latest
# 	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(golangci_version)
# 	find . -name '*.go' -type f -not -path "./vendor*" -not -path "*.git*" -not -path "./client/docs/statik/statik.go" -not -path "./tests/mocks/*" -not -name "*.pb.go" -not -name "*.pb.gw.go" -not -name "*.pulsar.go" -not -path "./crypto/keys/secp256k1/*" | xargs gofumpt -w -l
# 	$(golangci_lint_cmd) run --fix

lint-setup-pre-commit:
	@cp .git/hooks/pre-commit .git/hooks/pre-commit.bak 2>/dev/null || true
	@echo "Installing pre-commit hook..."
	@ln -sf ../../scripts/hooks/pre-commit.sh .git/hooks/pre-commit
	@echo "Pre-commit hook installed successfully"

CODESPELL_DOCKER_IMAGE=ixo-codespell

lint-build-image:
	@docker build -t $(CODESPELL_DOCKER_IMAGE) -f ./.infra/dockerfiles/Dockerfile.codespell .

lint-typo:
	@make lint-build-image
	@docker run -v $(PWD):/app $(CODESPELL_DOCKER_IMAGE) codespell

lint-fix-typo:
	@make lint-build-image
	@docker run -v $(PWD):/app $(CODESPELL_DOCKER_IMAGE) codespell -w
