package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"github.com/ixofoundation/ixo-blockchain/x/claims/types"
)

func (k Keeper) SetCollection(ctx sdk.Context, data types.Collection) {
	k.Set(ctx, []byte(data.Id), types.CollectionKey, data, k.Marshal)
}

func (k Keeper) GetCollection(ctx sdk.Context, id string) (types.Collection, error) {
	val, found := k.Get(ctx, []byte(id), types.CollectionKey, k.UnmarshalCollection)
	if !found {
		return types.Collection{}, sdkerrors.Wrapf(types.ErrCollectionNotFound, "for %s", id)
	}
	return val.(types.Collection), nil
}

func (k Keeper) UnmarshalCollection(value []byte) (interface{}, bool) {
	data := types.Collection{}
	k.Unmarshal(value, &data)
	return data, types.IsValidCollection(&data)
}

func (k Keeper) Marshal(value interface{}) (bytes []byte) {
	switch value := value.(type) {
	case types.Collection:
		bytes = k.cdc.MustMarshal(&value)
	case types.Claim:
		bytes = k.cdc.MustMarshal(&value)
	case types.Dispute:
		bytes = k.cdc.MustMarshal(&value)
	}
	return
}

// Unmarshal unmarshal a byte slice to a struct, return false in case of errors
func (k Keeper) Unmarshal(data []byte, val codec.ProtoMarshaler) bool {
	if len(data) == 0 {
		return false
	}
	if err := k.cdc.Unmarshal(data, val); err != nil {
		return false
	}
	return true
}

func (k Keeper) SetClaim(ctx sdk.Context, data types.Claim) {
	k.Set(ctx, []byte(data.ClaimId), types.ClaimKey, data, k.Marshal)
}

func (k Keeper) GetClaim(ctx sdk.Context, id string) (types.Claim, error) {
	val, found := k.Get(ctx, []byte(id), types.ClaimKey, k.UnmarshalClaim)
	if !found {
		return types.Claim{}, sdkerrors.Wrapf(types.ErrClaimNotFound, "for %s", id)
	}
	return val.(types.Claim), nil
}

func (k Keeper) UnmarshalClaim(value []byte) (interface{}, bool) {
	data := types.Claim{}
	k.Unmarshal(value, &data)
	return data, types.IsValidClaim(&data)
}

func (k Keeper) SetDispute(ctx sdk.Context, data types.Dispute) {
	k.Set(ctx, []byte(data.Data.Proof), types.DisputeKey, data, k.Marshal)
}

func (k Keeper) GetDispute(ctx sdk.Context, proof string) (types.Dispute, error) {
	val, found := k.Get(ctx, []byte(proof), types.DisputeKey, k.UnmarshalDispute)
	if !found {
		return types.Dispute{}, sdkerrors.Wrapf(types.ErrDisputeNotFound, "for proof %s", proof)
	}
	return val.(types.Dispute), nil
}

func (k Keeper) UnmarshalDispute(value []byte) (interface{}, bool) {
	data := types.Dispute{}
	k.Unmarshal(value, &data)
	return data, types.IsValidDispute(&data)
}
