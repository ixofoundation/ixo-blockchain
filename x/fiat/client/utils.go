package client

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	fiatTypes "github.com/ixofoundation/ixo-cosmos/x/fiat/internal/types"
)

func BuildIssueFiatMsg(issuerAddress, toAddress sdk.AccAddress, tranasctionID string, transactionAmount int64) sdk.Msg {
	issueFiat := fiatTypes.NewIssueFiat(issuerAddress, toAddress, tranasctionID, transactionAmount)
	msg := fiatTypes.NewMsgIssueFiats([]fiatTypes.IssueFiat{issueFiat})
	return msg
}

func BuildRedeemFiatMsg(redeemerAddress, issuerAddress sdk.AccAddress, amount int64) sdk.Msg {
	redeemFiat := fiatTypes.NewRedeemFiat(redeemerAddress, issuerAddress, amount)
	msg := fiatTypes.NewMsgRedeemFiats([]fiatTypes.RedeemFiat{redeemFiat})
	return msg
}

func BuildSendFiatMsg(fromAddress, toAddress sdk.AccAddress, amount int64) sdk.Msg {
	sendFiat := fiatTypes.NewSendFiat(fromAddress, toAddress, amount)
	msg := fiatTypes.NewMsgSendFiats([]fiatTypes.SendFiat{sendFiat})
	return msg
}
