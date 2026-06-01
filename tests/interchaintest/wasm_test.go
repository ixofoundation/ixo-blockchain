//go:build interchaintest

package interchaintest

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/strangelove-ventures/interchaintest/v8/chain/cosmos"
	"github.com/stretchr/testify/require"
)

// TestIxoWasm_FullScenario boots ONE chain and walks the wasm contract
// surface end-to-end on a single bond. Each subtest uploads ONLY the
// contract it needs — uploading all 5 in one chain triggers OOM in the
// validator container under amd64 emulation on Apple Silicon.
//
// Subsumes earlier separate Docker bootstraps:
//
//	TestIxoWasm_StoreContracts (collapsed into the per-subtest upload
//	  + ordering checks below)
//	TestIxoWasm_Cw20FullFlow
//	TestIxoWasm_Cw721FullFlow
//
// The deterministic code-ID-1-cw721 = entity.NftContractAddress
// observation is preserved as a comment in the cw721 subtest; we use
// cw721_base + a 2-instance dance there to dodge the entity ante.
func TestIxoWasm_FullScenario(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping interchaintest in -short mode")
	}
	t.Parallel()

	chain, users, ctx := SetupIxoChain(t, 2)
	deployer, recipient := users[0], users[1]

	t.Run("cw20: upload + instantiate + transfer + balance smart-queries", func(t *testing.T) {
		cw20CodeID := UploadContract(t, ctx, chain, deployer, "cw20_base.wasm")
		require.NotEmpty(t, cw20CodeID)

		initMsg := fmt.Sprintf(`{
  "name": "ixo-wasm-test",
  "symbol": "IXOC",
  "decimals": 6,
  "initial_balances": [{"address": %q, "amount": "1000000"}],
  "mint": null
}`, deployer.FormattedAddress())

		// Instantiate twice — see the cw721 subtest below for the
		// "entity ante reserves instance-1's deterministic address"
		// rationale. The cw20 contract hits the same trap (the
		// reservation is keyed off instance_id, not the wasm bytes).
		for i := 0; i < 2; i++ {
			out, err := chain.GetNode().ExecTx(ctx, deployer.KeyName(),
				"wasm", "instantiate", cw20CodeID, initMsg,
				"--label", fmt.Sprintf("ixo-wasm-test-cw20-%d", i),
				"--admin", deployer.FormattedAddress(),
				"--gas", "auto", "--gas-adjustment", "1.5",
			)
			require.NoError(t, err, "instantiate cw20 round %d: %s", i, out)
		}
		contract := lastContractByCode(t, ctx, chain, cw20CodeID)

		require.Equal(t, "1000000",
			queryCw20Balance(t, ctx, chain, contract, deployer.FormattedAddress()),
			"deployer must hold the initial cw20 balance")
		require.Equal(t, "0",
			queryCw20Balance(t, ctx, chain, contract, recipient.FormattedAddress()),
			"recipient must hold zero before transfer")

		transferMsg := fmt.Sprintf(`{"transfer":{"recipient":%q,"amount":"250000"}}`,
			recipient.FormattedAddress())
		out, err := chain.GetNode().ExecTx(ctx, deployer.KeyName(),
			"wasm", "execute", contract, transferMsg,
			"--gas", "auto", "--gas-adjustment", "1.5",
		)
		require.NoError(t, err, "wasm execute transfer: %s", out)

		require.Equal(t, "750000",
			queryCw20Balance(t, ctx, chain, contract, deployer.FormattedAddress()))
		require.Equal(t, "250000",
			queryCw20Balance(t, ctx, chain, contract, recipient.FormattedAddress()))
	})
	if t.Failed() {
		return
	}

	t.Run("cw721_base: upload + instantiate + mint + transfer_nft + owner_of", func(t *testing.T) {
		// We use cw721_base instead of the bundled cw721 because the
		// entity module reserves the deterministic address of
		// code-id-1 instance-1 (the first cw721 instantiate from the
		// keyring's deterministic creator) as
		// `params.NftContractAddress`, and an ante decorator at
		// x/entity/ante/decorators.go blocks any direct wasm execute
		// against that address.
		//
		// Even with cw721_base (a different wasm file → different
		// checksum → different code id), instance-1 from this user
		// can collide with the reserved address in some chain layouts;
		// instantiate twice and use the second.
		cw721CodeID := UploadContract(t, ctx, chain, deployer, "cw721_base.wasm")
		require.NotEmpty(t, cw721CodeID)

		initMsg := fmt.Sprintf(`{"name":"ixo-nft","symbol":"NFT","minter":%q}`,
			deployer.FormattedAddress())
		for i := 0; i < 2; i++ {
			out, err := chain.GetNode().ExecTx(ctx, deployer.KeyName(),
				"wasm", "instantiate", cw721CodeID, initMsg,
				"--label", fmt.Sprintf("ixo-nft-cw721-%d", i),
				"--admin", deployer.FormattedAddress(),
				"--gas", "auto", "--gas-adjustment", "1.5",
			)
			require.NoError(t, err, "instantiate cw721_base round %d: %s", i, out)
		}

		stdout, _, err := chain.GetNode().ExecQuery(ctx,
			"wasm", "list-contract-by-code", cw721CodeID, "--output", "json")
		require.NoError(t, err)
		var list struct {
			Contracts []string `json:"contracts"`
		}
		require.NoError(t, json.Unmarshal(stdout, &list))
		require.GreaterOrEqual(t, len(list.Contracts), 2)
		contract := list.Contracts[len(list.Contracts)-1]

		const tokenID = "ixo-token-1"
		mintMsg := fmt.Sprintf(`{"mint":{"token_id":%q,"owner":%q,"token_uri":"ipfs://test","extension":null}}`,
			tokenID, deployer.FormattedAddress())
		out, err := chain.GetNode().ExecTx(ctx, deployer.KeyName(),
			"wasm", "execute", contract, mintMsg,
			"--gas", "auto", "--gas-adjustment", "1.5",
		)
		require.NoError(t, err, "cw721 mint: %s", out)

		require.Equal(t, deployer.FormattedAddress(),
			queryCw721Owner(t, ctx, chain, contract, tokenID))

		transferMsg := fmt.Sprintf(`{"transfer_nft":{"recipient":%q,"token_id":%q}}`,
			recipient.FormattedAddress(), tokenID)
		out, err = chain.GetNode().ExecTx(ctx, deployer.KeyName(),
			"wasm", "execute", contract, transferMsg,
			"--gas", "auto", "--gas-adjustment", "1.5",
		)
		require.NoError(t, err, "cw721 transfer_nft: %s", out)

		require.Equal(t, recipient.FormattedAddress(),
			queryCw721Owner(t, ctx, chain, contract, tokenID))
	})
}

