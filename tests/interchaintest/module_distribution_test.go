//go:build interchaintest

package interchaintest

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestIxoDistribution_FullScenario boots ONE chain and walks the
// distribution module's full delegator-and-validator integration path:
//
//	staking validators → delegate 1M uixo →
//	  query distribution-rewards (well-formed, may be zero) →
//	  query delegations (validator visible) →
//	  query validator-distribution-info → query commission →
//	  set-withdraw-addr (no-op, exercises the msg path).
//
// We don't assert non-zero rewards/commission because the chain's mint
// module uses a `day` epoch identifier and the test window is ~60s —
// inflation hasn't fired yet. The keeper-level tests in x/distribution
// cover the math directly. This test guards the L3 integration
// boundary: every CLI surface resolves, the validator-distribution
// query plumbing is intact, and SetWithdrawAddr msg goes through the
// ante chain cleanly.
func TestIxoDistribution_FullScenario(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping interchaintest in -short mode")
	}
	t.Parallel()

	chain, users, ctx := SetupIxoChain(t, 1)
	delegator := users[0]

	var validatorAddr string
	t.Run("setup: discover validator", func(t *testing.T) {
		stdout, _, err := chain.GetNode().ExecQuery(ctx,
			"staking", "validators", "--output", "json")
		require.NoError(t, err)
		var vals struct {
			Validators []struct {
				OperatorAddress string `json:"operator_address"`
			} `json:"validators"`
		}
		require.NoError(t, json.Unmarshal(stdout, &vals))
		require.NotEmpty(t, vals.Validators)
		validatorAddr = vals.Validators[0].OperatorAddress
	})
	if t.Failed() {
		return
	}

	t.Run("delegate to start accruing rewards", func(t *testing.T) {
		out, err := chain.GetNode().ExecTx(ctx, delegator.KeyName(),
			"staking", "delegate", validatorAddr, "1000000uixo",
			"--gas", "auto", "--gas-adjustment", "1.5",
		)
		require.NoError(t, err, "delegate: %s", out)
	})

	WaitBlocks(t, ctx, chain, 8)

	t.Run("distribution-rewards query is well-formed", func(t *testing.T) {
		stdout, _, err := chain.GetNode().ExecQuery(ctx,
			"distribution", "rewards", delegator.FormattedAddress(),
			"--output", "json")
		require.NoError(t, err, "distribution rewards: %s", stdout)

		var rewardsResp struct {
			Total json.RawMessage `json:"total"`
		}
		require.NoError(t, json.Unmarshal(stdout, &rewardsResp))
		require.NotNil(t, rewardsResp.Total,
			"distribution rewards response must include `total` field; raw: %s", stdout)
	})

	t.Run("delegations query reflects the new delegation", func(t *testing.T) {
		stdout, _, err := chain.GetNode().ExecQuery(ctx,
			"staking", "delegations", delegator.FormattedAddress(),
			"--output", "json")
		require.NoError(t, err)
		require.Contains(t, string(stdout), validatorAddr,
			"delegator's delegation to %s must show up in staking delegations query: %s",
			validatorAddr, stdout)
	})

	t.Run("validator-distribution-info query resolves", func(t *testing.T) {
		stdout, _, err := chain.GetNode().ExecQuery(ctx,
			"distribution", "validator-distribution-info", validatorAddr,
			"--output", "json")
		require.NoError(t, err, "validator-distribution-info: %s", stdout)
		require.Contains(t, string(stdout), "operator_address",
			"validator-distribution-info must include operator_address: %s", stdout)
	})

	t.Run("commission query resolves with a commission key", func(t *testing.T) {
		stdout, _, err := chain.GetNode().ExecQuery(ctx,
			"distribution", "commission", validatorAddr,
			"--output", "json")
		require.NoError(t, err, "commission: %s", stdout)
		require.Contains(t, string(stdout), "commission",
			"commission response must include commission key: %s", stdout)
	})

	t.Run("set-withdraw-addr exercises the SetWithdrawAddr msg", func(t *testing.T) {
		// Setting the withdraw address to ourselves is semantically a
		// no-op but goes through the same ante-handler chain as a real
		// change. Bump gas-adjustment because auto-estimation
		// undershoots distribution txs.
		out, err := chain.GetNode().ExecTx(ctx, delegator.KeyName(),
			"distribution", "set-withdraw-addr", delegator.FormattedAddress(),
			"--gas", "auto", "--gas-adjustment", "2.5",
		)
		require.NoError(t, err, "set-withdraw-addr: %s", out)
	})
}
