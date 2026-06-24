# Changelog

All notable changes to the ixo Blockchain are documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).
The version suffix of the Go module path tracks the major chain-upgrade version
(e.g. `v7.x.y` → module `github.com/ixofoundation/ixo-blockchain/v8`).

## [v7.0.0] - "Opus"

On-chain upgrade name: **`Opus`**. Go module path bumped `v6 → v7`.

### Added
- **liquidstake — multi-pool**: the module moves from a single global pool to N
  governance-created pools, each with its own LST denom, validator whitelist,
  admin, fees, pause flag and proxy account. New `MsgCreatePool` / `MsgUpdatePool`
  / `MsgUpdateModuleParams` / `MsgSetPoolPaused` / `MsgSetModulePaused`.
- **x/names**: new governance-managed name-service module mapping human-readable
  handles to DIDs, with self-register and registrar-on-behalf flows, status
  lifecycle, reverse lookup, and wasm-whitelisted queries.
- **claims — team member budgets**: per-member periodic spend budgets on shared
  subscription collections (`MsgSetCollectionMembers` / `MsgRemoveCollectionMembers`).
- **claims — FLAGGED evaluation status**: non-terminal outcome letting evaluators
  defer a final call (e.g. AI oracle → human review); funds stay escrowed.
- **claims — dispute resolution + performance deposits**: target-role disputes,
  adjudicators authorized by DID key (`MsgAdjudicateDispute`), rolling per-agent
  deposits with a withdrawal lock, and configurable award/dismiss splits.
- **claims — `MsgUpdateCollectionQuota`**: admin update for a collection's quota,
  guarded so it can never be set below the current claim count.

### Changed
- **iid**: `MsgCreateIidDocument` is restricted to signer-owned account DIDs and
  wasm-contract DIDs; module-reserved namespaces (e.g. `did:ixo:entity:`) are
  rejected.
- Three-layer LST denom-collision protection across liquidstake, the bank mint
  restriction, and bonds.
- Go module path `github.com/ixofoundation/ixo-blockchain/v6 → /v7`.

### Migrations
- liquidstake: legacy single-pool state migrated to pool `zero` (`uzero` LST);
  existing delegations untouched.
- claims: store migration v3 → v4 stamps pre-v7 disputes as `DISMISSED`.
- New `x/names` store key added.

### Tooling
- Comprehensive multi-layer test suite: L1 (keeper unit), L2 (simulator),
  L3 (interchaintest E2E), with CI wiring.
- Legacy `MsgUpdateParams` proto re-added (decode-only, not in the Msg service)
  so the v7 RPC can render historical pre-v7 transactions for indexers.

[v7.0.0]: https://github.com/ixofoundation/ixo-blockchain/releases/tag/v7.0.0
