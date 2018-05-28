package rest

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"

	base58 "github.com/btcsuite/btcutil/base58"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/core"
	"github.com/cosmos/cosmos-sdk/wire"
	"github.com/gorilla/mux"
	"github.com/ixofoundation/ixo-cosmos/types"
	"github.com/ixofoundation/ixo-cosmos/x/ixo"
	"github.com/ixofoundation/ixo-cosmos/x/ixo/sovrin"
	"github.com/ixofoundation/ixo-cosmos/x/project"
	"github.com/tendermint/go-crypto/keys"
)

type commander struct {
	storeName string
	cdc       *wire.Codec
	decoder   project.ProjectDocDecoder
}

type sendBody struct {
	Data string `json:"data"`
}

type AccDetails struct {
	Did     string `json:"did"`
	Account string `json:"account"`
	Balance int64  `json:"balance"`
}

//CreateProjectRequestHandler create project handler
func CreateProjectRequestHandler(storeName string, cdc *wire.Codec, kb keys.Keybase) func(http.ResponseWriter, *http.Request) {
	ctx := context.NewCoreContextFromViper()
	return func(w http.ResponseWriter, r *http.Request) {
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

// Example curl request
/*
curl -X POST -G \
-H "Content-Type: application/json" \
-H "Accept: application/json" \
-d projectDoc='{"did":"","pubKey":"","title":"ReforestationCongo","shortDescription":"DescriptionaboutReforestation","longDescription":"DescriptionaboutReforestationlong","impactAction":"treesplanted","createdOn":"2018-05-14T13:56:16+00:00","createdBy":"","country":"CO","sdgs":["12.1","8.2"],"impactsRequired":"34","claimTemplate":"default","serviceURI":"http://localhost:8080/pds","socialMedia":{"facebookLink":"","instagramLink":"","twitterLink":""},"webLink":"","image":""}' \
-d didDoc='{"did":"CCzPRoyPQsTxVwoAwTZXcK","verifyKey":"77GSw8G26F1e3qwtmzzTvWicZSCCkFKK43NSntpfuJKx","encryptionPublicKey":"AhKmLwrPdPMY3yeBpPqUy8qsphgXGaFEWHNgeUxKa3bV","secret":{"seed":"ea25949b56257a8f16435af37d333fb11258fe9f7a1c2a8eebbebb4d0ea2ae85","signKey":"Gm1dz5ToFcw3Ur7aRqpfzXh9kFJ8C6FZTTueCSGaZDH6","encryptionPrivateKey":"Gm1dz5ToFcw3Ur7aRqpfzXh9kFJ8C6FZTTueCSGaZDH6"}}' \
http://localhost:1317/project*/

func QueryProjectDocRequestHandler(storeName string, cdc *wire.Codec, decoder project.ProjectDocDecoder) func(http.ResponseWriter, *http.Request) {
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
		projectDoc, err := c.decoder(res)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("Could't parse query result. Result: %s. Error: %s", res, err.Error())))
			return
		}

		// print out whole projectDoc
		output, err := json.MarshalIndent(projectDoc, "", "  ")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("Could't marshall query result. Error: %s", err.Error())))
			return
		}

		w.Write(output)
	}
}

func QueryProjectAccountsRequestHandler(storeName string, cdc *wire.Codec, decoder project.ProjectDocDecoder) func(http.ResponseWriter, *http.Request) {
	c := commander{storeName, cdc, decoder}
	ctx := context.NewCoreContextFromViper()
	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		projectDid := vars["projectDid"]

		var buffer bytes.Buffer
		buffer.WriteString("ACC-")
		buffer.WriteString(projectDid)
		key := buffer.Bytes()

		res, err := ctx.Query(key, c.storeName)
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
		var f interface{}
		err = json.Unmarshal(res, &f)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("Could't parse query result. Result: %s. Error: %s", res, err.Error())))
			return
		}
		accMap := f.(map[string]interface{})

		accDetails := make([]AccDetails, len(accMap))
		i := 0
		for k, v := range accMap {
			addr := v.(string)

			balance, err := GetAccountBalance(ctx, addr, types.GetAccountDecoder(cdc))
			if err != nil {
				panic(err)
			}
			accDetails[i] = AccDetails{Did: k, Account: addr, Balance: balance}
			i = i + 1
		}

		// print out whole didDoc
		output, err := json.MarshalIndent(accDetails, "", "  ")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("Could't marshall query result. Error: %s", err.Error())))
			return
		}

		w.Write(output)
	}

}

func GetAccountBalance(ctx core.CoreContext, addr string, decoder sdk.AccountDecoder) (int64, error) {
	bz, err := hex.DecodeString(addr)
	if err != nil {
		return 0, err
	}
	key := sdk.Address(bz)

	res, err := ctx.Query(key, "main")
	if err != nil {
		return 0, err
	}

	// decode the value
	account, err := decoder(res)
	if err != nil {
		return 0, err
	}

	baseAcc := account.(*types.AppAccount)

	return baseAcc.Coins.AmountOf(project.COIN_DENOM), nil

}
