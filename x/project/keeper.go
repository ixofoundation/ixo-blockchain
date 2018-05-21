package project

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-cosmos/x/ixo"
)

const moduleName = "project"

// ProjectKeeper manages dids
type ProjectKeeper struct {
	pm SealedProjectMapper
}

// NewKeeper returns a new Keeper
func NewKeeper(pm SealedProjectMapper) ProjectKeeper {
	return ProjectKeeper{pm: pm}
}

// GetDidDoc returns the did_doc at the addr.
func (pk ProjectKeeper) GetProjectDoc(ctx sdk.Context, did ixo.Did) ixo.ProjectDoc {
	projectDoc := pk.pm.GetProjectDoc(ctx, did)
	return projectDoc
}

// GetAllDids returns all the dids.
func (pk ProjectKeeper) GetAllDids(ctx sdk.Context) []ixo.Did {
	didDoc := pk.pm.GetAllDids(ctx)
	return didDoc
}

// AddDidDoc adds the did_doc at the addr.
func (pk ProjectKeeper) AddProjectDoc(ctx sdk.Context, newProjectDoc ixo.ProjectDoc) (ixo.ProjectDoc, sdk.Error) {
	projectDoc := pk.pm.GetProjectDoc(ctx, newProjectDoc.GetProjectDid())
	if projectDoc == nil {
		pk.pm.SetProjectDoc(ctx, newProjectDoc)
		return newProjectDoc, nil
	} else {
		return projectDoc, nil
	}
}
