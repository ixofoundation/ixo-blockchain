# Messages

In this section we describe the processing of the liquidstake messages and the corresponding updates to state. All created/modified state objects are defined within the [state](./02_state.md) section.

The messages fall into three groups by who is allowed to sign them:

- **End-user** — anyone (or specifically the pool admin in the case of `MsgLiquidStake`).
- **Pool admin** — the address stored in the pool's `whitelist_admin_address` (or the chain governance address, except for `MsgUpdateWeightedRewardsReceivers` which is admin-only).
- **Governance** — only the chain governance module address (typically reached via `cosmos.gov.v1.MsgSubmitProposal` carrying the inner Msg).

## End-user operations

### MsgLiquidStake

`MsgLiquidStake` deposits native bond-denom tokens (e.g. `uixo`) into a pool's proxy account, delegates them in proportion to the pool's whitelisted validator weights, and mints the pool's LST denom 1:1 to the delegator.

```go
type MsgLiquidStake struct {
    DelegatorAddress string
    PoolId           string
    Amount           sdk.Coin
}
```

The field's descriptions is as follows:

- `delegator_address` - the signer. Must equal `Pool(pool_id).whitelist_admin_address`. Other addresses are rejected with `ErrRestrictedToWhitelistedAdminAddress`.
- `pool_id` - the pool to stake into. Must reference an existing pool (`ErrPoolNotFound` otherwise) and be a syntactically valid pool ID (`ErrInvalidPoolID` otherwise).
- `amount` - the deposit. Denom must equal the chain's bond denom (e.g. `uixo`). Amount must be at least `ModuleParams.min_liquid_stake_amount` (else `ErrLessThanMinLiquidStakeAmount`).

Additional pre-conditions:
- Neither `Pool.paused` nor `ModuleParams.module_paused` may be true.
- The pool must have at least one active liquid validator and the active set must hold at least 33.33% of the total weight.

### MsgLiquidUnstake

`MsgLiquidUnstake` burns the pool's LST denom and initiates unbonding from the pool's whitelisted validators back to the delegator over the chain's standard unbonding period.

```go
type MsgLiquidUnstake struct {
    DelegatorAddress string
    PoolId           string
    Amount           sdk.Coin
}
```

The field's descriptions is as follows:

- `delegator_address` - any holder of the LST denom.
- `pool_id` - the pool to unstake from. Validated; both `pool_id` and `amount.denom` must reference the same pool. `amount.denom` must equal `Pool(pool_id).liquid_bond_denom` (else `ErrPoolDenomMismatch`).
- `amount` - the LST coin to burn.

The unbonding amount in native tokens is `amount * NetAmount / TotalSupply * (1 - Pool.unstake_fee_rate)`. The amount is split across the pool's liquid validators in proportion to their current liquid-token holdings (inactive validators are drained first via `PrioritiseInactiveLiquidValidators`). For each validator the keeper performs an LSM tokenize → bank-send → redeem → undelegate sequence so the unbonding queue ends up owned by the user (not the pool's proxy account), letting the user retrieve their tokens through the normal Cosmos unbonding flow.

If the pool currently has no positive liquid token total but the proxy account has enough spendable balance, the keeper sends the requested amount immediately (returned as `unbonded_amount` in the response) instead of opening unbonding entries.

### MsgBurn

`MsgBurn` permanently destroys native `uixo` tokens. Unchanged from pre-v7; not pool-scoped.

```go
type MsgBurn struct {
    Burner string
    Amount sdk.Coin
}
```

The field's descriptions is as follows:

- `burner` - the signer.
- `amount` - the coin to burn. Denom must be `uixo`.

## Pool admin operations

### MsgUpdatePool

Updates a pool's mutable scalar/address fields (fee rates, fee account, admin address). `pool_id`, `liquid_bond_denom`, and `proxy_account_address` are immutable and not affected.

```go
type MsgUpdatePool struct {
    Authority             string
    PoolId                string
    UnstakeFeeRate        math.LegacyDec
    FeeAccountAddress     string
    AutocompoundFeeRate   math.LegacyDec
    WhitelistAdminAddress string
}
```

The field's descriptions is as follows:

- `authority` - signer. Must be either the chain governance address or the pool's current `whitelist_admin_address`. Other signers are rejected with `ErrorInvalidSigner`.
- `pool_id` - the pool to update. Must exist.
- `unstake_fee_rate`, `autocompound_fee_rate`, `fee_account_address`, `whitelist_admin_address` - replace the pool's current values in full. All four are required and validated.

