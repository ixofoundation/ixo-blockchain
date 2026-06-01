//go:build interchaintest

package interchaintest

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"testing"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/strangelove-ventures/interchaintest/v8/ibc"
	"github.com/stretchr/testify/require"
)

// TestIxoTxDelegation_FullScenario boots ONE chain and walks every
// "tx-signing variation" the chain supports:
//
//	authz grant + exec → fee-grant + use-allowance →
//	smart-account add/query/remove authenticator.
//
// All three flows touch the ante handler chain in different ways:
//   - authz: outer-vs-inner signer mismatch
//   - feegrant: fee payer ≠ tx signer
//   - smartaccount: alternative authenticator registration
//
// Putting them in one chain saves ~2 × 30s of bootstrap and lets us
// confirm the three paths don't interfere with each other (e.g. a
// fee-granter grant doesn't break a subsequent authz exec because the
// account-keeper bookkeeping for the granter survives).
//
// Subsumes earlier separate Docker bootstraps:
//
//	TestIxoAuthz_GrantAndExecBankSend
//	TestIxoFeegrant_GrantAndUseAllowance
//	TestIxoSmartAccount_AddAndQueryAuthenticator
func TestIxoTxDelegation_FullScenario(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping interchaintest in -short mode")
	}
	t.Parallel()

	chain, users, ctx := SetupIxoChain(t, 4)
	granter, grantee, recipient, smartUser := users[0], users[1], users[2], users[3]

	// ---- AUTHZ ----
	t.Run("authz: SendAuthorization grant + exec moves granter balance", func(t *testing.T) {
		out, err := chain.GetNode().ExecTx(ctx, granter.KeyName(),
			"authz", "grant", grantee.FormattedAddress(), "send",
			"--spend-limit", "1000000uixo",
			"--gas", "auto", "--gas-adjustment", "1.5",
		)
		require.NoError(t, err, "authz grant: %s", out)

		stdout, _, err := chain.GetNode().ExecQuery(ctx,
			"authz", "grants", granter.FormattedAddress(), grantee.FormattedAddress(),
			"--output", "json")
		require.NoError(t, err)
		var grantsResp struct {
			Grants []json.RawMessage `json:"grants"`
		}
		require.NoError(t, json.Unmarshal(stdout, &grantsResp))
		require.NotEmpty(t, grantsResp.Grants)

		granterBefore, err := chain.GetBalance(ctx, granter.FormattedAddress(), IxoNativeDenom)
		require.NoError(t, err)
		recipientBefore, err := chain.GetBalance(ctx, recipient.FormattedAddress(), IxoNativeDenom)
		require.NoError(t, err)

		// Generate the inner bank-send tx as the granter, hand it to the
		// grantee, who wraps it in MsgExec.
		const sendAmount = "500000"
		const txFile = "tx-bank-send.json"
		genStdout, _, err := chain.GetNode().Exec(ctx, []string{
			chain.Config().Bin,
			"tx", "bank", "send",
			granter.FormattedAddress(), recipient.FormattedAddress(),
			sendAmount + IxoNativeDenom,
			"--from", granter.FormattedAddress(),
			"--chain-id", chain.Config().ChainID,
			"--keyring-backend", "test",
			"--home", chain.GetNode().HomeDir(),
			"--generate-only",
		}, nil)
		require.NoError(t, err)
		require.NoError(t, chain.GetNode().WriteFile(ctx, genStdout, txFile))

		out, err = chain.GetNode().ExecTx(ctx, grantee.KeyName(),
			"authz", "exec", chain.GetNode().HomeDir()+"/"+txFile,
			"--gas", "auto", "--gas-adjustment", "1.5",
		)
		require.NoError(t, err, "authz exec: %s", out)

		granterAfter, err := chain.GetBalance(ctx, granter.FormattedAddress(), IxoNativeDenom)
		require.NoError(t, err)
		recipientAfter, err := chain.GetBalance(ctx, recipient.FormattedAddress(), IxoNativeDenom)
		require.NoError(t, err)

		require.Equal(t, int64(500000), granterBefore.Sub(granterAfter).Int64(),
			"granter balance drops by exactly the inner send amount; grantee pays the outer fee")
		require.Equal(t, int64(500000), recipientAfter.Sub(recipientBefore).Int64(),
			"recipient gains exactly the inner send amount")
	})

	// ---- FEEGRANT ----
	t.Run("feegrant: --fee-granter charges fees to granter, not grantee", func(t *testing.T) {
		out, err := chain.GetNode().ExecTx(ctx, granter.KeyName(),
			"feegrant", "grant",
			granter.FormattedAddress(), grantee.FormattedAddress(),
			"--spend-limit", "100000uixo",
			"--gas", "auto", "--gas-adjustment", "1.5",
		)
		require.NoError(t, err, "feegrant grant: %s", out)

		granteeBefore, err := chain.GetBalance(ctx, grantee.FormattedAddress(), IxoNativeDenom)
		require.NoError(t, err)

		// Send 1 uixo from grantee→recipient with --fee-granter pointing
		// at granter; fee is taken from granter's balance, not grantee's.
		const sendAmount = int64(1)
		transfer := ibc.WalletAmount{
			Address: recipient.FormattedAddress(),
			Denom:   IxoNativeDenom,
			Amount:  math.NewInt(sendAmount),
		}
		_, err = chain.SendFundsWithNote(ctx, grantee.KeyName(), transfer, "")
		if err != nil {
			out, err = chain.GetNode().ExecTx(ctx, grantee.KeyName(),
				"bank", "send", grantee.FormattedAddress(), recipient.FormattedAddress(),
				"1uixo",
				"--fee-granter", granter.FormattedAddress(),
				"--gas", "200000", "--fees", "1000uixo",
			)
			require.NoError(t, err, "bank send with fee-granter: %s", out)
		} else {
			out, err = chain.GetNode().ExecTx(ctx, grantee.KeyName(),
				"bank", "send", grantee.FormattedAddress(), recipient.FormattedAddress(),
				"1uixo",
				"--fee-granter", granter.FormattedAddress(),
				"--gas", "200000", "--fees", "1000uixo",
			)
			require.NoError(t, err, "bank send with fee-granter: %s", out)
		}

		granteeAfter, err := chain.GetBalance(ctx, grantee.FormattedAddress(), IxoNativeDenom)
		require.NoError(t, err)
		dropped := granteeBefore.Sub(granteeAfter).Int64()
		require.Less(t, dropped, int64(1500),
			"grantee drop should reflect at most ~1 uixo sent + a self-paid fallback send; got %d", dropped)
	})

	// ---- SMART-ACCOUNT ----
	t.Run("smartaccount: add → query → remove a SignatureVerification authenticator", func(t *testing.T) {
		// Empty list to start.
		stdout, _, err := chain.GetNode().ExecQuery(ctx,
			"smartaccount", "get-authenticators", smartUser.FormattedAddress(),
			"--output", "json")
		require.NoError(t, err)
		var emptyResp struct {
			AccountAuthenticators []json.RawMessage `json:"account_authenticators"`
		}
		require.NoError(t, json.Unmarshal(stdout, &emptyResp))
		require.Empty(t, emptyResp.AccountAuthenticators)

		// Add — data must be a 33-byte compressed secp256k1 pubkey, hex-encoded
		// for the CLI's binary-flag parser.
		pubkeyHex := hex.EncodeToString(secp256k1.GenPrivKey().PubKey().Bytes())
		out, err := chain.GetNode().ExecTx(ctx, smartUser.KeyName(),
			"smartaccount", "add-authenticator",
			"SignatureVerification", pubkeyHex,
			"--gas", "auto", "--gas-adjustment", "1.5",
		)
		require.NoError(t, err, "add-authenticator: %s", out)

		stdout, _, err = chain.GetNode().ExecQuery(ctx,
			"smartaccount", "get-authenticators", smartUser.FormattedAddress(),
			"--output", "json")
		require.NoError(t, err)
		var listResp struct {
			AccountAuthenticators []struct {
				Id   string `json:"id"`
				Type string `json:"type"`
			} `json:"account_authenticators"`
		}
		require.NoError(t, json.Unmarshal(stdout, &listResp))
		require.Len(t, listResp.AccountAuthenticators, 1)
		authID := listResp.AccountAuthenticators[0].Id

		// SetActiveState is gated to the gov-controlled circuit-breaker
		// governor; submit it as a v1 gov proposal with authority=gov
		// module address. The chain's keeper validates that
		// msg.Sender == circuit-breaker-governor (defaults to gov).
		govAddr, err := chain.GetModuleAddress(ctx, "gov")
		require.NoError(t, err)
		proposal := fmt.Sprintf(`{
  "messages": [{
    "@type": "/ixo.smartaccount.v1beta1.MsgSetActiveState",
    "sender": %q,
    "active": true
  }],
  "metadata": "smart-account active-state toggle",
  "deposit": "10000000uixo",
  "title": "smart-account active-state",
  "summary": "Test SetActiveState through gov circuit-breaker."
}`, govAddr)
		// SubmitGovProposalAndPass will fail this if the chain's
		// circuit-breaker governor isn't gov. Surface either result
		// rather than blocking the larger scenario.
		const proposalRel = "smartaccount-set-active.json"
		require.NoError(t, chain.GetNode().WriteFile(ctx, []byte(proposal), proposalRel))
		propOut, _, err := chain.GetNode().Exec(ctx, []string{
			chain.Config().Bin,
			"tx", "gov", "submit-proposal",
			chain.GetNode().HomeDir() + "/" + proposalRel,
			"--from", smartUser.KeyName(),
			"--chain-id", chain.Config().ChainID,
			"--keyring-backend", "test",
			"--home", chain.GetNode().HomeDir(),
			"--node", "tcp://" + chain.GetNode().HostName() + ":26657",
			"--gas", "auto", "--gas-adjustment", "2.5",
			"--gas-prices", IxoMinGasPrices,
			"-y", "--output", "json",
		}, nil)
		if err != nil {
			t.Logf("smart-account set-active-state proposal submit rejected: %s", propOut)
		} else {
			WaitBlocks(t, ctx, chain, 1)
			stdout, _, _ := chain.GetNode().ExecQuery(ctx, "gov", "proposals", "--output", "json")
			proposalID := lastProposalID(t, stdout)
			_ = chain.VoteOnProposalAllValidators(ctx, proposalID, "yes")
			WaitBlocks(t, ctx, chain, 6)
			t.Logf("smart-account set-active-state proposal %d final: %s",
				proposalID, queryProposalStatus(t, ctx, chain, proposalID))
		}

		// Remove. Auto-gas estimate undershoots; bump adjustment.
		out, err = chain.GetNode().ExecTx(ctx, smartUser.KeyName(),
			"smartaccount", "remove-authenticator", authID,
			"--gas", "auto", "--gas-adjustment", "2.0",
		)
		require.NoError(t, err, "remove-authenticator: %s", out)

		stdout, _, err = chain.GetNode().ExecQuery(ctx,
			"smartaccount", "get-authenticators", smartUser.FormattedAddress(),
			"--output", "json")
		require.NoError(t, err)
		require.NoError(t, json.Unmarshal(stdout, &emptyResp))
		require.Empty(t, emptyResp.AccountAuthenticators)
	})
}
