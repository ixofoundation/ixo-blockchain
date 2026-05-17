package keeper

import (
	errorsmod "cosmossdk.io/errors"
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ixofoundation/ixo-blockchain/v6/x/claims/types"
	entitytypes "github.com/ixofoundation/ixo-blockchain/v6/x/entity/types"
	iidtypes "github.com/ixofoundation/ixo-blockchain/v6/x/iid/types"
)

// --------------------------
// DISPUTE SUBJECT INDEX
// --------------------------
//
// Maps (subject_id, target_role) -> dispute proof CID. The proof CID is
// the key under which the primary Dispute record is stored (see
// claims.go SetDispute / GetDispute), so we can fetch the dispute itself
// by chaining: subject index -> proof -> primary record.
//
// Rules:
//   - one entry per (subject_id, target_role); writes overwrite the prior
//   - on adjudication the index entry is updated to point at the latest
//     dispute's proof so the recorded status governs "can someone dispute
//     this same (subject, role) again?"
//   - DISMISSED disputes can be superseded by a new filing; AWARDED ones
//     permanently block further disputes against the same role.

// SetDisputeSubjectIndex writes the (subject_id, target_role) -> proof
// pointer. Caller is responsible for ensuring the proof points to a
// dispute that currently lives in the primary store.
func (k Keeper) SetDisputeSubjectIndex(ctx sdk.Context, subjectId string, targetRole types.DisputeTargetRole, proof string) {
	key := types.DisputeSubjectIndexKeyCreate(subjectId, targetRole)
	store := ctx.KVStore(k.storeKey)
	store.Set(append(types.DisputeSubjectIndexKey, key...), []byte(proof))
}

// GetDisputeProofForSubject returns the proof CID of the latest dispute for
// the given (subject_id, target_role), if any.
func (k Keeper) GetDisputeProofForSubject(ctx sdk.Context, subjectId string, targetRole types.DisputeTargetRole) (string, bool) {
	key := types.DisputeSubjectIndexKeyCreate(subjectId, targetRole)
	store := ctx.KVStore(k.storeKey)
	val := store.Get(append(types.DisputeSubjectIndexKey, key...))
	if len(val) == 0 {
		return "", false
	}
	return string(val), true
}

// GetDisputeBySubject fetches the dispute record for (subject_id,
// target_role) by chaining the subject index to the primary store.
func (k Keeper) GetDisputeBySubject(ctx sdk.Context, subjectId string, targetRole types.DisputeTargetRole) (types.Dispute, error) {
	proof, found := k.GetDisputeProofForSubject(ctx, subjectId, targetRole)
	if !found {
		return types.Dispute{}, errorsmod.Wrapf(types.ErrDisputeNotFoundForSubjectRole, "subject %s role %s", subjectId, targetRole)
	}
	return k.GetDispute(ctx, proof)
}

// CanFileNewDisputeForSubject decides whether a new dispute can be filed
// against (subject_id, target_role). Returns nil if allowed, otherwise an
// error explaining why not. Rules:
//   - no prior dispute -> allowed
//   - latest dispute OPEN -> blocked (one at a time)
//   - latest dispute AWARDED -> blocked permanently
//   - latest dispute DISMISSED -> allowed (overwrites with new filing)
func (k Keeper) CanFileNewDisputeForSubject(ctx sdk.Context, subjectId string, targetRole types.DisputeTargetRole) error {
	existing, err := k.GetDisputeBySubject(ctx, subjectId, targetRole)
	if err != nil {
		// not found -> allowed
		return nil
	}
	switch existing.Status {
	case types.DisputeStatus_dispute_open:
		return errorsmod.Wrapf(types.ErrDisputeAlreadyOpenForSubjectRole,
			"subject %s role %s has an open dispute (proof %s)",
			subjectId, targetRole, existing.Data.Proof)
	case types.DisputeStatus_dispute_awarded:
		return errorsmod.Wrapf(types.ErrDisputeAlreadyAwardedForSubjectRole,
			"subject %s role %s was previously AWARDED; further disputes blocked",
			subjectId, targetRole)
	case types.DisputeStatus_dispute_dismissed:
		return nil
	default:
		// Defensive: should not happen on records written by this version.
		// Legacy migrated disputes use DISMISSED so they don't block; any
		// other unexpected value is treated as blocking to be safe.
		return errorsmod.Wrapf(types.ErrDisputeAlreadyOpenForSubjectRole,
			"subject %s role %s has a dispute in unknown status %d",
			subjectId, targetRole, int32(existing.Status))
	}
}

