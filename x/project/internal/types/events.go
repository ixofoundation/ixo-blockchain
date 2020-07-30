package types

const (
	EventTypeCreateProject       = "create_project"
	EventTypeUpdateProjectStatus = "update_project_status"

	AttributeKeyTxHash       = "txHash"
	AttributeKeySenderDid    = "sender_did"
	AttributeKeyProjectDid   = "project_did"
	AttributeKeyPubKey       = "pub_key"
	AttributeKeyData         = "data"
	AttributeKeyRecipientDid = "recipient_did"
	AttributeKeyAmount       = "amount"

	AttributeValueCategory = ModuleName
)
