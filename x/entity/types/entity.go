package types

import (
	"crypto/sha256"
	"encoding/hex"
	fmt "fmt"
	"strings"
	time "time"

	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	paramstypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

var (
	KeyNftContractAddress = []byte("NftContractAddress")
	KeyNftContractMinter  = []byte("NftContractMinter")
	KeyCreateSequence     = []byte("CreateSequence")
)

// IsEmpty tells if the trimmed input is empty
func IsEmpty(input string) bool {
	return strings.TrimSpace(input) == ""
}

func validateNftContractAddress(i interface{}) error {
	addr, ok := i.(string)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T expected string", i)
	}

	if len(addr) == 0 {
		return fmt.Errorf("nft contract addresses can not be empty cannot be empty")
	}

	_, err := sdk.AccAddressFromBech32(addr)
	if err != nil {
		return err
	}

	return nil
}

func validateCreateSequence(i interface{}) error {
	_, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T expected uint64", i)
	}

	return nil
}

// IsValidEntity tells if a Entity is valid,
// that is if it has a non empty versionId and a non-zero create date
func IsValidEntity(entity *Entity) bool {
	if entity == nil {
		return false
	}
	if IsEmpty(entity.Metadata.VersionId) {
		return false
	}
	if entity.Metadata.Created == nil || entity.Metadata.Created.IsZero() {
		return false
	}
	return true
}

// ParamTable for project module.
func ParamKeyTable() paramstypes.KeyTable {
	return paramstypes.NewKeyTable().RegisterParamSet(&Params{})
}

func NewParams(nftContractAddress string, nftContractMinter string, createSequence uint64) Params {
	return Params{
		NftContractAddress: nftContractAddress,
		NftContractMinter:  nftContractAddress,
		CreateSequence:     createSequence,
	}
}

func DefaultParams() Params {
	return Params{
		NftContractAddress: "ixo14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9sqa3vn7",
		NftContractMinter:  "ixo14hj2tavq8fpesdwxxcu44rty3hh90vhujrvcmstl4zr3txmfvw9sqa3vn7",
		CreateSequence:     0,
	}
}

// Implements params.ParamSet
func (p *Params) ParamSetPairs() paramstypes.ParamSetPairs {
	return paramstypes.ParamSetPairs{
		paramstypes.ParamSetPair{Key: KeyNftContractAddress, Value: &p.NftContractAddress, ValidatorFn: validateNftContractAddress},
		paramstypes.ParamSetPair{Key: KeyNftContractMinter, Value: &p.NftContractMinter, ValidatorFn: validateNftContractAddress},
		paramstypes.ParamSetPair{Key: KeyCreateSequence, Value: &p.CreateSequence, ValidatorFn: validateCreateSequence},
	}
}

func NewEntityMetadata(versionData []byte, created time.Time) EntityMetadata {
	m := EntityMetadata{
		Created: &created,
	}
	UpdateEntityMetadata(&m, versionData, created)
	return m
}

// UpdateEntityMetadata updates a entity metadata time and version id
func UpdateEntityMetadata(meta *EntityMetadata, versionData []byte, updated time.Time) {
	txH := sha256.Sum256(versionData)
	meta.VersionId = hex.EncodeToString(txH[:])
	meta.Updated = &updated
}

// Helper to get module account key in form of id#name
func GetModuleAccountKey(id, name string) string {
	return id + "#" + name
}

// Helper to get module account address
func GetModuleAccountAddress(id, name string) sdk.AccAddress {
	return authtypes.NewModuleAddress(GetModuleAccountKey(id, name))
}

// UnpackInterfaces implements UnpackInterfacesMessage.UnpackInterfaces
// Require to run this to add the Any for Authorization to cache
func (g Grant) UnpackInterfaces(unpacker cdctypes.AnyUnpacker) error {
	var a authz.Authorization
	return unpacker.UnpackAny(g.Authorization, &a)
}
