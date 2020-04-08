package keeper

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"errors"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"

	"github.com/ixofoundation/ixo-cosmos/x/fees"
	"github.com/ixofoundation/ixo-cosmos/x/ixo"
	"github.com/ixofoundation/ixo-cosmos/x/project/internal/types"
)

type Keeper struct {
	key           sdk.StoreKey
	cdc           *codec.Codec
	accountKeeper auth.AccountKeeper
	feesKeeper    fees.Keeper
	config        types.Config
}

func NewKeeper(cdc *codec.Codec, key sdk.StoreKey, accountKeeper auth.AccountKeeper, feesKeeper fees.Keeper) Keeper {
	return Keeper{
		key,
		cdc,
		accountKeeper,
		feesKeeper,
		types.Config{
			"ACC-",
			"TX-",
		},
	}
}

func (k Keeper) GetProjectDoc(ctx sdk.Context, did ixo.Did) (types.StoredProjectDoc, bool) {
	store := ctx.KVStore(k.key)
	bz := store.Get([]byte(did))
	if bz == nil {
		return nil, false
	}
	project := k.decodeProject(bz)

	return project, true
}

func (k Keeper) SetProjectDoc(ctx sdk.Context, project types.StoredProjectDoc) {
	addr := []byte(project.GetProjectDid())
	store := ctx.KVStore(k.key)
	bz := k.EncodeProject(project)
	store.Set(addr, bz)
}

func (k Keeper) AddProjectDoc(ctx sdk.Context, newProjectDoc types.StoredProjectDoc) (types.StoredProjectDoc,
	sdk.Error) {
	projectDoc, found := k.GetProjectDoc(ctx, newProjectDoc.GetProjectDid())
	if !found {
		k.SetProjectDoc(ctx, newProjectDoc)
		return newProjectDoc, nil
	} else {
		return projectDoc, nil
	}
}

func (k Keeper) UpdateProjectDoc(ctx sdk.Context, newProjectDoc types.StoredProjectDoc) (types.StoredProjectDoc, bool) {
	projectDoc, found := k.GetProjectDoc(ctx, newProjectDoc.GetProjectDid())
	if !found {
		return types.StoredProjectDoc(nil), false
	} else {
		projectDoc.SetStatus(newProjectDoc.GetStatus())
		k.SetProjectDoc(ctx, newProjectDoc)

		return projectDoc, true
	}
}

func (k Keeper) GetAccountMap(ctx sdk.Context, projectDid ixo.Did) map[string]interface{} {
	store := ctx.KVStore(k.key)
	key := generateAccountsKey(k, projectDid)
	bz := store.Get(key)
	if bz == nil {
		return make(map[string]interface{})
	} else {
		didMap := k.decodeAccountMap(bz)

		return didMap
	}
}

func (k Keeper) GetProjectWithdrawalTransactions(ctx sdk.Context, projectDid ixo.Did) []types.WithdrawalInfo {
	store := ctx.KVStore(k.key)
	key := generateWithdrawalsKey(k, projectDid)
	bz := store.Get(key)
	if bz == nil {
		return []types.WithdrawalInfo{}
	} else {
		txs := []types.WithdrawalInfo{}
		err := json.Unmarshal(bz, &txs)
		if err != nil {
			panic(err)
		}

		return txs
	}
}

func (k Keeper) AddProjectWithdrawalTransaction(ctx sdk.Context, projectDid ixo.Did, withdrawalInfo types.WithdrawalInfo) {
	store := ctx.KVStore(k.key)
	key := generateWithdrawalsKey(k, projectDid)

	txs := k.GetProjectWithdrawalTransactions(ctx, projectDid)
	newTxs := append(txs, withdrawalInfo)
	bz, err := json.Marshal(newTxs)
	if err != nil {
		panic(err)
	}

	store.Set(key, bz)
}

func (k Keeper) AddAccountToProjectAccounts(ctx sdk.Context, projectDid ixo.Did, accountId string, account auth.Account) {
	accMap := k.GetAccountMap(ctx, projectDid)
	_, found := accMap[accountId]
	if found {
		return
	}

	store := ctx.KVStore(k.key)
	key := generateAccountsKey(k, projectDid)
	accMap[accountId] = string(account.GetAddress().Bytes())
	bz := k.encodeAccountMap(accMap)
	store.Set(key, bz)
}

func (k Keeper) CreateNewAccount(ctx sdk.Context, projectDid ixo.Did, accountId string) auth.Account {

	src := []byte(projectDid + "/" + accountId)
	hexAddr := hex.EncodeToString(src)
	addr := sdk.AccAddress(hexAddr)

	acc := k.accountKeeper.NewAccountWithAddress(ctx, addr)
	if k.accountKeeper.GetAccount(ctx, addr) != nil {
		panic(errors.New("Generate account already exists"))
	}

	k.accountKeeper.SetAccount(ctx, acc)

	return acc
}
func (k Keeper) decodeProject(bz []byte) types.StoredProjectDoc {

	storedProjectDoc := types.MsgCreateProject{}
	err := k.cdc.UnmarshalBinaryLengthPrefixed(bz, &storedProjectDoc)
	if err != nil {
		panic(err)
	}

	return &storedProjectDoc

}
func (k Keeper) EncodeProject(storedProjectDoc types.StoredProjectDoc) []byte {
	bz, err := k.cdc.MarshalBinaryLengthPrefixed(storedProjectDoc)
	if err != nil {
		panic(err)
	}

	return bz
}

func (k Keeper) encodeAccountMap(accMap map[string]interface{}) []byte {
	json, err := json.Marshal(accMap)
	if err != nil {
		panic(err)
	}

	return []byte(json)
}

func (k Keeper) decodeAccountMap(accMapBytes []byte) map[string]interface{} {
	jsonBytes := []byte(accMapBytes)
	var f interface{}
	err := json.Unmarshal(jsonBytes, &f)
	if err != nil {
		panic(err)
	}

	m := f.(map[string]interface{})
	return m
}

func generateAccountsKey(k Keeper, did ixo.Did) []byte {
	var buffer bytes.Buffer
	buffer.WriteString(k.config.AccountMapPrefix)
	buffer.WriteString(did)
	return buffer.Bytes()
}

func generateWithdrawalsKey(k Keeper, did ixo.Did) []byte {
	var buffer bytes.Buffer
	buffer.WriteString(k.config.WithdrawalsPrefix)
	buffer.WriteString(did)
	return buffer.Bytes()
}
