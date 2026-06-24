# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

The ixo Blockchain is a Cosmos SDK Layer 1 blockchain (v6) for coordinating, financing, and verifying impacts. Built on Cosmos SDK v0.50.10, CometBFT v0.38.21, and IBC-Go v8.7.0 with CosmWasm smart contract support.

**Binary:** `ixod`
**Go Version:** 1.22.11
**Module:** `github.com/ixofoundation/ixo-blockchain/v8`

## Common Commands

### Build & Install
```bash
make build              # Build ixod binary to ./build/
make install            # Install ixod to $GOPATH/bin
make build-dev-install  # Install with debug symbols (for debugging)
```

### Testing
```bash
make test-unit          # Run unit tests (excludes e2e and simulator)
make test-race          # Run with race detector
make test-cover         # Run with coverage (outputs coverage.txt)

# Run a single test
go test -v -run TestFunctionName ./x/module/...

# Simulation tests
make test-sim-app           # Full app simulation
make test-sim-determinism   # State determinism validation
```

### Linting
```bash
make lint-all           # Run golangci-lint
make lint-format        # Run with auto-fix
```

### Protocol Buffers
```bash
make proto-gen          # Generate Go code from proto files
make proto-all          # Format, generate, swagger, and docs
```

### Local Development
```bash
make run                # Clean build and run single node
make run_with_all_data  # Run with demo data
docker-compose up       # Run 3-validator local network
```

## Architecture

### Application Structure
- **`/app`** - Main application (IxoApp extends BaseApp)
  - `app.go` - App initialization and wiring
  - `modules.go` - Module registration and ordering
  - `ante.go` - Custom ante handler chain
  - `keepers/` - All keeper initialization
  - `upgrades/` - Version upgrade handlers (v2-v6)

- **`/cmd/ixod`** - CLI entry point

- **`/x/`** - Custom Cosmos SDK modules (8 total):
  - **bonds** - Token bonding curves for automated market-making
  - **iid** - Interchain Identifiers (W3C DID documents)
  - **entity** - Digital twins with NFT backing and verifiable credentials
  - **claims** - W3C verifiable claims with evaluations and payments
  - **token** - Custom token minting, burning, and metadata
  - **epochs** - Configurable on-chain timers for periodic events
  - **mint** - Custom inflation/deflation mechanism
  - **smart-account** - Advanced account authentication with custom authenticators

- **`/proto/ixo/`** - Protobuf definitions organized by module
- **`/wasmbinding`** - CosmWasm custom bindings
- **`/ixomath`** - Custom math libraries (decimals, exponents)

### Module Pattern
Each module in `/x/{module}/` follows standard Cosmos SDK structure:
- `types/` - Messages, genesis, keys, errors, events
- `keeper/` - State management and business logic
- `client/cli/` - CLI commands
- `spec/` - Module specification docs (for bonds, iid, entity, claims, token, liquidstake)

### Key Dependencies
- Cosmos SDK modules: auth, bank, staking, gov, distribution, slashing, evidence, feegrant, authz, params, crisis, consensus, upgrade
- IBC modules: core, transfer, ICA (controller/host), fee
- IBC middleware: packet-forward-middleware, ibc-hooks, async-icq
- CosmWasm: wasmd v0.50.0 with wasmvm v1.5.4

### Keeper Initialization
Keepers are centralized in `/app/keepers/keepers.go`. The `AppKeepers` struct holds all module keepers and is embedded in `IxoApp`.

### Upgrade System
Upgrades are defined in `/app/upgrades/` with handlers registered in `app.go`. Each upgrade version (v2, v3, v4, v5, v6) has its own package with migration logic.

## Bech32 Address Prefixes
- Account: `ixo`
- Validator: `ixovaloper`
- Consensus: `ixovalcons`

## Ports (Docker Compose)
- API: 1317, 1318, 1319
- RPC: 26657, 26658, 26659
- gRPC: 9090, 9091, 9092
