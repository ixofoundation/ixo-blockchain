###############################################################################
###                                 Tests                                   ###
###############################################################################

PACKAGES_UNIT=$(shell go list ./... | grep -E -v 'tests/simulator|e2e')
PACKAGES_E2E := $(shell go list ./... | grep '/e2e' | awk -F'/e2e' '{print $$1 "/e2e"}' | uniq)
PACKAGES_SIM=$(shell go list ./... | grep '/tests/simulator')
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
	@VERSION=$(VERSION) go test -mod=readonly -run ^TestFullAppSimulation -v $(PACKAGES_SIM)

test-sim-determinism:
	@VERSION=$(VERSION) go test -mod=readonly -run ^TestAppStateDeterminism -v $(PACKAGES_SIM)

test-sim-bench:
	@VERSION=$(VERSION) go test -benchmem -run ^BenchmarkFullAppSimulation -bench ^BenchmarkFullAppSimulation -cpuprofile cpu.out $(PACKAGES_SIM)
