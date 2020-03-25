package project

import (
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/ixofoundation/ixo-cosmos/x/project/internal/types"
)

func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(types.MsgCreateProject{}, "project/CreateProject", nil)
	cdc.RegisterConcrete(types.MsgUpdateProjectStatus{}, "project/UpdateProjectStatus", nil)
	cdc.RegisterConcrete(types.MsgCreateAgent{}, "project/CreateAgent", nil)
	cdc.RegisterConcrete(types.MsgUpdateAgent{}, "project/UpdateAgent", nil)
	cdc.RegisterConcrete(types.MsgCreateClaim{}, "project/CreateClaim", nil)
	cdc.RegisterConcrete(types.MsgCreateEvaluation{}, "project/CreateEvaluation", nil)
	cdc.RegisterConcrete(types.MsgWithdrawFunds{}, "project/WithdrawFunds", nil)
}

var moduleCdc = codec.New()

func init() {
	RegisterCodec(moduleCdc)
}
