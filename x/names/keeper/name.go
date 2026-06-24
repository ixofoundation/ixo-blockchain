package keeper

import (
	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"

	"github.com/ixofoundation/ixo-blockchain/v8/x/names/types"
)

// SetNameRecord writes a NameRecord and (re)populates the owner reverse
// index. Callers must ensure the namespace exists.
func (k Keeper) SetNameRecord(ctx sdk.Context, record types.NameRecord) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.NameRecordKey(record.Namespace, record.NormalizedName), k.cdc.MustMarshal(&record))
	store.Set(types.OwnerIndexKey(record.OwnerDid, record.Namespace, record.NormalizedName), []byte{})
}

// GetNameRecord returns a NameRecord by (namespace, normalized_name).
func (k Keeper) GetNameRecord(ctx sdk.Context, namespace, normalizedName string) (types.NameRecord, bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.NameRecordKey(namespace, normalizedName))
	if bz == nil {
		return types.NameRecord{}, false
	}
	var r types.NameRecord
	k.cdc.MustUnmarshal(bz, &r)
	return r, true
}

// HasNameRecord reports whether a NameRecord at (namespace, normalized_name)
// exists, regardless of status.
func (k Keeper) HasNameRecord(ctx sdk.Context, namespace, normalizedName string) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.NameRecordKey(namespace, normalizedName))
}

// removeOwnerIndex deletes a reverse-index entry. Used when changing owners.
func (k Keeper) removeOwnerIndex(ctx sdk.Context, ownerDid, namespace, normalizedName string) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.OwnerIndexKey(ownerDid, namespace, normalizedName))
}

// PaginateNamesByNamespace returns paginated NameRecords under namespace.
func (k Keeper) PaginateNamesByNamespace(ctx sdk.Context, namespace string, pageReq *query.PageRequest) ([]types.NameRecord, *query.PageResponse, error) {
	store := ctx.KVStore(k.storeKey)
	nsStore := prefix.NewStore(store, types.NameRecordNamespacePrefix(namespace))

	var out []types.NameRecord
	pageRes, err := query.Paginate(nsStore, pageReq, func(_ []byte, value []byte) error {
		var r types.NameRecord
		if err := k.cdc.Unmarshal(value, &r); err != nil {
			return err
		}
		out = append(out, r)
		return nil
	})
	if err != nil {
		return nil, nil, err
	}
	return out, pageRes, nil
}

// PaginateNamesByOwner returns every NameRecord owned by ownerDid using the
// reverse index. The returned records may span multiple namespaces.
func (k Keeper) PaginateNamesByOwner(ctx sdk.Context, ownerDid string, pageReq *query.PageRequest) ([]types.NameRecord, *query.PageResponse, error) {
	store := ctx.KVStore(k.storeKey)
	ownerStore := prefix.NewStore(store, types.OwnerIndexPrefix(ownerDid))

	var out []types.NameRecord
	pageRes, err := query.Paginate(ownerStore, pageReq, func(key []byte, _ []byte) error {
		ns, name, ok := types.ParseOwnerIndexSuffix(key)
		if !ok {
			return nil
		}
		r, found := k.GetNameRecord(ctx, ns, name)
		if !found {
			return nil
		}
		out = append(out, r)
		return nil
	})
	if err != nil {
		return nil, nil, err
	}
	return out, pageRes, nil
}

// IterateAllNames walks every NameRecord (across all namespaces) in
// deterministic order. Used for genesis export.
func (k Keeper) IterateAllNames(ctx sdk.Context, cb func(types.NameRecord) bool) {
	store := ctx.KVStore(k.storeKey)
	iterator := storetypes.KVStorePrefixIterator(store, types.NameRecordKeyPrefix)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var r types.NameRecord
		k.cdc.MustUnmarshal(iterator.Value(), &r)
		if cb(r) {
			break
		}
	}
}

// GetAllNames returns every NameRecord (used for genesis export).
func (k Keeper) GetAllNames(ctx sdk.Context) []types.NameRecord {
	out := []types.NameRecord{}
	k.IterateAllNames(ctx, func(r types.NameRecord) bool {
		out = append(out, r)
		return false
	})
	return out
}
