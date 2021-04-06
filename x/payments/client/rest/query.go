package rest
//
//import (
//	"fmt"
//	"github.com/gorilla/mux"
//	"net/http"
//
//	"github.com/cosmos/cosmos-sdk/client/context"
//	"github.com/cosmos/cosmos-sdk/types/rest"
//
//	"github.com/ixofoundation/ixo-blockchain/x/payments/internal/keeper"
//	"github.com/ixofoundation/ixo-blockchain/x/payments/internal/types"
//)
//
//func registerQueryRoutes(cliCtx context.CLIContext, r *mux.Router) {
//	r.HandleFunc(fmt.Sprintf("/payments/templates/{%s}", RestPaymentTemplateId),
//		queryPaymentTemplateHandler(cliCtx)).Methods("GET")
//
//	r.HandleFunc(fmt.Sprintf("/payments/contracts/{%s}", RestPaymentContractId),
//		queryPaymentContractHandler(cliCtx)).Methods("GET")
//
//	r.HandleFunc(fmt.Sprintf("/payments/contracts_by_id_prefix/{%s}", RestPaymentContractsIdPrefix),
//		queryPaymentContractsByIdPrefixHandler(cliCtx)).Methods("GET")
//
//	r.HandleFunc(fmt.Sprintf("/payments/subscriptions/{%s}", RestSubscriptionId),
//		querySubscriptionHandler(cliCtx)).Methods("GET")
//}
//
//func queryPaymentTemplateHandler(cliCtx context.CLIContext) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		vars := mux.Vars(r)
//		templateId := vars[RestPaymentTemplateId]
//
//		bz, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s",
//			types.QuerierRoute, keeper.QueryPaymentTemplate, templateId), nil)
//		if err != nil {
//			w.WriteHeader(http.StatusInternalServerError)
//			_, _ = w.Write([]byte(fmt.Sprintf("Couldn't get query data %s", err.Error())))
//			return
//		}
//
//		var template types.PaymentTemplate
//		if err := cliCtx.Codec.UnmarshalJSON(bz, &template); err != nil {
//			w.WriteHeader(http.StatusInternalServerError)
//			_, _ = w.Write([]byte(fmt.Sprintf("Couldn't Unmarshal data %s", err.Error())))
//			return
//		}
//
//		rest.PostProcessResponse(w, cliCtx, template)
//	}
//}
//
//func queryPaymentContractHandler(cliCtx context.CLIContext) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		vars := mux.Vars(r)
//		contractId := vars[RestPaymentContractId]
//
//		bz, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s",
//			types.QuerierRoute, keeper.QueryPaymentContract, contractId), nil)
//		if err != nil {
//			w.WriteHeader(http.StatusInternalServerError)
//			_, _ = w.Write([]byte(fmt.Sprintf("Couldn't get query data %s", err.Error())))
//			return
//		}
//
//		var contract types.PaymentContract
//		if err := cliCtx.Codec.UnmarshalJSON(bz, &contract); err != nil {
//			w.WriteHeader(http.StatusInternalServerError)
//			_, _ = w.Write([]byte(fmt.Sprintf("Couldn't Unmarshal data %s", err.Error())))
//			return
//		}
//
//		rest.PostProcessResponse(w, cliCtx, contract)
//	}
//}
//
//func queryPaymentContractsByIdPrefixHandler(cliCtx context.CLIContext) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		vars := mux.Vars(r)
//		contractIdPrefix := vars[RestPaymentContractsIdPrefix]
//
//		bz, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s",
//			types.QuerierRoute, keeper.QueryPaymentContractsByIdPrefix, contractIdPrefix), nil)
//		if err != nil {
//			w.WriteHeader(http.StatusInternalServerError)
//			_, _ = w.Write([]byte(fmt.Sprintf("Couldn't get query data %s", err.Error())))
//			return
//		}
//
//		var contracts []types.PaymentContract
//		if err := cliCtx.Codec.UnmarshalJSON(bz, &contracts); err != nil {
//			w.WriteHeader(http.StatusInternalServerError)
//			_, _ = w.Write([]byte(fmt.Sprintf("Couldn't Unmarshal data %s", err.Error())))
//			return
//		}
//
//		rest.PostProcessResponse(w, cliCtx, contracts)
//	}
//}
//
//func querySubscriptionHandler(cliCtx context.CLIContext) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		vars := mux.Vars(r)
//		subscriptionId := vars[RestSubscriptionId]
//
//		bz, _, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s",
//			types.QuerierRoute, keeper.QuerySubscription, subscriptionId), nil)
//		if err != nil {
//			w.WriteHeader(http.StatusInternalServerError)
//			_, _ = w.Write([]byte(fmt.Sprintf("Couldn't get query data %s", err.Error())))
//			return
//		}
//
//		var subscription types.Subscription
//		if err := cliCtx.Codec.UnmarshalJSON(bz, &subscription); err != nil {
//			w.WriteHeader(http.StatusInternalServerError)
//			_, _ = w.Write([]byte(fmt.Sprintf("Couldn't Unmarshal data %s", err.Error())))
//			return
//		}
//
//		rest.PostProcessResponse(w, cliCtx, subscription)
//	}
//}
