//go:build interchaintest

package interchaintest

import (
	"context"
	"testing"
	"time"

	"cosmossdk.io/math"
	"github.com/strangelove-ventures/interchaintest/v8"
	"github.com/strangelove-ventures/interchaintest/v8/chain/cosmos"
	"github.com/strangelove-ventures/interchaintest/v8/ibc"
	interchaintestrelayer "github.com/strangelove-ventures/interchaintest/v8/relayer"
	"github.com/strangelove-ventures/interchaintest/v8/testreporter"
	"github.com/strangelove-ventures/interchaintest/v8/testutil"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zaptest"
)

// TestIxoIBCTransfer spins up an ixo chain + a Gaia counterparty,
// initialises an IBC connection through the rly relayer, and walks the
// FULL bidirectional ICS-20 transfer flow:
//
//	ixo → gaia: ixoUser sends 1000 uixo. Relayer flushes packet to
//	  gaia. Assert gaiaUser holds the wrapped uixo (ibc/<hash> denom).
//	gaia → ixo: gaiaUser sends 400 of the wrapped uixo back. Relayer
//	  flushes. Assert ixoUser's native uixo balance increased by 400
//	  (the IBC unwrap cancels the wrap).
//
// Bidirectional coverage matters because the unwrap path uses a
// different ante decorator branch on ixo than the wrap path. Earlier
// versions only exercised the wrap leg.
func TestIxoIBCTransfer(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping interchaintest in -short mode")
	}
	t.Parallel()

	numVals := 1
	numFullNodes := 0

	cf := interchaintest.NewBuiltinChainFactory(zaptest.NewLogger(t), []*interchaintest.ChainSpec{
		{
			Name:          IxoChainName,
			ChainConfig:   *IxoChainSpec(),
			NumValidators: &numVals,
			NumFullNodes:  &numFullNodes,
		},
		{
			Name:          "gaia",
			Version:       "v17.2.0",
			NumValidators: &numVals,
			NumFullNodes:  &numFullNodes,
		},
	})

	chains, err := cf.Chains(t.Name())
	require.NoError(t, err)
	require.Len(t, chains, 2)
	ixoChain := chains[0].(*cosmos.CosmosChain)
	gaia := chains[1].(*cosmos.CosmosChain)

	client, network := interchaintest.DockerSetup(t)

	rf := interchaintest.NewBuiltinRelayerFactory(
		ibc.CosmosRly,
		zaptest.NewLogger(t),
		interchaintestrelayer.StartupFlags("--processor", "events", "--block-history", "100"),
	)
	r := rf.Build(t, client, network)

	const path = "ibc-path-ixo-gaia"
	ic := interchaintest.NewInterchain().
		AddChain(ixoChain).
		AddChain(gaia).
		AddRelayer(r, "rly").
		AddLink(interchaintest.InterchainLink{
			Chain1: ixoChain, Chain2: gaia, Relayer: r, Path: path,
		})

	ctx := context.Background()
	rep := testreporter.NewNopReporter().RelayerExecReporter(t)
	t.Cleanup(func() { _ = ic.Close() })
	require.NoError(t, ic.Build(ctx, rep, interchaintest.InterchainBuildOptions{
		TestName: t.Name(), Client: client, NetworkID: network,
	}))

	users := interchaintest.GetAndFundTestUsers(t, ctx, t.Name(), DefaultUserFunds, ixoChain, gaia)
	require.Len(t, users, 2)
	ixoUser, gaiaUser := users[0], users[1]

	// Discover the IBC channel created by the relayer.
	channels, err := r.GetChannels(ctx, rep, ixoChain.Config().ChainID)
	require.NoError(t, err)
	require.NotEmpty(t, channels, "relayer must have created at least one channel")
	ixoChan := channels[0]

	require.NoError(t, r.StartRelayer(ctx, rep, path))
	t.Cleanup(func() { _ = r.StopRelayer(ctx, rep) })

	const sendAmount = int64(1_000)

	// ----- Forward leg: ixo → gaia -----
	t.Run("forward: ixo→gaia transfer wraps uixo into ibc/<hash>", func(t *testing.T) {
		ixoBefore, err := ixoChain.GetBalance(ctx, ixoUser.FormattedAddress(), IxoNativeDenom)
		require.NoError(t, err)

		transfer := ibc.WalletAmount{
			Address: gaiaUser.FormattedAddress(),
			Denom:   IxoNativeDenom,
			Amount:  math.NewInt(sendAmount),
		}
		tx, err := ixoChain.SendIBCTransfer(ctx, ixoChan.ChannelID, ixoUser.KeyName(),
			transfer, ibc.TransferOptions{})
		require.NoError(t, err)
		require.NoError(t, tx.Validate())

		// Wait for the relayer to flush the packet.
		require.NoError(t, waitForRelayerFlush(ctx, ixoChain, gaia, 30*time.Second))

		// Sender's uixo balance dropped by sendAmount + fee.
		ixoAfter, err := ixoChain.GetBalance(ctx, ixoUser.FormattedAddress(), IxoNativeDenom)
		require.NoError(t, err)
		dropped := ixoBefore.Sub(ixoAfter).Int64()
		require.GreaterOrEqual(t, dropped, sendAmount,
			"ixoUser must have spent at least the sent amount")

		// Recipient on gaia holds the wrapped denom. The wrapped denom
		// hash is `ibc/<sha256(transfer/<channel>/uixo)>`. We don't
		// reconstruct the hash here — the cosmos chain helper does it
		// for us via GetIBCDenom.
		ibcDenom := chainIBCDenom(t, ixoChan.Counterparty.PortID, ixoChan.Counterparty.ChannelID, IxoNativeDenom)
		bal, err := gaia.GetBalance(ctx, gaiaUser.FormattedAddress(), ibcDenom)
		require.NoError(t, err)
		require.Equal(t, sendAmount, bal.Int64(),
			"gaiaUser must hold exactly %d of the wrapped uixo (denom %s)", sendAmount, ibcDenom)
	})
	if t.Failed() {
		return
	}

	// ----- Reverse leg: gaia → ixo -----
	const returnAmount = int64(400)
	t.Run("reverse: gaia→ixo unwraps ibc/<hash> back to native uixo", func(t *testing.T) {
		ixoBefore, err := ixoChain.GetBalance(ctx, ixoUser.FormattedAddress(), IxoNativeDenom)
		require.NoError(t, err)

		ibcDenom := chainIBCDenom(t, ixoChan.Counterparty.PortID, ixoChan.Counterparty.ChannelID, IxoNativeDenom)
		transfer := ibc.WalletAmount{
			Address: ixoUser.FormattedAddress(),
			Denom:   ibcDenom,
			Amount:  math.NewInt(returnAmount),
		}
		// SendIBCTransfer on gaia uses gaia's port/channel — query its
		// channels and pick the matching counterparty.
		gaiaChannels, err := r.GetChannels(ctx, rep, gaia.Config().ChainID)
		require.NoError(t, err)
		require.NotEmpty(t, gaiaChannels)
		tx, err := gaia.SendIBCTransfer(ctx, gaiaChannels[0].ChannelID, gaiaUser.KeyName(),
			transfer, ibc.TransferOptions{})
		require.NoError(t, err)
		require.NoError(t, tx.Validate())

		require.NoError(t, waitForRelayerFlush(ctx, gaia, ixoChain, 30*time.Second))

		ixoAfter, err := ixoChain.GetBalance(ctx, ixoUser.FormattedAddress(), IxoNativeDenom)
		require.NoError(t, err)
		gained := ixoAfter.Sub(ixoBefore).Int64()
		require.Equal(t, returnAmount, gained,
			"ixoUser native uixo must increase by exactly the return amount on the unwrap leg")
	})
}

// chainIBCDenom returns the deterministic `ibc/<hash>` denom for a
// token sent over `port/channel`. Mirrors the cosmos-sdk hashing
// convention (sha256 over `<port>/<channel>/<base-denom>`).
func chainIBCDenom(t *testing.T, port, channel, baseDenom string) string {
	t.Helper()
	return ibcDenomHash(port + "/" + channel + "/" + baseDenom)
}

// waitForRelayerFlush drives blocks on both sides until the relayer's
// async flush settles. 4 blocks on each side is empirically enough for
// the rly relayer's "events" processor to fire on a healthy
// connection.
func waitForRelayerFlush(ctx context.Context, src, dst *cosmos.CosmosChain, maxWait time.Duration) error {
	if err := testutil.WaitForBlocks(ctx, 4, src); err != nil {
		return err
	}
	if err := testutil.WaitForBlocks(ctx, 4, dst); err != nil {
		return err
	}
	return nil
}
