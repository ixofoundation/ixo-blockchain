package keeper

import (
	"fmt"
	"strings"

	"cosmossdk.io/errors"
	storetypes "cosmossdk.io/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	"github.com/ixofoundation/ixo-blockchain/v7/x/liquidstake/types"
)

// ----------------------------------------------------------------------------
// Pool record CRUD
// ----------------------------------------------------------------------------

// GetPool fetches a Pool record by id.
func (k Keeper) GetPool(ctx sdk.Context, poolID string) (types.Pool, bool) {
	bz := ctx.KVStore(k.storeKey).Get(types.GetPoolKey(poolID))
	if bz == nil {
		return types.Pool{}, false
	}
	var p types.Pool
	k.cdc.MustUnmarshal(bz, &p)
	return p, true
}

// MustGetPool returns a Pool by id or wraps ErrPoolNotFound. Used inside
// keeper helpers where the caller has already passed an authorisation check
// or the missing-pool case is unrecoverable.
func (k Keeper) MustGetPool(ctx sdk.Context, poolID string) (types.Pool, error) {
	p, ok := k.GetPool(ctx, poolID)
	if !ok {
		return types.Pool{}, errors.Wrap(types.ErrPoolNotFound, poolID)
	}
	return p, nil
}

// SetPool persists a Pool record. Validation is the caller's responsibility
// — call Pool.Validate() before invoking this when accepting external input.
func (k Keeper) SetPool(ctx sdk.Context, p types.Pool) {
	ctx.KVStore(k.storeKey).Set(types.GetPoolKey(p.PoolId), k.cdc.MustMarshal(&p))
}

// HasPool reports whether a Pool with the given id exists.
func (k Keeper) HasPool(ctx sdk.Context, poolID string) bool {
	return ctx.KVStore(k.storeKey).Has(types.GetPoolKey(poolID))
}

// HasPoolWithDenom reports whether any pool currently uses the given LST
// denom. O(N) over registered pools — N is expected to stay small (single
// digits) so this stays well below 1ms even under heavy use.
func (k Keeper) HasPoolWithDenom(ctx sdk.Context, denom string) bool {
	found := false
	k.IteratePools(ctx, func(p types.Pool) bool {
		if p.LiquidBondDenom == denom {
			found = true
			return true
		}
		return false
	})
	return found
}

// IteratePools invokes fn for every persisted Pool in lexicographic pool_id
// order. Return true from fn to stop early.
func (k Keeper) IteratePools(ctx sdk.Context, fn func(p types.Pool) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	it := storetypes.KVStorePrefixIterator(store, types.PoolPrefix)
	defer it.Close()
	for ; it.Valid(); it.Next() {
		var p types.Pool
		k.cdc.MustUnmarshal(it.Value(), &p)
		if fn(p) {
			return
		}
	}
}

// GetAllPools returns every persisted Pool. Convenient for genesis export
// and admin queries; the Pools query uses a paginated path instead.
func (k Keeper) GetAllPools(ctx sdk.Context) []types.Pool {
	out := []types.Pool{}
	k.IteratePools(ctx, func(p types.Pool) bool {
		out = append(out, p)
		return false
	})
	return out
}

// ----------------------------------------------------------------------------
// Pool creation
// ----------------------------------------------------------------------------

