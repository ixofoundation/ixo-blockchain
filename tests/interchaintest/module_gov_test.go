//go:build interchaintest

package interchaintest

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestIxoGov_FullScenario boots ONE chain and walks every distinct
// gov-driven flow on the chain — each is a different `@type` payload
// running through the same submit→deposit→vote→tally→execute pipeline,
// and each historically lived in its own ~60-second Docker bootstrap.
//
// Subsumes earlier separate tests:
//
//	TestIxoNames_GovCreateAndQuery        →  create namespace
//	TestIxoUpgrade_SoftwareUpgradeProposal →  schedule software upgrade
//	TestIxoLiquidStake_GovCreatePoolAndQuery →  register liquidstake pool
//	TestIxoGov_RejectedProposalDoesNotApply →  vote-NO rollback
//
// Putting them in one chain saves ~3 × 30s of bootstrap and exercises
// proposal-id sequencing through the same gov keeper instance — a
// regression in proposal-id allocation would surface here as
// "wrong proposal id" rather than each test getting id=1 in isolation.
func TestIxoGov_FullScenario(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping interchaintest in -short mode")
	}
	t.Parallel()

	chain, users, ctx := SetupIxoChain(t, 1)
	proposer := users[0]

	govAddr, err := chain.GetModuleAddress(ctx, "gov")
	require.NoError(t, err)

	// 1. CREATE NAMESPACE — names module gov path
	t.Run("MsgCreateNamespace passes and the namespace is queryable", func(t *testing.T) {
		const nsName = "demo"
		proposal := fmt.Sprintf(`{
  "messages": [{
    "@type": "/ixo.names.v1beta1.MsgCreateNamespace",
    "authority": %q,
    "namespace": {
      "name": %q,
      "description": "demo namespace for interchaintest",
      "allow_self_register": true,
      "min_length": 3,
      "max_length": 32
    }
  }],
  "metadata": "create demo namespace",
  "deposit": "10000000uixo",
  "title": "register demo namespace",
  "summary": "Adds the demo namespace via x/names so subsequent name registrations have a target."
}`, govAddr, nsName)
		SubmitGovProposalAndPass(t, ctx, chain, proposer, proposal)

		stdout, _, err := chain.GetNode().ExecQuery(ctx,
			"names", "namespace", nsName, "--output", "json")
		require.NoError(t, err)
		var nsResp struct {
			Namespace struct {
				Name        string `json:"name"`
				Description string `json:"description"`
			} `json:"namespace"`
		}
		require.NoError(t, json.Unmarshal(stdout, &nsResp))
		require.Equal(t, nsName, nsResp.Namespace.Name)

		stdout, _, err = chain.GetNode().ExecQuery(ctx,
			"names", "namespaces", "--output", "json")
		require.NoError(t, err)
		var listResp struct {
			Namespaces []struct {
				Name string `json:"name"`
			} `json:"namespaces"`
		}
		require.NoError(t, json.Unmarshal(stdout, &listResp))
		found := false
		for _, ns := range listResp.Namespaces {
			if ns.Name == nsName {
				found = true
			}
		}
		require.True(t, found, "demo namespace must appear in namespaces list")
	})

	// 2. SOFTWARE UPGRADE — cosmos.upgrade gov path
	t.Run("MsgSoftwareUpgrade schedules the upgrade plan", func(t *testing.T) {
		const (
			upgradeName   = "v7-test"
			upgradeHeight = 1_000_000
		)
		proposal := fmt.Sprintf(`{
  "messages": [{
    "@type": "/cosmos.upgrade.v1beta1.MsgSoftwareUpgrade",
    "authority": %q,
    "plan": {
      "name": %q,
      "height": "%d",
      "info": "interchaintest upgrade dry-run"
    }
  }],
  "metadata": "software-upgrade dry-run",
  "deposit": "10000000uixo",
  "title": "schedule v7-test upgrade",
  "summary": "Schedules a no-op upgrade for height %d so the upgrade plan query has data."
}`, govAddr, upgradeName, upgradeHeight, upgradeHeight)
		SubmitGovProposalAndPass(t, ctx, chain, proposer, proposal)

		stdout, _, err := chain.GetNode().ExecQuery(ctx,
			"upgrade", "plan", "--output", "json")
		require.NoError(t, err)
		var planResp struct {
			Plan struct {
				Name   string `json:"name"`
				Height string `json:"height"`
			} `json:"plan"`
		}
		require.NoError(t, json.Unmarshal(stdout, &planResp))
		require.Equal(t, upgradeName, planResp.Plan.Name)
		require.Equal(t, fmt.Sprintf("%d", upgradeHeight), planResp.Plan.Height)
	})

	// 3. LIQUIDSTAKE POOL — custom ixo gov path
	t.Run("MsgCreatePool registers a liquidstake pool", func(t *testing.T) {
		const (
			poolID      = "zero"
			liquidDenom = "uzero"
		)
		feeAccount := proposer.FormattedAddress()
		adminAddress := proposer.FormattedAddress()
		proposal := fmt.Sprintf(`{
  "messages": [{
    "@type": "/ixo.liquidstake.v1beta1.MsgCreatePool",
    "authority": %q,
    "pool_id": %q,
    "liquid_bond_denom": %q,
    "initial_admin_address": %q,
    "initial_fee_account_address": %q
  }],
  "metadata": "create liquidstake zero pool",
  "deposit": "10000000uixo",
  "title": "register zero liquidstake pool",
  "summary": "Adds a new liquidstake pool that mints uzero against bonded uixo."
}`, govAddr, poolID, liquidDenom, adminAddress, feeAccount)
		SubmitGovProposalAndPass(t, ctx, chain, proposer, proposal)

		stdout, _, err := chain.GetNode().ExecQuery(ctx,
			"liquidstake", "pool", poolID, "--output", "json")
		require.NoError(t, err)
		var poolResp struct {
			Pool struct {
				PoolId          string `json:"pool_id"`
				LiquidBondDenom string `json:"liquid_bond_denom"`
			} `json:"pool"`
		}
		require.NoError(t, json.Unmarshal(stdout, &poolResp))
		require.Equal(t, poolID, poolResp.Pool.PoolId)
		require.Equal(t, liquidDenom, poolResp.Pool.LiquidBondDenom)

		stdout, _, err = chain.GetNode().ExecQuery(ctx,
			"liquidstake", "module-params", "--output", "json")
		require.NoError(t, err)
		require.NotEmpty(t, stdout)
	})

	// 4. REJECTION ROLLBACK — vote NO; namespace must NOT be created
	t.Run("MsgCreateNamespace voted NO does not apply", func(t *testing.T) {
		const nsName = "rejected-ns"
		proposalJSON := fmt.Sprintf(`{
  "messages": [{
    "@type": "/ixo.names.v1beta1.MsgCreateNamespace",
    "authority": %q,
    "namespace": {
      "name": %q,
      "description": "this proposal will be rejected",
      "allow_self_register": true,
      "min_length": 3,
      "max_length": 32
    }
  }],
  "metadata": "rejection test",
  "deposit": "10000000uixo",
  "title": "rejection test",
  "summary": "Submitted then voted NO — must not pass."
}`, govAddr, nsName)

		const proposalRel = "rejected-proposal.json"
		require.NoError(t, chain.GetNode().WriteFile(ctx, []byte(proposalJSON), proposalRel))
		out, err := chain.GetNode().ExecTx(ctx, proposer.KeyName(),
			"gov", "submit-proposal", chain.GetNode().HomeDir()+"/"+proposalRel,
			"--gas", "auto", "--gas-adjustment", "2.0")
		require.NoError(t, err, "submit-proposal: %s", out)

		stdout, _, err := chain.GetNode().ExecQuery(ctx, "gov", "proposals", "--output", "json")
		require.NoError(t, err)
		proposalID := lastProposalID(t, stdout)

		require.NoError(t, chain.VoteOnProposalAllValidators(ctx, proposalID, "no"))
		WaitBlocks(t, ctx, chain, 6)

		require.Equal(t, "PROPOSAL_STATUS_REJECTED",
			queryProposalStatus(t, ctx, chain, proposalID),
			"proposal %d must be rejected after a NO vote", proposalID)

		// Namespace must NOT exist.
		stdout, _, err = chain.GetNode().ExecQuery(ctx,
			"names", "namespace", nsName, "--output", "json")
		if err != nil {
			return // expected: not-found error
		}
		var nsResp struct {
			Namespace *struct {
				Name string `json:"name"`
			} `json:"namespace"`
		}
		require.NoError(t, json.Unmarshal(stdout, &nsResp))
		require.Nil(t, nsResp.Namespace,
			"rejected MsgCreateNamespace must not create a namespace; got %+v", nsResp.Namespace)
	})
}
