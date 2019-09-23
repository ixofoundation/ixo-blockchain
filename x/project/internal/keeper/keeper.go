package keeper

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	
	didTypes "github.com/ixofoundation/ixo-cosmos/x/did"
	"github.com/ixofoundation/ixo-cosmos/x/ixo"
	"github.com/ixofoundation/ixo-cosmos/x/project/internal/types"
)

type Keeper struct {
	cdc           *codec.Codec
	storeKey      sdk.StoreKey
	accountKeeper types.AccountKeeper
	feeKeeper     types.FeeKeeper
}

func NewKeeper(cdc *codec.Codec, key sdk.StoreKey, accountKeeper types.AccountKeeper, feeKeeper types.FeeKeeper) Keeper {
	return Keeper{
		cdc:           cdc,
		storeKey:      key,
		accountKeeper: accountKeeper,
		feeKeeper:     feeKeeper,
	}
}

func (k Keeper) GetProjectDoc(ctx sdk.Context, projectDid ixo.Did) (types.StoredProjectDoc, sdk.Error) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetProjectPrefixKey(projectDid)
	
	bz := store.Get(key)
	if bz == nil {
		return nil, didTypes.ErrorInvalidDid(types.DefaultCodeSpace, "Invalid ProjectDid Address")
	}
	
	var projectDoc types.CreateProjectMsg
	k.cdc.MustUnmarshalBinaryLengthPrefixed(bz, &projectDoc)
	
	return &projectDoc, nil
}

func (k Keeper) SetProjectDoc(ctx sdk.Context, projectDoc types.StoredProjectDoc) sdk.Error {
	existedDoc, err := k.GetProjectDoc(ctx, projectDoc.GetProjectDid())
	if existedDoc != nil {
		return didTypes.ErrorInvalidDid(types.DefaultCodeSpace, fmt.Sprintf("Project already exists %s", err))
	}
	
	k.AddProjectDoc(ctx, projectDoc)
	
	return nil
}

func (k Keeper) AddProjectDoc(ctx sdk.Context, projectDoc types.StoredProjectDoc) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetProjectPrefixKey(projectDoc.GetProjectDid())
	store.Set(key, k.cdc.MustMarshalBinaryLengthPrefixed(projectDoc))
}

func (k Keeper) UpdateProjectDoc(ctx sdk.Context, newProjectDoc types.StoredProjectDoc) (types.StoredProjectDoc, sdk.Error) {
	existedDoc, _ := k.GetProjectDoc(ctx, newProjectDoc.GetProjectDid())
	if existedDoc == nil {
		
		return nil, didTypes.ErrorInvalidDid(types.DefaultCodeSpace, "ProjectDoc details are not exist")
	} else {
		
		existedDoc.SetStatus(newProjectDoc.GetStatus())
		k.AddProjectDoc(ctx, newProjectDoc)
		
		return newProjectDoc, nil
	}
}

func (k Keeper) GetAccountMap(ctx sdk.Context, projectDid ixo.Did) map[string]interface{} {
	store := ctx.KVStore(k.storeKey)
	key := types.GetAccountPrefixKey(projectDid)
	
	bz := store.Get(key)
	if bz == nil {
		return make(map[string]interface{})
	} else {
		var accountMap map[string]interface{}
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
	accountMap[accountId] = string(account.GetAddress().Bytes())
	
	bz, err := json.Marshal(accountMap)
	if err != nil {
		panic(err)
	}
	
	store.Set(key, bz)
}

func (k Keeper) CreateNewAccount(ctx sdk.Context, projectDid ixo.Did, accountId string) (auth.Account, sdk.Error) {
	src := []byte(projectDid + "/" + accountId)
	hexAddress := hex.EncodeToString(src)
	address := sdk.AccAddress(hexAddress)
	
	if k.accountKeeper.GetAccount(ctx, address) != nil {
		return nil, sdk.ErrInvalidAddress("Generate account already exists")
	}
	
	account := k.accountKeeper.NewAccountWithAddress(ctx, address)
	k.accountKeeper.SetAccount(ctx, account)
	
	return account, nil
}

func (k Keeper) GetProjectWithdrawalTransactions(ctx sdk.Context, projectDid ixo.Did) ([]types.WithdrawalInfo, sdk.Error) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetWithdrawalPrefixKey(projectDid)
	
	bz := store.Get(key)
	if bz == nil {
		return []types.WithdrawalInfo{}, didTypes.ErrorInvalidDid(types.DefaultCodeSpace, "ProjectDoc doesn't exist")
	} else {
		txs := []types.WithdrawalInfo{}
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
