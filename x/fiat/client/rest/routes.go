package rest

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/gorilla/mux"
)

func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	r.HandleFunc("/fiatAccount/{address}", QueryFiatAccountHandlerFn(cliCtx)).Methods("GET")

	r.HandleFunc("/issueFiat", IssueFiatHandlerFunction(cliCtx)).Methods("POST")
	r.HandleFunc("/redeemFiat", RedeemiatHandlerFunction(cliCtx)).Methods("POST")
	r.HandleFunc("/sendFiat", SendiatHandlerFunction(cliCtx)).Methods("POST")
}
