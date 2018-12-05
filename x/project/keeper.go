package project

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"errors"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/wire"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/ixofoundation/ixo-cosmos/x/fees"
	"github.com/ixofoundation/ixo-cosmos/x/ixo"
)

// module users must specify coin denomination and reward (constant) per PoW solution
type Config struct {
	accountMapPrefix  string
	withdrawalsPrefix string
}

type StoredProjectDoc interface {
	GetEvaluatorPay() int64
	GetProjectDid() ixo.Did
	GetPubKey() string
	GetStatus() ProjectStatus
	SetStatus(status ProjectStatus)
}

// ProjectKeeper manages dids
type Keeper struct {
	key    sdk.StoreKey
	cdc    *wire.Codec
	am     auth.AccountMapper
	fk     fees.Keeper
	config Config
}

// NewKeeper returns a new Keeper
func NewKeeper(cdc *wire.Codec, key sdk.StoreKey, am auth.AccountMapper, fk fees.Keeper) Keeper {
	return Keeper{
		key,
		cdc,
		am,
		fk,
		Config{
			"ACC-",
			"TX-",
		},
	}
}

// GetDidDoc returns the did_doc at the addr.
func (k Keeper) GetProjectDoc(ctx sdk.Context, did ixo.Did) (StoredProjectDoc, bool) {
	store := ctx.KVStore(k.key)
	bz := store.Get([]byte(did))
	if bz == nil {
		return nil, false
	}
	project := k.decodeProject(bz)
	return project, true
}

func (k Keeper) SetProjectDoc(ctx sdk.Context, project StoredProjectDoc) {
	addr := []byte(project.GetProjectDid())
	store := ctx.KVStore(k.key)
	bz := k.encodeProject(project)
	store.Set(addr, bz)
}

// AddDidDoc adds the did_doc at the addr.
func (k Keeper) AddProjectDoc(ctx sdk.Context, newProjectDoc StoredProjectDoc) (StoredProjectDoc, sdk.Error) {
	projectDoc, found := k.GetProjectDoc(ctx, newProjectDoc.GetProjectDid())
	if !found {
		k.SetProjectDoc(ctx, newProjectDoc)
		return newProjectDoc, nil
	} else {
		return projectDoc, nil
	}
}

// AddDidDoc adds the did_doc at the addr.
func (k Keeper) UpdateProjectDoc(ctx sdk.Context, newProjectDoc StoredProjectDoc) (StoredProjectDoc, bool) {
	projectDoc, found := k.GetProjectDoc(ctx, newProjectDoc.GetProjectDid())
	if !found {
		return StoredProjectDoc(nil), false
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

// GetProjectWithdrawalTransactions returns all the transactions on this project.
func (k Keeper) GetProjectWithdrawalTransactions(ctx sdk.Context, projectDid ixo.Did) []WithdrawalInfo {
	store := ctx.KVStore(k.key)
	key := generateWithdrawalsKey(k, projectDid)
	bz := store.Get(key)
	if bz == nil {
		return []WithdrawalInfo{}
	} else {
		txs := []WithdrawalInfo{}
		err := json.Unmarshal(bz, &txs)
		if err != nil {
			panic(err)
		}
		return txs
	}
}

func (k Keeper) AddProjectWithdrawalTransaction(ctx sdk.Context, projectDid ixo.Did, withdrawalInfo WithdrawalInfo) {
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

	//create account with random address
	acc := k.am.NewAccountWithAddress(ctx, addr)
	if k.am.GetAccount(ctx, addr) != nil {
		panic(errors.New("Generate account already exists"))
	}
	// Store the account
	k.am.SetAccount(ctx, acc)

	return acc
}
func (k Keeper) decodeProject(bz []byte) StoredProjectDoc {

	storedProjectDoc := CreateProjectMsg{}
	err := k.cdc.UnmarshalBinary(bz, &storedProjectDoc)
	if err != nil {
		panic(err)
	}
	return &storedProjectDoc

}
func (k Keeper) encodeProject(storedProjectDoc StoredProjectDoc) []byte {
	bz, err := k.cdc.MarshalBinary(storedProjectDoc)
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
	buffer.WriteString(k.config.accountMapPrefix)
	buffer.WriteString(did)
	return buffer.Bytes()
}

func generateWithdrawalsKey(k Keeper, did ixo.Did) []byte {
	var buffer bytes.Buffer
	buffer.WriteString(k.config.withdrawalsPrefix)
	buffer.WriteString(did)
	return buffer.Bytes()
}
