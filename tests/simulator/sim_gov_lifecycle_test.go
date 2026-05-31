// In-process governance lifecycle test. Submits a real gov proposal
// directly to the gov keeper, votes via the validator, advances time
// to elapse the voting period, and verifies the proposal executed.
// This is the fast counterpart to L3's
// `module_names_test.go::TestIxoNames_GovCreateAndQuery` — boots in
// ~1s instead of ~60s, so it's worth running on every push.
//
// Run with:
//
//	go test -run TestIxoSim_GovProposalLifecycleInProcess ./tests/simulator/...
package simulator

import (
	"testing"
	"time"

	abci "github.com/cometbft/cometbft/abci/types"
	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/gov"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	govv1 "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	"github.com/stretchr/testify/require"

	"github.com/ixofoundation/ixo-blockchain/v7/app"
	namestypes "github.com/ixofoundation/ixo-blockchain/v7/x/names/types"
)

// TestIxoSim_GovProposalLifecycleInProcess submits a gov v1 proposal
// containing a names.MsgCreateNamespace, votes YES from the validator,
// advances time past the voting period, and asserts the namespace was
// created.
//
// This is the same flow as the L3 names-gov test, but driven through
// the keeper directly. It catches gov-keeper regressions without
// paying the Docker startup cost.
func TestIxoSim_GovProposalLifecycleInProcess(t *testing.T) {
	a := app.Setup(false)
	ctx := a.NewContextLegacy(false, cmtproto.Header{Time: time.Now()})

	govAcct := a.AccountKeeper.GetModuleAddress(govtypes.ModuleName).String()

	// Build the embedded MsgCreateNamespace payload.
	const namespaceName = "in-process-test"
	embedded := &namestypes.MsgCreateNamespace{
		Authority: govAcct,
		Namespace: &namestypes.Namespace{
			Name:              namespaceName,
			Description:       "registered via in-process gov lifecycle test",
			MinLength:         3,
			MaxLength:         32,
			AllowSelfRegister: true,
		},
	}

	// Mint funds to a fresh proposer so we can satisfy the deposit
	// without invoking gov's missing Minter permission.
	govParams, err := a.GovKeeper.Params.Get(ctx)
	require.NoError(t, err)
	deposit := sdk.NewCoins(govParams.MinDeposit...)

	proposer := sdk.AccAddress([]byte("proposer-test-2cosmos"))[:20]
	require.NoError(t, a.BankKeeper.MintCoins(ctx, "mint", deposit))
	require.NoError(t, a.BankKeeper.SendCoinsFromModuleToAccount(ctx, "mint", proposer, deposit))

	// Submit a v1 proposal with the embedded message.
	proposalMsgs := []sdk.Msg{embedded}
	proposal, err := a.GovKeeper.SubmitProposal(ctx, proposalMsgs, "",
		"in-process gov test", "submitted by TestIxoSim_GovProposalLifecycleInProcess",
		proposer, false)
	require.NoError(t, err, "SubmitProposal")

	// Add the deposit to activate the voting period.
	_, err = a.GovKeeper.AddDeposit(ctx, proposal.Id, proposer, deposit)
	require.NoError(t, err)

	// Vote YES from the validator.
	vals, err := a.StakingKeeper.GetAllValidators(ctx)
	require.NoError(t, err)
	require.NotEmpty(t, vals, "test app must have at least one validator")
	valOperAddr, err := sdk.ValAddressFromBech32(vals[0].OperatorAddress)
	require.NoError(t, err)
	voterAddr := sdk.AccAddress(valOperAddr)

	require.NoError(t, a.GovKeeper.AddVote(ctx, proposal.Id, voterAddr,
		govv1.NewNonSplitVoteOption(govv1.OptionYes), ""))

	// Advance time past the voting period and run gov EndBlocker so the
	// proposal is tallied and (if passed) its messages are executed.
	votingPeriod := *govParams.VotingPeriod
	ctx = ctx.WithBlockTime(ctx.BlockTime().Add(votingPeriod).Add(time.Second))
	require.NoError(t, gov.EndBlocker(ctx, a.GovKeeper))

	// The namespace must now exist.
	ns, found := a.NamesKeeper.GetNamespace(ctx, namespaceName)
	require.True(t, found,
		"namespace %q must exist after the gov proposal passes; gov flow regression?",
		namespaceName)
	require.Equal(t, namespaceName, ns.Name)
	require.Equal(t, uint32(3), ns.MinLength)
	require.Equal(t, uint32(32), ns.MaxLength)

	// And the proposal status must be PASSED.
	got, err := a.GovKeeper.Proposals.Get(ctx, proposal.Id)
	require.NoError(t, err)
	require.Equal(t, govv1.StatusPassed, got.Status,
		"proposal must reach PASSED after voting period; got %s", got.Status)

	// silence import-unused for abci (kept for tests that drive
	// FinalizeBlock directly).
	_ = abci.ResponseFinalizeBlock{}
}
