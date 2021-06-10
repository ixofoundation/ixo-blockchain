package rest

import (
	"fmt"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
	"github.com/ixofoundation/ixo-blockchain/x/payments/keeper"
	"github.com/ixofoundation/ixo-blockchain/x/payments/types"
)

func registerQueryRoutes(clientCtx client.Context, r *mux.Router) {
	r.HandleFunc(fmt.Sprintf("/payments/templates/{%s}", RestPaymentTemplateId),
		queryPaymentTemplateHandler(clientCtx)).Methods("GET")

	r.HandleFunc(fmt.Sprintf("/payments/contracts/{%s}", RestPaymentContractId),
		queryPaymentContractHandler(clientCtx)).Methods("GET")

	r.HandleFunc(fmt.Sprintf("/payments/contracts_by_id_prefix/{%s}", RestPaymentContractsIdPrefix),
		queryPaymentContractsByIdPrefixHandler(clientCtx)).Methods("GET")

	r.HandleFunc(fmt.Sprintf("/payments/subscriptions/{%s}", RestSubscriptionId),
		querySubscriptionHandler(clientCtx)).Methods("GET")
}

func queryPaymentTemplateHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		templateId := vars[RestPaymentTemplateId]

		bz, _, err := clientCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s",
			types.QuerierRoute, keeper.QueryPaymentTemplate, templateId), nil)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(fmt.Sprintf("Couldn't get query data %s", err.Error())))
			return
		}

		var template types.PaymentTemplate
		if err := clientCtx.LegacyAmino.UnmarshalJSON(bz, &template); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(fmt.Sprintf("Couldn't Unmarshal data %s", err.Error())))
			return
		}

		rest.PostProcessResponse(w, clientCtx, template)
	}
}

func queryPaymentContractHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		contractId := vars[RestPaymentContractId]

		bz, _, err := clientCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s",
			types.QuerierRoute, keeper.QueryPaymentContract, contractId), nil)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(fmt.Sprintf("Couldn't get query data %s", err.Error())))
			return
		}

		var contract types.PaymentContract
		if err := clientCtx.LegacyAmino.UnmarshalJSON(bz, &contract); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(fmt.Sprintf("Couldn't Unmarshal data %s", err.Error())))
			return
		}

		rest.PostProcessResponse(w, clientCtx, contract)
	}
}

func queryPaymentContractsByIdPrefixHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		contractIdPrefix := vars[RestPaymentContractsIdPrefix]

		bz, _, err := clientCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s",
			types.QuerierRoute, keeper.QueryPaymentContractsByIdPrefix, contractIdPrefix), nil)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(fmt.Sprintf("Couldn't get query data %s", err.Error())))
			return
		}

		var contracts []types.PaymentContract
		if err := clientCtx.LegacyAmino.UnmarshalJSON(bz, &contracts); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(fmt.Sprintf("Couldn't Unmarshal data %s", err.Error())))
			return
		}

		rest.PostProcessResponse(w, clientCtx, contracts)
	}
}

func querySubscriptionHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		subscriptionId := vars[RestSubscriptionId]

		bz, _, err := clientCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s",
			types.QuerierRoute, keeper.QuerySubscription, subscriptionId), nil)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(fmt.Sprintf("Couldn't get query data %s", err.Error())))
			return
		}

		var subscription types.Subscription
		if err := clientCtx.LegacyAmino.UnmarshalJSON(bz, &subscription); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_, _ = w.Write([]byte(fmt.Sprintf("Couldn't Unmarshal data %s", err.Error())))
			return
		}

		rest.PostProcessResponse(w, clientCtx, subscription)
	}
}
