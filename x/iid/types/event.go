package types

// NewIidDocumentCreatedEvent constructs a new did_created sdk.Event
func NewIidDocumentCreatedEvent(did, owner string) *IidDocumentCreatedEvent {
	return &IidDocumentCreatedEvent{
		Did:    did,
		Signer: owner,
	}
}

// NewIidDocumentUpdatedEvent constructs a new did_updated sdk.Event
func NewIidDocumentUpdatedEvent(did, signer string) *IidDocumentUpdatedEvent {
	return &IidDocumentUpdatedEvent{
		Did:    did,
		Signer: signer,
	}
}
