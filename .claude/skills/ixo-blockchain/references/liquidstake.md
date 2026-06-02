# Liquidstake module — `x/liquidstake`

**Proto package:** `ixo.liquidstake.v1beta1` · **TS typeUrl prefix:** `/ixo.liquidstake.v1beta1.` · **CLI:** `ixod tx liquidstake …` (no `ixod query liquidstake …` CLI — queries only via gRPC/REST)

> **SDK WARNING (read before writing any TS):** the published `@ixo/impactxclient-sdk` (latest checked: `2.4.10`) is generated from the **pre-v7 single-pool proto** and does NOT match this v7 chain. Concretely, in the published SDK: `MsgLiquidStake`/`MsgLiquidUnstake` have only `{ delegatorAddress, amount }` and **no `poolId`**; there is a `MsgUpdateParams` (now a decode-only legacy type on chain); and `MsgCreatePool`, `MsgUpdateModuleParams`, `MsgUpdatePool`, `MsgSetPoolPaused` are **absent**. Until the SDK is regenerated, the `ixo.liquidstake.v1beta1.Msg<Name>.fromPartial(...)` helpers below for those messages do not exist and pool-scoped messages will be missing the `pool_id` field. Treat the proto field names/types in this doc (taken from the v7 chain `tx.proto`) as the source of truth and construct the `Any` manually if your installed SDK is stale.

## Purpose
Liquid staking lets a delegator put native bond-denom tokens (`uixo`) to work earning staking rewards while holding a freely transferable derivative — the **liquid staking token (LST)**, denominated in the bank module (e.g. `uzero`). The underlying delegation is held by a module-controlled per-pool proxy account; the LST is the user-facing handle and is redeemed via `MsgLiquidUnstake`, which burns it and unbonds the proportional share back to the holder over the chain's standard unbonding period. v7+ supports an arbitrary number of independent **pools**, each with its own LST denom, validator whitelist, fees and admin; LSTs from different pools are not fungible. Design is pSTAKE/Stride-style.

