package types

const (
	EventTypeAddDidDoc     = "add_did_doc"
	EventTypeAddCredential = "add_credential"

	AttributeKeyDid          = "did"
	AttributeKeyPubKey       = "pub_key"
	AttributeKeyCredType     = "cred_type"
	AttributeKeyIssuer       = "issuer"
	AttributeKeyIssued       = "issued"
	AttributeKeyClaimID      = "claim"
	AttributeKeyKYCValidated = "kyc_validated"
	AttributeValueCategory   = ModuleName
)
