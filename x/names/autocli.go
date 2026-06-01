package names

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: "ixo.names.v1beta1.Query",
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod:      "Namespace",
					Use:            "namespace [name]",
					Short:          "Query a namespace by name",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "name"}},
				},
				{
					RpcMethod: "Namespaces",
					Use:       "namespaces",
					Short:     "List all namespaces",
				},
				{
					RpcMethod: "ResolveName",
					Use:       "resolve [namespace] [name]",
					Short:     "Resolve an active name to its NameRecord",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "namespace"},
						{ProtoField: "name"},
					},
				},
				{
					RpcMethod: "GetName",
					Use:       "get [namespace] [normalized_name]",
					Short:     "Get a NameRecord regardless of status",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{
						{ProtoField: "namespace"},
						{ProtoField: "normalized_name"},
					},
				},
				{
					RpcMethod:      "NamesByNamespace",
					Use:            "list-by-namespace [namespace]",
					Short:          "List names registered under a namespace",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "namespace"}},
				},
				{
					RpcMethod:      "NamesByOwner",
					Use:            "list-by-owner [owner_did]",
					Short:          "List names owned by a DID across all namespaces",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "owner_did"}},
				},
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:           "ixo.names.v1beta1.Msg",
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{},
		},
	}
}
