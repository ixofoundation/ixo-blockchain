package liquidstake

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"
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
		// Tx surface intentionally empty — the LiquidStake/LiquidUnstake/
		// Burn msgs carry gogoproto-generated Coin fields that crash
		// autocli's dynamicpb-based request builder with
		// "cosmos.base.v1beta1.Coin.denom: field descriptor does not
		// belong to this message". The manual cobra cmds in
		// client/cli/tx.go (registered via GetTxCmd) parse Coin via
		// sdk.ParseCoinNormalized and bypass the autocli path.
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:           "ixo.liquidstake.v1beta1.Msg",
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{},
		},
	}
}
