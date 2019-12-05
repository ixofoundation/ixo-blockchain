package rest

import (
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/gorilla/mux"
)

func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	r.HandleFunc("/fiatAccount/{address}", QueryFiatRequestHandlerFn(cliCtx)).Methods("GET")

	// r.HandleFunc("/changeBuyerBid", ChangeBuyerBidRequestHandlerFn(cliCtx, kafkaBool, kafkaState)).Methods("POST")
	// r.HandleFunc("/changeSellerBid", ChangeSellerBidRequestHandlerFn(cliCtx, kafkaBool, kafkaState)).Methods("POST")
	// r.HandleFunc("/confirmBuyerBid", ConfirmBuyerBidRequestHandlerFn(cliCtx, kafkaBool, kafkaState)).Methods("POST")
	// r.HandleFunc("/confirmSellerBid", ConfirmSellerBidRequestHandlerFn(cliCtx, kafkaBool, kafkaState)).Methods("POST")
}
