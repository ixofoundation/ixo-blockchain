//go:build interchaintest

package interchaintest

import (
	"context"
	"encoding/json"
	"strconv"
	"testing"

	"github.com/strangelove-ventures/interchaintest/v8/chain/cosmos"
	"github.com/stretchr/testify/require"
)

// TestIxoChainTime_FullScenario boots ONE chain and exercises the
// "chain time / inflation" query surface that x/mint and x/epochs
// share (mint subscribes to an epochs hook to drive its inflation).
//
// Subsumes two earlier separate Docker bootstraps
// (`TestIxoMint_QueryParamsAndProvisions` + `TestIxoEpochs_QueryEpochInfos`).
//
// Inflation actually firing isn't asserted here — the chain's default
// mint epoch identifier is `day` (24h) and the default epoch genesis
// only ships day/hour/week, so a 60s test can't observe a real tick.
// The keeper-level tests in x/mint and x/epochs drive inflation
// directly. This test guards the integration boundary: gRPC/CLI query
// path is reachable, params are populated with sane values, and the
// chain ships the expected default identifiers.
func TestIxoChainTime_FullScenario(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping interchaintest in -short mode")
	}
	t.Parallel()

	chain, _, ctx := SetupIxoChain(t, 1)

	t.Run("mint params + epoch-provisions queries resolve", func(t *testing.T) {
		stdout, _, err := chain.GetNode().ExecQuery(ctx,
			"mint", "params", "--output", "json")
		require.NoError(t, err, "mint params: %s", stdout)
		require.Contains(t, string(stdout), "epoch_identifier",
			"mint params response must include epoch_identifier")
		require.Contains(t, string(stdout), "uixo",
			"mint params response must reference uixo as mint denom")

		stdout, _, err = chain.GetNode().ExecQuery(ctx,
			"mint", "epoch-provisions", "--output", "json")
		require.NoError(t, err, "mint epoch-provisions: %s", stdout)
		require.NotEmpty(t, stdout)
	})

	t.Run("total uixo supply is positive on a fresh chain", func(t *testing.T) {
		require.Greater(t, totalUixoSupply(t, ctx, chain), uint64(0),
			"total uixo supply must be positive on a freshly-bootstrapped chain")
	})

	t.Run("epochs: default identifiers (day/hour/week) all resolve", func(t *testing.T) {
		stdout, _, err := chain.GetNode().ExecQuery(ctx,
			"epochs", "epoch-infos", "--output", "json")
		require.NoError(t, err, "epoch-infos: %s", stdout)

		var infos struct {
			Epochs []struct {
				Identifier   string `json:"identifier"`
				Duration     string `json:"duration"`
				CurrentEpoch string `json:"current_epoch"`
			} `json:"epochs"`
		}
		require.NoError(t, json.Unmarshal(stdout, &infos))
		require.NotEmpty(t, infos.Epochs)

		identifiers := make(map[string]bool, len(infos.Epochs))
		for _, e := range infos.Epochs {
			identifiers[e.Identifier] = true
			_, err := strconv.ParseInt(e.CurrentEpoch, 10, 64)
			require.NoError(t, err, "current_epoch must parse as int for %q, got %q",
				e.Identifier, e.CurrentEpoch)
			_, _, err = chain.GetNode().ExecQuery(ctx,
				"epochs", "current-epoch", e.Identifier, "--output", "json")
			require.NoError(t, err, "current-epoch %q must resolve", e.Identifier)
		}

		for _, want := range []string{"day", "hour", "week"} {
			require.True(t, identifiers[want],
				"default genesis must register %q epoch; got %v", want, identifiers)
		}
	})
}

// totalUixoSupply pulls the current total uixo supply from the bank
// module's supply query and returns it as a uint64.
func totalUixoSupply(t *testing.T, ctx context.Context, chain *cosmos.CosmosChain) uint64 {
	t.Helper()
	stdout, _, err := chain.GetNode().ExecQuery(ctx,
		"bank", "total-supply-of", IxoNativeDenom, "--output", "json")
	require.NoError(t, err, "bank total-supply-of: %s", stdout)
	var resp struct {
		Amount struct {
			Amount string `json:"amount"`
		} `json:"amount"`
	}
	require.NoError(t, json.Unmarshal(stdout, &resp))
	require.NotEmpty(t, resp.Amount.Amount)

	var n uint64
	for _, c := range resp.Amount.Amount {
		require.True(t, c >= '0' && c <= '9',
			"unexpected non-digit in supply amount %q", resp.Amount.Amount)
		n = n*10 + uint64(c-'0')
	}
	return n
}
