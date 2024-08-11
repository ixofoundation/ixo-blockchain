package keeper

import (
	"fmt"

	"cosmossdk.io/log"
	storetypes "cosmossdk.io/store/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/ixofoundation/ixo-blockchain/v3/x/entity/types"
)

// UnmarshalFn is a generic function to unmarshal bytes
type UnmarshalFn func(value []byte) (interface{}, bool)

// MarshalFn is a generic function to marshal bytes
type MarshalFn func(value interface{}) []byte

type Keeper struct {
	cdc            codec.BinaryCodec
	storeKey       storetypes.StoreKey
	memStoreKey    storetypes.StoreKey
	IidKeeper      types.IidKeeper
	WasmKeeper     types.WasmKeeper
	WasmViewKeeper types.WasmViewKeeper
	ParamSpace     paramstypes.Subspace
	AccountKeeper  types.AccountKeeper
	AuthzKeeper    types.AuthzKeeper
}

func NewKeeper(
	cdc codec.BinaryCodec,
	key storetypes.StoreKey,
	memStoreKey storetypes.StoreKey,
	iidKeeper types.IidKeeper,
	wasmKeeper types.WasmKeeper,
	wasViewKeeper types.WasmViewKeeper,
	paramSpace paramstypes.Subspace,
	accountKeeper types.AccountKeeper,
	authzKeeper types.AuthzKeeper,
) Keeper {
	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		cdc:            cdc,
		storeKey:       key,
		memStoreKey:    memStoreKey,
		IidKeeper:      iidKeeper,
		WasmKeeper:     wasmKeeper,
		ParamSpace:     paramSpace,
		AccountKeeper:  accountKeeper,
		AuthzKeeper:    authzKeeper,
		WasmViewKeeper: wasViewKeeper,
	}
}

// GetParams returns the total set of project parameters.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.ParamSpace.GetParamSet(ctx, &params)
	return params
}

// SetParams sets the total set of project parameters.
func (k Keeper) SetParams(ctx sdk.Context, params *types.Params) {
	k.ParamSpace.SetParamSet(ctx, params)
}

func (k Keeper) EntityExists(ctx sdk.Context, entityDid string) bool {
	// store := ctx.KVStore(k.storeKey)
	_, exists := k.IidKeeper.GetDidDocument(ctx, []byte(entityDid))
	return exists
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// Set sets a value in the db with a prefixed key
func (k Keeper) Set(ctx sdk.Context,
	key []byte,
	prefix []byte,
	i interface{},
	marshal MarshalFn,
) {
	store := ctx.KVStore(k.storeKey)
	store.Set(append(prefix, key...), marshal(i))
}

// Delete - deletes a value form the store
func (k Keeper) Delete(
	ctx sdk.Context,
	key []byte,
	prefix []byte,
) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(append(prefix, key...))
}

// Get gets an item from the store by bytes
func (k Keeper) Get(
	ctx sdk.Context,
	key []byte,
	prefix []byte,
	unmarshal UnmarshalFn,
) (i interface{}, found bool) {
	store := ctx.KVStore(k.storeKey)
	value := store.Get(append(prefix, key...))

	return unmarshal(value)
}

// GetAll values from with a prefix from the store
func (k Keeper) GetAll(
	ctx sdk.Context,
	prefix []byte,
) storetypes.Iterator {
	store := ctx.KVStore(k.storeKey)
	return storetypes.KVStorePrefixIterator(store, prefix)
}
