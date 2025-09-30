package liquidstake

import (
	"fmt"

	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"
	"github.com/cosmos/cosmos-sdk/version"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: "ixo.liquidstake.v1beta1.Query",
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "LiquidValidators",
					Use:       "liquid-validators",
					Short:     "Query all liquid validators",
				},
				{
					RpcMethod: "States",
					Use:       "states",
					Short:     "Queries states about net amount, mint rate",
				},
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Query the current liquidstake parameters",
				},
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service: "ixo.liquidstake.v1beta1.Msg",
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod:      "LiquidStake",
					Use:            "liquid-stake [amount]",
					Short:          "Liquid-stake IXO tokens",
					Example:        fmt.Sprintf(`$ %s tx liquidstake liquid-stake 1000000uixo --from mykey`, version.AppName),
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "amount"}},
				},
				{
					RpcMethod:      "LiquidUnstake",
					Use:            "liquid-unstake [amount]",
					Short:          "Liquid-unstake ZERO tokens",
					Example:        fmt.Sprintf(`$ %s tx liquidstake liquid-unstake 1000000uzero --from mykey`, version.AppName),
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "amount"}},
				},
				{
					RpcMethod:      "SetModulePaused",
					Use:            "pause-module [paused]",
					Short:          "Pause or unpause the liquidstake module for an emergency updates",
					Example:        fmt.Sprintf(`$ %s tx liquidstake pause-module true --from mykey`, version.AppName),
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "is_paused"}},
				},
				{
					RpcMethod:      "Burn",
					Use:            "burn [amount]",
					Short:          "Burn IXO tokens",
					Example:        fmt.Sprintf(`$ %s tx liquidstake burn 1000000uixo --from mykey`, version.AppName),
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "amount"}},
				},
			},
		},
	}
}
