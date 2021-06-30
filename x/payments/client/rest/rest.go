package rest

import (
	"github.com/cosmos/cosmos-sdk/client"
	clientrest "github.com/cosmos/cosmos-sdk/client/rest"
	"github.com/gorilla/mux"
)

const (
	RestPaymentTemplateId        = "payment_template_id"
	RestPaymentContractId        = "payment_contract_id"
	RestPaymentContractsIdPrefix = "payment_contracts_id_prefix"
	RestSubscriptionId           = "subscription_id"
)

func RegisterHandlers(clientCtx client.Context, rtr *mux.Router) {
	r := clientrest.WithHTTPDeprecationHeaders(rtr)
	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)
}
