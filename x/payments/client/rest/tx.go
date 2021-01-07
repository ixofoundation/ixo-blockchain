package rest

import (
	"github.com/cosmos/cosmos-sdk/client"
	//"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/client/tx"
	//"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/gorilla/mux"
	"github.com/ixofoundation/ixo-blockchain/x/did"
	"github.com/ixofoundation/ixo-blockchain/x/payments/internal/types"
	"net/http"
)

func registerTxRoutes(cliCtx client.Context, r *mux.Router) {
	r.HandleFunc("/payments/create_payment_template", createPaymentTemplateRequestHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/payments/create_payment_contract", createPaymentContractRequestHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/payments/create_subscription", createSubscriptionRequestHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/payments/set_payment_contract_authorisation", setPaymentContractAuthorisationRequestHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/payments/grant_discount", grantDiscountRequestHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/payments/revoke_discount", revokeDiscountRequestHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/payments/effect_payment", effectPaymentRequestHandler(cliCtx)).Methods("POST")
}

type createPaymentTemplateReq struct {
	BaseReq         rest.BaseReq          `json:"base_req" yaml:"base_req"`
	CreatorDid      did.Did               `json:"creator_did" yaml:"creator_did"`
	PaymentTemplate types.PaymentTemplate `json:"payment_template" yaml:"payment_template"`
}

func createPaymentTemplateRequestHandler(ctx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createPaymentTemplateReq
		if !rest.ReadRESTReq(w, r, ctx.LegacyAmino, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		msg := types.NewMsgCreatePaymentTemplate(req.PaymentTemplate, req.CreatorDid)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		//utils.WriteGenerateStdTxResponse(w, ctx, req.BaseReq, []sdk.Msg{msg})
		tx.WriteGeneratedTxResponse(ctx, w, req.BaseReq, []sdk.Msg{msg}...)
	}
}

type createPaymentContractReq struct {
	BaseReq           rest.BaseReq       `json:"base_req" yaml:"base_req"`
	CreatorDid        did.Did            `json:"creator_did" yaml:"creator_did"`
	PaymentTemplateId string             `json:"payment_template_id" yaml:"payment_template_id"`
	PaymentContractId string             `json:"payment_contract_id" yaml:"payment_contract_id"`
	Payer             sdk.AccAddress     `json:"payer" yaml:"payer"`
	Recipients        types.Distribution `json:"recipients" yaml:"recipients"`
	CanDeauthorise    bool               `json:"can_deauthorise" yaml:"can_deauthorise"`
	DiscountId        sdk.Uint           `json:"discount_id" yaml:"discount_id"`
}

func createPaymentContractRequestHandler(ctx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createPaymentContractReq
		if !rest.ReadRESTReq(w, r, ctx.LegacyAmino, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		msg := types.NewMsgCreatePaymentContract(req.PaymentTemplateId,
			req.PaymentContractId, req.Payer, req.Recipients,
			req.CanDeauthorise, req.DiscountId, req.CreatorDid)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		//utils.WriteGenerateStdTxResponse(w, ctx, req.BaseReq, []sdk.Msg{msg})
		tx.WriteGeneratedTxResponse(ctx, w, req.BaseReq, []sdk.Msg{msg}...)
	}
}

type createSubscriptionReq struct {
	BaseReq           rest.BaseReq `json:"base_req" yaml:"base_req"`
	CreatorDid        did.Did      `json:"creator_did" yaml:"creator_did"`
	SubscriptionId    string       `json:"subscription_id" yaml:"subscription_id"`
	PaymentContractId string       `json:"payment_contract_id" yaml:"payment_contract_id"`
	MaxPeriods        sdk.Uint     `json:"max_periods" yaml:"max_periods"`
	Period            types.Period `json:"period" yaml:"period"`
}

