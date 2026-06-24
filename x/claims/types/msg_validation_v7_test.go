package types_test

import (
	"testing"
	"time"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	iidtypes "github.com/ixofoundation/ixo-blockchain/v8/x/iid/types"

	"github.com/ixofoundation/ixo-blockchain/v8/x/claims/types"
)

// validBech32 returns a deterministic, well-formed bech32 ixo address
// for use in negative-validation tests. Avoids picking a "real" address
// that might accidentally be used elsewhere in the test suite.
func validBech32(label string) string {
	return sdk.AccAddress("v7-validation-" + label).String()
}

func validAdjudicator(did string, pct int64) *types.AdjudicationDid {
	return &types.AdjudicationDid{Did: did, RewardPercentage: math.LegacyNewDec(pct)}
}

// --------------------------------------------------------------------------
// MsgUpdateCollectionDisputeConfig
// --------------------------------------------------------------------------

func TestMsgUpdateCollectionDisputeConfig_ValidateBasic(t *testing.T) {
	admin := validBech32("admin1")

	// Happy path: empty config validates (the call clears the v7 surface
	// on a collection that has no open disputes).
	require.NoError(t, (&types.MsgUpdateCollectionDisputeConfig{
		CollectionId: "c1",
		AdminAddress: admin,
	}).ValidateBasic())

	// Invalid bech32 admin.
	require.Error(t, (&types.MsgUpdateCollectionDisputeConfig{
		CollectionId: "c1",
		AdminAddress: "not-bech32",
	}).ValidateBasic())

	// Empty collection_id.
	require.Error(t, (&types.MsgUpdateCollectionDisputeConfig{
		AdminAddress: admin,
	}).ValidateBasic())

	// Deposit set but no adjudicators → ValidateCollectionDisputeConfig
	// rejects (adjudicators required if any deposit/penalty is configured).
	err := (&types.MsgUpdateCollectionDisputeConfig{
		CollectionId:                "c1",
		AdminAddress:                admin,
		ServiceAgentDepositRequired: sdk.NewCoins(sdk.NewCoin("uixo", math.NewInt(1_000))),
	}).ValidateBasic()
	require.ErrorIs(t, err, types.ErrDisputeConfigInvalid)

	// Adjudicator with invalid DID.
	err = (&types.MsgUpdateCollectionDisputeConfig{
		CollectionId: "c1",
		AdminAddress: admin,
		Adjudicators: []*types.AdjudicationDid{validAdjudicator("not-a-did", 10)},
	}).ValidateBasic()
	require.ErrorIs(t, err, types.ErrDisputeConfigInvalid)

	// Adjudicator with reward_percentage out of [0,100].
	err = (&types.MsgUpdateCollectionDisputeConfig{
		CollectionId: "c1",
		AdminAddress: admin,
		Adjudicators: []*types.AdjudicationDid{validAdjudicator("did:ixo:adj", 150)},
	}).ValidateBasic()
	require.ErrorIs(t, err, types.ErrDisputeConfigInvalid)

	// Duplicate adjudicator DID.
	err = (&types.MsgUpdateCollectionDisputeConfig{
		CollectionId: "c1",
		AdminAddress: admin,
		Adjudicators: []*types.AdjudicationDid{
			validAdjudicator("did:ixo:dup", 10),
			validAdjudicator("did:ixo:dup", 20),
		},
	}).ValidateBasic()
	require.ErrorIs(t, err, types.ErrDisputeConfigInvalid)

	// Penalty > role deposit_required must reject.
	err = (&types.MsgUpdateCollectionDisputeConfig{
		CollectionId:                "c1",
		AdminAddress:                admin,
		ServiceAgentDepositRequired: sdk.NewCoins(sdk.NewCoin("uixo", math.NewInt(100))),
		PenaltyAmountPerDispute:     sdk.NewCoins(sdk.NewCoin("uixo", math.NewInt(200))),
		Adjudicators:                []*types.AdjudicationDid{validAdjudicator("did:ixo:adj", 10)},
	}).ValidateBasic()
	require.ErrorIs(t, err, types.ErrDisputeConfigInvalid)
}

