package did

import (
	"errors"
	"fmt"
	"reflect"

	sdk "github.com/cosmos/cosmos-sdk/types"
	wire "github.com/cosmos/cosmos-sdk/wire"
	"github.com/ixofoundation/ixo-cosmos/x/ixo"
)

const AllKey = "ALL"

// Implements DidMapper.
// This DidMapper encodes/decodes accounts using the
// go-wire (binary) encoding/decoding library.
type DidMapper struct {

	// The (unexposed) key used to access the store from the Context.
	key sdk.StoreKey

	// The prototypical ixo.DidDoc concrete type.
	proto ixo.DidDoc

	// The wire codec for binary encoding/decoding of dids.
	cdc *wire.Codec
}

// Create and return a sealed account mapper
func NewDidMapperSealed(key sdk.StoreKey, proto ixo.DidDoc) SealedDidMapper {
	cdc := wire.NewCodec()
	am := DidMapper{
		key:   key,
		proto: proto,
		cdc:   cdc,
	}
	wire.RegisterCrypto(cdc)

	// make DidMapper's WireCodec() inaccessible, return
	return am.Seal()
}

// Returns the go-wire codec.  You may need to register interfaces
// and concrete types here, if your app's ixo.DidDoc
// implementation includes interface fields.
// NOTE: It is not secure to expose the codec, so check out
// .Seal().
func (dm DidMapper) WireCodec() *wire.Codec {
	return dm.cdc
}

// Returns a "sealed" DidMapper.
// The codec is not accessible from a sealedDidMapper.
func (dm DidMapper) Seal() SealedDidMapper {
	return SealedDidMapper{dm}
}

func (dm DidMapper) NewDidDoc(ctx sdk.Context, msg AddDidMsg) ixo.DidDoc {
	return msg.DidDoc
}

func (dm DidMapper) GetDidDoc(ctx sdk.Context, addr ixo.Did) ixo.DidDoc {
	store := ctx.KVStore(dm.key)
	bz := store.Get([]byte(addr))
	if bz == nil {
		return nil
	}
	did := dm.decodeDid(bz)
	return did
}

func (dm DidMapper) GetAllDids(ctx sdk.Context) []ixo.Did {
	store := ctx.KVStore(dm.key)
	bz := store.Get([]byte(AllKey))
	if bz == nil {
		return []ixo.Did{}
	} else {
		dids := []ixo.Did{}
		err := dm.cdc.UnmarshalBinary(bz, &dids)
		if err != nil {
			panic(err)
		}
		return dids
	}
}

func (dm DidMapper) SetDidDoc(ctx sdk.Context, did ixo.DidDoc) {
	addr := []byte(did.GetDid())
	store := ctx.KVStore(dm.key)
	bz := dm.encodeDid(did)
	store.Set(addr, bz)
	dm.appendDidToAll(ctx, ixo.Did(did.GetDid()))
}

func (dm DidMapper) AddCredential(ctx sdk.Context, did ixo.Did, credential DidCredential) ixo.DidDoc {
	addr := []byte(did)
	didDoc := dm.GetDidDoc(ctx, did)
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
		store := ctx.KVStore(dm.key)
		bz := dm.encodeDid(baseDidDoc)
		store.Set(addr, bz)
	}
	return baseDidDoc
}

//----------------------------------------
// sealedDidMapper

type SealedDidMapper struct {
	DidMapper
}

// There's no way for external modules to mutate the
// sam.DidMapper.ctx from here, even with reflection.
func (sam SealedDidMapper) WireCodec() *wire.Codec {
	panic("DidMapper is sealed")
}

//----------------------------------------
// misc.

// Creates a new struct (or pointer to struct) from am.proto.
func (dm DidMapper) clonePrototype() ixo.DidDoc {
	protoRt := reflect.TypeOf(dm.proto)
	if protoRt.Kind() == reflect.Ptr {
		protoCrt := protoRt.Elem()
		if protoCrt.Kind() != reflect.Struct {
			panic("DidMapper requires a struct proto ixo.DidDoc, or a pointer to one")
		}
		protoRv := reflect.New(protoCrt)
		clone, ok := protoRv.Interface().(ixo.DidDoc)
		if !ok {
			panic(fmt.Sprintf("DidMapper requires a proto ixo.DidDoc, but %v doesn't implement ixo.DidDoc", protoRt))
		}
		return clone
	} else {
		protoRv := reflect.New(protoRt).Elem()
		clone, ok := protoRv.Interface().(ixo.DidDoc)
		if !ok {
			panic(fmt.Sprintf("DidMapper requires a proto ixo.DidDoc, but %v doesn't implement ixo.DidDoc", protoRt))
		}
		return clone
	}
}

func (dm DidMapper) encodeDid(didDoc ixo.DidDoc) []byte {
	bz, err := dm.cdc.MarshalBinary(didDoc)
	if err != nil {
		panic(err)
	}
	return bz
}

func (dm DidMapper) decodeDid(bz []byte) ixo.DidDoc {

	didDoc := BaseDidDoc{}
	err := dm.cdc.UnmarshalBinary(bz, &didDoc)
	if err != nil {
		panic(err)
	}
	return didDoc

}

func (dm DidMapper) appendDidToAll(ctx sdk.Context, newDid ixo.Did) {
	dids := dm.GetAllDids(ctx)
	store := ctx.KVStore(dm.key)
	newDids := append(dids, newDid)
	bz, err := dm.cdc.MarshalBinary(newDids)
	if err != nil {
		panic(err)
	}
	store.Set([]byte(AllKey), bz)
}
