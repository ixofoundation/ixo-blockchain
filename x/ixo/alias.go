package ixo

import "github.com/ixofoundation/ixo-blockchain/x/ixo/internal/types"

const (
	IxoNativeToken = types.IxoNativeToken
)

type (
	PubKeyGetter = types.PubKeyGetter
	IxoMsg       = types.IxoMsg
)

var (
	RegisterCodec = types.RegisterCodec

	// Auth
	NewDefaultPubKeyGetter           = types.NewDefaultPubKeyGetter
	NewDefaultAnteHandler            = types.NewDefaultAnteHandler
	ApproximateFeeForTx              = types.ApproximateFeeForTx
	GenerateOrBroadcastMsgs          = types.GenerateOrBroadcastMsgs
	SignAndBroadcastTxFromStdSignMsg = types.SignAndBroadcastTxFromStdSignMsg
	IxoSigVerificationGasConsumer    = types.IxoSigVerificationGasConsumer

	// Types
	IxoDecimals = types.IxoDecimals
)
