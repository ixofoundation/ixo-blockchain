package treasury

import (
	"github.com/ixofoundation/ixo-cosmos/x/treasury/internal/types"
)

const (
	ModuleName = types.ModuleName
	RouterKey  = types.RouterKey
	StoreKey   = types.StoreKey
)

type (
	MsgSend = types.MsgSend
)

var (
	// function aliases
	RegisterCodec = types.RegisterCodec

	// variable aliases
	ModuleCdc = types.ModuleCdc
)
