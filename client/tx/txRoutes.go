package tx
//
//import (
//	"encoding/base64"
//	"encoding/hex"
//	//"github.com/cosmos/cosmos-sdk/client/context"
//	sdk "github.com/cosmos/cosmos-sdk/types"
//	"github.com/cosmos/cosmos-sdk/types/rest"
//	"github.com/cosmos/cosmos-sdk/x/auth"
//	//"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
//	"github.com/gorilla/mux"
//	"github.com/ixofoundation/ixo-blockchain/x/did"
//	"github.com/ixofoundation/ixo-blockchain/x/ixo"
//	"github.com/ixofoundation/ixo-blockchain/x/project"
//	"io/ioutil"
//	"net/http"
//	"strings"
//)
//
//func RegisterTxRoutes(cliCtx context.CLIContext, r *mux.Router) {
//	r.HandleFunc("/txs/sign_data", SignDataRequest(cliCtx)).Methods("POST")
//	r.HandleFunc("/txs/decode", DecodeTxRequestHandlerFn(cliCtx)).Methods("POST")
//}
//
//type SignDataReq struct {
//	Msg    string `json:"msg" yaml:"msg"`
//	PubKey string `json:"pub_key" yaml:"pub_key"`
//}
//
//type SignDataResponse struct {
//	SignBytes string      `json:"sign_bytes" yaml:"sign_bytes"`
//	Fee       auth.StdFee `json:"fee" yaml:"fee"`
//}
//
//func SignDataRequest(cliCtx context.CLIContext) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		var req SignDataReq
//
//		body, err := ioutil.ReadAll(r.Body)
//		if err != nil {
//			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
//			return
//		}
//
//		err = cliCtx.Codec.UnmarshalJSON(body, &req)
//		if err != nil {
//			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
//			return
//		}
//
//		msgBytes, err := hex.DecodeString(strings.TrimPrefix(req.Msg, "0x"))
//		if err != nil {
//			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
//			return
//		}
//
//		var msg sdk.Msg
//		err = cliCtx.Codec.UnmarshalJSON(msgBytes, &msg)
//		if err != nil {
//			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
//			return
//		}
//
//		// all messages must be of type ixo.IxoMsg
//		ixoMsg, ok := msg.(ixo.IxoMsg)
//		if !ok {
//			rest.WriteErrorResponse(w, http.StatusBadRequest, "msg must be ixo.IxoMsg")
//			return
//		}
//		msgs := []sdk.Msg{ixoMsg}
//
//		// obtain stdSignMsg (create-project is a special case)
//		var stdSignMsg auth.StdSignMsg
//		switch ixoMsg.Type() {
//		case project.TypeMsgCreateProject:
//			stdSignMsg = ixoMsg.(project.MsgCreateProject).ToStdSignMsg(
//				project.MsgCreateProjectTotalFee)
//		default:
//			// Deduce and set signer address
//			signerAddress := did.VerifyKeyToAddr(req.PubKey)
//			cliCtx = cliCtx.WithFromAddress(signerAddress)
//
//			txBldr, err := utils.PrepareTxBuilder(auth.NewTxBuilderFromCLI(cliCtx.Input), cliCtx)
//			if err != nil {
//				rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
//				return
//			}
//
//			// Build the transaction
//			stdSignMsg, err = txBldr.BuildSignMsg(msgs)
//			if err != nil {
//				rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
//				return
//			}
//
//			// Create dummy tx with blank signature for fee approximation
//			signature := auth.StdSignature{}
//			tx := auth.NewStdTx(stdSignMsg.Msgs, stdSignMsg.Fee,
//				[]auth.StdSignature{signature}, stdSignMsg.Memo)
//
//			// Approximate fee
//			fee, err := ixo.ApproximateFeeForTx(cliCtx, tx, txBldr.ChainID())
//			if err != nil {
//				rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
//				return
//			}
//			stdSignMsg.Fee = fee
//		}
//
//		// Produce response from sign bytes and fees
//		output := SignDataResponse{
//			SignBytes: string(stdSignMsg.Bytes()),
//			Fee:       stdSignMsg.Fee,
//		}
//
//		rest.PostProcessResponseBare(w, cliCtx, output)
//	}
//}
//
//type (
//	// DecodeReq defines a tx decoding request.
//	DecodeReq struct {
//		Tx string `json:"tx"`
//	}
//
//	// DecodeResp defines a tx decoding response.
//	DecodeResp auth.StdTx
//)
//
//func DecodeTxRequestHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		var req DecodeReq
//
//		body, err := ioutil.ReadAll(r.Body)
//		if err != nil {
//			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
//			return
//		}
//
//		err = cliCtx.Codec.UnmarshalJSON(body, &req)
//		if err != nil {
//			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
//			return
//		}
//
//		txBytes, err := base64.StdEncoding.DecodeString(req.Tx)
//		if err != nil {
//			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
//			return
//		}
//
//		var stdTx auth.StdTx
//		err = cliCtx.Codec.UnmarshalBinaryLengthPrefixed(txBytes, &stdTx)
//		if err != nil {
//			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
//			return
//		}
//
//		response := DecodeResp(stdTx)
//		rest.PostProcessResponse(w, cliCtx, response)
//	}
//}
