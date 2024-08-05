package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	"github.com/cosmos/cosmos-sdk/x/authz"
)

// RegisterCodec registers the necessary x/claims interfaces and concrete types on the provided
// LegacyAmino codec. These types are used for Amino JSON serialization.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateCollection{}, "claims/CreateCollection", nil)
	cdc.RegisterConcrete(&MsgSubmitClaim{}, "claims/SubmitClaim", nil)
	cdc.RegisterConcrete(&MsgEvaluateClaim{}, "claims/EvaluateClaim", nil)
	cdc.RegisterConcrete(&MsgDisputeClaim{}, "claims/DisputeClaim", nil)
	cdc.RegisterConcrete(&MsgWithdrawPayment{}, "claims/WithdrawPayment", nil)
	cdc.RegisterConcrete(&MsgUpdateCollectionState{}, "claims/UpdateCollectionState", nil)
	cdc.RegisterConcrete(&MsgUpdateCollectionDates{}, "claims/UpdateCollectionDates", nil)
	cdc.RegisterConcrete(&MsgUpdateCollectionPayments{}, "claims/UpdateCollectionPayments", nil)
}

// RegisterInterfaces registers interfaces and implementations of the x/claims module.
func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateCollection{},
		&MsgSubmitClaim{},
		&MsgEvaluateClaim{},
		&MsgDisputeClaim{},
		&MsgWithdrawPayment{},
		&MsgUpdateCollectionState{},
		&MsgUpdateCollectionDates{},
		&MsgUpdateCollectionPayments{},
	)

	// TODO: make sure authz shouldn't also be registered as concrete type above?
	registry.RegisterImplementations(
		(*authz.Authorization)(nil),
		&SubmitClaimAuthorization{},
		&EvaluateClaimAuthorization{},
		&WithdrawPaymentAuthorization{},
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
