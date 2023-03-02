package types

// Constructs a new entity_created sdk.Event
func NewEntityCreatedEvent(entity *Entity, owner string) *EntityCreatedEvent {
	return &EntityCreatedEvent{
		Entity: entity,
		Owner:  owner,
	}
}

// Constructs a new entity_updated sdk.Event
func NewEntityUpdatedEvent(entity *Entity, owner string) *EntityUpdatedEvent {
	return &EntityUpdatedEvent{
		Entity: entity,
		Owner:  owner,
	}
}

// Constructs a new entity_updated_verified sdk.Event
func NewEntityVerifiedUpdatedEvent(id, owner string, entity_verified bool) *EntityVerifiedUpdatedEvent {
	return &EntityVerifiedUpdatedEvent{
		Id:             id,
		Owner:          owner,
		EntityVerified: entity_verified,
	}
}

// Constructs a new entity_transferred sdk.Event
func NewEntityTransferredEvent(id, from, to string) *EntityTransferredEvent {
	return &EntityTransferredEvent{
		Id:   id,
		From: from,
		To:   to,
	}
}
