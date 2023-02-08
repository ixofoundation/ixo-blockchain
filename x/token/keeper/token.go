package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	iidtypes "github.com/ixofoundation/ixo-blockchain/x/iid/types"
	"github.com/ixofoundation/ixo-blockchain/x/token/types"
)

func (k Keeper) SetMinter(ctx sdk.Context, value types.TokenMinter) error {
	minterBytes, err := value.Marshal()
	if err != nil {
		return err
	}

	key := fmt.Sprintf("%s#%s", value.MinterDid.Did(), value.ContractAddress)
	ctx.KVStore(k.storeKey).Set([]byte(key), minterBytes)
	return nil
}

func (k Keeper) GetMinterContract(ctx sdk.Context, minterDid iidtypes.DIDFragment, contractAddress string) (types.TokenMinter, error) {
	key := fmt.Sprintf("%s#%s", minterDid.Did(), contractAddress)
	raw := ctx.KVStore(k.storeKey).Get([]byte(key))
	var minterContract types.TokenMinter
	err := k.cdc.Unmarshal(raw, &minterContract)
	if err != nil {
		return types.TokenMinter{}, err
	}

	return minterContract, nil
}

func (k Keeper) GetMinterContracts(ctx sdk.Context, minterDid string) []*types.TokenMinter {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, []byte(minterDid))

	minterContracts := []*types.TokenMinter{}
	for ; iterator.Valid(); iterator.Next() {
		var minterContract types.TokenMinter
		k.cdc.MustUnmarshal(iterator.Value(), &minterContract)
		minterContracts = append(minterContracts, &minterContract)
	}

	return minterContracts
}

func (k Keeper) TokenExists(ctx sdk.Context, tokenDid string) bool {
	// store := ctx.KVStore(k.storeKey)
	_, exists := k.IidKeeper.GetDidDocument(ctx, []byte(tokenDid))
	return exists
}

func (k Keeper) GetTokenDocIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, types.TokenKey)
}
