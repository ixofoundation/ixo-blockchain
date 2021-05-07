package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(MsgCreateProject{}, "project/CreateProject", nil)
	cdc.RegisterConcrete(MsgUpdateProjectStatus{}, "project/UpdateProjectStatus", nil)
	cdc.RegisterConcrete(MsgCreateAgent{}, "project/CreateAgent", nil)
	cdc.RegisterConcrete(MsgUpdateAgent{}, "project/UpdateAgent", nil)
	cdc.RegisterConcrete(MsgCreateClaim{}, "project/CreateClaim", nil)
	cdc.RegisterConcrete(MsgCreateEvaluation{}, "project/CreateEvaluation", nil)
	cdc.RegisterConcrete(MsgWithdrawFunds{}, "project/WithdrawFunds", nil)

	cdc.RegisterConcrete(ProjectDoc{}, "project/ProjectDoc", nil)
	cdc.RegisterConcrete(AccountMap{}, "project/AccountMap", nil)
	cdc.RegisterConcrete(WithdrawalInfoDoc{}, "project/WithdrawalInfo", nil)
}

var (
	amino = codec.NewLegacyAmino()

	// ModuleCdc references the global x/gov module codec. Note, the codec should
	// ONLY be used in certain instances of tests and for JSON encoding as Amino is
	// still used for that purpose.
	//
	// The actual codec used for serialization should be provided to x/gov and
	// defined at the application level.
	ModuleCdc = codec.NewAminoCodec(amino)
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
}

