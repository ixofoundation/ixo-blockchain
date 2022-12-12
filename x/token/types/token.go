package types

import (
	fmt "fmt"
	"strconv"

	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var (
	KeyCw20ContractCode    = []byte("Cw20ContractCode")
	KeyCw721ContractCode   = []byte("Cw721ContractCode")
	KeyIxo1155ContractCode = []byte("Ixo1155ContractCode")
)

func parseCode(stringCode string) (uint64, error) {
	code, err := strconv.ParseUint(stringCode, 0, 64)
	if err != nil {
		return 0, err
	}

	return code, nil
}

func validateContractCode(i interface{}) error {
	codeString, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T expected string", i)
	}
	_, err := parseCode(codeString)
	if err != nil {
		return err
	}

	return nil
}

// ParamTable for project module.
func ParamKeyTable() paramstypes.KeyTable {
	return paramstypes.NewKeyTable().RegisterParamSet(&Params{})
}

func NewParams(nftContractAddress string, nftContractMinter string) Params {
	return Params{
		Cw20ContractCode:    "0",
		Cw721ContractCode:   "0",
		Ixo1155ContractCode: "0",
	}
}

// func (p Params) MustCw20ContractCode() (uint64, error) {

// }

// // default project module parameters
func DefaultParams() Params {
	return Params{
		Cw20ContractCode:    "0",
		Cw721ContractCode:   "0",
		Ixo1155ContractCode: "0",
	}
}

func (p *Params) GetCw20ContractCode() uint64 {
	code, err := parseCode(p.Cw20ContractCode)
	if err != nil {
		panic(err)
	}
	return code
}
func (p *Params) GetCw721ContractCode() uint64 {
	code, err := parseCode(p.Cw721ContractCode)
	if err != nil {
		panic(err)
	}
	return code
}
func (p *Params) GetIxo1155ContractCode() uint64 {
	code, err := parseCode(p.Ixo1155ContractCode)
	if err != nil {
		panic(err)
	}
	return code
}

// // Implements params.ParamSet
func (p *Params) ParamSetPairs() paramstypes.ParamSetPairs {
	return paramstypes.ParamSetPairs{
		{KeyCw20ContractCode, &p.Cw20ContractCode, validateContractCode},
		{KeyCw721ContractCode, &p.Cw721ContractCode, validateContractCode},
		{KeyIxo1155ContractCode, &p.Ixo1155ContractCode, validateContractCode},
	}
}
