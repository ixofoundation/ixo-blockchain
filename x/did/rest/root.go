package rest

import (
	"github.com/cosmos/cosmos-sdk/wire"
	"github.com/gorilla/mux"
)

func RegisterRoutes(r *mux.Router, cdc *wire.Codec, storeName string) {
	r.HandleFunc("/did/{did}", QueryDidDocRequestHandler(storeName, cdc, did.GetDidDocDecoder(cdc))).Methods("GET")
}
