package rest

import (
	"encoding/json"
	"fmt"
	"github.com/ixofoundation/ixo-blockchain/x/did"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/ixofoundation/ixo-blockchain/x/project/internal/keeper"
	"github.com/ixofoundation/ixo-blockchain/x/project/internal/types"
)

type AccDetails struct {
	Did     string  `json:"did" yaml:"did"`
	Account string  `json:"account" yaml:"account"`
	Balance sdk.Int `json:"balance" yaml:"balance"`
}

func registerQueryRoutes(cliCtx context.CLIContext, r *mux.Router) {
	r.HandleFunc("/project/{did}", queryProjectDocRequestHandler(cliCtx)).Methods("GET")
	r.HandleFunc("/projectAccounts/{projectDid}", queryProjectAccountsRequestHandler(cliCtx)).Methods("GET")
	r.HandleFunc("/projectTxs/{projectDid}", queryProjectTxsRequestHandler(cliCtx)).Methods("GET")
	r.HandleFunc("/projectParams", queryParamsRequestHandler(cliCtx)).Methods("GET")
}

func queryProjectDocRequestHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		didAddr := vars["did"]

		key := did.Did(didAddr)
		res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s", types.QuerierRoute,
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
		cliCtx.Codec.MustUnmarshalJSON(res, &projectDoc)

		bz, err := json.Marshal(projectDoc)
		_, _ = w.Write(bz)
	}
}

func queryProjectAccountsRequestHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		projectDid := vars["projectDid"]

		res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s",
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

func queryProjectTxsRequestHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		projectDid := vars["projectDid"]

		res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s",
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
			cliCtx.Codec.MustUnmarshalJSON(res, &txs)
		}

		bz, err := json.Marshal(txs)
		_, _ = w.Write(bz)
	}

}

func queryParamsRequestHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		bz, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", types.QuerierRoute,
			keeper.QueryParams), nil)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(fmt.Sprintf("Couldn't get query data %s", err.Error())))
			return
		}

		var params types.Params
		if err := cliCtx.Codec.UnmarshalJSON(bz, &params); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(fmt.Sprintf("Couldn't Unmarshal data %s", err.Error())))
			return
		}

		rest.PostProcessResponse(w, cliCtx, params)
	}
}
