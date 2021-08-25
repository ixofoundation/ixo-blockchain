# Payment module specification

## Contents

1. **[State](01_state.md)**
    - [Payment templates](01_state.md#payment-templates)
    - [Payment contracts](01_state.md#payment-contracts)
    - [Subscriptions](01_state.md#subscriptions)
1. **[Messages](02_messages.md)**
    - [MsgCreatePaymentTemplate](02_messages.md#MsgCreatePaymentTemplate)
    - [MsgCreatePaymentContract](02_messages.md#MsgCreatePaymentContract)
    - [MsgCreateSubscription](02_messages.md#MsgCreateSubscription)
    - [MsgSetPaymentContractAuthorisation](02_messages.md#MsgSetPaymentContractAuthorisation)
    - [MsgGrantDiscount](02_messages.md#MsgGrantDiscount)
    - [MsgRevokeDiscount](02_messages.md#MsgRevokeDiscount)
    - [MsgEffectPayment](02_messages.md#MsgEffectPayment)
1. **[End-Block](03_end_block.md)**
1. **[Events](04_events.md)**
    - [EndBlocker](04_events.md#endblocker)
    - [Handlers](04_events.md#handlers)
