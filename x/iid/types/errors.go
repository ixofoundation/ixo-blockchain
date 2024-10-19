package types

import (
	errorsmod "cosmossdk.io/errors"
)

// x/did module sentinel errors
var (
	ErrDidDocumentNotFound        = errorsmod.Register(ModuleName, 1100, "did document not found")
	ErrDidDocumentFound           = errorsmod.Register(ModuleName, 1101, "did document found")
	ErrInvalidDIDFormat           = errorsmod.Register(ModuleName, 1102, "input not compliant with the DID specifications (crf. https://www.w3.org/TR/did-core/#did-syntax)")
	ErrInvalidDIDURLFormat        = errorsmod.Register(ModuleName, 1103, "input not compliant with the DID URL specifications (crf. https://www.w3.org/TR/did-core/#did-url-syntax)")
	ErrInvalidRFC3986UriFormat    = errorsmod.Register(ModuleName, 1104, "input not compliant with the RFC3986 URI specifications (crf. https://datatracker.ietf.org/doc/html/rfc3986)")
	ErrEmptyRelationships         = errorsmod.Register(ModuleName, 1105, "a verification method should have at least one verification relationship. (cfr. https://www.w3.org/TR/did-core/#verification-relationships)")
	ErrUnauthorized               = errorsmod.Register(ModuleName, 1106, "the requester DID's verification method is not listed in the required relationship")
	ErrInvalidState               = errorsmod.Register(ModuleName, 1107, "the requested action is not applicable on the resource")
	ErrInvalidInput               = errorsmod.Register(ModuleName, 1108, "input is invalid")
	ErrVerificationMethodNotFound = errorsmod.Register(ModuleName, 1109, "verification method not found")
	ErrInvalidDidMethodFormat     = errorsmod.Register(ModuleName, 1110, "invalid did method format")
	ErrKeyFormatNotSupported      = errorsmod.Register(ModuleName, 1111, "key format not supported")
	ErrDidPubKeyMismatch          = errorsmod.Register(ModuleName, 1112, "did does not match public key in signature")
)