func createSubscriptionRequestHandler(ctx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req createSubscriptionReq
		if !rest.ReadRESTReq(w, r, ctx.LegacyAmino, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		msg := types.NewMsgCreateSubscription(req.SubscriptionId, req.PaymentContractId,
			req.MaxPeriods, req.Period, req.CreatorDid)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		//utils.WriteGenerateStdTxResponse(w, ctx, req.BaseReq, []sdk.Msg{msg})
		tx.WriteGeneratedTxResponse(ctx, w, req.BaseReq, []sdk.Msg{msg}...)
	}
}

type setPaymentContractAuthorisationReq struct {
	BaseReq           rest.BaseReq `json:"base_req" yaml:"base_req"`
	PayerDid          did.Did      `json:"payer_did" yaml:"payer_did"`
	PaymentContractId string       `json:"payment_contract_id" yaml:"payment_contract_id"`
	Authorised        bool         `json:"authorised" yaml:"authorised"`
}

func setPaymentContractAuthorisationRequestHandler(ctx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req setPaymentContractAuthorisationReq
		if !rest.ReadRESTReq(w, r, ctx.LegacyAmino, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		msg := types.NewMsgSetPaymentContractAuthorisation(req.PaymentContractId,
			req.Authorised, req.PayerDid)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		//utils.WriteGenerateStdTxResponse(w, ctx, req.BaseReq, []sdk.Msg{msg})
		tx.WriteGeneratedTxResponse(ctx, w, req.BaseReq, []sdk.Msg{msg}...)
	}
}

type grantDiscountReq struct {
	BaseReq           rest.BaseReq   `json:"base_req" yaml:"base_req"`
	SenderDid         did.Did        `json:"sender_did" yaml:"sender_did"`
	PaymentContractId string         `json:"payment_contract_id" yaml:"payment_contract_id"`
	DiscountId        sdk.Uint       `json:"discount_id" yaml:"discount_id"`
	Recipient         sdk.AccAddress `json:"recipient" yaml:"recipient"`
}

func grantDiscountRequestHandler(ctx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req grantDiscountReq
		if !rest.ReadRESTReq(w, r, ctx.LegacyAmino, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		msg := types.NewMsgGrantDiscount(req.PaymentContractId, req.DiscountId,
			req.Recipient, req.SenderDid)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		//utils.WriteGenerateStdTxResponse(w, ctx, req.BaseReq, []sdk.Msg{msg})
		tx.WriteGeneratedTxResponse(ctx, w, req.BaseReq, []sdk.Msg{msg}...)
	}
}

type revokeDiscountReq struct {
	BaseReq           rest.BaseReq   `json:"base_req" yaml:"base_req"`
	SenderDid         did.Did        `json:"sender_did" yaml:"sender_did"`
	PaymentContractId string         `json:"payment_contract_id" yaml:"payment_contract_id"`
	Holder            sdk.AccAddress `json:"holder" yaml:"holder"`
}

func revokeDiscountRequestHandler(ctx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req revokeDiscountReq
		if !rest.ReadRESTReq(w, r, ctx.LegacyAmino, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		msg := types.NewMsgRevokeDiscount(req.PaymentContractId, req.Holder, req.SenderDid)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		//utils.WriteGenerateStdTxResponse(w, ctx, req.BaseReq, []sdk.Msg{msg})
		tx.WriteGeneratedTxResponse(ctx, w, req.BaseReq, []sdk.Msg{msg}...)
	}
}

type effectPaymentReq struct {
	BaseReq           rest.BaseReq `json:"base_req" yaml:"base_req"`
	SenderDid         did.Did      `json:"sender_did" yaml:"sender_did"`
	PaymentContractId string       `json:"payment_contract_id" yaml:"payment_contract_id"`
}

func effectPaymentRequestHandler(ctx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req effectPaymentReq
		if !rest.ReadRESTReq(w, r, ctx.LegacyAmino, &req) {
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		msg := types.NewMsgEffectPayment(req.PaymentContractId, req.SenderDid)
		if err := msg.ValidateBasic(); err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		//utils.WriteGenerateStdTxResponse(w, ctx, req.BaseReq, []sdk.Msg{msg})
		tx.WriteGeneratedTxResponse(ctx, w, req.BaseReq, []sdk.Msg{msg}...)
	}
}
