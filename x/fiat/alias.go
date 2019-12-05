package fiat

import (
	types2 "github.com/ixofoundation/ixo-cosmos/types"
	"github.com/ixofoundation/ixo-cosmos/x/fiat/internal/keeper"
	"github.com/ixofoundation/ixo-cosmos/x/fiat/internal/types"
)

const (
	StoreKey     = types.StoreKey
	ModuleName   = types.ModuleName
	RouterKey    = types.RouterKey
	QuerierRoute = types.QuerierRoute

	DefaultCodeSpace = types.DefaultCodeSpace
)

var (
	RegisterCodec = types.RegisterCodec
	ModuleCdc     = types.ModuleCdc

	DefaultGenesisState = types.DefaultGenesisState
	ValidateGenesis     = types.ValidateGenesis

	NewQuerier = keeper.NewQuerier
	NewKeeper  = keeper.NewKeeper
)

type (
	GenesisState = types.GenesisState
	Keeper       = keeper.Keeper

	BaseFiatPeg = types2.BaseFiatPeg

	MsgIssueFiats  = types.MsgIssueFiats
	MsgRedeemFiats = types.MsgRedeemFiats
	MsgSendFiats   = types.MsgSendFiats

	IssueFiat = types.IssueFiat
)
