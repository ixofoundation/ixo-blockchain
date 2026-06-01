# ixo-blockchain test suite

This directory holds the layered test surfaces for the ixo blockchain.
The companion document `../tests.md` is the master plan and progress
tracker; this file is the operational quick-start.

## Layers

```
L3 · interchaintest  — Docker-based E2E. Slow (minutes). Nightly.
L2 · simulator       — In-process IxoApp. Random tx ops + invariants.
L1 · keeper unit     — Real keeper, gomock'd deps. Fast (seconds).
```

Each layer has its own Make target.

## Running

| Goal | Command |
|---|---|
| Quick local run (L1 only) | `make test-unit` |
| With race detector | `make test-race` |
| Coverage report | `make test-cover` |
| Apptesting integration | `make test-integration` |
| Simulator smoke (skip-by-default) | `make test-sim-app` |
| Simulator with state-determinism check | `make test-sim-determinism` |
| Simulator benchmarks | `make test-sim-bench` |
| Interchaintest E2E (Docker) | `make test-interchaintest` |
| Interchaintest skip-by-default | `make test-interchaintest-short` |
| Regenerate gomock mocks | `make mocks-gen` |

## Layer 1 — Keeper unit tests

Located in `x/{module}/keeper/*_test.go` and `x/{module}/types/*_test.go`.
Each module has a `KeeperTestSuite` embedding `app/apptesting.KeeperTestHelper`,
which boots a fresh in-memory `IxoApp` per test (see `app/apptesting/test_suite.go`).
Pattern:

```go
type KeeperTestSuite struct {
    apptesting.KeeperTestHelper
    msgServer types.MsgServer
}
func (s *KeeperTestSuite) SetupTest() {
    s.Setup()
    s.msgServer = keeper.NewMsgServerImpl(s.App.MyKeeper)
}
```

Wasm-dependent flows (CreateToken, CreateEntity NFT instantiation,
claims payment settlement) are intentionally NOT covered here — they
need a live wasmvm and live in `tests/interchaintest/`.

## Layer 2 — Simulator

`tests/simulator/sim_test.go` runs `simulation.SimulateFromSeed` against
the IxoApp. All four entry points (`TestFullAppSimulation`,
`TestAppStateDeterminism`, `TestAppImportExport`, etc.) skip-by-default
unless `-Enabled=true` is supplied; this is the cosmos-sdk convention.

Each custom module satisfies `module.AppModuleSimulation` (see
`x/{module}/module.go`) — `RegisterStoreDecoder`, `WeightedOperations`
(currently empty for ixo modules), `ProposalContents`, `ProposalMsgs`,
`GenerateGenesisState`. The decoders live in `x/{module}/simulation/decoder.go`.

Known issue: the live `-Enabled=true` run hits a height-handshake error.
See the entry in `tests.md::Bug Log` for details and resolution path.

## Layer 3 — interchaintest

`tests/interchaintest/` is a separate Go sub-module (own go.mod) so the
strangelove dep tree doesn't bleed into the main project. Build-tag
gated: `//go:build interchaintest`. The test inventory is documented in
`tests/interchaintest/README.md`.

Run:

```bash
docker build -t ixofoundation/ixo-blockchain:local .
export IXO_IMAGE=ixofoundation/ixo-blockchain:local
make test-interchaintest
```

## Adding a new test

1. **Keeper test** — add a `Test*` method on the existing
   `KeeperTestSuite` for the module. If the module needs a new fixture,
   add a helper to `keeper_test.go`.
2. **New `Msg*` server method** — also add a `simulation.NewWeighted*`
   entry to `x/{module}/simulation/operations.go` (currently a no-op
   stub) and an interchaintest test for the live-chain path.
3. **New module** — add `_test.go` files under `keeper/` and `types/`,
   create `simulation/decoder.go`, and wire `AppModuleSimulation` methods
   into `module.go`.

## CONTRIBUTING checklist for new code

- [ ] Add a keeper-level test with happy-path + at least one error path.
- [ ] Run `make test-unit && make test-race` locally.
- [ ] If the change affects a `Msg*` handler, also touch the matching
      simulation/operations.go entry (even if just a doc-comment update
      explaining why a stub remains).
- [ ] If the change affects a wasm-dependent flow, add a corresponding
      test under `tests/interchaintest/`.
