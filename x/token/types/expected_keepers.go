package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	iidtypes "github.com/ixofoundation/ixo-blockchain/v5/x/iid/types"
)

type IidKeeper interface {
	GetDidDocument(ctx sdk.Context, key []byte) (iidtypes.IidDocument, bool)
}

type WasmKeeper interface {
	Instantiate(
		ctx sdk.Context,
		codeID uint64,
		creator, admin sdk.AccAddress,
		initMsg []byte,
		label string,
		deposit sdk.Coins,
	) (sdk.AccAddress, []byte, error)
	Execute(ctx sdk.Context, contractAddress sdk.AccAddress, caller sdk.AccAddress, msg []byte, coins sdk.Coins) ([]byte, error)
}
