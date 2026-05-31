package keeper

import (
	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/ixofoundation/ixo-blockchain/v7/x/names/types"
)

// SetNamespace writes a Namespace to state.
func (k Keeper) SetNamespace(ctx sdk.Context, ns types.Namespace) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.NamespaceKey(ns.Name), k.cdc.MustMarshal(&ns))
}

// GetNamespace returns a Namespace by name. The bool indicates existence.
func (k Keeper) GetNamespace(ctx sdk.Context, name string) (types.Namespace, bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.NamespaceKey(name))
	if bz == nil {
		return types.Namespace{}, false
	}
	var ns types.Namespace
	k.cdc.MustUnmarshal(bz, &ns)
	return ns, true
}

// HasNamespace reports whether a Namespace with name exists.
func (k Keeper) HasNamespace(ctx sdk.Context, name string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.NamespaceKey(name))
}

// IterateNamespaces walks every Namespace in deterministic key order. The cb
// returns true to stop iteration.
func (k Keeper) IterateNamespaces(ctx sdk.Context, cb func(types.Namespace) bool) {
	store := ctx.KVStore(k.storeKey)
	iterator := storetypes.KVStorePrefixIterator(store, types.NamespaceKeyPrefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var ns types.Namespace
		k.cdc.MustUnmarshal(iterator.Value(), &ns)
		if cb(ns) {
			break
		}
	}
}

// GetAllNamespaces returns every Namespace.
func (k Keeper) GetAllNamespaces(ctx sdk.Context) []types.Namespace {
	out := []types.Namespace{}
	k.IterateNamespaces(ctx, func(ns types.Namespace) bool {
		out = append(out, ns)
		return false
	})
	return out
}

// PaginateNamespaces returns Namespaces using the supplied pagination request.
func (k Keeper) PaginateNamespaces(ctx sdk.Context, pageReq *query.PageRequest) ([]types.Namespace, *query.PageResponse, error) {
	store := ctx.KVStore(k.storeKey)
	nsStore := prefix.NewStore(store, types.NamespaceKeyPrefix)

	var out []types.Namespace
	pageRes, err := query.Paginate(nsStore, pageReq, func(_ []byte, value []byte) error {
		var ns types.Namespace
		if err := k.cdc.Unmarshal(value, &ns); err != nil {
			return err
		}
		out = append(out, ns)
		return nil
	})
	if err != nil {
		return nil, nil, err
	}
	return out, pageRes, nil
}
