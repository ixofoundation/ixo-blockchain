package keeper_test

import (
	"time"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ixofoundation/ixo-blockchain/v6/app/apptesting"
	claimskeeper "github.com/ixofoundation/ixo-blockchain/v6/x/claims/keeper"
	"github.com/ixofoundation/ixo-blockchain/v6/x/claims/types"
)

// v7 feature surface that lives off the wasm payment path:
//
//   - MsgUpdateCollectionQuota         (commit 7b74603d)
//   - MsgUpdateCollectionDisputeConfig (commit 1f2b56f5, MsgServer-side)
//   - MsgAddPerformanceDeposit         (commit 1f2b56f5)
//   - MsgWithdrawPerformanceDeposit    (commit 1f2b56f5)
//   - EvaluationStatus.FLAGGED         (commit be64ddd9, counter mutations)
//
// AdjudicateDispute and the full EvaluateClaim path stay in interchaintest
// (they need wasm-backed payments and IID ante DID-key auth respectively).

// --------------------------------------------------------------------------
// MsgUpdateCollectionQuota
// --------------------------------------------------------------------------

// Happy path: raise an unlimited collection to a positive cap.
func (s *KeeperTestSuite) TestMsgUpdateCollectionQuota_Raise() {
	s.SetupTest()
	admin, _ := s.seedAdminCollection("quota-raise")
	ms := claimskeeper.NewMsgServerImpl(&s.App.ClaimsKeeper)

	_, err := ms.UpdateCollectionQuota(s.Ctx, &types.MsgUpdateCollectionQuota{
		CollectionId: "quota-raise",
		Quota:        100,
		AdminAddress: admin,
	})
	s.Require().NoError(err)
	got, _ := s.App.ClaimsKeeper.GetCollection(s.Ctx, "quota-raise")
	s.Require().EqualValues(100, got.Quota)
}

// Zero quota means "unlimited" and is always permitted (it relaxes the gate).
func (s *KeeperTestSuite) TestMsgUpdateCollectionQuota_SetUnlimited() {
	s.SetupTest()
	admin, _ := s.seedAdminCollection("quota-zero")
	ms := claimskeeper.NewMsgServerImpl(&s.App.ClaimsKeeper)

	// First cap at 50, then relax to unlimited.
	_, err := ms.UpdateCollectionQuota(s.Ctx, &types.MsgUpdateCollectionQuota{
		CollectionId: "quota-zero",
		Quota:        50,
		AdminAddress: admin,
	})
	s.Require().NoError(err)
	_, err = ms.UpdateCollectionQuota(s.Ctx, &types.MsgUpdateCollectionQuota{
		CollectionId: "quota-zero",
		Quota:        0,
		AdminAddress: admin,
	})
	s.Require().NoError(err)
	got, _ := s.App.ClaimsKeeper.GetCollection(s.Ctx, "quota-zero")
	s.Require().EqualValues(0, got.Quota)
}

// Rejecting quota < count is the load-bearing rule: it prevents an admin
// from retroactively invalidating claims that are already in the collection.
func (s *KeeperTestSuite) TestMsgUpdateCollectionQuota_RejectBelowCount() {
	s.SetupTest()
	admin, _ := s.seedAdminCollection("quota-below")
	c, _ := s.App.ClaimsKeeper.GetCollection(s.Ctx, "quota-below")
	c.Count = 10
	s.App.ClaimsKeeper.SetCollection(s.Ctx, c)
	ms := claimskeeper.NewMsgServerImpl(&s.App.ClaimsKeeper)

	_, err := ms.UpdateCollectionQuota(s.Ctx, &types.MsgUpdateCollectionQuota{
		CollectionId: "quota-below",
		Quota:        5,
		AdminAddress: admin,
	})
	s.Require().ErrorIs(err, types.ErrCollectionQuotaBelowCount,
		"quota strictly below count must be rejected")

	// State must be unchanged.
	got, _ := s.App.ClaimsKeeper.GetCollection(s.Ctx, "quota-below")
	s.Require().EqualValues(0, got.Quota, "rejected update must not mutate quota")
	s.Require().EqualValues(10, got.Count, "rejected update must not mutate count")
}

// quota == count must succeed (boundary case — not < count).
func (s *KeeperTestSuite) TestMsgUpdateCollectionQuota_EqualToCount() {
	s.SetupTest()
	admin, _ := s.seedAdminCollection("quota-equal")
	c, _ := s.App.ClaimsKeeper.GetCollection(s.Ctx, "quota-equal")
	c.Count = 7
	s.App.ClaimsKeeper.SetCollection(s.Ctx, c)
	ms := claimskeeper.NewMsgServerImpl(&s.App.ClaimsKeeper)

	_, err := ms.UpdateCollectionQuota(s.Ctx, &types.MsgUpdateCollectionQuota{
		CollectionId: "quota-equal",
		Quota:        7,
		AdminAddress: admin,
	})
	s.Require().NoError(err)
}

// Stranger cannot set quota.
func (s *KeeperTestSuite) TestMsgUpdateCollectionQuota_Unauthorized() {
	s.SetupTest()
	_, _ = s.seedAdminCollection("quota-auth")
	stranger := apptesting.RandomAccountAddress().String()
	ms := claimskeeper.NewMsgServerImpl(&s.App.ClaimsKeeper)

	_, err := ms.UpdateCollectionQuota(s.Ctx, &types.MsgUpdateCollectionQuota{
		CollectionId: "quota-auth",
		Quota:        50,
		AdminAddress: stranger,
	})
	s.Require().ErrorContains(err, "unauthorized")
}

// Unknown collection_id returns the not-found error from GetCollection.
func (s *KeeperTestSuite) TestMsgUpdateCollectionQuota_UnknownCollection() {
	s.SetupTest()
	admin := apptesting.RandomAccountAddress().String()
	ms := claimskeeper.NewMsgServerImpl(&s.App.ClaimsKeeper)

	_, err := ms.UpdateCollectionQuota(s.Ctx, &types.MsgUpdateCollectionQuota{
		CollectionId: "does-not-exist",
		Quota:        1,
		AdminAddress: admin,
	})
	s.Require().Error(err)
}

