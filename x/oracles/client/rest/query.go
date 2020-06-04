package rest

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/types/rest"

	"github.com/ixofoundation/ixo-blockchain/x/oracles/internal/keeper"
	"github.com/ixofoundation/ixo-blockchain/x/oracles/internal/types"
)

func registerQueryRoutes(cliCtx context.CLIContext, r *mux.Router) {
	r.HandleFunc("/fees", queryFeesRequestHandler(cliCtx)).Methods("GET")
}

func queryFeesRequestHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		bz, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", types.QuerierRoute,
			keeper.QueryOracles), nil)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(fmt.Sprintf("Couldn't get query data %s", err.Error())))
			return
		}

		var params types.Oracles
		if err := cliCtx.Codec.UnmarshalJSON(bz, &params); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(fmt.Sprintf("Couldn't Unmarshal data %s", err.Error())))
			return
		}

		rest.PostProcessResponse(w, cliCtx, params)
	}
}
