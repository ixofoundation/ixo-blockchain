package types

import (
	fmt "fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
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
	return true
}

// IsValidTokenProperties tells if a TokenProperties is valid,
func IsValidTokenProperties(tokenProperties *TokenProperties) bool {
	if tokenProperties == nil {
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

// UpdateSupply updates token suply by subtracting or adding amount
func (token *Token) UpdateSupply(amount sdk.Int) error {
	if amount.IsNegative() {
		token.Supply = token.Supply.Sub(sdk.NewUint(amount.Abs().Uint64()))
	} else {
		token.Supply = token.Supply.Add(sdk.NewUint(amount.Abs().Uint64()))
	}
	return nil
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

func (batch *MintBatchData) GetTokenMintedEventBatch() *TokenMintedBatch {
	return &TokenMintedBatch{
		Id:     batch.Id,
		Amount: batch.Amount.String(),
	}
}

func Map[T, V any](ts []T, fn func(T) V) []V {
	result := make([]V, len(ts))
	for i, t := range ts {
		result[i] = fn(t)
	}
	return result
}
