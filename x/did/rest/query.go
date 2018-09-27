package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/wire"
	"github.com/gorilla/mux"
	"github.com/ixofoundation/ixo-cosmos/x/did"
	"github.com/ixofoundation/ixo-cosmos/x/ixo"
)

const storeName = "did"

func registerQueryRoutes(cliCtx context.CLIContext, r *mux.Router, cdc *wire.Codec) {
	r.HandleFunc(
		"/did/{did}",
		queryDidDocRequestHandler(cdc, did.GetDidDocDecoder(cdc))).Methods("GET")
	r.HandleFunc(
		"/did",
		queryAllDidsRequestHandler(cdc, did.GetDidDocDecoder(cdc)),
	).Methods("GET")
}

func queryDidDocRequestHandler(cdc *wire.Codec, decoder did.DidDocDecoder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.NewCLIContext().
			WithCodec(cdc).
			WithLogger(os.Stdout)

		vars := mux.Vars(r)
		didAddr := vars["did"]

		key := ixo.Did(didAddr)

		res, err := ctx.QueryStore([]byte(key), storeName)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("Could't query did. Error: %s", err.Error())))
			return
		}

		// the query will return empty if there is no data for this did
		if len(res) == 0 {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		// decode the value
		didDoc, err := decoder(res)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("Could't parse query result. Result: %s. Error: %s", res, err.Error())))
			return
		}

		// print out whole didDoc
		output, err := json.MarshalIndent(didDoc, "", "  ")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("Could't marshall query result. Error: %s", err.Error())))
			return
		}

		w.Write(output)
	}
}

func queryAllDidsRequestHandler(cdc *wire.Codec, decoder did.DidDocDecoder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.NewCLIContext().
			WithCodec(cdc).
			WithLogger(os.Stdout)

		allKey := "ALL"

		res, err := ctx.QueryStore([]byte(allKey), storeName)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("Could't query did. Error: %s", err.Error())))
			return
		}

		// the query will return empty if there is no data for this did
		if len(res) == 0 {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		// decode the value
		dids := []ixo.Did{}
		err = cdc.UnmarshalBinary(res, &dids)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("Could't parse query result. Result: %s. Error: %s", res, err.Error())))
			return
		}

		// print out whole didDoc
		output, err := json.MarshalIndent(dids, "", "  ")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("Could't marshall query result. Error: %s", err.Error())))
			return
		}

		w.Write(output)
	}

}
