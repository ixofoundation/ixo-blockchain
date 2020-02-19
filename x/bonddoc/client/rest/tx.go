package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/btcsuite/btcutil/base58"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"

	"github.com/ixofoundation/ixo-cosmos/x/bonddoc/internal/types"
	"github.com/ixofoundation/ixo-cosmos/x/ixo"
	"github.com/ixofoundation/ixo-cosmos/x/ixo/sovrin"
)

func registerTxRoutes(cliCtx context.CLIContext, r *mux.Router) {
	r.HandleFunc("/bond", createBondRequestHandler(cliCtx)).Methods("POST")
}

func createBondRequestHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		bondDocParam := r.URL.Query().Get("bondDoc")
		didDocParam := r.URL.Query().Get("didDoc")
		mode := r.URL.Query().Get("mode")
		var bondDoc types.BondDoc
		err := json.Unmarshal([]byte(bondDocParam), &bondDoc)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(fmt.Sprintf("Could not unmarshall bondDoc into struct. Error: %s", err.Error())))

			return
		}

		var didDoc sovrin.SovrinDid
		err = json.Unmarshal([]byte(didDocParam), &didDoc)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = w.Write([]byte(fmt.Sprintf("Could not unmarshall didDoc into struct. Error: %s", err.Error())))

			return
		}

		cliCtx = cliCtx.WithBroadcastMode(mode)
		msg := types.NewCreateBondMsg(bondDoc, didDoc)
		privKey := [64]byte{}
		copy(privKey[:], base58.Decode(didDoc.Secret.SignKey))
		copy(privKey[32:], base58.Decode(didDoc.VerifyKey))

		msgBytes, err := json.Marshal(msg)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(fmt.Sprintf("Could not marshall msg to json. Error: %s", err.Error())))
			return
		}

		signature := ixo.SignIxoMessage(msgBytes, didDoc.Did, privKey)
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

		rest.PostProcessResponse(w, cliCtx, output)
	}
}