// ValidateBasic-level: msg with empty collection_id is rejected at antehandler.
func (s *KeeperTestSuite) TestMsgUpdateCollectionQuota_ValidateBasic() {
	signer := apptesting.RandomAccountAddress().String()

	// Empty collection_id.
	err := (&types.MsgUpdateCollectionQuota{AdminAddress: signer}).ValidateBasic()
	s.Require().Error(err)

	// Invalid bech32 admin.
	err = (&types.MsgUpdateCollectionQuota{CollectionId: "x", AdminAddress: "not-bech32"}).ValidateBasic()
	s.Require().Error(err)

	// Happy path.
	err = (&types.MsgUpdateCollectionQuota{CollectionId: "x", AdminAddress: signer}).ValidateBasic()
	s.Require().NoError(err)
}

// --------------------------------------------------------------------------
// MsgUpdateCollectionDisputeConfig — admin path (no DID auth needed)
// --------------------------------------------------------------------------

func adjudicatorDid(did string, pct int64) *types.AdjudicationDid {
	return &types.AdjudicationDid{
		Did:              did,
		RewardPercentage: math.LegacyNewDec(pct),
	}
}

// Happy path: replace all deposit/penalty/adjudicator fields atomically.
func (s *KeeperTestSuite) TestMsgUpdateCollectionDisputeConfig_FullUpdate() {
	s.SetupTest()
	admin, _ := s.seedAdminCollection("dispute-cfg")
	ms := claimskeeper.NewMsgServerImpl(&s.App.ClaimsKeeper)

	period := 24 * time.Hour
	_, err := ms.UpdateCollectionDisputeConfig(s.Ctx, &types.MsgUpdateCollectionDisputeConfig{
		CollectionId:                "dispute-cfg",
		AdminAddress:                admin,
		ServiceAgentDepositRequired: sdk.NewCoins(sdk.NewCoin("uixo", math.NewInt(1_000_000))),
		EvaluatorDepositRequired:    sdk.NewCoins(sdk.NewCoin("uixo", math.NewInt(500_000))),
		DisputeDepositAmount:        sdk.NewCoins(sdk.NewCoin("uixo", math.NewInt(100_000))),
		PenaltyAmountPerDispute:     sdk.NewCoins(sdk.NewCoin("uixo", math.NewInt(50_000))),
		MinDepositPeriod:            period,
		Adjudicators:                []*types.AdjudicationDid{adjudicatorDid("did:ixo:adj-1", 20)},
	})
	s.Require().NoError(err)

	got, _ := s.App.ClaimsKeeper.GetCollection(s.Ctx, "dispute-cfg")
	s.Require().Equal("1000000uixo", got.ServiceAgentDepositRequired.String())
	s.Require().Equal("500000uixo", got.EvaluatorDepositRequired.String())
	s.Require().Equal("100000uixo", got.DisputeDepositAmount.String())
	s.Require().Equal("50000uixo", got.PenaltyAmountPerDispute.String())
	s.Require().Equal(period, got.MinDepositPeriod)
	s.Require().Len(got.Adjudicators, 1)
	s.Require().Equal("did:ixo:adj-1", got.Adjudicators[0].Did)
}

// Refuse to clear the adjudicator whitelist while any dispute is OPEN —
// otherwise those disputes (with locked deposits) become unsettleable.
func (s *KeeperTestSuite) TestMsgUpdateCollectionDisputeConfig_BlockClearWithOpenDisputes() {
	s.SetupTest()
	admin, _ := s.seedAdminCollection("dispute-cfg-block")
	c, _ := s.App.ClaimsKeeper.GetCollection(s.Ctx, "dispute-cfg-block")
	c.DisputesOpen = 1
	s.App.ClaimsKeeper.SetCollection(s.Ctx, c)
	ms := claimskeeper.NewMsgServerImpl(&s.App.ClaimsKeeper)

	_, err := ms.UpdateCollectionDisputeConfig(s.Ctx, &types.MsgUpdateCollectionDisputeConfig{
		CollectionId: "dispute-cfg-block",
		AdminAddress: admin,
		Adjudicators: nil, // clearing
	})
	s.Require().ErrorIs(err, types.ErrAdjudicationNotConfigured,
		"clearing adjudicators must be blocked while disputes_open > 0")

	// Confirm Adjudicators on Collection still empty (was never set).
	got, _ := s.App.ClaimsKeeper.GetCollection(s.Ctx, "dispute-cfg-block")
	s.Require().Empty(got.Adjudicators)
}

// Reducing (but not emptying) the whitelist must succeed even with open disputes.
func (s *KeeperTestSuite) TestMsgUpdateCollectionDisputeConfig_ReduceWhileOpenDisputes() {
	s.SetupTest()
	admin, _ := s.seedAdminCollection("dispute-cfg-reduce")
	c, _ := s.App.ClaimsKeeper.GetCollection(s.Ctx, "dispute-cfg-reduce")
	c.DisputesOpen = 2
	c.Adjudicators = []*types.AdjudicationDid{
		adjudicatorDid("did:ixo:keep", 10),
		adjudicatorDid("did:ixo:drop", 15),
	}
	s.App.ClaimsKeeper.SetCollection(s.Ctx, c)
	ms := claimskeeper.NewMsgServerImpl(&s.App.ClaimsKeeper)

	_, err := ms.UpdateCollectionDisputeConfig(s.Ctx, &types.MsgUpdateCollectionDisputeConfig{
		CollectionId: "dispute-cfg-reduce",
		AdminAddress: admin,
		Adjudicators: []*types.AdjudicationDid{adjudicatorDid("did:ixo:keep", 10)},
	})
	s.Require().NoError(err, "shrinking the whitelist is permitted while disputes are open")

	got, _ := s.App.ClaimsKeeper.GetCollection(s.Ctx, "dispute-cfg-reduce")
	s.Require().Len(got.Adjudicators, 1)
	s.Require().Equal("did:ixo:keep", got.Adjudicators[0].Did)
}

// Stranger cannot mutate dispute config.
func (s *KeeperTestSuite) TestMsgUpdateCollectionDisputeConfig_Unauthorized() {
	s.SetupTest()
	_, _ = s.seedAdminCollection("dispute-cfg-auth")
	stranger := apptesting.RandomAccountAddress().String()
	ms := claimskeeper.NewMsgServerImpl(&s.App.ClaimsKeeper)

	_, err := ms.UpdateCollectionDisputeConfig(s.Ctx, &types.MsgUpdateCollectionDisputeConfig{
		CollectionId: "dispute-cfg-auth",
		AdminAddress: stranger,
	})
	s.Require().ErrorContains(err, "unauthorized")
}

