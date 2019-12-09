package rest

import (
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/keys"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"

	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
)

func SignAndBroadcast(br rest.BaseReq, cliCtx context.CLIContext,
	mode, password string, msgs []sdk.Msg) ([]byte, sdk.Error) {

	cdc := cliCtx.Codec
	gasAdj, _, err := ParseFloat64OrReturnBadRequest(br.GasAdjustment, client.DefaultGasAdjustment)
	if err != nil {
		return nil, sdk.NewError(DefaultCodeSpace, http.StatusBadRequest, err.Error())
	}

	simAndExec, gas, err := client.ParseGas(br.Gas)
	if err != nil {
		return nil, sdk.NewError(DefaultCodeSpace, http.StatusBadRequest, err.Error())
	}

	keyBase, err := keys.NewKeyBaseFromHomeFlag()
	if err != nil {
		return nil, sdk.NewError(DefaultCodeSpace, http.StatusBadRequest, err.Error())
	}

	txBldr := auth.NewTxBuilder(
		utils.GetTxEncoder(cdc), br.AccountNumber, br.Sequence, gas, gasAdj,
		br.Simulate, br.ChainID, br.Memo, br.Fees, br.GasPrices,
	)
	txBldr = txBldr.WithKeybase(keyBase)

	if br.Simulate || simAndExec {
		if gasAdj < 0 {
			return nil, sdk.NewError(DefaultCodeSpace, http.StatusBadRequest, "Error invalid gas adjustment")
		}

		txBldr, err = utils.EnrichWithGas(txBldr, cliCtx, msgs)
		if err != nil {
			return nil, sdk.NewError(DefaultCodeSpace, http.StatusBadRequest, err.Error())
		}

		if br.Simulate {
			return SimulationResponse(cdc, txBldr.Gas())
		}
	}

	stdMsg, err := txBldr.BuildSignMsg(msgs)
	if err != nil {
		return nil, sdk.NewError(DefaultCodeSpace, http.StatusBadRequest, err.Error())
	}

	stdTx := auth.NewStdTx(stdMsg.Msgs, stdMsg.Fee, nil, stdMsg.Memo)

	stdTx, err = SignStdTxFromRest(txBldr, cliCtx, cliCtx.GetFromName(), stdTx, true, false, password)
	if err != nil {
		return nil, sdk.NewError(DefaultCodeSpace, http.StatusBadRequest, err.Error())
	}

	return BroadcastRest(cliCtx, cdc, stdTx, mode)

}

func SignAndBroadcastMultiple(brs []rest.BaseReq, cliCtxs []context.CLIContext,
	mode []string, passwords []string, msgs []sdk.Msg) ([]byte, error) {

	var stdTxs types.StdTx
	for i, _ := range brs {

		cdc := cliCtxs[i].Codec
		gasAdj, _, err := ParseFloat64OrReturnBadRequest(brs[i].GasAdjustment, client.DefaultGasAdjustment)
		if err != nil {
			return nil, sdk.NewError(DefaultCodeSpace, http.StatusBadRequest, err.Error())
		}

		simAndExec, gas, err := client.ParseGas(brs[i].Gas)
		if err != nil {
			return nil, sdk.NewError(DefaultCodeSpace, http.StatusBadRequest, err.Error())
		}

		keyBase, err := keys.NewKeyBaseFromHomeFlag()
		if err != nil {
			return nil, sdk.NewError(DefaultCodeSpace, http.StatusBadRequest, err.Error())
		}

		txBldr := auth.NewTxBuilder(
			utils.GetTxEncoder(cdc), brs[i].AccountNumber, brs[i].Sequence, gas, gasAdj,
			brs[i].Simulate, brs[i].ChainID, brs[i].Memo, brs[i].Fees, brs[i].GasPrices,
		)

		txBldr = txBldr.WithKeybase(keyBase)

		if brs[i].Simulate || simAndExec {
			if gasAdj < 0 {
				return nil, sdk.NewError(DefaultCodeSpace, http.StatusBadRequest, "Error invalid gas adjustment")
			}

			txBldr, err = utils.EnrichWithGas(txBldr, cliCtxs[i], []sdk.Msg{msgs[i]})
			if err != nil {
				return nil, sdk.NewError(DefaultCodeSpace, http.StatusBadRequest, err.Error())
			}

			if brs[i].Simulate {
				return SimulationResponse(cdc, txBldr.Gas())
			}
		}

		var count = uint64(0)
		for j := 0; j < i; j++ {
			if txBldr.AccountNumber() == brs[j].AccountNumber {
				count++
			}
		}
		txBldr = txBldr.WithSequence(count)

		stdMsg, err := txBldr.BuildSignMsg(msgs)
		if err != nil {
			return nil, sdk.NewError(DefaultCodeSpace, http.StatusBadRequest, err.Error())
		}

		stdTx := auth.NewStdTx(stdMsg.Msgs, stdMsg.Fee, nil, stdMsg.Memo)

		stdTx, err = SignStdTxFromRest(txBldr, cliCtxs[i], cliCtxs[i].GetFromName(), stdTx, true, false, passwords[i])
		if err != nil {
			return nil, sdk.NewError(DefaultCodeSpace, http.StatusBadRequest, err.Error())
		}

		if i == 0 {
			stdTxs.Msgs = stdTx.Msgs
			stdTxs.Fee = stdTx.Fee
			stdTxs.Memo = stdTx.Memo
		}
		stdTxs.Signatures = append(stdTxs.Signatures, stdTx.Signatures...)
	}

	return BroadcastRest(cliCtxs[0], cliCtxs[0].Codec, stdTxs, mode[0])

}
