# Concepts

## Liquid staking

Liquid staking lets a delegator put native staking tokens to work earning rewards while still holding a transferable receipt token (the **liquid staking token**, or **LST**) that represents their claim on the underlying delegation. Unlike a raw `MsgDelegate` from `cosmos.staking`, the underlying delegation is held by a module-controlled proxy account; the LST is the user-facing handle.

Every LST is denominated by the chain's bank module (e.g. `uzero`, `ucarbon`) and is freely transferable like any other native coin. A holder can redeem their LST at any time via `MsgLiquidUnstake`, which burns the LST and unbonds the corresponding share of the proxy's delegations back to the holder over the chain's standard unbonding period.

## Pools

In v7+ the module supports an arbitrary number of independent **pools**. Each pool has:

- A unique, immutable `pool_id` (lowercase alphanumeric + dashes, 2–16 chars).
- A unique, immutable `liquid_bond_denom` (the LST denom, e.g. `uzero`).
- A unique proxy account address derived from the pool_id at creation time, which holds all delegations for the pool.
- A whitelisted validator set with target weights summing to exactly `10000`.
- An admin (`whitelist_admin_address`) authorised to operate the pool.
- An `unstake_fee_rate`, an `autocompound_fee_rate`, and a `fee_account_address` that receives the autocompound fee.
- A `weighted_rewards_receivers` list — addresses that receive a configured share of each autocompound's rewards before the remainder is re-staked.
- A per-pool `paused` flag.

Pools are created by governance via `MsgCreatePool`. Once created, the admin can update mutable fields without further governance involvement. Pool IDs and denoms are immutable after creation.

Each pool maintains an independent `NetAmountState` (supply, rate, etc.). LSTs from different pools are **not fungible** with each other — their values can and will diverge based on each pool's validator performance, fee structure, and timing.

## Mint and unstake rates

Within a pool:

- **Mint rate** is fixed at **1 native token = 1 LST**. The pool admin (only address allowed to call `LiquidStake`) deposits N native tokens into the proxy account and receives N LSTs.
- **Unstake rate** is `NetAmount / TotalSupply`, computed at the moment of unstake. As autocompounded rewards push `NetAmount` above the LST supply, the unstake rate appreciates above 1.0; slashing pushes it below.

Because `LiquidStake` is currently restricted to the pool's admin, this asymmetry has no dilution effect — there is only one staker per pool, so they always receive back proportionally what they contributed plus accrued rewards. If `LiquidStake` is ever opened to multiple stakers per pool, the mint formula must be revisited (a new staker would otherwise claim a pro-rata share of already-accrued rewards).

## Per-pool proxy account

Each pool's delegations are held by a derived module sub-account:

```
proxyAccount = address.Module("liquidstake", "-LiquidStakeProxyAcc-" + poolID)
```

This guarantees a unique 32-byte LSM-compatible address per pool. The address is computed at `MsgCreatePool` time and stored in `Pool.proxy_account_address` so the keeper never needs to re-derive it.

The legacy `zero` pool — migrated from pre-v7 state — is the one exception: its proxy account is the original pre-v7 `LiquidStakeProxyAcc = address.Module("liquidstake", "-LiquidStakeProxyAcc")` (without the pool-id suffix), so existing on-chain delegations from the original deployment remain accessible without any state migration of staking-module records.

## Whitelisted validators and weights

Each pool keeps its own `WhitelistedValidator[]` list. Each entry has a `validator_address` and a `target_weight`; weights across the list must sum to exactly `10000` (`TotalValidatorWeight`).

`LiquidStake` only succeeds if the **active subset** of whitelisted validators (those that exist on-chain, are not tombstoned, and are bonded) collectively hold at least `33.33%` (`ActiveLiquidValidatorsWeightQuorum`) of the total weight. This prevents staking when the whitelist is essentially empty due to validator removal or tombstoning.

Per-validator delegations from the proxy account are kept in lockstep with the target weights by the **rebalance** epoch (see below).

## Autocompounding and rebalancing epochs

The module subscribes to two epoch identifiers from `x/epochs`:

