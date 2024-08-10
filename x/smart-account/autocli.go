package authenticator

import (
	"fmt"

	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"
	"github.com/cosmos/cosmos-sdk/version"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: "ixo.smartaccount.v1beta1.Query",
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod:      "GetAuthenticator",
					Use:            "get-authenticator [account] [id]",
					Short:          "Query authenticator by account and id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "account"}, {ProtoField: "authenticator_id"}},
				},
				{
					RpcMethod:      "GetAuthenticators",
					Use:            "get-authenticators [account]",
					Short:          "Query authenticators by account",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "account"}},
				},
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Query the current smartaccount parameters",
				},
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service: "ixo.smartaccount.v1beta1.Msg",
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod:      "AddAuthenticator",
					Use:            "add-authenticator [authenticator-type] [data]",
					Short:          "Add an authenticator for an address",
					Example:        fmt.Sprintf(`$ %s tx smartaccount add-authenticator SigVerification <pubkey> --from mykey`, version.AppName),
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "authenticator_type"}, {ProtoField: "data"}},
				},
				{
					RpcMethod:      "RemoveAuthenticator",
					Use:            "remove-authenticator [id]",
					Short:          "Remove an authenticator for an address",
					Example:        fmt.Sprintf(`$ %s tx smartaccount remove-authenticator 1 --from mykey`, version.AppName),
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod:      "SetActiveState",
					Use:            "set-active-state [active]",
					Short:          "Set the smart account active state",
					Example:        fmt.Sprintf(`$ %s tx smartaccount set-active-state true --from mykey`, version.AppName),
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "active"}},
				},
			},
		},
	}
}
