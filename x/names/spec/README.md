# Names module specification

This document specifies the names module, a custom Ixo Cosmos SDK module.

The names module is a chain-level **name service** that maps human-readable handles to DIDs (e.g. `yoid:alice` → `did:ixo:…`). Names live in **governance-managed namespaces** that each define their own validation rules, registrar set, and self-register policy. Records bind a normalised name to a DID, are never hard-deleted (status transitions are used instead so audit history is preserved), and support both user self-registration and registrar-on-behalf flows for custodial / USSD / oracle-attested use cases.

The module was introduced in chain upgrade `v7` (IXO-1123) alongside the existing `iid` module, which it does not replace — `iid.alsoKnownAs` is still the W3C-style soft hint, while the names module is the canonical, indexed, governance-controlled handle registry.

## Contents

1. **[Concepts](01_concepts.md)**

   - [Names and namespaces](01_concepts.md#names-and-namespaces)
   - [Normalisation and uniqueness](01_concepts.md#normalisation-and-uniqueness)
   - [Self-register vs registrar-only namespaces](01_concepts.md#self-register-vs-registrar-only-namespaces)
   - [Registrars and on-behalf registration](01_concepts.md#registrars-and-on-behalf-registration)
   - [Verified records and off-chain attestation](01_concepts.md#verified-records-and-off-chain-attestation)
   - [Status lifecycle (no hard-delete)](01_concepts.md#status-lifecycle-no-hard-delete)
   - [Transfers and registrar override](01_concepts.md#transfers-and-registrar-override)
   - [Reverse lookup by owner DID](01_concepts.md#reverse-lookup-by-owner-did)
   - [Relationship to the iid module](01_concepts.md#relationship-to-the-iid-module)
   - [Reserved fields for future expansion](01_concepts.md#reserved-fields-for-future-expansion)

2. **[State](02_state.md)**

   - [Storage layout](02_state.md#storage-layout)
   - [Types](02_state.md#types)
     - [Namespace](02_state.md#namespace)
     - [NameRecord](02_state.md#namerecord)
     - [NameStatus enum](02_state.md#namestatus)

3. **[Messages](03_messages.md)**

   - [Governance](03_messages.md#governance-operations)
     - [MsgCreateNamespace](03_messages.md#msgcreatenamespace)
     - [MsgUpdateNamespace](03_messages.md#msgupdatenamespace)
   - [User](03_messages.md#user-operations)
     - [MsgRegisterName](03_messages.md#msgregistername)
     - [MsgTransferName](03_messages.md#msgtransfername)
   - [Registrar](03_messages.md#registrar-operations)
     - [MsgRegisterNameByRegistrar](03_messages.md#msgregisternamebyregistrar)
     - [MsgUpdateNameByRegistrar](03_messages.md#msgupdatenamebyregistrar)
     - [MsgSetNameStatus](03_messages.md#msgsetnamestatus)

4. **[Events](04_events.md)**

   - [Dual-emission pattern](04_events.md#dual-emission-pattern)
   - [NamespaceCreatedEvent](04_events.md#namespacecreatedevent)
   - [NamespaceUpdatedEvent](04_events.md#namespaceupdatedevent)
   - [NameRegisteredEvent](04_events.md#nameregisteredevent)
   - [NameUpdatedEvent](04_events.md#nameupdatedevent)
   - [NameTransferredEvent](04_events.md#nametransferredevent)
   - [NameStatusChangedEvent](04_events.md#namestatuschangedevent)

5. **[Parameters](05_params.md)**

   - [Module-wide parameters](05_params.md#module-wide-parameters)
   - [Per-namespace fields](05_params.md#per-namespace-fields)
   - [Module-wide constants](05_params.md#module-wide-constants)
