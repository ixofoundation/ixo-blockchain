package v7

import (
	"fmt"

	storetypes "cosmossdk.io/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	liquidstakekeeper "github.com/ixofoundation/ixo-blockchain/v6/x/liquidstake/keeper"
	liquidstaketypes "github.com/ixofoundation/ixo-blockchain/v6/x/liquidstake/types"
)

// migrateLiquidStakeToMultiPool reshapes the pre-v7 single-pool liquidstake
// state into the v7 multi-pool layout.
//
// Steps, in order:
//
//  1. Read the legacy types.Params record from KV[0x01].
//  2. Construct ModuleParams from the legacy single-pool fields.
//  3. Persist the legacy "zero" Pool with proxy_account_address set to the
//     PRE-V7 LiquidStakeProxyAcc bech32 — preserves all existing
//     delegations and unbondings without state migration of staking module
//     records.
//  4. Re-key every legacy LiquidValidator record (prefix 0x02 + valAddr)
//     under the new per-pool prefix (0x11 + lp("zero") + valAddr).
//  5. Delete the legacy 0x02 records.
//  6. Overwrite KV[0x01] with the marshaled ModuleParams.
//
// Idempotency note: this migration must run exactly once per chain (chained
// to the v7 SoftwareUpgradeProposal). Re-running it after legacy state has
// already been consumed would write empty/zero ModuleParams over a real
// configuration; the upgrade keeper guarantees single execution.
func migrateLiquidStakeToMultiPool(ctx sdk.Context, k liquidstakekeeper.Keeper) error {
	// Step 1: read legacy params.
	legacy := k.GetLegacyParams(ctx)

	// Step 2: build new global ModuleParams. Min stake amount carries over
	// (it was already a single value); module_paused inherits the legacy
	// pause flag so an emergency pause survives the upgrade.
	moduleParams := liquidstaketypes.ModuleParams{
		MinLiquidStakeAmount: legacy.MinLiquidStakeAmount,
		ModulePaused:         legacy.ModulePaused,
	}
	if moduleParams.MinLiquidStakeAmount.IsNil() {
		moduleParams.MinLiquidStakeAmount = liquidstaketypes.DefaultMinLiquidStakeAmount
	}
	if err := moduleParams.Validate(); err != nil {
		return fmt.Errorf("v7 migration: invalid migrated module_params: %w", err)
	}

	// Skip the pool-construction path entirely if the legacy state never
	// configured a liquid bond denom — that means liquidstake was added
	// to the binary but never initialised on this chain. Persist
	// ModuleParams only and exit; future pools come in via MsgCreatePool.
	if legacy.LiquidBondDenom == "" {
		if err := k.SetModuleParams(ctx, moduleParams); err != nil {
			return fmt.Errorf("v7 migration: failed to persist module_params: %w", err)
		}
		ctx.Logger().Info(
			"v7 multi-pool liquidstake migration: no legacy denom configured, no pool migrated",
		)
		return nil
	}

	// Step 3: build the legacy "zero" Pool. ProxyAccountAddress is the
	// pre-v7 LiquidStakeProxyAcc string so existing delegations remain
	// addressable via this pool's proxy.
	zeroPool := liquidstaketypes.Pool{
		PoolId:                   LegacyPoolID,
		LiquidBondDenom:          legacy.LiquidBondDenom,
		ProxyAccountAddress:      liquidstaketypes.LegacyLiquidStakeProxyAcc().String(),
		WhitelistedValidators:    legacy.WhitelistedValidators,
		UnstakeFeeRate:           legacy.UnstakeFeeRate,
		FeeAccountAddress:        legacy.FeeAccountAddress,
		AutocompoundFeeRate:      legacy.AutocompoundFeeRate,
		WhitelistAdminAddress:    legacy.WhitelistAdminAddress,
		Paused:                   false, // global pause carried into ModuleParams; per-pool starts clean
		WeightedRewardsReceivers: legacy.WeightedRewardsReceivers,
	}
	// Legacy v4 setup may have left WhitelistedValidators / receivers nil
	// rather than empty slices; pool.Validate tolerates that, but the
	// keeper expects non-nil slices for downstream iteration.
	if zeroPool.WhitelistedValidators == nil {
		zeroPool.WhitelistedValidators = []liquidstaketypes.WhitelistedValidator{}
	}
	if zeroPool.WeightedRewardsReceivers == nil {
		zeroPool.WeightedRewardsReceivers = []liquidstaketypes.WeightedAddress{}
	}
	if err := zeroPool.Validate(); err != nil {
		return fmt.Errorf("v7 migration: invalid migrated zero pool: %w", err)
	}

	// Step 4 + 5: collect legacy LiquidValidator entries (prefix 0x02), then
	// re-key each one under the per-pool prefix and delete the legacy key.
	store := ctx.KVStore(k.GetStoreKey())

	type legacyLV struct {
		key   []byte // legacy KV key (full)
		value []byte // marshaled LiquidValidator
	}
	collected := []legacyLV{}
	{
		it := storetypes.KVStorePrefixIterator(store, liquidstaketypes.LegacyLiquidValidators_)
		defer it.Close()
		for ; it.Valid(); it.Next() {
			// Copy the key/value bytes; the iterator backing array is reused
			// across iterations. Local names avoid shadowing the outer
			// keeper parameter `k`.
			keyCopy := append([]byte{}, it.Key()...)
			valCopy := append([]byte{}, it.Value()...)
			collected = append(collected, legacyLV{key: keyCopy, value: valCopy})
		}
	}

	for _, entry := range collected {
		store.Delete(entry.key)
		lv := liquidstaketypes.MustUnmarshalLiquidValidator(k.GetCodec(), entry.value)
		k.SetLiquidValidator(ctx, LegacyPoolID, lv)
	}

	// Step 3 (persist) + Step 6: write Pool, then ModuleParams. ModuleParams
	// is written last so a partial failure leaves the legacy Params intact
	// for a retry — Cosmos SDK upgrade transactions are atomic, so this is
	// belt-and-braces.
	k.SetPool(ctx, zeroPool)

	// Claim the legacy "uzero" denom in bank's metadata registry so future
	// bond / token-factory / etc. attempts to mint into the same denom are
	// rejected at their own front-door checks. Mirrors the metadata claim
	// performed by registerPool for any pool created via MsgCreatePool.
	k.RegisterLSTDenomMetadata(ctx, LegacyPoolID, zeroPool.LiquidBondDenom)

	if err := k.SetModuleParams(ctx, moduleParams); err != nil {
		return fmt.Errorf("v7 migration: failed to persist module_params: %w", err)
	}

	ctx.Logger().Info(
		"v7 multi-pool liquidstake migration complete",
		"pool_id", LegacyPoolID,
		"liquid_bond_denom", zeroPool.LiquidBondDenom,
		"proxy_account_address", zeroPool.ProxyAccountAddress,
		"liquid_validators_migrated", len(collected),
	)
	return nil
}
