package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ixofoundation/ixo-blockchain/x/did/exported"
	"github.com/ixofoundation/ixo-blockchain/x/did/internal/types"
)

type Keeper struct {
	storeKey sdk.StoreKey
	cdc      *codec.Codec
}

func NewKeeper(cdc *codec.Codec, key sdk.StoreKey) Keeper {
	return Keeper{
		storeKey: key,
		cdc:      cdc,
	}
}

func (k Keeper) GetDidDoc(ctx sdk.Context, did exported.Did) (exported.DidDoc, error) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetDidPrefixKey(did)
	bz := store.Get(key)
	if bz == nil {
		return nil, sdkerrors.Wrap(types.ErrorInvalidDid, did)
	}

	var didDoc types.BaseDidDoc
	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &didDoc)

	return didDoc, nil
}

func (k Keeper) MustGetDidDoc(ctx sdk.Context, did exported.Did) exported.DidDoc {
	didDoc, err := k.GetDidDoc(ctx, did)
	if err != nil {
		panic(err)
	}
	return didDoc
}

func (k Keeper) SetDidDoc(ctx sdk.Context, did exported.DidDoc) (err error) {
	existedDidDoc, err := k.GetDidDoc(ctx, did.GetDid())
	if existedDidDoc != nil {
		return sdkerrors.Wrap(types.ErrorInvalidDid, "Did already exists")
	}

	k.AddDidDoc(ctx, did)
	return nil
}

func (k Keeper) AddDidDoc(ctx sdk.Context, did exported.DidDoc) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetDidPrefixKey(did.GetDid())
	store.Set(key, k.cdc.MustMarshalBinaryLengthPrefixed(did))
}

func (k Keeper) AddCredentials(ctx sdk.Context, did exported.Did, credential exported.DidCredential) (err error) {
	existedDid, err := k.GetDidDoc(ctx, did)
	if err != nil {
		return err
	}

	baseDidDoc := existedDid.(types.BaseDidDoc)
	credentials := baseDidDoc.GetCredentials()

	for _, data := range credentials {
		if data.Issuer == credential.Issuer && data.CredType[0] == credential.CredType[0] && data.CredType[1] == credential.CredType[1] && data.Claim.KYCValidated == credential.Claim.KYCValidated {
			return sdkerrors.Wrap(types.ErrorInvalidCredentials, "credentials already exist")
		}
	}

	baseDidDoc.AddCredential(credential)
	k.AddDidDoc(ctx, baseDidDoc)

	return nil
}

func (k Keeper) GetAllDidDocs(ctx sdk.Context) (didDocs []exported.DidDoc) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.DidKey)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var didDoc types.BaseDidDoc
		k.cdc.MustUnmarshalBinaryLengthPrefixed(iterator.Value(), &didDoc)
		didDocs = append(didDocs, &didDoc)
	}

	return didDocs
}

func (k Keeper) GetAddDids(ctx sdk.Context) (dids []exported.Did) {
	didDocs := k.GetAllDidDocs(ctx)
	for _, did := range didDocs {
		dids = append(dids, did.GetDid())
	}

	return dids
}
