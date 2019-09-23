package project

import (
	"github.com/cosmos/cosmos-sdk/codec"
	
	"github.com/ixofoundation/ixo-cosmos/x/project/internal/types"
)

func Registercodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(types.CreateProjectMsg{}, "project/CreateProject", nil)
	cdc.RegisterConcrete(types.UpdateProjectStatusMsg{}, "project/UpdateProjectStatus", nil)
	cdc.RegisterConcrete(types.CreateAgentMsg{}, "project/CreateAgent", nil)
	cdc.RegisterConcrete(types.UpdateAgentMsg{}, "project/UpdateAgent", nil)
	cdc.RegisterConcrete(types.CreateClaimMsg{}, "project/CreateClaim", nil)
	cdc.RegisterConcrete(types.CreateEvaluationMsg{}, "project/CreateEvaluation", nil)
	cdc.RegisterConcrete(types.WithdrawFundsMsg{}, "project/WithdrawFunds", nil)
}

var moduleCdc = codec.New()

func init() {
	Registercodec(moduleCdc)
}
