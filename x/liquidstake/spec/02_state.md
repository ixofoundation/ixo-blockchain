# State

## Storage layout

The liquidstake module owns one KV store with the following key prefixes:

| Prefix | Key | Value |
| --- | --- | --- |
| `0x01` | `ModuleParamsKey` (single) | `ProtocolBuffer(ModuleParams)` |
| `0x10` | `PoolPrefix + lp(poolID)` | `ProtocolBuffer(Pool)` |
| `0x11` | `LiquidValidatorPrefix + lp(poolID) + lp(valOperatorAddress)` | `ProtocolBuffer(LiquidValidator)` |

`lp(...)` denotes a length-prefixed byte slice (`address.MustLengthPrefix`). Length-prefixing the `pool_id` segment guarantees that no pool's storage key is a byte-prefix of another's — essential for safe iteration with `KVStorePrefixIterator(store, prefix)`.

Prefix `0x02` was the pre-v7 `LiquidValidator` prefix. The v7 upgrade migration deletes any remaining `0x02` records and re-keys them under `0x11`. The byte is reserved (never reused) so historical state can always be unambiguously identified.

The `NetAmountState` is **never persisted**; it is recomputed on demand from the pool's proxy-account state, the bank module's supply, and the staking module's delegation data.

# Types

### ModuleParams

Global, module-wide parameters. Exactly one record exists; written by the genesis init or `MsgUpdateModuleParams` / `MsgSetModulePaused`.

```go
type ModuleParams struct {
    MinLiquidStakeAmount math.Int
    ModulePaused         bool
}
```

The field's descriptions is as follows:

- `min_liquid_stake_amount` - minimum amount (in the chain bond denom) that a single `MsgLiquidStake` may stake. Applied identically across every pool. Defaults to `1000000` (1 IXO) so dust amounts don't waste gas with truncation losses.
- `module_paused` - global emergency kill switch. When true, every pool halts its `LiquidStake` / `LiquidUnstake` / autocompounding / rebalancing — overriding each pool's own `paused` flag. Used for chain-wide migrations or critical incidents.

### Pool

A single liquid staking pool. Each pool is independent — its own LST denom, validator whitelist, fees, admin, paused flag, and proxy account. Pools' LSTs are not fungible with each other.

```go
type Pool struct {
    PoolId                   string
    LiquidBondDenom          string
    ProxyAccountAddress      string
    WhitelistedValidators    []WhitelistedValidator
    UnstakeFeeRate           math.LegacyDec
    FeeAccountAddress        string
    AutocompoundFeeRate      math.LegacyDec
    WhitelistAdminAddress    string
    Paused                   bool
    WeightedRewardsReceivers []WeightedAddress
}
```

The field's descriptions is as follows:

