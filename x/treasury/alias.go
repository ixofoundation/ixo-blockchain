package treasury

import (
	"github.com/ixofoundation/ixo-blockchain/x/treasury/internal/keeper"
	"github.com/ixofoundation/ixo-blockchain/x/treasury/internal/types"
)

const (
	ModuleName   = types.ModuleName
	QuerierRoute = types.QuerierRoute
	RouterKey    = types.RouterKey
	StoreKey     = types.StoreKey
)

type (
	Keeper = keeper.Keeper

	MsgSend           = types.MsgSend
	MsgOracleTransfer = types.MsgOracleTransfer
	MsgOracleMint     = types.MsgOracleMint
	MsgOracleBurn     = types.MsgOracleBurn
)

var (
	// function aliases
	NewKeeper     = keeper.NewKeeper
	RegisterCodec = types.RegisterCodec

	// variable aliases
	ModuleCdc = types.ModuleCdc
)
