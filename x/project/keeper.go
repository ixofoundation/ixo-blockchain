package project

import (
	"encoding/hex"
	"encoding/json"
	"errors"

	server "github.com/cosmos/cosmos-sdk/server"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/wire"
	"github.com/cosmos/cosmos-sdk/x/auth"

	"github.com/ixofoundation/ixo-cosmos/x/ixo"
)

// module users must specify coin denomination and reward (constant) per PoW solution
type Config struct {
	accountMapPrefix string
}

// ProjectKeeper manages dids
type Keeper struct {
	key    sdk.StoreKey
	cdc    *wire.Codec
	am     auth.AccountMapper
	config Config
}

// NewKeeper returns a new Keeper
func NewKeeper(cdc *wire.Codec, key sdk.StoreKey, am auth.AccountMapper) Keeper {
	return Keeper{
		key,
		cdc,
		am,
		Config{"ACC-"},
	}
}

// GetDidDoc returns the did_doc at the addr.
func (k Keeper) GetProjectDoc(ctx sdk.Context, did ixo.Did) (ixo.StoredProjectDoc, bool) {
	projectDoc, found := k.GetProjectDoc(ctx, did)
	return projectDoc, found
}

func (k Keeper) SetProjectDoc(ctx sdk.Context, project ixo.StoredProjectDoc) {
	addr := []byte(project.GetProjectDid())
	store := ctx.KVStore(k.key)
	bz := k.encodeProject(project)
	store.Set(addr, bz)
}

// AddDidDoc adds the did_doc at the addr.
func (k Keeper) AddProjectDoc(ctx sdk.Context, newProjectDoc ixo.StoredProjectDoc) (ixo.StoredProjectDoc, sdk.Error) {
	projectDoc, found := k.GetProjectDoc(ctx, newProjectDoc.GetProjectDid())
	if !found {
		k.SetProjectDoc(ctx, newProjectDoc)
		return newProjectDoc, nil
	} else {
		return projectDoc, nil
	}
}

func (k Keeper) GetAccountMap(ctx sdk.Context, projectDid ixo.Did) map[string]interface{} {
	store := ctx.KVStore(k.key)
	key := generateAccountsKey(projectDid)
	bz := store.Get(key)
	if bz == nil {
		return make(map[string]interface{})
	} else {
		didMap := k.decodeAccountMap(bz)
		return didMap
	}
}

func (k Keeper) AddAccountToAccountProjectAccounts(ctx sdk.Context, projectDid ixo.Did, accountDid ixo.Did, account auth.Account) {
	accMap := k.GetAccountMap(ctx, projectDid)
	_, found := accMap[accountDid]
	if found {
		return
	}

	store := ctx.KVStore(k.key)
	key := generateAccountsKey(projectDid)
	accountAddrString := hex.EncodeToString(account.GetAddress())
	accMap[string(accountDid)] = accountAddrString
	bz := k.encodeAccountMap(accMap)
	store.Set(key, bz)
}

func (k Keeper) CreateNewAccount(ctx sdk.Context, projectDid ixo.Did, accountDid ixo.Did) auth.Account {
	// generate secret and address
	addr, _, err := server.GenerateCoinKey()
	if err != nil {
		panic(err)
	}
	//create account with random address
	acc := k.am.NewAccountWithAddress(ctx, addr)
	if k.am.GetAccount(ctx, addr) != nil {
		panic(errors.New("Generate account already exists"))
	}
	// Store the account
	k.am.SetAccount(ctx, acc)

	return acc
}

func (k Keeper) encodeProject(storedProjectDoc ixo.StoredProjectDoc) []byte {
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
