package types

import (
	"crypto/sha256"
	"encoding/hex"
	"strings"
	time "time"

	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
)

// IsEmpty tells if the trimmed input is empty
func IsEmpty(input string) bool {
	return strings.TrimSpace(input) == ""
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

func (e Entity) GetAdminAccount() (*EntityAccount, error) {
	for _, acc := range e.Accounts {
		if acc.Name == EntityAdminAccountName {
			return acc, nil
		}
	}
	return nil, ErrAccountNotFound
}

func (e Entity) ContainsAccountName(name string) bool {
	for _, acc := range e.Accounts {
		if acc.Name == name {
			return true
		}
	}
	return false
}

func (e Entity) ContainsAccountAddress(address string) bool {
	for _, acc := range e.Accounts {
		if acc.Address == address {
			return true
		}
	}
	return false
}