// --------------------------
// ACTIVE-DISPUTE INDEX (per-agent presence index)
// --------------------------
//
// Key: collectionId + "/" + agentAddress + "/" + subjectId
// Value: empty (presence-only)
//
// Used by SubmitClaim / EvaluateClaim / WithdrawPerformanceDeposit to gate
// the actor: "any OPEN dispute against this agent on this collection?" is
// a prefix scan with limit 1, O(1) gas.

// SetActiveDispute marks (collection, agent, subject) as having an OPEN
// dispute. Called on MsgDisputeClaim.
func (k Keeper) SetActiveDispute(ctx sdk.Context, collectionId, agentAddress, subjectId string) {
	key := types.ActiveDisputeKeyCreate(collectionId, agentAddress, subjectId)
	store := ctx.KVStore(k.storeKey)
	store.Set(append(types.ActiveDisputeKey, key...), []byte{1})
}

// RemoveActiveDispute clears the presence entry on adjudication.
func (k Keeper) RemoveActiveDispute(ctx sdk.Context, collectionId, agentAddress, subjectId string) {
	key := types.ActiveDisputeKeyCreate(collectionId, agentAddress, subjectId)
	store := ctx.KVStore(k.storeKey)
	store.Delete(append(types.ActiveDisputeKey, key...))
}

// HasActiveDisputeAgainstAgent reports whether the agent has any OPEN
// dispute targeting them on this collection. Prefix scan, exits on first
// entry.
func (k Keeper) HasActiveDisputeAgainstAgent(ctx sdk.Context, collectionId, agentAddress string) bool {
	prefix := types.ActiveDisputeAgentPrefix(collectionId, agentAddress)
	iter := k.GetAll(ctx, append(types.ActiveDisputeKey, prefix...))
	defer iter.Close()
	return iter.Valid()
}

// GetActiveDisputeSubjectsForAgent returns the subject ids of every OPEN
// dispute targeting an agent on a collection. Mostly useful for indexers
// / debugging; the gate check above is the hot path.
//
// Note: storetypes.KVStorePrefixIterator (used inside k.GetAll) returns
// FULL keys including the table prefix byte — it does not strip. So the
// keys we see here look like `\x07{collectionId}/{agentAddress}/{subjectId}`
// and we slice off the known-length composite prefix to extract the
// subjectId.
func (k Keeper) GetActiveDisputeSubjectsForAgent(ctx sdk.Context, collectionId, agentAddress string) []string {
	fullPrefix := append(types.ActiveDisputeKey, types.ActiveDisputeAgentPrefix(collectionId, agentAddress)...)
	iter := k.GetAll(ctx, fullPrefix)
	defer iter.Close()
	prefixLen := len(fullPrefix)
	var subjects []string
	for ; iter.Valid(); iter.Next() {
		key := iter.Key()
		if len(key) > prefixLen {
			subjects = append(subjects, string(key[prefixLen:]))
		}
	}
	return subjects
}

