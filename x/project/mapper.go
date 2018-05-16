package project

import (
	"fmt"
	"reflect"

	sdk "github.com/cosmos/cosmos-sdk/types"
	wire "github.com/cosmos/cosmos-sdk/wire"
	"github.com/ixofoundation/ixo-cosmos/x/ixo"
)

const AllKey = "ALL"

// Implements ProjectMapper.
// This ProjectMapper encodes/decodes accounts using the
// go-wire (binary) encoding/decoding library.
type ProjectMapper struct {

	// The (unexposed) key used to access the store from the Context.
	key sdk.StoreKey

	// The prototypical ixo.ProjectDoc concrete type.
	proto ixo.ProjectDoc

	// The wire codec for binary encoding/decoding of projects.
	cdc *wire.Codec
}

// Create and return a sealed account mapper
func NewProjectMapperSealed(key sdk.StoreKey, proto ixo.ProjectDoc) SealedProjectMapper {
	cdc := wire.NewCodec()
	pm := ProjectMapper{
		key:   key,
		proto: proto,
		cdc:   cdc,
	}
	wire.RegisterCrypto(cdc)

	// make ProjectMapper's WireCodec() inaccessible, return
	return pm.Seal()
}

type SealedProjectMapper struct {
	ProjectMapper
}

// Returns the go-wire codec.  You may need to register interfaces
// and concrete types here, if your app's ixo.ProjectDoc
// implementation includes interface fields.
// NOTE: It is not secure to expose the codec, so check out
// .Seal().
func (pm ProjectMapper) WireCodec() *wire.Codec {
	return pm.cdc
}

// Returns a "sealed" ProjectMapper.
// The codec is not accessible from a sealedProjectMapper.
func (pm ProjectMapper) Seal() SealedProjectMapper {
	return SealedProjectMapper{pm}
}

func (pm ProjectMapper) NewProjectDoc(ctx sdk.Context, msg AddProjectMsg) ixo.ProjectDoc {
	project := pm.clonePrototype()
	project.SetDid(msg.ProjectDoc.GetDid())
	project.SetCreatedBy(msg.ProjectDoc.GetCreatedBy())
	return project
}

func (pm ProjectMapper) GetProjectDoc(ctx sdk.Context, addr ixo.Did) ixo.ProjectDoc {
	store := ctx.KVStore(pm.key)
	bz := store.Get([]byte(addr))
	if bz == nil {
		return nil
	}
	project := pm.decodeProject(bz)
	return project
}

func (pm ProjectMapper) GetAllDids(ctx sdk.Context) []ixo.Did {
	store := ctx.KVStore(pm.key)
	bz := store.Get([]byte(AllKey))
	if bz == nil {
		return []ixo.Did{}
	} else {
		dids := []ixo.Did{}
		err := pm.cdc.UnmarshalBinary(bz, &dids)
		if err != nil {
			panic(err)
		}
		return dids
	}
}

func (pm ProjectMapper) SetProjectDoc(ctx sdk.Context, project ixo.ProjectDoc) {
	addr := []byte(project.GetDid())
	store := ctx.KVStore(pm.key)
	bz := pm.encodeProject(project)
	store.Set(addr, bz)
}

// There's no way for external modules to mutate the
// spm.ProjectMapper.ctx from here, even with reflection.
func (spm SealedProjectMapper) WireCodec() *wire.Codec {
	panic("ProjectMapper is sealed")
}

// Creates a new struct (or pointer to struct) from am.proto.
func (pm ProjectMapper) clonePrototype() ixo.ProjectDoc {
	protoRt := reflect.TypeOf(pm.proto)
	if protoRt.Kind() == reflect.Ptr {
		protoCrt := protoRt.Elem()
		if protoCrt.Kind() != reflect.Struct {
			panic("ProjectMapper requires a struct proto ixo.ProjectDoc, or a pointer to one")
		}
		protoRv := reflect.New(protoCrt)
		clone, ok := protoRv.Interface().(ixo.ProjectDoc)
		if !ok {
			panic(fmt.Sprintf("accountMapper requires a proto ixo.ProjectDoc, but %v doesn't implement ixo.ProjectDoc", protoRt))
		}
		return clone
	} else {
		protoRv := reflect.New(protoRt).Elem()
		clone, ok := protoRv.Interface().(ixo.ProjectDoc)
		if !ok {
			panic(fmt.Sprintf("accountMapper requires a proto ixo.ProjectDoc, but %v doesn't implement ixo.ProjectDoc", protoRt))
		}
		return clone
	}
}

func (pm ProjectMapper) encodeProject(projectDoc ixo.ProjectDoc) []byte {
	bz, err := pm.cdc.MarshalBinary(projectDoc)
	if err != nil {
		panic(err)
	}
	return bz
}

func (pm ProjectMapper) decodeProject(bz []byte) ixo.ProjectDoc {

	projectDoc := ProjectDoc{}
	err := pm.cdc.UnmarshalBinary(bz, &projectDoc)
	if err != nil {
		panic(err)
	}
	return projectDoc

}

func (pm ProjectMapper) appendDidToAll(ctx sdk.Context, newDid ixo.Did) {
	dids := pm.GetAllDids(ctx)
	store := ctx.KVStore(pm.key)
	newDids := append(dids, newDid)
	bz, err := pm.cdc.MarshalBinary(newDids)
	if err != nil {
		panic(err)
	}
	store.Set([]byte(AllKey), bz)
}
