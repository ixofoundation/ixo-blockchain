package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgCreateProject{}, "ixo-cosmos/MsgCreateProject", nil)
	cdc.RegisterConcrete(MsgCreateAgent{}, "ixo-cosmos/MsgCreateAgent", nil)
	cdc.RegisterConcrete(MsgCreateClaim{}, "ixo-cosmos/CreateClaimMsg", nil)
	cdc.RegisterConcrete(MsgCreateEvaluation{}, "ixo-cosmos/CreateEvaluationMsg", nil)
	cdc.RegisterConcrete(MsgUpdateAgent{}, "ixo-cosmos/UpdateAgentMsg", nil)
	cdc.RegisterConcrete(MsgUpdateProjectStatus{}, "ixo-cosmos/UpdateProjectStatusMsg", nil)
}

var ModuleCdc *codec.Codec

func init() {
	ModuleCdc = codec.New()
	ModuleCdc.Seal()
}
