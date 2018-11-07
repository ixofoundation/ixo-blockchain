package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/wire"
	"github.com/gorilla/mux"
	"github.com/ixofoundation/ixo-cosmos/x/contracts"
)

const storeName = "params"

func registerQueryRoutes(cliCtx context.CLIContext, r *mux.Router, cdc *wire.Codec) {
	r.HandleFunc(
		"/contracts",
		queryContractsRequestHandler(cdc)).Methods("GET")
}

func queryContractsRequestHandler(cdc *wire.Codec) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.NewCLIContext().
			WithCodec(cdc).
			WithLogger(os.Stdout)

		m := make(map[string]string)

		var res string
		for _, k := range contracts.AllContracts {
			bz, err := ctx.QueryStore([]byte(contracts.MakeContractKey(k)), storeName)
			if err == nil {
				m[k], _ = string(bz[:]), res
			}
		}

		output, err := json.MarshalIndent(m, "", "  ")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("Couldn't marshall query result. Error: %s", err.Error())))
			return
		}

		w.Write(output)
	}
}
