package types

const (
	EventTypePaymentContractAuthorisation = "payment_contract_authorisation"
	EventTypeCreatePaymentTemplate        = "create_payment_template"
	EventTypeCreatePaymentContract        = "create_payment_contract"
	EventTypeCreateSubscription           = "create_subscription"
	EventTypeGrantDiscount                = "grant_discount"
	EventTypeRevokeDiscount               = "revoke_discount"
	EventTypeEffectPayment                = "effect_payment"

	AttributeKeyPayerDid           = "payer_did"
	AttributeKeyPaymentContractId  = "payment_contract-id"
	AttributeKeyAuthorised         = "authorised"
	AttributeKeyCreatorDid         = "creator_did"
	AttributeKeyPaymentTemplateId  = "payment_template_id"
	AttributeKeyPayer              = "payer"
	AttributeKeyCanDeauthorise     = "can_deauthorise"
	AttributeKeyDiscountId         = "discount_id"
	AttributeKeySubscriptionId     = "attribute_key"
	AttributeKeyMaxPeriods         = "max_periods"
	AttributeKeyPeriod             = "period"
	AttributeKeySenderDid          = "sender_did"
	AttributeKeyRecipient          = "recipient"
	AttributeKeyHolder             = "holder"
	AttributeKeyAttributeKeyId     = "payment_id"
	AttributeKeyPaymentAmount      = "payment_amount"
	AttributeKeyPaymentMinimum     = "payment_minimum"
	AttributeKeyPaymentMaximum     = "payment_maximum"
	AttributeKeyDiscounts          = "discounts"
	AttributeKeyWalletDistribution = "wallet_distribution"

	AttributeKeyInputFromPayRemainderPool = "input_from_pay_remainder_pool"
	AttributeKeyInputFromPayer            = "input_from_payer"
	AttributeKeyOutputToPayRemainderPool  = "output_to_pay_remainder_pool"
	AttributeKeyOutputToPayees            = "output_to_payees"
	AttributeValueCategory                = ModuleName
)
