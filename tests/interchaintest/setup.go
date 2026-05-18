//go:build interchaintest

// Package interchaintest provides Docker-based end-to-end tests for the ixo
// blockchain. Build-tag gated; run with:
//
//	make test-interchaintest
//
// or directly:
//
//	cd tests/interchaintest && go test -tags interchaintest -timeout 60m ./...
package interchaintest

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"testing"

	sdkmath "cosmossdk.io/math"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	interchaintest "github.com/strangelove-ventures/interchaintest/v8"
	"github.com/strangelove-ventures/interchaintest/v8/chain/cosmos"
	"github.com/strangelove-ventures/interchaintest/v8/ibc"
	"github.com/strangelove-ventures/interchaintest/v8/testutil"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"
)

// Chain configuration constants — matched to the IxoApp's bech32 prefixes
// and native token (see lib/ixo/types.go and app/keepers).
const (
	IxoBech32       = "ixo"
	IxoNativeDenom  = "uixo"
	IxoChainName    = "ixo"
	IxoBinaryName   = "ixod"
	IxoMinGasPrices = "0.0025uixo"

	// VotingPeriod and MaxDepositPeriod are intentionally short so gov
	// flows complete within a single test rather than blocking the suite
	// on the chain's mainnet voting cadence.
	VotingPeriod     = "8s"
	MaxDepositPeriod = "8s"

	// ixoImageEnv is the env var holding the locally-built ixod Docker image
	// reference. Defaults to ixofoundation/ixo-blockchain:local; override
	// when testing a PR build.
	ixoImageEnv = "IXO_IMAGE"
)

// IxoChainSpec returns an interchaintest *ibc.ChainConfig for ixod with
// genesis modifications that make the chain test-friendly:
//   - short gov voting/deposit periods (8s) so proposals settle quickly
//   - wasm code-upload + instantiate permissions opened to everybody so
//     tests don't need to wrap every Store/Instantiate in a gov proposal
//   - min deposit denominated in uixo
func IxoChainSpec() *ibc.ChainConfig {
	image := os.Getenv(ixoImageEnv)
	if image == "" {
		image = "ixofoundation/ixo-blockchain:local"
	}
	// Repository must be tag-less; the Version field below supplies the tag.
	// If the env var came in pre-tagged (e.g. `repo:tag`), split it.
	repo, tag := image, "local"
	if idx := strings.LastIndexByte(image, ':'); idx > 0 && !strings.Contains(image[idx:], "/") {
		repo, tag = image[:idx], image[idx+1:]
	}
	return &ibc.ChainConfig{
		Type:    "cosmos",
		Name:    IxoChainName,
		ChainID: "ixo-localnet-1",
		Images: []ibc.DockerImage{
			{Repository: repo, Version: tag, UIDGID: "1025:1025"},
		},
		Bin:            IxoBinaryName,
		Bech32Prefix:   IxoBech32,
		Denom:          IxoNativeDenom,
		CoinType:       "118",
		GasPrices: IxoMinGasPrices,
		// GasAdjustment is what interchaintest's helpers (VoteOnProposal,
		// SendIBCTransfer, etc.) tack on to `--gas auto` when they don't
		// expose a per-call override. The custom ixo modules are
		// gas-heavy on the verification path (IID auth, smart-account
		// authenticator lookup), so 2.5 keeps txes from going OOG.
		GasAdjustment: 2.5,
		TrustingPeriod: "112h",
		NoHostMount:    false,
		ModifyGenesis:  cosmos.ModifyGenesis(defaultGenesisKV()),
	}
}

// defaultGenesisKV returns the genesis-key overrides that make the chain
// usable in a test scenario.
func defaultGenesisKV() []cosmos.GenesisKV {
	return []cosmos.GenesisKV{
		{Key: "app_state.gov.params.voting_period", Value: VotingPeriod},
		{Key: "app_state.gov.params.max_deposit_period", Value: MaxDepositPeriod},
		{Key: "app_state.gov.params.expedited_voting_period", Value: VotingPeriod},
		{Key: "app_state.gov.params.min_deposit.0.denom", Value: IxoNativeDenom},
		{Key: "app_state.gov.params.min_deposit.0.amount", Value: "1000000"},
		{Key: "app_state.gov.params.expedited_min_deposit.0.denom", Value: IxoNativeDenom},
		{Key: "app_state.gov.params.expedited_min_deposit.0.amount", Value: "5000000"},

		// Wasm: open code-upload + instantiate so tests can deploy contracts
		// directly without a gov proposal. AccessTypeEverybody is the
		// permissive setting; tests that want to exercise the gov-only path
		// should override these in their own ChainConfig.
		{Key: "app_state.wasm.params.code_upload_access.permission",
			Value: wasmtypes.AccessTypeEverybody.String()},
		{Key: "app_state.wasm.params.instantiate_default_permission",
			Value: wasmtypes.AccessTypeEverybody.String()},
	}
}

// DefaultUserFunds is the per-test-account balance handed out by
// interchaintest.GetAndFundTestUsers. Generous so cosmos-sdk fees and
// per-msg gas don't drain the account mid-test.
var DefaultUserFunds = sdkmath.NewInt(10_000_000_000)

