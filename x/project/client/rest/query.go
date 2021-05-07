package rest

import (
	"encoding/json"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/ixofoundation/ixo-blockchain/x/did"
	"github.com/ixofoundation/ixo-blockchain/x/project/keeper"
	"github.com/ixofoundation/ixo-blockchain/x/project/types"
	"net/http"

	//"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/types/rest"
)

type AccDetails struct {
	Did     string  `json:"did" yaml:"did"`
	Account string  `json:"account" yaml:"account"`
	Balance sdk.Int `json:"balance" yaml:"balance"`
}

func registerQueryRoutes(clientCtx client.Context, r *mux.Router) {
	r.HandleFunc("/project/{did}", queryProjectDocRequestHandler(clientCtx)).Methods("GET")
	r.HandleFunc("/projectAccounts/{projectDid}", queryProjectAccountsRequestHandler(clientCtx)).Methods("GET")
	r.HandleFunc("/projectTxs/{projectDid}", queryProjectTxsRequestHandler(clientCtx)).Methods("GET")
	r.HandleFunc("/projectParams", queryParamsRequestHandler(clientCtx)).Methods("GET")
}

func queryProjectDocRequestHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		didAddr := vars["did"]

		key := did.Did(didAddr)
		res, _, err := clientCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s", types.QuerierRoute,
			keeper.QueryProjectDoc, key), nil)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(fmt.Sprintf("Could't query did. Error: %s", err.Error())))
			return
		}

		if len(res) == 0 {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		var projectDoc types.ProjectDoc
		clientCtx.LegacyAmino.MustUnmarshalJSON(res, &projectDoc)

		bz, err := json.Marshal(projectDoc)
		_, _ = w.Write(bz)
	}
}

func queryProjectAccountsRequestHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		projectDid := vars["projectDid"]

		res, _, err := clientCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s",
			types.QuerierRoute, keeper.QueryProjectAccounts, projectDid), nil)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(fmt.Sprintf("Could't query did. Error: %s", err.Error())))
			return
		}

		if len(res) == 0 {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		var f interface{}
		err = json.Unmarshal(res, &f)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(fmt.Sprintf("Could't parse query result. Result: %s. Error: %s", res, err.Error())))
			return
		}

		accMap := f.(map[string]interface{})
		bz, err := json.Marshal(accMap)
		_, _ = w.Write(bz)
	}

}

func queryProjectTxsRequestHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		projectDid := vars["projectDid"]

		res, _, err := clientCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s",
			types.QuerierRoute, keeper.QueryProjectTx, projectDid), nil)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(fmt.Sprintf("Could't query did. Error: %s", err.Error())))
			return
		}

		var txs []types.WithdrawalInfoDoc
		if len(res) == 0 {
			w.WriteHeader(http.StatusNotFound)
			return
		} else {
			clientCtx.LegacyAmino.MustUnmarshalJSON(res, &txs)
		}

		bz, err := json.Marshal(txs)
		_, _ = w.Write(bz)
	}

}

func queryParamsRequestHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		bz, _, err := clientCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", types.QuerierRoute,
			keeper.QueryParams), nil)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(fmt.Sprintf("Couldn't get query data %s", err.Error())))
			return
		}

		var params types.Params
		if err := clientCtx.LegacyAmino.UnmarshalJSON(bz, &params); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(fmt.Sprintf("Couldn't Unmarshal data %s", err.Error())))
			return
		}

		rest.PostProcessResponse(w, clientCtx, params)
	}
}
