package keeper

import (
	"fmt"
	"reflect"
	
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Keeper struct {
	cdc *codec.Codec
	key sdk.StoreKey
}

func NewKeeper(cdc *codec.Codec, key sdk.StoreKey) Keeper {
	return Keeper{
		cdc: cdc,
		key: key,
	}
}

func InitKeeper(ctx sdk.Context, cdc *codec.Codec, key sdk.StoreKey, params ...interface{}) Keeper {
	if len(params)%2 != 0 {
		panic("Odd params list length for InitKeeper")
	}
	
	k := NewKeeper(cdc, key)
	
	for i := 0; i < len(params); i += 2 {
		k.Set(ctx, params[i].(string), params[i+1])
	}
	
	return k
}

func (k Keeper) get(ctx sdk.Context, key string, ptr interface{}) error {
	store := ctx.KVStore(k.key)
	bz := store.Get([]byte(key))
	return k.cdc.UnmarshalBinaryLengthPrefixed(bz, ptr)
}

func (k Keeper) getRaw(ctx sdk.Context, key string) []byte {
	store := ctx.KVStore(k.key)
	return store.Get([]byte(key))
}

func (k Keeper) Set(ctx sdk.Context, key string, param interface{}) error {
	store := ctx.KVStore(k.key)
	bz := store.Get([]byte(key))
	if bz != nil {
		ptrty := reflect.PtrTo(reflect.TypeOf(param))
		ptr := reflect.New(ptrty).Interface()
		
		if k.cdc.UnmarshalBinaryLengthPrefixed(bz, ptr) != nil {
			return fmt.Errorf("Type mismatch with stored param and provided param")
		}
	}
	
	bz, err := k.cdc.MarshalBinaryLengthPrefixed(param)
	if err != nil {
		return err
	}
	store.Set([]byte(key), bz)
	return nil
}

func (k Keeper) setRaw(ctx sdk.Context, key string, param []byte) {
	store := ctx.KVStore(k.key)
	store.Set([]byte(key), param)
}

func (k Keeper) Getter() Getter {
	return Getter{k}
}

func (k Keeper) Setter() Setter {
	return Setter{Getter{k}}
}

type Getter struct {
	k Keeper
}

func (k Getter) Get(ctx sdk.Context, key string, ptr interface{}) error {
	return k.k.get(ctx, key, ptr)
}

func (k Getter) GetRaw(ctx sdk.Context, key string) []byte {
	return k.k.getRaw(ctx, key)
}

func (k Getter) GetString(ctx sdk.Context, key string) (res string, err error) {
	store := ctx.KVStore(k.k.key)
	bz := store.Get([]byte(key))
	err = k.k.cdc.UnmarshalBinaryLengthPrefixed(bz, &res)
	return
}

func (k Getter) GetBool(ctx sdk.Context, key string) (res bool, err error) {
	store := ctx.KVStore(k.k.key)
	bz := store.Get([]byte(key))
	err = k.k.cdc.UnmarshalBinaryLengthPrefixed(bz, &res)
	return
}

func (k Getter) GetInt16(ctx sdk.Context, key string) (res int16, err error) {
	store := ctx.KVStore(k.k.key)
	bz := store.Get([]byte(key))
	err = k.k.cdc.UnmarshalBinaryLengthPrefixed(bz, &res)
	return
}

func (k Getter) GetInt32(ctx sdk.Context, key string) (res int32, err error) {
	store := ctx.KVStore(k.k.key)
	bz := store.Get([]byte(key))
	err = k.k.cdc.UnmarshalBinaryLengthPrefixed(bz, &res)
	return
}

func (k Getter) GetInt64(ctx sdk.Context, key string) (res int64, err error) {
	store := ctx.KVStore(k.k.key)
	bz := store.Get([]byte(key))
	err = k.k.cdc.UnmarshalBinaryLengthPrefixed(bz, &res)
	return
}

func (k Getter) GetUint16(ctx sdk.Context, key string) (res uint16, err error) {
	store := ctx.KVStore(k.k.key)
	bz := store.Get([]byte(key))
	err = k.k.cdc.UnmarshalBinaryLengthPrefixed(bz, &res)
	return
}

func (k Getter) GetUint32(ctx sdk.Context, key string) (res uint32, err error) {
	store := ctx.KVStore(k.k.key)
	bz := store.Get([]byte(key))
	err = k.k.cdc.UnmarshalBinaryLengthPrefixed(bz, &res)
	return
}

func (k Getter) GetUint64(ctx sdk.Context, key string) (res uint64, err error) {
	store := ctx.KVStore(k.k.key)
	bz := store.Get([]byte(key))
	err = k.k.cdc.UnmarshalBinaryLengthPrefixed(bz, &res)
	return
}

func (k Getter) GetDec(ctx sdk.Context, key string) (res sdk.Dec, err error) {
	store := ctx.KVStore(k.k.key)
	bz := store.Get([]byte(key))
	err = k.k.cdc.UnmarshalBinaryLengthPrefixed(
		bz, &res)
	return
}

func (k Getter) GetUint(ctx sdk.Context, key string) (res sdk.Uint, err error) {
	store := ctx.KVStore(k.k.key)
	bz := store.Get([]byte(key))
	err = k.k.cdc.UnmarshalBinaryLengthPrefixed(bz, &res)
	return
}

func (k Getter) GetRat(ctx sdk.Context, key string) (res sdk.Int, err error) {
	store := ctx.KVStore(k.k.key)
	bz := store.Get([]byte(key))
	err = k.k.cdc.UnmarshalBinaryLengthPrefixed(bz, &res)
	return
}

func (k Getter) GetStringWithDefault(ctx sdk.Context, key string, def string) (res string) {
	store := ctx.KVStore(k.k.key)
	bz := store.Get([]byte(key))
	if bz == nil {
		return def
	}
	k.k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &res)
	return
}

