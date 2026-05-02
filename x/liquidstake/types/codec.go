package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
)

// RegisterLegacyAminoCodec registers liquidstake's concrete types under the
// LegacyAmino codec for Amino JSON serialization (used by ledger signing,
// among other things).
func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(&MsgLiquidStake{}, "liquidstake/MsgLiquidStake", nil)
	cdc.RegisterConcrete(&MsgLiquidUnstake{}, "liquidstake/MsgLiquidUnstake", nil)
	cdc.RegisterConcrete(&MsgCreatePool{}, "liquidstake/MsgCreatePool", nil)
	cdc.RegisterConcrete(&MsgUpdateModuleParams{}, "liquidstake/MsgUpdateModuleParams", nil)
	cdc.RegisterConcrete(&MsgUpdatePool{}, "liquidstake/MsgUpdatePool", nil)
	cdc.RegisterConcrete(&MsgUpdateWhitelistedValidators{}, "liquidstake/MsgUpdateWhitelistedValidators", nil)
	cdc.RegisterConcrete(&MsgUpdateWeightedRewardsReceivers{}, "liquidstake/MsgUpdateWeightedRewardsReceivers", nil)
	cdc.RegisterConcrete(&MsgSetPoolPaused{}, "liquidstake/MsgSetPoolPaused", nil)
	cdc.RegisterConcrete(&MsgSetModulePaused{}, "liquidstake/MsgSetModulePaused", nil)
	cdc.RegisterConcrete(&MsgBurn{}, "liquidstake/MsgBurn", nil)
}

// RegisterInterfaces registers liquidstake's sdk.Msg implementations and the
// MsgService descriptor with the interface registry.
func RegisterInterfaces(registry types.InterfaceRegistry) {
	registry.RegisterImplementations(
		(*sdk.Msg)(nil),
		&MsgLiquidStake{},
		&MsgLiquidUnstake{},
		&MsgCreatePool{},
		&MsgUpdateModuleParams{},
		&MsgUpdatePool{},
		&MsgUpdateWhitelistedValidators{},
		&MsgUpdateWeightedRewardsReceivers{},
		&MsgSetPoolPaused{},
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
