package keeper_test

import (
	"time"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/ixofoundation/ixo-blockchain/v7/app/apptesting"
	claimskeeper "github.com/ixofoundation/ixo-blockchain/v7/x/claims/keeper"
	"github.com/ixofoundation/ixo-blockchain/v7/x/claims/types"
)

// These tests cover the admin-driven, non-wasm collection management flows:
// UpdateCollectionState / Dates / Intents, SetCollectionMembers,
// RemoveCollectionMembers. The wasm-payment-settlement flows (SubmitClaim,
// EvaluateClaim, DisputeClaim, WithdrawPayment, ClaimIntent) are covered in
// tests/interchaintest/ where a real cw20/cw1155 wasm artefact is available.

func (s *KeeperTestSuite) seedAdminCollection(id string) (string, types.Collection) {
	admin := apptesting.RandomAccountAddress().String()
	c := types.Collection{
		Id:       id,
		Entity:   "did:ixo:entity-claim",
		Admin:    admin,
		Protocol: "did:ixo:protocol-claim",
		State:    types.CollectionState_open,
		Intents:  types.CollectionIntentOptions_allow,
	}
	s.App.ClaimsKeeper.SetCollection(s.Ctx, c)
	return admin, c
}

func (s *KeeperTestSuite) TestMsgUpdateCollectionState() {
	s.SetupTest()
	admin, _ := s.seedAdminCollection("upd-state")
	ms := claimskeeper.NewMsgServerImpl(&s.App.ClaimsKeeper)

	_, err := ms.UpdateCollectionState(s.Ctx, &types.MsgUpdateCollectionState{
		CollectionId: "upd-state",
		State:        types.CollectionState_paused,
		AdminAddress: admin,
	})
	s.Require().NoError(err)
	got, _ := s.App.ClaimsKeeper.GetCollection(s.Ctx, "upd-state")
	s.Require().Equal(types.CollectionState_paused, got.State)
}

func (s *KeeperTestSuite) TestMsgUpdateCollectionState_Unauthorized() {
	s.SetupTest()
	_, _ = s.seedAdminCollection("upd-state-auth")
	ms := claimskeeper.NewMsgServerImpl(&s.App.ClaimsKeeper)
	stranger := apptesting.RandomAccountAddress().String()

	_, err := ms.UpdateCollectionState(s.Ctx, &types.MsgUpdateCollectionState{
		CollectionId: "upd-state-auth",
		State:        types.CollectionState_paused,
		AdminAddress: stranger,
	})
	s.Require().ErrorContains(err, "unauthorized")
}

func (s *KeeperTestSuite) TestMsgUpdateCollectionDates() {
	s.SetupTest()
	admin, _ := s.seedAdminCollection("upd-dates")
	ms := claimskeeper.NewMsgServerImpl(&s.App.ClaimsKeeper)

	start := s.Ctx.BlockTime()
	end := start.Add(48 * time.Hour)
	_, err := ms.UpdateCollectionDates(s.Ctx, &types.MsgUpdateCollectionDates{
		CollectionId: "upd-dates",
		StartDate:    &start,
		EndDate:      &end,
		AdminAddress: admin,
	})
	s.Require().NoError(err)
	got, _ := s.App.ClaimsKeeper.GetCollection(s.Ctx, "upd-dates")
	s.Require().NotNil(got.StartDate)
	s.Require().NotNil(got.EndDate)
}

func (s *KeeperTestSuite) TestMsgUpdateCollectionIntents() {
	s.SetupTest()
	admin, _ := s.seedAdminCollection("upd-intents")
	ms := claimskeeper.NewMsgServerImpl(&s.App.ClaimsKeeper)

	_, err := ms.UpdateCollectionIntents(s.Ctx, &types.MsgUpdateCollectionIntents{
		CollectionId: "upd-intents",
		Intents:      types.CollectionIntentOptions_required,
		AdminAddress: admin,
	})
	s.Require().NoError(err)
	got, _ := s.App.ClaimsKeeper.GetCollection(s.Ctx, "upd-intents")
	s.Require().Equal(types.CollectionIntentOptions_required, got.Intents)
}

