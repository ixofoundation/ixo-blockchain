package rest

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/rest"

	"github.com/gorilla/mux"
)

func RegisterHandlers(cliCtx client.Context, rtr *mux.Router) {
	r := rest.WithHTTPDeprecationHeaders(rtr)
	registerQueryRoutes(cliCtx, r)
	registerTxHandlers(cliCtx, r)
}
