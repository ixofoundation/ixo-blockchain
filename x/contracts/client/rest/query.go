package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"

	"github.com/ixofoundation/ixo-cosmos/x/contracts/internal/keeper"
	"github.com/ixofoundation/ixo-cosmos/x/contracts/internal/types"
)

func registerQueryRoutes(cliCtx context.CLIContext, r *mux.Router) {
	r.HandleFunc("/contracts", queryContractsRequestHandler(cliCtx)).Methods("GET")
}

func queryContractsRequestHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		bz, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", types.QuerierRoute,
			keeper.QueryAllContracts), nil)

		contracts := make(map[string]string)
		err = json.Unmarshal(bz, &contracts)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(fmt.Sprintf("Couldn't Unmarshall query result. Error: %s", err.Error())))
			return
		}

		rest.PostProcessResponse(w, cliCtx, contracts)
	}
}
