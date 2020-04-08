package keeper

import (
	"encoding/json"
	"github.com/ixofoundation/ixo-cosmos/x/fees"
	"github.com/tendermint/tendermint/crypto"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"

	"github.com/ixofoundation/ixo-cosmos/x/did"
	"github.com/ixofoundation/ixo-cosmos/x/ixo"
	"github.com/ixofoundation/ixo-cosmos/x/project/internal/types"
)

type Keeper struct {
	cdc           *codec.Codec
	storeKey      sdk.StoreKey
	accountKeeper auth.AccountKeeper
	feeKeeper     fees.Keeper
}

func NewKeeper(cdc *codec.Codec, key sdk.StoreKey, accountKeeper auth.AccountKeeper, feeKeeper fees.Keeper) Keeper {
	return Keeper{
		cdc:           cdc,
		storeKey:      key,
		accountKeeper: accountKeeper,
		feeKeeper:     feeKeeper,
	}
}

func (k Keeper) GetProjectDocIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStorePrefixIterator(store, types.ProjectKey)
}

func (k Keeper) MustGetProjectDocByKey(ctx sdk.Context, key []byte) types.StoredProjectDoc {
	store := ctx.KVStore(k.storeKey)
	if !store.Has(key) {
		panic("project doc not found")
	}

	bz := store.Get(key)
	var projectDoc types.MsgCreateProject
	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &projectDoc)

	return &projectDoc
}

func (k Keeper) ProjectDocExists(ctx sdk.Context, projectDid ixo.Did) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetProjectPrefixKey(projectDid))
}

func (k Keeper) GetProjectDoc(ctx sdk.Context, projectDid ixo.Did) (types.StoredProjectDoc, sdk.Error) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetProjectPrefixKey(projectDid)

	bz := store.Get(key)
	if bz == nil {
		return nil, did.ErrorInvalidDid(types.DefaultCodeSpace, "Invalid ProjectDid Address")
	}

	var projectDoc types.MsgCreateProject
	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &projectDoc)

	return &projectDoc, nil
}

func (k Keeper) SetProjectDoc(ctx sdk.Context, projectDoc types.StoredProjectDoc) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetProjectPrefixKey(projectDoc.GetProjectDid())
	store.Set(key, k.cdc.MustMarshalBinaryLengthPrefixed(projectDoc))
}

func (k Keeper) UpdateProjectDoc(ctx sdk.Context, newProjectDoc types.StoredProjectDoc) (types.StoredProjectDoc, sdk.Error) {
	existedDoc, _ := k.GetProjectDoc(ctx, newProjectDoc.GetProjectDid())
	if existedDoc == nil {

		return nil, did.ErrorInvalidDid(types.DefaultCodeSpace, "ProjectDoc details are not exist")
	} else {

		existedDoc.SetStatus(newProjectDoc.GetStatus())
		k.SetProjectDoc(ctx, newProjectDoc)

		return newProjectDoc, nil
	}
}

func (k Keeper) SetAccountMap(ctx sdk.Context, projectDid ixo.Did, accountMap types.AccountMap) {
	store := ctx.KVStore(k.storeKey)
	bz, err := json.Marshal(accountMap)
	if err != nil {
		panic(err)
	}
	store.Set(types.GetAccountPrefixKey(projectDid), bz)
}

func (k Keeper) GetAccountMap(ctx sdk.Context, projectDid ixo.Did) types.AccountMap {
	store := ctx.KVStore(k.storeKey)
	key := types.GetAccountPrefixKey(projectDid)

	bz := store.Get(key)
	if bz == nil {
		return make(map[string]sdk.AccAddress)
	} else {
		var accountMap map[string]sdk.AccAddress
		if err := json.Unmarshal(bz, &accountMap); err != nil {
			panic(err)
		}

		return accountMap
	}
}

func (k Keeper) AddAccountToProjectAccounts(ctx sdk.Context, projectDid ixo.Did, accountId string, account auth.Account) {
	accountMap := k.GetAccountMap(ctx, projectDid)
	_, found := accountMap[accountId]
	if found {
		return
	}

	store := ctx.KVStore(k.storeKey)
	key := types.GetAccountPrefixKey(projectDid)
	accountMap[accountId] = account.GetAddress()

	bz, err := json.Marshal(accountMap)
	if err != nil {
		panic(err)
	}

	store.Set(key, bz)
}

func (k Keeper) CreateNewAccount(ctx sdk.Context, projectDid ixo.Did, accountId string) (auth.Account, sdk.Error) {
	key := projectDid + "/" + accountId
	address := sdk.AccAddress(crypto.AddressHash([]byte(key)))

	if k.accountKeeper.GetAccount(ctx, address) != nil {
		return nil, sdk.ErrInvalidAddress("Generate account already exists")
	}

	account := k.accountKeeper.NewAccountWithAddress(ctx, address)
	k.accountKeeper.SetAccount(ctx, account)

	return account, nil
}

func (k Keeper) SetProjectWithdrawalTransactions(ctx sdk.Context, projectDid ixo.Did, txs []types.WithdrawalInfo) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshalBinaryLengthPrefixed(txs)
	store.Set(types.GetWithdrawalPrefixKey(projectDid), bz)
}

func (k Keeper) GetProjectWithdrawalTransactions(ctx sdk.Context, projectDid ixo.Did) ([]types.WithdrawalInfo, sdk.Error) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetWithdrawalPrefixKey(projectDid)

	bz := store.Get(key)
	if bz == nil {
		return []types.WithdrawalInfo{}, did.ErrorInvalidDid(types.DefaultCodeSpace, "ProjectDoc doesn't exist")
	} else {
		var txs []types.WithdrawalInfo
		k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &txs)

		return txs, nil
	}
}

func (k Keeper) AddProjectWithdrawalTransaction(ctx sdk.Context, projectDid ixo.Did, info types.WithdrawalInfo) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetWithdrawalPrefixKey(projectDid)

	txs, _ := k.GetProjectWithdrawalTransactions(ctx, projectDid)
	txs = append(txs, info)

	store.Set(key, k.cdc.MustMarshalBinaryLengthPrefixed(txs))
}
