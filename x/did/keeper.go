package did

import (
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/wire"
	"github.com/ixofoundation/ixo-cosmos/x/ixo"
)

// module users must specify coin denomination and reward (constant) per PoW solution
type Config struct {
	allKey string
}

// DidKeeper manages dids
type Keeper struct {
	key    sdk.StoreKey
	cdc    *wire.Codec
	config Config
}

func NewKeeper(cdc *wire.Codec, key sdk.StoreKey) Keeper {
	return Keeper{key, cdc, Config{"ALL"}}
}

// GetDidDoc returns the did_doc at the addr.
func (k Keeper) GetDidDoc(ctx sdk.Context, did ixo.Did) ixo.DidDoc {
	store := ctx.KVStore(k.key)
	bz := store.Get([]byte(did))
	if bz == nil {
		return nil
	}
	decodedDid := k.decodeDid(bz)
	return decodedDid
}

// GetAllDids returns all the dids.
func (k Keeper) GetAllDids(ctx sdk.Context) []ixo.Did {
	store := ctx.KVStore(k.key)
	bz := store.Get([]byte(k.config.allKey))
	if bz == nil {
		return []ixo.Did{}
	} else {
		dids := []ixo.Did{}
		err := k.cdc.UnmarshalBinary(bz, &dids)
		if err != nil {
			panic(err)
		}
		return dids
	}
}

// AddDidDoc adds the did_doc at the addr.
func (k Keeper) AddDidDoc(ctx sdk.Context, newDidDoc ixo.DidDoc) (ixo.DidDoc, sdk.Error) {
	didDoc := k.GetDidDoc(ctx, newDidDoc.GetDid())
	if didDoc == nil {
		k.SetDidDoc(ctx, newDidDoc)
		return newDidDoc, nil
	} else {
		panic(errors.New("Did already exists"))
	}
}

func (k Keeper) SetDidDoc(ctx sdk.Context, did ixo.DidDoc) {
	addr := []byte(did.GetDid())
	store := ctx.KVStore(k.key)
	bz := k.encodeDid(did)
	store.Set(addr, bz)
	k.appendDidToAll(ctx, ixo.Did(did.GetDid()))
}

func (k Keeper) AddCredential(ctx sdk.Context, did ixo.Did, credential DidCredential) (ixo.DidDoc, sdk.Error) {
	addr := []byte(did)
	didDoc := k.GetDidDoc(ctx, did)
	if didDoc == nil {
		panic(errors.New("Did does not exist"))
	}
	baseDidDoc := didDoc.(BaseDidDoc)
	credentials := baseDidDoc.GetCredentials()
	found := false
	for _, v := range credentials {
		if v.Issuer == credential.Issuer && v.CredType[0] == credential.CredType[0] && v.CredType[1] == credential.CredType[1] && v.Claim.KYCValidated == credential.Claim.KYCValidated {
			found = true
		}
	}
	if !found {
		baseDidDoc.AddCredential(credential)
		store := ctx.KVStore(k.key)
		bz := k.encodeDid(baseDidDoc)
		store.Set(addr, bz)
	}
	return baseDidDoc, nil
}

func (k Keeper) appendDidToAll(ctx sdk.Context, newDid ixo.Did) {
	dids := k.GetAllDids(ctx)
	store := ctx.KVStore(k.key)
	newDids := append(dids, newDid)
	bz, err := k.cdc.MarshalBinary(newDids)
	if err != nil {
		panic(err)
	}
	store.Set([]byte(k.config.allKey), bz)
}

func (k Keeper) encodeDid(didDoc ixo.DidDoc) []byte {
	bz, err := k.cdc.MarshalBinary(didDoc)
	if err != nil {
		panic(err)
	}
	return bz
}

func (k Keeper) decodeDid(bz []byte) ixo.DidDoc {

	didDoc := BaseDidDoc{}
	err := k.cdc.UnmarshalBinary(bz, &didDoc)
	if err != nil {
		panic(err)
	}
	return didDoc

}
