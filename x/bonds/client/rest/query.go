package rest

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
	"github.com/ixofoundation/ixo-blockchain/x/bonds/keeper"
	"github.com/ixofoundation/ixo-blockchain/x/bonds/types"
	"net/http"
)

func registerQueryRoutes(clientCtx client.Context, r *mux.Router, queryRoute string) {
	r.HandleFunc("/bonds", queryBondsHandler(clientCtx, queryRoute)).Methods("GET")
	r.HandleFunc("/bonds_detailed", queryBondsDetailedHandler(clientCtx, queryRoute)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/bonds/{%s}", RestBondDid), queryBondHandler(clientCtx, queryRoute)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/bonds/{%s}/batch", RestBondDid), queryBatchHandler(clientCtx, queryRoute)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/bonds/{%s}/last_batch", RestBondDid),	queryLastBatchHandler(clientCtx, queryRoute)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/bonds/{%s}/current_price", RestBondDid), queryCurrentPriceHandler(clientCtx, queryRoute)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/bonds/{%s}/current_reserve", RestBondDid), queryCurrentReserveHandler(clientCtx, queryRoute)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/bonds/{%s}/price/{%s}", RestBondDid, RestBondAmount), queryCustomPriceHandler(clientCtx, queryRoute)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/bonds/{%s}/buy_price/{%s}", RestBondDid, RestBondAmount), queryBuyPriceHandler(clientCtx, queryRoute)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/bonds/{%s}/sell_return/{%s}", RestBondDid, RestBondAmount), querySellReturnHandler(clientCtx, queryRoute)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/bonds/{%s}/swap_return/{%s}/{%s}", RestBondDid, RestFromTokenWithAmount, RestToToken), querySwapReturnHandler(clientCtx, queryRoute)).Methods("GET")
	r.HandleFunc(fmt.Sprintf("/bonds/{%s}/alpha_maximums", RestBondDid), queryAlphaMaximumsHandler(clientCtx, queryRoute)).Methods("GET")
	r.HandleFunc("/bonds/params", queryParamsRequestHandler(clientCtx)).Methods("GET")
}

func queryBondsHandler(clientCtx client.Context, queryRoute string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		res, _, err := clientCtx.QueryWithData(
			fmt.Sprintf("custom/%s/%s",
				queryRoute, keeper.QueryBonds), nil)
		if rest.CheckNotFoundError(w, err) {
			return
		}
		rest.PostProcessResponse(w, clientCtx, res)
	}
}

func queryBondsDetailedHandler(clientCtx client.Context, queryRoute string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, clientCtx, r)
		if !ok {
			return
		}

		res, height, err := cliCtx.QueryWithData(
			fmt.Sprintf("custom/%s/%s",
				queryRoute, keeper.QueryBondsDetailed), nil)
		if rest.CheckNotFoundError(w, err) {
			return
		}

		cliCtx = cliCtx.WithHeight(height)

		rest.PostProcessResponse(w, cliCtx, res)
	}
}

func queryBondHandler(clientCtx client.Context, queryRoute string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		bondDid := vars[RestBondDid]

		res, _, err := clientCtx.QueryWithData(
			fmt.Sprintf("custom/%s/%s/%s",
				queryRoute, keeper.QueryBond, bondDid), nil)
		if rest.CheckNotFoundError(w, err) {
			return
		}

		rest.PostProcessResponse(w, clientCtx, res)
	}
}

func queryBatchHandler(clientCtx client.Context, queryRoute string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		bondDid := vars[RestBondDid]

		res, _, err := clientCtx.QueryWithData(
			fmt.Sprintf("custom/%s/%s/%s",
				queryRoute, keeper.QueryBatch, bondDid), nil)
		if rest.CheckNotFoundError(w, err) {
			return
		}

		rest.PostProcessResponse(w, clientCtx, res)
	}
}

func queryLastBatchHandler(clientCtx client.Context, queryRoute string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		bondDid := vars[RestBondDid]

		res, _, err := clientCtx.QueryWithData(
			fmt.Sprintf("custom/%s/%s/%s",
				queryRoute, keeper.QueryLastBatch, bondDid), nil)
		if rest.CheckNotFoundError(w, err) {
			return
		}

		rest.PostProcessResponse(w, clientCtx, res)
	}
}

