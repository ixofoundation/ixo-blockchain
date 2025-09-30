package types

import (
	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
	iidtypes "github.com/ixofoundation/ixo-blockchain/v6/x/iid/types"
	"github.com/ixofoundation/ixo-blockchain/v6/x/token/types/contracts/ixo1155"
)

// IsValidToken tells if a Token is valid,
func IsValidToken(token *Token) bool {
	if token == nil {
		return false
	}
	if iidtypes.IsEmpty(token.Name) {
		return false
	}
	_, err := sdk.AccAddressFromBech32(token.ContractAddress)
	if err != nil {
		return false
	}
	_, err = sdk.AccAddressFromBech32(token.Minter)
	if err != nil {
		return false
	}
	if !iidtypes.IsValidDID(token.Class) {
		return false
	}
	return true
}

// IsValidTokenProperties tells if a TokenProperties is valid,
func IsValidTokenProperties(tokenProperties *TokenProperties) bool {
	if tokenProperties == nil {
		return false
	}
	if iidtypes.IsEmpty(tokenProperties.Id) {
		return false
	}
	if iidtypes.IsEmpty(tokenProperties.Name) {
		return false
	}
	return true
}

type MintBatchData struct {
	Id         string
	Uri        string
	Name       string
	Index      string
	Amount     math.Uint
	Collection string
	TokenData  []*TokenData
}

func (batch *MintBatchData) GetWasmMintBatch() ixo1155.Batch {
	return []string{batch.Id, batch.Amount.String(), batch.Uri}
}

func (batch *MintBatchData) GetTokenMintedEventBatch() *TokenBatch {
	return &TokenBatch{
		Id:     batch.Id,
		Amount: batch.Amount,
	}
}

func (batch *TokenBatch) GetWasmTransferBatch() ixo1155.Batch {
	return []string{batch.Id, batch.Amount.String(), ""}
}
func Map[T, V any](ts []T, fn func(T) V) []V {
	result := make([]V, len(ts))
	for i, t := range ts {
		result[i] = fn(t)
	}
	return result
}
