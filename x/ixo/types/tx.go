package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-blockchain/x/did/exported"
)

type IxoMsg interface {
	sdk.Msg
	GetSignerDid() exported.Did
	Type() string
}