// --------------------------------------------------------------------------
// MsgAddPerformanceDeposit / MsgWithdrawPerformanceDeposit
// --------------------------------------------------------------------------

// seedDepositCollection creates a collection backed by a real module-account
// escrow + funds the agent so we can exercise SendCoins through the bank
// keeper without wasm.
func (s *KeeperTestSuite) seedDepositCollection(id string, minDepositPeriod time.Duration) (admin string, escrow sdk.AccAddress, agent sdk.AccAddress) {
	adminAddr := apptesting.RandomAccountAddress()
	escrow = apptesting.RandomAccountAddress()
	agent = apptesting.RandomAccountAddress()

	c := types.Collection{
		Id:               id,
		Entity:           "did:ixo:entity-deposit",
		Admin:            adminAddr.String(),
		Protocol:         "did:ixo:protocol-deposit",
		State:            types.CollectionState_open,
		EscrowAccount:    escrow.String(),
		MinDepositPeriod: minDepositPeriod,
	}
	s.App.ClaimsKeeper.SetCollection(s.Ctx, c)

	s.FundAcc(agent, sdk.NewCoins(sdk.NewCoin("uixo", math.NewInt(10_000_000))))
	return adminAddr.String(), escrow, agent
}

// AddPerformanceDeposit moves funds agent → escrow and writes the balance.
// Idempotent across calls: balance accumulates.
func (s *KeeperTestSuite) TestMsgAddPerformanceDeposit_Topup() {
	s.SetupTest()
	_, escrow, agent := s.seedDepositCollection("dep-topup", 0)
	ms := claimskeeper.NewMsgServerImpl(&s.App.ClaimsKeeper)

	resp, err := ms.AddPerformanceDeposit(s.Ctx, &types.MsgAddPerformanceDeposit{
		CollectionId: "dep-topup",
		AgentAddress: agent.String(),
		Amount:       sdk.NewCoins(sdk.NewCoin("uixo", math.NewInt(1_000))),
	})
	s.Require().NoError(err)
	s.Require().Equal("1000uixo", resp.NewBalance.String())

	bal := s.App.BankKeeper.GetBalance(s.Ctx, escrow, "uixo")
	s.Require().Equal(int64(1_000), bal.Amount.Int64(), "funds must land in collection escrow")

	// Second top-up accumulates.
	resp, err = ms.AddPerformanceDeposit(s.Ctx, &types.MsgAddPerformanceDeposit{
		CollectionId: "dep-topup",
		AgentAddress: agent.String(),
		Amount:       sdk.NewCoins(sdk.NewCoin("uixo", math.NewInt(2_500))),
	})
	s.Require().NoError(err)
	s.Require().Equal("3500uixo", resp.NewBalance.String())

	stored, err := s.App.ClaimsKeeper.GetAgentDepositBalance(s.Ctx, "dep-topup", agent.String())
	s.Require().NoError(err)
	s.Require().Equal("3500uixo", stored.Amount.String())
}

// min_deposit_period rolls withdrawable_at forward on every top-up to
// max(current, now + period). This closes the deposit-submit-withdraw
// exploit window.
func (s *KeeperTestSuite) TestMsgAddPerformanceDeposit_WithdrawableAtRollForward() {
	s.SetupTest()
	period := time.Hour
	_, _, agent := s.seedDepositCollection("dep-lock", period)
	ms := claimskeeper.NewMsgServerImpl(&s.App.ClaimsKeeper)

	t0 := s.Ctx.BlockTime()
	_, err := ms.AddPerformanceDeposit(s.Ctx, &types.MsgAddPerformanceDeposit{
		CollectionId: "dep-lock",
		AgentAddress: agent.String(),
		Amount:       sdk.NewCoins(sdk.NewCoin("uixo", math.NewInt(500))),
	})
	s.Require().NoError(err)
	bal, err := s.App.ClaimsKeeper.GetAgentDepositBalance(s.Ctx, "dep-lock", agent.String())
	s.Require().NoError(err)
	s.Require().NotNil(bal.WithdrawableAt)
	s.Require().True(bal.WithdrawableAt.Equal(t0.Add(period)), "first top-up sets withdrawable_at = blocktime + period")

	// Advance time 30min (still inside the lock window), top up again — the
	// new top-up extends the lock to (t0 + 30min) + period, which is strictly
	// later than the first lock at (t0 + period).
	s.Ctx = s.Ctx.WithBlockTime(t0.Add(30 * time.Minute))
	_, err = ms.AddPerformanceDeposit(s.Ctx, &types.MsgAddPerformanceDeposit{
		CollectionId: "dep-lock",
		AgentAddress: agent.String(),
		Amount:       sdk.NewCoins(sdk.NewCoin("uixo", math.NewInt(500))),
	})
	s.Require().NoError(err)
	bal, _ = s.App.ClaimsKeeper.GetAgentDepositBalance(s.Ctx, "dep-lock", agent.String())
	s.Require().True(bal.WithdrawableAt.After(t0.Add(period)),
		"top-up while still locked must extend the lock further into the future")
}

// min_deposit_period == 0 → no lock recorded (legacy/disabled mode).
func (s *KeeperTestSuite) TestMsgAddPerformanceDeposit_LegacyNoLock() {
	s.SetupTest()
	_, _, agent := s.seedDepositCollection("dep-nolock", 0)
	ms := claimskeeper.NewMsgServerImpl(&s.App.ClaimsKeeper)

	_, err := ms.AddPerformanceDeposit(s.Ctx, &types.MsgAddPerformanceDeposit{
		CollectionId: "dep-nolock",
		AgentAddress: agent.String(),
		Amount:       sdk.NewCoins(sdk.NewCoin("uixo", math.NewInt(123))),
	})
	s.Require().NoError(err)
	bal, _ := s.App.ClaimsKeeper.GetAgentDepositBalance(s.Ctx, "dep-nolock", agent.String())
	s.Require().Nil(bal.WithdrawableAt, "period=0 means withdrawable_at is never set")
}

// Negative path: amount must be strictly positive.
func (s *KeeperTestSuite) TestMsgAddPerformanceDeposit_RejectZeroAmount() {
	s.SetupTest()
	_, _, agent := s.seedDepositCollection("dep-zero", 0)
	ms := claimskeeper.NewMsgServerImpl(&s.App.ClaimsKeeper)

	_, err := ms.AddPerformanceDeposit(s.Ctx, &types.MsgAddPerformanceDeposit{
		CollectionId: "dep-zero",
		AgentAddress: agent.String(),
		Amount:       sdk.Coins{},
	})
	s.Require().ErrorIs(err, types.ErrAgentDepositAmountInvalid)
}