### MsgUpdateWhitelistedValidators

Replaces the pool's validator whitelist. Sum of `target_weight` values across the supplied list must equal `10000` (`TotalValidatorWeight`). Each listed validator must exist on chain and be in `Bonded` status.

```go
type MsgUpdateWhitelistedValidators struct {
    Authority             string
    PoolId                string
    WhitelistedValidators []WhitelistedValidator
}
```

The field's descriptions is as follows:

- `authority` - signer. Either chain governance or the pool's current admin.
- `pool_id` - the pool whose whitelist is being replaced.
- `whitelisted_validators` - the new whitelist. Each entry validated for unique address, positive weight, on-chain existence, and bonded status. Sum of weights must equal `10000`.

### MsgUpdateWeightedRewardsReceivers

Replaces the pool's weighted-rewards receivers list (addresses that receive a portion of each autocompound's rewards before re-staking). **Pool admin only** — governance is not a valid signer here, matching pre-v7 behaviour.

```go
type MsgUpdateWeightedRewardsReceivers struct {
    Authority                string
    PoolId                   string
    WeightedRewardsReceivers []WeightedAddress
}
```

The field's descriptions is as follows:

- `authority` - signer. Must equal the pool's current `whitelist_admin_address`.
- `pool_id` - the pool whose receivers list is being replaced.
- `weighted_rewards_receivers` - the new list. Each entry has a valid bech32 address and a positive weight; sum of weights must not exceed `1`.

### MsgSetPoolPaused

Toggles a single pool's `Paused` flag. Other pools are unaffected.

```go
type MsgSetPoolPaused struct {
    Authority string
    PoolId    string
    IsPaused  bool
}
```

The field's descriptions is as follows:

- `authority` - signer. Either chain governance or the pool's current admin.
- `pool_id` - the pool to toggle.
- `is_paused` - the target value of `Pool.paused`.

## Governance operations

### MsgCreatePool

Registers a new liquid staking pool. The proxy account is derived deterministically from `pool_id`. The pool starts with empty `whitelisted_validators` and `weighted_rewards_receivers`, zero fee rates, and `paused: false`. The admin must subsequently call `MsgUpdateWhitelistedValidators` before any `MsgLiquidStake` will succeed (the active-weight quorum check requires a non-empty active set).

`MsgCreatePool` also (atomically) registers a `banktypes.Metadata` entry for the new LST denom — the standard "this denom is claimed" signal that other modules consult.

```go
type MsgCreatePool struct {
    Authority                string
    PoolId                   string
    LiquidBondDenom          string
    InitialAdminAddress      string
    InitialFeeAccountAddress string
}
```

The field's descriptions is as follows:

- `authority` - signer. Must be the chain governance module address.
- `pool_id` - identifier for the new pool. 2–16 chars, lowercase alphanumeric plus internal dashes (no leading/trailing dash). Must be globally unique; rejected with `ErrDuplicatePoolID` otherwise.
- `liquid_bond_denom` - the LST denom for the new pool. Must pass `sdk.ValidateDenom`. Rejected if it:
  - is the chain's bond denom (`uixo`),
  - matches the IBC voucher pattern `^ibc/[0-9A-F]{64}$`,
  - is already used by another pool (`ErrDuplicateLiquidBondDenom`),
  - has non-zero bank supply,
  - already has bank denom metadata registered.
- `initial_admin_address` - bech32 address that becomes the pool's `whitelist_admin_address`. Required (cannot be empty) — a pool with no admin would never accept a `MsgLiquidStake`.
- `initial_fee_account_address` - bech32 address that receives the autocompound fee. Required.

The response includes the derived `proxy_account_address` as a client convenience.

### MsgUpdateModuleParams

Updates the global `ModuleParams` (currently `min_liquid_stake_amount` and `module_paused`).

```go
type MsgUpdateModuleParams struct {
    Authority    string
    ModuleParams ModuleParams
}
```

The field's descriptions is as follows:

- `authority` - signer. Must be the chain governance module address.
- `module_params` - replaces the current ModuleParams in full. Validated for non-nil, non-negative `MinLiquidStakeAmount`.

### MsgSetModulePaused

Toggles the global `ModuleParams.module_paused` kill switch. When true, every pool halts regardless of its per-pool flag.

```go
type MsgSetModulePaused struct {
    Authority string
    IsPaused  bool
}
```

The field's descriptions is as follows:

- `authority` - signer. Must be the chain governance module address.
- `is_paused` - target value of `ModuleParams.module_paused`.