// --------------------------------------------------------------------------
// MsgAddPerformanceDeposit
// --------------------------------------------------------------------------

func TestMsgAddPerformanceDeposit_ValidateBasic(t *testing.T) {
	agent := validBech32("addagent")

	// Happy path.
	require.NoError(t, (&types.MsgAddPerformanceDeposit{
		CollectionId: "c1",
		AgentAddress: agent,
		Amount:       sdk.NewCoins(sdk.NewCoin("uixo", math.NewInt(1))),
	}).ValidateBasic())

	// Invalid bech32 agent.
	require.Error(t, (&types.MsgAddPerformanceDeposit{
		CollectionId: "c1",
		AgentAddress: "not-bech32",
		Amount:       sdk.NewCoins(sdk.NewCoin("uixo", math.NewInt(1))),
	}).ValidateBasic())

	// Empty collection_id.
	require.Error(t, (&types.MsgAddPerformanceDeposit{
		AgentAddress: agent,
		Amount:       sdk.NewCoins(sdk.NewCoin("uixo", math.NewInt(1))),
	}).ValidateBasic())

	// Zero / empty amount.
	err := (&types.MsgAddPerformanceDeposit{
		CollectionId: "c1",
		AgentAddress: agent,
		Amount:       sdk.Coins{},
	}).ValidateBasic()
	require.ErrorIs(t, err, types.ErrAgentDepositAmountInvalid)
}

// --------------------------------------------------------------------------
// MsgWithdrawPerformanceDeposit
// --------------------------------------------------------------------------

func TestMsgWithdrawPerformanceDeposit_ValidateBasic(t *testing.T) {
	agent := validBech32("wdagent")

	// Happy path: empty amount means "withdraw full balance" — permitted.
	require.NoError(t, (&types.MsgWithdrawPerformanceDeposit{
		CollectionId: "c1",
		AgentAddress: agent,
	}).ValidateBasic())

	// Happy path: explicit partial amount.
	require.NoError(t, (&types.MsgWithdrawPerformanceDeposit{
		CollectionId: "c1",
		AgentAddress: agent,
		Amount:       sdk.NewCoins(sdk.NewCoin("uixo", math.NewInt(50))),
	}).ValidateBasic())

	// Invalid bech32 agent.
	require.Error(t, (&types.MsgWithdrawPerformanceDeposit{
		CollectionId: "c1",
		AgentAddress: "not-bech32",
	}).ValidateBasic())

	// Empty collection_id.
	require.Error(t, (&types.MsgWithdrawPerformanceDeposit{
		AgentAddress: agent,
	}).ValidateBasic())

	// Zero-amount Coins (caller supplied an explicit empty coin rather than
	// omitting the field) must reject — otherwise we'd silently no-op while
	// the user thought they were withdrawing.
	err := (&types.MsgWithdrawPerformanceDeposit{
		CollectionId: "c1",
		AgentAddress: agent,
		Amount:       sdk.Coins{sdk.Coin{Denom: "uixo", Amount: math.NewInt(0)}},
	}).ValidateBasic()
	require.ErrorIs(t, err, types.ErrAgentDepositAmountInvalid)
}

// --------------------------------------------------------------------------
// MsgAdjudicateDispute
// --------------------------------------------------------------------------