// Withdraw round-trip: top-up, advance past lock, pull all funds out.
func (s *KeeperTestSuite) TestMsgWithdrawPerformanceDeposit_FullWithdraw() {
	s.SetupTest()
	_, escrow, agent := s.seedDepositCollection("dep-wd", time.Minute)
	ms := claimskeeper.NewMsgServerImpl(&s.App.ClaimsKeeper)

	_, err := ms.AddPerformanceDeposit(s.Ctx, &types.MsgAddPerformanceDeposit{
		CollectionId: "dep-wd",
		AgentAddress: agent.String(),
		Amount:       sdk.NewCoins(sdk.NewCoin("uixo", math.NewInt(700))),
	})
	s.Require().NoError(err)

	// Advance past the lock window.
	s.Ctx = s.Ctx.WithBlockTime(s.Ctx.BlockTime().Add(2 * time.Minute))

	resp, err := ms.WithdrawPerformanceDeposit(s.Ctx, &types.MsgWithdrawPerformanceDeposit{
		CollectionId: "dep-wd",
		AgentAddress: agent.String(),
		Amount:       sdk.Coins{}, // empty == full balance
	})
	s.Require().NoError(err)
	s.Require().Equal("700uixo", resp.Withdrawn.String())
	s.Require().Equal("", resp.RemainingBalance.String())
	s.Require().Equal(int64(0), s.App.BankKeeper.GetBalance(s.Ctx, escrow, "uixo").Amount.Int64())
}

// Withdrawal partial: balance partially drains, lock window irrelevant once
// the deposit is past withdrawable_at.
func (s *KeeperTestSuite) TestMsgWithdrawPerformanceDeposit_Partial() {
	s.SetupTest()
	_, _, agent := s.seedDepositCollection("dep-part", 0)
	ms := claimskeeper.NewMsgServerImpl(&s.App.ClaimsKeeper)

	_, _ = ms.AddPerformanceDeposit(s.Ctx, &types.MsgAddPerformanceDeposit{
		CollectionId: "dep-part",
		AgentAddress: agent.String(),
		Amount:       sdk.NewCoins(sdk.NewCoin("uixo", math.NewInt(1_000))),
	})

	resp, err := ms.WithdrawPerformanceDeposit(s.Ctx, &types.MsgWithdrawPerformanceDeposit{
		CollectionId: "dep-part",
		AgentAddress: agent.String(),
		Amount:       sdk.NewCoins(sdk.NewCoin("uixo", math.NewInt(300))),
	})
	s.Require().NoError(err)
	s.Require().Equal("300uixo", resp.Withdrawn.String())
	s.Require().Equal("700uixo", resp.RemainingBalance.String())
}

// Lock gate: withdrawing while inside the min_deposit_period window fails.
func (s *KeeperTestSuite) TestMsgWithdrawPerformanceDeposit_LockedReject() {
	s.SetupTest()
	_, _, agent := s.seedDepositCollection("dep-locked", time.Hour)
	ms := claimskeeper.NewMsgServerImpl(&s.App.ClaimsKeeper)

	_, _ = ms.AddPerformanceDeposit(s.Ctx, &types.MsgAddPerformanceDeposit{
		CollectionId: "dep-locked",
		AgentAddress: agent.String(),
		Amount:       sdk.NewCoins(sdk.NewCoin("uixo", math.NewInt(500))),
	})

	_, err := ms.WithdrawPerformanceDeposit(s.Ctx, &types.MsgWithdrawPerformanceDeposit{
		CollectionId: "dep-locked",
		AgentAddress: agent.String(),
		Amount:       sdk.Coins{},
	})
	s.Require().ErrorIs(err, types.ErrAgentDepositLocked,
		"withdrawal inside min_deposit_period must be rejected")
}

// Withdraw exceeding balance rejects without partial settlement.
func (s *KeeperTestSuite) TestMsgWithdrawPerformanceDeposit_OverBalance() {
	s.SetupTest()
	_, _, agent := s.seedDepositCollection("dep-over", 0)
	ms := claimskeeper.NewMsgServerImpl(&s.App.ClaimsKeeper)

	_, _ = ms.AddPerformanceDeposit(s.Ctx, &types.MsgAddPerformanceDeposit{
		CollectionId: "dep-over",
		AgentAddress: agent.String(),
		Amount:       sdk.NewCoins(sdk.NewCoin("uixo", math.NewInt(100))),
	})

	_, err := ms.WithdrawPerformanceDeposit(s.Ctx, &types.MsgWithdrawPerformanceDeposit{
		CollectionId: "dep-over",
		AgentAddress: agent.String(),
		Amount:       sdk.NewCoins(sdk.NewCoin("uixo", math.NewInt(101))),
	})
	s.Require().ErrorIs(err, types.ErrAgentDepositAmountExceedsBalance)
}

// Withdraw without any prior deposit fails (no balance entry to draw from).
func (s *KeeperTestSuite) TestMsgWithdrawPerformanceDeposit_NoBalance() {
	s.SetupTest()
	_, _, agent := s.seedDepositCollection("dep-nobal", 0)
	ms := claimskeeper.NewMsgServerImpl(&s.App.ClaimsKeeper)

	_, err := ms.WithdrawPerformanceDeposit(s.Ctx, &types.MsgWithdrawPerformanceDeposit{
		CollectionId: "dep-nobal",
		AgentAddress: agent.String(),
		Amount:       sdk.NewCoins(sdk.NewCoin("uixo", math.NewInt(1))),
	})
	s.Require().Error(err)
}

// --------------------------------------------------------------------------
// Dispute subject index — direct keeper coverage
// --------------------------------------------------------------------------
//
// MsgDisputeClaim and MsgAdjudicateDispute require IID-key auth + wasm
// payment teardown, so they live in interchaintest. The keeper-level
// subject-index transitions (open/dismissed/awarded) are testable here
// against the bare Dispute store.

