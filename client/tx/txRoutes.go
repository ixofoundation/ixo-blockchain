package tx
//
//import (
//	"encoding/base64"
//	"encoding/hex"
//	"fmt"
//	"github.com/cosmos/cosmos-sdk/client"
//	"github.com/cosmos/cosmos-sdk/client/flags"
//	sdk "github.com/cosmos/cosmos-sdk/types"
//	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
//	"github.com/cosmos/cosmos-sdk/types/rest"
//	"github.com/cosmos/cosmos-sdk/x/auth"
//	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
//	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"
//	"github.com/cosmos/cosmos-sdk/x/auth/signing"
//	"github.com/cosmos/cosmos-sdk/x/auth/types"
//	"github.com/gorilla/mux"
//	"github.com/ixofoundation/ixo-blockchain/x/did"
//	"github.com/ixofoundation/ixo-blockchain/x/ixo"
//	ixotypes "github.com/ixofoundation/ixo-blockchain/x/ixo/types"
//	"github.com/ixofoundation/ixo-blockchain/x/project"
//	"github.com/spf13/viper"
//	"io/ioutil"
//	"net/http"
//	"strings"
//	clienttx "github.com/cosmos/cosmos-sdk/client/tx"
//	clientrest "github.com/cosmos/cosmos-sdk/client/rest"
//
//)
//
//func RegisterTxRoutes(clientCtx client.Context, r *mux.Router) {
//	r.HandleFunc("/txs/sign_data", SignDataRequest(clientCtx)).Methods("POST")
//	r.HandleFunc("/txs/decode", DecodeTxRequestHandlerFn(clientCtx)).Methods("POST")
//}
//
//type SignDataReq struct {
//	Msg    string `json:"msg" yaml:"msg"`
//	PubKey string `json:"pub_key" yaml:"pub_key"`
//}
//
//type SignDataResponse struct {
//	SignBytes string          `json:"sign_bytes" yaml:"sign_bytes"`
//	Fee       legacytx.StdFee `json:"fee" yaml:"fee"`
//}
//
//func SignDataRequest(clientCtx client.Context) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		var req SignDataReq
//
//		body, err := ioutil.ReadAll(r.Body)
//		if err != nil {
//			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
//			return
//		}
//
//		err = clientCtx.LegacyAmino.UnmarshalJSON(body, &req)
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
//		err = clientCtx.LegacyAmino.UnmarshalJSON(msgBytes, &msg)
//		if err != nil {
//			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
//			return
//		}
//
//		// all messages must be of type ixo.IxoMsg
//		ixoMsg, ok := msg.(ixotypes.IxoMsg)
//		if !ok {
//			rest.WriteErrorResponse(w, http.StatusBadRequest, "msg must be ixo.IxoMsg")
//			return
//		}
//		msgs := []sdk.Msg{ixoMsg}
//
//		// obtain stdSignMsg (create-project is a special case)
//		var stdSignMsg legacytx.StdSignMsg
//		switch ixoMsg.Type() {
//		case project.TypeMsgCreateProject:
//			stdSignMsg = ixoMsg.(project.MsgCreateProject).ToStdSignMsg(
//				project.MsgCreateProjectTotalFee)
//		default:
//			// Deduce and set signer address
//			signerAddress := did.VerifyKeyToAddr(req.PubKey)
//			clientCtx = clientCtx.WithFromAddress(signerAddress)
//
//			// Create tx builder with:
//			// - TxEncoder set to GetTxEncoder(...)
//			// - Acc no. and sequence set to 0 (set by PrepareTxBuilder)
//			// - Gas and gas adjustment set to 0 (set by ApproximateFeeForTx)
//			// - Simulate set to false
//			// - Chain ID obtained using viper.GetString(...)
//			// - Memo set to empty string (custom memo not supported)
//			// - Fees and gas prices set to nil (set by ApproximateFeeForTx)
//			chainId := viper.GetString(flags.FlagChainID)
//			txBldr := types.NewTxBuilder(
//				utils.GetTxEncoder(clientCtx.Codec),
//				0, 0, 0, 0, false, chainId, "", nil, nil)
//
//			txBldr, err = utils.PrepareTxBuilder(txBldr, clientCtx)
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
//			fee, err := ixo.ApproximateFeeForTx(clientCtx, tx, txBldr.ChainID())
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
//		rest.PostProcessResponseBare(w, clientCtx, output)
//	}
//}
//
////type (
////	// DecodeReq defines a tx decoding request.
////	DecodeReq struct {
////		Tx string `json:"tx"`
////	}
////
////	// DecodeResp defines a tx decoding response.
////	DecodeResp auth.StdTx
////)
//
////func DecodeTxRequestHandlerFn(clientCtx client.Context) http.HandlerFunc {
////	return func(w http.ResponseWriter, r *http.Request) {
////		var req DecodeReq
////
////		body, err := ioutil.ReadAll(r.Body)
////		if err != nil {
////			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
////			return
////		}
////
////		err = clientCtx.LegacyAmino.UnmarshalJSON(body, &req)
////		if err != nil {
////			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
////			return
////		}
////
////		txBytes, err := base64.StdEncoding.DecodeString(req.Tx)
////		if err != nil {
////			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
////			return
////		}
////
////		var stdTx auth.StdTx
////		err = clientCtx.LegacyAmino.UnmarshalBinaryLengthPrefixed(txBytes, &stdTx)
////		if err != nil {
////			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
////			return
////		}
////
////		response := DecodeResp(stdTx)
////		rest.PostProcessResponse(w, clientCtx, response)
////	}
////}
//
//type (
//	// DecodeReq defines a tx decoding request.
//	DecodeReq struct {
//		Tx string `json:"tx"`
//	}
//
//	// DecodeResp defines a tx decoding response.
//	DecodeResp legacytx.StdTx
//)
//
//// DecodeTxRequestHandlerFn returns the decode tx REST handler. In particular,
//// it takes base64-decoded bytes, decodes it from the Amino wire protocol,
//// and responds with a json-formatted transaction.
//// Copied from Cosmos SDK.
//func DecodeTxRequestHandlerFn(clientCtx client.Context) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//		var req DecodeReq
//
//		body, err := ioutil.ReadAll(r.Body)
//		if rest.CheckBadRequestError(w, err) {
//			return
//		}
//
//		// NOTE: amino is used intentionally here, don't migrate it
//		err = clientCtx.LegacyAmino.UnmarshalJSON(body, &req)
//		if rest.CheckBadRequestError(w, err) {
//			return
//		}
//
//		txBytes, err := base64.StdEncoding.DecodeString(req.Tx)
//		if rest.CheckBadRequestError(w, err) {
//			return
//		}
//
//		stdTx, err := convertToStdTx(w, clientCtx, txBytes)
//		if err != nil {
//			// Error is already returned by convertToStdTx.
//			return
//		}
//
//		response := DecodeResp(stdTx)
//
//		err = checkAminoMarshalError(clientCtx, response, "/ixo/txs/decode")
//		if err != nil {
//			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
//
//			return
//		}
//
//		rest.PostProcessResponse(w, clientCtx, response)
//	}
//}
//
//// convertToStdTx converts tx proto binary bytes retrieved from Tendermint into
//// a StdTx. Returns the StdTx, as well as a flag denoting if the function
//// successfully converted or not.
//func convertToStdTx(w http.ResponseWriter, clientCtx client.Context, txBytes []byte) (legacytx.StdTx, error) {
//	txI, err := clientCtx.TxConfig.TxDecoder()(txBytes)
//	if rest.CheckBadRequestError(w, err) {
//		return legacytx.StdTx{}, err
//	}
//
//	tx, ok := txI.(signing.Tx)
//	if !ok {
//		rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("%+v is not backwards compatible with %T", tx, legacytx.StdTx{}))
//		return legacytx.StdTx{}, sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "expected %T, got %T", (signing.Tx)(nil), txI)
//	}
//
//	stdTx, err := clienttx.ConvertTxToStdTx(clientCtx.LegacyAmino, tx)
//	if rest.CheckBadRequestError(w, err) {
//		return legacytx.StdTx{}, err
//	}
//
//	return stdTx, nil
//}
//
//// checkAminoMarshalError checks if there are errors with marshalling non-amino
//// txs with amino.
//func checkAminoMarshalError(ctx client.Context, resp interface{}, grpcEndPoint string) error {
//	// LegacyAmino used intentionally here to handle the SignMode errors
//	marshaler := ctx.LegacyAmino
//
//	_, err := marshaler.MarshalJSON(resp)
//	if err != nil {
//
//		// If there's an unmarshalling error, we assume that it's because we're
//		// using amino to unmarshal a non-amino tx.
//		return fmt.Errorf("this transaction cannot be displayed via legacy REST endpoints, because it does not support"+
//			" Amino serialization. Please either use CLI, gRPC, gRPC-gateway, or directly query the Tendermint RPC"+
//			" endpoint to query this transaction. The new REST endpoint (via gRPC-gateway) is %s. Please also see the"+
//			"REST endpoints migration guide at %s for more info", grpcEndPoint, clientrest.DeprecationURL)
//
//	}
//
//	return nil
//}