func (s *KeeperTestSuite) TestMsgSetCollectionMembers_AddNew() {
	s.SetupTest()
	admin, _ := s.seedAdminCollection("members")
	ms := claimskeeper.NewMsgServerImpl(&s.App.ClaimsKeeper)

	member := apptesting.RandomAccountAddress().String()
	period := 30 * 24 * time.Hour
	_, err := ms.SetCollectionMembers(s.Ctx, &types.MsgSetCollectionMembers{
		CollectionId: "members",
		AdminAddress: admin,
		Members: []*types.CollectionMemberInput{
			{
				MemberAddress:    member,
				Period:           period,
				PeriodSpendLimit: sdk.NewCoins(sdk.NewCoin("uixo", math.NewInt(1_000))),
			},
		},
	})
	s.Require().NoError(err)

	budget, err := s.App.ClaimsKeeper.GetMemberBudget(s.Ctx, "members", member)
	s.Require().NoError(err)
	s.Require().Equal(period, budget.Period)
	s.Require().Equal(member, budget.MemberAddress)
	s.AssertEventEmitted(s.Ctx, "ixo.claims.v1beta1.MemberBudgetCreatedEvent", 1)
}

func (s *KeeperTestSuite) TestMsgSetCollectionMembers_UpdateExisting() {
	s.SetupTest()
	admin, _ := s.seedAdminCollection("members-update")
	ms := claimskeeper.NewMsgServerImpl(&s.App.ClaimsKeeper)

	member := apptesting.RandomAccountAddress().String()
	period := 30 * 24 * time.Hour
	// First call: creates the member budget.
	_, err := ms.SetCollectionMembers(s.Ctx, &types.MsgSetCollectionMembers{
		CollectionId: "members-update",
		AdminAddress: admin,
		Members: []*types.CollectionMemberInput{
			{MemberAddress: member, Period: period, PeriodSpendLimit: sdk.NewCoins(sdk.NewCoin("uixo", math.NewInt(1_000)))},
		},
	})
	s.Require().NoError(err)

	// Second call with a different limit — must update without error.
	_, err = ms.SetCollectionMembers(s.Ctx, &types.MsgSetCollectionMembers{
		CollectionId: "members-update",
		AdminAddress: admin,
		Members: []*types.CollectionMemberInput{
			{MemberAddress: member, Period: period, PeriodSpendLimit: sdk.NewCoins(sdk.NewCoin("uixo", math.NewInt(5_000)))},
		},
	})
	s.Require().NoError(err)

	got, err := s.App.ClaimsKeeper.GetMemberBudget(s.Ctx, "members-update", member)
	s.Require().NoError(err)
	s.Require().Equal(int64(5_000), got.PeriodSpendLimit.AmountOf("uixo").Int64())
}

func (s *KeeperTestSuite) TestMsgRemoveCollectionMembers() {
	s.SetupTest()
	admin, _ := s.seedAdminCollection("members-remove")
	ms := claimskeeper.NewMsgServerImpl(&s.App.ClaimsKeeper)

	member := apptesting.RandomAccountAddress().String()
	_, err := ms.SetCollectionMembers(s.Ctx, &types.MsgSetCollectionMembers{
		CollectionId: "members-remove",
		AdminAddress: admin,
		Members: []*types.CollectionMemberInput{
			{MemberAddress: member, Period: 24 * time.Hour, PeriodSpendLimit: sdk.NewCoins(sdk.NewCoin("uixo", math.NewInt(100)))},
		},
	})
	s.Require().NoError(err)

	_, err = ms.RemoveCollectionMembers(s.Ctx, &types.MsgRemoveCollectionMembers{
		CollectionId:    "members-remove",
		AdminAddress:    admin,
		MemberAddresses: []string{member},
	})
	s.Require().NoError(err)

	_, err = s.App.ClaimsKeeper.GetMemberBudget(s.Ctx, "members-remove", member)
	s.Require().Error(err, "removed member must not be retrievable")
}
