# interchaintest E2E suite

This directory holds the Docker-based end-to-end test suite for ixo-blockchain.
It targets `github.com/strangelove-ventures/interchaintest/v8` and is the
counterpart to the in-process unit/keeper tests under `x/*/keeper/`.

## Why a separate sub-module

interchaintest pulls in a large Docker / chain-spec dependency tree
(`testcontainers`, `lo`, `dockertest`, the full Cosmos chain registry, etc.).
We isolate it in `tests/interchaintest/go.mod` (this directory) so the main
project go.mod stays clean and `go test ./...` doesn't transitively download
those deps for normal contributors.

The sub-module follows Juno's `interchaintest/` pattern (see
`/Users/michael/dev/ixo/other/juno/interchaintest`).

## Build tag

Tests are gated by `//go:build interchaintest` so `make test-unit` (which
greps out `tests/interchaintest`) never tries to run them. Run them with:

```bash
make test-interchaintest
# or:
go test -tags interchaintest -timeout 60m ./tests/interchaintest/...
```

## Running locally

You'll need:

1. **Docker** running.
2. A local **ixod** image. Build with:
   ```bash
   docker build -t ixofoundation/ixo-blockchain:local .
   ```
   Then export `IXO_IMAGE=ixofoundation/ixo-blockchain:local` before running
   the tests.

The `setup.go` `IxoChainSpec()` helper returns a `*ibc.ChainConfig` pointing
at that image with the right denom / bech32 prefixes.

## Test inventory

Each test boots ONE chain and runs a multi-step scenario flow with
`t.Run` subtests for step-level reporting (matches the
`ixo-multiclient-sdk` Jest pattern). Earlier per-msg tests were
consolidated to amortise chain-boot cost and exercise cross-msg state
propagation — the silent-drop class of bug only shows up when many
messages run against the same persistent state.

### Helpers

- `setup.go` — `IxoChainSpec`, `SetupIxoChain`, `UploadAllContracts`, `SubmitGovProposalAndPass`, `VoteOnLatestProposalAndPass`. Genesis tweaks: 8s voting/deposit periods, wasm permissions=Everybody.
- `helpers.go` — module-level helpers (`CreateIidDoc`, `QueryIidDocument`, `IidExec`, etc.) so the scenario flows read as a story rather than CLI plumbing.
- `contracts/` — bundled wasm artefacts (cw721, ixo1155, cw20_base, cw721_base, cw4_group) sha256-pinned in `checksums.txt`, verified by `bootstrap.sh`.

### Chain-level flows

- `basic_test.go::TestIxoBasicStart` — chain start, single-validator block production, balance assertion.
- `ibc_test.go::TestIxoIBCTransfer` — ixo ↔ Gaia ICS-20 transfer with rly relayer.

### Module flows

