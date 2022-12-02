package keeper

import (
	"fmt"

	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"
	authzkeeper "github.com/cosmos/cosmos-sdk/x/authz/keeper"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	iidkeeper "github.com/ixofoundation/ixo-blockchain/x/iid/keeper"
	iidtypes "github.com/ixofoundation/ixo-blockchain/x/iid/types"
	"github.com/ixofoundation/ixo-blockchain/x/token/types"
)

type Keeper struct {
	cdc           codec.BinaryCodec
	storeKey      sdk.StoreKey
	IidKeeper     iidkeeper.Keeper
	WasmKeeper    wasmtypes.ContractOpsKeeper
	AccountKeeper authkeeper.AccountKeeper
	AuthzKeeper   authzkeeper.Keeper
	ParamSpace    paramstypes.Subspace
}

func NewKeeper(cdc codec.BinaryCodec, key sdk.StoreKey, iidKeeper iidkeeper.Keeper, wasmKeeper wasmkeeper.Keeper,
	accountKeeper authkeeper.AccountKeeper, authzKeeper authzkeeper.Keeper, paramSpace paramstypes.Subspace) Keeper {

	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		cdc:           cdc,
		storeKey:      key,
		IidKeeper:     iidKeeper,
		WasmKeeper:    wasmkeeper.NewDefaultPermissionKeeper(wasmKeeper),
		AccountKeeper: accountKeeper,
		AuthzKeeper:   authzKeeper,
		ParamSpace:    paramSpace,
	}
}

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

// func (k Keeper) Mint()

func (k Keeper) TransferToken(ctx sdk.Context, msg *types.MsgTransferToken) (*types.MsgTransferTokenResponse, error) {
	return &types.MsgTransferTokenResponse{}, nil
}

// GetParams returns the total set of project parameters.
func (k Keeper) GetParams(ctx sdk.Context) (params types.Params) {
	k.ParamSpace.GetParamSet(ctx, &params)
	return params
}

// SetParams sets the total set of project parameters.
func (k Keeper) SetParams(ctx sdk.Context, params *types.Params) {
	k.ParamSpace.SetParamSet(ctx, params)
}

func (k Keeper) GetTokenDocIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, types.TokenKey)
}

func (k Keeper) TokenExists(ctx sdk.Context, tokenDid string) bool {
	// store := ctx.KVStore(k.storeKey)
	_, exists := k.IidKeeper.GetDidDocument(ctx, []byte(tokenDid))
	return exists
}