func TestMsgAdjudicateDispute_ValidateBasic(t *testing.T) {
	signer := validBech32("adj1")

	// Happy path: no penalty (collection has fixed penalty), no data.
	require.NoError(t, (&types.MsgAdjudicateDispute{
		SubjectId:          "claim-1",
		TargetRole:         types.DisputeTargetRole_target_submitter,
		AdjudicatorDid:     "did:ixo:adj-1",
		AdjudicatorAddress: signer,
		Outcome:            types.DisputeStatus_dispute_awarded,
	}).ValidateBasic())

	// Happy path: with full DisputeData and explicit penalty.
	require.NoError(t, (&types.MsgAdjudicateDispute{
		SubjectId:          "claim-1",
		TargetRole:         types.DisputeTargetRole_target_evaluator,
		AdjudicatorDid:     "did:ixo:adj-1",
		AdjudicatorAddress: signer,
		Outcome:            types.DisputeStatus_dispute_dismissed,
		Data: &types.DisputeData{
			Type:  "application/vnd.ixo+json",
			Proof: "ipfs://cid",
			Uri:   "ipfs://uri",
		},
		PenaltyAmount: sdk.NewCoins(sdk.NewCoin("uixo", math.NewInt(50))),
	}).ValidateBasic())

	// Invalid bech32 signer.
	require.Error(t, (&types.MsgAdjudicateDispute{
		SubjectId:          "claim-1",
		TargetRole:         types.DisputeTargetRole_target_submitter,
		AdjudicatorDid:     "did:ixo:adj-1",
		AdjudicatorAddress: "not-bech32",
		Outcome:            types.DisputeStatus_dispute_awarded,
	}).ValidateBasic())

	// Empty subject_id.
	require.Error(t, (&types.MsgAdjudicateDispute{
		TargetRole:         types.DisputeTargetRole_target_submitter,
		AdjudicatorDid:     "did:ixo:adj-1",
		AdjudicatorAddress: signer,
		Outcome:            types.DisputeStatus_dispute_awarded,
	}).ValidateBasic())

	// Invalid adjudicator DID.
	err := (&types.MsgAdjudicateDispute{
		SubjectId:          "claim-1",
		TargetRole:         types.DisputeTargetRole_target_submitter,
		AdjudicatorDid:     "not-a-did",
		AdjudicatorAddress: signer,
		Outcome:            types.DisputeStatus_dispute_awarded,
	}).ValidateBasic()
	require.ErrorIs(t, err, iidtypes.ErrInvalidDIDFormat)

	// UNSPECIFIED target_role is rejected (only legacy migrated disputes
	// carry this value).
	err = (&types.MsgAdjudicateDispute{
		SubjectId:          "claim-1",
		TargetRole:         types.DisputeTargetRole_target_unspecified,
		AdjudicatorDid:     "did:ixo:adj-1",
		AdjudicatorAddress: signer,
		Outcome:            types.DisputeStatus_dispute_awarded,
	}).ValidateBasic()
	require.ErrorIs(t, err, types.ErrDisputeTargetRoleInvalid)

	// OPEN outcome is rejected (resolutions must be AWARDED or DISMISSED).
	err = (&types.MsgAdjudicateDispute{
		SubjectId:          "claim-1",
		TargetRole:         types.DisputeTargetRole_target_submitter,
		AdjudicatorDid:     "did:ixo:adj-1",
		AdjudicatorAddress: signer,
		Outcome:            types.DisputeStatus_dispute_open,
	}).ValidateBasic()
	require.ErrorIs(t, err, types.ErrAdjudicationInvalidOutcome)

	// Zero-coin penalty rejects (must omit field to defer to collection fixed).
	err = (&types.MsgAdjudicateDispute{
		SubjectId:          "claim-1",
		TargetRole:         types.DisputeTargetRole_target_submitter,
		AdjudicatorDid:     "did:ixo:adj-1",
		AdjudicatorAddress: signer,
		Outcome:            types.DisputeStatus_dispute_awarded,
		PenaltyAmount:      sdk.Coins{sdk.Coin{Denom: "uixo", Amount: math.NewInt(0)}},
	}).ValidateBasic()
	require.ErrorIs(t, err, types.ErrPenaltyAmountInvalid)

	// Partial DisputeData (missing Uri) rejects.
	require.Error(t, (&types.MsgAdjudicateDispute{
		SubjectId:          "claim-1",
		TargetRole:         types.DisputeTargetRole_target_submitter,
		AdjudicatorDid:     "did:ixo:adj-1",
		AdjudicatorAddress: signer,
		Outcome:            types.DisputeStatus_dispute_awarded,
		Data: &types.DisputeData{
			Type:  "application/vnd.ixo+json",
			Proof: "ipfs://cid",
			// Uri intentionally empty
		},
	}).ValidateBasic())
}

