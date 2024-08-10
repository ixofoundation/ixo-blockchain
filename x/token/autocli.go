package token

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: "ixo.token.v1beta1.Query",
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod:      "TokenMetadata",
					Use:            "get-token-metadata [id]",
					Short:          "Query minted token metadata",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod:      "TokenList",
					Use:            "list-tokens [minter]",
					Short:          "List all token docs for a minter",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "minter"}},
				},
				{
					RpcMethod:      "TokenDoc",
					Use:            "show-token-doc [minter] [contract-address]",
					Short:          "Query for a token doc",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "minter"}, {ProtoField: "contract_address"}},
				},
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Query the current token parameters",
				},
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:           "ixo.token.v1beta1.Msg",
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{},
		},
	}
}
