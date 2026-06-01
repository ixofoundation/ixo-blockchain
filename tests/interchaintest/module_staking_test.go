//go:build interchaintest

package interchaintest

import (
	"encoding/json"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestIxoStaking_FullScenario boots ONE chain and walks the full
// delegator lifecycle on the chain's single bonded validator:
//
//	query validators → delegate 1M uixo → query delegations →
//	  partial unbond (starts an unbonding entry) → query
//	  unbonding-delegations (entry exists) → cancel-unbond against
//	  that entry → query unbonding-delegations (entry gone).
//
// Mirrors the SDK's `cosmos.ts::stakingBasic` pattern. The chain has
// one validator out-of-the-box (interchaintest's gentx), so redelegate
// (which needs ≥2 validators) is intentionally not exercised here.
func TestIxoStaking_FullScenario(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping interchaintest in -short mode")
	}
	t.Parallel()

	chain, users, ctx := SetupIxoChain(t, 1)
	delegator := users[0]

	var validatorAddr string
	t.Run("query validators returns the genesis validator", func(t *testing.T) {
		stdout, _, err := chain.GetNode().ExecQuery(ctx,
			"staking", "validators", "--output", "json")
		require.NoError(t, err)
		var resp struct {
			Validators []struct {
				OperatorAddress string `json:"operator_address"`
				Status          string `json:"status"`
			} `json:"validators"`
		}
		require.NoError(t, json.Unmarshal(stdout, &resp))
		require.NotEmpty(t, resp.Validators, "fresh chain must have at least one validator")
		validatorAddr = resp.Validators[0].OperatorAddress
		require.Equal(t, "BOND_STATUS_BONDED", resp.Validators[0].Status,
			"genesis validator must start in BONDED state")
	})
	if t.Failed() {
		return
	}

	const stakeAmount = "1000000" // 1 IXO

	t.Run("delegate 1 IXO to the validator", func(t *testing.T) {
		out, err := chain.GetNode().ExecTx(ctx, delegator.KeyName(),
			"staking", "delegate", validatorAddr, stakeAmount+IxoNativeDenom,
			"--gas", "auto", "--gas-adjustment", "1.5",
		)
		require.NoError(t, err, "delegate: %s", out)
	})

	t.Run("query delegations confirms the delegation", func(t *testing.T) {
		stdout, _, err := chain.GetNode().ExecQuery(ctx,
			"staking", "delegations", delegator.FormattedAddress(),
			"--output", "json")
		require.NoError(t, err, "delegations: %s", stdout)
		require.Contains(t, string(stdout), validatorAddr,
			"delegations must include the validator we just delegated to")
		require.Contains(t, string(stdout), stakeAmount,
			"delegations must show the staked amount")
	})

	const unbondAmount = "400000" // 0.4 IXO

	var creationHeight int64
	t.Run("unbond a fraction starts an unbonding-delegation entry", func(t *testing.T) {
		out, err := chain.GetNode().ExecTx(ctx, delegator.KeyName(),
			"staking", "unbond", validatorAddr, unbondAmount+IxoNativeDenom,
			"--gas", "auto", "--gas-adjustment", "1.5",
		)
		require.NoError(t, err, "unbond: %s", out)

		WaitBlocks(t, ctx, chain, 2)

		stdout, _, err := chain.GetNode().ExecQuery(ctx,
			"staking", "unbonding-delegations", delegator.FormattedAddress(),
			"--output", "json")
		require.NoError(t, err, "unbonding-delegations: %s", stdout)
		var resp struct {
			UnbondingResponses []struct {
				Entries []struct {
					CreationHeight string `json:"creation_height"`
					Balance        string `json:"balance"`
				} `json:"entries"`
			} `json:"unbonding_responses"`
		}
		require.NoError(t, json.Unmarshal(stdout, &resp))
		require.NotEmpty(t, resp.UnbondingResponses,
			"unbonding-delegations must include the new entry")
		require.NotEmpty(t, resp.UnbondingResponses[0].Entries)
		entry := resp.UnbondingResponses[0].Entries[0]
		require.Equal(t, unbondAmount, entry.Balance,
			"unbonding entry balance must equal the unbonded amount")
		creationHeight, err = strconv.ParseInt(entry.CreationHeight, 10, 64)
		require.NoError(t, err)
		require.Greater(t, creationHeight, int64(0))
	})
	if t.Failed() {
		return
	}

	t.Run("cancel-unbond rolls the entry back into delegation", func(t *testing.T) {
		out, err := chain.GetNode().ExecTx(ctx, delegator.KeyName(),
			"staking", "cancel-unbond", validatorAddr,
			unbondAmount+IxoNativeDenom,
			strconv.FormatInt(creationHeight, 10),
			"--gas", "auto", "--gas-adjustment", "1.5",
		)
		require.NoError(t, err, "cancel-unbond: %s", out)

		WaitBlocks(t, ctx, chain, 2)

		// Unbonding entry must be gone.
		stdout, _, err := chain.GetNode().ExecQuery(ctx,
			"staking", "unbonding-delegations", delegator.FormattedAddress(),
			"--output", "json")
		require.NoError(t, err)
		var resp struct {
			UnbondingResponses []json.RawMessage `json:"unbonding_responses"`
		}
		require.NoError(t, json.Unmarshal(stdout, &resp))
		require.Empty(t, resp.UnbondingResponses,
			"unbonding-delegation must be gone after cancel-unbond")

		// Full stake amount is back in the delegation.
		stdout, _, err = chain.GetNode().ExecQuery(ctx,
			"staking", "delegations", delegator.FormattedAddress(),
			"--output", "json")
		require.NoError(t, err)
		require.Contains(t, string(stdout), stakeAmount,
			"after cancel-unbond, delegation must hold the full original amount: %s", stdout)
	})

	t.Run("query delegator-validators returns the validator", func(t *testing.T) {
		stdout, _, err := chain.GetNode().ExecQuery(ctx,
			"staking", "delegator-validators", delegator.FormattedAddress(),
			"--output", "json")
		require.NoError(t, err)
		require.Contains(t, string(stdout), validatorAddr)
	})
}