func (s *KeeperTestSuite) TestCanFileNewDisputeForSubject_Open() {
	s.SetupTest()
	const subject = "claim-1"

	// No prior dispute → allowed.
	s.Require().NoError(
		s.App.ClaimsKeeper.CanFileNewDisputeForSubject(s.Ctx, subject, types.DisputeTargetRole_target_submitter))

	// Now seed an OPEN dispute and confirm the same call is blocked.
	// IsValidDispute requires both Proof and Uri to be non-empty.
	d := types.Dispute{
		SubjectId:  subject,
		Type:       1,
		Status:     types.DisputeStatus_dispute_open,
		Data:       &types.DisputeData{Proof: "proof-open-1", Uri: "ipfs://uri-open"},
		TargetRole: types.DisputeTargetRole_target_submitter,
	}
	s.App.ClaimsKeeper.SetDispute(s.Ctx, d)
	s.App.ClaimsKeeper.SetDisputeSubjectIndex(s.Ctx, subject, types.DisputeTargetRole_target_submitter, d.Data.Proof)

	err := s.App.ClaimsKeeper.CanFileNewDisputeForSubject(s.Ctx, subject, types.DisputeTargetRole_target_submitter)
	s.Require().ErrorIs(err, types.ErrDisputeAlreadyOpenForSubjectRole)

	// Different role is independent — still allowed.
	s.Require().NoError(
		s.App.ClaimsKeeper.CanFileNewDisputeForSubject(s.Ctx, subject, types.DisputeTargetRole_target_evaluator))
}

// AWARDED permanently blocks; DISMISSED allows supersede.
func (s *KeeperTestSuite) TestCanFileNewDisputeForSubject_AwardedVsDismissed() {
	s.SetupTest()
	const subject = "claim-2"

	award := types.Dispute{
		SubjectId:  subject,
		Type:       1,
		Status:     types.DisputeStatus_dispute_awarded,
		Data:       &types.DisputeData{Proof: "proof-award", Uri: "ipfs://uri-award"},
		TargetRole: types.DisputeTargetRole_target_submitter,
	}
	s.App.ClaimsKeeper.SetDispute(s.Ctx, award)
	s.App.ClaimsKeeper.SetDisputeSubjectIndex(s.Ctx, subject, types.DisputeTargetRole_target_submitter, award.Data.Proof)

	err := s.App.ClaimsKeeper.CanFileNewDisputeForSubject(s.Ctx, subject, types.DisputeTargetRole_target_submitter)
	s.Require().ErrorIs(err, types.ErrDisputeAlreadyAwardedForSubjectRole)

	// Replace with DISMISSED — supersede is allowed.
	dismissed := types.Dispute{
		SubjectId:  subject,
		Type:       1,
		Status:     types.DisputeStatus_dispute_dismissed,
		Data:       &types.DisputeData{Proof: "proof-dismiss", Uri: "ipfs://uri-dismiss"},
		TargetRole: types.DisputeTargetRole_target_evaluator,
	}
	s.App.ClaimsKeeper.SetDispute(s.Ctx, dismissed)
	s.App.ClaimsKeeper.SetDisputeSubjectIndex(s.Ctx, subject, types.DisputeTargetRole_target_evaluator, dismissed.Data.Proof)

	s.Require().NoError(
		s.App.ClaimsKeeper.CanFileNewDisputeForSubject(s.Ctx, subject, types.DisputeTargetRole_target_evaluator))
}

// --------------------------------------------------------------------------
// EvaluationStatus.FLAGGED — re-evaluation gate negative paths
// --------------------------------------------------------------------------
//
// EvaluateClaim's full happy path requires wasm (payment processing),
// so it lives in interchaintest. The re-evaluation gate runs BEFORE the
// payment path, so the negative cases — self-reflag, terminal-locks,
// history-aware blocking — are reachable from the keeper.
//
// All three negative tests share this fixture: a collection whose admin is
// the msg's AdminAddress, and a claim already flagged by `firstAgent`.

func (s *KeeperTestSuite) seedFlaggedClaim(collectionID, claimID, firstAgent string, alsoInHistory ...string) (admin string) {
	adminAddr := apptesting.RandomAccountAddress()
	admin = adminAddr.String()

	s.App.ClaimsKeeper.SetCollection(s.Ctx, types.Collection{
		Id:       collectionID,
		Entity:   "did:ixo:entity-flagged",
		Admin:    admin,
		Protocol: "did:ixo:protocol-flagged",
		State:    types.CollectionState_open,
		Payments: &types.Payments{
			Submission: &types.Payment{},
			Evaluation: &types.Payment{},
			Approval:   &types.Payment{},
			Rejection:  &types.Payment{},
		},
	})

	priorEvaluation := &types.Evaluation{
		ClaimId:      claimID,
		CollectionId: collectionID,
		AgentAddress: firstAgent,
		AgentDid:     "did:ixo:agent-1",
		Status:       types.EvaluationStatus_flagged,
	}
	history := make([]*types.Evaluation, 0, len(alsoInHistory))
	for _, h := range alsoInHistory {
		history = append(history, &types.Evaluation{
			ClaimId:      claimID,
			CollectionId: collectionID,
			AgentAddress: h,
			Status:       types.EvaluationStatus_flagged,
		})
	}

	s.App.ClaimsKeeper.SetClaim(s.Ctx, types.Claim{
		CollectionId:      collectionID,
		ClaimId:           claimID,
		AgentAddress:      apptesting.RandomAccountAddress().String(),
		AgentDid:          "did:ixo:claim-author",
		Evaluation:        priorEvaluation,
		EvaluationHistory: history,
	})
	return admin
}

// Re-flagging by the agent who originally flagged the claim is rejected —
// they already registered their flag.
func (s *KeeperTestSuite) TestEvaluateClaim_FLAGGED_RejectSelfReflagCurrent() {
	s.SetupTest()
	firstAgent := apptesting.RandomAccountAddress().String()
	admin := s.seedFlaggedClaim("flag-self-cur", "claim-flag-self-cur", firstAgent)
	ms := claimskeeper.NewMsgServerImpl(&s.App.ClaimsKeeper)

	_, err := ms.EvaluateClaim(s.Ctx, &types.MsgEvaluateClaim{
		ClaimId:      "claim-flag-self-cur",
		CollectionId: "flag-self-cur",
		AdminAddress: admin,
		AgentDid:     "did:ixo:agent-1",
		AgentAddress: firstAgent, // <-- same as the prior flagger
		Status:       types.EvaluationStatus_flagged,
	})
	s.Require().ErrorIs(err, types.ErrSelfReFlag,
		"the same agent must not be allowed to re-flag (current evaluation match)")
}