// --------------------------
// ADJUDICATOR AUTHORIZATION
// --------------------------
//
// Authorization for MsgAdjudicateDispute follows the same pattern as every
// other DID-gated message in this module: MsgAdjudicateDispute implements
// IidTxMsg via GetIidController() returning AdjudicatorDid, and the IID
// ante (VerifyIidControllersAgainstSignature) checks that the signer's
// address is registered on the DID document under the Authentication or
// AssertionMethod verification relationships.
//
// The keeper re-checks the same relationship as defense-in-depth (so
// in-keeper unit tests that bypass the ante are still safe), and
// additionally decides payout routing: if the DID resolves to an entity in
// the entity module, the 20% adjudicator share is paid into the entity's
// EntityAdjudicatorRevenueAccountName account (auto-created on first
// payout, mirroring the oracle-revenue pattern); otherwise it goes
// directly to the signer.

// AdjudicatorAuthorization is the result of a successful authorization
// check; carries the validated signer and the payout-routing flag.
type AdjudicatorAuthorization struct {
	// AdjudicatorDid is the validated DID (a member of the collection's
	// adjudicators whitelist).
	AdjudicatorDid string
	// AdjudicatorAddress is the validated signer — registered as a
	// verification method on AdjudicatorDid's DID document.
	AdjudicatorAddress sdk.AccAddress
	// IsEntity is true iff AdjudicatorDid resolves to an entity in the
	// entity module. Drives payout routing only (entity revenue account
	// vs direct payout). Independent of how the signer was authorised.
	IsEntity bool
}

// AuthorizeAdjudicator verifies that signer is a registered key on
// adjudicatorDid (defense-in-depth — the IID ante already runs the same
// check at tx time) and resolves payout routing.
func (k Keeper) AuthorizeAdjudicator(ctx sdk.Context, adjudicatorDid string, signer sdk.AccAddress) (AdjudicatorAuthorization, error) {
	iidDoc, ok := k.IidKeeper.GetDidDocument(ctx, []byte(adjudicatorDid))
	if !ok {
		return AdjudicatorAuthorization{}, errorsmod.Wrapf(types.ErrAdjudicatorNotAuthorized,
			"did document %s not found", adjudicatorDid)
	}
	if !iidDoc.HasRelationship(iidtypes.NewBlockchainAccountID(signer.String()),
		iidtypes.Authentication,
		iidtypes.AssertionMethod,
	) {
		return AdjudicatorAuthorization{}, errorsmod.Wrapf(types.ErrAdjudicatorNotAuthorized,
			"address %s not authorized for did %s", signer.String(), adjudicatorDid)
	}

	// Payout routing: does the DID resolve to an entity?
	_, _, entityErr := k.EntityKeeper.ResolveEntity(ctx, adjudicatorDid)
	return AdjudicatorAuthorization{
		AdjudicatorDid:     adjudicatorDid,
		AdjudicatorAddress: signer,
		IsEntity:           entityErr == nil,
	}, nil
}

// AdjudicatorPayoutAddress returns the address where the adjudicator share
// should be paid. For entity-resolved adjudicator DIDs, looks up (or creates)
// the EntityAdjudicatorRevenueAccountName account on the entity, following
// the same pattern as the oracle-revenue account in processPayment().
func (k Keeper) AdjudicatorPayoutAddress(ctx sdk.Context, auth AdjudicatorAuthorization) (sdk.AccAddress, error) {
	if !auth.IsEntity {
		return auth.AdjudicatorAddress, nil
	}

	_, entity, err := k.EntityKeeper.ResolveEntity(ctx, auth.AdjudicatorDid)
	if err != nil {
		return nil, errorsmod.Wrapf(types.ErrInternalError,
			"failed to resolve adjudicator entity %s: %s", auth.AdjudicatorDid, err)
	}

	accountStr, err := entity.GetEntityAccountByName(types.EntityAdjudicatorRevenueAccountName)
	if err == nil {
		return sdk.AccAddressFromBech32(accountStr)
	}

	// Not yet created — create it now, mirroring the oracle-revenue flow.
	address, err := k.EntityKeeper.CreateNewAccount(ctx, entity.Id, types.EntityAdjudicatorRevenueAccountName)
	if err != nil {
		return nil, err
	}
	entity.Accounts = append(entity.Accounts, &entitytypes.EntityAccount{
		Name:    types.EntityAdjudicatorRevenueAccountName,
		Address: address.String(),
	})
	entitytypes.UpdateEntityMetadata(entity.Metadata, ctx.TxBytes(), ctx.BlockTime())
	k.EntityKeeper.SetEntity(ctx, []byte(entity.Id), entity)

	// Emit the same shape of events the oracle path uses so indexers
	// recognise the entity / account being created.
	if err := ctx.EventManager().EmitTypedEvents(
		&entitytypes.EntityUpdatedEvent{Entity: &entity, Signer: ""},
		&entitytypes.EntityAccountCreatedEvent{
			Id:             entity.Id,
			Signer:         "",
			AccountName:    types.EntityAdjudicatorRevenueAccountName,
			AccountAddress: address.String(),
		},
	); err != nil {
		return nil, err
	}

	return address, nil
}

