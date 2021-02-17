package rest

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
	"github.com/ixofoundation/ixo-blockchain/x/bonds/internal/keeper"
	"github.com/ixofoundation/ixo-blockchain/x/bonds/internal/types"
	"net/http"
	"strconv"
)

func registerQueryRoutes(cliCtx context.CLIContext, r *mux.Router, queryRoute string) {
	r.HandleFunc(
		"/bonds", queryBondsHandler(cliCtx, queryRoute),
	).Methods("GET")

	r.HandleFunc(
		"/bonds_detailed",
		queryBondsDetailedHandler(cliCtx, queryRoute),
	).Methods("GET")

	r.HandleFunc(
		fmt.Sprintf("/bonds_detailed/{%s}", RestHeight),
		queryBondsDetailedHandlerWithHeight(cliCtx, queryRoute),
	).Methods("GET")

	r.HandleFunc(
		fmt.Sprintf("/bonds/{%s}", RestBondDid),
		queryBondHandler(cliCtx, queryRoute),
	).Methods("GET")

	r.HandleFunc(
		fmt.Sprintf("/bonds/{%s}/batch", RestBondDid),
		queryBatchHandler(cliCtx, queryRoute),
	).Methods("GET")

	r.HandleFunc(
		fmt.Sprintf("/bonds/{%s}/last_batch", RestBondDid),
		queryLastBatchHandler(cliCtx, queryRoute),
	).Methods("GET")

	r.HandleFunc(
		fmt.Sprintf("/bonds/{%s}/current_price", RestBondDid),
		queryCurrentPriceHandler(cliCtx, queryRoute),
	).Methods("GET")

	r.HandleFunc(
		fmt.Sprintf("/bonds/{%s}/current_reserve", RestBondDid),
		queryCurrentReserveHandler(cliCtx, queryRoute),
	).Methods("GET")

	r.HandleFunc(
		fmt.Sprintf("/bonds/{%s}/price/{%s}", RestBondDid, RestBondAmount),
		queryCustomPriceHandler(cliCtx, queryRoute),
	).Methods("GET")

	r.HandleFunc(
		fmt.Sprintf("/bonds/{%s}/buy_price/{%s}", RestBondDid, RestBondAmount),
		queryBuyPriceHandler(cliCtx, queryRoute),
	).Methods("GET")

	r.HandleFunc(
		fmt.Sprintf("/bonds/{%s}/sell_return/{%s}", RestBondDid, RestBondAmount),
		querySellReturnHandler(cliCtx, queryRoute),
	).Methods("GET")

	r.HandleFunc(
		fmt.Sprintf("/bonds/{%s}/swap_return/{%s}/{%s}", RestBondDid, RestFromTokenWithAmount, RestToToken),
		querySwapReturnHandler(cliCtx, queryRoute),
	).Methods("GET")

	r.HandleFunc(
		fmt.Sprintf("/bonds/{%s}/alpha_maximums", RestBondDid),
		queryAlphaMaximumsHandler(cliCtx, queryRoute),
	).Methods("GET")

	r.HandleFunc(
		"/bonds/params",
		queryParamsRequestHandler(cliCtx),
	).Methods("GET")
}

