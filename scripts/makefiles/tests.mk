###############################################################################
###                                 Tests                                   ###
###############################################################################

PACKAGES_UNIT=$(shell go list ./... | grep -E -v 'tests/simulator|tests/interchaintest|e2e')
PACKAGES_E2E := $(shell go list ./... | grep '/e2e' | awk -F'/e2e' '{print $$1 "/e2e"}' | uniq)
PACKAGES_SIM=$(shell go list ./... | grep '/tests/simulator')
PACKAGES_INTEGRATION=$(shell go list ./... | grep -E 'apptesting|tests/integration')
TEST_PACKAGES=./...

test-help:
	@echo "test subcommands"
	@echo ""
	@echo "Usage:"
	@echo "  make test-[command]"
	@echo ""
	@echo "Available Commands:"
	@echo "  all                Run all tests"
	@echo "  unit               Run unit tests"
	@echo "  race               Run race tests"
	@echo "  cover              Run coverage tests"
	@echo "  benchmark          Run benchmark tests"
	@echo "  sim-app            Run sim app tests"
	@echo "  sim-bench          Run sim benchmark tests"
	@echo "  sim-determinism    Run sim determinism tests"
	@echo "  sim-suite          Run sim suite tests"
	@echo "  integration        Run integration tests (apptesting suite)"
	@echo "  interchaintest     Run interchaintest E2E (Docker, slow)"
	@echo "  mocks-gen          Regenerate gomock mock keepers across modules"

test: test-help

test-all: test-race test-covertest-unit test-build

test-unit:
	@VERSION=$(VERSION) SKIP_WASM_WSL_TESTS=$(SKIP_WASM_WSL_TESTS) go test -mod=readonly -tags='ledger test_ledger_mock norace' $(PACKAGES_UNIT)

test-race:
	@VERSION=$(VERSION) go test -mod=readonly -race -tags='ledger test_ledger_mock' $(PACKAGES_UNIT)

test-cover:
	@VERSION=$(VERSION) go test -mod=readonly -timeout 30m -coverprofile=coverage.txt -tags='norace' -covermode=atomic $(PACKAGES_UNIT)

test-benchmark:
	@go test -mod=readonly -bench=. $(PACKAGES_UNIT)

test-sim-suite:
	@VERSION=$(VERSION) go test -mod=readonly $(PACKAGES_SIM)

test-sim-app:
	@VERSION=$(VERSION) go test -mod=readonly -timeout 30m $(PACKAGES_SIM) -run ^TestFullAppSimulation -Enabled=true -NumBlocks=100 -BlockSize=200 -Period=5 -Seed=99 -v

test-sim-app-fast:
	@VERSION=$(VERSION) go test -mod=readonly $(PACKAGES_SIM) -run ^TestFullAppSimulation -Enabled=true -NumBlocks=10 -BlockSize=20 -Period=1 -Seed=99 -v

test-sim-determinism:
	@VERSION=$(VERSION) go test -mod=readonly $(PACKAGES_SIM) -run ^TestAppStateDeterminism -Enabled=true -NumBlocks=10 -BlockSize=20 -Period=1 -v

test-sim-import-export:
	@VERSION=$(VERSION) go test -mod=readonly $(PACKAGES_SIM) -run ^TestAppImportExport -Enabled=true -NumBlocks=10 -BlockSize=20 -Period=1 -Seed=99 -v

test-sim-bench:
	@VERSION=$(VERSION) go test -benchmem -run ^BenchmarkFullAppSimulation -bench ^BenchmarkFullAppSimulation -cpuprofile cpu.out $(PACKAGES_SIM)

test-integration:
	@VERSION=$(VERSION) go test -mod=readonly -tags='ledger test_ledger_mock norace' -timeout 30m $(PACKAGES_INTEGRATION)

test-interchaintest:
	@VERSION=$(VERSION) cd tests/interchaintest && go test -mod=readonly -tags='interchaintest' -timeout 60m ./...

test-interchaintest-short:
	@VERSION=$(VERSION) cd tests/interchaintest && go test -mod=readonly -tags='interchaintest' -short ./...

###############################################################################
###                                 Mocks                                   ###
###############################################################################

# mocks-gen regenerates uber/mock mock files for each module's expected_keepers.go.
# Generated files live alongside their source under x/<module>/testutil/.
# Re-run after editing any expected_keepers interface.
mocks-gen:
	@echo "Regenerating module keeper mocks..."
	@for mod in bonds claims entity epochs iid liquidstake mint names smart-account token; do \
		if [ -f "x/$$mod/types/expected_keepers.go" ]; then \
			echo "  -> x/$$mod"; \
			mkdir -p x/$$mod/testutil; \
			mockgen -source=x/$$mod/types/expected_keepers.go -package testutil -destination x/$$mod/testutil/expected_keepers_mocks.go; \
		fi; \
	done
	@echo "Done."
