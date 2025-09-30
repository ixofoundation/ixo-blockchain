package app

import (
	txsigning "cosmossdk.io/x/tx/signing"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"

	smartaccountkeeper "github.com/ixofoundation/ixo-blockchain/v6/x/smart-account/keeper"
	smartaccountpost "github.com/ixofoundation/ixo-blockchain/v6/x/smart-account/post"
)

func NewPostHandler(
	cdc codec.Codec,
	smartAccountKeeper *smartaccountkeeper.Keeper,
	accountKeeper authkeeper.AccountKeeper,
	sigModeHandler *txsigning.HandlerMap,
) sdk.PostHandler {
	return sdk.ChainPostDecorators(
		smartaccountpost.NewAuthenticatorPostDecorator(
			cdc,
			smartAccountKeeper,
			accountKeeper,
			sigModeHandler,
			// Add an empty handler here to enable a circuit breaker pattern
			sdk.ChainPostDecorators(sdk.Terminator{}), //nolint
		),
	)
}
