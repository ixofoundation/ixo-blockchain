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
     - [Intents](02_state.md#intents)
     - [Member Budgets](02_state.md#member-budgets)
   - [Types](02_state.md#types)
     - [Collection](02_state.md#collection)
     - [Payments](02_state.md#payments)
     - [Payment](02_state.md#payment)
     - [Contract1155Payment](02_state.md#contract1155payment) DEPRECATED, use [CW1155Payment](02_state.md#cw1155payment) instead
     - [CW1155Payment](02_state.md#cw1155payment)
     - [CW1155IntentPayment](02_state.md#cw1155intentpayment)
     - [CW20Payment](02_state.md#cw20payment)
     - [CW20Output](02_state.md#cw20output)
     - [Claim](02_state.md#claim)
     - [ClaimPayments](02_state.md#claimpayments)
     - [Evaluation](02_state.md#evaluation)
     - [Dispute](02_state.md#dispute)
     - [DisputeData](02_state.md#disputedata)
     - [Intent](02_state.md#intent)
     - [MemberBudget](02_state.md#memberbudget)
     - [Enums](02_state.md#enums)
       - [CollectionState](02_state.md#collectionstate)
       - [CollectionIntentOptions](02_state.md#collectionintentoptions)
       - [EvaluationStatus](02_state.md#evaluationstatus)
       - [IntentStatus](02_state.md#intentstatus)
       - [PaymentType](02_state.md#paymenttype)
       - [PaymentStatus](02_state.md#paymentstatus)
     - [Authz Types](02_state.md#authz-types)
       - [WithdrawPaymentAuthorization](02_state.md#withdrawpaymentauthorization)
       - [WithdrawPaymentConstraints](02_state.md#withdrawpaymentconstraints)
       - [SubmitClaimAuthorization](02_state.md#submitclaimauthorization)
       - [SubmitClaimConstraints](02_state.md#submitclaimconstraints)
       - [EvaluateClaimAuthorization](02_state.md#evaluateclaimauthorization)
       - [EvaluateClaimConstraints](02_state.md#evaluateclaimconstraints)
       - [CreateClaimAuthorizationAuthorization](02_state.md#createclaimauthorizationauthorization)
       - [CreateClaimAuthorizationType](02_state.md#createclaimauthorizationtype)
       - [CreateClaimAuthorizationConstraints](02_state.md#createclaimauthorizationconstraints)

3. **[Messages](03_messages.md)**

   - [Messages](03_messages.md#messages)
     - [MsgCreateCollection](03_messages.md#msgcreatecollection)
     - [MsgUpdateCollectionState](03_messages.md#msgupdatecollectionstate)
     - [MsgUpdateCollectionDates](03_messages.md#msgupdatecollectiondates)
     - [MsgUpdateCollectionPayments](03_messages.md#msgupdatecollectionpayments)
     - [MsgUpdateCollectionIntents](03_messages.md#msgupdatecollectionintents)
     - [MsgSubmitClaim](03_messages.md#msgsubmitclaim)
     - [MsgEvaluateClaim](03_messages.md#msgevaluateclaim)
     - [MsgDisputeClaim](03_messages.md#msgdisputeclaim)
     - [MsgWithdrawPayment](03_messages.md#msgwithdrawpayment)
     - [MsgClaimIntent](03_messages.md#msgclaimintent)
     - [MsgCreateClaimAuthorization](03_messages.md#msgcreateclaimauthorization)
     - [MsgSetCollectionMembers](03_messages.md#msgsetcollectionmembers)
     - [MsgRemoveCollectionMembers](03_messages.md#msgremovecollectionmembers)

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
     - [IntentSubmittedEvent](04_events.md#intentsubmittedevent)
     - [IntentUpdatedEvent](04_events.md#intentupdatedevent)
     - [ClaimAuthorizationCreatedEvent](04_events.md#claimauthorizationcreatedevent)
     - [MemberBudgetCreatedEvent](04_events.md#memberbudgetcreatedevent)
     - [MemberBudgetUpdatedEvent](04_events.md#memberbudgetupdatedevent)
     - [MemberBudgetRemovedEvent](04_events.md#memberbudgetremovedevent)

5. **[Parameters](05_params.md)**

6. **[Future Improvements](06_future_improvements.md)**