func (k Getter) GetBoolWithDefault(ctx sdk.Context, key string, def bool) (res bool) {
	store := ctx.KVStore(k.k.key)
	bz := store.Get([]byte(key))
	if bz == nil {
		return def
	}
	k.k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &res)
	return
}

func (k Getter) GetInt16WithDefault(ctx sdk.Context, key string, def int16) (res int16) {
	store := ctx.KVStore(k.k.key)
	bz := store.Get([]byte(key))
	if bz == nil {
		return def
	}
	k.k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &res)
	return
}

func (k Getter) GetInt32WithDefault(ctx sdk.Context, key string, def int32) (res int32) {
	store := ctx.KVStore(k.k.key)
	bz := store.Get([]byte(key))
	if bz == nil {
		return def
	}
	k.k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &res)
	return
}

func (k Getter) GetInt64WithDefault(ctx sdk.Context, key string, def int64) (res int64) {
	store := ctx.KVStore(k.k.key)
	bz := store.Get([]byte(key))
	if bz == nil {
		return def
	}
	k.k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &res)
	return
}

func (k Getter) GetUint16WithDefault(ctx sdk.Context, key string, def uint16) (res uint16) {
	store := ctx.KVStore(k.k.key)
	bz := store.Get([]byte(key))
	if bz == nil {
		return def
	}
	k.k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &res)
	return
}

func (k Getter) GetUint32WithDefault(ctx sdk.Context, key string, def uint32) (res uint32) {
	store := ctx.KVStore(k.k.key)
	bz := store.Get([]byte(key))
	if bz == nil {
		return def
	}
	k.k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &res)
	return
}

func (k Getter) GetUint64WithDefault(ctx sdk.Context, key string, def uint64) (res uint64) {
	store := ctx.KVStore(k.k.key)
	bz := store.Get([]byte(key))
	if bz == nil {
		return def
	}
	k.k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &res)
	return
}

func (k Getter) GetIntWithDefault(ctx sdk.Context, key string, def sdk.Int) (res sdk.Int) {
	store := ctx.KVStore(k.k.key)
	bz := store.Get([]byte(key))
	if bz == nil {
		return def
	}
	k.k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &res)
	return
}

func (k Getter) GetUintWithDefault(ctx sdk.Context, key string, def sdk.Uint) (res sdk.Uint) {
	store := ctx.KVStore(k.k.key)
	bz := store.Get([]byte(key))
	if bz == nil {
		return def
	}
	k.k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &res)
	return
}

func (k Getter) GetRatWithDefault(ctx sdk.Context, key string, def sdk.Int) (res sdk.Int) {
	store := ctx.KVStore(k.k.key)
	bz := store.Get([]byte(key))
	if bz == nil {
		return def
	}
	k.k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &res)
	return
}

type Setter struct {
	Getter
}

func (k Setter) Set(ctx sdk.Context, key string, param interface{}) error {
	return k.k.Set(ctx, key, param)
}

func (k Setter) SetRaw(ctx sdk.Context, key string, param []byte) {
	k.k.setRaw(ctx, key, param)
}

func (k Setter) SetString(ctx sdk.Context, key string, param string) {
	if err := k.k.Set(ctx, key, param); err != nil {
		panic(err)
	}
}

func (k Setter) SetBool(ctx sdk.Context, key string, param bool) {
	if err := k.k.Set(ctx, key, param); err != nil {
		panic(err)
	}
}

func (k Setter) SetInt16(ctx sdk.Context, key string, param int16) {
	if err := k.k.Set(ctx, key, param); err != nil {
		panic(err)
	}
}

func (k Setter) SetInt32(ctx sdk.Context, key string, param int32) {
	if err := k.k.Set(ctx, key, param); err != nil {
		panic(err)
	}
}

func (k Setter) SetInt64(ctx sdk.Context, key string, param int64) {
	if err := k.k.Set(ctx, key, param); err != nil {
		panic(err)
	}
}

func (k Setter) SetUint16(ctx sdk.Context, key string, param uint16) {
	if err := k.k.Set(ctx, key, param); err != nil {
		panic(err)
	}
}

func (k Setter) SetUint32(ctx sdk.Context, key string, param uint32) {
	if err := k.k.Set(ctx, key, param); err != nil {
		panic(err)
	}
}

func (k Setter) SetUint64(ctx sdk.Context, key string, param uint64) {
	if err := k.k.Set(ctx, key, param); err != nil {
		panic(err)
	}
}

func (k Setter) SetDec(ctx sdk.Context, key string, param sdk.Dec) {
	if err := k.k.Set(ctx, key, param); err != nil {
		panic(err)
	}
}

func (k Setter) SetUint(ctx sdk.Context, key string, param sdk.Uint) {
	if err := k.k.Set(ctx, key, param); err != nil {
		panic(err)
	}
}