// ContractsDir is the absolute path to tests/interchaintest/contracts/,
// resolved relative to setup.go regardless of the test's working directory.
func ContractsDir() string {
	_, file, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(file), "contracts")
}

// ContractPath returns the absolute path to a wasm artefact in the
// contracts/ directory, e.g. ContractPath("cw721.wasm").
func ContractPath(name string) string { return filepath.Join(ContractsDir(), name) }

// ptr takes the address of a value literal — interchaintest's ChainSpec
// wants *int for validator counts.
func ptr[T any](v T) *T { return &v }

// SetupIxoChain spins up a single-validator ixo chain with the default
// genesis modifications, builds the interchain, funds `numUsers` test
// users, and returns the live chain plus those users.
func SetupIxoChain(t *testing.T, numUsers int) (*cosmos.CosmosChain, []ibc.Wallet, context.Context) {
	t.Helper()
	ctx := context.Background()

	cf := interchaintest.NewBuiltinChainFactory(zaptest.NewLogger(t), []*interchaintest.ChainSpec{
		{
			Name:          IxoChainName,
			Version:       "local",
			ChainConfig:   *IxoChainSpec(),
			NumValidators: ptr(1),
			NumFullNodes:  ptr(0),
		},
	})

	chains, err := cf.Chains(t.Name())
	require.NoError(t, err)
	chain := chains[0].(*cosmos.CosmosChain)

	// DockerSetup must run before Build — it creates the docker client
	// and a per-test docker network that interchaintest hands to the
	// chain so it can mount volumes and expose ports. Without it the
	// CosmosChain's docker client stays nil and NewChainNode panics
	// inside VolumeCreate.
	client, network := interchaintest.DockerSetup(t)

	ic := interchaintest.NewInterchain().AddChain(chain)
	t.Cleanup(func() { _ = ic.Close() })

	require.NoError(t, ic.Build(ctx, nil, interchaintest.InterchainBuildOptions{
		TestName:         t.Name(),
		Client:           client,
		NetworkID:        network,
		SkipPathCreation: true,
	}))

	// GetAndFundTestUsers returns one wallet per *chain*, not N wallets per
	// chain. To fund N distinct accounts on the same chain, call it once per
	// account with a unique key-name prefix so the keyring entries don't
	// collide.
	users := make([]ibc.Wallet, 0, numUsers)
	for i := 0; i < numUsers; i++ {
		prefix := fmt.Sprintf("%s-u%d", t.Name(), i)
		got := interchaintest.GetAndFundTestUsers(t, ctx, prefix, DefaultUserFunds, chain)
		require.Len(t, got, 1, "GetAndFundTestUsers must return one wallet per call")
		users = append(users, got[0])
	}

	return chain, users, ctx
}

// UploadContract uploads a wasm artefact from tests/interchaintest/contracts/
// and returns its assigned code id.
func UploadContract(t *testing.T, ctx context.Context, chain *cosmos.CosmosChain,
	caller ibc.Wallet, filename string,
) string {
	t.Helper()
	codeID, err := chain.StoreContract(ctx, caller.KeyName(), ContractPath(filename))
	require.NoError(t, err, "uploading %s", filename)
	require.NotEmpty(t, codeID, "stored code id for %s must be non-empty", filename)
	return codeID
}

// UploadAllContracts stores every bundled wasm artefact in the order the
// SDK's Proposals.instantiateModulesProposals expects (cw721, ixo1155,
// cw20_base, cw721_base, cw4_group). Returns a map {filename → codeID}.
// On a freshly-bootstrapped chain the resulting code IDs are 1..5 in that
// order — matching the SDK's contract.constants.ts.
func UploadAllContracts(t *testing.T, ctx context.Context, chain *cosmos.CosmosChain,
	caller ibc.Wallet,
) map[string]string {
	t.Helper()
	out := make(map[string]string, 5)
	for _, name := range []string{
		"cw721.wasm",
		"ixo1155.wasm",
		"cw20_base.wasm",
		"cw721_base.wasm",
		"cw4_group.wasm",
	} {
		out[name] = UploadContract(t, ctx, chain, caller, name)
	}
	return out
}

// WaitBlocks is a small wrapper around interchaintest's WaitForBlocks.
func WaitBlocks(t *testing.T, ctx context.Context, chain *cosmos.CosmosChain, n int) {
	t.Helper()
	require.NoError(t, testutil.WaitForBlocks(ctx, n, chain))
}