- `module_iid_test.go::TestIxoIID_FullScenario` — 18+ subtests: covers every iid Msg* (CreateIidDocument, UpdateIidDocument, Add/Delete-Service, Add/Delete-LinkedResource, Add/Delete-Controller, AddVerification, RevokeVerification, Add/Delete-IidContext, Add/Delete-AccordedRight, Add/Delete-LinkedClaim, Add/Delete-LinkedEntity, SetVerificationRelationships, DeactivateIID), plus 3 rejection paths and the silent-drop regression for MsgUpdateIidDocument and MsgDeactivateIID.State.
- `module_bonds_test.go::TestIxoBonds_FullScenario` — 13 subtests covering most of x/bonds: power-curve bond (create, edit, buy, sell, max-supply rejection, update-state SETTLE, withdraw-share, make-outcome-payment) + a SECOND alpha-bond config covering SetNextAlpha and WithdrawReserve. OracleDid silent-drop regression.
- `module_bank_test.go::TestIxoBank_FullScenario` — 3 subtests: happy-path send + blocked module account rejection (gov) + allowed module account success (distribution).
- `multi_message_test.go::TestIxoMultiMessage_AtomicityScenario` — 2 subtests: two-good-MsgSends bundle delivers atomically; one-good + one-over-balance bundle reverts the WHOLE tx.
- `wasm_test.go::TestIxoWasm_FullScenario` — 3 subtests: upload all 5 contracts (deterministic code ID order), cw20 instantiate+transfer+balance smart-queries, cw721_base instantiate+mint+transfer+owner_of smart-queries.
- `module_gov_test.go::TestIxoGov_FullScenario` — 4 subtests, each a different gov proposal type: MsgCreateNamespace (names) + MsgSoftwareUpgrade (cosmos.upgrade) + MsgCreatePool (liquidstake) + voted-NO MsgCreateNamespace (rejection rollback).
- `module_chaintime_test.go::TestIxoChainTime_FullScenario` — 3 subtests: mint params + epoch-provisions queries + total uixo supply > 0 + epochs default identifiers (day/hour/week) all resolve.
- `module_txdelegation_test.go::TestIxoTxDelegation_FullScenario` — 3 subtests: authz SendAuthorization grant+exec, feegrant `--fee-granter` charges granter not grantee, smart-account add+query+remove SignatureVerification authenticator.
- `module_distribution_test.go::TestIxoDistribution_FullScenario` — 6 subtests: discover validator → delegate 1M uixo → query rewards (well-formed) → query delegations → query validator-distribution-info → query commission → set-withdraw-addr msg path.
- `module_staking_test.go::TestIxoStaking_FullScenario` — 6 subtests: query validators → delegate → query delegations → unbond fraction (creates unbonding-delegation entry) → cancel-unbond rolls it back → query delegator-validators.
- `module_slashing_test.go::TestIxoSlashing_FullScenario` — 2 subtests: signing-infos returns the validator's record (zero missed blocks, not tombstoned), slashing params resolve with sane defaults.
- `module_entity_test.go::TestIxoEntity_FullScenario` — 8 subtests covering all entity user-facing msgs: upload cw721 + gov-set NFT contract → create-entity (asset) → entity-list → update-entity-verified → create-entity-account → update-entity → grant-entity-account-authz (raw tx) → revoke-entity-account-authz (raw tx). (TransferEntity intentionally deferred — needs recipient-side auth setup.)
- `module_liquidstake_test.go::TestIxoLiquidStake_FullScenario` — 9 subtests covering all user/admin-facing liquidstake msgs: gov-create pool → admin update-whitelisted-validators (raw tx) → user liquid-stake → states query → liquid-unstake → pause-pool blocks stakes → unpause restores → user burn → admin update-pool sets new fee rate.
- `module_names_test.go::TestIxoNames_FullScenario` — 12 subtests covering all names msgs: gov-create namespace (yoid self-register + twitter registrar-only) + gov-update-namespace → user RegisterName → resolve normalises case → by-namespace + by-owner queries → TransferName → non-owner transfer rejected → registrar RegisterNameByRegistrar → registrar UpdateNameByRegistrar → registrar SetNameStatus → length validation rejects.
- `module_token_test.go::TestIxoToken_FullScenario` — 10 subtests covering all token msgs: upload cw721+ixo1155 → gov-set entity NFT → gov-set token ixo1155 code → register iid → create-entity → CreateToken → MintToken → TransferToken → RetireToken → PauseToken (toggle) → StopToken → CancelToken-after-stop rejected.
- `module_claims_test.go::TestIxoClaims_RejectsCreateForUnknownEntity` — MsgCreateCollection against unknown entity DID is rejected (entity-not-found path). Full claims flow (submit-claim, evaluate-claim, etc.) requires deeper entity + cw20 setup; covered by keeper-level tests.

## Status

Full-flow tests for all 10 custom modules + IBC + wasm + upgrade + gov.
Run locally with a built `ixofoundation/ixo-blockchain:local` Docker
image, or via the nightly `interchaintest E2E` GitHub Actions matrix.
