package project

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"reflect"

	sdk "github.com/cosmos/cosmos-sdk/types"
	wire "github.com/cosmos/cosmos-sdk/wire"
	"github.com/ixofoundation/ixo-cosmos/x/ixo"
)

const AccountMapPrefix = "ACC-"

// Implements ProjectMapper.
// This ProjectMapper encodes/decodes accounts using the
// go-wire (binary) encoding/decoding library.
type ProjectMapper struct {

	// The (unexposed) key used to access the store from the Context.
	key sdk.StoreKey

	// The prototypical ixo.ProjectDoc concrete type.
	proto ixo.StoredProjectDoc

	// The wire codec for binary encoding/decoding of projects.
	cdc *wire.Codec
}

// Create and return a sealed account mapper
func NewProjectMapperSealed(key sdk.StoreKey, proto ixo.StoredProjectDoc) SealedProjectMapper {
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

// Returns the array of account maps for a project
func (pm ProjectMapper) GetAccountMap(ctx sdk.Context, projectDid ixo.Did) map[string]interface{} {
	store := ctx.KVStore(pm.key)
	key := generateAccountsKey(projectDid)
	bz := store.Get(key)
	if bz == nil {
		return make(map[string]interface{})
	}

	didMap := pm.decodeAccountMap(bz)
	return didMap
}

func (pm ProjectMapper) AddAccountToAccountMap(ctx sdk.Context, projectDid ixo.Did, accountDid ixo.Did, accountAddr sdk.Address) {
	accMap := pm.GetAccountMap(ctx, projectDid)
	_, found := accMap[accountDid]
	if found {
		return
	}

	store := ctx.KVStore(pm.key)
	key := generateAccountsKey(projectDid)
	accountAddrString := hex.EncodeToString(accountAddr)
	accMap[string(accountDid)] = accountAddrString
	bz := pm.encodeAccountMap(accMap)
	store.Set(key, bz)
}

func (pm ProjectMapper) decodeAccountMap(accMapBytes []byte) map[string]interface{} {
	jsonBytes := []byte(accMapBytes)
	var f interface{}
	err := json.Unmarshal(jsonBytes, &f)
	if err != nil {
		panic(err)
	}
	m := f.(map[string]interface{})
	return m
}

func (pm ProjectMapper) encodeAccountMap(accMap map[string]interface{}) []byte {
	json, err := json.Marshal(accMap)
	if err != nil {
		panic(err)
	}
	return []byte(json)
}

func (pm ProjectMapper) GetProjectDoc(ctx sdk.Context, addr ixo.Did) (ixo.StoredProjectDoc, bool) {
	store := ctx.KVStore(pm.key)
	bz := store.Get([]byte(addr))
	if bz == nil {
		return nil, false
	}
	project := pm.decodeProject(bz)
	return project, true
}

func (pm ProjectMapper) SetProjectDoc(ctx sdk.Context, project ixo.StoredProjectDoc) {
	addr := []byte(project.GetProjectDid())
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
func (pm ProjectMapper) clonePrototype() ixo.StoredProjectDoc {
	protoRt := reflect.TypeOf(pm.proto)
	if protoRt.Kind() == reflect.Ptr {
		protoCrt := protoRt.Elem()
		if protoCrt.Kind() != reflect.Struct {
			panic("ProjectMapper requires a struct proto ixo.StoredProjectDoc, or a pointer to one")
		}
		protoRv := reflect.New(protoCrt)
		clone, ok := protoRv.Interface().(ixo.StoredProjectDoc)
		if !ok {
			panic(fmt.Sprintf("accountMapper requires a proto ixo.ProjectDoc, but %v doesn't implement ixo.StoredProjectDoc", protoRt))
		}
		return clone
	} else {
		protoRv := reflect.New(protoRt).Elem()
		clone, ok := protoRv.Interface().(ixo.StoredProjectDoc)
		if !ok {
			panic(fmt.Sprintf("accountMapper requires a proto ixo.ProjectDoc, but %v doesn't implement ixo.StoredProjectDoc", protoRt))
		}
		return clone
	}
}

func (pm ProjectMapper) encodeProject(storedProjectDoc ixo.StoredProjectDoc) []byte {
	bz, err := pm.cdc.MarshalBinary(storedProjectDoc)
	if err != nil {
		panic(err)
	}
	return bz
}

func (pm ProjectMapper) decodeProject(bz []byte) ixo.StoredProjectDoc {

	storedProjectDoc := StoredProjectDoc{}
	err := pm.cdc.UnmarshalBinary(bz, &storedProjectDoc)
	if err != nil {
		panic(err)
	}
	return storedProjectDoc

}
func generateAccountsKey(did ixo.Did) []byte {
	var buffer bytes.Buffer
	buffer.WriteString(AccountMapPrefix)
	buffer.WriteString(did)
	return buffer.Bytes()
}
