package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-cosmos/x/ixo"
	"github.com/tendermint/tendermint/crypto"
)

const (
	TRUE  = "true"
	FALSE = "false"
)

type BondsMsg interface {
	sdk.Msg
	IsNewDid() bool
}

func DidToAddr(did ixo.Did) sdk.AccAddress {
	return sdk.AccAddress(crypto.AddressHash([]byte(did)))
}
