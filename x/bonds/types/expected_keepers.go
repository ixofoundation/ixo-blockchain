package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	did "github.com/ixofoundation/ixo-blockchain/lib/legacydid"
)

type IidKeeper interface {
	MustGetDidDoc(ctx sdk.Context, did did.Did) did.DidDoc
	GetDidDoc(ctx sdk.Context, did did.Did) (did.DidDoc, error)
}
