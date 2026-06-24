package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	iidtypes "github.com/ixofoundation/ixo-blockchain/v8/x/iid/types"
)

// validateAddr returns an error if addr is not a valid bech32 address.
func validateAddr(addr string) error {
	if _, err := sdk.AccAddressFromBech32(addr); err != nil {
		return errorsmod.Wrapf(ErrInvalidRequest, "invalid address %q: %s", addr, err)
	}
	return nil
}

// validateDID returns an error if did is not a valid DID.
func validateDID(did string) error {
	if !iidtypes.IsValidDID(did) {
		return errorsmod.Wrapf(ErrInvalidDID, "invalid DID %q", did)
	}
	return nil
}

// ValidateBasic for MsgCreateNamespace.
func (m *MsgCreateNamespace) ValidateBasic() error {
	if err := validateAddr(m.Authority); err != nil {
		return errorsmod.Wrap(ErrInvalidAuthority, err.Error())
	}
	if m.Namespace == nil {
		return errorsmod.Wrap(ErrInvalidRequest, "namespace is required")
	}
	if err := ValidateNamespace(*m.Namespace); err != nil {
		return err
	}
	for _, r := range m.Namespace.RegistrarAccounts {
		if err := validateAddr(r); err != nil {
			return err
		}
	}
	return nil
}

// ValidateBasic for MsgUpdateNamespace.
func (m *MsgUpdateNamespace) ValidateBasic() error {
	if err := validateAddr(m.Authority); err != nil {
		return errorsmod.Wrap(ErrInvalidAuthority, err.Error())
	}
	if m.Namespace == nil {
		return errorsmod.Wrap(ErrInvalidRequest, "namespace is required")
	}
	if err := ValidateNamespace(*m.Namespace); err != nil {
		return err
	}
	for _, r := range m.Namespace.RegistrarAccounts {
		if err := validateAddr(r); err != nil {
			return err
		}
	}
	return nil
}

// ValidateBasic for MsgRegisterName.
func (m *MsgRegisterName) ValidateBasic() error {
	if err := validateAddr(m.Signer); err != nil {
		return err
	}
	if m.Namespace == "" {
		return errorsmod.Wrap(ErrInvalidRequest, "namespace is required")
	}
	if m.Name == "" {
		return errorsmod.Wrap(ErrInvalidName, "name is required")
	}
	return validateDID(m.OwnerDid)
}

// ValidateBasic for MsgRegisterNameByRegistrar.
func (m *MsgRegisterNameByRegistrar) ValidateBasic() error {
	if err := validateAddr(m.Registrar); err != nil {
		return err
	}
	if m.Namespace == "" {
		return errorsmod.Wrap(ErrInvalidRequest, "namespace is required")
	}
	if m.Name == "" {
		return errorsmod.Wrap(ErrInvalidName, "name is required")
	}
	if err := ValidateRecordMetadata(m.EvidenceHash, m.Source); err != nil {
		return err
	}
	return validateDID(m.OwnerDid)
}

// ValidateBasic for MsgUpdateNameByRegistrar.
func (m *MsgUpdateNameByRegistrar) ValidateBasic() error {
	if err := validateAddr(m.Registrar); err != nil {
		return err
	}
	if m.Namespace == "" {
		return errorsmod.Wrap(ErrInvalidRequest, "namespace is required")
	}
	if m.NormalizedName == "" {
		return errorsmod.Wrap(ErrInvalidName, "normalized_name is required")
	}
	return ValidateRecordMetadata(m.EvidenceHash, m.Source)
}

// ValidateBasic for MsgTransferName.
func (m *MsgTransferName) ValidateBasic() error {
	if err := validateAddr(m.Signer); err != nil {
		return err
	}
	if m.Namespace == "" {
		return errorsmod.Wrap(ErrInvalidRequest, "namespace is required")
	}
	if m.NormalizedName == "" {
		return errorsmod.Wrap(ErrInvalidName, "normalized_name is required")
	}
	return validateDID(m.NewOwnerDid)
}

// ValidateBasic for MsgSetNameStatus.
func (m *MsgSetNameStatus) ValidateBasic() error {
	if err := validateAddr(m.Signer); err != nil {
		return err
	}
	if m.Namespace == "" {
		return errorsmod.Wrap(ErrInvalidRequest, "namespace is required")
	}
	if m.NormalizedName == "" {
		return errorsmod.Wrap(ErrInvalidName, "normalized_name is required")
	}
	switch m.Status {
	case NAME_STATUS_ACTIVE,
		NAME_STATUS_SUSPENDED,
		NAME_STATUS_REVOKED,
		NAME_STATUS_TOMBSTONED:
		return nil
	default:
		return errorsmod.Wrapf(ErrInvalidStatusTransition, "unsupported target status %s", m.Status)
	}
}
