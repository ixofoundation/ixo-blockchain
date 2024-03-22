# Iid module specification

This document specifies the iid module, a custom Ixo Cosmos SDK module.

The IID (Interchain Identifier) Module establishes a decentralized identity mechanism, ensuring a standardized approach for all entities within the system. By harnessing the power of DIDs (Decentralized Identifiers) and IIDs, this module facilitates a robust, secure, and universally recognizable identity framework, paving the way for a seamless integration across various platforms and networks.

## Contents

1. **[Concepts](01_concepts.md)**

   - [Concepts](01_concepts.md#concepts)
     - [DID Documents](01_concepts.md#did-documents)
       - [Key Components of DID Documents:](01_concepts.md#key-components-of-did-documents)
       - [Advantages:](01_concepts.md#advantages)
   - [IID Module](01_concepts.md#iid-module)
     - [Context](01_concepts.md#context)
     - [Concepts](01_concepts.md#concepts-1)
       - [DIDs and IIDs](01_concepts.md#dids-and-iids)
       - [IID Document](01_concepts.md#iid-document)
       - [IID Registry](01_concepts.md#iid-registry)
       - [IID Method](01_concepts.md#iid-method)
       - [IID Resolver](01_concepts.md#iid-resolver)
       - [IID Deactivation](01_concepts.md#iid-deactivation)

2. **[State](02_state.md)**

   - [State](02_state.md#state)
     - [Iids](02_state.md#iids)
   - [Types](02_state.md#types)
     - [IidDocument](02_state.md#iiddocument)
     - [Verification Method](02_state.md#verification-method)
     - [Service](02_state.md#service)
     - [Accorded Right](02_state.md#accorded-right)
     - [Linked Resource](02_state.md#linked-resource)
     - [Linked Entity](02_state.md#linked-entity)
     - [Linked Claim](02_state.md#linked-claim)
     - [Context](02_state.md#context)
     - [IidMetadata](02_state.md#iidmetadata)
     - [Verification](02_state.md#verification)

3. **[Messages](03_messages.md)**

   - [Messages](03_messages.md#messages)
     - [MsgCreateIidDocument](03_messages.md#msgcreateiiddocument)
     - [MsgUpdateIidDocument](03_messages.md#msgupdateiiddocument)
     - [MsgAddVerification](03_messages.md#msgaddverification)
     - [MsgSetVerificationRelationships](03_messages.md#msgsetverificationrelationships)
     - [MsgRevokeVerification](03_messages.md#msgrevokeverification)
     - [MsgAddService](03_messages.md#msgaddservice)
     - [MsgDeleteService](03_messages.md#msgdeleteservice)
     - [MsgAddController](03_messages.md#msgaddcontroller)
     - [MsgDeleteController](03_messages.md#msgdeletecontroller)
     - [MsgAddLinkedResource](03_messages.md#msgaddlinkedresource)
     - [MsgDeleteLinkedResource](03_messages.md#msgdeletelinkedresource)
     - [MsgAddLinkedClaim](03_messages.md#msgaddlinkedclaim)
     - [MsgDeleteLinkedClaim](03_messages.md#msgdeletelinkedclaim)
     - [MsgAddLinkedEntity](03_messages.md#msgaddlinkedentity)
     - [MsgDeleteLinkedEntity](03_messages.md#msgdeletelinkedentity)
     - [MsgAddAccordedRight](03_messages.md#msgaddaccordedright)
     - [MsgDeleteAccordedRight](03_messages.md#msgdeleteaccordedright)
     - [MsgAddIidContext](03_messages.md#msgaddiidcontext)
     - [MsgDeleteIidContext](03_messages.md#msgdeleteiidcontext)
     - [MsgDeactivateIID](03_messages.md#msgdeactivateiid)

4. **[Events](04_events.md)**

   - [Events](04_events.md#events)
     - [IidDocumentCreatedEvent](04_events.md#iiddocumentcreatedevent)
     - [IidDocumentUpdatedEvent](04_events.md#iiddocumentupdatedevent)

5. **[Parameters](05_params.md)**

6. **[Future Improvements](06_future_improvements.md)**
