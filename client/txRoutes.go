package client

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/ixofoundation/ixo-blockchain/x/project"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/gorilla/mux"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	core "github.com/tendermint/tendermint/rpc/core/types"

	genutilrest "github.com/cosmos/cosmos-sdk/x/genutil/client/rest"
	"github.com/ixofoundation/ixo-blockchain/x/ixo"
)

func QueryTxCmd(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{

		Use:   "tx [hash]",
		Short: "Query for a transaction by hash in a committed block",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			output, err := QueryTx(cliCtx, args[0])
			if err != nil {
				return err
			}

			if output.Empty() {
				return fmt.Errorf("No transaction found with hash %s", args[0])
			}

			return cliCtx.PrintOutput(output)
		},
	}

	cmd.Flags().StringP(flags.FlagNode, "n", "tcp://localhost:26657", "Node to connect to")
	_ = viper.BindPFlag(flags.FlagNode, cmd.Flags().Lookup(flags.FlagNode))
	cmd.Flags().Bool(flags.FlagTrustNode, false, "Trust connected full node (don't verify proofs for responses)")
	_ = viper.BindPFlag(flags.FlagTrustNode, cmd.Flags().Lookup(flags.FlagTrustNode))

	return cmd
}

func RegisterTxRoutes(cliCtx context.CLIContext, r *mux.Router) {
	r.HandleFunc("/txs/{hash}", QueryTxRequestHandlerFn(cliCtx)).Methods("GET")
	r.HandleFunc("/txs", QueryTxsRequestHandlerFn(cliCtx)).Methods("GET")
	r.HandleFunc("/txs", BroadcastTxRequest(cliCtx)).Methods("POST")
	r.HandleFunc("/sign_data/{msg}", SignDataRequest(cliCtx)).Methods("GET")
}

func QueryTxRequestHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		hashHexStr := vars["hash"]

		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		output, err := QueryTx(cliCtx, hashHexStr)
		if err != nil {
			if strings.Contains(err.Error(), hashHexStr) {
				rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
				return
			}
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		if output.Empty() {
			rest.WriteErrorResponse(w, http.StatusNotFound, fmt.Sprintf("no transaction found with hash %s", hashHexStr))
		}

		data, err := json.Marshal(output)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, fmt.Sprintf("parse error,%s", err.Error()))
		}

		_, _ = w.Write(data)
	}
}

func QueryTxsRequestHandlerFn(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest,
				sdk.AppendMsgToErr("could not parse query parameters", err.Error()))
			return
		}

		// if the height query param is set to zero, query for genesis transactions
		heightStr := r.FormValue("height")
		if heightStr != "" {
			if height, err := strconv.ParseInt(heightStr, 10, 64); err == nil && height == 0 {
				genutilrest.QueryGenesisTxs(cliCtx, w)
				return
			}
		}

		var (
			events      []string
			txs         []sdk.TxResponse
			page, limit int
		)

		cliCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, cliCtx, r)
		if !ok {
			return
		}

		if len(r.Form) == 0 {
			rest.PostProcessResponseBare(w, cliCtx, txs)
			return
		}

		events, page, limit, err = rest.ParseHTTPArgs(r)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		searchResult, err := utils.QueryTxsByEvents(cliCtx, events, page, limit)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		rest.PostProcessResponseBare(w, cliCtx, searchResult)
	}
}

func QueryTx(cliCtx context.CLIContext, hashHexStr string) (sdk.TxResponse, error) {
	hash, err := hex.DecodeString(hashHexStr)
	if err != nil {
		return sdk.TxResponse{}, err
	}

	node, err := cliCtx.GetNode()
	if err != nil {
		return sdk.TxResponse{}, err
	}

	resTx, err := node.Tx(hash, !cliCtx.TrustNode)
	if err != nil {
		return sdk.TxResponse{}, err
	}

	if !cliCtx.TrustNode {
		if err = ValidateTxResult(cliCtx, resTx); err != nil {
			return sdk.TxResponse{}, err
		}
	}

	resBlocks, err := getBlocksForTxResults(cliCtx, []*core.ResultTx{resTx})
	if err != nil {
		return sdk.TxResponse{}, err
	}

	out, err := formatTxResult(cliCtx.Codec, resTx, resBlocks[resTx.Height])
	if err != nil {
		return out, err
	}

	return out, nil
}

type BroadcastReq struct {
	Tx   string `json:"tx" yaml:"tx"`
	Mode string `json:"mode" yaml:"mode"`
}