## Concepts & state
- **Pool** (`liquidstake.proto` `Pool`): one independent liquid-staking instance. Immutable `pool_id` (lowercase alphanumeric + dashes, 2–16 chars), immutable `liquid_bond_denom` (the LST denom), immutable `proxy_account_address` (derived from `pool_id`), `whitelisted_validators`, `unstake_fee_rate`, `fee_account_address`, `autocompound_fee_rate`, `whitelist_admin_address`, `paused`, `weighted_rewards_receivers`. Each pool keeps an independent supply/rate.
- **Proxy account**: per-pool module sub-account holding all of the pool's delegations/unbondings/reward withdrawals. Derived `address.Module("liquidstake", "-LiquidStakeProxyAcc-"+poolID)`; the migrated `zero` pool keeps the pre-v7 `LiquidStakeProxyAcc` (no `-poolID` suffix) so existing delegations survive untouched.
- **Liquid bond denom / derivative (LST)**: the bank coin minted 1:1 on stake (e.g. `uzero`). Globally unique per pool. Mint rate is fixed at 1 native = 1 LST; unstake rate = `NetAmount / TotalSupply` (appreciates as rewards autocompound, drops on slashing).
- **WhitelistedValidator** (`liquidstake.proto`): `{ validator_address, target_weight }`. `target_weight` values across the pool's list must sum to exactly `10000` (`TotalValidatorWeight`). `LiquidStake` only succeeds if the active subset (exists on-chain, bonded, not tombstoned) holds ≥ 33.33% (`ActiveLiquidValidatorsWeightQuorum`) of total weight.
- **WeightedAddress** (`liquidstake.proto`): `{ address, weight }` — receivers of a fraction of each autocompound's rewards before the remainder is restaked. Sum of `weight` must not exceed `1`.
- **ModuleParams** (`liquidstake.proto`): global, single record — `{ min_liquid_stake_amount (math.Int), module_paused (bool) }`. `min_liquid_stake_amount` defaults `1000000`; `module_paused` is a global kill switch overriding every pool's own flag.
- **LiquidValidator / LiquidValidatorState** (`liquidstake.proto`): persisted record is `{ operator_address }` keyed by `(pool_id, valAddr)`; the query-time state adds `weight (math.Int)`, `status (ValidatorStatus)`, `del_shares (math.LegacyDec)`, `liquid_tokens (math.Int)`.
- **NetAmountState** (`liquidstake.proto`): per-pool, **never persisted** — recomputed each time from proxy-account/bank/staking state. Holds `stake_rate`, `unstake_rate`, `stkixo_total_supply`, `net_amount`, `total_del_shares`, `total_liquid_tokens`, `total_remaining_rewards`, `total_unbonding_balance`, `proxy_acc_balance`.
- **Rebalancing & autocompounding**: two `x/epochs` hooks fire per pool (production ids `"hour"` and `"day"`). Autocompound withdraws rewards, takes `autocompound_fee_rate` to `fee_account_address`, pays `weighted_rewards_receivers`, restakes the rest. Rebalance reshuffles delegations toward target weights (only if a validator's delta exceeds `RebalancingTrigger`, 0.1%). A paused pool (per-pool or global) is skipped.
- **Pause flags**: `Pool.paused` (per-pool; gov or pool admin) and `ModuleParams.module_paused` (global; gov only). Either being true halts the affected pool's stake/unstake/autocompound/rebalance.

## Messages
Source order from `tx.proto` service `Msg`. All ten messages are registered in the `Msg` service. Pool-scoped messages carry `pool_id`. `MsgStakeToLP` does **not exist** in this module (no such RPC or message type). `MsgUpdateParams` exists in `tx.proto` only as a legacy decode-only type (no RPC, no handler — see "Module gotchas").

### MsgLiquidStake
- **Purpose:** Deposit native bond-denom tokens into a pool's proxy account, delegate them per the pool's validator weights, and mint the pool's LST 1:1 to the delegator.
- **Signer / auth:** `delegator_address` signs. **Must equal `Pool(pool_id).whitelist_admin_address`** — `LiquidStake` is restricted to the pool admin (else `ErrRestrictedToWhitelistedAdminAddress`).
- **Fields:**

| Field | Type | Req | Description |
|---|---|---|---|
| delegator_address | string (`cosmos.AddressString`) | yes | Signer; must equal the pool's `whitelist_admin_address` |
| pool_id | string | yes | Pool to stake into; selects which LST is minted |
| amount | `cosmos.base.v1beta1.Coin` (non-nullable) | yes | Native staking coin to stake; denom must equal the chain bond denom (`uixo`) |

- **CLI:** `ixod tx liquidstake liquid-stake [pool-id] [amount] [flags]` — `delegator_address` is set from the `--from` key.
- **TS SDK:**
  ```ts
  const msg = {
    typeUrl: '/ixo.liquidstake.v1beta1.MsgLiquidStake',
    value: ixo.liquidstake.v1beta1.MsgLiquidStake.fromPartial({
      delegatorAddress: adminAddress,
      poolId: 'zero',
      amount: { denom: 'uixo', amount: '1000000' },
    }),
  };
  ```
  Note: the published SDK's `MsgLiquidStake` has no `poolId` field (see SDK warning).
- **Gotchas:** `ValidateBasic` rejects an invalid bech32 `delegator_address`, an invalid `pool_id` (`ErrInvalidPoolID` — must pass `ValidatePoolID`), a zero `amount`, or an invalid `amount`. Server-side: pool must exist (`ErrPoolNotFound`); amount must be ≥ `ModuleParams.min_liquid_stake_amount` (`ErrLessThanMinLiquidStakeAmount`); neither `Pool.paused` nor `ModuleParams.module_paused` may be true; pool must have an active validator set ≥ 33.33% weight (`ErrActiveLiquidValidatorsWeightQuorumNotReached`).

### MsgLiquidUnstake
- **Purpose:** Burn the pool's LST and initiate unbonding from the pool's validators back to the delegator over the unbonding period (or immediate send if the proxy has spendable balance and no positive liquid tokens).
- **Signer / auth:** `delegator_address` signs. Any holder of the LST denom (not admin-restricted).
- **Fields:**

| Field | Type | Req | Description |
|---|---|---|---|
| delegator_address | string (`cosmos.AddressString`) | yes | Signer; any LST holder |
| pool_id | string | yes | Pool to unstake from; must correspond to `amount.denom` |
| amount | `cosmos.base.v1beta1.Coin` (non-nullable) | yes | LST coin to burn; denom must equal `Pool(pool_id).liquid_bond_denom` |

- **Response:** `MsgLiquidUnstakeResponse { completion_time: google.protobuf.Timestamp }`.
- **CLI:** `ixod tx liquidstake liquid-unstake [pool-id] [amount] [flags]` — `delegator_address` set from `--from`.
- **TS SDK:**
  ```ts
  const msg = {
    typeUrl: '/ixo.liquidstake.v1beta1.MsgLiquidUnstake',
    value: ixo.liquidstake.v1beta1.MsgLiquidUnstake.fromPartial({
      delegatorAddress: holderAddress,
      poolId: 'zero',
      amount: { denom: 'uzero', amount: '1000000' },
    }),
  };
  ```
  Note: the published SDK's `MsgLiquidUnstake` has no `poolId` field (see SDK warning).
- **Gotchas:** `ValidateBasic` rejects invalid bech32, invalid `pool_id` (`ErrInvalidPoolID`), zero/invalid `amount`. Server-side: `amount.denom` must equal the pool's `liquid_bond_denom` (`ErrPoolDenomMismatch`); pool/module must not be paused; if neither liquid tokens nor proxy balance suffice, fails with `ErrInsufficientProxyAccBalance`. Unbonding native amount = `amount * NetAmount / TotalSupply * (1 - unstake_fee_rate)`; inactive validators are drained first.

### MsgCreatePool
- **Purpose:** Register a new liquid-staking pool (own LST denom, derived proxy account, admin, fee config). Also atomically registers bank `Metadata` claiming the LST denom.
- **Signer / auth:** `authority` signs. **Governance authority only** (must equal the gov module address; `ErrorInvalidSigner` otherwise).
- **Fields:**

| Field | Type | Req | Description |
|---|---|---|---|
| authority | string (`cosmos.AddressString`) | yes | Signer; must be the gov module address |
| pool_id | string | yes | Immutable id; lowercase alphanumeric + internal dashes, 2–16 chars, globally unique |
| liquid_bond_denom | string | yes | LST denom; must pass `sdk.ValidateDenom` and be globally unique |
| initial_admin_address | string (`cosmos.AddressString`) | yes | Becomes the pool's `whitelist_admin_address`; cannot be empty |
| initial_fee_account_address | string (`cosmos.AddressString`) | yes | Becomes `fee_account_address`; cannot be empty |

- **Response:** `MsgCreatePoolResponse { proxy_account_address: string }` (the derived address, returned for convenience).
- **CLI:** No CLI command (governance-only; submit via `cosmos.gov.v1.MsgSubmitProposal` carrying this Msg).
- **TS SDK:**
  ```ts
  const msg = {
    typeUrl: '/ixo.liquidstake.v1beta1.MsgCreatePool',
    value: ixo.liquidstake.v1beta1.MsgCreatePool.fromPartial({
      authority: govModuleAddress,
      poolId: 'qi',
      liquidBondDenom: 'uqi',
      initialAdminAddress: adminAddress,
      initialFeeAccountAddress: feeAddress,
    }),
  };
  ```
  Note: absent from the published SDK (see SDK warning).
- **Gotchas:** `ValidateBasic` checks bech32 `authority`, `pool_id` (`ErrInvalidPoolID`), `liquid_bond_denom` (`ErrInvalidLiquidBondDenom`), and non-empty admin/fee addresses. Server-side `registerPool` rejects a `liquid_bond_denom` that is the chain bond denom (`uixo`), matches the IBC voucher pattern `^ibc/[0-9A-F]{64}$`, duplicates another pool's denom (`ErrDuplicateLiquidBondDenom`), already has non-zero bank supply, or already has bank metadata; duplicate `pool_id` → `ErrDuplicatePoolID`. New pools start with empty whitelist — `MsgUpdateWhitelistedValidators` is required before staking succeeds.

### MsgUpdateModuleParams
- **Purpose:** Replace the global `ModuleParams` (min stake amount + global pause flag) in full.
- **Signer / auth:** `authority` signs. **Governance authority only.**
- **Fields:**

| Field | Type | Req | Description |
|---|---|---|---|
| authority | string (`cosmos.AddressString`) | yes | Signer; must be the gov module address |
| module_params | `ModuleParams` (non-nullable) | yes | Replacement params: `{ min_liquid_stake_amount (math.Int), module_paused (bool) }` |

- **CLI:** No CLI command (governance-only; submit via gov proposal).
- **TS SDK:**
  ```ts
  const msg = {
    typeUrl: '/ixo.liquidstake.v1beta1.MsgUpdateModuleParams',
    value: ixo.liquidstake.v1beta1.MsgUpdateModuleParams.fromPartial({
      authority: govModuleAddress,
      moduleParams: { minLiquidStakeAmount: '1000000', modulePaused: false },
    }),
  };
  ```
  Note: absent from the published SDK (see SDK warning).
- **Gotchas:** `ValidateBasic` checks bech32 `authority` then `ModuleParams.Validate()` (non-nil, non-negative `MinLiquidStakeAmount`). `module_params` replaces the record wholesale — supply both fields.

### MsgUpdatePool
- **Purpose:** Update a pool's mutable scalar/address fields (fee rates, fee account, admin). `pool_id`, `liquid_bond_denom`, `proxy_account_address` are immutable and untouched. Whitelist, weighted receivers, and paused flag have their own messages.
- **Signer / auth:** `authority` signs. **Governance OR the pool's current `whitelist_admin_address`** (`ErrorInvalidSigner` otherwise).
- **Fields:**

| Field | Type | Req | Description |
|---|---|---|---|
| authority | string (`cosmos.AddressString`) | yes | Signer; gov or pool admin |
| pool_id | string | yes | Pool to update; must exist |
| unstake_fee_rate | string (`cosmos.Dec` / `math.LegacyDec`, non-nullable) | yes | Replaces the pool's unstake fee rate |
| fee_account_address | string (`cosmos.AddressString`) | yes | Replaces the pool's fee account |
| autocompound_fee_rate | string (`cosmos.Dec` / `math.LegacyDec`, non-nullable) | yes | Replaces the pool's autocompound fee rate |
| whitelist_admin_address | string (`cosmos.AddressString`) | yes | Replaces the pool's admin address |

- **CLI:** No CLI command (construct via SDK / gov proposal).
- **TS SDK:**
  ```ts
  const msg = {
    typeUrl: '/ixo.liquidstake.v1beta1.MsgUpdatePool',
    value: ixo.liquidstake.v1beta1.MsgUpdatePool.fromPartial({
      authority: adminAddress,
      poolId: 'zero',
      unstakeFeeRate: '0.001000000000000000',
      feeAccountAddress: feeAddress,
      autocompoundFeeRate: '0.100000000000000000',
      whitelistAdminAddress: adminAddress,
    }),
  };
  ```
  Note: absent from the published SDK (see SDK warning).
- **Gotchas:** All four mutable fields are replaced in full and required. `ValidateBasic` validates bech32 `authority`, `pool_id`, both fee rates (`validateUnstakeFeeRate`/`validateAutocompoundFeeRate`, range `[0,1]`) and addresses. Server re-runs `pool.Validate()` and wraps failures as `ErrInvalidRequest`.

### MsgUpdateWhitelistedValidators
- **Purpose:** Replace the pool's validator whitelist. Target weights must sum to `10000`; each validator must exist and be `Bonded`.
- **Signer / auth:** `authority` signs. **Governance OR the pool's current admin.**
- **Fields:**

| Field | Type | Req | Description |
|---|---|---|---|
| authority | string (`cosmos.AddressString`) | yes | Signer; gov or pool admin |
| pool_id | string | yes | Pool whose whitelist is replaced |
| whitelisted_validators | repeated `WhitelistedValidator` (non-nullable) | yes | New validator set; each `{ validator_address (cosmos.AddressString), target_weight (cosmos.Int / math.Int) }`; weights sum to `10000` |

- **CLI:** No CLI command (construct via SDK / gov proposal).
- **TS SDK:**
  ```ts
  const msg = {
    typeUrl: '/ixo.liquidstake.v1beta1.MsgUpdateWhitelistedValidators',
    value: ixo.liquidstake.v1beta1.MsgUpdateWhitelistedValidators.fromPartial({
      authority: adminAddress,
      poolId: 'zero',
      whitelistedValidators: [
        { validatorAddress: 'ixovaloper1...', targetWeight: '10000' },
      ],
    }),
  };
  ```
  Note: the published SDK has `MsgUpdateWhitelistedValidators` but **without** the `poolId` field (single-pool layout) — see SDK warning.
- **Gotchas:** `ValidateBasic` runs `ValidateWhitelistedValidators` (each entry parses, positive non-nil `target_weight`, unique addresses) — failures wrap `ErrWhitelistedValidatorsList`. Server additionally requires every validator to exist on chain and be `stakingtypes.Bonded`, and the weight sum to equal `TotalValidatorWeight` (`10000`); all errors wrap `ErrWhitelistedValidatorsList`.

### MsgUpdateWeightedRewardsReceivers
- **Purpose:** Replace the pool's weighted-rewards receivers list (addresses paid a share of each autocompound before restaking).
- **Signer / auth:** `authority` signs. **Pool admin only** — `authority` must equal `Pool.whitelist_admin_address`; governance is NOT a valid signer here (matches pre-v7 admin-only constraint).
- **Fields:**

| Field | Type | Req | Description |
|---|---|---|---|
| authority | string (`cosmos.AddressString`) | yes | Signer; must equal the pool's `whitelist_admin_address` |
| pool_id | string | yes | Pool whose receivers list is replaced |
| weighted_rewards_receivers | repeated `WeightedAddress` (non-nullable) | yes | New list; each `{ address (cosmos.AddressString), weight (math.LegacyDec) }`; weight sum must not exceed `1` |

- **CLI:** No CLI command (construct via SDK).
- **TS SDK:**
  ```ts
  const msg = {
    typeUrl: '/ixo.liquidstake.v1beta1.MsgUpdateWeightedRewardsReceivers',
    value: ixo.liquidstake.v1beta1.MsgUpdateWeightedRewardsReceivers.fromPartial({
      authority: adminAddress,
      poolId: 'zero',
      weightedRewardsReceivers: [
        { address: 'ixo1...', weight: '0.500000000000000000' },
      ],
    }),
  };
  ```
  Note: the published SDK has `MsgUpdateWeightedRewardsReceivers` but **without** the `poolId` field — see SDK warning.
- **Gotchas:** `ValidateBasic` validates bech32 `authority`, `pool_id`, and `ValidateWeightedRewardsReceivers` (valid address, positive weight, sum ≤ `1`). Server enforces admin-only: a gov-address signer is rejected with `ErrorInvalidSigner` ("expected pool admin").

### MsgSetPoolPaused
- **Purpose:** Toggle a single pool's `paused` flag. Other pools unaffected.
- **Signer / auth:** `authority` signs. **Governance OR the pool's current admin.**
- **Fields:**

| Field | Type | Req | Description |
|---|---|---|---|
| authority | string (`cosmos.AddressString`) | yes | Signer; gov or pool admin |
| pool_id | string | yes | Pool to toggle |
| is_paused | bool | yes | Target value of `Pool.paused` |

- **CLI:** `ixod tx liquidstake pause-pool [pool-id] [paused] [flags]` — `[paused]` parsed as a bool (`strconv.ParseBool`); `authority` set from `--from`.
- **TS SDK:**
  ```ts
  const msg = {
    typeUrl: '/ixo.liquidstake.v1beta1.MsgSetPoolPaused',
    value: ixo.liquidstake.v1beta1.MsgSetPoolPaused.fromPartial({
      authority: adminAddress,
      poolId: 'zero',
      isPaused: true,
    }),
  };
  ```
  Note: absent from the published SDK (see SDK warning).
- **Gotchas:** `ValidateBasic` checks bech32 `authority` and `pool_id`. Server requires gov-or-admin (`authorisedByGovOrPoolAdmin`). The global `ModuleParams.module_paused` overrides this — a pool can be effectively paused regardless of `Pool.paused`.

### MsgSetModulePaused
- **Purpose:** Toggle the global `ModuleParams.module_paused` kill switch; when true every pool halts regardless of its per-pool flag.
- **Signer / auth:** `authority` signs. **Governance authority only.**
- **Fields:**

| Field | Type | Req | Description |
|---|---|---|---|
| authority | string (`cosmos.AddressString`) | yes | Signer; must be the gov module address |
| is_paused | bool | yes | Target value of `ModuleParams.module_paused` |

- **CLI:** No CLI command (governance-only; submit via gov proposal).
- **TS SDK:**
  ```ts
  const msg = {
    typeUrl: '/ixo.liquidstake.v1beta1.MsgSetModulePaused',
    value: ixo.liquidstake.v1beta1.MsgSetModulePaused.fromPartial({
      authority: govModuleAddress,
      isPaused: true,
    }),
  };
  ```
  Note: present in the published SDK with the same `{ authority, isPaused }` shape.
- **Gotchas:** `ValidateBasic` checks bech32 `authority` only. Server requires gov authority (`authorisedByGov`). No `pool_id` — this is module-wide.

### MsgBurn
- **Purpose:** Permanently destroy the signer's native `uixo` tokens. Module-level (not pool-scoped). Debug helper.
- **Signer / auth:** `burner` signs. Any account (burns its own coins).
- **Fields:**

| Field | Type | Req | Description |
|---|---|---|---|
| burner | string (`cosmos.AddressString`) | yes | Signer whose coins are burned |
| amount | `cosmos.base.v1beta1.Coin` (non-nullable) | yes | Coin to burn; denom must be `uixo` |

- **CLI:** `ixod tx liquidstake burn [amount] [flags]` — `burner` set from `--from`.
- **TS SDK:**
  ```ts
  const msg = {
    typeUrl: '/ixo.liquidstake.v1beta1.MsgBurn',
    value: ixo.liquidstake.v1beta1.MsgBurn.fromPartial({
      burner: burnerAddress,
      amount: { denom: 'uixo', amount: '1000000' },
    }),
  };
  ```
  Note: present in the published SDK with the same shape.
- **Gotchas:** `ValidateBasic` rejects invalid bech32 `burner`, zero/invalid `amount`, and any denom other than `uixo` ("burning amount must be in uixo"). Server re-checks the `uixo` denom (`ErrInvalidRequest`).

## Queries
No `ixod query liquidstake` CLI (the module registers no query command); use gRPC/REST. Service `ixo.liquidstake.v1beta1.Query`:

| Query | gRPC method | CLI | Args | Returns |
|---|---|---|---|---|
| Module params | `ModuleParams` | none | — (`QueryModuleParamsRequest{}`) | `ModuleParams` (`REST GET /ixo/liquidstake/v1beta1/module_params`) |
| Single pool | `Pool` | none | `pool_id` | `Pool` (`REST GET /ixo/liquidstake/v1beta1/pools/{pool_id}`) |
| All pools | `Pools` | none | `pagination` (PageRequest) | `repeated Pool` + `PageResponse` (`REST GET /ixo/liquidstake/v1beta1/pools`) |
| Liquid validators (one pool) | `LiquidValidators` | none | `pool_id` | `repeated LiquidValidatorState` (`REST GET /ixo/liquidstake/v1beta1/pools/{pool_id}/validators`) |
| Net amount state (one pool) | `States` | none | `pool_id` | `NetAmountState` (`REST GET /ixo/liquidstake/v1beta1/pools/{pool_id}/states`) |

## Events
Typed (proto) events from `event.proto`:
- **ModuleParamsUpdatedEvent** `{ module_params, authority }` — global `ModuleParams` changed (via `MsgUpdateModuleParams` or `MsgSetModulePaused`).
- **PoolCreatedEvent** `{ pool_id, pool, authority }` — new pool registered via `MsgCreatePool`.
- **PoolUpdatedEvent** `{ pool_id, pool, authority }` — a pool's config changed via any of `MsgUpdatePool`, `MsgUpdateWhitelistedValidators`, `MsgUpdateWeightedRewardsReceivers`, `MsgSetPoolPaused`.
- **LiquidStakeEvent** `{ delegator, liquid_amount (Coin), stk_ixo_minted_amount (Coin), pool_id }` — emitted on a successful `MsgLiquidStake`.
- **LiquidUnstakeEvent** `{ delegator, unstake_amount (Coin), unbonding_amount (Coin), unbonded_amount (Coin), completion_time (Timestamp), pool_id }` — emitted on a successful `MsgLiquidUnstake`.
- **AddLiquidValidatorEvent** `{ validator, pool_id }` — a newly whitelisted validator is activated for a pool.
- **RebalancedLiquidStakeEvent** `{ delegator, redelegation_count (uint32), redelegation_fail_count (uint32), pool_id }` — after a pool's rebalancing pass.
- **AutocompoundStakingRewardsEvent** `{ delegator, total_amount (Coin), fee_amount (Coin), redelegate_amount (Coin), weighted_rewards_amount (Coin), pool_id }` — after a pool's autocompound epoch hook runs successfully.

## Module gotchas
- **Admin/governance-only vs user-callable.** User-callable: `MsgLiquidUnstake` (any LST holder), `MsgBurn` (any account). Pool-admin-restricted `MsgLiquidStake` (only `Pool.whitelist_admin_address` may stake — single staker per pool by design, so the mint/unstake rate asymmetry causes no dilution). Gov-or-pool-admin: `MsgUpdatePool`, `MsgUpdateWhitelistedValidators`, `MsgSetPoolPaused`. **Pool-admin-only** (gov NOT allowed): `MsgUpdateWeightedRewardsReceivers`. **Governance-only:** `MsgCreatePool`, `MsgUpdateModuleParams`, `MsgSetModulePaused`. The gov-only and gov-or-admin update messages are typically submitted via `cosmos.gov.v1.MsgSubmitProposal` carrying the inner Msg.
- **Derivative denom.** Each pool's LST is a distinct bank coin (`Pool.liquid_bond_denom`, e.g. `uzero`); LSTs from different pools are NOT fungible. Stake mints 1:1; unstake burns at `NetAmount/TotalSupply`. Three layers (create-time guards, bank metadata claim, bank `MintCoinsRestriction`) stop other modules minting an LST denom.
- **Interaction with `x/staking`.** Delegations live in the per-pool proxy account, not the user's account. `LiquidUnstake` uses an LSM tokenize → bank-send → redeem → undelegate sequence so the unbonding queue ends up owned by the user (normal Cosmos unbonding flow). Whitelisted validators must be `Bonded`; the active set must hold ≥ 33.33% weight for stakes to succeed. Autocompound/rebalance epochs keep per-validator delegations aligned with `target_weight`.
- **Legacy `MsgUpdateParams` / `Params`.** `tx.proto` defines `MsgUpdateParams` and `liquidstake.proto` defines `Params` (`option deprecated = true`) **only** for backward-compat: they are registered in `codec.go` so the v7 node can *decode* pre-v7 historical txs/state, but there is **no RPC entry and no msg-server handler** — you cannot submit a new `MsgUpdateParams` on v7. Use `MsgUpdateModuleParams` + `MsgUpdatePool`/`MsgCreatePool` instead. (The stale published SDK still exposes `MsgUpdateParams` — do not use it for new txs.)
- **`zero` pool.** The pre-v7 single-pool deployment is migrated to a pool with `pool_id = "zero"` and `liquid_bond_denom = "uzero"`, keeping the original proxy account so existing delegations are preserved.
