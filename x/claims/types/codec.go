package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	"github.com/cosmos/cosmos-sdk/x/authz"
)

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateCollection{}, "claims/CreateCollection", nil)
	cdc.RegisterConcrete(&MsgSubmitClaim{}, "claims/SubmitClaim", nil)
	cdc.RegisterConcrete(&MsgEvaluateClaim{}, "claims/EvaluateClaim", nil)
	cdc.RegisterConcrete(&MsgDisputeClaim{}, "claims/DisputeClaim", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateCollection{},
		&MsgSubmitClaim{},
		&MsgEvaluateClaim{},
		&MsgDisputeClaim{},
	)

	registry.RegisterImplementations(
		(*authz.Authorization)(nil),
		&SubmitClaimAuthorization{},
		&EvaluateClaimAuthorization{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(cdctypes.NewInterfaceRegistry())
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}
