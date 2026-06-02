# Claims module — `x/claims`

**Proto package:** `ixo.claims.v1beta1` · **TS typeUrl prefix:** `/ixo.claims.v1beta1.` · **CLI:** `ixod tx claims …` / `ixod query claims …`

## Purpose
The claims module manages W3C Verifiable Claims through their full impact-verification lifecycle. A **Collection** (tied to an entity DID + a protocol/oracle DID) groups claims submitted under one protocol; service agents submit **Claims**, oracle/evaluation agents record **Evaluations**, and **Payments** flow to agents/oracles on submission, evaluation, approval, or rejection. v7 adds **Intents** (pre-work payment guarantees with escrow), **Disputes** with performance-deposit staking and adjudication, and **team member budgets**. Agents act on behalf of the entity via cosmos `authz` grants (`SubmitClaimAuthorization`, `EvaluateClaimAuthorization`), so almost every claim/evaluate message is signed by the collection **admin** account and executed through `MsgExec`.

## Concepts & state
- **Collection** (`claims.proto` `Collection`): incrementing `id` (from `Params.collection_sequence`); holds `entity` DID, `admin` address (the authz grantor, an entity `admin` account), `protocol` DID, dates, `quota`, internal counters, `state` (`CollectionState`), `payments` (`Payments`), `intents` (`CollectionIntentOptions`), `escrow_account`, plus v7 dispute/deposit config.
- **Claim** (`Claim`): keyed by user-supplied `claim_id` (cid hash). Carries `agent_did`/`agent_address`, `submission_date`, current `evaluation`, `payments_status` (`ClaimPayments`), `use_intent`, custom payment amounts, `member_address`, and `evaluation_history` (superseded evaluations after FLAGGED re-evaluation).
- **Evaluation** (`Evaluation`): result of evaluating a claim — `oracle` DID, evaluator `agent_did`/`agent_address`, `status` (`EvaluationStatus`), `reason` code, `verification_proof` cid, custom payout amounts.
- **Dispute** (`Dispute`): keyed by `data.proof` cid; `subject_id` (claim id), `target_role` (`DisputeTargetRole`), `disputer_address`/`disputer_did`, locked `dispute_deposit`, `status` (`DisputeStatus`), populated `resolution` (`DisputeResolution`) on adjudication.
- **Intent** (`Intent`): incrementing `id` (from `Params.intent_sequence`); a service agent's pre-declared intention to submit a claim, with payment moved to escrow. `status` (`IntentStatus`), `expire_at`, `member_address`, `from_address`/`escrow_address`.
- **Payments / Payment** (`Payments`, `Payment`): `Payments` bundles four optional `Payment` slots — `submission`, `evaluation`, `approval`, `rejection`. Each `Payment` has `account` (entity account to pay from), `amount` (`Coins`), `cw20_payment[]`, `cw1155_payment[]`, `timeout_ns` (0 = immediate, else delayed authz withdrawal), and `is_oracle_payment` (APPROVAL only; splits via network fees).
- **CW20Payment** (`CW20Payment`): `address` (CW20 contract), `amount` (uint64, field number 3).
- **CW1155Payment** (`CW1155Payment`): `address`, `token_id[]` (repeated; empty = any token id), `amount` (uint64). Preferred over the deprecated `Contract1155Payment`.
- **Contract1155Payment** (`Contract1155Payment`, DEPRECATED): `address`, `token_id` (single string), `amount` (uint32). Replaced by `CW1155Payment`.
- **Payment timeouts / escrow**: `timeout_ns > 0` defers payout — an authz `WithdrawPaymentAuthorization` is granted with `release_date = created + timeout`; grantee later runs `MsgWithdrawPayment` via `MsgExec`. Intent funds + agent deposits + locked dispute deposits all share the collection's one `escrow_account`.
- **AgentDepositBalance** (`AgentDepositBalance`): rolling per-(collection,agent) performance-deposit balance held in escrow; gates submit/evaluate when the collection requires it; slashed on lost disputes; `withdrawable_at` lock from `min_deposit_period`.
- **MemberBudget** (`MemberBudget`): per-(collection,member) periodic spend cap (`period_spend_limit`, `period_cw20_spend_limit`) with lazy reset; opt-in "team/enterprise" mode.
- **EvaluationStatus** enum: `PENDING=0`, `APPROVED=1`, `REJECTED=2`, `DISPUTED=3` (deprecated for new txs), `INVALIDATED=4`, `FLAGGED=5` (non-terminal; no payment, quota still consumed, re-evaluatable).
- **CollectionState** enum: `OPEN=0`, `PAUSED=1`, `CLOSED=2`.
- **CollectionIntentOptions** enum: `ALLOW=0`, `DENY=1`, `REQUIRED=2`.
- **IntentStatus** enum: `ACTIVE=0`, `FULFILLED=1`, `EXPIRED=2`.
- **PaymentType** enum: `SUBMISSION=0`, `APPROVAL=1`, `EVALUATION=2`, `REJECTION=3`.
- **PaymentStatus** enum: `NO_PAYMENT=0`, `PROMISED=1`, `AUTHORIZED=2`, `GUARANTEED=3`, `PAID=4`, `FAILED=5`, `DISPUTED_PAYMENT=6`.
- **DisputeTargetRole** enum: `DISPUTE_TARGET_ROLE_UNSPECIFIED=0`, `DISPUTE_TARGET_ROLE_SUBMITTER=1`, `DISPUTE_TARGET_ROLE_EVALUATOR=2`.
- **DisputeStatus** enum: `DISPUTE_STATUS_OPEN=0`, `DISPUTE_STATUS_AWARDED=1`, `DISPUTE_STATUS_DISMISSED=2`.
- **CreateClaimAuthorizationType** enum (`authz.proto`): `ALL=0`, `SUBMIT=1`, `EVALUATE=2`.
- **authz constraints**: `SubmitClaimAuthorization`, `EvaluateClaimAuthorization`, `WithdrawPaymentAuthorization`, `CreateClaimAuthorizationAuthorization` (see [Authz authorizations](#authz-authorizations)).

## Messages

### MsgCreateCollection
- **Purpose:** Create a new Collection defining protocol, dates, quota, state, payments, intent policy, and optional dispute/deposit config.
- **Signer / auth:** `signer` field. Keeper requires `entity` to resolve, `protocol` DID to exist, and `signer` to be the entity NFT owner (`CheckIfOwner`). The persisted `admin` is the entity's `admin` account (`EntityAdminAccountName`), not `signer`. All payment `account`s must be entity accounts of `entity`.
- **Fields:**

| Field | Type | Req | Description |
|---|---|---|---|
| entity | string | yes | DID of the entity the claims are created for |
| signer | string | yes | signer address (must be entity NFT owner) |
| protocol | string | yes | DID of the claim protocol |
| start_date | google.protobuf.Timestamp (stdtime) | no | date after which claims may be submitted |
| end_date | google.protobuf.Timestamp (stdtime) | no | date after which no more claims (zero = none) |
| quota | uint64 | no | max claims, 0 = unlimited |
| state | CollectionState (enum) | yes | open / paused / closed |
| payments | Payments | yes | submission/evaluation/approval/rejection payments |
| intents | CollectionIntentOptions (enum) | yes | allow / deny / required |
| service_agent_deposit_required | repeated Coin (Coins, non-nullable) | no | min SA performance deposit to submit |
| evaluator_deposit_required | repeated Coin (Coins) | no | min evaluator deposit to evaluate |
| dispute_deposit_amount | repeated Coin (Coins) | no | stake a disputer locks per dispute |
| penalty_amount_per_dispute | repeated Coin (Coins) | no | fixed AWARDED penalty (≤ each deposit-required) |
| min_deposit_period | google.protobuf.Duration (stdduration, non-nullable) | no | deposit withdrawal lock after top-up |
| adjudicators | repeated AdjudicationDid | no | whitelist; required non-empty if any deposit/penalty field set |

  `AdjudicationDid`: `did` (string), `reward_percentage` (string, LegacyDec, range `[0,100]`).
  `Payments`: `submission`/`evaluation`/`approval`/`rejection`, each a `Payment` (see Concepts). `Payment`: `account` (string), `amount` (Coins), `contract_1155_payment` (Contract1155Payment, DEPRECATED), `timeout_ns` (Duration), `cw20_payment` (repeated CW20Payment), `is_oracle_payment` (bool), `cw1155_payment` (repeated CW1155Payment).
- **CLI:** `ixod tx claims create-collection [create-collection-doc]` — single arg is raw JSON of `MsgCreateCollection`. Plus standard tx flags (`--from`, `--chain-id`, `--fees`/`--gas`, etc.).
- **TS SDK:**
  ```ts
  const msg = {
    typeUrl: '/ixo.claims.v1beta1.MsgCreateCollection',
    value: ixo.claims.v1beta1.MsgCreateCollection.fromPartial({
      entity, signer, protocol, state, payments, intents,
    }),
  };
  ```
- **Gotchas:** `ErrDidDocumentNotFound` if entity/protocol missing; `unauthorized` if signer not owner; evaluation payment may not have CW20 (`ErrCollectionEvalCW20Error`, rejected when `len > 1`) or CW1155 (`ErrCollectionEvalCW1155Error`); `ErrCollNotEntityAcc` if any payment account is not an entity account. Escrow account is auto-created.

### MsgSubmitClaim
- **Purpose:** Submit a new Claim to a Collection; fires SUBMISSION payment if configured.
- **Signer / auth:** `admin_address` (must equal `collection.Admin`). The actual SA acts via a `SubmitClaimAuthorization` granted to them by the admin; submission is normally wrapped in `MsgExec` so the collection module account signs as admin.
- **Fields:**

| Field | Type | Req | Description |
|---|---|---|---|
| collection_id | string | yes | Collection this claim belongs to |
| claim_id | string | yes | unique claim id (cid hash) |
| agent_did | string (casttype DIDFragment) | yes | DID of submitting agent |
| agent_address | string | yes | address of submitting agent |
| admin_address | string | yes | signs msg; validated against collection admin |
| use_intent | bool | no | use active intent; overrides custom amounts |
| amount | repeated Coin (Coins, non-nullable) | no | custom approval amount; empty = collection default |
| cw20_payment | repeated CW20Payment | no | custom CW20 approval payment |
| cw1155_payment | repeated CW1155Payment | no | custom CW1155 approval payment |
| member_address | string | no | team member this claim is for; must match intent's |
- **CLI:** `ixod tx claims submit-claim [submit-claim-doc]` — raw JSON of `MsgSubmitClaim` + tx flags.
- **TS SDK:**
  ```ts
  const msg = {
    typeUrl: '/ixo.claims.v1beta1.MsgSubmitClaim',
    value: ixo.claims.v1beta1.MsgSubmitClaim.fromPartial({
      collectionId, claimId, agentDid, agentAddress, adminAddress,
    }),
  };
  ```
- **Gotchas:** `ErrClaimDuplicate` if `claim_id` exists; `ErrClaimUnauthorized` if `admin_address != collection.Admin`; `ErrCollectionNotOpen` if state ≠ open; `ErrClaimCollectionNotStarted` / `ErrClaimCollectionEnded`; `ErrClaimCollectionQuotaReached`; intent-required/deny violations (`ErrInvalidRequest`); `ErrAgentHasActiveDispute` (open dispute on SUBMITTER role); `ErrAgentDepositInsufficient`; `use_intent=true` needs an active non-expired intent (`ErrIntentNotFound`); `member_address` must equal the intent's (`ErrMemberAddressMismatch`) and must be empty when `use_intent=false`; oracle approval payment without intent rejects non-native (`ErrOraclePaymentOnlyNative`).

### MsgEvaluateClaim
- **Purpose:** Record an Evaluation for a claim; fires EVALUATION payment (to evaluator/oracle) and, on APPROVED, the APPROVAL payment to the submitter.
- **Signer / auth:** `admin_address` (must equal `collection.Admin`). The evaluator acts via an `EvaluateClaimAuthorization`; executed through `MsgExec` with the admin as grantor.
- **Fields:**

| Field | Type | Req | Description |
|---|---|---|---|
| claim_id | string | yes | claim to evaluate |
| collection_id | string | yes | must match the claim's collection |
| oracle | string | yes | DID of the Oracle entity |
| agent_did | string (casttype DIDFragment) | yes | DID of evaluating agent |
| agent_address | string | yes | address of evaluating agent |
| admin_address | string | yes | signs msg; validated against collection admin |
| status | EvaluationStatus (enum) | yes | approved/rejected/invalidated/flagged (DISPUTED rejected) |
| reason | uint32 | no | evaluator-defined reason code |
| verification_proof | string | no | cid of evaluation VC |
| amount | repeated Coin (Coins, non-nullable) | no | custom approval payout; ignored if intent used |
| cw20_payment | repeated CW20Payment | no | custom CW20 approval payout |
| cw1155_payment | repeated CW1155Payment | no | custom CW1155 approval payout |
- **CLI:** `ixod tx claims evaluate-claim [evaluate-claim-doc]` — raw JSON of `MsgEvaluateClaim` + tx flags.
- **TS SDK:**
  ```ts
  const msg = {
    typeUrl: '/ixo.claims.v1beta1.MsgEvaluateClaim',
    value: ixo.claims.v1beta1.MsgEvaluateClaim.fromPartial({
      claimId, collectionId, oracle, agentDid, agentAddress, adminAddress, status,
    }),
  };
  ```
- **Gotchas:** `ErrEvaluateWrongCollection` if collection mismatch; `status = DISPUTED` rejected with `ErrEvaluationStatusDisputedDeprecated` (use `MsgDisputeClaim`); terminal prior evaluation locks the claim (`ErrClaimDuplicateEvaluation`); only `FLAGGED` claims may be re-evaluated; same agent can't re-flag (`ErrSelfReFlag`, checks current + history); `ErrClaimUnauthorized`; `ErrAgentHasActiveDispute` (EVALUATOR role); `ErrAgentDepositInsufficient`; `INVALIDATED` and `FLAGGED` skip payments (only `INVALIDATED` skips quota decrement); intent claims override custom amounts with intent amounts and pay APPROVED out of escrow.

### MsgDisputeClaim
- **Purpose:** File a dispute against a claim's submitter or evaluator role.
- **Signer / auth:** `agent_address`. Open to **anyone with a registered IID** (the IID ante requires `agent_address` to be a key on `agent_did`); no module-level admin/controller/authz check. Economic gate is the staked `dispute_deposit_amount`.
- **Fields:**

| Field | Type | Req | Description |
|---|---|---|---|
| subject_id | string | yes | claim id being disputed |
| agent_did | string (casttype DIDFragment) | yes | DID of disputer (not persisted on record) |
| agent_address | string | yes | address of disputer |
| dispute_type | int32 | no | client-interpreted type |
| data | DisputeData | yes | dispute payload; `data.proof` is the record key |
| target_role | DisputeTargetRole (enum) | yes | SUBMITTER or EVALUATOR (UNSPECIFIED rejected) |

  `DisputeData`: `uri` (string), `type` (string), `proof` (string), `encrypted` (bool).
- **CLI:** `ixod tx claims dispute-claim [dispute-claim-doc]` — raw JSON of `MsgDisputeClaim` + tx flags.
- **TS SDK:**
  ```ts
  const msg = {
    typeUrl: '/ixo.claims.v1beta1.MsgDisputeClaim',
    value: ixo.claims.v1beta1.MsgDisputeClaim.fromPartial({
      subjectId, agentDid, agentAddress, data, targetRole,
    }),
  };
  ```
- **Gotchas:** `ErrDisputeDuplicate` if `data.proof` exists; EVALUATOR role requires an existing terminal evaluation (`ErrDisputeTargetEvaluatorNoEvaluation`); disputing a FLAGGED evaluation rejected (`ErrDisputeTargetEvaluatorFlagged`); one OPEN dispute per `(subject_id, target_role)` — AWARDED permanently blocks, DISMISSED allows refiling (`ErrDisputeAlreadyOpenForSubjectRole` / `ErrDisputeAlreadyAwardedForSubjectRole`); locking a non-zero deposit requires a non-empty `adjudicators` whitelist (`ErrAdjudicationNotConfigured`); deposit is moved into escrow and snapshotted on the record.

### MsgWithdrawPayment
- **Purpose:** Execute a delayed (timeout-based) payment withdrawal via a multi-send, after the payment's release date.
- **Signer / auth:** `admin_address` (must equal `collection.Admin`). For delayed payouts the receiver holds a `WithdrawPaymentAuthorization` and runs this through `MsgExec` after `release_date`.
- **Fields:**

| Field | Type | Req | Description |
|---|---|---|---|
| claim_id | string | yes | claim the withdrawal is for |
| inputs | repeated cosmos.bank.v1beta1.Input (non-nullable) | no | multi-send inputs |
| outputs | repeated cosmos.bank.v1beta1.Output (non-nullable) | no | multi-send outputs |
| payment_type | PaymentType (enum) | yes | submission/approval/evaluation/rejection |
| contract_1155_payment | Contract1155Payment (DEPRECATED) | no | use cw1155_payment instead |
| toAddress | string | no | contract payment recipient |
| fromAddress | string | no | contract payment source |
| release_date | google.protobuf.Timestamp (stdtime) | no | earliest execution date |
| admin_address | string | yes | signs msg; validated against collection admin |
| cw20_payment | repeated CW20Payment | no | CW20 payments to make |
| cw1155_payment | repeated CW1155Payment | no | CW1155 payments to make |
- **CLI:** `ixod tx claims withdraw-payment [withdraw-payment-doc]` — raw JSON of `MsgWithdrawPayment` + tx flags.
- **TS SDK:**
  ```ts
  const msg = {
    typeUrl: '/ixo.claims.v1beta1.MsgWithdrawPayment',
    value: ixo.claims.v1beta1.MsgWithdrawPayment.fromPartial({
      claimId, paymentType, adminAddress, fromAddress, toAddress,
    }),
  };
  ```
- **Gotchas:** `from_address` and any input `address` may not be the collection escrow account (`ErrInvalidRequest`) — escrow funds only flow through intents; `ErrClaimUnauthorized` if admin mismatch.

### MsgUpdateCollectionState
- **Purpose:** Update a Collection's `state` (open/paused/closed).
- **Signer / auth:** `admin_address` = `collection.Admin`.
- **Fields:**

| Field | Type | Req | Description |
|---|---|---|---|
| collection_id | string | yes | collection to update |
| state | CollectionState (enum) | yes | new state |
| admin_address | string | yes | signer; must equal collection admin |
- **CLI:** `ixod tx claims update-collection-state [json]` — raw JSON of `MsgUpdateCollectionState` + tx flags.
- **TS SDK:**
  ```ts
  const msg = {
    typeUrl: '/ixo.claims.v1beta1.MsgUpdateCollectionState',
    value: ixo.claims.v1beta1.MsgUpdateCollectionState.fromPartial({
      collectionId, state, adminAddress,
    }),
  };
  ```
- **Gotchas:** `ErrClaimUnauthorized` on admin mismatch.

### MsgUpdateCollectionDates
- **Purpose:** Update a Collection's `start_date` / `end_date`.
- **Signer / auth:** `admin_address` = `collection.Admin`.
- **Fields:**

| Field | Type | Req | Description |
|---|---|---|---|
| collection_id | string | yes | collection to update |
| start_date | google.protobuf.Timestamp (stdtime) | no | new start date |
| end_date | google.protobuf.Timestamp (stdtime) | no | new end date (zero = none) |
| admin_address | string | yes | signer; must equal collection admin |
- **CLI:** `ixod tx claims update-collection-dates [json]` — raw JSON of `MsgUpdateCollectionDates` + tx flags.
- **TS SDK:**
  ```ts
  const msg = {
    typeUrl: '/ixo.claims.v1beta1.MsgUpdateCollectionDates',
    value: ixo.claims.v1beta1.MsgUpdateCollectionDates.fromPartial({
      collectionId, adminAddress,
    }),
  };
  ```
- **Gotchas:** `ErrClaimUnauthorized` on admin mismatch.

### MsgUpdateCollectionPayments
- **Purpose:** Replace a Collection's `payments` config.
- **Signer / auth:** `admin_address` = `collection.Admin`.
- **Fields:**

| Field | Type | Req | Description |
|---|---|---|---|
| collection_id | string | yes | collection to update |
| payments | Payments | yes | new payments |
| admin_address | string | yes | signer; must equal collection admin |
- **CLI:** `ixod tx claims update-collection-payments [json]` — raw JSON of `MsgUpdateCollectionPayments` + tx flags.
- **TS SDK:**
  ```ts
  const msg = {
    typeUrl: '/ixo.claims.v1beta1.MsgUpdateCollectionPayments',
    value: ixo.claims.v1beta1.MsgUpdateCollectionPayments.fromPartial({
      collectionId, payments, adminAddress,
    }),
  };
  ```
- **Gotchas:** Same payment validations as create: evaluation payment CW20 rejected when `len > 1` (`ErrCollectionEvalCW20Error`), CW1155 when `len > 1` (`ErrCollectionEvalCW1155Error`); accounts must be entity accounts (`ErrCollNotEntityAcc`); `ErrClaimUnauthorized` on admin mismatch.

### MsgUpdateCollectionIntents
- **Purpose:** Update a Collection's intent policy (`intents`).
- **Signer / auth:** `admin_address` = `collection.Admin`.
- **Fields:**

| Field | Type | Req | Description |
|---|---|---|---|
| collection_id | string | yes | collection to update |
| intents | CollectionIntentOptions (enum) | yes | allow / deny / required |
| admin_address | string | yes | signer; must equal collection admin |
- **CLI:** `ixod tx claims update-collection-intents [json]` — raw JSON of `MsgUpdateCollectionIntents` + tx flags.
- **TS SDK:**
  ```ts
  const msg = {
    typeUrl: '/ixo.claims.v1beta1.MsgUpdateCollectionIntents',
    value: ixo.claims.v1beta1.MsgUpdateCollectionIntents.fromPartial({
      collectionId, intents, adminAddress,
    }),
  };
  ```
- **Gotchas:** `ErrClaimUnauthorized` on admin mismatch.

### MsgUpdateCollectionQuota
- **Purpose:** Update a Collection's max-claim `quota`.
- **Signer / auth:** `admin_address` = `collection.Admin`.
- **Fields:**

| Field | Type | Req | Description |
|---|---|---|---|
| collection_id | string | yes | collection to update |
| quota | uint64 | yes | new max; 0 = unlimited; must be 0 or ≥ count |
| admin_address | string | yes | signer; must equal collection admin |
- **CLI:** `ixod tx claims update-collection-quota [json]` — raw JSON of `MsgUpdateCollectionQuota` + tx flags.
- **TS SDK:**
  ```ts
  const msg = {
    typeUrl: '/ixo.claims.v1beta1.MsgUpdateCollectionQuota',
    value: ixo.claims.v1beta1.MsgUpdateCollectionQuota.fromPartial({
      collectionId, quota, adminAddress,
    }),
  };
  ```
- **Gotchas:** new quota below current `count` (and non-zero) rejected with `ErrCollectionQuotaBelowCount`; `ErrClaimUnauthorized` on admin mismatch.

### MsgClaimIntent
- **Purpose:** Create an intent to submit a claim, moving the approval amount into escrow as a payment guarantee.
- **Signer / auth:** `agent_address`. Agent must hold a `SubmitClaimAuthorization` from the collection admin with a constraint matching `(collection_id, member_address)`.
- **Fields:**

| Field | Type | Req | Description |
|---|---|---|---|
| agent_did | string (casttype DIDFragment) | yes | DID of service agent |
| agent_address | string | yes | submitter of this message |
| collection_id | string | yes | collection this intent is for |
| amount | repeated Coin (Coins, non-nullable) | no | desired claim amount; empty = collection APPROVAL default |
| cw20_payment | repeated CW20Payment | no | custom CW20 payment |
| cw1155_payment | repeated CW1155Payment | no | custom CW1155 payment |
| member_address | string | no | team member; required iff collection has member budgets |
- **CLI:** `ixod tx claims claim-intent [json]` — raw JSON of `MsgClaimIntent` + tx flags. (Note: a `submit-intent` command function exists in source but is NOT wired into the CLI; use `claim-intent`.)
- **TS SDK:**
  ```ts
  const msg = {
    typeUrl: '/ixo.claims.v1beta1.MsgClaimIntent',
    value: ixo.claims.v1beta1.MsgClaimIntent.fromPartial({
      agentDid, agentAddress, collectionId,
    }),
  };
  ```
- **Gotchas:** `ErrIntentExists` if agent already has an active intent on the collection; `ErrIntentUnauthorized` if intents denied / no authz / no matching constraint / amounts exceed constraint maxes; `ErrMemberAddressRequired` / `ErrMemberAddressNotAllowed` for member-budget mismatch; `ErrMemberBudgetNotFound` / `ErrMemberBudgetExceeded` on budget checks; response returns `intent_id` and `expire_at` (created + constraint `intent_duration_ns`).

### MsgCreateClaimAuthorization
- **Purpose:** Create a `SubmitClaimAuthorization` or `EvaluateClaimAuthorization` on a grantee, on behalf of an entity admin account.
- **Signer / auth:** `admin_address` (= `collection.Admin`). The creator must hold a `CreateClaimAuthorizationAuthorization` (meta-authorization) with a constraint permitting the requested type/collection/member. The grant is routed through the entity keeper (`MsgGrantEntityAccountAuthz`).
- **Fields:**

| Field | Type | Req | Description |
|---|---|---|---|
| creator_address | string | yes | creator (holder of meta-authorization) |
| creator_did | string (casttype DIDFragment) | yes | DID of creator |
| grantee_address | string | yes | who receives the authorization |
| admin_address | string | yes | signs msg; validated against collection admin |
| collection_id | string | yes | collection the authorization applies to |
| auth_type | CreateClaimAuthorizationType (enum) | yes | SUBMIT or EVALUATE (ALL rejected) |
| agent_quota | uint64 | no | quota for the created authorization |
| max_amount | repeated Coin (Coins, non-nullable) | no | max amount settable in the authorization |
| max_cw20_payment | repeated CW20Payment | no | max CW20 settable |
| expiration | google.protobuf.Timestamp (stdtime) | no | authorization expiry (removes constraints when hit) |
| intent_duration_ns | google.protobuf.Duration (stdduration, non-nullable) | no | max intent duration (submit) |
| before_date | google.protobuf.Timestamp (stdtime) | no | evaluate cutoff date |
| max_cw1155_payment | repeated CW1155Payment | no | max CW1155 settable |
| member_address | string | no | team member; must match meta-auth constraint |
- **CLI:** `ixod tx claims create-claim-authorization [json]` — raw JSON of `MsgCreateClaimAuthorization` + tx flags.
- **TS SDK:**
  ```ts
  const msg = {
    typeUrl: '/ixo.claims.v1beta1.MsgCreateClaimAuthorization',
    value: ixo.claims.v1beta1.MsgCreateClaimAuthorization.fromPartial({
      creatorAddress, creatorDid, granteeAddress, adminAddress, collectionId, authType,
    }),
  };
  ```
- **Gotchas:** `ErrClaimUnauthorized` on admin mismatch; `auth_type = ALL` rejected (`ErrInvalidRequest` — create submit/evaluate separately); a pre-existing `GenericAuthorization` for the grantee/route must be removed first; an existing typed authorization gets the new constraint appended; SUBMIT maps to `SubmitClaimConstraints`, EVALUATE to `EvaluateClaimConstraints` (with `max_amount`→`max_custom_amount`, etc.); `member_address` anti-spoofing enforced in `Accept()`.

### MsgSetCollectionMembers
- **Purpose:** Add or update one or more team member budgets on a collection.
- **Signer / auth:** `admin_address` = `collection.Admin`.
- **Fields:**

| Field | Type | Req | Description |
|---|---|---|---|
| collection_id | string | yes | collection to add/update members on |
| admin_address | string | yes | signer; must equal collection admin |
| members | repeated CollectionMemberInput | yes | member budgets to set (no duplicates) |

  `CollectionMemberInput`: `member_address` (string), `period` (Duration, stdduration, non-nullable, ≥ 24h), `period_spend_limit` (Coins, non-nullable), `period_cw20_spend_limit` (repeated CW20Payment), `reset_period_spent` (bool).
- **CLI:** `ixod tx claims set-collection-members [set-collection-members-doc]` — raw JSON of `MsgSetCollectionMembers` + tx flags.
- **TS SDK:**
  ```ts
  const msg = {
    typeUrl: '/ixo.claims.v1beta1.MsgSetCollectionMembers',
    value: ixo.claims.v1beta1.MsgSetCollectionMembers.fromPartial({
      collectionId, adminAddress, members,
    }),
  };
  ```
- **Gotchas:** `ErrClaimUnauthorized` on admin mismatch; a member entry with all-zero spend limits is rejected (`ErrMemberBudgetZero` — use remove instead); `period` shorter than 24h (`MinMemberBudgetPeriod`) rejected; existing members preserve `period_spent`/`period_reset_at` unless `reset_period_spent`; new members emit `MemberBudgetCreatedEvent`, updates emit `MemberBudgetUpdatedEvent`. Adding the first member turns the collection into team mode (all future intents need `member_address`).

### MsgRemoveCollectionMembers
- **Purpose:** Remove one or more member budgets from a collection.
- **Signer / auth:** `admin_address` = `collection.Admin`.
- **Fields:**

| Field | Type | Req | Description |
|---|---|---|---|
| collection_id | string | yes | collection to remove members from |
| admin_address | string | yes | signer; must equal collection admin |
| member_addresses | repeated string | yes | member addresses to remove |
- **CLI:** `ixod tx claims remove-collection-members [remove-collection-members-doc]` — raw JSON of `MsgRemoveCollectionMembers` + tx flags.
- **TS SDK:**
  ```ts
  const msg = {
    typeUrl: '/ixo.claims.v1beta1.MsgRemoveCollectionMembers',
    value: ixo.claims.v1beta1.MsgRemoveCollectionMembers.fromPartial({
      collectionId, adminAddress, memberAddresses,
    }),
  };
  ```
- **Gotchas:** `ErrClaimUnauthorized` on admin mismatch; each address must currently exist (`ErrMemberBudgetNotFound`) — whole tx rolls back atomically otherwise; does NOT revoke the members' authz grants; emits `MemberBudgetRemovedEvent` per member.

### MsgUpdateCollectionDisputeConfig
- **Purpose:** Replace the full dispute / performance-deposit configuration on a collection.
- **Signer / auth:** `admin_address` = `collection.Admin` (typically via `MsgExec` admin pattern).
- **Fields:**

| Field | Type | Req | Description |
|---|---|---|---|
| collection_id | string | yes | collection to update |
| admin_address | string | yes | signer; must equal collection admin |
| service_agent_deposit_required | repeated Coin (Coins, non-nullable) | no | min SA deposit (full replacement) |
| evaluator_deposit_required | repeated Coin (Coins) | no | min evaluator deposit |
| dispute_deposit_amount | repeated Coin (Coins) | no | disputer stake per dispute |
| penalty_amount_per_dispute | repeated Coin (Coins) | no | fixed AWARDED penalty (≤ each deposit-required) |
| min_deposit_period | google.protobuf.Duration (stdduration, non-nullable) | no | deposit withdrawal lock |
| adjudicators | repeated AdjudicationDid | no | whitelist (full replacement) |
- **CLI:** `ixod tx claims update-collection-dispute-config [json]` — raw JSON of `MsgUpdateCollectionDisputeConfig` + tx flags.
- **TS SDK:**
  ```ts
  const msg = {
    typeUrl: '/ixo.claims.v1beta1.MsgUpdateCollectionDisputeConfig',
    value: ixo.claims.v1beta1.MsgUpdateCollectionDisputeConfig.fromPartial({
      collectionId, adminAddress,
    }),
  };
  ```
- **Gotchas:** all fields are full replacements (not merges); `ValidateBasic`/`ValidateCollectionDisputeConfig` enforces penalty ≤ each non-empty deposit-required, non-empty `adjudicators` if any deposit/penalty/stake field set, unique valid DIDs, `reward_percentage` in `[0,100]`, non-negative `min_deposit_period`; clearing `adjudicators` while `disputes_open > 0` rejected (`ErrAdjudicationNotConfigured`); does NOT affect in-flight disputes or existing `withdrawable_at`.

### MsgAddPerformanceDeposit
- **Purpose:** Top up an agent's performance-deposit balance on a collection (funds → escrow).
- **Signer / auth:** `agent_address` (the balance owner; no admin authz required — anyone funds their own balance).
- **Fields:**

| Field | Type | Req | Description |
|---|---|---|---|
| collection_id | string | yes | collection the deposit is held on |
| agent_address | string | yes | balance owner / payer / signer |
| amount | repeated Coin (Coins, non-nullable) | yes | amount to add (strictly positive) |
- **CLI:** `ixod tx claims add-performance-deposit [json]` — raw JSON of `MsgAddPerformanceDeposit` + tx flags.
- **TS SDK:**
  ```ts
  const msg = {
    typeUrl: '/ixo.claims.v1beta1.MsgAddPerformanceDeposit',
    value: ixo.claims.v1beta1.MsgAddPerformanceDeposit.fromPartial({
      collectionId, agentAddress, amount,
    }),
  };
  ```
- **Gotchas:** permitted even with active disputes (needed to clear arrears); first top-up emits `AgentDepositBalanceCreatedEvent`, later ones `AgentDepositBalanceUpdatedEvent`; each top-up rolls `withdrawable_at` to `max(current, now + min_deposit_period)`; response returns `new_balance`.

### MsgWithdrawPerformanceDeposit
- **Purpose:** Withdraw some/all of an agent's performance-deposit balance back to their wallet.
- **Signer / auth:** `agent_address` (must own the balance).
- **Fields:**

| Field | Type | Req | Description |
|---|---|---|---|
| collection_id | string | yes | collection the balance is held on |
| agent_address | string | yes | balance owner / signer |
| amount | repeated Coin (Coins, non-nullable) | no | amount to withdraw; empty = full balance |
- **CLI:** `ixod tx claims withdraw-performance-deposit [json]` — raw JSON of `MsgWithdrawPerformanceDeposit` + tx flags.
- **TS SDK:**
  ```ts
  const msg = {
    typeUrl: '/ixo.claims.v1beta1.MsgWithdrawPerformanceDeposit',
    value: ixo.claims.v1beta1.MsgWithdrawPerformanceDeposit.fromPartial({
      collectionId, agentAddress,
    }),
  };
  ```
- **Gotchas:** rejected while any OPEN dispute targets the agent (`ErrAgentDepositBalanceCannotWithdraw`); rejected if still inside the lock window (`ErrAgentDepositLocked`); amount must be ≤ balance; full drain deletes the entry and emits `AgentDepositBalanceRemovedEvent`, partial emits `AgentDepositBalanceUpdatedEvent`; response returns `withdrawn` and `remaining_balance`.

### MsgAdjudicateDispute
- **Purpose:** Settle an OPEN dispute as AWARDED or DISMISSED, applying penalty/payout math.
- **Signer / auth:** `adjudicator_address`. `adjudicator_did` must be in `collection.adjudicators`; signer must be a registered key on that DID document (`Authentication`/`AssertionMethod`, enforced by IID ante + keeper `AuthorizeAdjudicator`).
- **Fields:**

| Field | Type | Req | Description |
|---|---|---|---|
| subject_id | string | yes | claim id of the dispute |
| target_role | DisputeTargetRole (enum) | yes | SUBMITTER or EVALUATOR — identifies the dispute |
| adjudicator_did | string | yes | adjudicating DID (must be whitelisted) |
| adjudicator_address | string | yes | signer; key on adjudicator_did |
| outcome | DisputeStatus (enum) | yes | AWARDED or DISMISSED (OPEN rejected) |
| data | DisputeData | no | adjudicator opinion doc; if set all 3 strings required |
| penalty_amount | repeated Coin (Coins, non-nullable) | no | AWARDED penalty; ignored if collection has fixed penalty |
- **CLI:** `ixod tx claims adjudicate-dispute [json]` — raw JSON of `MsgAdjudicateDispute` + tx flags.
- **TS SDK:**
  ```ts
  const msg = {
    typeUrl: '/ixo.claims.v1beta1.MsgAdjudicateDispute',
    value: ixo.claims.v1beta1.MsgAdjudicateDispute.fromPartial({
      subjectId, targetRole, adjudicatorDid, adjudicatorAddress, outcome,
    }),
  };
  ```
- **Gotchas:** `ErrDisputeNotOpen` if not OPEN; `ErrAdjudicatorDidNotApproved` if DID not whitelisted; `ErrAdjudicatorNotAuthorized` if signer not a key on the DID; AWARDED without fixed collection penalty requires `penalty_amount` (`ErrPenaltyAmountRequired`), capped at the loser's role deposit-required (`ErrPenaltyAmountExceedsCap`); AWARDED slashes the target agent's balance (winner share to disputer, adjudicator share to payout address per that adjudicator's `reward_percentage`), disputer deposit refunded; DISMISSED makes the disputer's deposit the pot (target agent is winner); actual slash always `min(intended, balance)`; response returns `actual_penalty_paid`.

## Authz authorizations
Granted via the standard cosmos `authz` `MsgGrant` (here usually routed through the entity module's `MsgGrantEntityAccountAuthz`). All four implement `cosmos.authz.v1beta1.Authorization` and carry an `admin` field = the collection admin (entity admin account, the grantor).

- **SubmitClaimAuthorization** (`/ixo.claims.v1beta1.SubmitClaimAuthorization`) — authorizes `MsgSubmitClaim` (and gates `MsgClaimIntent`). Fields: `admin` (string), `constraints` (repeated `SubmitClaimConstraints`). `SubmitClaimConstraints`: `collection_id` (string), `agent_quota` (uint64), `max_amount` (Coins, non-nullable), `max_cw20_payment` (repeated CW20Payment), `intent_duration_ns` (Duration, stdduration, non-nullable), `max_cw1155_payment` (repeated CW1155Payment), `member_address` (string).
- **EvaluateClaimAuthorization** (`/ixo.claims.v1beta1.EvaluateClaimAuthorization`) — authorizes `MsgEvaluateClaim`. Fields: `admin` (string), `constraints` (repeated `EvaluateClaimConstraints`). `EvaluateClaimConstraints`: `collection_id` (string), `claim_ids` (repeated string; empty = any), `agent_quota` (uint64), `before_date` (Timestamp, stdtime), `max_custom_amount` (Coins, field 10), `max_custom_cw20_payment` (repeated CW20Payment, field 11), `max_custom_cw1155_payment` (repeated CW1155Payment, field 12).
- **WithdrawPaymentAuthorization** (`/ixo.claims.v1beta1.WithdrawPaymentAuthorization`) — authorizes `MsgWithdrawPayment` (delayed/timeout payouts). Fields: `admin` (string), `constraints` (repeated `WithdrawPaymentConstraints`). `WithdrawPaymentConstraints`: `claim_id` (string), `inputs`/`outputs` (repeated bank Input/Output, non-nullable), `payment_type` (PaymentType), `contract_1155_payment` (Contract1155Payment, DEPRECATED), `toAddress`/`fromAddress` (string), `release_date` (Timestamp, stdtime), `cw20_payment` (repeated CW20Payment), `cw1155_payment` (repeated CW1155Payment). Auto-created by the keeper on delayed payments; the receiver is the grantee.
- **CreateClaimAuthorizationAuthorization** (`/ixo.claims.v1beta1.CreateClaimAuthorizationAuthorization`) — meta-authorization that authorizes `MsgCreateClaimAuthorization`. Fields: `admin` (string), `constraints` (repeated `CreateClaimAuthorizationConstraints`). `CreateClaimAuthorizationConstraints`: `max_authorizations` (uint64), `max_agent_quota` (uint64), `max_amount` (Coins), `max_cw20_payment` (repeated CW20Payment), `expiration` (Timestamp, stdtime), `collection_ids` (repeated string; empty = all), `allowed_auth_types` (CreateClaimAuthorizationType enum), `max_intent_duration_ns` (Duration, stdduration, non-nullable), `max_cw1155_payment` (repeated CW1155Payment), `member_address` (string, anti-spoofing).

## Queries
No query CLI exists (the `client/cli` package contains only `tx.go`). Use gRPC / REST. All gRPC methods are on `ixo.claims.v1beta1.Query`:

| Query | gRPC method | CLI | Args | Returns |
|---|---|---|---|---|
| Params | `Params` | none | — | `Params` |
| Collection | `Collection` | none | `id` | `Collection` |
| Collection list | `CollectionList` | none | pagination | `[]Collection` + page |
| Claim | `Claim` | none | `id` | `Claim` |
| Claim list | `ClaimList` | none | pagination | `[]Claim` + page |
| Dispute | `Dispute` | none | `proof` | `Dispute` |
| Dispute list | `DisputeList` | none | pagination | `[]Dispute` + page |
| Intent | `Intent` | none | `agentAddress`, `collectionId`, `id` | `Intent` |
| Intent list | `IntentList` | none | pagination | `[]Intent` + page |
| Collection member | `CollectionMember` | none | `collectionId`, `memberAddress` | `MemberBudget` |
| Collection member list | `CollectionMemberList` | none | `collectionId`, pagination | `[]MemberBudget` + page |
| Dispute by subject | `DisputeBySubject` | none | `subjectId`, `targetRole` | `Dispute` |
| Disputes for subject | `DisputeListForSubject` | none | `subjectId` | `[]Dispute` |
| Agent deposit balance | `AgentDepositBalance` | none | `collectionId`, `agentAddress` | `AgentDepositBalance` |
| Agent deposit balance list | `AgentDepositBalanceList` | none | `collectionId`, pagination | `[]AgentDepositBalance` + page |

REST paths are under `/ixo/claims/…` (e.g. `/ixo/claims/collection/{id}`, `/ixo/claims/claims`, `/ixo/claims/dispute-by-subject/{subjectId}/{targetRole}`).

## Events
Typed events from `event.proto`:
- `CollectionCreatedEvent` — on collection creation.
- `CollectionUpdatedEvent` — on any collection update (state/dates/payments/intents/quota/dispute-config/counter changes).
- `ClaimSubmittedEvent` — on claim submission.
- `ClaimUpdatedEvent` — on claim update (evaluation, payment withdrawal).
- `ClaimEvaluatedEvent` — on claim evaluation (carries the `Evaluation`).
- `ClaimDisputedEvent` — on dispute filing (carries the `Dispute`).
- `PaymentWithdrawnEvent` — when a withdrawal payout executes (with `cw20_outputs` and `cw1155_payments`).
- `PaymentWithdrawCreatedEvent` — when a delayed-withdrawal authorization is created.
- `IntentSubmittedEvent` — on intent submission.
- `IntentUpdatedEvent` — on intent update (fulfilled/expired).
- `ClaimAuthorizationCreatedEvent` — on `MsgCreateClaimAuthorization` (`creator`, `creator_did`, `grantee`, `admin`, `collection_id`, `auth_type`).
- `MemberBudgetCreatedEvent` / `MemberBudgetUpdatedEvent` / `MemberBudgetRemovedEvent` — team budget lifecycle.
- `AgentDepositBalanceCreatedEvent` / `AgentDepositBalanceUpdatedEvent` / `AgentDepositBalanceRemovedEvent` — performance-deposit lifecycle.
- `DisputeResolvedEvent` — when `MsgAdjudicateDispute` settles a dispute (AWARDED or DISMISSED).

## Module gotchas
- **Prerequisites:** a Collection requires both an existing entity DID (`entity`) and a protocol DID (`protocol`) registered in the iid module; the `signer` of `MsgCreateCollection` must be the entity NFT owner. The persisted `admin` is the entity's `admin` account, not the signer.
- **Admin-signed via authz:** `MsgSubmitClaim`, `MsgEvaluateClaim`, and the collection-update messages are signed by `admin_address` = `collection.Admin`. Real service/evaluation agents act via `SubmitClaimAuthorization` / `EvaluateClaimAuthorization` grants and execute through cosmos `authz` `MsgExec` (the collection module account signs as admin). Oracle agents must be granted before they can submit/evaluate.
- **Payment accounts must be entity accounts:** every `Payment.account` must be an entity account of the collection's `entity` (`ErrCollNotEntityAcc`). Evaluation payments may not carry CW20/CW1155.
- **Escrow:** the collection's single `escrow_account` holds intent funds, agent performance-deposit balances, and locked dispute deposits. `MsgWithdrawPayment` explicitly forbids escrow as a `from`/input address.
- **Payments flow on evaluation:** APPROVED triggers the APPROVAL payout to the submitter (from the approval account, or from escrow if intent-funded); EVALUATION pays the evaluator/oracle; REJECTED triggers REJECTION; INVALIDATED/FLAGGED pay nothing. `timeout_ns > 0` defers a payment to a `WithdrawPaymentAuthorization` claimable after `release_date`.
- **Oracle payments** (`is_oracle_payment`, APPROVAL only) split via network fees; without an intent only native coins are allowed (`ErrOraclePaymentOnlyNative`), and CW1155 is never allowed for oracle payments.
- **Intents** require a matching `SubmitClaimAuthorization` constraint and, in team collections, a `member_address`; funds move to escrow on intent creation and are released to the agent on APPROVED or refunded on any other terminal status / expiry.
- **Disputes** are open to any registered IID, gated economically by `dispute_deposit_amount`; adjudication requires a whitelisted `adjudicators` DID whose key signs `MsgAdjudicateDispute`.
