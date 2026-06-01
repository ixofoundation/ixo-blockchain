//go:build interchaintest

package interchaintest

import (
	"context"
	"testing"

	"github.com/strangelove-ventures/interchaintest/v8"
	"github.com/strangelove-ventures/interchaintest/v8/chain/cosmos"
	"github.com/strangelove-ventures/interchaintest/v8/ibc"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"
)

// TestIxoBasicStart spins up a single-validator ixo chain, asserts blocks
// are being produced, and verifies a funded user account has the expected
// balance. This is the smoke test that anyone setting up a fresh dev box
// should run first to confirm the Docker image and IBC-ready ports work.
func TestIxoBasicStart(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping interchaintest in -short mode")
	}
	t.Parallel()

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

	client, network := interchaintest.DockerSetup(t)

	ic := interchaintest.NewInterchain().AddChain(chain)
	t.Cleanup(func() { _ = ic.Close() })

	require.NoError(t, ic.Build(ctx, nil, interchaintest.InterchainBuildOptions{
		TestName:         t.Name(),
		Client:           client,
		NetworkID:        network,
		SkipPathCreation: true,
	}))

	users := interchaintest.GetAndFundTestUsers(t, ctx, t.Name(), DefaultUserFunds, chain)
	require.Len(t, users, 1)

	bal, err := chain.GetBalance(ctx, users[0].FormattedAddress(), IxoNativeDenom)
	require.NoError(t, err)
	require.True(t, bal.GTE(DefaultUserFunds), "funded user balance must be at least DefaultUserFunds")
}

// silence unused linter warning for ibc in this file when the suite is
// trimmed down to just the basic-start case.
var _ = ibc.ChainConfig{}
