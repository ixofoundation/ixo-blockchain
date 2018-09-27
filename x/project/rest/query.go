package rest

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/wire"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/gorilla/mux"
	"github.com/ixofoundation/ixo-cosmos/types"
	"github.com/ixofoundation/ixo-cosmos/x/ixo"
	"github.com/ixofoundation/ixo-cosmos/x/project"
)

type AccDetails struct {
	Did     string  `json:"did"`
	Account string  `json:"account"`
	Balance sdk.Int `json:"balance"`
}

func registerQueryRoutes(cliCtx context.CLIContext, r *mux.Router, cdc *wire.Codec) {
	r.HandleFunc(
		"/project/{did}",
		queryProjectDocRequestHandler(cdc, project.GetProjectDocDecoder(cdc)),
	).Methods("GET")
	r.HandleFunc(
		"/projectAccounts/{projectDid}",
		queryProjectAccountsRequestHandler(cdc, project.GetProjectDocDecoder(cdc)),
	).Methods("GET")
}

func queryProjectDocRequestHandler(cdc *wire.Codec, decoder project.ProjectDocDecoder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.NewCLIContext().
			WithCodec(cdc).
			WithLogger(os.Stdout)

		vars := mux.Vars(r)
		didAddr := vars["did"]

		key := ixo.Did(didAddr)

		res, err := ctx.QueryStore([]byte(key), storeName)
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
		projectDoc, err := decoder(res)
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

func queryProjectAccountsRequestHandler(cdc *wire.Codec, decoder project.ProjectDocDecoder) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.NewCLIContext().
			WithCodec(cdc).
			WithLogger(os.Stdout)

		vars := mux.Vars(r)
		projectDid := vars["projectDid"]

		var buffer bytes.Buffer
		buffer.WriteString("ACC-")
		buffer.WriteString(projectDid)
		key := buffer.Bytes()

		res, err := ctx.QueryStore(key, storeName)
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

			balance, err := getAccountBalance(ctx, addr, types.GetAccountDecoder(cdc))
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

func getAccountBalance(ctx context.CLIContext, addr string, decoder auth.AccountDecoder) (sdk.Int, error) {
	bz, err := hex.DecodeString(addr)
	if err != nil {
		return sdk.NewInt(0), err
	}

	res, err := ctx.QueryStore(bz, "main")
	if err != nil {
		return sdk.NewInt(0), err
	}

	// decode the value
	account, err := decoder(res)
	if err != nil {
		return sdk.NewInt(0), err
	}

	baseAcc := account.(*types.AppAccount)

	return baseAcc.Coins.AmountOf(project.COIN_DENOM), nil

}