// queryCw20Balance hits the cw20 contract's `balance` smart query and
// returns the amount as a string.
func queryCw20Balance(t *testing.T, ctx context.Context, chain *cosmos.CosmosChain,
	contract, address string,
) string {
	t.Helper()
	q := fmt.Sprintf(`{"balance":{"address":%q}}`, address)
	stdout, _, err := chain.GetNode().ExecQuery(ctx,
		"wasm", "contract-state", "smart", contract, q, "--output", "json")
	require.NoError(t, err)
	var resp struct {
		Data struct {
			Balance string `json:"balance"`
		} `json:"data"`
	}
	require.NoError(t, json.Unmarshal(stdout, &resp))
	return resp.Data.Balance
}

// queryCw721Owner runs the cw721 contract's owner_of smart query and
// returns the resolved bech32 owner address.
func queryCw721Owner(t *testing.T, ctx context.Context, chain *cosmos.CosmosChain,
	contract, tokenID string,
) string {
	t.Helper()
	q := fmt.Sprintf(`{"owner_of":{"token_id":%q}}`, tokenID)
	stdout, _, err := chain.GetNode().ExecQuery(ctx,
		"wasm", "contract-state", "smart", contract, q, "--output", "json")
	require.NoError(t, err)
	var resp struct {
		Data struct {
			Owner string `json:"owner"`
		} `json:"data"`
	}
	require.NoError(t, json.Unmarshal(stdout, &resp))
	return resp.Data.Owner
}

// lastContractByCode returns the most-recently-instantiated contract
// address for the given code id. Used by the wasm scenario subtests so
// they pick the second-or-later instance — the first instance of any
// code id is reserved by the entity ante decorator.
func lastContractByCode(t *testing.T, ctx context.Context, chain *cosmos.CosmosChain, codeID string) string {
	t.Helper()
	stdout, _, err := chain.GetNode().ExecQuery(ctx,
		"wasm", "list-contract-by-code", codeID, "--output", "json")
	require.NoError(t, err)
	var list struct {
		Contracts []string `json:"contracts"`
	}
	require.NoError(t, json.Unmarshal(stdout, &list))
	require.NotEmpty(t, list.Contracts, "no contracts for code %s", codeID)
	return list.Contracts[len(list.Contracts)-1]
}
