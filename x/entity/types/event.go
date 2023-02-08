package types

// Constructs a new entity_created sdk.Event
func NewEntityCreatedEvent(id string, owner string) *EntityCreatedEvent {
	return &EntityCreatedEvent{
		Id:    id,
		Owner: owner,
	}
}

// Constructs a new entity_updated sdk.Event
func NewEntityUpdatedEvent(id string, signer string) *EntityUpdatedEvent {
	return &EntityUpdatedEvent{
		Id:    id,
		Signer: signer,
	}
}

// Constructs a new entity_updated_verified sdk.Event
func NewEntityVerifiedUpdatedEvent(id string, signer string, entity_verified bool) *EntityVerifiedUpdatedEvent {
	return &EntityVerifiedUpdatedEvent{
		Id:    id,
		Signer: signer,
		EntityVerified: entity_verified,
	}
}

// Constructs a new entity_transferred sdk.Event
func NewEntityTransferredEvent(id string, owner string) *EntityTransferredEvent {
	return &EntityTransferredEvent{
		Id:    id,
		Owner: owner,
	}
}