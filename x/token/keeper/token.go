package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	iidtypes "github.com/ixofoundation/ixo-blockchain/x/iid/types"
	"github.com/ixofoundation/ixo-blockchain/x/token/types"
)

func (k Keeper) SetToken(ctx sdk.Context, value types.Token) {
	key := fmt.Sprintf("%s#%s", value.MinterDid.Did(), value.ContractAddress)
	k.Set(ctx, []byte(key), types.TokenKey, value, k.Marshal)
}

func (k Keeper) GetToken(ctx sdk.Context, minterDid iidtypes.DIDFragment, contractAddress string) (types.Token, error) {
	key := fmt.Sprintf("%s#%s", minterDid.Did(), contractAddress)
	val, found := k.Get(ctx, []byte(key), types.TokenKey, k.UnmarshalToken)
	if !found {
		return types.Token{}, sdkerrors.Wrapf(types.ErrTokenNotFound, "token not found for %s", minterDid)
	}
	return val.(types.Token), nil
}

func (k Keeper) UnmarshalToken(value []byte) (interface{}, bool) {
	data := types.Token{}
	k.Unmarshal(value, &data)
	return data, types.IsValidToken(&data)
}

func (k Keeper) SetTokenProperties(ctx sdk.Context, value types.TokenProperties) {
	k.Set(ctx, []byte(value.Id), types.TokenPropertiesKey, value, k.Marshal)
}

func (k Keeper) GetTokenProperties(ctx sdk.Context, id string) (types.TokenProperties, error) {
	val, found := k.Get(ctx, []byte(id), types.TokenKey, k.UnmarshalTokenProperties)
	if !found {
		return types.TokenProperties{}, sdkerrors.Wrapf(types.ErrTokenPropertiesNotFound, "token properties not found for %s", id)
	}
	return val.(types.TokenProperties), nil
}

func (k Keeper) UnmarshalTokenProperties(value []byte) (interface{}, bool) {
	data := types.TokenProperties{}
	k.Unmarshal(value, &data)
	return data, types.IsValidTokenProperties(&data)
}

func (k Keeper) GetMinterTokens(ctx sdk.Context, minterDid string) []*types.Token {
	iterator := k.GetAll(ctx, append([]byte(types.TokenKey), []byte(minterDid)...))
	minterTokens := []*types.Token{}
	for ; iterator.Valid(); iterator.Next() {
		var minterToken types.Token
		k.cdc.MustUnmarshal(iterator.Value(), &minterToken)
		minterTokens = append(minterTokens, &minterToken)
	}

	return minterTokens
}

// helper function to check if there are any tokens with provded name, return true if it is a duplicate name
func (k Keeper) CheckTokensDuplicateName(ctx sdk.Context, name string) bool {
	iterator := k.GetTokenIterator(ctx)
	for ; iterator.Valid(); iterator.Next() {
		var token types.Token
		k.cdc.MustUnmarshal(iterator.Value(), &token)
		if token.Name == name {
			return true
		}
	}
	return false
}

func (k Keeper) GetTokens(ctx sdk.Context, minterDid string) []*types.Token {
	iterator := k.GetTokenIterator(ctx)
	tokens := []*types.Token{}
	for ; iterator.Valid(); iterator.Next() {
		var token types.Token
		k.cdc.MustUnmarshal(iterator.Value(), &token)
		tokens = append(tokens, &token)
	}

	return tokens
}

func (k Keeper) GetTokenIterator(ctx sdk.Context) sdk.Iterator {
	return k.GetAll(ctx, append([]byte(types.TokenKey)))
}
