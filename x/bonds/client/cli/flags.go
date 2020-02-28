package cli

import (
	"github.com/ixofoundation/ixo-cosmos/x/bonds/internal/types"
	flag "github.com/spf13/pflag"
)

const (
	FlagToken                  = "token"
	FlagName                   = "name"
	FlagDescription            = "description"
	FlagFunctionType           = "function-type"
	FlagFunctionParameters     = "function-parameters"
	FlagReserveTokens          = "reserve-tokens"
	FlagTxFeePercentage        = "tx-fee-percentage"
	FlagExitFeePercentage      = "exit-fee-percentage"
	FlagFeeAddress             = "fee-address"
	FlagMaxSupply              = "max-supply"
	FlagOrderQuantityLimits    = "order-quantity-limits"
	FlagSanityRate             = "sanity-rate"
	FlagSanityMarginPercentage = "sanity-margin-percentage"
	FlagAllowSells             = "allow-sells"
	FlagBatchBlocks            = "batch-blocks"
	FlagBondDid                = "bond-did"
)

var (
	fsBondGeneral = flag.NewFlagSet("", flag.ContinueOnError)
	fsBondCreate  = flag.NewFlagSet("", flag.ContinueOnError)
	fsBondEdit    = flag.NewFlagSet("", flag.ContinueOnError)
)

func init() {

	fsBondGeneral.String(FlagToken, "", "The bond's token")

	fsBondCreate.String(FlagName, "", "The bond's name")
	fsBondCreate.String(FlagDescription, "", "The bond's description")
	fsBondCreate.String(FlagFunctionType, "", "The type of function that the bond will be")
	fsBondCreate.String(FlagFunctionParameters, "", "The parameters that will define the function")
	fsBondCreate.String(FlagReserveTokens, "", "The token(s) that will serve as the reserve token(s)")
	fsBondCreate.String(FlagTxFeePercentage, "", "The percentage fee charged on buys and sells")
	fsBondCreate.String(FlagExitFeePercentage, "", "The percentage fee charged on sells")
	fsBondCreate.String(FlagFeeAddress, "", "The address that will hold any charged fees")
	fsBondCreate.String(FlagMaxSupply, "", "The maximum supply that can be achieved")
	fsBondCreate.String(FlagOrderQuantityLimits, "", "The max number of tokens bought/sold/swapped per order")
	fsBondCreate.String(FlagSanityRate, "", "For swappers, this is the typical t1 per t2 rate")
	fsBondCreate.String(FlagSanityMarginPercentage, "", "For swappers, this is the acceptable deviation from the sanity rate")
	fsBondCreate.String(FlagAllowSells, "", "Whether or not sells will be allowed")
	fsBondCreate.String(FlagBatchBlocks, "", "The duration in terms of blocks of each orders batch")
	fsBondCreate.String(FlagBondDid, "", "Bond's Sovrin DID")

	fsBondEdit.String(FlagName, types.DoNotModifyField, "The bond's name")
	fsBondEdit.String(FlagDescription, types.DoNotModifyField, "The bond's description")
	fsBondEdit.String(FlagOrderQuantityLimits, types.DoNotModifyField, "The max number of tokens bought/sold/swapped per order")
	fsBondEdit.String(FlagSanityRate, types.DoNotModifyField, "For swappers, this is the typical t1 per t2 rate")
	fsBondEdit.String(FlagSanityMarginPercentage, types.DoNotModifyField, "For swappers, this is the acceptable deviation from the sanity rate")
	fsBondEdit.String(FlagBondDid, "", "Bond's Sovrin DID")
}
