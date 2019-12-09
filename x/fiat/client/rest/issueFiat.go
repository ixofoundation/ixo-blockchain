package rest

import (
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	// rest2 "github.com/ixofoundation/ixo-cosmos/client/rest"
	"github.com/ixofoundation/ixo-cosmos/x/fiat/client"
	fiatTypes "github.com/ixofoundation/ixo-cosmos/x/fiat/internal/types"
)

type IssueFiatReq struct {
	BaseReq           rest.BaseReq `json:"base_req"`
	To                string       `json:"to" valid:"required~Enter the ToAddress,matches(^commit[a-z0-9]{39}$)~ToAddress is Invalid"`
	GasAdjustment     string       `json:"gasAdjustment"`
	TransactionID     string       `json:"transactionID" valid:"required~Enter the TransactionID,  matches(^[A-Z0-9]+$)~transactionID is Invalid,length(2|40)~TransactionID length should be 2 to 40"`
	TransactionAmount int64        `json:"transactionAmount" valid:"required~Enter the TransactionAmount,matches(^[1-9]{1}[0-9]*$)~Invalid TransactionAmount"`
	Password          string       `json:"password" valid:"required~Enter the Password"`
	Mode              string       `json:"mode"`
}

func IssueFiatHandlerFunction(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req IssueFiatReq
		if !rest.ReadRESTReq(w, r, cliCtx.Codec, &req) {
			return
		}

		_, err := govalidator.ValidateStruct(req)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, sdk.NewError(fiatTypes.DefaultCodeSpace, http.StatusBadRequest, err.Error()).Error())
			return
		}

		req.BaseReq = req.BaseReq.Sanitize()
		if !req.BaseReq.ValidateBasic(w) {
			return
		}

		fromAddr, name, err := context.GetFromFields(req.BaseReq.From, false)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		cliCtx = cliCtx.WithFromAddress(fromAddr)
		cliCtx = cliCtx.WithFromName(name)

		to, err := sdk.AccAddressFromBech32(req.To)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		_ = client.BuildIssueFiatMsg(fromAddr, to, req.TransactionID, req.TransactionAmount)

		// output, err := rest2.SignAndBroadcast(req.BaseReq, cliCtx, req.Mode, req.Password, []sdk.Msg{msg})
		// if err != nil {
		// 	rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
		// 	return
		// }
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte("output"))
	}

}