// Re-flagging by an agent who flagged earlier and is now in
// evaluation_history (a bob-flag intervened) is also rejected. This guards
// against the flag-bomb attack: tester→bob→tester would otherwise create
// two history entries from the same agent.
func (s *KeeperTestSuite) TestEvaluateClaim_FLAGGED_RejectSelfReflagHistory() {
	s.SetupTest()
	tester := apptesting.RandomAccountAddress().String()
	bob := apptesting.RandomAccountAddress().String()
	// Tester flagged first (so they're now in history); bob's flag is the
	// current evaluation.
	admin := s.seedFlaggedClaim("flag-self-hist", "claim-flag-self-hist", bob, tester)
	ms := claimskeeper.NewMsgServerImpl(&s.App.ClaimsKeeper)

	_, err := ms.EvaluateClaim(s.Ctx, &types.MsgEvaluateClaim{
		ClaimId:      "claim-flag-self-hist",
		CollectionId: "flag-self-hist",
		AdminAddress: admin,
		AgentDid:     "did:ixo:agent-tester",
		AgentAddress: tester, // <-- already in evaluation_history
		Status:       types.EvaluationStatus_flagged,
	})
	s.Require().ErrorIs(err, types.ErrSelfReFlag,
		"history-aware self-reflag must be blocked across an intervening flag from another agent")
}

// A terminal prior evaluation (anything other than FLAGGED) locks the
// claim — the duplicate-evaluation check fires before the self-reflag
// check.
func (s *KeeperTestSuite) TestEvaluateClaim_FLAGGED_TerminalLocked() {
	s.SetupTest()
	admin := apptesting.RandomAccountAddress().String()
	s.App.ClaimsKeeper.SetCollection(s.Ctx, types.Collection{
		Id:       "flag-locked",
		Entity:   "did:ixo:entity-locked",
		Admin:    admin,
		Protocol: "did:ixo:protocol-locked",
		State:    types.CollectionState_open,
		Payments: &types.Payments{
			Submission: &types.Payment{},
			Evaluation: &types.Payment{},
			Approval:   &types.Payment{},
			Rejection:  &types.Payment{},
		},
	})
	s.App.ClaimsKeeper.SetClaim(s.Ctx, types.Claim{
		CollectionId: "flag-locked",
		ClaimId:      "claim-locked",
		AgentAddress: apptesting.RandomAccountAddress().String(),
		AgentDid:     "did:ixo:claim-author",
		Evaluation: &types.Evaluation{
			ClaimId:      "claim-locked",
			CollectionId: "flag-locked",
			AgentAddress: apptesting.RandomAccountAddress().String(),
			Status:       types.EvaluationStatus_approved,
		},
	})
	ms := claimskeeper.NewMsgServerImpl(&s.App.ClaimsKeeper)

	_, err := ms.EvaluateClaim(s.Ctx, &types.MsgEvaluateClaim{
		ClaimId:      "claim-locked",
		CollectionId: "flag-locked",
		AdminAddress: admin,
		AgentDid:     "did:ixo:agent-late",
		AgentAddress: apptesting.RandomAccountAddress().String(),
		Status:       types.EvaluationStatus_flagged,
	})
	s.Require().ErrorIs(err, types.ErrClaimDuplicateEvaluation,
		"prior APPROVED evaluation is terminal; flagging on top must be rejected")
}

// --------------------------------------------------------------------------
// v7 genesis round-trip: AgentDepositBalances + dispute target_role + the
// derived (subject, role) -> proof index and active-dispute presence index.
// --------------------------------------------------------------------------

