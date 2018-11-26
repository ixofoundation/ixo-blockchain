package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	base58 "github.com/btcsuite/btcutil/base58"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/crypto/keys"
	"github.com/cosmos/cosmos-sdk/wire"
	"github.com/gorilla/mux"
	"github.com/ixofoundation/ixo-cosmos/x/ixo"
	"github.com/ixofoundation/ixo-cosmos/x/ixo/sovrin"
	"github.com/ixofoundation/ixo-cosmos/x/project"
)

func registerTxRoutes(cliCtx context.CLIContext, r *mux.Router, cdc *wire.Codec, kb keys.Keybase) {
	r.HandleFunc(
		"/project",
		createProjectRequestHandler(cdc, kb, cliCtx),
	).Methods("POST")
}

//CreateProjectRequestHandler create project handler
func createProjectRequestHandler(cdc *wire.Codec, kb keys.Keybase, cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.NewCLIContext().
			WithCodec(cdc).
			WithLogger(os.Stdout)

		// collect data
		projectDocParam := r.URL.Query().Get("projectDoc")
		didDocParam := r.URL.Query().Get("didDoc")

		projectDoc := project.ProjectDoc{}
		err := json.Unmarshal([]byte(projectDocParam), &projectDoc)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("Could not unmarshall projectDoc into struct. Error: %s", err.Error())))
			return
		}

		sovrinDid := sovrin.SovrinDid{}
		sovrinErr := json.Unmarshal([]byte(didDocParam), &sovrinDid)
		if sovrinErr != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("Could not unmarshall didDoc into struct. Error: %s", err.Error())))
			return
		}

		// create the message
		msg := project.NewCreateProjectMsg(projectDoc, sovrinDid)

		// Force the length to 64
		privKey := [64]byte{}
		copy(privKey[:], base58.Decode(sovrinDid.Secret.SignKey))
		copy(privKey[32:], base58.Decode(sovrinDid.VerifyKey))

		msgBytes, err := json.Marshal(msg)
		if err != nil {
			panic(err)
		}
		//Create the Signature
		signature := ixo.SignIxoMessage(msgBytes, sovrinDid.Did, privKey)
		tx := ixo.NewIxoTxSingleMsg(msg, signature)

		fmt.Println("*******TRANSACTION******* \n", tx.String())

		bz, err := cdc.MarshalJSON(tx)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("Could not marshall tx to binary. Error: %s", err.Error())))
			return
		}

		res, err := ctx.BroadcastTx(bz)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("Could not broadcast tx. Error: %s", err.Error())))
			return
		}

		output, err := json.MarshalIndent(res, "", "  ")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		w.Write(output)
	}
}
