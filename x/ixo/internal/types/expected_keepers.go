package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-blockchain/x/did/exported"
)

// DidKeeper defines the did contract that must be fulfilled throughout the ixo module
type DidKeeper interface {
	GetDidDoc(ctx sdk.Context, did exported.Did) (exported.DidDoc, sdk.Error)
	SetDidDoc(ctx sdk.Context, did exported.DidDoc) (err sdk.Error)
	AddDidDoc(ctx sdk.Context, did exported.DidDoc)
	AddCredentials(ctx sdk.Context, did exported.Did, credential exported.DidCredential) (err sdk.Error)
	GetAllDidDocs(ctx sdk.Context) (didDocs []exported.DidDoc)
	GetAddDids(ctx sdk.Context) (dids []exported.Did)
}
