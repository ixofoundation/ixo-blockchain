package rest

import (
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
	didexported "github.com/ixofoundation/ixo-blockchain/lib/legacydid"
	"github.com/ixofoundation/ixo-blockchain/x/payments/types"
)

func registerTxHandlers(cliCtx client.Context, r *mux.Router) {
	r.HandleFunc("/payments/create_payment_template", newCreatePaymentTemplateRequestHandlerFn(cliCtx)).Methods("POST")
	r.HandleFunc("/payments/create_payment_contract", newCreatePaymentContractRequestHandlerFn(cliCtx)).Methods("POST")
	r.HandleFunc("/payments/create_subscription", newCreateSubscriptionRequestHandlerFn(cliCtx)).Methods("POST")
	r.HandleFunc("/payments/set_payment_contract_authorisation", newSetPaymentContractAuthorisationRequestHandlerFn(cliCtx)).Methods("POST")
	r.HandleFunc("/payments/grant_discount", newGrantDiscountRequestHandlerFn(cliCtx)).Methods("POST")
	r.HandleFunc("/payments/revoke_discount", newRevokeDiscountRequestHandlerFn(cliCtx)).Methods("POST")
	r.HandleFunc("/payments/effect_payment", newEffectPaymentRequestHandlerFn(cliCtx)).Methods("POST")
}

type createPaymentTemplateReq struct {
	BaseReq         rest.BaseReq          `json:"base_req" yaml:"base_req"`
	CreatorDid      didexported.Did       `json:"creator_did" yaml:"creator_did"`
	PaymentTemplate types.PaymentTemplate `json:"payment_template" yaml:"payment_template"`
}

func newCreatePaymentTemplateRequestHandlerFn(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createPaymentTemplateReq
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		msg := types.NewMsgCreatePaymentTemplate(req.PaymentTemplate, req.CreatorDid)
		if rest.CheckBadRequestError(w, msg.ValidateBasic()) {
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

type createPaymentContractReq struct {
	BaseReq           rest.BaseReq       `json:"base_req" yaml:"base_req"`
	CreatorDid        didexported.Did    `json:"creator_did" yaml:"creator_did"`
	PaymentTemplateId string             `json:"payment_template_id" yaml:"payment_template_id"`
	PaymentContractId string             `json:"payment_contract_id" yaml:"payment_contract_id"`
	Payer             sdk.AccAddress     `json:"payer" yaml:"payer"`
	Recipients        types.Distribution `json:"recipients" yaml:"recipients"`
	CanDeauthorise    bool               `json:"can_deauthorise" yaml:"can_deauthorise"`
	DiscountId        sdk.Uint           `json:"discount_id" yaml:"discount_id"`
}

func newCreatePaymentContractRequestHandlerFn(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createPaymentContractReq
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		msg := types.NewMsgCreatePaymentContract(req.PaymentTemplateId,
			req.PaymentContractId, req.Payer, req.Recipients,
			req.CanDeauthorise, req.DiscountId, req.CreatorDid)
		if rest.CheckBadRequestError(w, msg.ValidateBasic()) {
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

type createSubscriptionReq struct {
	BaseReq           rest.BaseReq    `json:"base_req" yaml:"base_req"`
	CreatorDid        didexported.Did `json:"creator_did" yaml:"creator_did"`
	SubscriptionId    string          `json:"subscription_id" yaml:"subscription_id"`
	PaymentContractId string          `json:"payment_contract_id" yaml:"payment_contract_id"`
	MaxPeriods        sdk.Uint        `json:"max_periods" yaml:"max_periods"`
	Period            types.Period    `json:"period" yaml:"period"`
}

func newCreateSubscriptionRequestHandlerFn(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createSubscriptionReq
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		msg := types.NewMsgCreateSubscription(req.SubscriptionId, req.PaymentContractId,
			req.MaxPeriods, req.Period, req.CreatorDid)
		if rest.CheckBadRequestError(w, msg.ValidateBasic()) {
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

type setPaymentContractAuthorisationReq struct {
	BaseReq           rest.BaseReq    `json:"base_req" yaml:"base_req"`
	PayerDid          didexported.Did `json:"payer_did" yaml:"payer_did"`
	PaymentContractId string          `json:"payment_contract_id" yaml:"payment_contract_id"`
	Authorised        bool            `json:"authorised" yaml:"authorised"`
}

func newSetPaymentContractAuthorisationRequestHandlerFn(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req setPaymentContractAuthorisationReq
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		msg := types.NewMsgSetPaymentContractAuthorisation(req.PaymentContractId,
			req.Authorised, req.PayerDid)
		if rest.CheckBadRequestError(w, msg.ValidateBasic()) {
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

type grantDiscountReq struct {
	BaseReq           rest.BaseReq    `json:"base_req" yaml:"base_req"`
	SenderDid         didexported.Did `json:"sender_did" yaml:"sender_did"`
	PaymentContractId string          `json:"payment_contract_id" yaml:"payment_contract_id"`
	DiscountId        sdk.Uint        `json:"discount_id" yaml:"discount_id"`
	Recipient         sdk.AccAddress  `json:"recipient" yaml:"recipient"`
}

func newGrantDiscountRequestHandlerFn(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req grantDiscountReq
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		msg := types.NewMsgGrantDiscount(req.PaymentContractId, req.DiscountId,
			req.Recipient, req.SenderDid)
		if rest.CheckBadRequestError(w, msg.ValidateBasic()) {
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

type revokeDiscountReq struct {
	BaseReq           rest.BaseReq    `json:"base_req" yaml:"base_req"`
	SenderDid         didexported.Did `json:"sender_did" yaml:"sender_did"`
	PaymentContractId string          `json:"payment_contract_id" yaml:"payment_contract_id"`
	Holder            sdk.AccAddress  `json:"holder" yaml:"holder"`
}

func newRevokeDiscountRequestHandlerFn(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req revokeDiscountReq
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		msg := types.NewMsgRevokeDiscount(req.PaymentContractId, req.Holder, req.SenderDid)
		if rest.CheckBadRequestError(w, msg.ValidateBasic()) {
			return
		}

		//utils.WriteGenerateStdTxResponse(w, ctx, req.BaseReq, []sdk.Msg{msg})
		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}

type effectPaymentReq struct {
	BaseReq           rest.BaseReq    `json:"base_req" yaml:"base_req"`
	SenderDid         didexported.Did `json:"sender_did" yaml:"sender_did"`
	PaymentContractId string          `json:"payment_contract_id" yaml:"payment_contract_id"`
}

func newEffectPaymentRequestHandlerFn(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req effectPaymentReq
		if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		msg := types.NewMsgEffectPayment(req.PaymentContractId, req.SenderDid)
		if rest.CheckBadRequestError(w, msg.ValidateBasic()) {
			return
		}

		tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}
