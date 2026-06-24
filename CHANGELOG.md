# Changelog

All notable changes to the ixo Blockchain are documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).
The version suffix of the Go module path tracks the major chain-upgrade version
(e.g. `v8.x.y` → module `github.com/ixofoundation/ixo-blockchain/v8`).

## [v8.0.0] - "Alpha"

On-chain upgrade name: **`Alpha`**. Emergency security release, enacted on mainnet
via expedited governance proposal #482 at height 17871000. Go module path bumped
`v7 → v8`.

### Security
- **x/bonds disabled**: in response to a disclosed vulnerability (advisory
  `GHSA-w3rp-4cm2-4wgc`), the bonds module is disabled at the ante,
  message-server and module levels. All bonds messages are rejected; existing
  bond state is retained but no new state transitions are processed.
- **iid ante — `authz.MsgExec` recursion**: the IID ante resolves the effective
  signer through single- and double-nested `authz.MsgExec` delegation, so the
  "signer must control the message DID" check cannot be bypassed via delegation.
- **entity authorization**: the entity keeper binds the signing account to the
  entity's controller DID, and the entity ante decorator handles nested
  `authz.MsgExec`.
- **ICA host allow-list**: the Interchain Accounts host `AllowMessages` set is
  restricted to an ante-safe allow-list.

### Changed
- **claims — IID-ante scoping**: `MsgSubmitClaim`, `MsgEvaluateClaim`,
  `MsgCreateClaimAuthorization`, `MsgClaimIntent` and `MsgDisputeClaim` are no
  longer `IidTxMsg`; their `*_did` field is attribution only and authorization is
  enforced in the keeper, so delegated and module-account agent flows are not
  broken. `MsgAdjudicateDispute` remains `IidTxMsg` (the keeper already requires
  the signer to control the adjudicator DID).
- **x/token**: batch minting/transfer paths hardened.
- Go module path `github.com/ixofoundation/ixo-blockchain/v7 → /v8`.

### Migrations
- Upgrade handler `Alpha`: disables the bonds module and restricts the ICA host
  `AllowMessages` set to an ante-safe allow-list.

### Tooling
- Added end-to-end "no-ante" harnesses (`x/iid`, `x/entity`, `x/claims`),
  `x/token` batch tests, and a claims IID-ante membership regression guard.

[v8.0.0]: https://github.com/ixofoundation/ixo-blockchain/releases/tag/v8.0.0

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
