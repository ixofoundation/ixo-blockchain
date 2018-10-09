package project

import (
	"github.com/cosmos/cosmos-sdk/wire"
)

// Register concrete types on wire codec
func RegisterWire(cdc *wire.Codec) {
	cdc.RegisterConcrete(CreateProjectMsg{}, "project/CreateProject", nil)
	cdc.RegisterConcrete(UpdateProjectStatusMsg{}, "project/UpdateProjectStatus", nil)
	cdc.RegisterConcrete(CreateAgentMsg{}, "project/CreateAgent", nil)
	cdc.RegisterConcrete(UpdateAgentMsg{}, "project/UpdateAgent", nil)
	cdc.RegisterConcrete(CreateClaimMsg{}, "project/CreateClaim", nil)
	cdc.RegisterConcrete(CreateEvaluationMsg{}, "project/CreateEvaluation", nil)
	cdc.RegisterConcrete(FundProjectMsg{}, "project/FundProject", nil)
	cdc.RegisterConcrete(WithdrawFundsMsg{}, "project/WithdrawFunds", nil)
}

var msgCdc = wire.NewCodec()

func init() {
	RegisterWire(msgCdc)
}
