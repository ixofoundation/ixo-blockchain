package liquidstake

import (
	"fmt"

	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"
	"github.com/cosmos/cosmos-sdk/version"
)

// AutoCLIOptions defines CLI command surface generated from the v7 multi-pool
// proto definitions. Pool-scoped commands take pool_id as the first
// positional argument; module-scoped commands take none.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: "ixo.liquidstake.v1beta1.Query",
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "ModuleParams",
					Use:       "module-params",
					Short:     "Query the global liquidstake module parameters",
				},
				{
					RpcMethod:      "Pool",
					Use:            "pool [pool-id]",
					Short:          "Query a single liquid staking pool by id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "pool_id"}},
				},
				{
					RpcMethod: "Pools",
					Use:       "pools",
					Short:     "Query every registered liquid staking pool",
				},
				{
					RpcMethod:      "LiquidValidators",
					Use:            "liquid-validators [pool-id]",
					Short:          "Query a pool's liquid validators with state",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "pool_id"}},
				},
				{
					RpcMethod:      "States",
					Use:            "states [pool-id]",
					Short:          "Query a pool's net amount and mint/burn rates",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "pool_id"}},
				},
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service: "ixo.liquidstake.v1beta1.Msg",
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "LiquidStake",
					Use:       "liquid-stake [pool-id] [amount]",
					Short:     "Liquid-stake native tokens into a pool",
					Example:   fmt.Sprintf(`$ %s tx liquidstake liquid-stake zero 1000000uixo --from mykey`, version.AppName),
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "pool_id"},
						{ProtoField: "amount"},
					},
				},
				{
					RpcMethod: "LiquidUnstake",
					Use:       "liquid-unstake [pool-id] [amount]",
					Short:     "Liquid-unstake the LST denom of a pool back into native tokens",
					Example:   fmt.Sprintf(`$ %s tx liquidstake liquid-unstake zero 1000000uzero --from mykey`, version.AppName),
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "pool_id"},
						{ProtoField: "amount"},
					},
				},
				{
					RpcMethod: "SetPoolPaused",
					Use:       "pause-pool [pool-id] [paused]",
					Short:     "Pause or unpause a single pool (governance or pool admin)",
					Example:   fmt.Sprintf(`$ %s tx liquidstake pause-pool zero true --from mykey`, version.AppName),
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "pool_id"},
						{ProtoField: "is_paused"},
					},
				},
				{
					RpcMethod:      "SetModulePaused",
					Use:            "pause-module [paused]",
					Short:          "Toggle the global module-paused kill switch (governance only)",
					Example:        fmt.Sprintf(`$ %s tx liquidstake pause-module true --from mykey`, version.AppName),
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "is_paused"}},
				},
				{
					RpcMethod:      "Burn",
					Use:            "burn [amount]",
					Short:          "Burn native uixo tokens",
					Example:        fmt.Sprintf(`$ %s tx liquidstake burn 1000000uixo --from mykey`, version.AppName),
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "amount"}},
				},
			},
		},
	}
}
