package types

func NewNamespaceCreatedEvent(ns Namespace, authority string) *NamespaceCreatedEvent {
	return &NamespaceCreatedEvent{Namespace: &ns, Authority: authority}
}

func NewNamespaceUpdatedEvent(ns Namespace, authority string) *NamespaceUpdatedEvent {
	return &NamespaceUpdatedEvent{Namespace: &ns, Authority: authority}
}

func NewNameRegisteredEvent(record NameRecord, registeredBy string) *NameRegisteredEvent {
	return &NameRegisteredEvent{Record: &record, RegisteredBy: registeredBy}
}

func NewNameUpdatedEvent(record NameRecord, updatedBy string) *NameUpdatedEvent {
	return &NameUpdatedEvent{Record: &record, UpdatedBy: updatedBy}
}

func NewNameTransferredEvent(namespace, normalizedName, fromDid, toDid, transferredBy string) *NameTransferredEvent {
	return &NameTransferredEvent{
		Namespace:      namespace,
		NormalizedName: normalizedName,
		FromOwnerDid:   fromDid,
		ToOwnerDid:     toDid,
		TransferredBy:  transferredBy,
	}
}

func NewNameStatusChangedEvent(namespace, normalizedName string, oldStatus, newStatus NameStatus, changedBy, reason string) *NameStatusChangedEvent {
	return &NameStatusChangedEvent{
		Namespace:      namespace,
		NormalizedName: normalizedName,
		OldStatus:      oldStatus,
		NewStatus:      newStatus,
		ChangedBy:      changedBy,
		Reason:         reason,
	}
}
