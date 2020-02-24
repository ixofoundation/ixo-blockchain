package types

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	TRUE  = "true"
	FALSE = "false"
)

type BondsMsg interface {
	sdk.Msg
	IsNewDid() bool
}
