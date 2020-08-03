package types

const (
	EventTypeCreateProject       = "create_project"
	EventTypeUpdateProjectStatus = "update_project_status"

	AttributeKeyTxHash          = "tx_hash"
	AttributeKeySenderDid       = "sender_did"
	AttributeKeyProjectDid      = "project_did"
	AttributeKeyPubKey          = "pub_key"
	AttributeKeyData            = "data"
	AttributeKeyRecipientDid    = "recipient_did"
	AttributeKeyAmount          = "amount"
	AttributeKeyIsRefund        = "is_refund"
	AttributeKeyClaimID         = "claim_iD"
	AttributeKeyStatus          = "status"
	AttributeKeyAgentDid        = "agent_did"
	AttributeKeyRole            = "role"
	AttributeKeyEthFundingTxnID = "eth_funding_txn_id"

	AttributeValueCategory = ModuleName
)
