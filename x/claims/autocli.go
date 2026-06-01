package claims

import (
	autocliv1 "cosmossdk.io/api/cosmos/autocli/v1"
)

// AutoCLIOptions implements the autocli.HasAutoCLIConfig interface.
func (am AppModule) AutoCLIOptions() *autocliv1.ModuleOptions {
	return &autocliv1.ModuleOptions{
		Query: &autocliv1.ServiceCommandDescriptor{
			Service: "ixo.claims.v1beta1.Query",
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{
				{
					RpcMethod: "CollectionList",
					Use:       "collection-list",
					Short:     "Query for all collections",
				},
				{
					RpcMethod:      "Collection",
					Use:            "collection [id]",
					Short:          "Query for a collection",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod: "ClaimList",
					Use:       "claim-list",
					Short:     "Query for all claims",
				},
				{
					RpcMethod:      "Claim",
					Use:            "claim [id]",
					Short:          "Query for a claim",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "id"}},
				},
				{
					RpcMethod: "DisputeList",
					Use:       "dispute-list",
					Short:     "Query for all disputes",
				},
				{
					RpcMethod:      "Dispute",
					Use:            "dispute [proof-id]",
					Short:          "Query for a dispute",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "proof"}},
				},
				{
					RpcMethod: "IntentList",
					Use:       "intent-list",
					Short:     "Query for all intents",
				},
				{
					RpcMethod:      "Intent",
					Use:            "intent [agent-address] [collection-id] [id]",
					Short:          "Query for an intent",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "agentAddress"}, {ProtoField: "collectionId"}, {ProtoField: "id"}},
				},
				{
					RpcMethod:      "CollectionMemberList",
					Use:            "collection-member-list [collection-id]",
					Short:          "Query for all member budgets on a collection",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "collectionId"}},
				},
				{
					RpcMethod:      "CollectionMember",
					Use:            "collection-member [collection-id] [member-address]",
					Short:          "Query for a member's budget on a collection",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "collectionId"}, {ProtoField: "memberAddress"}},
				},
				{
					RpcMethod:      "DisputeBySubject",
					Use:            "dispute-by-subject [subject-id] [target-role]",
					Short:          "Query a dispute by (subject_id, target_role) via the v7 subject index",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "subjectId"}, {ProtoField: "targetRole"}},
				},
				{
					RpcMethod:      "DisputeListForSubject",
					Use:            "dispute-list-for-subject [subject-id]",
					Short:          "List every dispute (any role) for a subject id",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "subjectId"}},
				},
				{
					RpcMethod:      "AgentDepositBalance",
					Use:            "agent-deposit-balance [collection-id] [agent-address]",
					Short:          "Query a single agent's performance-deposit balance on a collection",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "collectionId"}, {ProtoField: "agentAddress"}},
				},
				{
					RpcMethod:      "AgentDepositBalanceList",
					Use:            "agent-deposit-balance-list [collection-id]",
					Short:          "List every agent performance-deposit balance for a collection",
					PositionalArgs: []*autocliv1.PositionalArgDescriptor{{ProtoField: "collectionId"}},
				},
				{
					RpcMethod: "Params",
					Use:       "params",
					Short:     "Query the current claims parameters",
				},
			},
		},
		Tx: &autocliv1.ServiceCommandDescriptor{
			Service:           "ixo.claims.v1beta1.Msg",
			RpcCommandOptions: []*autocliv1.RpcCommandOptions{},
		},
	}
}
