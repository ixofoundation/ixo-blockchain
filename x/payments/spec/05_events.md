# Events

The Payment module emits the following events:

## Endblocker
| Type                            | Attribute Key            | Attribute Value       |
|---------------------------------|--------------------------|-----------------------|
| effect_payment                  | sender_did               | {sender_did}          |
| effect_payment                  | payment_contract_id      | {payment_contract_id} |


## Handler

## MsgCreatePaymentTemplate
| Type                            | Attribute Key            | Attribute Value       |
|---------------------------------|--------------------------|-----------------------|
| create_payment_template         | creator_did              | {creator_did}         |
| create_payment_template         | id                       | {id}                  |
| create_payment_template         | payment_amount           | {payment_amount}      |
| create_payment_template         | payment_minimum          | {payment_minimum}     |
| create_payment_template         | payment_maximum,         | {payment_maximum}     |
| create_payment_template         | discounts                | {discounts}           |
| create_payment_template         | wallet_distribution      | {wallet_distribution} |

## MsgCreatePaymentContract
| Type                            | Attribute Key            | Attribute Value       |
|---------------------------------|--------------------------|-----------------------|
| create_payment_contract         | creator_did              | {creator_did}         |
| create_payment_contract         | payment_template_id      | {payment_template_id} |
| create_payment_contract         | payment_contract_id      | {payment_contract_id} |
| create_payment_contract         | payer                    | {payer}               |
| create_payment_contract         | discount_id              | {wallet_distribution} |
| create_payment_contract         | can_deauthorise          | {can_deauthorise}     |

## MsgGrantDiscount
| Type                            | Attribute Key            | Attribute Value       |
|---------------------------------|--------------------------|-----------------------|
| grant_discount                  | sender_did               | {sender_did}          |
| grant_discount                  | payment_contract_id      | {payment_contract_id} |
| grant_discount                  | discount_id              | {discount_id}         |
| grant_discount                  | Recipient                | {Recipient}           |

## MsgSetPaymentContractAuthorisation
| Type                            | Attribute Key            | Attribute Value       |
|---------------------------------|--------------------------|-----------------------|
| payment_contract_authorisation  | payer-did                | {payer-did}           |
| payment_contract_authorisation  | payment_contract_id      | {payment_contract_id} |
| payment_contract_authorisation  | authorised               | {authorised}          |

## MsgCreateSubscription
| Type                            | Attribute Key            | Attribute Value       |
|---------------------------------|--------------------------|-----------------------|
| create_subscription             | subscription_id          | {subscription_id}     | 
| create_subscription             | payment_contract_id      | {payment_contract_id} |
| create_subscription             | max_periods              | {max_periods}         |
| create_subscription             | key_period               | {key_period}          |
| create_subscription             | creator_did              | {creator_did}         |

## MsgRevokeDiscount
| Type                            | Attribute Key            | Attribute Value       |
|---------------------------------|--------------------------|-----------------------|
| revoke_discount                 | sender_did               | {sender_did}          |
| revoke_discount                 | payment_contract_id      | {payment_contract_id} |
| revoke_discount                 | holder                   | {holder}              |

## MsgEffectPayment
| Type                            | Attribute Key            | Attribute Value       |
|---------------------------------|--------------------------|-----------------------|
| effect_payment                  | sender_did               | {sender_did}          |
| effect_payment                  | payment_contract_id      | {payment_contract_id} |
