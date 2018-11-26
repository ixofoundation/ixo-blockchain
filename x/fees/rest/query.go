package rest

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/ixofoundation/ixo-cosmos/x/fees"

	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/wire"
	"github.com/gorilla/mux"
)

const storeName = "params"

func registerQueryRoutes(cliCtx context.CLIContext, r *mux.Router, cdc *wire.Codec) {
	r.HandleFunc(
		"/fees",
		queryFeesRequestHandler(cdc)).Methods("GET")
}

func queryFeesRequestHandler(cdc *wire.Codec) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.NewCLIContext().
			WithCodec(cdc).
			WithLogger(os.Stdout)

		m := make(map[string]float64)

		var res sdk.Rat
		for _, k := range fees.AllFees {
			bz, err := ctx.QueryStore([]byte(fees.MakeFeeKey(k)), storeName)
			if err == nil {
				cdc.UnmarshalBinary(bz, &res)
				m[k], _ = res.Float64()
			}
		}

		output, err := json.MarshalIndent(m, "", "  ")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(fmt.Sprintf("Could't marshall query result. Error: %s", err.Error())))
			return
		}

		w.Write(output)
	}
}
