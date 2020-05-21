package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-blockchain/x/fees/internal/types"
)

// -------------------------------------------------------- Fees

func (k Keeper) GetFeeIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, types.FeeKeyPrefix)
}

func (k Keeper) MustGetFeeByKey(ctx sdk.Context, key []byte) types.Fee {
	store := ctx.KVStore(k.storeKey)
	if !store.Has(key) {
		panic("fee not found")
	}

	bz := store.Get(key)
	var fee types.Fee
	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &fee)

	return fee
}

func (k Keeper) FeeExists(ctx sdk.Context, feeId uint64) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetFeeKey(feeId))
}

func (k Keeper) GetFee(ctx sdk.Context, feeId uint64) (types.Fee, sdk.Error) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetFeeKey(feeId)

	bz := store.Get(key)
	if bz == nil {
		return types.Fee{}, sdk.ErrInternal("invalid fee")
	}

	var fee types.Fee
	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &fee)

	return fee, nil
}

func (k Keeper) SubmitFee(ctx sdk.Context, content types.FeeContent) (types.Fee, sdk.Error) {
	feeId, err := k.GetFeeID(ctx)
	if err != nil {
		return types.Fee{}, err
	}

	fee := types.NewFee(feeId, content)

	k.SetFee(ctx, fee)
	k.SetFeeID(ctx, feeId+1)

	return fee, nil
}

func (k Keeper) SetFee(ctx sdk.Context, fee types.Fee) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetFeeKey(fee.Id)
	store.Set(key, k.cdc.MustMarshalBinaryLengthPrefixed(fee))
}

// GetFeeID gets the highest fee ID
func (k Keeper) GetFeeID(ctx sdk.Context) (feeId uint64, err sdk.Error) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.FeeIdKey)
	if bz == nil {
		return 0, types.ErrInvalidGenesis(types.DefaultCodespace, "initial fee ID hasn't been set")
	}
	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &feeId)
	return feeId, nil
}

// Set the fee ID
func (k Keeper) SetFeeID(ctx sdk.Context, feeId uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(feeId)
	store.Set(types.FeeIdKey, bz)
}

// -------------------------------------------------------- FeeContracts

func (k Keeper) GetFeeContractIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, types.FeeContractKeyPrefix)
}

func (k Keeper) MustGetFeeContractByKey(ctx sdk.Context, key []byte) types.FeeContract {
	store := ctx.KVStore(k.storeKey)
	if !store.Has(key) {
		panic("fee contract not found")
	}

	bz := store.Get(key)
	var feeContract types.FeeContract
	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &feeContract)

	return feeContract
}

func (k Keeper) FeeContractExists(ctx sdk.Context, feeContractId uint64) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetFeeContractKey(feeContractId))
}

func (k Keeper) GetFeeContract(ctx sdk.Context, feeContractId uint64) (types.FeeContract, sdk.Error) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetFeeContractKey(feeContractId)

	bz := store.Get(key)
	if bz == nil {
		return types.FeeContract{}, sdk.ErrInternal("invalid fee contract")
	}

	var feeContract types.FeeContract
	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &feeContract)

	return feeContract, nil
}

func (k Keeper) SubmitFeeContract(ctx sdk.Context, content types.FeeContractContent) (types.FeeContract, sdk.Error) {
	feeContractId, err := k.GetFeeContractID(ctx)
	if err != nil {
		return types.FeeContract{}, err
	}

	feeContract := types.NewFeeContract(feeContractId, content)

	k.SetFeeContract(ctx, feeContract)
	k.SetFeeContractID(ctx, feeContractId+1)

	return feeContract, nil
}

func (k Keeper) SetFeeContract(ctx sdk.Context, feeContract types.FeeContract) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetFeeContractKey(feeContract.Id)
	store.Set(key, k.cdc.MustMarshalBinaryLengthPrefixed(feeContract))
}

// GetFeeContractID gets the highest fee contract ID
func (k Keeper) GetFeeContractID(ctx sdk.Context) (feeContractId uint64, err sdk.Error) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.FeeContractIdKey)
	if bz == nil {
		return 0, types.ErrInvalidGenesis(types.DefaultCodespace, "initial fee contract ID hasn't been set")
	}
	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &feeContractId)
	return feeContractId, nil
}

// Set the fee contract ID
func (k Keeper) SetFeeContractID(ctx sdk.Context, feeContractId uint64) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(feeContractId)
	store.Set(types.FeeContractIdKey, bz)
}

func (k Keeper) SetFeeContractAuthorised(ctx sdk.Context, feeContractId uint64,
	authorised bool) sdk.Error {
	feeContract, err := k.GetFeeContract(ctx, feeContractId)
	if err != nil {
		return err
	}

	// If de-authorising, check if can be de-authorised
	if !authorised && !feeContract.Content.CanDeauthorise {
		return types.ErrFeeContractCannotBeDeauthorised(types.DefaultCodespace)
	}

	// Set authorised state
	feeContract.Content.Authorised = authorised
	k.SetFeeContract(ctx, feeContract)

	return nil
}