func queryBondsHandler(cliCtx context.CLIContext, queryRoute string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, _, err := cliCtx.QueryWithData(
			fmt.Sprintf("custom/%s/%s",
				queryRoute, keeper.QueryBonds), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}
		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func queryBondsDetailedHandler(cliCtx context.CLIContext, queryRoute string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, _, err := cliCtx.QueryWithData(
			fmt.Sprintf("custom/%s/%s",
				queryRoute, keeper.QueryBondsDetailed), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}
		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func queryBondsDetailedHandlerWithHeight(cliCtx context.CLIContext, queryRoute string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		heightStr := vars[RestHeight]

		height, err := strconv.ParseInt(heightStr, 10, 64)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest,
				fmt.Errorf("invalid height").Error())
			return
		}
		cliCtx = cliCtx.WithHeight(height)

		res, _, err := cliCtx.QueryWithData(
			fmt.Sprintf("custom/%s/%s",
				queryRoute, keeper.QueryBondsDetailed), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}
		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func queryBondHandler(cliCtx context.CLIContext, queryRoute string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		bondDid := vars[RestBondDid]

		res, _, err := cliCtx.QueryWithData(
			fmt.Sprintf("custom/%s/%s/%s",
				queryRoute, keeper.QueryBond, bondDid), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func queryBatchHandler(cliCtx context.CLIContext, queryRoute string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		bondDid := vars[RestBondDid]

		res, _, err := cliCtx.QueryWithData(
			fmt.Sprintf("custom/%s/%s/%s",
				queryRoute, keeper.QueryBatch, bondDid), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func queryLastBatchHandler(cliCtx context.CLIContext, queryRoute string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		bondDid := vars[RestBondDid]

		res, _, err := cliCtx.QueryWithData(
			fmt.Sprintf("custom/%s/%s/%s",
				queryRoute, keeper.QueryLastBatch, bondDid), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func queryCurrentPriceHandler(cliCtx context.CLIContext, queryRoute string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		bondDid := vars[RestBondDid]

		res, _, err := cliCtx.QueryWithData(
			fmt.Sprintf("custom/%s/%s/%s",
				queryRoute, keeper.QueryCurrentPrice, bondDid), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func queryCurrentReserveHandler(cliCtx context.CLIContext, queryRoute string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		bondDid := vars[RestBondDid]

		res, _, err := cliCtx.QueryWithData(
			fmt.Sprintf("custom/%s/%s/%s",
				queryRoute, keeper.QueryCurrentReserve, bondDid), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func queryCustomPriceHandler(cliCtx context.CLIContext, queryRoute string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		bondDid := vars[RestBondDid]
		bondAmount := vars[RestBondAmount]

		res, _, err := cliCtx.QueryWithData(
			fmt.Sprintf("custom/%s/%s/%s/%s",
				queryRoute, keeper.QueryCustomPrice, bondDid, bondAmount), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func queryBuyPriceHandler(cliCtx context.CLIContext, queryRoute string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		bondDid := vars[RestBondDid]
		bondAmount := vars[RestBondAmount]

		res, _, err := cliCtx.QueryWithData(
			fmt.Sprintf("custom/%s/%s/%s/%s",
				queryRoute, keeper.QueryBuyPrice, bondDid, bondAmount), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func querySellReturnHandler(cliCtx context.CLIContext, queryRoute string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		bondDid := vars[RestBondDid]
		bondAmount := vars[RestBondAmount]

		res, _, err := cliCtx.QueryWithData(
			fmt.Sprintf("custom/%s/%s/%s/%s",
				queryRoute, keeper.QuerySellReturn, bondDid, bondAmount), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func querySwapReturnHandler(cliCtx context.CLIContext, queryRoute string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		bondDid := vars[RestBondDid]
		fromTokenWithAmount := vars[RestFromTokenWithAmount]
		toToken := vars[RestToToken]

		reserveCoinWithAmount, err := sdk.ParseCoin(fromTokenWithAmount)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		res, _, err := cliCtx.QueryWithData(
			fmt.Sprintf("custom/%s/%s/%s/%s/%s/%s",
				queryRoute, keeper.QuerySwapReturn, bondDid, reserveCoinWithAmount.Denom,
				reserveCoinWithAmount.Amount.String(), toToken), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func queryAlphaMaximumsHandler(cliCtx context.CLIContext, queryRoute string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		bondDid := vars[RestBondDid]

		res, _, err := cliCtx.QueryWithData(
			fmt.Sprintf("custom/%s/%s/%s",
				queryRoute, keeper.QueryAlphaMaximums, bondDid), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func queryParamsRequestHandler(cliCtx context.CLIContext) http.HandlerFunc {
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
