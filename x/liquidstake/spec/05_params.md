# Parameters

The liquidstake module distinguishes between **module-wide parameters** (single record, governance-managed via [MsgUpdateModuleParams](./03_messages.md#msgupdatemoduleparams) and [MsgSetModulePaused](./03_messages.md#msgsetmodulepaused)) and **per-pool fields** (managed by each pool's admin via the various `MsgUpdate*` messages — see [Messages](./03_messages.md)). It also relies on a few **module-wide constants** that are baked into the binary and not changeable through state.

## ModuleParams

Single record at KV slot `0x01`:

| Key                      | Type      | Default      | Notes                                                                                                                              |
| ------------------------ | --------- | ------------ | ---------------------------------------------------------------------------------------------------------------------------------- |
| `min_liquid_stake_amount` | `math.Int`| `1_000_000`  | Minimum bond-denom amount accepted by `MsgLiquidStake`. Applied identically across every pool. Default = 1 IXO.                    |
| `module_paused`           | `bool`    | `false`      | Global emergency kill switch. When `true`, every pool halts stake/unstake/autocompound/rebalance, overriding per-pool `paused`.    |

### `min_liquid_stake_amount`

Bound that protects against dust-amount stakes whose mint/unbond rounding would lose meaningful value. Applied uniformly so every pool inherits the same floor.

### `module_paused`

Reserved for chain-wide incidents (e.g. mid-upgrade migrations, suspected exploits). Routine pool pauses should use the per-pool `paused` flag instead.

## Pool fields

Per-pool fields are not "params" in the Cosmos SDK x/params sense — they are stored on each `Pool` record (KV prefix `0x10`). Each pool has its own settings, set at create time and modifiable thereafter:

| Field                        | Type                       | Mutability                                                          | Notes                                                                                                  |
| ---------------------------- | -------------------------- | ------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------------ |
| `pool_id`                    | `string`                   | immutable                                                           | Set at `MsgCreatePool`. 2–16 chars, lowercase alphanumeric + internal dashes.                          |
| `liquid_bond_denom`          | `string`                   | immutable                                                           | Set at `MsgCreatePool`. Globally unique. Cannot be the chain bond denom or in IBC namespace.           |
| `proxy_account_address`      | `string`                   | immutable                                                           | Derived from `pool_id`. The migrated `zero` pool uses the legacy pre-v7 address.                       |
| `whitelisted_validators`     | `[]WhitelistedValidator`   | gov OR pool admin via `MsgUpdateWhitelistedValidators`              | `target_weight` values must sum to exactly `10000`.                                                    |
| `unstake_fee_rate`           | `math.LegacyDec` `[0, 1]`  | gov OR pool admin via `MsgUpdatePool`                               | Default `0`.                                                                                           |
| `fee_account_address`        | `string`                   | gov OR pool admin via `MsgUpdatePool`                               | Receives the autocompound fee.                                                                        |
| `autocompound_fee_rate`      | `math.LegacyDec` `[0, 1]`  | gov OR pool admin via `MsgUpdatePool`                               | Default `0`. Applied to per-cycle compounded rewards.                                                  |
| `whitelist_admin_address`    | `string`                   | gov OR pool admin via `MsgUpdatePool`                               | Required (cannot be empty). Only address allowed to call `MsgLiquidStake`.                            |
| `paused`                     | `bool`                     | gov OR pool admin via `MsgSetPoolPaused`                            | Per-pool kill switch. Combines OR-wise with `ModuleParams.module_paused`.                              |
| `weighted_rewards_receivers` | `[]WeightedAddress`        | pool admin only via `MsgUpdateWeightedRewardsReceivers` (no gov)    | Sum of weights ≤ 1. Receives a fraction of each autocompound's rewards before re-staking.              |

## Module-wide constants

These live in `x/liquidstake/types/keys.go` and `x/liquidstake/types/pool.go`. They cannot be changed through state — only through a chain upgrade.

| Constant                              | Value                                                          | Notes                                                                                                       |
| ------------------------------------- | -------------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------- |
| `MinPoolIDLength` / `MaxPoolIDLength` | `2` / `16`                                                     | Bounds on `pool_id` length.                                                                                 |
| `TotalValidatorWeight`                | `10000`                                                        | The required sum of `WhitelistedValidator.target_weight` values within a pool.                              |
| `ActiveLiquidValidatorsWeightQuorum`  | `0.3333` (`math.LegacyDec`)                                    | Minimum fraction of `TotalValidatorWeight` that must be held by *active* validators for `LiquidStake` to succeed. |
| `RebalancingTrigger`                  | `0.001` (`math.LegacyDec`)                                     | Minimum fractional gap before the rebalance epoch will issue any redelegation.                              |
| `AutocompoundEpoch`                   | `"hour"` in production / `"2min"` for local testing            | Epoch identifier consumed for the autocompound hook. Pin to a faster epoch only on test chains.             |
| `RebalanceEpoch`                      | `"day"` in production / `"2min"` for local testing             | Epoch identifier consumed for the rebalance hook. Same caveat as above.                                     |
| `LegacyLiquidStakeProxyAcc`           | `address.Module("liquidstake", "-LiquidStakeProxyAcc")`        | The pre-v7 single-pool proxy account. Used by the v7 migration as the `zero` pool's `proxy_account_address`. |
| `PoolProxyAcc(poolID)`                | `address.Module("liquidstake", "-LiquidStakeProxyAcc-"+poolID)`| Derivation function for newly-created pools' proxy accounts.                                                |
