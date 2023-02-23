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
func NewEntityVerifiedUpdatedEvent(entity *Entity, owner string, entity_verified bool) *EntityVerifiedUpdatedEvent {
	return &EntityVerifiedUpdatedEvent{
		Entity:         entity,
		Owner:          owner,
		EntityVerified: entity_verified,
	}
}

// Constructs a new entity_transferred sdk.Event
func NewEntityTransferredEvent(entity *Entity, owner string) *EntityTransferredEvent {
	return &EntityTransferredEvent{
		Entity: entity,
		Owner:  owner,
	}
}
