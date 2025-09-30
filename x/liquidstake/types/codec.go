package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

// RegisterLegacyAminoCodec registers the necessary x/liquidstake interfaces and concrete types
// on the provided LegacyAmino codec. These types are used for Amino JSON serialization.
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgLiquidStake{}, "liquidstake/MsgLiquidStake", nil)
	cdc.RegisterConcrete(&MsgLiquidUnstake{}, "liquidstake/MsgLiquidUnstake", nil)
	cdc.RegisterConcrete(&MsgUpdateParams{}, "liquidstake/MsgUpdateParams", nil)
	cdc.RegisterConcrete(&MsgUpdateWhitelistedValidators{}, "liquidstake/MsgUpdateWhitelistedValidators", nil)
	cdc.RegisterConcrete(&MsgUpdateWeightedRewardsReceivers{}, "liquidstake/MsgUpdateWeightedRewardsReceivers", nil)
	cdc.RegisterConcrete(&MsgSetModulePaused{}, "liquidstake/MsgSetModulePaused", nil)
	cdc.RegisterConcrete(&MsgBurn{}, "liquidstake/MsgBurn", nil)
}

// RegisterInterfaces registers the x/liquidstake interfaces types with the interface registry.
func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgLiquidStake{},
		&MsgLiquidUnstake{},
		&MsgUpdateParams{},
		&MsgUpdateWhitelistedValidators{},
		&MsgUpdateWeightedRewardsReceivers{},
		&MsgSetModulePaused{},
		&MsgBurn{},
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
