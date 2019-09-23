package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(CreateProjectMsg{}, "ixo-cosmos/MsgCreateProject", nil)
	cdc.RegisterConcrete(CreateAgentMsg{}, "ixo-cosmos/MsgCreateAgent", nil)
	cdc.RegisterConcrete(CreateClaimMsg{}, "ixo-cosmos/CreateClaimMsg", nil)
	cdc.RegisterConcrete(CreateEvaluationMsg{}, "ixo-cosmos/CreateEvaluationMsg", nil)
	cdc.RegisterConcrete(UpdateAgentMsg{}, "ixo-cosmos/UpdateAgentMsg", nil)
	cdc.RegisterConcrete(UpdateProjectStatusMsg{}, "ixo-cosmos/UpdateProjectStatusMsg", nil)
}

var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	ModuleCdc.Seal()
}