// --------------------------
// COLLECTION HELPERS
// --------------------------

// LookupAdjudicator returns the AdjudicationDid entry for the given DID
// from a collection's adjudicators whitelist. Used by MsgAdjudicateDispute
// to (a) verify the caller's DID is whitelisted and (b) read that
// adjudicator's reward_percentage for the penalty split.
func LookupAdjudicator(c types.Collection, did string) (*types.AdjudicationDid, bool) {
	for _, a := range c.Adjudicators {
		if a != nil && a.Did == did {
			return a, true
		}
	}
	return nil, false
}

// DepositRequiredForRole returns the configured deposit-required for the
// given role on a collection. Empty Coins if no requirement is set.
func DepositRequiredForRole(c types.Collection, role types.DisputeTargetRole) sdk.Coins {
	switch role {
	case types.DisputeTargetRole_target_submitter:
		return c.ServiceAgentDepositRequired
	case types.DisputeTargetRole_target_evaluator:
		return c.EvaluatorDepositRequired
	default:
		return sdk.NewCoins()
	}
}

// PenaltyPotForAwarded computes the actual penalty amount for an AWARDED
// outcome: min(intended, loser balance per denom). The pot is bounded by
// what is actually slashable — a balance shorter than intended results in
// a partial pay-out.
func (k Keeper) PenaltyPotForAwarded(
	ctx sdk.Context,
	collectionId string,
	loserAddress string,
	intended sdk.Coins,
) sdk.Coins {
	balance, err := k.GetAgentDepositBalance(ctx, collectionId, loserAddress)
	if err != nil {
		return sdk.NewCoins()
	}
	return minCoinsPerDenom(intended, balance.Amount)
}

// SplitPenalty divides `pot` according to the adjudicator's reward_percentage.
// Returns (winnerAmount, adjudicatorAmount). Rounds adjudicator share
// down so winner gets the remainder — preserves the invariant that
// winner + adjudicator == pot exactly.
//
// If pot is zero, returns (zero, zero).
// percentage is expected in [0, 100]; the caller validates upstream.
func SplitPenalty(pot sdk.Coins, percentage math.LegacyDec) (winner, adjudicator sdk.Coins) {
	if pot.IsZero() {
		return sdk.NewCoins(), sdk.NewCoins()
	}
	// Guard against an unset percentage; treat as zero (winner takes all).
	pct := percentage
	if pct.IsNil() {
		pct = math.LegacyZeroDec()
	}

	adjudicator = sdk.NewCoins()
	for _, c := range pot {
		share := math.LegacyNewDecFromInt(c.Amount).Mul(pct).Quo(types.OneHundred).TruncateInt()
		if share.IsPositive() {
			adjudicator = adjudicator.Add(sdk.NewCoin(c.Denom, share))
		}
	}
	winner, _ = pot.SafeSub(adjudicator...)
	// Defensive: drop any residual zero/negative coins (SafeSub shouldn't
	// produce negatives given adjudicator ≤ pot per-denom, but be safe).
	winner = winner.Sort()
	return winner, adjudicator
}