func (s *KeeperTestSuite) TestGenesisRoundTrip_V7State() {
	s.SetupTest()

	collID := "coll-genesis"
	adminAcc := apptesting.RandomAccountAddress()
	submitterAcc := apptesting.RandomAccountAddress()
	evaluatorAcc := apptesting.RandomAccountAddress()
	depositorAcc := apptesting.RandomAccountAddress()

	// Seed: collection, claim with terminal evaluation, OPEN dispute against
	// submitter, AWARDED dispute against evaluator, one performance-deposit
	// balance.
	s.App.ClaimsKeeper.SetCollection(s.Ctx, types.Collection{
		Id:       collID,
		Entity:   "did:ixo:entity-gen",
		Admin:    adminAcc.String(),
		Protocol: "did:ixo:protocol-gen",
		State:    types.CollectionState_open,
	})
	s.App.ClaimsKeeper.SetClaim(s.Ctx, types.Claim{
		ClaimId:      "claim-gen",
		CollectionId: collID,
		AgentAddress: submitterAcc.String(),
		AgentDid:     "did:ixo:claim-author",
		Evaluation: &types.Evaluation{
			ClaimId:      "claim-gen",
			CollectionId: collID,
			AgentAddress: evaluatorAcc.String(),
			Status:       types.EvaluationStatus_approved,
		},
	})

	openDispute := types.Dispute{
		SubjectId:  "claim-gen",
		Type:       1,
		Status:     types.DisputeStatus_dispute_open,
		Data:       &types.DisputeData{Proof: "proof-open-gen", Uri: "ipfs://uri-open-gen"},
		TargetRole: types.DisputeTargetRole_target_submitter,
	}
	awardedDispute := types.Dispute{
		SubjectId:  "claim-gen",
		Type:       1,
		Status:     types.DisputeStatus_dispute_awarded,
		Data:       &types.DisputeData{Proof: "proof-award-gen", Uri: "ipfs://uri-award-gen"},
		TargetRole: types.DisputeTargetRole_target_evaluator,
	}
	s.App.ClaimsKeeper.SetDispute(s.Ctx, openDispute)
	s.App.ClaimsKeeper.SetDispute(s.Ctx, awardedDispute)

	s.App.ClaimsKeeper.SetAgentDepositBalance(s.Ctx, types.AgentDepositBalance{
		CollectionId: collID,
		AgentAddress: depositorAcc.String(),
		Amount:       sdk.NewCoins(sdk.NewCoin("uixo", math.NewInt(777))),
	})

	// Export.
	export := s.App.ClaimsKeeper.ExportGenesis(s.Ctx)
	s.Require().Len(export.Collections, 1, "exported genesis must include the collection")
	s.Require().Len(export.Claims, 1, "exported genesis must include the claim")
	s.Require().Len(export.Disputes, 2, "exported genesis must include both disputes")
	s.Require().Len(export.AgentDepositBalances, 1,
		"exported genesis must include the agent deposit balance (v7 addition)")
	s.Require().Equal("777uixo", export.AgentDepositBalances[0].Amount.String())

	// Confirm both target_roles are preserved in export — the v7
	// MsgDisputeClaim wire path threads target_role into Dispute, so
	// round-tripping the role is the load-bearing v7 genesis invariant.
	var sawSubmitter, sawEvaluator bool
	for _, d := range export.Disputes {
		switch d.TargetRole {
		case types.DisputeTargetRole_target_submitter:
			sawSubmitter = true
		case types.DisputeTargetRole_target_evaluator:
			sawEvaluator = true
		}
	}
	s.Require().True(sawSubmitter, "export must preserve target_role=SUBMITTER")
	s.Require().True(sawEvaluator, "export must preserve target_role=EVALUATOR")

	// Now re-init a clean state from the exported genesis.
	s.SetupTest()
	s.App.ClaimsKeeper.InitGenesis(s.Ctx, *export)

	// Collection + claim + balances round-trip.
	got, err := s.App.ClaimsKeeper.GetCollection(s.Ctx, collID)
	s.Require().NoError(err)
	s.Require().Equal(adminAcc.String(), got.Admin)

	gotClaim, err := s.App.ClaimsKeeper.GetClaim(s.Ctx, "claim-gen")
	s.Require().NoError(err)
	s.Require().NotNil(gotClaim.Evaluation)
	s.Require().Equal(types.EvaluationStatus_approved, gotClaim.Evaluation.Status)

	bal, err := s.App.ClaimsKeeper.GetAgentDepositBalance(s.Ctx, collID, depositorAcc.String())
	s.Require().NoError(err)
	s.Require().Equal("777uixo", bal.Amount.String())

	// Derived index: subject->proof for the SUBMITTER-targeted dispute.
	// InitGenesis rebuilds this from the dispute records.
	proof, ok := s.App.ClaimsKeeper.GetDisputeProofForSubject(s.Ctx, "claim-gen",
		types.DisputeTargetRole_target_submitter)
	s.Require().True(ok)
	s.Require().Equal("proof-open-gen", proof)

	// Derived index for the EVALUATOR-targeted AWARDED dispute.
	proof, ok = s.App.ClaimsKeeper.GetDisputeProofForSubject(s.Ctx, "claim-gen",
		types.DisputeTargetRole_target_evaluator)
	s.Require().True(ok)
	s.Require().Equal("proof-award-gen", proof)

	// Active-dispute presence index for the submitter (OPEN status; rebuilt).
	s.Require().True(
		s.App.ClaimsKeeper.HasActiveDisputeAgainstAgent(s.Ctx, collID, submitterAcc.String()),
		"OPEN-against-submitter dispute must rebuild the per-agent presence index")

	// AWARDED disputes don't appear in the presence index (only OPEN ones
	// do — the index gates "can the agent still submit?").
	s.Require().False(
		s.App.ClaimsKeeper.HasActiveDisputeAgainstAgent(s.Ctx, collID, evaluatorAcc.String()),
		"AWARDED-against-evaluator dispute must NOT rebuild presence (only OPEN does)")
}

// Legacy (target_role=UNSPECIFIED) disputes must round-trip — they're the
// migrated v3→v4 records — without being added to the subject index (they
// can't gate future filings under v7 semantics).
func (s *KeeperTestSuite) TestGenesisRoundTrip_LegacyDisputeIgnoresIndex() {
	s.SetupTest()

	legacy := types.Dispute{
		SubjectId:  "claim-legacy",
		Type:       1,
		Status:     types.DisputeStatus_dispute_dismissed,
		Data:       &types.DisputeData{Proof: "proof-legacy", Uri: "ipfs://uri-legacy"},
		TargetRole: types.DisputeTargetRole_target_unspecified,
	}
	s.App.ClaimsKeeper.SetDispute(s.Ctx, legacy)

	export := s.App.ClaimsKeeper.ExportGenesis(s.Ctx)
	s.SetupTest()
	s.App.ClaimsKeeper.InitGenesis(s.Ctx, *export)

	// Legacy disputes survive in storage so they're queryable / auditable.
	got, err := s.App.ClaimsKeeper.GetDispute(s.Ctx, "proof-legacy")
	s.Require().NoError(err)
	s.Require().Equal(types.DisputeTargetRole_target_unspecified, got.TargetRole)

	// But the subject index is NOT populated (target_role is UNSPECIFIED),
	// so a new dispute against the same subject CAN still be filed under
	// v7 rules.
	s.Require().NoError(
		s.App.ClaimsKeeper.CanFileNewDisputeForSubject(s.Ctx, "claim-legacy",
			types.DisputeTargetRole_target_submitter))
}

// --------------------------------------------------------------------------
// Active-dispute presence index — used by SubmitClaim / EvaluateClaim /
// WithdrawPerformanceDeposit to gate the actor.
// --------------------------------------------------------------------------

func (s *KeeperTestSuite) TestActiveDisputeIndex_PresenceAndScan() {
	s.SetupTest()
	const (
		coll  = "coll-active"
		agent = "ixo1agent-a"
	)

	// No entries → not present.
	s.Require().False(s.App.ClaimsKeeper.HasActiveDisputeAgainstAgent(s.Ctx, coll, agent))

	// Add three open disputes against the same agent (three subjects);
	// presence flag should fire on a single prefix lookup.
	s.App.ClaimsKeeper.SetActiveDispute(s.Ctx, coll, agent, "subj-1")
	s.App.ClaimsKeeper.SetActiveDispute(s.Ctx, coll, agent, "subj-2")
	s.App.ClaimsKeeper.SetActiveDispute(s.Ctx, coll, agent, "subj-3")
	s.Require().True(s.App.ClaimsKeeper.HasActiveDisputeAgainstAgent(s.Ctx, coll, agent))

	// Iterator must return all three subjects.
	subjects := s.App.ClaimsKeeper.GetActiveDisputeSubjectsForAgent(s.Ctx, coll, agent)
	s.Require().ElementsMatch([]string{"subj-1", "subj-2", "subj-3"}, subjects)

	// Different agent on same collection: still empty.
	s.Require().False(s.App.ClaimsKeeper.HasActiveDisputeAgainstAgent(s.Ctx, coll, "ixo1agent-b"))

	// Different collection: also empty.
	s.Require().False(s.App.ClaimsKeeper.HasActiveDisputeAgainstAgent(s.Ctx, "other-coll", agent))

	// Remove two of three — presence still fires.
	s.App.ClaimsKeeper.RemoveActiveDispute(s.Ctx, coll, agent, "subj-1")
	s.App.ClaimsKeeper.RemoveActiveDispute(s.Ctx, coll, agent, "subj-2")
	s.Require().True(s.App.ClaimsKeeper.HasActiveDisputeAgainstAgent(s.Ctx, coll, agent))

	// Remove the last → presence clears.
	s.App.ClaimsKeeper.RemoveActiveDispute(s.Ctx, coll, agent, "subj-3")
	s.Require().False(s.App.ClaimsKeeper.HasActiveDisputeAgainstAgent(s.Ctx, coll, agent))
}

