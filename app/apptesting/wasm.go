package apptesting

import (
	"os"

	wasmkeeper "github.com/CosmWasm/wasmd/x/wasm/keeper"
	wasmtypes "github.com/CosmWasm/wasmd/x/wasm/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// StoreContractCode reads a wasm artefact from disk, stores it on-chain via the
// wasm keeper as `creator`, and returns the assigned codeID.
//
// Tests that need real wasm execution (entity NFT contract, claims contract
// payments) should usually go through interchaintest instead — this helper is
// kept for the rare case where in-process wasm is justified.
func (s *KeeperTestHelper) StoreContractCode(creator sdk.AccAddress, path string) uint64 {
	wasmCode, err := os.ReadFile(path)
	s.Require().NoError(err)

	contractKeeper := wasmkeeper.NewDefaultPermissionKeeper(s.App.WasmKeeper)
	codeID, _, err := contractKeeper.Create(s.Ctx, creator, wasmCode, &wasmtypes.AccessConfig{
		Permission: wasmtypes.AccessTypeEverybody,
	})
	s.Require().NoError(err)
	return codeID
}

// InstantiateContract calls Instantiate with the given init message. Returns
// the contract address.
func (s *KeeperTestHelper) InstantiateContract(creator, admin sdk.AccAddress, codeID uint64, label string, initMsg []byte, funds sdk.Coins) sdk.AccAddress {
	contractKeeper := wasmkeeper.NewDefaultPermissionKeeper(s.App.WasmKeeper)
	addr, _, err := contractKeeper.Instantiate(s.Ctx, codeID, creator, admin, initMsg, label, funds)
	s.Require().NoError(err)
	return addr
}
