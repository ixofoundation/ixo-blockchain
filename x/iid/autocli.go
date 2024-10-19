package iid

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: "ixo.iid.v1beta1.Query",
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "IidDocuments",
					Use:       "iids",
					Short:     "Query for all iids",
				},
				{
					RpcMethod:      "IidDocument",
					Use:            "iid [id]",
					Short:          "Query for an iid",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:           "ixo.iid.v1beta1.Msg",
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{},
		},
	}
}