// --------------------------------------------------------------------------
// HasAgentMetDepositRequirement — deposit-gate comparison logic
// --------------------------------------------------------------------------

func (s *KeeperTestSuite) TestHasAgentMetDepositRequirement() {
	s.SetupTest()
	const coll = "coll-req"
	agent := apptesting.RandomAccountAddress().String()

	// No required → always allowed (escape hatch).
	s.Require().True(s.App.ClaimsKeeper.HasAgentMetDepositRequirement(s.Ctx, coll, agent, sdk.Coins{}))

	// Required but no balance → blocked.
	required := sdk.NewCoins(sdk.NewCoin("uixo", math.NewInt(100)))
	s.Require().False(s.App.ClaimsKeeper.HasAgentMetDepositRequirement(s.Ctx, coll, agent, required))

	// Add a balance just below requirement → still blocked.
	s.App.ClaimsKeeper.SetAgentDepositBalance(s.Ctx, types.AgentDepositBalance{
		CollectionId: coll,
		AgentAddress: agent,
		Amount:       sdk.NewCoins(sdk.NewCoin("uixo", math.NewInt(99))),
	})
	s.Require().False(s.App.ClaimsKeeper.HasAgentMetDepositRequirement(s.Ctx, coll, agent, required))

	// Balance == requirement → allowed (≥).
	s.App.ClaimsKeeper.SetAgentDepositBalance(s.Ctx, types.AgentDepositBalance{
		CollectionId: coll,
		AgentAddress: agent,
		Amount:       sdk.NewCoins(sdk.NewCoin("uixo", math.NewInt(100))),
	})
	s.Require().True(s.App.ClaimsKeeper.HasAgentMetDepositRequirement(s.Ctx, coll, agent, required))

	// Balance well above requirement → allowed.
	s.App.ClaimsKeeper.SetAgentDepositBalance(s.Ctx, types.AgentDepositBalance{
		CollectionId: coll,
		AgentAddress: agent,
		Amount:       sdk.NewCoins(sdk.NewCoin("uixo", math.NewInt(1_000_000))),
	})
	s.Require().True(s.App.ClaimsKeeper.HasAgentMetDepositRequirement(s.Ctx, coll, agent, required))

	// Multi-denom requirement: must meet ALL denoms.
	multi := sdk.NewCoins(
		sdk.NewCoin("uixo", math.NewInt(10)),
		sdk.NewCoin("uatom", math.NewInt(20)),
	)
	s.App.ClaimsKeeper.SetAgentDepositBalance(s.Ctx, types.AgentDepositBalance{
		CollectionId: coll,
		AgentAddress: agent,
		Amount:       sdk.NewCoins(sdk.NewCoin("uixo", math.NewInt(1_000_000))), // only uixo present
	})
	s.Require().False(s.App.ClaimsKeeper.HasAgentMetDepositRequirement(s.Ctx, coll, agent, multi),
		"missing one of the required denoms → must reject")

	s.App.ClaimsKeeper.SetAgentDepositBalance(s.Ctx, types.AgentDepositBalance{
		CollectionId: coll,
		AgentAddress: agent,
		Amount: sdk.NewCoins(
			sdk.NewCoin("uixo", math.NewInt(1_000_000)),
			sdk.NewCoin("uatom", math.NewInt(20)),
		),
	})
	s.Require().True(s.App.ClaimsKeeper.HasAgentMetDepositRequirement(s.Ctx, coll, agent, multi))
}

// --------------------------------------------------------------------------
// AgentDepositBalance — query + iteration lifecycle
// --------------------------------------------------------------------------

func (s *KeeperTestSuite) TestAgentDepositBalance_OrZero() {
	s.SetupTest()
	const coll = "coll-zero"
	agent := apptesting.RandomAccountAddress().String()

	// Not set → OrZero returns a zero-valued entry without erroring.
	bal := s.App.ClaimsKeeper.GetAgentDepositBalanceOrZero(s.Ctx, coll, agent)
	s.Require().Equal(coll, bal.CollectionId)
	s.Require().Equal(agent, bal.AgentAddress)
	s.Require().True(bal.Amount.IsZero())

	// GetAgentDepositBalance on unset entry errors.
	_, err := s.App.ClaimsKeeper.GetAgentDepositBalance(s.Ctx, coll, agent)
	s.Require().Error(err)

	// Set then read back.
	s.App.ClaimsKeeper.SetAgentDepositBalance(s.Ctx, types.AgentDepositBalance{
		CollectionId: coll,
		AgentAddress: agent,
		Amount:       sdk.NewCoins(sdk.NewCoin("uixo", math.NewInt(42))),
	})
	got, err := s.App.ClaimsKeeper.GetAgentDepositBalance(s.Ctx, coll, agent)
	s.Require().NoError(err)
	s.Require().Equal("42uixo", got.Amount.String())
}

// LookupAdjudicator returns the entry that matches by DID, ignoring the
// percentage. Used by the adjudication handler to fetch the per-DID
// reward_percentage.
func (s *KeeperTestSuite) TestLookupAdjudicator() {
	collection := types.Collection{
		Adjudicators: []*types.AdjudicationDid{
			adjudicatorDid("did:ixo:a", 10),
			adjudicatorDid("did:ixo:b", 25),
		},
	}
	got, ok := claimskeeper.LookupAdjudicator(collection, "did:ixo:b")
	s.Require().True(ok)
	s.Require().Equal("did:ixo:b", got.Did)
	s.Require().Equal(int64(25), got.RewardPercentage.TruncateInt64())

	_, ok = claimskeeper.LookupAdjudicator(collection, "did:ixo:not-listed")
	s.Require().False(ok)
}
