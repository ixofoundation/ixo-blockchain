package did

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-cosmos/x/ixo"
)

const moduleName = "did"

// DidKeeper manages dids
type DidKeeper struct {
	dm SealedDidMapper
}

// NewKeeper returns a new Keeper
func NewKeeper(dm SealedDidMapper) DidKeeper {
	return DidKeeper{dm: dm}
}

// GetDidDoc returns the did_doc at the addr.
func (dk DidKeeper) GetDidDoc(ctx sdk.Context, did ixo.Did) ixo.DidDoc {
	didDoc := dk.dm.GetDidDoc(ctx, did)
	return didDoc
}

// AddDidDoc adds the did_doc at the addr.
func (dk DidKeeper) AddDidDoc(ctx sdk.Context, newDidDoc ixo.DidDoc) (ixo.DidDoc, sdk.Error) {
	didDoc := dk.dm.GetDidDoc(ctx, newDidDoc.GetDid())
	if didDoc == nil {
		dk.dm.SetDidDoc(ctx, newDidDoc)
		return newDidDoc, nil
	} else {
		return didDoc, nil
	}

}
