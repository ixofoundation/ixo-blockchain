package rest

import (
	"fmt"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"

	fiatTypes "github.com/ixofoundation/ixo-cosmos/x/fiat/internal/types"
)

func QueryFiatAccountHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		bech32addr := vars["address"]

		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		res, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s", fiatTypes.QuerierRoute, "queryFiatAccount", bech32addr), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, fiatTypes.ErrQuery(fiatTypes.DefaultCodeSpace, bech32addr).Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx, res)
	}
}
