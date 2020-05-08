package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ixofoundation/ixo-blockchain/x/did/internal/types"
	"github.com/ixofoundation/ixo-blockchain/x/ixo"
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

func (k Keeper) GetDidDoc(ctx sdk.Context, did ixo.Did) (ixo.DidDoc, sdk.Error) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetDidPrefixKey(did)
	bz := store.Get(key)
	if bz == nil {
		return nil, types.ErrorInvalidDid(types.DefaultCodeSpace, "Invalid Did Address")
	}

	var didDoc types.BaseDidDoc
	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &didDoc)

	return didDoc, nil
}

func (k Keeper) SetDidDoc(ctx sdk.Context, did ixo.DidDoc) (err sdk.Error) {
	existedDidDoc, err := k.GetDidDoc(ctx, did.GetDid())
	if existedDidDoc != nil {
		return types.ErrorInvalidDid(types.DefaultCodeSpace, "Did already exists")
	}

	k.AddDidDoc(ctx, did)
	return nil
}

func (k Keeper) AddDidDoc(ctx sdk.Context, did ixo.DidDoc) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetDidPrefixKey(did.GetDid())
	store.Set(key, k.cdc.MustMarshalBinaryLengthPrefixed(did))
}

func (k Keeper) AddCredentials(ctx sdk.Context, did ixo.Did, credential types.DidCredential) (err sdk.Error) {
	existedDid, err := k.GetDidDoc(ctx, did)
	if err != nil {
		return err
	}

	baseDidDoc := existedDid.(types.BaseDidDoc)
	credentials := baseDidDoc.GetCredentials()

	for _, data := range credentials {
		if data.Issuer == credential.Issuer && data.CredType[0] == credential.CredType[0] && data.CredType[1] == credential.CredType[1] && data.Claim.KYCValidated == credential.Claim.KYCValidated {
			return types.ErrorInvalidCredentials(types.DefaultCodeSpace, "credentials already exist")
		}
	}

	baseDidDoc.AddCredential(credential)
	k.AddDidDoc(ctx, baseDidDoc)

	return nil
}

func (k Keeper) GetAllDidDocs(ctx sdk.Context) (didDocs []ixo.DidDoc) {
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

func (k Keeper) GetAddDids(ctx sdk.Context) (dids []ixo.Did) {
	didDocs := k.GetAllDidDocs(ctx)
	for _, did := range didDocs {
		dids = append(dids, did.GetDid())
	}

	return dids
}
