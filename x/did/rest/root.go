package rest

import (
	"github.com/cosmos/cosmos-sdk/wire"
	"github.com/gorilla/mux"
	"github.com/ixofoundation/ixo-cosmos/x/did"
)

func RegisterRoutes(r *mux.Router, cdc *wire.Codec, storeName string) {
	r.HandleFunc("/did/{did}", QueryDidDocRequestHandler(storeName, cdc, did.GetDidDocDecoder(cdc))).Methods("GET")
	r.HandleFunc("/did", QueryAllDidsRequestHandler(storeName, cdc, did.GetDidDocDecoder(cdc))).Methods("GET")
}
