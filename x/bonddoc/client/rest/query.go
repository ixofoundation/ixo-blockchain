package rest

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/gorilla/mux"

	"github.com/ixofoundation/ixo-blockchain/x/bonddoc/internal/keeper"
	"github.com/ixofoundation/ixo-blockchain/x/bonddoc/internal/types"
	"github.com/ixofoundation/ixo-blockchain/x/ixo"
)

func registerQueryRoutes(cliCtx context.CLIContext, r *mux.Router) {
	r.HandleFunc("/bond/{did}", queryBondDocRequestHandler(cliCtx)).Methods("GET")
}

func queryBondDocRequestHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		didAddr := vars["did"]

		key := ixo.Did(didAddr)
		res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s", types.QuerierRoute,
			keeper.QueryBondDoc, key), nil)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(fmt.Sprintf("Could't query did. Error: %s", err.Error())))
			return
		}

		if len(res) == 0 {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		var bondDoc types.MsgCreateBond
		cliCtx.Codec.MustUnmarshalJSON(res, &bondDoc)

		bz, err := json.Marshal(bondDoc)
		_, _ = w.Write(bz)
	}
}
