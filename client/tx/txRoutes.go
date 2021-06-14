package tx

import (
	"encoding/hex"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	clienttx "github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/auth/legacy/legacytx"
	"github.com/gorilla/mux"
	didexported "github.com/ixofoundation/ixo-blockchain/x/did/exported"
	ixotypes "github.com/ixofoundation/ixo-blockchain/x/ixo/types"
	projecttypes "github.com/ixofoundation/ixo-blockchain/x/project/types"
)

var (
	approximationGasAdjustment = float64(1.5)
	expectedMinGasPrices       = "0.025" + ixotypes.IxoNativeToken
)

func RegisterTxRoutes(clientCtx client.Context, r *mux.Router) {
	r.HandleFunc("/txs/sign_data", SignDataRequest(clientCtx)).Methods("POST")
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
		case projecttypes.TypeMsgCreateProject:
			var stdSignMsg legacytx.StdSignMsg
			stdSignMsg = ixoMsg.(*projecttypes.MsgCreateProject).ToStdSignMsg(
				projecttypes.MsgCreateProjectTotalFee)
			stdSignMsg.ChainID = clientCtx.ChainID

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

			txfForGasCalc, err = clienttx.PrepareFactory(clientCtx, txfForGasCalc)
			if rest.CheckInternalServerError(w, err) {
				return
			}

			_, gasAmt, err := clienttx.CalculateGas(clientCtx.QueryWithData, txfForGasCalc, msg)
			if rest.CheckBadRequestError(w, err) {
				return
			}
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
