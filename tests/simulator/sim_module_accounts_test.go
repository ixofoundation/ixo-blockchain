// Module-account-coverage tests. Every module that needs a system
// account (mint, distribution, gov, fee_collector, the staking pools,
// the custom ixo module accounts) must register that account in
// `maccPerms` and have a queryable address. If any module's account is
// silently dropped from maccPerms, that module's keeper can't move
// coins and bugs are subtle: txs fail at runtime with cryptic "module
// account not found" errors deep in the bank keeper.
//
// Run with:
//
//	go test -run TestIxoSim_AllModuleAccountsRegistered ./tests/simulator/...
package simulator

import (
	"testing"

	cmtproto "github.com/cometbft/cometbft/proto/tendermint/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/stretchr/testify/require"

	"github.com/ixofoundation/ixo-blockchain/v7/app"
)

// expectedModuleAccounts lists every module that MUST have a registered
// module account. Any drop from this list (in app/modules.go::
// moduleAccountPermissions) breaks fee distribution, inflation, gov
// burn, or the custom ixo modules' bookkeeping.
var expectedModuleAccounts = []string{
	authtypes.FeeCollectorName, // "fee_collector" — receives all tx fees
	"distribution",
	"mint",
	"gov",
	"bonded_tokens_pool",
	"not_bonded_tokens_pool",
	"transfer",                  // ibc transfer escrow
	"liquidstake",               // custom liquid staking
	"bonds_mint_burn_account",         // custom bonds: mints / burns curve tokens
	"bonds_reserve_account",           // custom bonds: holds reserve coins
	"batches_intermediary_account",    // custom bonds: per-batch holding
	"wasm",                      // wasm contract execution
	"smartaccount",              // custom smart accounts
}

// TestIxoSim_AllModuleAccountsRegistered confirms every module account
// the chain claims to need has a registered ADDRESS in maccPerms. Some
// accounts (e.g. ibc transfer escrow) are only persisted to the
// AccountKeeper on first use, so we only check address-resolution here,
// not GetAccount; presence in maccPerms is what guarantees the bank
// keeper can find the address by module name.
func TestIxoSim_AllModuleAccountsRegistered(t *testing.T) {
	a := app.Setup(false)
	_ = a.NewContextLegacy(false, cmtproto.Header{})

	for _, name := range expectedModuleAccounts {
		addr := a.AccountKeeper.GetModuleAddress(name)
		require.NotNil(t, addr,
			"module account %q must be registered (look at app/keepers/keepers.go::maccPerms)", name)
		require.NotEmpty(t, addr,
			"module account %q address must be non-empty", name)
	}
}

// TestIxoSim_BlockedAccountsParseAsValidBech32 checks that every entry
// in app.BlockedAddresses is a parseable bech32 address. A typo in
// BlockedAddresses would silently fail to block transfers, so we verify
// the format up front. We don't require GetAccount lookup because some
// module accounts are created lazily on first use.
func TestIxoSim_BlockedAccountsParseAsValidBech32(t *testing.T) {
	a := app.Setup(false)
	_ = a.NewContextLegacy(false, cmtproto.Header{})

	blocked := a.BlockedAddresses()
	require.NotEmpty(t, blocked, "BlockedAddresses must contain at least the gov account")

	for addr, isBlocked := range blocked {
		if !isBlocked {
			continue
		}
		_, err := a.AccountKeeper.AddressCodec().StringToBytes(addr)
		require.NoError(t, err,
			"BlockedAddresses entry %q must be a valid bech32 address (typos silently fail to block transfers)", addr)
	}
}