// registerPool persists a brand-new Pool. Enforces pool_id uniqueness, denom
// uniqueness across pools, and address validity. The proxy account is
// derived deterministically from pool_id; the v7 upgrade is the only path
// that overrides that derivation (for the legacy "zero" pool, by writing
// directly via SetPool).
//
// Newly created pools start with empty whitelists / receivers, zero fee
// rates, and paused=false. Validators must subsequently be added with
// MsgUpdateWhitelistedValidators before any LiquidStake will succeed
// (the active-weight-quorum check requires a non-empty active set).
//
// Package-private so the embedded Keeper inside msgServer doesn't shadow
// the MsgServer.CreatePool interface method.
func (k Keeper) registerPool(
	ctx sdk.Context,
	poolID, liquidBondDenom string,
	adminAddress, feeAccountAddress string,
) (types.Pool, error) {
	if err := types.ValidatePoolID(poolID); err != nil {
		return types.Pool{}, errors.Wrap(types.ErrInvalidPoolID, err.Error())
	}
	if k.HasPool(ctx, poolID) {
		return types.Pool{}, errors.Wrap(types.ErrDuplicatePoolID, poolID)
	}
	if k.HasPoolWithDenom(ctx, liquidBondDenom) {
		return types.Pool{}, errors.Wrap(types.ErrDuplicateLiquidBondDenom, liquidBondDenom)
	}

	// Reject denoms in the IBC voucher namespace. Only the IBC transfer
	// module legitimately mints into ibc/<HASH>; allowing a liquidstake
	// pool there would silently shadow vouchers, and any IBC transfer
	// would later inflate stkixoTotalSupply and break the unstake rate.
	if types.IsIBCDenom(liquidBondDenom) {
		return types.Pool{}, errors.Wrapf(types.ErrDenomAlreadyInUse,
			"denom %s is in the IBC voucher namespace and cannot be a liquid bond denom",
			liquidBondDenom)
	}

	// Reject the chain's native bond denom — it's reserved for raw staking
	// and a pool that minted "uixo" would shadow real IXO balances.
	bondDenom, err := k.stakingKeeper.BondDenom(ctx)
	if err != nil {
		panic(err) // should never happen
	}
	if liquidBondDenom == bondDenom {
		return types.Pool{}, errors.Wrapf(types.ErrDenomAlreadyInUse,
			"%s is the chain bond denom and cannot be used as a liquid bond denom",
			liquidBondDenom)
	}

	// Reject any denom that already has non-zero bank supply. The
	// liquidstake module account has the Minter permission, so without
	// this guard we could mint into a denom that some other module / IBC
	// transfer / token-factory already issued. The pool's NetAmountState
	// computes stkixoTotalSupply via bankKeeper.GetSupply(denom), which
	// would silently include those externally-minted tokens — letting
	// outsiders unstake against delegations they never funded.
	if !k.bankKeeper.GetSupply(ctx, liquidBondDenom).Amount.IsZero() {
		return types.Pool{}, errors.Wrapf(types.ErrDenomAlreadyInUse,
			"denom %s already has non-zero supply in the bank module",
			liquidBondDenom)
	}

	// Reject any denom with bank metadata registered. Metadata is the
	// chain's record of "this denom belongs to module X" — even if it
	// currently has zero supply (e.g. uixo before genesis distribution,
	// IBC vouchers between full round-trips). Catches the gap between
	// the supply check and a denom being briefly empty.
	if k.bankKeeper.HasDenomMetaData(ctx, liquidBondDenom) {
		return types.Pool{}, errors.Wrapf(types.ErrDenomAlreadyInUse,
			"denom %s already has bank denom metadata registered",
			liquidBondDenom)
	}

	pool := types.Pool{
		PoolId:                   poolID,
		LiquidBondDenom:          liquidBondDenom,
		ProxyAccountAddress:      types.PoolProxyAcc(poolID).String(),
		WhitelistedValidators:    []types.WhitelistedValidator{},
		UnstakeFeeRate:           types.DefaultUnstakeFeeRate,
		FeeAccountAddress:        feeAccountAddress,
		AutocompoundFeeRate:      types.DefaultAutocompoundFeeRate,
		WhitelistAdminAddress:    adminAddress,
		Paused:                   false,
		WeightedRewardsReceivers: []types.WeightedAddress{},
	}
	if err := pool.Validate(); err != nil {
		// MsgCreatePool.ValidateBasic has already run, so reaching this
		// branch implies a bug or genesis-level inconsistency rather than
		// a bad pool_id. Surface the underlying validation error verbatim.
		return types.Pool{}, fmt.Errorf("registerPool %q: %w", poolID, err)
	}

	k.SetPool(ctx, pool)
	k.RegisterLSTDenomMetadata(ctx, poolID, liquidBondDenom)
	return pool, nil
}

