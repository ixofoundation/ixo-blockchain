package types

// NewIidDocumentCreatedEvent constructs a new did_created sdk.Event
func NewIidDocumentCreatedEvent(iidDocument *IidDocument) *IidDocumentCreatedEvent {
	return &IidDocumentCreatedEvent{
		IidDocument: iidDocument,
	}
}

// NewIidDocumentUpdatedEvent constructs a new did_updated sdk.Event
func NewIidDocumentUpdatedEvent(did, signer string) *IidDocumentUpdatedEvent {
	return &IidDocumentUpdatedEvent{
		Did:    did,
		Signer: signer,
	}
}