// SubmitGovProposalAndPass submits a JSON gov proposal file, votes yes from
// the validator, and waits for the voting period to elapse. Returns the
// proposal id.
func SubmitGovProposalAndPass(t *testing.T, ctx context.Context, chain *cosmos.CosmosChain,
	proposer ibc.Wallet, proposalJSON string,
) uint64 {
	t.Helper()

	// Push the proposal JSON into the validator container's home dir.
	const proposalRel = "proposal.json"
	require.NoError(t, chain.GetNode().WriteFile(ctx, []byte(proposalJSON), proposalRel))

	out, err := chain.GetNode().ExecTx(ctx, proposer.KeyName(),
		"gov", "submit-proposal", chain.GetNode().HomeDir()+"/"+proposalRel,
		"--gas", "auto", "--gas-adjustment", "2.0")
	require.NoError(t, err, "submit-proposal: %s", out)

	// Find the most recent proposal id by listing proposals.
	stdout, _, err := chain.GetNode().ExecQuery(ctx, "gov", "proposals", "--output", "json")
	require.NoError(t, err)
	proposalID := lastProposalID(t, stdout)

	// Vote with EVERY validator's signing key — fresh test users have funded
	// balances but no bonded stake, so a yes vote from `proposer` alone
	// can't reach quorum + threshold. interchaintest's
	// VoteOnProposalAllValidators iterates every validator node's `valKey`.
	require.NoError(t, chain.VoteOnProposalAllValidators(ctx, proposalID, "yes"))

	// Wait for voting period (8s) + a little slack.
	WaitBlocks(t, ctx, chain, 6)

	// Verify the proposal actually passed.
	require.Equal(t, "PROPOSAL_STATUS_PASSED", queryProposalStatus(t, ctx, chain, proposalID),
		"proposal %d did not pass", proposalID)

	return proposalID
}

// queryProposalStatus returns the gov v1 status string for a single
// proposal id. Robust to whether the SDK CLI wraps the proposal in a
// top-level `proposal` field or returns it flat — both shapes have been
// observed across cosmos-sdk minor versions.
func queryProposalStatus(t *testing.T, ctx context.Context, chain *cosmos.CosmosChain, id uint64) string {
	t.Helper()
	out, _, err := chain.GetNode().ExecQuery(ctx,
		"gov", "proposal", fmt.Sprintf("%d", id), "--output", "json")
	require.NoError(t, err)

	// Try wrapped first.
	var wrapped struct {
		Proposal struct {
			Status string `json:"status"`
		} `json:"proposal"`
	}
	if json.Unmarshal(out, &wrapped) == nil && wrapped.Proposal.Status != "" {
		return wrapped.Proposal.Status
	}
	// Fall back to flat.
	var flat struct {
		Status string `json:"status"`
	}
	require.NoError(t, json.Unmarshal(out, &flat),
		"proposal query response did not match either shape: %s", out)
	return flat.Status
}

// VoteOnLatestProposalAndPass votes on the most recently submitted gov
// proposal and waits for it to pass. Use this with CLI proposal helpers
// (e.g. `token update-token-params`, `entity update-entity-params`) that
// submit a gov proposal under the hood — caller submits the proposal,
// then calls this to drive it through voting. Returns the proposal id.
func VoteOnLatestProposalAndPass(t *testing.T, ctx context.Context, chain *cosmos.CosmosChain,
	_ ibc.Wallet,
) uint64 {
	t.Helper()

	// Find the latest proposal. Wait until it's actually in voting
	// period — proposal-deposit transitions happen in EndBlock, so a
	// proposal submitted in the same block as our gov-list call may
	// still be in DEPOSIT_PERIOD when the validator's vote arrives
	// ("inactive proposal" rejection). Poll for up to ~10s.
	WaitBlocks(t, ctx, chain, 1)
	stdout, _, err := chain.GetNode().ExecQuery(ctx, "gov", "proposals", "--output", "json")
	require.NoError(t, err)
	proposalID := lastProposalID(t, stdout)

	for tries := 0; tries < 6; tries++ {
		status := queryProposalStatus(t, ctx, chain, proposalID)
		if status == "PROPOSAL_STATUS_VOTING_PERIOD" {
			break
		}
		if status == "PROPOSAL_STATUS_PASSED" || status == "PROPOSAL_STATUS_REJECTED" ||
			status == "PROPOSAL_STATUS_FAILED" {
			break
		}
		WaitBlocks(t, ctx, chain, 1)
	}

	require.NoError(t, chain.VoteOnProposalAllValidators(ctx, proposalID, "yes"))

	WaitBlocks(t, ctx, chain, 6)

	require.Equal(t, "PROPOSAL_STATUS_PASSED", queryProposalStatus(t, ctx, chain, proposalID),
		"proposal %d did not pass", proposalID)

	return proposalID
}

// lastProposalID extracts the highest proposal_id from a `gov proposals`
// JSON response. Robust to the SDK's habit of returning string-encoded ints
// and to the v1 → v1beta1 schema drift.
func lastProposalID(t *testing.T, out []byte) uint64 {
	t.Helper()
	type respLite struct {
		Proposals []struct {
			ID         string `json:"id,omitempty"`
			ProposalID string `json:"proposal_id,omitempty"`
		} `json:"proposals"`
	}
	var r respLite
	require.NoError(t, json.Unmarshal(out, &r), "unmarshal gov proposals response")
	require.NotEmpty(t, r.Proposals, "no proposals found in response")
	last := r.Proposals[len(r.Proposals)-1]
	idStr := last.ID
	if idStr == "" {
		idStr = last.ProposalID
	}
	id, err := strconv.ParseUint(idStr, 10, 64)
	require.NoError(t, err, "proposal id %q is not a uint64", idStr)
	return id
}
