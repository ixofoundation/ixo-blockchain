package types

import (
	"encoding/hex"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-cosmos/x/ixo"
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
	return sdk.AccAddress(hex.EncodeToString([]byte(did)))
}
