package rest

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/ixofoundation/ixo-blockchain/x/did/exported"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/gorilla/mux"

	"github.com/ixofoundation/ixo-blockchain/x/did/keeper"
	"github.com/ixofoundation/ixo-blockchain/x/did/types"
)

func registerQueryRoutes(clientCtx client.Context, r *mux.Router) {
	// The .* is necessary so that a slash in the did gets included as part of the did
	r.HandleFunc("/didToAddr/{did:.*}", queryAddressFromDidRequestHandlerFn(clientCtx)).Methods("GET")
	r.HandleFunc("/pubKeyToAddr/{pubKey}", queryAddressFromBase58EncodedPubkeyRequestHandlerFn(clientCtx)).Methods("GET")
	r.HandleFunc("/did/{did}", queryDidDocRequestHandlerFn(clientCtx)).Methods("GET")
	r.HandleFunc("/did", queryAllDidsRequestHandlerFn(clientCtx)).Methods("GET")
	r.HandleFunc("/allDidDocs", queryAllDidDocsRequestHandlerFn(clientCtx)).Methods("GET")
}

func queryAddressFromBase58EncodedPubkeyRequestHandlerFn(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)

		if !types.IsValidPubKey(vars["pubKey"]) {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("input is not a valid base-58 encoded pubKey"))
			return
		}

		accAddress := exported.VerifyKeyToAddr(vars["pubKey"])

		rest.PostProcessResponse(w, clientCtx, accAddress)
	}
}

func queryAddressFromDidRequestHandlerFn(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		didAddr := vars["did"]
		key := exported.Did(didAddr)
		res, _, err := clientCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s", types.QuerierRoute,
			keeper.QueryDidDoc, key), nil)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(err.Error()))
			return
		}
		if len(res) == 0 {
			w.WriteHeader(http.StatusNoContent)
			_, _ = w.Write([]byte("No data for respected did address."))
			return
		}

		var didDoc types.BaseDidDoc
		clientCtx.LegacyAmino.MustUnmarshalJSON(res, &didDoc)
		addressFromDid := didDoc.Address()

		rest.PostProcessResponse(w, clientCtx, addressFromDid)
	}
}

func queryDidDocRequestHandlerFn(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		didAddr := vars["did"]
		key := exported.Did(didAddr)
		res, _, err := clientCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s", types.QuerierRoute,
			keeper.QueryDidDoc, key), nil)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(fmt.Sprintf("Could't query did. Error: %s", err.Error())))
			return
		}
		if len(res) == 0 {
			w.WriteHeader(http.StatusNoContent)
			_, _ = w.Write([]byte("No data for respected did address."))
			return
		}

		var didDoc types.BaseDidDoc
		clientCtx.LegacyAmino.MustUnmarshalJSON(res, &didDoc)


		rest.PostProcessResponse(w, clientCtx, didDoc)
	}
}

func queryAllDidsRequestHandlerFn(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		res, _, err := clientCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", types.QuerierRoute,
			keeper.QueryAllDids), nil)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(fmt.Sprintf("Could't query did. Error: %s", err.Error())))
			return
		}

		if len(res) == 0 {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		var dids []exported.Did
		clientCtx.LegacyAmino.MustUnmarshalJSON(res, &dids)

		rest.PostProcessResponse(w, clientCtx, dids)
	}
}

func queryAllDidDocsRequestHandlerFn(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		res, _, err := clientCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", types.QuerierRoute,
			keeper.QueryAllDidDocs), nil)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(fmt.Sprintf("Could't query did. Error: %s", err.Error())))
			return
		}

		if len(res) == 0 {
			w.WriteHeader(http.StatusNoContent)
			_, _ = w.Write([]byte("No data present."))
			return
		}

		var didDocs []types.BaseDidDoc
		clientCtx.LegacyAmino.MustUnmarshalJSON(res, &didDocs)

		rest.PostProcessResponse(w, clientCtx, didDocs)
	}
}
