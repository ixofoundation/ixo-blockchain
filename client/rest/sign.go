package rest

import (
	"bytes"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/cosmos/cosmos-sdk/x/auth"
)

func SignStdTxFromRest(txBldr auth.TxBuilder, cliCtx context.CLIContext, name string, stdTx auth.StdTx, appendSig bool, offline bool, password string) (auth.StdTx, error) {

	var signedStdTx auth.StdTx

	info, err := txBldr.Keybase().Get(name)
	if err != nil {
		return signedStdTx, err
	}

	addr := info.GetPubKey().Address()

	if !isTxSigner(sdk.AccAddress(addr), stdTx.GetSigners()) {
		return signedStdTx, fmt.Errorf("%s: %s", "Error invalid signer", name)
	}

	ad := sdk.AccAddress(addr)
	if !offline {
		txBldr, err = populateAccountFromState(txBldr, cliCtx, ad)
		if err != nil {
			return signedStdTx, err
		}
	}

	return txBldr.SignStdTx(name, password, stdTx, appendSig)
}

func isTxSigner(user sdk.AccAddress, signers []sdk.AccAddress) bool {
	for _, s := range signers {
		if bytes.Equal(user.Bytes(), s.Bytes()) {
			return true
		}
	}

	return false
}

func populateAccountFromState(
	txBldr auth.TxBuilder, cliCtx context.CLIContext, addr sdk.AccAddress,
) (auth.TxBuilder, error) {
	num, seq, err := auth.NewAccountRetriever(cliCtx).GetAccountNumberSequence(addr)
	if err != nil {
		return txBldr, err
	}

	return txBldr.WithAccountNumber(num).WithSequence(seq + txBldr.Sequence()), nil
}
