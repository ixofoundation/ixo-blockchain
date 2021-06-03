package tx

import (
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	clientrest "github.com/cosmos/cosmos-sdk/client/rest"
	clienttx "github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"
	"github.com/cosmos/cosmos-sdk/x/auth/signing"
	"github.com/gorilla/mux"
	didexported "github.com/ixofoundation/ixo-blockchain/x/did/exported"
	ixotypes "github.com/ixofoundation/ixo-blockchain/x/ixo/types"
	projecttypes "github.com/ixofoundation/ixo-blockchain/x/project/types"
	"io/ioutil"
	"net/http"
	"strings"
)

var (
	approximationGasAdjustment = float64(1.5)
	expectedMinGasPrices       = "0.025" + ixotypes.IxoNativeToken
)

func RegisterTxRoutes(clientCtx client.Context, r *mux.Router) {
	r.HandleFunc("/txs/sign_data", SignDataRequest(clientCtx)).Methods("POST")
	r.HandleFunc("/txs/decode", DecodeTxRequestHandlerFn(clientCtx)).Methods("POST")
}

type SignDataReq struct {
	Msg    string `json:"msg" yaml:"msg"`
	PubKey string `json:"pub_key" yaml:"pub_key"`
}

type SignDataResponse struct {
	SignBytes string          `json:"sign_bytes" yaml:"sign_bytes"`
	Fee       legacytx.StdFee `json:"fee" yaml:"fee"`
}

func SignDataRequest(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req SignDataReq

		body, err := ioutil.ReadAll(r.Body)
		if rest.CheckBadRequestError(w, err){
			return
		}

		err = clientCtx.LegacyAmino.UnmarshalJSON(body, &req)
		if rest.CheckBadRequestError(w, err){
			return
		}

		msgBytes, err := hex.DecodeString(strings.TrimPrefix(req.Msg, "0x"))
		if rest.CheckInternalServerError(w, err){
			return
		}

		var msg sdk.Msg
		err = clientCtx.LegacyAmino.UnmarshalJSON(msgBytes, &msg)
		if rest.CheckBadRequestError(w, err){
			return
		}

		// all messages must be of type ixo.IxoMsg
		ixoMsg, ok := msg.(ixotypes.IxoMsg)
		if !ok {
			rest.WriteErrorResponse(w, http.StatusBadRequest, "msg must be ixo.IxoMsg")
			return
		}

		output := SignDataResponse{}

		switch ixoMsg.Type(){
		// TODO (Stef) Case not working
		case projecttypes.TypeMsgCreateProject:
			var stdSignMsg legacytx.StdSignMsg
			stdSignMsg = ixoMsg.(*projecttypes.MsgCreateProject).ToStdSignMsg(
				projecttypes.MsgCreateProjectTotalFee)

			output.SignBytes = string(stdSignMsg.Bytes())
			output.Fee       = stdSignMsg.Fee
		default:
			// Deduce and set signer address
			signerAddress := didexported.VerifyKeyToAddr(req.PubKey)
			clientCtx = clientCtx.WithFromAddress(signerAddress)

			// Set gas adjustment and fees
			gasAdjustment := approximationGasAdjustment
			fees := sdk.NewCoins(sdk.NewCoin(ixotypes.IxoNativeToken, sdk.OneInt()))

			chainId := clientCtx.ChainID
			txf := clienttx.Factory{}.
				WithTxConfig(clientCtx.TxConfig).
				WithAccountRetriever(clientCtx.AccountRetriever).
				WithKeybase(clientCtx.Keyring).
				WithChainID(chainId).
				WithSimulateAndExecute(false).
				WithMemo("").
				WithGasPrices(expectedMinGasPrices)

			txf, err := clienttx.PrepareFactory(clientCtx, txf)
			if rest.CheckInternalServerError(w, err) {
				return
			}

			txfForGasCalc := clienttx.Factory{}.
				WithTxConfig(clientCtx.TxConfig).
				WithAccountRetriever(clientCtx.AccountRetriever).
				WithKeybase(clientCtx.Keyring).
				WithChainID(chainId).
				WithSimulateAndExecute(true).
				WithMemo("").
				WithGasAdjustment(gasAdjustment).
				WithFees(fees.String())

			_, gasAmt, err := clienttx.CalculateGas(clientCtx.QueryWithData, txfForGasCalc, msg)
			txf = txf.WithGas(gasAmt)

			tx, err := clienttx.BuildUnsignedTx(txf, msg)
			if rest.CheckBadRequestError(w, err) {
				return
			}

			stdFee := legacytx.NewStdFee(gasAmt, tx.GetTx().GetFee())
			bytes := legacytx.StdSignBytes(txf.ChainID(), txf.AccountNumber(), txf.Sequence(),
				txf.TimeoutHeight(), stdFee, []sdk.Msg{msg}, txf.Memo())

			// Produce response from sign bytes and fees
			output.SignBytes = string(bytes)
			output.Fee       = stdFee
		}

		rest.PostProcessResponseBare(w, clientCtx, output)
	}
}


