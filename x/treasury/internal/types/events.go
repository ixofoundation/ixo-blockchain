package types

const (
	EventTypeSend           = "send"
	EventTypeOracleBurn     = "oracle_burn"
	EventTypeOracleMint     = "oracle_mint"
	EventTypeOracleTransfer = "oracle_transfer"

	AttributeKeyTxHashFromDid = "from_did"
	AttributeKeyToDidOrAddr   = "to_did_or_addr"
	AttributeKeyAmount        = "amount"
	AttributeKeyOracleDid     = "oracle_did"
	AttributeKeyProof         = "proof"

	AttributeValueCategory = ModuleName
)
