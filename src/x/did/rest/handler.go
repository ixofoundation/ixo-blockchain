package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"

	base58 "github.com/btcsuite/btcutil/base58"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/wire"
	did "github.com/ixofoundation/ixo-cosmos/x/did"
	"github.com/ixofoundation/ixo-cosmos/x/ixo"
	"github.com/ixofoundation/ixo-cosmos/x/ixo/sovrin"
	"github.com/tendermint/go-crypto/keys"
)

type commander struct {
	storeName string
	cdc       *wire.Codec
	decoder   did.DidDocDecoder
}

//CreateProjectRequestHandler create project handler
func CreateDidRequestHandler(storeName string, cdc *wire.Codec, kb keys.Keybase) func(http.ResponseWriter, *http.Request) {
	ctx := context.NewCoreContextFromViper()
	return func(w http.ResponseWriter, r *http.Request) {
		// collect data
		didDocParam := r.URL.Query().Get("didDoc")

		sovrinDid := sovrin.SovrinDid{}
		err := json.Unmarshal([]byte(didDocParam), &sovrinDid)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(fmt.Sprintf("Could not unmarshall didDoc into struct. Error: %s", err.Error())))
			return
		}

		// create the message
		msg := did.NewAddDidMsg(sovrinDid.Did, sovrinDid.VerifyKey)

		// Force the length to 64
		privKey := [64]byte{}
		copy(privKey[:], base58.Decode(sovrinDid.Secret.SignKey))
		copy(privKey[32:], base58.Decode(sovrinDid.VerifyKey))

		//Create the Signature
		signature := ixo.SignIxoMessage(msg, sovrinDid.Did, privKey)

		tx := ixo.NewIxoTx(msg, signature)

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

func QueryDidDocRequestHandler(storeName string, cdc *wire.Codec, decoder did.DidDocDecoder) func(http.ResponseWriter, *http.Request) {
	c := commander{storeName, cdc, decoder}
	ctx := context.NewCoreContextFromViper()
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		didAddr := vars["did"]

		key := ixo.Did(didAddr)

		res, err := ctx.Query([]byte(key), c.storeName)
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
		didDoc, err := c.decoder(res)
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

func QueryAllDidsRequestHandler(storeName string, cdc *wire.Codec, decoder did.DidDocDecoder) func(http.ResponseWriter, *http.Request) {
	c := commander{storeName, cdc, decoder}
	ctx := context.NewCoreContextFromViper()
	return func(w http.ResponseWriter, r *http.Request) {
		allKey := "ALL"

		res, err := ctx.Query([]byte(allKey), c.storeName)
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
