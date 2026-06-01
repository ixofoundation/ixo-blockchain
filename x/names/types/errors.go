package types

import errorsmod "cosmossdk.io/errors"

var (
	ErrNamespaceNotFound       = errorsmod.Register(ModuleName, 1100, "namespace not found")
	ErrNamespaceExists         = errorsmod.Register(ModuleName, 1101, "namespace already exists")
	ErrInvalidNamespace        = errorsmod.Register(ModuleName, 1102, "invalid namespace configuration")
	ErrNameNotFound            = errorsmod.Register(ModuleName, 1103, "name not found")
	ErrNameTaken               = errorsmod.Register(ModuleName, 1104, "name already taken")
	ErrInvalidName             = errorsmod.Register(ModuleName, 1105, "invalid name")
	ErrSelfRegisterNotAllowed  = errorsmod.Register(ModuleName, 1106, "namespace does not allow self-registration")
	ErrUnauthorized            = errorsmod.Register(ModuleName, 1107, "signer not authorized for this action")
	ErrNotRegistrar            = errorsmod.Register(ModuleName, 1108, "signer is not a registrar of this namespace")
	ErrExpiryNotAllowed        = errorsmod.Register(ModuleName, 1109, "namespace does not allow non-zero valid_until")
	ErrInvalidStatusTransition = errorsmod.Register(ModuleName, 1110, "invalid status transition")
	ErrInvalidDID              = errorsmod.Register(ModuleName, 1111, "invalid DID")
	ErrInvalidAuthority        = errorsmod.Register(ModuleName, 1112, "invalid authority")
	ErrInvalidRequest          = errorsmod.Register(ModuleName, 1113, "invalid request")
)
