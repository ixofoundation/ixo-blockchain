package types

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"

	// "github.com/cosmos/cosmos-sdk/x/auth/ante"
	// authkeeper "github.com/cosmos/cosmos-sdk/x/auth/keeper"

	// bankkeeper "github.com/cosmos/cosmos-sdk/x/bank/keeper"
	"github.com/ixofoundation/ixo-blockchain/x/did/exported"
	"github.com/spf13/pflag"
)

func GenerateOrBroadcastTxCLI(clientCtx client.Context, flagSet *pflag.FlagSet, ixoDid exported.IxoDid, msg sdk.Msg) error {
	txf := tx.NewFactoryCLI(clientCtx, flagSet)
	return GenerateOrBroadcastTxWithFactory(clientCtx, txf, ixoDid, msg)
}

func GenerateOrBroadcastTxWithFactory(clientCtx client.Context, txf tx.Factory, ixoDid exported.IxoDid, msg sdk.Msg) error {
	if clientCtx.GenerateOnly {
		return tx.GenerateOrBroadcastTxWithFactory(clientCtx, txf, msg) // like old PrintUnsignedStdTx
	}

	return tx.BroadcastTx(clientCtx, txf)
}
