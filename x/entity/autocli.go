package entity

// TODO: add autocli in all modules

// import (
// 	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"
// )

// // AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
// func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
// 	return &autocliv1.ModuleOptions{
// 		Query: &autocliv1.ServiceCommandDescriptor{
// 			Service: entityv1beta1.Query_ServiceDesc.ServiceName,
// 			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
// 				{
// 					RpcMethod: "Params",
// 					Use:       "params",
// 					Short:     "Query the current entity parameters",
// 				},
// 				{
// 					RpcMethod:      "Entity",
// 					Use:            "entity [id]",
// 					Short:          "Query for an entity",
// 					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
// 				},
// 			},
// 		},
// 	}
// }
