package rest

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/gorilla/mux"
)

const (
	RestFeeId          = "fee_id"
	RestFeeContractId  = "fee_contract_id"
	RestSubscriptionId = "subscription_id"
)

func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	registerQueryRoutes(cliCtx, r)
	registerTxRoutes(cliCtx, r)
}
