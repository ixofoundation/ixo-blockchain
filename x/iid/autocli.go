package iid

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
//
// The Query surface is left intentionally empty — autocli's CLI
// rendering uses `aminojson.NewEncoder` over `dynamicpb.Message`,
// which returns empty `{}` for gogoproto-generated nested message
// types (Service, VerificationMethod, IidMetadata). The manual
// GetQueryCmd in client/cli/query.go is what cosmos-sdk wires into
// `ixod query iid ...`, and it uses the gogo jsonpb path which
// renders nested types correctly.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:           "ixo.iid.v1beta1.Msg",
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{},
		},
	}
}
