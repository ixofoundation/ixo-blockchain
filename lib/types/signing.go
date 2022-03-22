package types

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"

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

type PubKeyGetter func(ctx sdk.Context, msg IxoMsg) (cryptotypes.PubKey, error)

func SignAndBroadcastTxFromStdSignMsg(clientCtx client.Context,
	msg sdk.Msg, ixoDid exported.IxoDid, flagSet *pflag.FlagSet) (*sdk.TxResponse, error) {

	// txf := tx.NewFactoryCLI(clientCtx, flagSet)
	// txf = txf.WithFees(fees).WithGasPrices("").WithGas(0)

	// tx, err := tx.BuildUnsignedTx(txf, msg)
	// if err != nil {
	// 	return nil, err
	// }

	// if !clientCtx.SkipConfirm {
	// 	out, err := clientCtx.TxConfig.TxJSONEncoder()(tx.GetTx())
	// 	if err != nil {
	// 		return nil, err
	// 	}

	// 	_, _ = fmt.Fprintf(os.Stderr, "%s\n\n", out)

	// 	buf := bufio.NewReader(os.Stdin)
	// 	ok, err := input.GetConfirmation("confirm transaction before signing and broadcasting", buf, os.Stderr)

	// 	if err != nil || !ok {
	// 		_, _ = fmt.Fprintf(os.Stderr, "%s\n", "cancelled transaction")
	// 		return nil, err
	// 	}
	// }

	// // err = Sign(txf, clientCtx, tx, true, ixoDid)
	// // if err != nil {
	// // 	return nil, err
	// // }

	// txBytes, err := clientCtx.TxConfig.TxEncoder()(tx.GetTx())
	// if err != nil {
	// 	return nil, err
	// }

	// // broadcast to a Tendermint node
	// res, err := clientCtx.BroadcastTx(txBytes)
	// if err != nil {
	// 	return &sdk.TxResponse{}, err
	// }

	return nil, nil
}
