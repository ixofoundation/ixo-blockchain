package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	//"github.com/golang/protobuf/ptypes/any"
	"github.com/ixofoundation/ixo-blockchain/x/did/exported"
	"github.com/ixofoundation/ixo-blockchain/x/did/types"
	//codectypes "github.com/cosmos/cosmos-sdk/codec/types"
)

type Keeper struct {
	storeKey sdk.StoreKey
	cdc codec.BinaryMarshaler
}

func NewKeeper(cdc codec.BinaryMarshaler, key sdk.StoreKey) Keeper {
	return Keeper{
		storeKey: key,
		cdc:      cdc,
	}
}

// MarshalDidDoc protobuf serializes a DidDoc interface
func (k Keeper) MarshalDidDoc(did exported.DidDoc) ([]byte, error) {
	return k.cdc.MarshalInterface(did)
}

// UnmarshalDidDoc returns a DidDoc interface from raw encoded did document
// bytes of a Proto-based DidDoc type
func (k Keeper) UnmarshalDidDoc(bz []byte) (exported.DidDoc, error){
	var dd exported.DidDoc
	return dd, k.cdc.UnmarshalInterface(bz, &dd)
}

func (k Keeper) decodeDidDoc(bz []byte) exported.DidDoc {
	dd, err := k.UnmarshalDidDoc(bz)
	if err != nil {
		panic(err)
	}

	return dd
}

func (k Keeper) IterateDidDocs(ctx sdk.Context, cb func(account exported.DidDoc) (stop bool)) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.DidKey)

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		account := k.decodeDidDoc(iterator.Value())

		if cb(account) {
			break
		}
	}
}

func (k Keeper) GetDidDoc(ctx sdk.Context, did exported.Did) (exported.DidDoc, error) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetDidPrefixKey(did)
	bz := store.Get(key)
	if bz == nil {
		return nil, sdkerrors.Wrap(types.ErrInvalidDid, did)
	}

	return k.decodeDidDoc(bz), nil
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
		return sdkerrors.Wrap(types.ErrInvalidDid, "DID already exists")
	}

	k.AddDidDoc(ctx, did)
	return nil
}

func (k Keeper) AddDidDoc(ctx sdk.Context, did exported.DidDoc) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetDidPrefixKey(did.GetDid())

	dd, err := k.MarshalDidDoc(did)
	if err != nil {
		panic(err)
	}

	store.Set(key, dd)
}

func (k Keeper) AddCredentials(ctx sdk.Context, did exported.Did, credential types.DidCredential) (err error) {
	existedDid, err := k.GetDidDoc(ctx, did)
	if err != nil {
		return err
	}

	baseDidDoc := existedDid.(*types.BaseDidDoc)
	credentials := baseDidDoc.GetCredentials()

	for _, data := range credentials {
		if data.Issuer == credential.Issuer && data.CredType[0] == credential.CredType[0] && data.CredType[1] == credential.CredType[1] && data.Claim.KYCValidated == credential.Claim.KYCValidated {
			return sdkerrors.Wrap(types.ErrInvalidCredentials, "credentials already exist")
		}
	}

	baseDidDoc.AddCredential(&credential)
	k.AddDidDoc(ctx, baseDidDoc)

	return nil
}

func (k Keeper) GetAllDidDocs(ctx sdk.Context) (didDocs []exported.DidDoc) {
	k.IterateDidDocs(ctx, func(dd exported.DidDoc) (stop bool) {
		didDocs = append(didDocs, dd)
		return false
	})

	return didDocs
}

func (k Keeper) GetAllDids(ctx sdk.Context) (dids []exported.Did) {
	didDocs := k.GetAllDidDocs(ctx)
	for _, did := range didDocs {
		dids = append(dids, did.GetDid())
	}

	return dids
}
