package cli

import (
	flag "github.com/spf13/pflag"
)

// noLint
const (
	FlagTo                = "toAddress"
	FlagRedeemerAddress   = "redeemerAddress"
	FlagTransactionAmount = "transactionAmount"
	FlagTransactionID     = "transactionID"
	FlagAmount            = "amount"
	FlagAddress           = "address"
)

var (
	fsTo                = flag.NewFlagSet("", flag.ContinueOnError)
	fsRedeemerAddress   = flag.NewFlagSet("", flag.ContinueOnError)
	fsTransactionAmount = flag.NewFlagSet("", flag.ContinueOnError)
	fsTransactionID     = flag.NewFlagSet("", flag.ContinueOnError)
	fsAmount            = flag.NewFlagSet("", flag.ContinueOnError)
	fsAddress           = flag.NewFlagSet("", flag.ContinueOnError)
)

func init() {
	fsTo.String(FlagTo, "", "Address to send or issue fiats")
	fsRedeemerAddress.String(FlagRedeemerAddress, "", "Address from which fiats  to be redeemed")
	fsTransactionAmount.String(FlagTransactionAmount, "", "Amount to be issued.")
	fsTransactionID.String(FlagTransactionID, "", "Transaction ID from the bank.")
	fsAmount.String(FlagAmount, "", "Amount to send or issue fiats")
	fsAddress.String(FlagAddress, "", "Address to query")
}
