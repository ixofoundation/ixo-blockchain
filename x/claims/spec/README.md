# Claims module specification

This document specifies the claims module, a custom Ixo Cosmos SDK module.

The Claims Module provides an advanced structure for handling Verifiable Claims (VCs), cryptographic attestations regarding a subject. By aligning with the W3C standard and incorporating unique ixo system identifiers, this module offers a comprehensive solution for creating, evaluating, and managing claims. It enables entities to define protocols, authorize agents, and maintain a verifiable registry, ensuring authenticity and transparency in all claim-related processes.

## Contents

1. **[Concepts](01_concepts.md)**

   - [Concepts](01_concepts.md#concepts)
     - [Verifiable Credentials (VCs)](01_concepts.md#verifiable-credentials-vcs)
   - [Claims Module](01_concepts.md#claims-module)

2. **[State](02_state.md)**

   - [State](02_state.md#state)
     - [Collections](02_state.md#collections)
     - [Claims](02_state.md#claims)
     - [Disputes](02_state.md#disputes)
   - [Types](02_state.md#types)
     - [Collection](02_state.md#collection)
     - [Payments](02_state.md#payments)
     - [Payment](02_state.md#payment)
     - [Contract1155Payment](02_state.md#contract1155payment) DEPRECATED, use [CW1155Payment](02_state.md#cw1155payment) instead
     - [CW1155Payment](02_state.md#cw1155payment)
     - [CW1155IntentPayment](02_state.md#cw1155intentpayment)
     - [Claim](02_state.md#claim)
     - [ClaimPayments](02_state.md#claimpayments)
     - [Evaluation](02_state.md#evaluation)
     - [Dispute](02_state.md#dispute)
     - [DisputeData](02_state.md#disputedata)
     - [Enums](02_state.md#enums)
       - [CollectionState](02_state.md#collectionstate)
       - [EvaluationStatus](02_state.md#evaluationstatus)
       - [PaymentType](02_state.md#paymenttype)
       - [PaymentStatus](02_state.md#paymentstatus)
     - [Authz Types](02_state.md#authz-types)
       - [WithdrawPaymentAuthorization](02_state.md#withdrawpaymentauthorization)
       - [WithdrawPaymentConstraints](02_state.md#withdrawpaymentconstraints)
       - [SubmitClaimAuthorization](02_state.md#submitclaimauthorization)
       - [SubmitClaimConstraints](02_state.md#submitclaimconstraints)
       - [EvaluateClaimAuthorization](02_state.md#evaluateclaimauthorization)
       - [EvaluateClaimConstraints](02_state.md#evaluateclaimconstraints)

3. **[Messages](03_messages.md)**

   - [Messages](03_messages.md#messages)
     - [MsgCreateCollection](03_messages.md#msgcreatecollection)
     - [MsgSubmitClaim](03_messages.md#msgsubmitclaim)
     - [MsgEvaluateClaim](03_messages.md#msgevaluateclaim)
     - [MsgDisputeClaim](03_messages.md#msgdisputeclaim)
     - [MsgWithdrawPayment](03_messages.md#msgwithdrawpayment)

4. **[Events](04_events.md)**

   - [Events](04_events.md#events)
     - [CollectionCreatedEvent](04_events.md#collectioncreatedevent)
     - [CollectionUpdatedEvent](04_events.md#collectionupdatedevent)
     - [ClaimSubmittedEvent](04_events.md#claimsubmittedevent)
     - [ClaimUpdatedEvent](04_events.md#claimupdatedevent)
     - [ClaimEvaluatedEvent](04_events.md#claimevaluatedevent)
     - [ClaimDisputedEvent](04_events.md#claimdisputedevent)
     - [PaymentWithdrawnEvent](04_events.md#paymentwithdrawnevent)
     - [PaymentWithdrawCreatedEvent](04_events.md#paymentwithdrawcreatedevent)

5. **[Parameters](05_params.md)**

6. **[Future Improvements](06_future_improvements.md)**
