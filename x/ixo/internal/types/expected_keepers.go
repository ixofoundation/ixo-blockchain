package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-blockchain/x/did/exported"
)

// DidKeeper defines the did contract that must be fulfilled throughout the ixo module
type DidKeeper interface {
	GetDidDoc(ctx sdk.Context, did exported.Did) (exported.DidDoc, error)
	SetDidDoc(ctx sdk.Context, did exported.DidDoc) (err error)
	AddDidDoc(ctx sdk.Context, did exported.DidDoc)
	AddCredentials(ctx sdk.Context, did exported.Did, credential exported.DidCredential) (err error)
	GetAllDidDocs(ctx sdk.Context) (didDocs []exported.DidDoc)
	GetAllDids(ctx sdk.Context) (dids []exported.Did)
}