func BroadcastTxRequest(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req BroadcastReq

		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		err = cliCtx.Codec.UnmarshalJSON(body, &req)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// The only line in this function different from that in Cosmos SDK
		// is the one below. Instead of codec (JSON) marshalling, hex is used
		// so that the DefaultTxDecoder can successfully recognize the IxoTx
		//
		// txBytes, err := cliCtx.Codec.MarshalBinaryLengthPrefixed(req.Tx)

		txBytes, err := hex.DecodeString(strings.TrimPrefix(req.Tx, "0x"))
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		cliCtx = cliCtx.WithBroadcastMode(req.Mode)

		res, err := cliCtx.BroadcastTx(txBytes)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		rest.PostProcessResponseBare(w, cliCtx, res)
	}
}

type SignDataResponse struct {
	SignBytes string      `json:"sign_bytes" yaml:"sign_bytes"`
	Fee       auth.StdFee `json:"fee" yaml:"fee"`
}

func SignDataRequest(cliCtx context.CLIContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		vars := mux.Vars(r)
		msgParam := vars["msg"]

		msgBytes, err := hex.DecodeString(strings.TrimPrefix(msgParam, "0x"))
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		var msg sdk.Msg
		err = cliCtx.Codec.UnmarshalJSON(msgBytes, &msg)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
			return
		}

		// all messages must be of type ixo.IxoMsg
		ixoMsg, ok := msg.(ixo.IxoMsg)
		if !ok {
			rest.WriteErrorResponse(w, http.StatusBadRequest, sdk.ErrInternal("msg must be ixo.IxoMsg").Error())
			return
		}
		msgs := []sdk.Msg{ixoMsg}

		// obtain stdSignMsg (create-project is a special case)
		var stdSignMsg auth.StdSignMsg
		switch ixoMsg.Type() {
		case project.TypeMsgCreateProject:
			stdSignMsg = project.GetProjectCreationStdSignMsg(msgs)
		default:
			// Deduce and set signer address
			signerAddress := ixo.DidToAddr(ixoMsg.GetSignerDid())
			cliCtx = cliCtx.WithFromAddress(signerAddress)

			txBldr, err := utils.PrepareTxBuilder(auth.NewTxBuilderFromCLI(), cliCtx)
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
				return
			}

			// Build the transaction
			stdSignMsg, err = txBldr.BuildSignMsg(msgs)
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
				return
			}

			// Create dummy tx with blank signature for fee approximation
			signature := ixo.IxoSignature{}
			signature.Created = signature.Created.Add(1) // maximizes signature length
			tx := ixo.NewIxoTxSingleMsg(
				stdSignMsg.Msgs[0], stdSignMsg.Fee, signature, stdSignMsg.Memo)

			// Approximate fee
			fee, err := ixo.ApproximateFeeForTx(cliCtx, tx, txBldr.ChainID())
			if err != nil {
				rest.WriteErrorResponse(w, http.StatusBadRequest, err.Error())
				return
			}
			stdSignMsg.Fee = fee
		}

		// Produce response from sign bytes and fees
		output := SignDataResponse{
			SignBytes: string(stdSignMsg.Bytes()),
			Fee:       stdSignMsg.Fee,
		}

		rest.PostProcessResponseBare(w, cliCtx, output)
	}
}

func getBlocksForTxResults(cliCtx context.CLIContext, resTxs []*core.ResultTx) (map[int64]*core.ResultBlock, error) {
	node, err := cliCtx.GetNode()
	if err != nil {
		return nil, err
	}

	resBlocks := make(map[int64]*core.ResultBlock)
	for _, resTx := range resTxs {
		if _, ok := resBlocks[resTx.Height]; !ok {
			resBlock, err := node.Block(&resTx.Height)
			if err != nil {
				return nil, err
			}

			resBlocks[resTx.Height] = resBlock
		}
	}

	return resBlocks, nil
}

func formatTxResult(cdc *codec.Codec, resTx *core.ResultTx, resBlock *core.ResultBlock) (sdk.TxResponse, error) {
	tx, err := parseTx(cdc, resTx.Tx)
	if err != nil {
		return sdk.TxResponse{}, err
	}

	return sdk.NewResponseResultTx(resTx, tx, resBlock.Block.Time.Format(time.RFC3339)), nil
}

func parseTx(cdc *codec.Codec, txBytes []byte) (sdk.Tx, error) {
	return ixo.DefaultTxDecoder(cdc)(txBytes)
}

func ValidateTxResult(cliCtx context.CLIContext, resTx *core.ResultTx) error {
	if !cliCtx.TrustNode {
		check, err := cliCtx.Verify(resTx.Height)
		if err != nil {
			return err
		}

		err = resTx.Proof.Validate(check.Header.DataHash)
		if err != nil {
			return err
		}
	}

	return nil
}