type (
	// DecodeReq defines a tx decoding request.
	DecodeReq struct {
		Tx string `json:"tx"`
	}

	// DecodeResp defines a tx decoding response.
	DecodeResp legacytx.StdTx
)

// DecodeTxRequestHandlerFn returns the decode tx REST handler. In particular,
// it takes base64-decoded bytes, decodes it from the Amino wire protocol,
// and responds with a json-formatted transaction.
// Copied from Cosmos SDK.
func DecodeTxRequestHandlerFn(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req DecodeReq

		body, err := ioutil.ReadAll(r.Body)
		if rest.CheckBadRequestError(w, err) {
			return
		}

		// NOTE: amino is used intentionally here, don't migrate it
		err = clientCtx.LegacyAmino.UnmarshalJSON(body, &req)
		if rest.CheckBadRequestError(w, err) {
			return
		}

		txBytes, err := base64.StdEncoding.DecodeString(req.Tx)
		if rest.CheckBadRequestError(w, err) {
			return
		}

		stdTx, err := convertToStdTx(w, clientCtx, txBytes)
		if err != nil {
			// Error is already returned by convertToStdTx.
			return
		}

		response := DecodeResp(stdTx)

		err = checkAminoMarshalError(clientCtx, response, "/cosmos/tx/v1beta1/txs/decode")
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())

			return
		}

		rest.PostProcessResponse(w, clientCtx, response)
	}
}

// convertToStdTx converts tx proto binary bytes retrieved from Tendermint into
// a StdTx. Returns the StdTx, as well as a flag denoting if the function
// successfully converted or not.
// Copied from Cosmos SDK.
func convertToStdTx(w http.ResponseWriter, clientCtx client.Context, txBytes []byte) (legacytx.StdTx, error) {
	txI, err := clientCtx.TxConfig.TxDecoder()(txBytes)
	if rest.CheckBadRequestError(w, err) {
		return legacytx.StdTx{}, err
	}

	tx, ok := txI.(signing.Tx)
	if !ok {
		rest.WriteErrorResponse(w, http.StatusBadRequest, fmt.Sprintf("%+v is not backwards compatible with %T", tx, legacytx.StdTx{}))
		return legacytx.StdTx{}, sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "expected %T, got %T", (signing.Tx)(nil), txI)
	}

	stdTx, err := clienttx.ConvertTxToStdTx(clientCtx.LegacyAmino, tx)
	if rest.CheckBadRequestError(w, err) {
		return legacytx.StdTx{}, err
	}

	return stdTx, nil
}

// checkAminoMarshalError checks if there are errors with marshalling non-amino
// txs with amino.
// Copied from Cosmos SDK.
func checkAminoMarshalError(ctx client.Context, resp interface{}, grpcEndPoint string) error {
	// LegacyAmino used intentionally here to handle the SignMode errors
	marshaler := ctx.LegacyAmino

	_, err := marshaler.MarshalJSON(resp)
	if err != nil {

		// If there's an unmarshalling error, we assume that it's because we're
		// using amino to unmarshal a non-amino tx.
		return fmt.Errorf("this transaction cannot be displayed via legacy REST endpoints, because it does not support"+
			" Amino serialization. Please either use CLI, gRPC, gRPC-gateway, or directly query the Tendermint RPC"+
			" endpoint to query this transaction. The new REST endpoint (via gRPC-gateway) is %s. Please also see the"+
			"REST endpoints migration guide at %s for more info", grpcEndPoint, clientrest.DeprecationURL)

	}

	return nil
}