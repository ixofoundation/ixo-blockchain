package rest

import (
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/gorilla/mux"
	"github.com/ixofoundation/ixo-blockchain/x/did"
	"net/http"
	"strings"

	"github.com/ixofoundation/ixo-blockchain/x/payments/internal/types"
)

func registerTxRoutes(cliCtx context.CLIContext, r *mux.Router) {
	r.HandleFunc("/payments/create_payment_template", createPaymentTemplateHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/payments/create_payment_contract", createPaymentContractHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/payments/create_subscription", createSubscriptionHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/payments/set_payment_contract_authorisation", setPaymentContractAuthorisationHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/payments/grant_discount", grantDiscountHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/payments/revoke_discount", revokeDiscountHandler(cliCtx)).Methods("POST")
	r.HandleFunc("/payments/effect_payment", effectPaymentHandler(cliCtx)).Methods("POST")
}

const (
	TRUE  = "true"
	FALSE = "false"
)

func parseBool(boolStr, boolName string) (bool, error) {
	boolStr = strings.ToLower(strings.TrimSpace(boolStr))
	if boolStr == TRUE {
		return true, nil
	} else if boolStr == FALSE {
		return false, nil
	} else {
		return false, types.ErrInvalidArgument(types.DefaultCodespace, ""+
			fmt.Sprintf("%s is not a valid bool (true/false)", boolName))
	}
}

type CreatePaymentTemplateReq struct {
	BaseReq         rest.BaseReq          `json:"base_req" yaml:"base_req"`
	CreatorDid      did.Did               `json:"creator_did" yaml:"creator_did"`
	PaymentTemplate types.PaymentTemplate `json:"payment_template" yaml:"payment_template"`
}

func createPaymentTemplateHandler(ctx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req CreatePaymentTemplateReq
		if !rest.ReadRESTReq(w, r, ctx.Codec, &req) {
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
		utils.WriteGenerateStdTxResponse(w, ctx, req.BaseReq, []sdk.Msg{msg})
	}
}

type CreatePaymentContractReq struct {
	BaseReq           rest.BaseReq   `json:"base_req" yaml:"base_req"`
	CreatorDid        did.Did        `json:"creator_did" yaml:"creator_did"`
	PaymentTemplateId string         `json:"payment_template_id" yaml:"payment_template_id"`
	PaymentContractId string         `json:"payment_contract_id" yaml:"payment_contract_id"`
	Payer             sdk.AccAddress `json:"payer" yaml:"payer"`
	CanDeauthorise    bool           `json:"can_deauthorise" yaml:"can_deauthorise"`
	DiscountId        sdk.Uint       `json:"discount_id" yaml:"discount_id"`
}

func createPaymentContractHandler(ctx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req CreatePaymentContractReq
		if !rest.ReadRESTReq(w, r, ctx.Codec, &req) {
			return
		}
		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}
		msg := types.NewMsgCreatePaymentContract(req.PaymentTemplateId, req.PaymentContractId,
			req.Payer, req.CanDeauthorise, req.DiscountId, req.CreatorDid)
		utils.WriteGenerateStdTxResponse(w, ctx, req.BaseReq, []sdk.Msg{msg})
	}
}

type CreateSubscriptionReq struct {
	BaseReq           rest.BaseReq `json:"base_req" yaml:"base_req"`
	CreatorDid        did.Did      `json:"creator_did" yaml:"creator_did"`
	SubscriptionId    string       `json:"subscription_id" yaml:"subscription_id"`
	PaymentContractId string       `json:"payment_contract_id" yaml:"payment_contract_id"`
	MaxPeriods        sdk.Uint     `json:"max_periods" yaml:"max_periods"`
	Period            types.Period `json:"period" yaml:"period"`
}

func createSubscriptionHandler(ctx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req CreateSubscriptionReq
		if !rest.ReadRESTReq(w, r, ctx.Codec, &req) {
			return
		}
		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}
		msg := types.NewMsgCreateSubscription(req.SubscriptionId, req.PaymentContractId,
			req.MaxPeriods, req.Period, req.CreatorDid)
		utils.WriteGenerateStdTxResponse(w, ctx, req.BaseReq, []sdk.Msg{msg})
	}
}

type SetPaymentContractAuthorisationReq struct {
	BaseReq           rest.BaseReq `json:"base_req" yaml:"base_req"`
	PayerDid          did.Did      `json:"payer_did" yaml:"payer_did"`
	PaymentContractId string       `json:"payment_contract_id" yaml:"payment_contract_id"`
	Authorised        bool         `json:"authorised" yaml:"authorised"`
}

func setPaymentContractAuthorisationHandler(ctx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req SetPaymentContractAuthorisationReq
		if !rest.ReadRESTReq(w, r, ctx.Codec, &req) {
			return
		}
		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}
		msg := types.NewMsgSetPaymentContractAuthorisation(req.PaymentContractId,
			req.Authorised, req.PayerDid)
		utils.WriteGenerateStdTxResponse(w, ctx, req.BaseReq, []sdk.Msg{msg})
	}
}

type GrantDiscountReq struct {
	BaseReq           rest.BaseReq   `json:"base_req" yaml:"base_req"`
	SenderDid         did.Did        `json:"sender_did" yaml:"sender_did"`
	PaymentContractId string         `json:"payment_contract_id" yaml:"payment_contract_id"`
	DiscountId        sdk.Uint       `json:"discount_id" yaml:"discount_id"`
	Recipient         sdk.AccAddress `json:"recipient" yaml:"recipient"`
}

func grantDiscountHandler(ctx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req GrantDiscountReq
		if !rest.ReadRESTReq(w, r, ctx.Codec, &req) {
			return
		}
		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}
		msg := types.NewMsgGrantDiscount(req.PaymentContractId, req.DiscountId,
			req.Recipient, req.SenderDid)
		utils.WriteGenerateStdTxResponse(w, ctx, req.BaseReq, []sdk.Msg{msg})
	}
}

type RevokeDiscountReq struct {
	BaseReq           rest.BaseReq   `json:"base_req" yaml:"base_req"`
	SenderDid         did.Did        `json:"sender_did" yaml:"sender_did"`
	PaymentContractId string         `json:"payment_contract_id" yaml:"payment_contract_id"`
	Holder            sdk.AccAddress `json:"holder" yaml:"holder"`
}

func revokeDiscountHandler(ctx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req RevokeDiscountReq
		if !rest.ReadRESTReq(w, r, ctx.Codec, &req) {
			return
		}
		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}
		msg := types.NewMsgRevokeDiscount(req.PaymentContractId, req.Holder, req.SenderDid)
		utils.WriteGenerateStdTxResponse(w, ctx, req.BaseReq, []sdk.Msg{msg})
	}
}

type EffectPaymentReq struct {
	BaseReq           rest.BaseReq `json:"base_req" yaml:"base_req"`
	SenderDid         did.Did      `json:"sender_did" yaml:"sender_did"`
	PaymentContractId string       `json:"payment_contract_id" yaml:"payment_contract_id"`
}

func effectPaymentHandler(ctx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req EffectPaymentReq
		if !rest.ReadRESTReq(w, r, ctx.Codec, &req) {
			return
		}
		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}
		msg := types.NewMsgEffectPayment(req.PaymentContractId, req.SenderDid)
		utils.WriteGenerateStdTxResponse(w, ctx, req.BaseReq, []sdk.Msg{msg})
	}
}
