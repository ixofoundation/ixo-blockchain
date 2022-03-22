package types

const (
	EventTypeCreateProject       = "create_project"
	EventTypeUpdateProjectStatus = "update_project_status"
	EventTypeUpdateProjectDoc    = "update_project_doc"
	EventTypeCreateAgent         = "create_agent"
	EventTypeUpdateAgent         = "update_agent"
	EventTypeCreateClaim         = "create_claim"
	EventTypeCreateEvaluation    = "create_evaluation"
	EventTypeWithdrawFunds       = "withdraw_funds"

	AttributeKeyTxHash          = "tx_hash"
	AttributeKeySenderDid       = "sender_did"
	AttributeKeyProjectDid      = "project_did"
	AttributeKeyPubKey          = "pub_key"
	AttributeKeyRecipientDid    = "recipient_did"
	AttributeKeyAmount          = "amount"
	AttributeKeyIsRefund        = "is_refund"
	AttributeKeyClaimID         = "claim_id"
	AttributeKeyClaimTemplateID = "claim_template_id"
	AttributeKeyClaimStatus     = "claim_status"
	AttributeKeyAgentDid        = "agent_did"
	AttributeKeyAgentRole       = "role"
	AttributeKeyEthFundingTxnID = "eth_funding_txn_id"
	AttributeKeyUpdatedStatus   = "updated_status"

	AttributeValueCategory = ModuleName
)
