package types

// NewDidDocumentCreatedEvent constructs a new did_created sdk.Event
func NewIidDocumentCreatedEvent(did, owner string) *IidDocumentCreatedEvent {
	return &IidDocumentCreatedEvent{
		Did:    did,
		Signer: owner,
	}
}

// NewDidDocumentUpdatedEvent constructs a new did_created sdk.Event
func NewIidDocumentUpdatedEvent(did, signer string) *IidDocumentUpdatedEvent {
	return &IidDocumentUpdatedEvent{
		Did:    did,
		Signer: signer,
	}
}
