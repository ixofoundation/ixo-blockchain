# Events

In this section we describe the typed protobuf events emitted by the names module. All events are dispatched via `EventManager.EmitTypedEvent` (or `EmitTypedEvents` for the dual-emit cases below) and follow the standard Cosmos SDK convention where the event type string is the fully-qualified message name (e.g. `ixo.names.v1beta1.NameRegisteredEvent`).

Indexers should subscribe to all six event types to maintain a complete picture: two for namespace state, four for name records. The two state-mirror events (`NamespaceCreatedEvent`/`UpdatedEvent`, `NameRegisteredEvent`/`UpdatedEvent`) carry the full updated record so an indexer can refresh its state-mirror table from a single hook regardless of which message caused the change. The two action events (`NameTransferredEvent`, `NameStatusChangedEvent`) carry the *delta* and are intended for audit-history tables.

## Dual-emission pattern

`MsgTransferName` and `MsgSetNameStatus` each emit **two** events in a single call:

1. A `NameUpdatedEvent` carrying the post-change full `NameRecord`.
2. The action-specific event (`NameTransferredEvent` or `NameStatusChangedEvent`) carrying the actor and delta.

This lets an indexer have **one** state-mirror handler keyed on `NameUpdatedEvent` that refreshes the `NameRecord` row regardless of *which* message mutated it, while the action events drive append-only audit tables. The same pattern is used by the entity module (`EntityUpdatedEvent` + `EntityVerifiedUpdatedEvent` on verify, `EntityUpdatedEvent` + `EntityAccountCreatedEvent` on account creation).

`MsgRegisterName` and `MsgRegisterNameByRegistrar` emit only `NameRegisteredEvent` — registration is the canonical "create" event and an indexer can both insert into its state table and append to its registration-history table from the same payload.

`MsgUpdateNameByRegistrar` emits only `NameUpdatedEvent` — there is no separate action event for verification-metadata updates because the diff is fully captured in the record itself.

## NamespaceCreatedEvent

Emitted when a new namespace is registered through [MsgCreateNamespace](./03_messages.md#msgcreatenamespace).

```go
type NamespaceCreatedEvent struct {
    Namespace *Namespace  // tag 1
    Authority string      // tag 2
}
```

The field's descriptions is as follows:

- `namespace` - the full [Namespace](./02_state.md#namespace) record at creation time.
- `authority` - the signer of the gov proposal that created the namespace (i.e., the chain governance module address). Lets indexers attribute namespace mutations to a specific governance proposal.

## NamespaceUpdatedEvent

Emitted when an existing namespace's configuration is replaced via [MsgUpdateNamespace](./03_messages.md#msgupdatenamespace).

```go
type NamespaceUpdatedEvent struct {
    Namespace *Namespace  // tag 1
    Authority string      // tag 2
}
```

The field's descriptions is as follows:

- `namespace` - the full updated [Namespace](./02_state.md#namespace) record.
- `authority` - the gov authority signer of the update.

## NameRegisteredEvent

Emitted on every successful [MsgRegisterName](./03_messages.md#msgregistername) or [MsgRegisterNameByRegistrar](./03_messages.md#msgregisternamebyregistrar).

```go
type NameRegisteredEvent struct {
    Record       *NameRecord  // tag 1
    RegisteredBy string       // tag 2
}
```

The field's descriptions is as follows:

- `record` - the full [NameRecord](./02_state.md#namerecord) at registration time. For self-register, `verified=false`, `verified_by=""`, `source="self"`. For registrar-on-behalf, the registrar's chosen `verified`/`evidence_hash`/`source` are filled in and `verified_by = registrar address`.
- `registered_by` - the address that submitted the tx. For self-register this is the controller of `owner_did`; for registrar-on-behalf this is the registrar account.

This event drives both the state-mirror insert (via `record`) and the registration-history append (via `record` + `registered_by`).

## NameUpdatedEvent

Emitted whenever a `NameRecord` mutates in any way:

- on every [MsgUpdateNameByRegistrar](./03_messages.md#msgupdatenamebyregistrar) (single-emit),
- on every [MsgTransferName](./03_messages.md#msgtransfername) (dual-emit alongside `NameTransferredEvent`),
- on every [MsgSetNameStatus](./03_messages.md#msgsetnamestatus) (dual-emit alongside `NameStatusChangedEvent`).

```go
type NameUpdatedEvent struct {
    Record    *NameRecord  // tag 1
    UpdatedBy string       // tag 2
}
```

The field's descriptions is as follows:

- `record` - the full [NameRecord](./02_state.md#namerecord) **after** the mutation. Indexers can use this to upsert their state-mirror table without needing to know which message triggered the change.
- `updated_by` - the signer of the message (registrar address, owner-controlling address, or gov authority depending on the message).

## NameTransferredEvent

Emitted alongside `NameUpdatedEvent` on every successful [MsgTransferName](./03_messages.md#msgtransfername).

```go
type NameTransferredEvent struct {
    Namespace      string  // tag 1
    NormalizedName string  // tag 2
    FromOwnerDid   string  // tag 3
    ToOwnerDid     string  // tag 4
    TransferredBy  string  // tag 5
}
```

The field's descriptions is as follows:

- `namespace` - the namespace containing the transferred record.
- `normalized_name` - the canonical name that was transferred.
- `from_owner_did` - the DID that owned the record before the transfer.
- `to_owner_did` - the DID that owns it after the transfer.
- `transferred_by` - the signer. For owner-driven transfers this is an address controlling `from_owner_did`. For registrar-override transfers this is a registrar address — distinguishable from the previous case by the absence of a verification-method link between `transferred_by` and `from_owner_did`.

## NameStatusChangedEvent

Emitted alongside `NameUpdatedEvent` on every successful [MsgSetNameStatus](./03_messages.md#msgsetnamestatus).

```go
type NameStatusChangedEvent struct {
    Namespace      string      // tag 1
    NormalizedName string      // tag 2
    OldStatus      NameStatus  // tag 3
    NewStatus      NameStatus  // tag 4
    ChangedBy      string      // tag 5
    Reason         string      // tag 6
}
```

The field's descriptions is as follows:

- `namespace` - the namespace containing the record.
- `normalized_name` - the canonical name.
- `old_status` - the [NameStatus](./02_state.md#namestatus) immediately before the change.
- `new_status` - the [NameStatus](./02_state.md#namestatus) the record now holds.
- `changed_by` - the signer (registrar address or gov authority).
- `reason` - the free-form `reason` string from the message. Surfaced for moderation audit; may be empty.
