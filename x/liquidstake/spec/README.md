# Liquidstake module specification

This document specifies the liquidstake module, a custom Ixo Cosmos SDK module.

The liquidstake module mints **liquid staking tokens (LSTs)** that represent a delegator's claim on a pool of native-token delegations to a curated, weighted set of validators. Each **pool** is an independent liquid-staking instance with its own LST denom, validator whitelist, admin, fees, paused flag, and proxy account; pools are created through governance and operated by their pool admin. Validators delegated through a pool earn rewards, which are auto-compounded back into the pool on a configurable epoch.

The `v7` chain upgrade reshaped this module from a single global pool (the original ZERO/uzero pool, deployed in `v4`) into a multi-pool layout while preserving the legacy `zero` pool's proxy account address and existing delegations through migration.

## Contents

1. **[Concepts](01_concepts.md)**

   - [Liquid staking](01_concepts.md#liquid-staking)
   - [Pools](01_concepts.md#pools)
   - [Mint and unstake rates](01_concepts.md#mint-and-unstake-rates)
   - [Per-pool proxy account](01_concepts.md#per-pool-proxy-account)
   - [Whitelisted validators and weights](01_concepts.md#whitelisted-validators-and-weights)
   - [Autocompounding and rebalancing epochs](01_concepts.md#autocompounding-and-rebalancing-epochs)
   - [Pause flags (per-pool and global)](01_concepts.md#pause-flags-per-pool-and-global)
   - [Denom collision defences](01_concepts.md#denom-collision-defences)
   - [v7 upgrade migration](01_concepts.md#v7-upgrade-migration)

2. **[State](02_state.md)**

   - [Storage layout](02_state.md#storage-layout)
   - [Types](02_state.md#types)
     - [ModuleParams](02_state.md#moduleparams)
     - [Pool](02_state.md#pool)
     - [WhitelistedValidator](02_state.md#whitelistedvalidator)
     - [WeightedAddress](02_state.md#weightedaddress)
     - [LiquidValidator](02_state.md#liquidvalidator)
     - [LiquidValidatorState](02_state.md#liquidvalidatorstate)
     - [NetAmountState](02_state.md#netamountstate)
     - [ValidatorStatus enum](02_state.md#validatorstatus)
     - [Legacy Params (deprecated)](02_state.md#legacy-params-deprecated)

3. **[Messages](03_messages.md)**

   - [End-user](03_messages.md#end-user-operations)
     - [MsgLiquidStake](03_messages.md#msgliquidstake)
     - [MsgLiquidUnstake](03_messages.md#msgliquidunstake)
     - [MsgBurn](03_messages.md#msgburn)
   - [Pool admin](03_messages.md#pool-admin-operations)
     - [MsgUpdatePool](03_messages.md#msgupdatepool)
     - [MsgUpdateWhitelistedValidators](03_messages.md#msgupdatewhitelistedvalidators)
     - [MsgUpdateWeightedRewardsReceivers](03_messages.md#msgupdateweightedrewardsreceivers)
     - [MsgSetPoolPaused](03_messages.md#msgsetpoolpaused)
   - [Governance](03_messages.md#governance-operations)
     - [MsgCreatePool](03_messages.md#msgcreatepool)
     - [MsgUpdateModuleParams](03_messages.md#msgupdatemoduleparams)
     - [MsgSetModulePaused](03_messages.md#msgsetmodulepaused)

4. **[Events](04_events.md)**

   - [PoolCreatedEvent](04_events.md#poolcreatedevent)
   - [PoolUpdatedEvent](04_events.md#poolupdatedevent)
   - [ModuleParamsUpdatedEvent](04_events.md#moduleparamsupdatedevent)
   - [LiquidStakeEvent](04_events.md#liquidstakeevent)
   - [LiquidUnstakeEvent](04_events.md#liquidunstakeevent)
   - [AddLiquidValidatorEvent](04_events.md#addliquidvalidatorevent)
   - [RebalancedLiquidStakeEvent](04_events.md#rebalancedliquidstakeevent)
   - [AutocompoundStakingRewardsEvent](04_events.md#autocompoundstakingrewardsevent)

5. **[Parameters](05_params.md)**

   - [ModuleParams](05_params.md#moduleparams)
   - [Pool fields](05_params.md#pool-fields)
   - [Module-wide constants](05_params.md#module-wide-constants)
