package rest

import (
	"encoding/json"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/gorilla/mux"
	"net/http"

	"github.com/ixofoundation/ixo-cosmos/x/ixo"
	"github.com/ixofoundation/ixo-cosmos/x/ixo/sovrin"
	"github.com/ixofoundation/ixo-cosmos/x/treasury/internal/types"
)

func registerTxRoutes(cliCtx context.CLIContext, r *mux.Router) {
	r.HandleFunc("/treasury/send", sendRequestHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/treasury/sendOnBehalfOf", sendOnBehalfOfRequestHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/treasury/mint", mintRequestHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/treasury/burn", burnRequestHandler(cliCtx)).Methods("POST")
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

		privKey := [64]byte{}
		copy(privKey[:], base58.Decode(sovrinDid.Secret.SignKey))
		copy(privKey[32:], base58.Decode(sovrinDid.VerifyKey))

		msgBytes, err := json.Marshal(msg)
		if err != nil {
			panic(err)
		}

		signature := ixo.SignIxoMessage(msgBytes, sovrinDid.Did, privKey)
		tx := ixo.NewIxoTxSingleMsg(msg, signature)

		bz, err := cliCtx.Codec.MarshalJSON(tx)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(fmt.Sprintf("Could not marshall tx to binary. Error: %s", err.Error())))

			return
		}

		res, err := cliCtx.BroadcastTx(bz)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(fmt.Sprintf("Could not broadcast tx. Error: %s", err.Error())))

			return
		}

		output, err := json.MarshalIndent(res, "", "  ")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(err.Error()))

			return
		}

		_, _ = w.Write(output)
	}
}

func sendOnBehalfOfRequestHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		fromDidParam := r.URL.Query().Get("fromDid")
		toDidParam := r.URL.Query().Get("toDid")
		amountParam := r.URL.Query().Get("amount")
		oracleDidParam := r.URL.Query().Get("oracleDid")

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

		msg := types.NewMsgSendOnBehalfOf(fromDidParam, toDidParam, coins, oracleDid)

		privKey := [64]byte{}
		copy(privKey[:], base58.Decode(oracleDid.Secret.SignKey))
		copy(privKey[32:], base58.Decode(oracleDid.VerifyKey))

		msgBytes, err := json.Marshal(msg)
		if err != nil {
			panic(err)
		}

		signature := ixo.SignIxoMessage(msgBytes, oracleDid.Did, privKey)
		tx := ixo.NewIxoTxSingleMsg(msg, signature)

		bz, err := cliCtx.Codec.MarshalJSON(tx)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(fmt.Sprintf("Could not marshall tx to binary. Error: %s", err.Error())))

			return
		}

		res, err := cliCtx.BroadcastTx(bz)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(fmt.Sprintf("Could not broadcast tx. Error: %s", err.Error())))

			return
		}

		output, err := json.MarshalIndent(res, "", "  ")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(err.Error()))

			return
		}

		_, _ = w.Write(output)
	}
}

func mintRequestHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		toDidParam := r.URL.Query().Get("toDid")
		amountParam := r.URL.Query().Get("amount")
		oracleDidParam := r.URL.Query().Get("oracleDid")

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

		msg := types.NewMsgMint(toDidParam, coins, oracleDid)

		privKey := [64]byte{}
		copy(privKey[:], base58.Decode(oracleDid.Secret.SignKey))
		copy(privKey[32:], base58.Decode(oracleDid.VerifyKey))

		msgBytes, err := json.Marshal(msg)
		if err != nil {
			panic(err)
		}

		signature := ixo.SignIxoMessage(msgBytes, oracleDid.Did, privKey)
		tx := ixo.NewIxoTxSingleMsg(msg, signature)

		bz, err := cliCtx.Codec.MarshalJSON(tx)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(fmt.Sprintf("Could not marshall tx to binary. Error: %s", err.Error())))

			return
		}

		res, err := cliCtx.BroadcastTx(bz)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(fmt.Sprintf("Could not broadcast tx. Error: %s", err.Error())))

			return
		}

		output, err := json.MarshalIndent(res, "", "  ")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(err.Error()))

			return
		}

		_, _ = w.Write(output)
	}
}

func burnRequestHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")

		fromDidParam := r.URL.Query().Get("fromDid")
		amountParam := r.URL.Query().Get("amount")
		oracleDidParam := r.URL.Query().Get("oracleDid")

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

		msg := types.NewMsgBurn(fromDidParam, coins, oracleDid)

		privKey := [64]byte{}
		copy(privKey[:], base58.Decode(oracleDid.Secret.SignKey))
		copy(privKey[32:], base58.Decode(oracleDid.VerifyKey))

		msgBytes, err := json.Marshal(msg)
		if err != nil {
			panic(err)
		}

		signature := ixo.SignIxoMessage(msgBytes, oracleDid.Did, privKey)
		tx := ixo.NewIxoTxSingleMsg(msg, signature)

		bz, err := cliCtx.Codec.MarshalJSON(tx)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(fmt.Sprintf("Could not marshall tx to binary. Error: %s", err.Error())))

			return
		}

		res, err := cliCtx.BroadcastTx(bz)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(fmt.Sprintf("Could not broadcast tx. Error: %s", err.Error())))

			return
		}

		output, err := json.MarshalIndent(res, "", "  ")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(err.Error()))

			return
		}

		_, _ = w.Write(output)
	}
}
