package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	"github.com/cosmos/cosmos-sdk/x/authz"
	govtypesv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

// RegisterCodec registers the necessary x/token interfaces and concrete types on the provided
// LegacyAmino codec. These types are used for Amino JSON serialization.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgCreateToken{}, "token/CreateToken", nil)
	cdc.RegisterConcrete(&MsgMintToken{}, "token/MintToken", nil)
	cdc.RegisterConcrete(&MsgTransferToken{}, "token/TransferToken", nil)
	cdc.RegisterConcrete(&MsgRetireToken{}, "token/RetireToken", nil)
	cdc.RegisterConcrete(&MsgTransferCredit{}, "token/TransferCredit", nil)
	cdc.RegisterConcrete(&MsgCancelToken{}, "token/CancelToken", nil)

	// gov proposals
	cdc.RegisterConcrete(&SetTokenContractCodes{}, "token/SetTokenContractCodes", nil)
}

// RegisterInterfaces registers interfaces and implementations of the x/token module.
func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgCreateToken{},
		&MsgMintToken{},
		&MsgTransferToken{},
		&MsgRetireToken{},
		&MsgTransferCredit{},
		&MsgCancelToken{},
	)

	registry.RegisterImplementations(
		(*govtypesv1.Content)(nil),
		&SetTokenContractCodes{},
	)

	registry.RegisterImplementations(
		(*authz.Authorization)(nil),
		&MintAuthorization{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewProtoCodec(types.NewInterfaceRegistry())
)

func init() {
	RegisterLegacyAminoCodec(amino)
	cryptocodec.RegisterCrypto(amino)
	amino.Seal()
}
