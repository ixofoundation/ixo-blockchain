package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ixofoundation/ixo-blockchain/v6/x/claims/types"
)

// InitGenesis initializes the x/claims module's state from a provided genesis state.
func (k Keeper) InitGenesis(ctx sdk.Context, gs types.GenesisState) {
	// Initialise params
	k.SetParams(ctx, &gs.Params)

	// save collections to the store
	for _, c := range gs.Collections {
		k.SetCollection(ctx, c)
	}

	// save claims to the store
	for _, c := range gs.Claims {
		k.SetClaim(ctx, c)
	}

	// save disputes to the store
	for _, d := range gs.Disputes {
		k.SetDispute(ctx, d)
		// Rebuild the (subject_id, target_role) -> proof index and, for
		// disputes still OPEN, the per-agent active-dispute presence index.
		// We don't export those indices separately — they're derived state.
		// Legacy disputes carry target_role=UNSPECIFIED and don't get
		// indexed (they can't block future filings under v7 semantics).
		if d.TargetRole == types.DisputeTargetRole_target_submitter ||
			d.TargetRole == types.DisputeTargetRole_target_evaluator {
			if d.Data != nil && d.Data.Proof != "" {
				k.SetDisputeSubjectIndex(ctx, d.SubjectId, d.TargetRole, d.Data.Proof)
			}
			if d.Status == types.DisputeStatus_dispute_open {
				// Derive the targeted agent from current claim state. Same
				// rule msg_server uses; safe because EVALUATOR-targeted
				// disputes can only exist against terminal evaluations,
				// which are immutable.
				if claim, err := k.GetClaim(ctx, d.SubjectId); err == nil {
					var agent string
					switch d.TargetRole {
					case types.DisputeTargetRole_target_submitter:
						agent = claim.AgentAddress
					case types.DisputeTargetRole_target_evaluator:
						if claim.Evaluation != nil {
							agent = claim.Evaluation.AgentAddress
						}
					}
					if agent != "" {
						k.SetActiveDispute(ctx, claim.CollectionId, agent, d.SubjectId)
					}
				}
			}
		}
	}

	// save intents to the store
	for _, i := range gs.Intents {
		k.SetIntent(ctx, i)
	}

	// save member budgets to the store
	for _, mb := range gs.MemberBudgets {
		k.SetMemberBudget(ctx, mb)
	}

	// save agent deposit balances
	for _, b := range gs.AgentDepositBalances {
		k.SetAgentDepositBalance(ctx, b)
	}
}

// ExportGenesis returns the x/claims module's exported genesis.
func (k Keeper) ExportGenesis(ctx sdk.Context) *types.GenesisState {
	params := k.GetParams(ctx)

	collections := k.GetCollections(ctx)
	claims := k.GetClaims(ctx)
	disputes := k.GetDisputes(ctx)
	intents := k.GetIntents(ctx)
	memberBudgets := k.GetAllMemberBudgets(ctx)
	balances := k.GetAllAgentDepositBalances(ctx)

	return &types.GenesisState{
		Params:               params,
		Collections:          collections,
		Disputes:             disputes,
		Claims:               claims,
		Intents:              intents,
		MemberBudgets:        memberBudgets,
		AgentDepositBalances: balances,
	}
}
