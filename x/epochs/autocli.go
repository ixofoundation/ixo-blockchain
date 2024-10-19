package epochs

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: "ixo.epochs.v1beta1.Query",
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "EpochInfos",
					Use:       "epoch-infos",
					Short:     "Query running epoch infos.",
				},
				{
					RpcMethod:      "CurrentEpoch",
					Use:            "current-epoch [identifier]",
					Short:          "Query current epoch by specified identifier.",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "identifier"}},
				},
			},
		},
	}
}
