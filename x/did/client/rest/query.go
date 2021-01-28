package rest

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/ixofoundation/ixo-blockchain/x/did/exported"
	"net/http"

	//"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/gorilla/mux"

	"github.com/ixofoundation/ixo-blockchain/x/did/internal/keeper"
	"github.com/ixofoundation/ixo-blockchain/x/did/internal/types"
)

func registerQueryRoutes(cliCtx /*context*/client.Context, r *mux.Router) {
	// The .* is necessary so that a slash in the did gets included as part of the did
	r.HandleFunc("/didToAddr/{did:.*}", queryAddressFromDidRequestHandlerFn(cliCtx)).Methods("GET")
	r.HandleFunc("/pubKeyToAddr/{pubKey}", queryAddressFromBase58EncodedPubkeyRequestHandlerFn(cliCtx)).Methods("GET")
	r.HandleFunc("/did/{did}", queryDidDocRequestHandlerFn(cliCtx)).Methods("GET")
	r.HandleFunc("/did", queryAllDidsRequestHandlerFn(cliCtx)).Methods("GET")
	r.HandleFunc("/allDidDocs", queryAllDidDocsRequestHandlerFn(cliCtx)).Methods("GET")
}

func queryAddressFromBase58EncodedPubkeyRequestHandlerFn(cliCtx /*context*/client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)

		if !types.IsValidPubKey(vars["pubKey"]) {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte("input is not a valid base-58 encoded pubKey"))
			return
		}

		accAddress := exported.VerifyKeyToAddr(vars["pubKey"])

		rest.PostProcessResponse(w, cliCtx, accAddress)
	}
}

func queryAddressFromDidRequestHandlerFn(cliCtx /*context*/client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		didAddr := vars["did"]
		key := exported.Did(didAddr)
		res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s", types.QuerierRoute,
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
		//cliCtx.Codec.MustUnmarshalJSON(res, &didDoc)
		cliCtx.LegacyAmino.MustUnmarshalJSON(res, &didDoc)
		addressFromDid := didDoc.Address()

		rest.PostProcessResponse(w, cliCtx, addressFromDid)
	}
}

func queryDidDocRequestHandlerFn(cliCtx /*context*/client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		didAddr := vars["did"]
		key := exported.Did(didAddr)
		res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s", types.QuerierRoute,
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
		//cliCtx.Codec.MustUnmarshalJSON(res, &didDoc)
		cliCtx.LegacyAmino.MustUnmarshalJSON(res, &didDoc)


		rest.PostProcessResponse(w, cliCtx, didDoc)
	}
}

func queryAllDidsRequestHandlerFn(cliCtx /*context*/client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", types.QuerierRoute,
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
		//cliCtx.Codec.MustUnmarshalJSON(res, &dids)
		cliCtx.LegacyAmino.MustUnmarshalJSON(res, &dids)

		rest.PostProcessResponse(w, cliCtx, dids)
	}
}

func queryAllDidDocsRequestHandlerFn(cliCtx /*context*/client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", types.QuerierRoute,
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
		//cliCtx.Codec.MustUnmarshalJSON(res, &didDocs)
		cliCtx.LegacyAmino.MustUnmarshalJSON(res, &didDocs)

		rest.PostProcessResponse(w, cliCtx, didDocs)
	}
}
