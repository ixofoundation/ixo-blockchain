package rest

import (
	"github.com/cosmos/cosmos-sdk/client"
	clientrest "github.com/cosmos/cosmos-sdk/client/rest"
	"github.com/gorilla/mux"
)

func RegisterHandlers(clientCtx client.Context, rtr *mux.Router) {
	r := clientrest.WithHTTPDeprecationHeaders(rtr)
	registerQueryRoutes(clientCtx, r)
	registerTxHandlers(clientCtx, r)
}
