package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// type interface to use within ExecuteOnDidWithRelationships function that also gets used by other modules and thus their expected
// keepers have limited object capabilities, this is minimum required to be able to use the ExecuteOnDidWithRelationships function
type IidKeeper interface {
	GetDidDocument(ctx sdk.Context, key []byte) (IidDocument, bool)
	SetDidDocument(ctx sdk.Context, key []byte, document IidDocument)
}
