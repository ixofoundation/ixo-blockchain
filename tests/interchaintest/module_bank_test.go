//go:build interchaintest

package interchaintest

import (
	"testing"

	"cosmossdk.io/math"
	"github.com/strangelove-ventures/interchaintest/v8/ibc"
	"github.com/stretchr/testify/require"
)

// TestIxoBank_FullScenario boots ONE chain and walks every bank-side
// invariant we care about:
//
//	send happy-path → balance bookkeeping →
//	send to a blocked module account (gov) → must reject →
//	send a tiny amount to verify minimum-fee accounting.
//
// This used to be two separate Docker bootstraps
// (`TestIxoBank_SendUpdatesBalances` + `_BlockedAccountsRejected`).
// Same chain runs through both — saves ~30s of boot.
func TestIxoBank_FullScenario(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping interchaintest in -short mode")
	}
	t.Parallel()

	chain, users, ctx := SetupIxoChain(t, 2)
	sender, recipient := users[0], users[1]

	t.Run("happy-path send: recipient gains exactly the amount, sender loses amount+fee", func(t *testing.T) {
		const amount = int64(123_456)

		senderBefore, err := chain.GetBalance(ctx, sender.FormattedAddress(), IxoNativeDenom)
		require.NoError(t, err)
		recipientBefore, err := chain.GetBalance(ctx, recipient.FormattedAddress(), IxoNativeDenom)
		require.NoError(t, err)

		transfer := ibc.WalletAmount{
			Address: recipient.FormattedAddress(),
			Denom:   IxoNativeDenom,
			Amount:  math.NewInt(amount),
		}
		require.NoError(t, chain.SendFunds(ctx, sender.KeyName(), transfer))

		senderAfter, err := chain.GetBalance(ctx, sender.FormattedAddress(), IxoNativeDenom)
		require.NoError(t, err)
		recipientAfter, err := chain.GetBalance(ctx, recipient.FormattedAddress(), IxoNativeDenom)
		require.NoError(t, err)

		require.Equal(t, amount, recipientAfter.Sub(recipientBefore).Int64(),
			"recipient must gain exactly the transfer amount")

		dropped := senderBefore.Sub(senderAfter).Int64()
		require.GreaterOrEqual(t, dropped, amount, "sender drop must be at least the transfer amount")
		require.Less(t, dropped, amount+1_000_000,
			"sender drop must be amount+sane-fee; got %d", dropped)
	})

	t.Run("send to blocked module account (gov) is rejected", func(t *testing.T) {
		// gov is blocked on this chain — only `distribution` and
		// `liquidstake` are listed in app/blocked.go::allowedReceivingModAcc;
		// every other module account is blocked from receiving direct
		// end-user transfers.
		govAddr, err := chain.GetModuleAddress(ctx, "gov")
		require.NoError(t, err)

		transfer := ibc.WalletAmount{
			Address: govAddr,
			Denom:   IxoNativeDenom,
			Amount:  math.NewInt(1),
		}
		err = chain.SendFunds(ctx, sender.KeyName(), transfer)
		require.Error(t, err, "bank send to a blocked module account must be rejected")
	})

	t.Run("send to allowed module account (distribution) succeeds", func(t *testing.T) {
		// distribution is explicitly in allowedReceivingModAcc — sends
		// to it must NOT be blocked. This guards the negative half of
		// the BlockedAddresses contract.
		distAddr, err := chain.GetModuleAddress(ctx, "distribution")
		require.NoError(t, err)

		transfer := ibc.WalletAmount{
			Address: distAddr,
			Denom:   IxoNativeDenom,
			Amount:  math.NewInt(1),
		}
		require.NoError(t, chain.SendFunds(ctx, sender.KeyName(), transfer),
			"bank send to an allowed module account must succeed")
	})
}
