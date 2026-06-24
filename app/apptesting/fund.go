package apptesting

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktestutil "github.com/cosmos/cosmos-sdk/x/bank/testutil"

	minttypes "github.com/ixofoundation/ixo-blockchain/v8/x/mint/types"
)

// FundAcc tops up `acc` with `amounts`, minting fresh coins via the mint
// module so the source has the correct authority.
func (s *KeeperTestHelper) FundAcc(acc sdk.AccAddress, amounts sdk.Coins) {
	s.Require().NoError(banktestutil.FundAccount(s.Ctx, s.App.BankKeeper, acc, amounts))
}

// FundModuleAcc tops up the module account named `moduleName`.
func (s *KeeperTestHelper) FundModuleAcc(moduleName string, amounts sdk.Coins) {
	s.Require().NoError(banktestutil.FundModuleAccount(s.Ctx, s.App.BankKeeper, moduleName, amounts))
}

// MintCoins mints fresh coins into the mint module account. Wraps the
// underlying bank keeper call so tests don't need to import banktypes.
func (s *KeeperTestHelper) MintCoins(coins sdk.Coins) {
	s.Require().NoError(s.App.BankKeeper.MintCoins(s.Ctx, minttypes.ModuleName, coins))
}
