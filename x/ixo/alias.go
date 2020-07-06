package ixo

import "github.com/ixofoundation/ixo-blockchain/x/ixo/internal/types"

const (
	IxoNativeToken = types.IxoNativeToken
)

type (
	PubKeyGetter = types.PubKeyGetter

	IxoTx        = types.IxoTx
	IxoSignature = types.IxoSignature
	IxoMsg       = types.IxoMsg
)

var (
	// Auth
	NewDefaultPubKeyGetter           = types.NewDefaultPubKeyGetter
	ProcessSig                       = types.ProcessSig
	NewDefaultAnteHandler            = types.NewDefaultAnteHandler
	ApproximateFeeForTx              = types.ApproximateFeeForTx
	GenerateOrBroadcastMsgs          = types.GenerateOrBroadcastMsgs
	CompleteAndBroadcastTxRest       = types.CompleteAndBroadcastTxRest
	SignAndBroadcastTxFromStdSignMsg = types.SignAndBroadcastTxFromStdSignMsg

	// Types
	IxoDecimals = types.IxoDecimals

	// Tx
	DefaultTxDecoder  = types.DefaultTxDecoder
	NewIxoTxSingleMsg = types.NewIxoTxSingleMsg
)
