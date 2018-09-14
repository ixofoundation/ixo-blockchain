package project

import (
	"errors"

	server "github.com/cosmos/cosmos-sdk/server"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ixofoundation/ixo-cosmos/x/ixo"
)

const moduleName = "project"

// ProjectKeeper manages dids
type ProjectKeeper struct {
	pm SealedProjectMapper
	am sdk.AccountMapper
}

// NewKeeper returns a new Keeper
func NewKeeper(pm SealedProjectMapper, am sdk.AccountMapper) ProjectKeeper {
	return ProjectKeeper{
		pm: pm,
		am: am,
	}
}

// GetDidDoc returns the did_doc at the addr.
func (pk ProjectKeeper) GetProjectDoc(ctx sdk.Context, did ixo.Did) (ixo.StoredProjectDoc, bool) {
	projectDoc, found := pk.pm.GetProjectDoc(ctx, did)
	return projectDoc, found
}

// AddDidDoc adds the did_doc at the addr.
func (pk ProjectKeeper) AddProjectDoc(ctx sdk.Context, newProjectDoc ixo.StoredProjectDoc) (ixo.StoredProjectDoc, sdk.Error) {
	projectDoc, found := pk.pm.GetProjectDoc(ctx, newProjectDoc.GetProjectDid())
	if !found {
		pk.pm.SetProjectDoc(ctx, newProjectDoc)
		return newProjectDoc, nil
	}
	return projectDoc, nil
}

func (pk ProjectKeeper) GetAccountMap(ctx sdk.Context, projectDid ixo.Did) map[string]interface{} {
	return pk.pm.GetAccountMap(ctx, projectDid)
}

func (pk ProjectKeeper) AddAccountToAccountProjectAccounts(ctx sdk.Context, projectDid ixo.Did, accountDid ixo.Did, account sdk.Account) {
	pk.pm.AddAccountToAccountMap(ctx, projectDid, accountDid, account.GetAddress())
}

func (pk ProjectKeeper) CreateNewAccount(ctx sdk.Context, projectDid ixo.Did, accountDid ixo.Did) sdk.Account {
	// generate secret and address
	addr, _, err := server.GenerateCoinKey()
	if err != nil {
		panic(err)
	}
	//create account with random address
	acc := pk.am.NewAccountWithAddress(ctx, addr)
	if pk.am.GetAccount(ctx, addr) != nil {
		panic(errors.New("Generate account already exists"))
	}
	// Store the account
	pk.am.SetAccount(ctx, acc)

	return acc
}
