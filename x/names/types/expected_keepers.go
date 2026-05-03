package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	iidtypes "github.com/ixofoundation/ixo-blockchain/v6/x/iid/types"
)

// IidKeeper is the subset of x/iid we depend on for verifying that a tx signer
// controls a given DID.
type IidKeeper interface {
	GetDidDocument(ctx sdk.Context, key []byte) (iidtypes.IidDocument, bool)
}
