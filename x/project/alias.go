package project

import (
	"github.com/ixofoundation/ixo-cosmos/x/project/internal/keeper"
	"github.com/ixofoundation/ixo-cosmos/x/project/internal/types"
)

const (
	ModuleName   = types.ModuleName
	QuerierRoute = types.QuerierRoute
	RouterKey    = types.RouterKey
	StoreKey     = types.StoreKey

	DefaultCodeSpace = types.DefaultCodeSpace
	PaidoutStatus    = types.PaidoutStatus
	FundedStatus     = types.FundedStatus
)

type (
	Keeper                 = keeper.Keeper
	MsgCreateProject       = types.MsgCreateProject
	MsgUpdateProjectStatus = types.MsgUpdateProjectStatus
	MsgCreateAgent         = types.MsgCreateAgent
	MsgUpdateAgent         = types.MsgUpdateAgent
	MsgCreateClaim         = types.MsgCreateClaim
	MsgCreateEvaluation    = types.MsgCreateEvaluation
	MsgWithdrawFunds       = types.MsgWithdrawFunds
	StoredProjectDoc       = types.StoredProjectDoc
	WithdrawalInfo         = types.WithdrawalInfo
	AccountMap             = types.AccountMap
)

var (
	// function aliases
	NewKeeper     = keeper.NewKeeper
	NewQuerier    = keeper.NewQuerier
	RegisterCodec = types.RegisterCodec

	// variable aliases
	ModuleCdc = types.ModuleCdc
)
