package project

import (
	"github.com/ixofoundation/ixo-blockchain/x/project/internal/keeper"
	"github.com/ixofoundation/ixo-blockchain/x/project/internal/types"
)

const (
	ModuleName        = types.ModuleName
	DefaultParamspace = types.DefaultParamspace
	QuerierRoute      = types.QuerierRoute
	RouterKey         = types.RouterKey
	StoreKey          = types.StoreKey

	PaidoutStatus = types.PaidoutStatus
	FundedStatus  = types.FundedStatus

	TypeMsgCreateProject = types.TypeMsgCreateProject

	MsgCreateProjectTotalFee       = types.MsgCreateProjectTotalFee
	MsgCreateProjectTransactionFee = types.MsgCreateProjectTransactionFee
)

type (
	Keeper       = keeper.Keeper
	GenesisState = types.GenesisState

	MsgCreateProject       = types.MsgCreateProject
	MsgUpdateProjectStatus = types.MsgUpdateProjectStatus
	MsgCreateAgent         = types.MsgCreateAgent
	MsgUpdateAgent         = types.MsgUpdateAgent
	MsgCreateClaim         = types.MsgCreateClaim
	MsgCreateEvaluation    = types.MsgCreateEvaluation
	MsgWithdrawFunds       = types.MsgWithdrawFunds

	ProjectDoc       = types.ProjectDoc
	StoredProjectDoc = types.StoredProjectDoc

	WithdrawalInfo    = types.WithdrawalInfo
	AccountMap        = types.AccountMap
	GenesisAccountMap = types.GenesisAccountMap
	InternalAccountID = types.InternalAccountID
)

var (
	// function aliases
	NewKeeper      = keeper.NewKeeper
	NewQuerier     = keeper.NewQuerier
	ParamKeyTable  = types.ParamKeyTable
	NewParams      = types.NewParams
	DefaultParams  = types.DefaultParams
	ValidateParams = types.ValidateParams
	RegisterCodec  = types.RegisterCodec

	NewProjectDoc = types.NewProjectDoc

	NewGenesisState     = types.NewGenesisState
	DefaultGenesisState = types.DefaultGenesisState
	ValidateGenesis     = types.ValidateGenesis

	// variable aliases
	ModuleCdc = types.ModuleCdc
)
