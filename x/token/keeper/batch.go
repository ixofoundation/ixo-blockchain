package keeper

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"

	iidtypes "github.com/ixofoundation/ixo-blockchain/v8/x/iid/types"
	"github.com/ixofoundation/ixo-blockchain/v8/x/token/types"
)

// ValidateTokenBatch enforces, inside the keeper (not just in ValidateBasic),
// the invariants for a token batch used by Transfer/Retire/TransferCredit/Cancel:
//   - the batch is non-empty,
//   - every element has a non-empty id and a positive amount,
//   - every element resolves to the SAME contract (the documented
//     "all tokens in the same contract" invariant).
//
// These checks live in the keeper because the handlers index msg.Tokens[0] and
// build a single batch executed against that one contract; relying on
// ValidateBasic alone is unsafe for any route that reaches the msg server
// without it and lets a mixed-contract batch corrupt supply accounting. The
// resolved common contract address is returned for the caller's convenience.
func (k Keeper) ValidateTokenBatch(ctx sdk.Context, tokens []*types.TokenBatch) (string, error) {
	if len(tokens) == 0 {
		return "", errorsmod.Wrap(types.ErrTokenBatchInvalid, "token batch cannot be empty")
	}

	var contractAddress string
	for i, batch := range tokens {
		if batch == nil || iidtypes.IsEmpty(batch.Id) {
			return "", errorsmod.Wrap(types.ErrTokenBatchInvalid, "token id is empty for a batch element")
		}
		if batch.Amount.IsZero() {
			return "", errorsmod.Wrap(types.ErrTokenAmountIncorrect, "token amount must be greater than 0")
		}

		_, token, err := k.GetTokenById(ctx, batch.Id)
		if err != nil {
			return "", err
		}

		if i == 0 {
			contractAddress = token.ContractAddress
		} else if token.ContractAddress != contractAddress {
			return "", errorsmod.Wrapf(types.ErrTokenContractMismatch,
				"token %s belongs to contract %s, expected %s", batch.Id, token.ContractAddress, contractAddress)
		}
	}

	return contractAddress, nil
}
