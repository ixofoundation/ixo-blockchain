# Events

In this section we describe the processing of the liquidstake events. Clients can subscribe to these to track pool lifecycle, stake/unstake activity, and per-epoch hook execution. All events are typed protobuf messages emitted via `EventManager.EmitTypedEvent`.

The five **per-tx and per-epoch events** carried over from pre-v7 (`LiquidStakeEvent`, `LiquidUnstakeEvent`, `AddLiquidValidatorEvent`, `RebalancedLiquidStakeEvent`, `AutocompoundStakingRewardsEvent`) preserve their pre-v7 field tag numbers exactly, so v7 codegen decodes historical v6 events cleanly with `pool_id` simply showing as the empty string. Where individual fields' types changed (`string` → `cosmos.base.v1beta1.Coin`, `string` → `google.protobuf.Timestamp`, `string` → `uint32`), the wire encoding for that tag changes — but indexers consume events as `(key, value)` attribute pairs derived from field names, not from raw wire bytes, so the impact is contained to the attribute *value format*: `"100000000uixo"` becomes `{"denom":"uixo","amount":"100000000"}`. Indexers should fork on attribute presence (and value format) when handling v6 vs v7 events.

### PoolCreatedEvent

Emitted when a new pool is registered through [MsgCreatePool](./03_messages.md#msgcreatepool).

```go
type PoolCreatedEvent struct {
    PoolId    string  // tag 1
    Pool      *Pool   // tag 2
    Authority string  // tag 3
}
```

The field's descriptions is as follows:

- `pool_id` - the new pool's identifier.
- `pool` - the full [Pool](./02_state.md#pool) record at creation time.
- `authority` - the signer of the gov proposal that created the pool (i.e., the chain governance module address).

### PoolUpdatedEvent

Emitted when an existing pool's configuration changes via any of [MsgUpdatePool](./03_messages.md#msgupdatepool), [MsgUpdateWhitelistedValidators](./03_messages.md#msgupdatewhitelistedvalidators), [MsgUpdateWeightedRewardsReceivers](./03_messages.md#msgupdateweightedrewardsreceivers), or [MsgSetPoolPaused](./03_messages.md#msgsetpoolpaused).

```go
type PoolUpdatedEvent struct {
    PoolId    string  // tag 1
    Pool      *Pool   // tag 2
    Authority string  // tag 3
}
```

The field's descriptions is as follows:

- `pool_id` - identifier of the pool that was updated.
- `pool` - the full updated [Pool](./02_state.md#pool) record.
- `authority` - the signer of the update message (chain governance address or the pool's admin, depending on the message).

### ModuleParamsUpdatedEvent

Emitted when the global `ModuleParams` change via [MsgUpdateModuleParams](./03_messages.md#msgupdatemoduleparams) or [MsgSetModulePaused](./03_messages.md#msgsetmodulepaused).

```go
type ModuleParamsUpdatedEvent struct {
    ModuleParams *ModuleParams  // tag 1
    Authority    string         // tag 2
}
```

The field's descriptions is as follows:

- `module_params` - the full updated [ModuleParams](./02_state.md#moduleparams) record.
- `authority` - the signer (chain governance module address).

### LiquidStakeEvent

Emitted on every successful [MsgLiquidStake](./03_messages.md#msgliquidstake).

```go
type LiquidStakeEvent struct {
    Delegator          string    // tag 1 (preserved from v6)
    LiquidAmount       sdk.Coin  // tag 2 (preserved from v6; type upgraded from string)
    StkIxoMintedAmount sdk.Coin  // tag 3 (preserved from v6; type upgraded from string)
    PoolId             string    // tag 4 (NEW in v7, appended)
}
```

The field's descriptions is as follows:

- `delegator` - the signer / pool admin.
- `liquid_amount` - the deposited bond-denom coin (e.g. `{denom: "uixo", amount: "100000000"}`). v6 emitted this as the formatted string `"100000000uixo"` at the same tag number; v7 onwards emits a typed Coin.
- `stk_ixo_minted_amount` - the LST coin minted to the delegator. Same v6→v7 type upgrade as `liquid_amount`.
- `pool_id` - the pool staked into. Empty for v6 events; an indexer should default to `"zero"` (the migrated legacy pool) when absent.

### LiquidUnstakeEvent

Emitted on every successful [MsgLiquidUnstake](./03_messages.md#msgliquidunstake).

```go
type LiquidUnstakeEvent struct {
    Delegator       string                   // tag 1 (preserved from v6)
    UnstakeAmount   sdk.Coin                 // tag 2 (preserved; type upgraded from string)
    UnbondingAmount sdk.Coin                 // tag 3 (preserved; type upgraded from string)
    UnbondedAmount  sdk.Coin                 // tag 4 (preserved; type upgraded from string)
    CompletionTime  time.Time                // tag 5 (preserved; type upgraded from RFC3339 string)
    PoolId          string                   // tag 6 (NEW in v7, appended)
}
```

The field's descriptions is as follows:

- `delegator` - the signer / LST holder.
- `unstake_amount` - the LST coin burned (`{denom: "uzero", amount: ...}`). v6 emitted this as `"<amount><LSTDenom>"` at the same tag.
- `unbonding_amount` - the bond-denom coin that has entered the standard unbonding queue (will be available at `completion_time`).
- `unbonded_amount` - the bond-denom coin returned immediately from the proxy account's spendable balance (the rare case where the pool has no live delegations but enough liquid balance). Usually `{denom: "uixo", amount: "0"}`.
- `completion_time` - the Cosmos `google.protobuf.Timestamp` at which the unbonding completes. Matches `MsgLiquidUnstakeResponse.completion_time`. Zero / unset if `unbonded_amount > 0` was returned immediately. v6 emitted this as an RFC3339-formatted string at the same tag.
- `pool_id` - the pool unstaked from. Empty for v6 events; default to `"zero"`.

### AddLiquidValidatorEvent

Emitted when a newly whitelisted validator is registered as a `LiquidValidator` for a pool. This happens during `BeginBlock` (via `UpdateLiquidValidatorSet`) the first time a whitelisted validator is observed in a queryable, active state.

```go
type AddLiquidValidatorEvent struct {
    Validator string  // tag 1 (preserved from v6)
    PoolId    string  // tag 2 (NEW in v7, appended)
}
```

The field's descriptions is as follows:

- `validator` - the bech32 validator operator address newly added to the pool's `LiquidValidators` set.
- `pool_id` - the pool that just registered the validator. Empty for v6 events; default to `"zero"`.

### RebalancedLiquidStakeEvent

Emitted by the rebalance epoch hook (`RebalanceEpoch`, default `"day"`) for each pool that performed at least one redelegation in this cycle.

```go
type RebalancedLiquidStakeEvent struct {
    Delegator             string  // tag 1 (preserved from v6)
    RedelegationCount     uint32  // tag 2 (preserved; type upgraded from string)
    RedelegationFailCount uint32  // tag 3 (preserved; type upgraded from string)
    PoolId                string  // tag 4 (NEW in v7, appended)
}
```

The field's descriptions is as follows:

- `delegator` - the pool's proxy account address (which performed the redelegations).
- `redelegation_count` - total redelegations attempted this cycle. v6 emitted this as a `strconv.Itoa`-formatted string at the same tag; v7 onwards is a typed `uint32`.
- `redelegation_fail_count` - subset of the above that failed (e.g. due to transitive-redelegation rejection). Failed entries are retried on the next epoch. Same v6→v7 type upgrade as `redelegation_count`.
- `pool_id` - the pool that rebalanced. Empty for v6 events; default to `"zero"`.

### AutocompoundStakingRewardsEvent

Emitted by the autocompound epoch hook (`AutocompoundEpoch`, default `"hour"`) for each pool that successfully completed an autocompound cycle (rewards withdrawal → fee deduction → weighted distribution → re-stake).

```go
type AutocompoundStakingRewardsEvent struct {
    Delegator             string    // tag 1 (preserved from v6)
    TotalAmount           sdk.Coin  // tag 2 (preserved; type upgraded from string)
    FeeAmount             sdk.Coin  // tag 3 (preserved; type upgraded from string)
    RedelegateAmount      sdk.Coin  // tag 4 (preserved; type upgraded from string)
    WeightedRewardsAmount sdk.Coin  // tag 5 (preserved; type upgraded from string)
    PoolId                string    // tag 6 (NEW in v7, appended)
}
```

The field's descriptions is as follows:

- `delegator` - the pool's proxy account address.
- `total_amount` - total bond-denom amount available for compounding this cycle (the proxy account's spendable balance after withdrawing rewards), as a typed Coin.
- `fee_amount` - portion sent to the pool's `fee_account_address` (= `total_amount * autocompound_fee_rate`).
- `redelegate_amount` - portion re-staked across the pool's active validators.
- `weighted_rewards_amount` - portion distributed across `weighted_rewards_receivers`.
- `pool_id` - the pool that auto-compounded. Empty for v6 events; default to `"zero"`.

The three component coin amounts (`fee_amount`, `redelegate_amount`, `weighted_rewards_amount`) sum to `total_amount` (modulo rounding).
