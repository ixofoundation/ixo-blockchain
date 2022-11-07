package rest

import (
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	govrest "github.com/cosmos/cosmos-sdk/x/gov/client/rest"
)

func ProposalRESTHandler(clientCtx client.Context) govrest.ProposalRESTHandler {
	return govrest.ProposalRESTHandler{
		SubRoute: "token_param_change",
		Handler:  postProposalHandlerFn(clientCtx),
	}
}

func postProposalHandlerFn(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// var req paramscutils.ParamChangeProposalReq
		// if !rest.ReadRESTReq(w, r, clientCtx.LegacyAmino, &req) {
		// 	return
		// }

		// req.BaseReq = req.BaseReq.Sanitize()
		// if !req.BaseReq.ValidateBasic(w) {
		// 	return
		// }
		//
		// content := types.NewParams("")

		// msg, err := govtypes.NewMsgSubmitProposal(content, req.Deposit, req.Proposer)
		// if rest.CheckBadRequestError(w, err) {
		// 	return
		// }
		// if rest.CheckBadRequestError(w, msg.ValidateBasic()) {
		// 	return
		// }

		// tx.WriteGeneratedTxResponse(clientCtx, w, req.BaseReq, msg)
	}
}
