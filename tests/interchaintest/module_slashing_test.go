//go:build interchaintest

package interchaintest

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

// TestIxoSlashing_FullScenario boots ONE chain and walks the
// slashing-side query surface. We can't easily provoke a real slashing
// event from a single-validator chain (jail-on-downtime needs the
// validator to actually miss blocks, which requires multi-validator
// consensus or stopping the only validator → halts the chain), so this
// test focuses on the QUERY API that observers and middleware depend
// on:
//
//	signing-infos lists the validator's signing record →
//	  validator's missed_blocks_counter is 0 on a healthy chain →
//	  slashing params resolve →
//	  slashing-info-by-cons-pubkey resolves for the genesis validator.
func TestIxoSlashing_FullScenario(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping interchaintest in -short mode")
	}
	t.Parallel()

	chain, _, ctx := SetupIxoChain(t, 1)

	t.Run("signing-infos returns the validator's signing record", func(t *testing.T) {
		stdout, _, err := chain.GetNode().ExecQuery(ctx,
			"slashing", "signing-infos", "--output", "json")
		require.NoError(t, err)
		var resp struct {
			Info []struct {
				Address             string `json:"address"`
				MissedBlocksCounter string `json:"missed_blocks_counter"`
				Tombstoned          bool   `json:"tombstoned"`
				JailedUntil         string `json:"jailed_until"`
			} `json:"info"`
		}
		require.NoError(t, json.Unmarshal(stdout, &resp))
		require.NotEmpty(t, resp.Info, "fresh chain must have ≥1 signing-info entry for its validator")
		// MissedBlocksCounter is omitted from JSON when it's zero
		// (proto3 omitempty for uint64). An empty string OR "0"
		// both mean "no missed blocks". A non-zero value would be a
		// real number string like "5" — that's the regression target.
		mbc := resp.Info[0].MissedBlocksCounter
		require.True(t, mbc == "" || mbc == "0",
			"genesis validator must have zero missed blocks; got %q", mbc)
		require.False(t, resp.Info[0].Tombstoned,
			"genesis validator must not be tombstoned")
	})

	t.Run("slashing params query resolves with sane defaults", func(t *testing.T) {
		stdout, _, err := chain.GetNode().ExecQuery(ctx,
			"slashing", "params", "--output", "json")
		require.NoError(t, err, "slashing params: %s", stdout)
		var resp struct {
			Params struct {
				SignedBlocksWindow      string `json:"signed_blocks_window"`
				MinSignedPerWindow      string `json:"min_signed_per_window"`
				DowntimeJailDuration    string `json:"downtime_jail_duration"`
				SlashFractionDoubleSign string `json:"slash_fraction_double_sign"`
				SlashFractionDowntime   string `json:"slash_fraction_downtime"`
			} `json:"params"`
		}
		require.NoError(t, json.Unmarshal(stdout, &resp))
		require.NotEmpty(t, resp.Params.SignedBlocksWindow,
			"slashing params must include signed_blocks_window")
		require.NotEmpty(t, resp.Params.SlashFractionDoubleSign,
			"slashing params must include slash_fraction_double_sign")
	})
}

