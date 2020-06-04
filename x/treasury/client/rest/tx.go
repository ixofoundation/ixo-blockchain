package rest

import (
	"encoding/json"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
	"net/http"

	"github.com/ixofoundation/ixo-blockchain/x/ixo"
	"github.com/ixofoundation/ixo-blockchain/x/ixo/sovrin"
	"github.com/ixofoundation/ixo-blockchain/x/treasury/internal/types"
)

func registerTxRoutes(cliCtx context.CLIContext, r *mux.Router) {
	r.HandleFunc("/treasury/send", sendRequestHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/treasury/oracleTransfer", oracleTransferRequestHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/treasury/oracleMint", oracleMintRequestHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/treasury/oracleBurn", oracleBurnRequestHandler(cliCtx)).Methods("POST")
}

func sendRequestHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		toDidParam := r.URL.Query().Get("toDid")
		amountParam := r.URL.Query().Get("amount")
		sovrinDidParam := r.URL.Query().Get("sovrinDid")

		mode := r.URL.Query().Get("mode")
		cliCtx = cliCtx.WithBroadcastMode(mode)

		coins, err := sdk.ParseCoins(amountParam)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(err.Error()))

			return
		}

		sovrinDid := sovrin.SovrinDid{}
		err = json.Unmarshal([]byte(sovrinDidParam), &sovrinDid)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(fmt.Sprintf("Could not unmarshall didDoc into struct. Error: %s", err.Error())))

			return
		}

		msg := types.NewMsgSend(toDidParam, coins, sovrinDid)

		output, err := ixo.SignAndBroadcastRest(cliCtx, msg, sovrinDid)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(err.Error()))
			return
		}

		rest.PostProcessResponse(w, cliCtx, output)
	}
}

func oracleTransferRequestHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		fromDidParam := r.URL.Query().Get("fromDid")
		toDidParam := r.URL.Query().Get("toDid")
		amountParam := r.URL.Query().Get("amount")
		oracleDidParam := r.URL.Query().Get("oracleDid")
		proofParam := r.URL.Query().Get("proof")

		mode := r.URL.Query().Get("mode")
		cliCtx = cliCtx.WithBroadcastMode(mode)

		coins, err := sdk.ParseCoins(amountParam)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(err.Error()))

			return
		}

		oracleDid := sovrin.SovrinDid{}
		err = json.Unmarshal([]byte(oracleDidParam), &oracleDid)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(fmt.Sprintf("Could not unmarshall didDoc into struct. Error: %s", err.Error())))

			return
		}

		msg := types.NewMsgOracleTransfer(fromDidParam, toDidParam, coins, oracleDid, proofParam)

		output, err := ixo.SignAndBroadcastRest(cliCtx, msg, oracleDid)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(err.Error()))
			return
		}

		rest.PostProcessResponse(w, cliCtx, output)
	}
}

func oracleMintRequestHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		toDidParam := r.URL.Query().Get("toDid")
		amountParam := r.URL.Query().Get("amount")
		oracleDidParam := r.URL.Query().Get("oracleDid")
		proofParam := r.URL.Query().Get("proof")

		mode := r.URL.Query().Get("mode")
		cliCtx = cliCtx.WithBroadcastMode(mode)

		coins, err := sdk.ParseCoins(amountParam)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(err.Error()))

			return
		}

		oracleDid := sovrin.SovrinDid{}
		err = json.Unmarshal([]byte(oracleDidParam), &oracleDid)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(fmt.Sprintf("Could not unmarshall didDoc into struct. Error: %s", err.Error())))

			return
		}

		msg := types.NewMsgOracleMint(toDidParam, coins, oracleDid, proofParam)

		output, err := ixo.SignAndBroadcastRest(cliCtx, msg, oracleDid)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(err.Error()))
			return
		}

		rest.PostProcessResponse(w, cliCtx, output)
	}
}

func oracleBurnRequestHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		fromDidParam := r.URL.Query().Get("fromDid")
		amountParam := r.URL.Query().Get("amount")
		oracleDidParam := r.URL.Query().Get("oracleDid")
		proofParam := r.URL.Query().Get("proof")

		mode := r.URL.Query().Get("mode")
		cliCtx = cliCtx.WithBroadcastMode(mode)

		coins, err := sdk.ParseCoins(amountParam)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(err.Error()))

			return
		}

		oracleDid := sovrin.SovrinDid{}
		err = json.Unmarshal([]byte(oracleDidParam), &oracleDid)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(fmt.Sprintf("Could not unmarshall didDoc into struct. Error: %s", err.Error())))

			return
		}

		msg := types.NewMsgOracleBurn(fromDidParam, coins, oracleDid, proofParam)

		output, err := ixo.SignAndBroadcastRest(cliCtx, msg, oracleDid)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(err.Error()))
			return
		}

		rest.PostProcessResponse(w, cliCtx, output)
	}
}