- `AutocompoundEpoch` — every tick withdraws all unwithdrawn rewards from each pool's proxy account, deducts `autocompound_fee_rate` (sent to `fee_account_address`), distributes the configured `weighted_rewards_receivers` portions, and re-stakes the remainder.
- `RebalanceEpoch` — every tick reshuffles delegations among each pool's whitelisted validators so their actual stake matches the target weights. Only triggers if any single validator's delta exceeds `RebalancingTrigger` (0.1% of total liquid tokens) — avoids churn from small reward fluctuations.

Production identifiers are `"hour"` and `"day"` respectively. They can be temporarily collapsed to a faster identifier (e.g. `"2min"`) for local testing — see `x/liquidstake/types/keys.go`.

The two hooks fire **per pool** during `BeforeEpochStart`. A pool whose `paused` flag is true (or where the global `module_paused` is true) is skipped.

## Pause flags (per-pool and global)

Two independent pause toggles:

- **`Pool.paused`** — per-pool. Halts that one pool's `LiquidStake`, `LiquidUnstake`, autocompounding, and rebalancing. Other pools are unaffected. Settable by gov or pool admin via `MsgSetPoolPaused`.
- **`ModuleParams.module_paused`** — global emergency kill switch. When true, **every pool** halts regardless of its per-pool flag. Settable by gov only via `MsgSetModulePaused`.

Either being true halts the affected pool. Use the per-pool flag for routine ops (e.g. while changing validator set during instability); reserve the global flag for chain-wide incidents.

## Denom collision defences

Three layers prevent another module from minting tokens of an LST denom (which would corrupt that pool's `NetAmountState.stkixoTotalSupply` and break the unstake rate):

1. **`registerPool` create-time guards** — at pool registration, the keeper rejects:
   - Any denom matching the IBC voucher pattern `^ibc/[0-9A-F]{64}$`
   - The chain bond denom (e.g. `uixo`)
   - Any denom with non-zero bank supply
   - Any denom that already has bank metadata registered

2. **Bank denom metadata claim** — `RegisterLSTDenomMetadata` is called from `registerPool` (and from the v7 migration for the `zero` pool). This writes a `banktypes.Metadata` record claiming the denom in the bank module. Any other module that consults `bank.HasDenomMetaData` before issuing tokens (e.g. `bonds.CreateBond`, after the `v7` chain change) will see the claim and refuse. This closes the gap where a freshly-created pool has zero supply and the bank-supply check above wouldn't fire.

3. **Bank `MintCoinsRestriction`** (installed in `app/keepers/keepers.go`) — final, enduring backstop. For every `bank.MintCoins` call across the entire chain: if the denom matches a registered pool's `liquid_bond_denom` AND the calling context lacks the liquidstake authorisation sentinel (set only by liquidstake's own `LiquidStake` handler), the mint is rejected. Catches any future module that bypasses checks 1 and 2. Beware: in some EndBlocker contexts (notably bonds' batch processor) a rejection here will halt the chain — so layers 1 and 2 are the front-door checks that should normally fire first.

## v7 upgrade migration

The `v7` chain upgrade (handler in `app/upgrades/v7/`) reshapes the pre-existing single-pool state into the multi-pool layout:

1. Reads the legacy `Params` record from `0x01`.
2. Constructs `ModuleParams` (carrying over `min_liquid_stake_amount` and `module_paused`).
3. Constructs the `zero` Pool from the legacy fields, with `proxy_account_address` set to the **pre-v7** `LiquidStakeProxyAcc` bech32 — this preserves all existing delegations untouched.
4. Re-keys every legacy `LiquidValidator` record from prefix `0x02 + valAddr` to `0x11 + lp("zero") + valAddr`.
5. Calls `RegisterLSTDenomMetadata("zero", "uzero")` so the legacy denom is also claimed in bank metadata going forward.
6. Writes the new `ModuleParams` record to `0x01` (overwriting the legacy `Params`).

The migration is idempotent in practice (the `upgrade` keeper guarantees single execution) and skips the pool-construction path entirely if no legacy `LiquidBondDenom` was ever set — useful for chains that have the binary but never initialised liquidstake state.

If at the time of upgrade `uzero` had non-zero supply on chain, the migration's `RegisterLSTDenomMetadata` call still claims the metadata cleanly because metadata isn't currently set for `uzero` (mainnet check confirmed empty).
