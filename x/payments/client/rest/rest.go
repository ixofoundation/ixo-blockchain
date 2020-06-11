package rest

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/gorilla/mux"
)

const (
	RestPaymentTemplateId = "payment_template_id"
	RestPaymentContractId = "payment_contract_id"
	RestSubscriptionId    = "subscription_id"
)

func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	registerQueryRoutes(cliCtx, r)
	registerTxRoutes(cliCtx, r)
}
