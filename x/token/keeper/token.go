package keeper

import (
	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/store/prefix"
	storetypes "cosmossdk.io/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-blockchain/v5/x/token/types"
)

func tokenKey(minter, contract_address string) []byte {
	key := minter + contract_address
	return []byte(key)
}

func (k Keeper) SetToken(ctx sdk.Context, value types.Token) {
	key := tokenKey(value.Minter, value.ContractAddress)
	k.Set(ctx, key, types.TokenKey, value, k.Marshal)
}

func (k Keeper) GetToken(ctx sdk.Context, minter, contractAddress string) (types.Token, error) {
	key := tokenKey(minter, contractAddress)
	val, found := k.Get(ctx, key, types.TokenKey, k.UnmarshalToken)
	if !found {
		return types.Token{}, errorsmod.Wrapf(types.ErrTokenNotFound, "token not found minter %s and contract address %s", minter, contractAddress)
	}
	token, ok := val.(types.Token)
	if !ok {
		return types.Token{}, errorsmod.Wrapf(types.ErrTokenNotFound, "token not found minter %s and contract address %s", minter, contractAddress)
	}
	return token, nil
}

func (k Keeper) UnmarshalToken(value []byte) (interface{}, bool) {
	data := types.Token{}
	k.Unmarshal(value, &data)
	return data, types.IsValidToken(&data)
}

func (k Keeper) Marshal(value interface{}) (bytes []byte) {
	switch value := value.(type) {
	case types.Token:
		bytes = k.cdc.MustMarshal(&value)
	case types.TokenProperties:
		bytes = k.cdc.MustMarshal(&value)
	}
	return
}

func (k Keeper) SetTokenProperties(ctx sdk.Context, value types.TokenProperties) {
	k.Set(ctx, []byte(value.Id), types.TokenPropertiesKey, value, k.Marshal)
}

func (k Keeper) GetTokenProperties(ctx sdk.Context, id string) (types.TokenProperties, error) {
	val, found := k.Get(ctx, []byte(id), types.TokenPropertiesKey, k.UnmarshalTokenProperties)
	if !found {
		return types.TokenProperties{}, errorsmod.Wrapf(types.ErrTokenPropertiesNotFound, "token properties not found for %s", id)
	}
	tokenProperties, ok := val.(types.TokenProperties)
	if !ok {
		return types.TokenProperties{}, errorsmod.Wrapf(types.ErrTokenPropertiesNotFound, "token properties not found for %s", id)
	}
	return tokenProperties, nil
}

func (k Keeper) UnmarshalTokenProperties(value []byte) (interface{}, bool) {
	data := types.TokenProperties{}
	k.Unmarshal(value, &data)
	return data, types.IsValidTokenProperties(&data)
}

func (k Keeper) GetMinterTokens(ctx sdk.Context, minter string) []*types.Token {
	iterator := k.GetAll(ctx, append(types.TokenKey, []byte(minter)...))
	minterTokens := []*types.Token{}
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var minterToken types.Token
		k.cdc.MustUnmarshal(iterator.Value(), &minterToken)
		minterTokens = append(minterTokens, &minterToken)
	}

	return minterTokens
}

// helper function to get the Token and TokenProperties from tokenId
func (k Keeper) GetTokenById(ctx sdk.Context, id string) (*types.TokenProperties, *types.Token, error) {
	tokenProperties, err := k.GetTokenProperties(ctx, id)
	if err != nil {
		return nil, nil, errorsmod.Wrapf(err, "no TokenProperties for token %s", id)
	}

	token, found := k.GetTokenByName(ctx, tokenProperties.Name)
	if !found {
		return nil, nil, errorsmod.Wrapf(types.ErrTokenNotFound, "no Token found for name %s", tokenProperties.Name)
	}

	return &tokenProperties, token, err
}

// helper function to get the token with provided name
func (k Keeper) GetTokenByName(ctx sdk.Context, name string) (*types.Token, bool) {
	iterator := k.GetTokenIterator(ctx)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var token types.Token
		k.cdc.MustUnmarshal(iterator.Value(), &token)
		if token.Name == name {
			return &token, true
		}
	}
	return nil, false
}

// helper function to check if there are any tokens with provided name, return true if it is a duplicate name
func (k Keeper) CheckTokensDuplicateName(ctx sdk.Context, name string) bool {
	iterator := k.GetTokenIterator(ctx)
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var token types.Token
		k.cdc.MustUnmarshal(iterator.Value(), &token)
		if token.Name == name {
			return true
		}
	}
	return false
}

func (k Keeper) GetTokens(ctx sdk.Context) []types.Token {
	iterator := k.GetTokenIterator(ctx)
	tokens := []types.Token{}
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var token types.Token
		k.cdc.MustUnmarshal(iterator.Value(), &token)
		tokens = append(tokens, token)
	}

	return tokens
}

func (k Keeper) GetTokenIterator(ctx sdk.Context) storetypes.Iterator {
	return k.GetAll(ctx, types.TokenKey)
}

func (k Keeper) GetMinterTokensIterator(ctx sdk.Context, minter string) storetypes.Iterator {
	return k.GetAll(ctx, append(types.TokenKey, []byte(minter)...))
}

func (k Keeper) GetMinterTokensStore(ctx sdk.Context, minter string) prefix.Store {
	return prefix.NewStore(ctx.KVStore(k.storeKey), append(types.TokenKey, []byte(minter)...))
}

func (k Keeper) GetTokenPropertiesAll(ctx sdk.Context) []types.TokenProperties {
	iterator := k.GetTokenPropertiesAllIterator(ctx)
	tokenProperties := []types.TokenProperties{}
	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var tp types.TokenProperties
		k.cdc.MustUnmarshal(iterator.Value(), &tp)
		tokenProperties = append(tokenProperties, tp)
	}

	return tokenProperties
}

func (k Keeper) GetTokenPropertiesAllIterator(ctx sdk.Context) storetypes.Iterator {
	return k.GetAll(ctx, types.TokenPropertiesKey)
}
