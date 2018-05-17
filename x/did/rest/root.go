package rest

import (
	"github.com/cosmos/cosmos-sdk/wire"
	"github.com/gorilla/mux"
	"github.com/ixofoundation/ixo-cosmos/x/did"
	keys "github.com/tendermint/go-crypto/keys"
)

func RegisterRoutes(r *mux.Router, cdc *wire.Codec, kb keys.Keybase, storeName string) {
	r.HandleFunc("/did", CreateDidRequestHandler(storeName, cdc, kb)).Methods("POST")
	r.HandleFunc("/did/{did}", QueryDidDocRequestHandler(storeName, cdc, did.GetDidDocDecoder(cdc))).Methods("GET")
	r.HandleFunc("/did", QueryAllDidsRequestHandler(storeName, cdc, did.GetDidDocDecoder(cdc))).Methods("GET")
}
