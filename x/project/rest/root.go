package rest

import (
	"github.com/cosmos/cosmos-sdk/wire"
	"github.com/gorilla/mux"
	"github.com/ixofoundation/ixo-cosmos/x/project"
	keys "github.com/tendermint/go-crypto/keys"
)

// RegisterRoutes - Central function to define routes that get registered by the main application
func RegisterRoutes(r *mux.Router, cdc *wire.Codec, kb keys.Keybase, storeName string) {
	r.HandleFunc("/project", CreateProjectRequestHandler(storeName, cdc, kb)).Methods("POST")
	r.HandleFunc("/project/{did}", QueryProjectDocRequestHandler(storeName, cdc, project.GetProjectDocDecoder(cdc))).Methods("GET")
	//	r.HandleFunc("/project", QueryAllProjectDocRequestHandler(storeName, cdc, project.GetProjectDocDecoder(cdc))).Methods("GET")
}
