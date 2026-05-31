package keeper

import (
	"fmt"

	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ixofoundation/ixo-blockchain/v7/x/liquidstake/types"
)

// Keeper is the persistent state interface for the liquidstake module. The
// store at storeKey holds: ModuleParams (single record at 0x01), Pool records
// (one per pool, prefixed 0x10), and per-pool LiquidValidator records
// (prefixed 0x11). The router is used to invoke staking and distribution
// MsgServer handlers in-process; authority is the gov module address.
type Keeper struct {
	cdc      codec.BinaryCodec
	storeKey storetypes.StoreKey

	accountKeeper  types.AccountKeeper
	bankKeeper     types.BankKeeper
	stakingKeeper  types.StakingKeeper
	distrKeeper    types.DistrKeeper
	slashingKeeper types.SlashingKeeper

	router    *baseapp.MsgServiceRouter
	authority string
}

// NewKeeper constructs a Keeper. Panics if the module account is not present
// in account-keeper state — that would indicate maccPerms is misconfigured
// in app wiring.
func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey storetypes.StoreKey,
	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,
	stakingKeeper types.StakingKeeper,
	distrKeeper types.DistrKeeper,
	slashingKeeper types.SlashingKeeper,
	router *baseapp.MsgServiceRouter,
	authority string,
) Keeper {
	// ensure liquidstake module account is set
	if addr := accountKeeper.GetModuleAddress(types.ModuleName); addr == nil {
		panic(fmt.Sprintf("%s module account has not been set", types.ModuleName))
	}

	return Keeper{
		cdc:            cdc,
		storeKey:       storeKey,
		accountKeeper:  accountKeeper,
		bankKeeper:     bankKeeper,
		stakingKeeper:  stakingKeeper,
		distrKeeper:    distrKeeper,
		slashingKeeper: slashingKeeper,
		router:         router,
		authority:      authority,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// Authority returns the bech32-encoded address authorised to update global
// module parameters and create pools (typically x/gov).
func (k Keeper) Authority() string { return k.authority }

// SetModuleParams persists global module parameters after validating them.
func (k Keeper) SetModuleParams(ctx sdk.Context, params types.ModuleParams) error {
	if err := params.Validate(); err != nil {
		return err
	}
	ctx.KVStore(k.storeKey).Set(types.ModuleParamsKey, k.cdc.MustMarshal(&params))
	return nil
}

// GetModuleParams returns the current global module parameters. If the store
// is empty (e.g. no genesis init has run yet), returns the zero-value
// ModuleParams; callers should treat that as "module not yet initialised".
func (k Keeper) GetModuleParams(ctx sdk.Context) (params types.ModuleParams) {
	bz := ctx.KVStore(k.storeKey).Get(types.ModuleParamsKey)
	if bz == nil {
		return params
	}
	k.cdc.MustUnmarshal(bz, &params)
	return params
}

// IsModulePaused is a convenience wrapper checking the global kill switch.
// Callers must combine this with the per-pool Pool.Paused flag — either
// being true halts the affected pool's operations.
func (k Keeper) IsModulePaused(ctx sdk.Context) bool {
	return k.GetModuleParams(ctx).ModulePaused
}

func (k Keeper) GetCodec() codec.BinaryCodec       { return k.cdc }
func (k Keeper) Router() *baseapp.MsgServiceRouter { return k.router }

// GetStoreKey returns the module's storeKey. Exported strictly for use by
// upgrade migrations that need to iterate or rewrite raw KV entries
// (e.g. re-keying legacy records under a new prefix). Application code
// must not use this; it bypasses every keeper-level invariant.
func (k Keeper) GetStoreKey() storetypes.StoreKey { return k.storeKey }

// ----------------------------------------------------------------------------
// Legacy single-pool Params helpers
// ----------------------------------------------------------------------------
//
// These helpers operate on the same KV slot (types.ModuleParamsKey, 0x01)
// as ModuleParams but using the deprecated single-pool types.Params layout.
// They exist solely to keep the historical v4 upgrade compilable and to
// give the v7 migration a typed accessor for pre-upgrade state. New code
// must use {Get,Set}ModuleParams instead.

// SetLegacyParams writes a legacy types.Params record to the KV store.
// Used by the historical v4 upgrade handler.
func (k Keeper) SetLegacyParams(ctx sdk.Context, params types.Params) error {
	ctx.KVStore(k.storeKey).Set(types.ModuleParamsKey, k.cdc.MustMarshal(&params))
	return nil
}

// GetLegacyParams reads a legacy types.Params record from the KV store.
// Used by the v7 migration to read pre-upgrade state.
func (k Keeper) GetLegacyParams(ctx sdk.Context) (params types.Params) {
	bz := ctx.KVStore(k.storeKey).Get(types.ModuleParamsKey)
	if bz == nil {
		return params
	}
	k.cdc.MustUnmarshal(bz, &params)
	return params
}
