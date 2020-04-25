package treasury

import (
	"github.com/ixofoundation/ixo-cosmos/x/treasury/internal/keeper"
	"github.com/ixofoundation/ixo-cosmos/x/treasury/internal/types"
)

const (
	ModuleName = types.ModuleName
	RouterKey  = types.RouterKey
	StoreKey   = types.StoreKey
)

type (
	Keeper = keeper.Keeper

	MsgSend           = types.MsgSend
	MsgSendOnBehalfOf = types.MsgSendOnBehalfOf
	MsgMint           = types.MsgMint
	MsgBurn           = types.MsgBurn
)

var (
	// function aliases
	NewKeeper     = keeper.NewKeeper
	RegisterCodec = types.RegisterCodec

	// variable aliases
	ModuleCdc = types.ModuleCdc
)
