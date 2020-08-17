package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ixofoundation/ixo-blockchain/x/did/exported"
)

// DidKeeper defines the did contract that must be fulfilled throughout the ixo module
type DidKeeper interface {
	GetDidDoc(ctx sdk.Context, did exported.Did) (exported.DidDoc, error)
	SetDidDoc(ctx sdk.Context, did exported.DidDoc) (err sdkerrors.Error)
	AddDidDoc(ctx sdk.Context, did exported.DidDoc)
	AddCredentials(ctx sdk.Context, did exported.Did, credential exported.DidCredential) (err sdkerrors.Error)
	GetAllDidDocs(ctx sdk.Context) (didDocs []exported.DidDoc)
	GetAddDids(ctx sdk.Context) (dids []exported.Did)
}