func queryCurrentPriceHandler(clientCtx client.Context, queryRoute string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		bondDid := vars[RestBondDid]

		res, _, err := clientCtx.QueryWithData(
			fmt.Sprintf("custom/%s/%s/%s",
				queryRoute, keeper.QueryCurrentPrice, bondDid), nil)
		if rest.CheckNotFoundError(w, err) {
			return
		}

		rest.PostProcessResponse(w, clientCtx, res)
	}
}

func queryCurrentReserveHandler(clientCtx client.Context, queryRoute string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		bondDid := vars[RestBondDid]

		res, _, err := clientCtx.QueryWithData(
			fmt.Sprintf("custom/%s/%s/%s",
				queryRoute, keeper.QueryCurrentReserve, bondDid), nil)
		if rest.CheckNotFoundError(w, err) {
			return
		}

		rest.PostProcessResponse(w, clientCtx, res)
	}
}

func queryCustomPriceHandler(clientCtx client.Context, queryRoute string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		bondDid := vars[RestBondDid]
		bondAmount := vars[RestBondAmount]

		res, _, err := clientCtx.QueryWithData(
			fmt.Sprintf("custom/%s/%s/%s/%s",
				queryRoute, keeper.QueryCustomPrice, bondDid, bondAmount), nil)
		if rest.CheckNotFoundError(w, err) {
			return
		}

		rest.PostProcessResponse(w, clientCtx, res)
	}
}

func queryBuyPriceHandler(clientCtx client.Context, queryRoute string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		bondDid := vars[RestBondDid]
		bondAmount := vars[RestBondAmount]

		res, _, err := clientCtx.QueryWithData(
			fmt.Sprintf("custom/%s/%s/%s/%s",
				queryRoute, keeper.QueryBuyPrice, bondDid, bondAmount), nil)
		if rest.CheckNotFoundError(w, err) {
			return
		}

		rest.PostProcessResponse(w, clientCtx, res)
	}
}

func querySellReturnHandler(clientCtx client.Context, queryRoute string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		bondDid := vars[RestBondDid]
		bondAmount := vars[RestBondAmount]

		res, _, err := clientCtx.QueryWithData(
			fmt.Sprintf("custom/%s/%s/%s/%s",
				queryRoute, keeper.QuerySellReturn, bondDid, bondAmount), nil)
		if rest.CheckNotFoundError(w, err) {
			return
		}

		rest.PostProcessResponse(w, clientCtx, res)
	}
}

func querySwapReturnHandler(clientCtx client.Context, queryRoute string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		bondDid := vars[RestBondDid]
		fromTokenWithAmount := vars[RestFromTokenWithAmount]
		toToken := vars[RestToToken]

		reserveCoinWithAmount, err := sdk.ParseCoinNormalized(fromTokenWithAmount)
		if rest.CheckNotFoundError(w, err) {
			return
		}

		res, _, err := clientCtx.QueryWithData(
			fmt.Sprintf("custom/%s/%s/%s/%s/%s/%s",
				queryRoute, keeper.QuerySwapReturn, bondDid, reserveCoinWithAmount.Denom,
				reserveCoinWithAmount.Amount.String(), toToken), nil)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
			return
		}

		rest.PostProcessResponse(w, clientCtx, res)
	}
}

func queryAlphaMaximumsHandler(clientCtx client.Context, queryRoute string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		bondDid := vars[RestBondDid]

		res, _, err := clientCtx.QueryWithData(
			fmt.Sprintf("custom/%s/%s/%s",
				queryRoute, keeper.QueryAlphaMaximums, bondDid), nil)
		if rest.CheckNotFoundError(w, err) {
			return
		}

		rest.PostProcessResponse(w, clientCtx, res)
	}
}

func queryParamsRequestHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		bz, _, err := clientCtx.QueryWithData(fmt.Sprintf("custom/%s/%s", types.QuerierRoute,
			keeper.QueryParams), nil)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(fmt.Sprintf("Couldn't get query data %s", err.Error())))
			return
		}

		var params types.Params
		if err := clientCtx.LegacyAmino.UnmarshalJSON(bz, &params); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(fmt.Sprintf("Couldn't Unmarshal data %s", err.Error())))
			return
		}

		rest.PostProcessResponse(w, clientCtx, params)
	}
}
