package rest

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/types/rest"

	"github.com/ixofoundation/ixo-blockchain/x/fees/internal/keeper"
	"github.com/ixofoundation/ixo-blockchain/x/fees/internal/types"
)

func queryFeeParamsHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		bz, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", types.QuerierRoute,
			keeper.QueryParams), nil)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(fmt.Sprintf("Couldn't get query data %s", err.Error())))

			return
		}

		var params types.Params
		if err := cliCtx.Codec.UnmarshalJSON(bz, &params); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(fmt.Sprintf("Couldn't Unmarshal data %s", err.Error())))

			return
		}

		rest.PostProcessResponse(w, cliCtx, params)
	}
}

func queryFeeHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		feeId := vars[RestFeeId]

		bz, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s",
			types.QuerierRoute, keeper.QueryFee, feeId), nil)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(fmt.Sprintf("Couldn't get query data %s", err.Error())))

			return
		}

		var fee types.Fee
		if err := cliCtx.Codec.UnmarshalJSON(bz, &fee); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(fmt.Sprintf("Couldn't Unmarshal data %s", err.Error())))

			return
		}

		rest.PostProcessResponse(w, cliCtx, fee)
	}
}

func queryFeeContractHandler(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		feeContractId := vars[RestFeeContractId]

		bz, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s",
			types.QuerierRoute, keeper.QueryFeeContract, feeContractId), nil)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(fmt.Sprintf("Couldn't get query data %s", err.Error())))

			return
		}

		var feeContract types.FeeContract
		if err := cliCtx.Codec.UnmarshalJSON(bz, &feeContract); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(fmt.Sprintf("Couldn't Unmarshal data %s", err.Error())))

			return
		}

		rest.PostProcessResponse(w, cliCtx, feeContract)
	}
}