// --------------------------------------------------------------------------
// MsgDisputeClaim — v7 target_role gate
// --------------------------------------------------------------------------

func TestMsgDisputeClaim_ValidateBasic_TargetRole(t *testing.T) {
	agent := validBech32("disp")

	// Happy path: SUBMITTER.
	require.NoError(t, (&types.MsgDisputeClaim{
		AgentDid:     "did:ixo:disputer",
		AgentAddress: agent,
		SubjectId:    "claim-1",
		DisputeType:  1,
		TargetRole:   types.DisputeTargetRole_target_submitter,
		Data: &types.DisputeData{
			Type:  "application/vnd.ixo+json",
			Proof: "ipfs://cid",
			Uri:   "ipfs://uri",
		},
	}).ValidateBasic())

	// Happy path: EVALUATOR.
	require.NoError(t, (&types.MsgDisputeClaim{
		AgentDid:     "did:ixo:disputer",
		AgentAddress: agent,
		SubjectId:    "claim-1",
		DisputeType:  1,
		TargetRole:   types.DisputeTargetRole_target_evaluator,
		Data: &types.DisputeData{
			Type:  "application/vnd.ixo+json",
			Proof: "ipfs://cid",
			Uri:   "ipfs://uri",
		},
	}).ValidateBasic())

	// UNSPECIFIED target_role rejects (new disputes must pick a role).
	err := (&types.MsgDisputeClaim{
		AgentDid:     "did:ixo:disputer",
		AgentAddress: agent,
		SubjectId:    "claim-1",
		DisputeType:  1,
		TargetRole:   types.DisputeTargetRole_target_unspecified,
		Data: &types.DisputeData{
			Type:  "application/vnd.ixo+json",
			Proof: "ipfs://cid",
			Uri:   "ipfs://uri",
		},
	}).ValidateBasic()
	require.ErrorIs(t, err, types.ErrDisputeTargetRoleInvalid)

	// Missing Data rejects.
	require.Error(t, (&types.MsgDisputeClaim{
		AgentDid:     "did:ixo:disputer",
		AgentAddress: agent,
		SubjectId:    "claim-1",
		DisputeType:  1,
		TargetRole:   types.DisputeTargetRole_target_submitter,
	}).ValidateBasic())
}

// --------------------------------------------------------------------------
// ValidateCollectionDisputeConfig — direct edge cases
// --------------------------------------------------------------------------

func TestValidateCollectionDisputeConfig(t *testing.T) {
	// Fully empty is permitted (no dispute features configured).
	require.NoError(t, types.ValidateCollectionDisputeConfig(types.CollectionDisputeConfig{}))

	// Negative MinDepositPeriod is rejected by Collection validation later;
	// the config helper itself accepts non-negative durations.
	require.NoError(t, types.ValidateCollectionDisputeConfig(types.CollectionDisputeConfig{
		MinDepositPeriod: time.Hour,
		Adjudicators:     []*types.AdjudicationDid{validAdjudicator("did:ixo:a", 10)},
	}))

	// Nil entry in adjudicators slice is rejected.
	err := types.ValidateCollectionDisputeConfig(types.CollectionDisputeConfig{
		Adjudicators: []*types.AdjudicationDid{nil},
	})
	require.ErrorIs(t, err, types.ErrDisputeConfigInvalid)

	// Adjudicators-only config (no deposits) is permitted.
	require.NoError(t, types.ValidateCollectionDisputeConfig(types.CollectionDisputeConfig{
		Adjudicators: []*types.AdjudicationDid{validAdjudicator("did:ixo:a", 0)},
	}))

	// Adjudicators-only with multiple unique entries permitted.
	require.NoError(t, types.ValidateCollectionDisputeConfig(types.CollectionDisputeConfig{
		Adjudicators: []*types.AdjudicationDid{
			validAdjudicator("did:ixo:a", 5),
			validAdjudicator("did:ixo:b", 15),
			validAdjudicator("did:ixo:c", 100),
		},
	}))
}
