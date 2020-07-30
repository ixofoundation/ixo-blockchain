package types

const (
	EventTypePaymentContractAuthorisation = "payment_contract_authorisation"
	EventTypeCreatePaymentTemplate        = "create_payment_template"
	EventTypeCreatePaymentContract        = "create_payment_contract"
	EventTypeCreateSubscription           = "create_subscription"
	EventTypeGrantDiscount                = "grant-discount"
	EventTypeRevokeDiscount               = "revoke_discount"
	EventTypeEffectPayment                = "effect_payment"

	AttributeKeyPayerDid          = "payer_did"
	AttributeKeyPaymentContractId = "payment_contract-id"
	AttributeKeyAuthorised        = "authorised"
	AttributeKeyCreatorDid        = "creator_did"
	AttributeKeyPaymentTemplate   = "payment_template"
	AttributeKeyPaymentTemplateId = "payment_template_id"
	AttributeKeyPayer             = "payer"
	AttributeKeyDeAuthorise       = "can_de_authorise"
	AttributeKeyDiscountId        = "discountId"
	AttributeKeySubscriptionId    = "attribute_key"
	AttributeKeyMaxPeriods        = "max_periods"
	AttributeKeyPeriod            = "period"
	AttributeKeySenderDid         = "sender_did"
	AttributeKeyRecipient         = "recipient"
	AttributeKeyHolder            = "holder"
	AttributeValueCategory        = ModuleName
)