// RegisterLSTDenomMetadata claims an LST denom in bank's metadata registry.
// This is the standard "this denom belongs to module X" signal that other
// modules (e.g. bonds at create time) consult via HasDenomMetaData. Closes
// the gap where a freshly created pool has zero supply yet, leaving the
// supply-only front-door check open until the first LiquidStake.
//
// Called by registerPool for new pools and by the v7 migration for the
// legacy "zero" pool. The MintCoinsRestriction in app/keepers/keepers.go
// remains the enduring chain-level backstop for any future module that
// might bypass the metadata check.
//
// Denom-units heuristic: if liquidBondDenom follows the standard Cosmos
// "u"-prefix convention (e.g. "uzero", "ucarbon"), we declare TWO units
// — the micro base at exponent 0 and the human-readable display at
// exponent 6 — matching how uixo/ixo, uatom/atom, etc. are conventionally
// modelled. For non-u-prefixed denoms we fall back to a single base unit
// at exponent 0, which is still a valid Metadata record.
func (k Keeper) RegisterLSTDenomMetadata(ctx sdk.Context, poolID, liquidBondDenom string) {
	denomUnits, displayDenom, displaySymbol := lstDenomUnits(liquidBondDenom)
	k.bankKeeper.SetDenomMetaData(ctx, banktypes.Metadata{
		Description: fmt.Sprintf("Liquid staking token for pool %q", poolID),
		DenomUnits:  denomUnits,
		Base:        liquidBondDenom,
		Display:     displayDenom,
		Name:        fmt.Sprintf("liquidstake LST (%s)", poolID),
		Symbol:      strings.ToUpper(displaySymbol),
	})
}

// lstDenomUnits returns the denom-unit slice, display denom string, and
// symbol string to use when registering metadata for an LST denom.
//
// For a "u"-prefixed denom (e.g. "uzero"), declares both the base unit
// (uzero, exp 0) and the display unit (zero, exp 6). For other shapes
// (e.g. "stkfoo"), declares a single unit at the base; display == base
// and symbol == base.
func lstDenomUnits(liquidBondDenom string) (units []*banktypes.DenomUnit, display, symbol string) {
	if len(liquidBondDenom) > 1 && strings.HasPrefix(liquidBondDenom, "u") {
		display = liquidBondDenom[1:] // strip the "u" prefix
		return []*banktypes.DenomUnit{
			{Denom: liquidBondDenom, Exponent: 0},
			{Denom: display, Exponent: 6},
		}, display, display
	}
	return []*banktypes.DenomUnit{
		{Denom: liquidBondDenom, Exponent: 0},
	}, liquidBondDenom, liquidBondDenom
}

// ----------------------------------------------------------------------------
// Per-pool LiquidValidator CRUD
// ----------------------------------------------------------------------------

// GetLiquidValidator fetches a single per-pool LiquidValidator record.
func (k Keeper) GetLiquidValidator(ctx sdk.Context, poolID string, addr sdk.ValAddress) (val types.LiquidValidator, found bool) {
	bz := ctx.KVStore(k.storeKey).Get(types.GetLiquidValidatorKey(poolID, addr))
	if bz == nil {
		return val, false
	}
	val = types.MustUnmarshalLiquidValidator(k.cdc, bz)
	return val, true
}

// SetLiquidValidator persists a per-pool LiquidValidator record.
func (k Keeper) SetLiquidValidator(ctx sdk.Context, poolID string, val types.LiquidValidator) {
	bz := types.MustMarshalLiquidValidator(k.cdc, &val)
	ctx.KVStore(k.storeKey).Set(types.GetLiquidValidatorKey(poolID, val.GetOperator()), bz)
}

// RemoveLiquidValidator deletes a per-pool LiquidValidator record.
func (k Keeper) RemoveLiquidValidator(ctx sdk.Context, poolID string, val types.LiquidValidator) {
	ctx.KVStore(k.storeKey).Delete(types.GetLiquidValidatorKey(poolID, val.GetOperator()))
}

// IterateLiquidValidatorsForPool invokes fn for every LiquidValidator
// belonging to the given pool. Return true from fn to stop early.
func (k Keeper) IterateLiquidValidatorsForPool(ctx sdk.Context, poolID string, fn func(val types.LiquidValidator) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	it := storetypes.KVStorePrefixIterator(store, types.GetLiquidValidatorsByPoolPrefix(poolID))
	defer it.Close()
	for ; it.Valid(); it.Next() {
		val := types.MustUnmarshalLiquidValidator(k.cdc, it.Value())
		if fn(val) {
			return
		}
	}
}

// GetAllLiquidValidatorsForPool returns every persisted LiquidValidator for
// the given pool. Pagination is unnecessary here since the active-set is
// bounded by the validator whitelist (capped to a small number of operators).
func (k Keeper) GetAllLiquidValidatorsForPool(ctx sdk.Context, poolID string) types.LiquidValidators {
	out := types.LiquidValidators{}
	k.IterateLiquidValidatorsForPool(ctx, poolID, func(val types.LiquidValidator) bool {
		out = append(out, val)
		return false
	})
	return out
}
