package ixo

import "github.com/ixofoundation/ixo-blockchain/x/ixo/internal/types"

const (
	IxoNativeToken = "uixo"
)

type (
	PubKeyGetter = types.PubKeyGetter
	IxoMsg       = types.IxoMsg
)

var (
	RegisterCodec = types.RegisterCodec

	// Auth
	NewDefaultAnteHandler            = types.NewDefaultAnteHandler
	ProcessSig                       = types.ProcessSig
	NewSetPubKeyDecorator            = types.NewSetPubKeyDecorator
	NewDeductFeeDecorator            = types.NewDeductFeeDecorator
	NewSigGasConsumeDecorator        = types.NewSigGasConsumeDecorator
	NewSigVerificationDecorator      = types.NewSigVerificationDecorator
	NewIncrementSequenceDecorator    = types.NewIncrementSequenceDecorator
	ApproximateFeeForTx              = types.ApproximateFeeForTx
	GenerateOrBroadcastMsgs          = types.GenerateOrBroadcastMsgs
	SignAndBroadcastTxFromStdSignMsg = types.SignAndBroadcastTxFromStdSignMsg
	IxoSigVerificationGasConsumer    = types.IxoSigVerificationGasConsumer
)
