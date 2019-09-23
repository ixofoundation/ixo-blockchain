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
	CreateProjectMsg       = types.CreateProjectMsg
	UpdateProjectStatusMsg = types.UpdateProjectStatusMsg
	CreateAgentMsg         = types.CreateAgentMsg
	UpdateAgentMsg         = types.UpdateAgentMsg
	CreateClaimMsg         = types.CreateClaimMsg
	CreateEvaluationMsg    = types.CreateEvaluationMsg
	WithdrawFundsMsg       = types.WithdrawFundsMsg
	StoredProjectDoc       = types.StoredProjectDoc
	WithdrawalInfo         = types.WithdrawalInfo
	AccountMap             = types.AccountMap
)

var (
	NewKeeper = keeper.NewKeeper
	ModuleCdc = types.ModuleCdc
)
