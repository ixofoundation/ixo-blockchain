package rest

import (
	"fmt"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"

	rest2 "github.com/ixofoundation/ixo-cosmos/client/rest"
	"github.com/ixofoundation/ixo-cosmos/types"
	fiatTypes "github.com/ixofoundation/ixo-cosmos/x/fiat/internal/types"
)

func QueryFiatRequestHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		pegHashStr := vars["pegHash"]

		pegHashHex, err := types.GetPegHashHex(pegHashStr)
		if err != nil {
			rest2.WriteErrorResponse(w, fiatTypes.ErrPegHashHex(fiatTypes.DefaultCodeSpace, pegHashStr))
			return
		}

		res, height, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", fiatTypes.QuerierRoute, pegHashStr), pegHashHex)
		if err != nil {
			rest2.WriteErrorResponse(w, fiatTypes.ErrQuery(fiatTypes.DefaultCodeSpace, pegHashStr))
			return
		}

		cliCtx = cliCtx.WithHeight(height)
		rest.PostProcessResponse(w, cliCtx, res)
	}
}
