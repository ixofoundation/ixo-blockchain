package rest

import (
	"encoding/json"
	"net/http"

	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/x/auth"
)

func BroadcastRest(cliCtx context.CLIContext, cdc *codec.Codec, stdTx auth.StdTx, mode string) ([]byte, sdk.Error) {

	txBytes, err := cdc.MarshalBinaryLengthPrefixed(stdTx)
	if err != nil {
		return nil, sdk.NewError(DefaultCodeSpace, http.StatusInternalServerError, err.Error())
	}
	cliCtx = cliCtx.WithBroadcastMode(mode)

	res, err := cliCtx.BroadcastTx(txBytes)
	if err != nil {
		var abci []sdk.ABCIMessageLog
		_err := json.Unmarshal([]byte(err.Error()), &abci)
		if _err != nil {
			panic(_err)
		}

		var _log Log
		_err = json.Unmarshal([]byte(abci[0].Log), &_log)
		if _err != nil {
			panic(_err)
		}

		return nil, sdk.NewError(DefaultCodeSpace, http.StatusInternalServerError, _log.Message)
	}

	return PostProcessResponse(cliCtx, res)
}