- `pool_id` - the unique, immutable identifier for the pool (lowercase alphanumeric + dashes, 2–16 chars). Used in storage keys, message routing, event emission, and proxy account derivation.
- `liquid_bond_denom` - the LST denom (e.g. `uzero`, `ucarbon`). Must be globally unique across pools and not collide with any other denom in active bank use. Immutable.
- `proxy_account_address` - bech32-encoded address of the per-pool delegation proxy account. All delegations, redelegations, unbondings, and reward withdrawals for this pool flow through this account. Derived deterministically from `pool_id` at creation; immutable. The migrated `zero` pool exceptionally stores the **pre-v7** `LiquidStakeProxyAcc` to preserve existing delegations.
- `whitelisted_validators` - the validators eligible for delegation from this pool. `target_weight` values across the list must sum to exactly `10000` whenever the list is set via `MsgUpdateWhitelistedValidators`.
- `unstake_fee_rate` - fraction (`math.LegacyDec`, range `[0, 1]`) deducted from the unbonding amount on `MsgLiquidUnstake`. The deducted fraction effectively boosts the pool's `NetAmount` for remaining holders.
- `fee_account_address` - bech32 address that accumulates `autocompound_fee_rate` of each autocompound's rewards.
- `autocompound_fee_rate` - fraction (`math.LegacyDec`, range `[0, 1]`) of accrued staking rewards taken as a protocol fee on each autocompound and sent to `fee_account_address`.
- `whitelist_admin_address` - bech32 address authorised to update mutable pool fields (whitelisted validators, weighted receivers, paused, fees, etc.) and the only address allowed to call `MsgLiquidStake` against this pool. Required at create time (cannot be empty).
- `paused` - per-pool safety toggle. When true, this pool's stake/unstake, autocompounding, and rebalancing are halted; other pools unaffected. Settable by gov or pool admin.
- `weighted_rewards_receivers` - list of [WeightedAddress](#weightedaddress) entries that receive a fraction of each autocompound's rewards (after the autocompound fee is taken, before the remainder is re-staked). Sum of weights must not exceed `1`.

### WhitelistedValidator

```go
type WhitelistedValidator struct {
    ValidatorAddress string
    TargetWeight     math.Int
}
```

The field's descriptions is as follows:

- `validator_address` - bech32-encoded validator operator address (`ixovaloper...`).
- `target_weight` - the validator's share of the pool's total stake. Sum across the pool's list must equal `10000` (`TotalValidatorWeight`).

### WeightedAddress

```go
type WeightedAddress struct {
    Address string
    Weight  math.LegacyDec
}
```

The field's descriptions is as follows:

- `address` - bech32 address that receives this fraction of the pool's autocompound rewards.
- `weight` - fraction (`math.LegacyDec`, range `(0, 1]`) of the pool's per-cycle distributable rewards (after the autocompound fee) sent to this address. Sum of weights across the receivers list must not exceed `1`; the unused fraction is restaked.

### LiquidValidator

The persisted record for a validator currently delegated to from a pool. One record per `(pool_id, validator_address)` combination.

```go
type LiquidValidator struct {
    OperatorAddress string
}
```

The field's descriptions is as follows:

- `operator_address` - bech32-encoded validator operator address. The pool ownership is implicit in the storage key (`0x11 + lp(poolID) + lp(valAddr)`).

### LiquidValidatorState

Returned by the `LiquidValidators` query. Combines the persisted record with live state computed at query time.

```go
type LiquidValidatorState struct {
    OperatorAddress string
    Weight          math.Int
    Status          ValidatorStatus
    DelShares       math.LegacyDec
    LiquidTokens    math.Int
}
```

The field's descriptions is as follows:

- `operator_address` - bech32 validator operator address.
- `weight` - the validator's `target_weight` from the pool's whitelist if it is currently active, otherwise `0`.
- `status` - a [ValidatorStatus](#validatorstatus) enum value.
- `del_shares` - the proxy account's delegation shares to this validator (`math.LegacyDec`).
- `liquid_tokens` - the slashing-applied token-equivalent of `del_shares` for this validator (`math.Int`).

### NetAmountState

Per-pool live-computed state. Read by the `States` query and used internally to compute mint/burn rates. Never persisted.

```go
type NetAmountState struct {
    StakeRate              math.LegacyDec
    UnstakeRate            math.LegacyDec
    StkixoTotalSupply      math.Int
    NetAmount              math.LegacyDec
    TotalDelShares         math.LegacyDec
    TotalLiquidTokens      math.Int
    TotalRemainingRewards  math.LegacyDec
    TotalUnbondingBalance  math.Int
    ProxyAccBalance        math.Int
}
```

The field's descriptions is as follows:

- `stake_rate` - LSTs minted per 1 native token. Always `1.0` by design (see [Mint and unstake rates](01_concepts.md#mint-and-unstake-rates)).
- `unstake_rate` - native tokens returned per 1 LST burned, equals `net_amount / stkixo_total_supply`. Diverges from `1.0` as rewards accrue or slashing occurs.
- `stkixo_total_supply` - total supply of the pool's `liquid_bond_denom`, queried from the bank module.
- `net_amount` - `total_liquid_tokens + total_unbonding_balance`. The pool's redeemable backing.
- `total_del_shares` - sum of delegation shares the proxy account holds across all of the pool's liquid validators.
- `total_liquid_tokens` - slashing-applied token-equivalent of `total_del_shares`.
- `total_remaining_rewards` - sum of unwithdrawn staking rewards owed to the proxy account.
- `total_unbonding_balance` - sum of unbonding amounts (slashing applied) for the proxy account.
- `proxy_acc_balance` - spendable bond-denom balance currently in the proxy account (rewards withdrawn but not yet redelegated, or freshly received deposits before delegation).

### ValidatorStatus

```go
const (
    ValidatorStatusUnspecified ValidatorStatus = 0
    ValidatorStatusActive      ValidatorStatus = 1
    ValidatorStatusInactive    ValidatorStatus = 2
)
```

A liquid validator is `Active` when:

- It is in the pool's `whitelisted_validators` list, AND
- The Cosmos SDK staking keeper recognises it (it exists, is bonded/unbonding/unbonded, has non-nil delegation shares and tokens, and its exchange rate is valid), AND
- It is not tombstoned by the slashing module.

Otherwise it is `Inactive`. Inactive validators in a pool's set are drained first when `MsgLiquidUnstake` runs (see `PrioritiseInactiveLiquidValidators` in the keeper).

### Legacy Params (deprecated)

The pre-v7 single-pool layout used a single `Params` record at `0x01`:

```go
// DEPRECATED: kept so the v7 migration can unmarshal pre-upgrade state.
type Params struct {
    LiquidBondDenom          string
    WhitelistedValidators    []WhitelistedValidator
    UnstakeFeeRate           math.LegacyDec
    MinLiquidStakeAmount     math.Int
    FeeAccountAddress        string
    AutocompoundFeeRate      math.LegacyDec
    WhitelistAdminAddress    string
    ModulePaused             bool
    WeightedRewardsReceivers []WeightedAddress
}
```

The proto definition is retained with `option deprecated = true` so the v7 migration handler can deserialise the pre-upgrade KV slot and project its fields into the new `ModuleParams` (`MinLiquidStakeAmount`, `ModulePaused`) and `Pool` (`zero`) records. After the migration runs, KV slot `0x01` holds `ModuleParams` and the `Params` type is unused by application code.
