package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
	iidtypes "github.com/ixofoundation/ixo-blockchain/x/iid/types"
	"github.com/ixofoundation/ixo-blockchain/x/token/types/contracts/ixo1155"
)

var (
	KeyIxo1155ContractCode = []byte("Ixo1155ContractCode")
)

func validateContractCode(i interface{}) error {
	_, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T expected uint64", i)
	}

	return nil
}

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

// ParamTable for module.
func ParamKeyTable() paramstypes.KeyTable {
	return paramstypes.NewKeyTable().RegisterParamSet(&Params{})
}

func NewParams(ixo1155ContractCode uint64) Params {
	return Params{
		Ixo1155ContractCode: ixo1155ContractCode,
	}
}

func DefaultParams() Params {
	return Params{
		Ixo1155ContractCode: 0,
	}
}

// Implements params.ParamSet
func (p *Params) ParamSetPairs() paramstypes.ParamSetPairs {
	return paramstypes.ParamSetPairs{
		paramstypes.ParamSetPair{Key: KeyIxo1155ContractCode, Value: &p.Ixo1155ContractCode, ValidatorFn: validateContractCode},
	}
}

type MintBatchData struct {
	Id         string
	Uri        string
	Name       string
	Index      string
	Amount     sdk.Uint
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
