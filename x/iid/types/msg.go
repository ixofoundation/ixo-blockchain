package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// msg types
const (
	TypeMsgCreateDidDocument = "create-did"
)

var _ sdk.Msg = &MsgCreateIidDocument{}

// NewMsgCreateDidDocument creates a new MsgCreateDidDocument instance
func NewMsgCreateIidDocument(
	id string,
	verifications []*Verification,
	services []*Service,
	rights []*AccordedRight,
	resources []*LinkedResource,
	entity []*LinkedEntity,
	signerAccount string,
	didContexts []*Context,
) *MsgCreateIidDocument {
	return &MsgCreateIidDocument{
		Id:             id,
		Verifications:  verifications,
		Services:       services,
		AccordedRight:  rights,
		LinkedResource: resources,
		LinkedEntity:   entity,
		Context:        didContexts,
		Signer:         signerAccount,
	}
}

// Route implements sdk.Msg
func (MsgCreateIidDocument) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (MsgCreateIidDocument) Type() string {
	return TypeMsgCreateDidDocument
}

func (msg MsgCreateIidDocument) GetSignBytes() []byte {
	panic("IBC messages do not support amino")
}

// GetSigners implements sdk.Msg
func (msg MsgCreateIidDocument) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

// --------------------------
// UPDATE IDENTIFIER
// --------------------------

// msg types
const (
	TypeMsgUpdateDidDocument = "update-did"
)

func NewMsgUpdateDidDocument(
	didDoc *IidDocument,
	signerAccount string,
) *MsgUpdateIidDocument {
	return &MsgUpdateIidDocument{
		Doc:    didDoc,
		Signer: signerAccount,
	}
}

// Route implements sdk.Msg
func (MsgUpdateIidDocument) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (MsgUpdateIidDocument) Type() string {
	return TypeMsgUpdateDidDocument
}

func (msg MsgUpdateIidDocument) GetSignBytes() []byte {
	panic("IBC messages do not support amino")
}

// GetSigners implements sdk.Msg
func (msg MsgUpdateIidDocument) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

// --------------------------
// ADD VERIFICATION
// --------------------------
// msg types
const (
	TypeMsgAddVerification = "add-verification"
)

var _ sdk.Msg = &MsgAddVerification{}

// NewMsgAddVerification creates a new MsgAddVerification instance
func NewMsgAddVerification(
	id string,
	verification *Verification,
	signerAccount string,
) *MsgAddVerification {
	return &MsgAddVerification{
		Id:           id,
		Verification: verification,
		Signer:       signerAccount,
	}
}

// Route implements sdk.Msg
func (MsgAddVerification) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (MsgAddVerification) Type() string {
	return TypeMsgAddVerification
}

func (msg MsgAddVerification) GetSignBytes() []byte {
	panic("IBC messages do not support amino")
}

// GetSigners implements sdk.Msg
func (msg MsgAddVerification) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

// --------------------------
// REVOKE VERIFICATION
// --------------------------

// msg types
const (
	TypeMsgRevokeVerification = "revoke-verification"
)

var _ sdk.Msg = &MsgRevokeVerification{}

// NewMsgRevokeVerification creates a new MsgRevokeVerification instance
func NewMsgRevokeVerification(
	id string,
	methodID string,
	signerAccount string,
) *MsgRevokeVerification {
	return &MsgRevokeVerification{
		Id:       id,
		MethodId: methodID,
		Signer:   signerAccount,
	}
}

// Route implements sdk.Msg
func (MsgRevokeVerification) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (MsgRevokeVerification) Type() string {
	return TypeMsgRevokeVerification
}

func (msg MsgRevokeVerification) GetSignBytes() []byte {
	panic("IBC messages do not support amino")
}

// GetSigners implements sdk.Msg
func (msg MsgRevokeVerification) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

// --------------------------
// SET VERIFICATION RELATIONSHIPS
// --------------------------
// msg types
const (
	TypeMsgSetVerificationRelationships = "set-verification-relationships"
)

func NewMsgSetVerificationRelationships(
	id string,
	methodID string,
	relationships []string,
	signerAccount string,
) *MsgSetVerificationRelationships {
	return &MsgSetVerificationRelationships{
		Id:            id,
		MethodId:      methodID,
		Relationships: relationships,
		Signer:        signerAccount,
	}
}

// Route implements sdk.Msg
func (MsgSetVerificationRelationships) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (MsgSetVerificationRelationships) Type() string {
	return TypeMsgSetVerificationRelationships
}

func (msg MsgSetVerificationRelationships) GetSignBytes() []byte {
	panic("IBC messages do not support amino")
}

// GetSigners implements sdk.Msg
func (msg MsgSetVerificationRelationships) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

// --------------------------
// ADD SERVICE
// --------------------------

// msg types
const (
	TypeMsgAddService = "add-service"
)

var _ sdk.Msg = &MsgAddService{}

// NewMsgAddService creates a new MsgAddService instance
func NewMsgAddService(
	id string,
	service *Service,
	signerAccount string,
) *MsgAddService {
	return &MsgAddService{
		Id:          id,
		ServiceData: service,
		Signer:      signerAccount,
	}
}
func NewMsgAddLinkedResource(
	id string,
	linkedResource *LinkedResource,
	signerAccount string,
) *MsgAddLinkedResource {
	return &MsgAddLinkedResource{
		Id:             id,
		LinkedResource: linkedResource,
		Signer:         signerAccount,
	}
}
func NewMsgAddLinkedEntity(
	id string,
	linkedResource *LinkedEntity,
	signerAccount string,
) *MsgAddLinkedEntity {
	return &MsgAddLinkedEntity{
		Id:           id,
		LinkedEntity: linkedResource,
		Signer:       signerAccount,
	}
}

// Route implements sdk.Msg
func (MsgAddService) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (MsgAddService) Type() string {
	return TypeMsgAddService
}

func (msg MsgAddService) GetSignBytes() []byte {
	panic("IBC messages do not support amino")
}

// GetSigners implements sdk.Msg
func (msg MsgAddService) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

// Route implements sdk.Msg
func (MsgAddLinkedResource) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (MsgAddLinkedResource) Type() string {
	return TypeMsgAddService
}

func (msg MsgAddLinkedResource) GetSignBytes() []byte {
	panic("IBC messages do not support amino")
}

// GetSigners implements sdk.Msg
func (msg MsgAddLinkedResource) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}
func (msg MsgAddLinkedEntity) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

// --------------------------
// DELETE SERVICE
// --------------------------

// msg types
const (
	TypeMsgDeleteService = "delete-service"
)

func NewMsgDeleteService(
	id string,
	serviceID string,
	signerAccount string,
) *MsgDeleteService {
	return &MsgDeleteService{
		Id:        id,
		ServiceId: serviceID,
		Signer:    signerAccount,
	}
}
func NewMsgDeleteLinkedResource(
	id string,
	resourceID string,
	signerAccount string,
) *MsgDeleteLinkedResource {
	return &MsgDeleteLinkedResource{
		Id:         id,
		ResourceId: resourceID,
		Signer:     signerAccount,
	}
}
func NewMsgDeleteLinkedEntity(
	id string,
	entityID string,
	signerAccount string,
) *MsgDeleteLinkedEntity {
	return &MsgDeleteLinkedEntity{
		Id:       id,
		EntityId: entityID,
		Signer:   signerAccount,
	}
}

// Route implements sdk.Msg
func (MsgDeleteService) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (MsgDeleteService) Type() string {
	return TypeMsgDeleteService
}

func (msg MsgDeleteService) GetSignBytes() []byte {
	panic("IBC messages do not support amino")
}

// GetSigners implements sdk.Msg
func (msg MsgDeleteService) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

func (MsgDeleteLinkedResource) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (MsgDeleteLinkedResource) Type() string {
	return TypeMsgDeleteService
}

func (msg MsgDeleteLinkedResource) GetSignBytes() []byte {
	panic("IBC messages do not support amino")
}

// GetSigners implements sdk.Msg
func (msg MsgDeleteLinkedResource) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}
func (msg MsgDeleteLinkedEntity) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

// --------------------------
// ADD RIGHT
// --------------------------

// msg types
const (
	TypeMsgAddRight = "add-right"
)

var _ sdk.Msg = &MsgAddAccordedRight{}

// NewMsgAddAccordedRight creates a new MsgAddAccordedright instance
func NewMsgAddAccordedRight(
	id string,
	right *AccordedRight,
	signerAccount string,
) *MsgAddAccordedRight {
	return &MsgAddAccordedRight{
		Id:            id,
		AccordedRight: right,
		Signer:        signerAccount,
	}
}

// Route implements sdk.Msg
func (MsgAddAccordedRight) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (MsgAddAccordedRight) Type() string {
	return TypeMsgAddService
}

func (msg MsgAddAccordedRight) GetSignBytes() []byte {
	panic("IBC messages do not support amino")
}

// GetSigners implements sdk.Msg
func (msg MsgAddAccordedRight) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

// --------------------------
// DELETE RIGHT
// --------------------------

// msg types
const (
	TypeMsgDeleteAccordedRight = "delete-right"
)

func NewMsgDeleteAccordedRight(
	id string,
	rightID string,
	signerAccount string,
) *MsgDeleteAccordedRight {
	return &MsgDeleteAccordedRight{
		Id:      id,
		RightId: rightID,
		Signer:  signerAccount,
	}
}

// Route implements sdk.Msg
func (MsgDeleteAccordedRight) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (MsgDeleteAccordedRight) Type() string {
	return TypeMsgDeleteService
}

func (msg MsgDeleteAccordedRight) GetSignBytes() []byte {
	panic("IBC messages do not support amino")
}

// GetSigners implements sdk.Msg
func (msg MsgDeleteAccordedRight) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

// --------------------------
// ADD CONTROLLER
// --------------------------

// msg types
const (
	TypeMsgAddController = "add-controller"
)

func NewMsgAddController(
	id string,
	controllerDID string,
	signerAccount string,
) *MsgAddController {
	return &MsgAddController{
		Id:            id,
		ControllerDid: controllerDID,
		Signer:        signerAccount,
	}
}

// Route implements sdk.Msg
func (MsgAddController) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (MsgAddController) Type() string {
	return TypeMsgAddController
}

func (msg MsgAddController) GetSignBytes() []byte {
	panic("IBC messages do not support amino")
}

// GetSigners implements sdk.Msg
func (msg MsgAddController) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

// --------------------------
// DELETE CONTROLLER
// --------------------------

// msg types
const (
	TypeMsgDeleteController = "delete-controller"
)

func NewMsgDeleteController(
	id string,
	controllerDID string,
	signerAccount string,
) *MsgDeleteController {
	return &MsgDeleteController{
		Id:            id,
		ControllerDid: controllerDID,
		Signer:        signerAccount,
	}
}

// Route implements sdk.Msg
func (MsgDeleteController) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (MsgDeleteController) Type() string {
	return TypeMsgDeleteController
}

func (msg MsgDeleteController) GetSignBytes() []byte {
	panic("IBC messages do not support amino")
}

// GetSigners implements sdk.Msg
func (msg MsgDeleteController) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

// --------------------------
// ADD CONTEXT
// --------------------------

// msg types
const (
	TypeMsgAddContext = "add-did-context"
)

var _ sdk.Msg = &MsgAddIidContext{}

// NewMsgAddService creates a new MsgAddService instance
func NewMsgAddDidContext(
	id string,
	context *Context,
	signerAccount string,
) *MsgAddIidContext {
	return &MsgAddIidContext{
		Id:      id,
		Context: context,
		Signer:  signerAccount,
	}
}

// Route implements sdk.Msg
func (MsgAddIidContext) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (MsgAddIidContext) Type() string {
	return TypeMsgAddService
}

func (msg MsgAddIidContext) GetSignBytes() []byte {
	panic("IBC messages do not support amino")
}

// GetSigners implements sdk.Msg
func (msg MsgAddIidContext) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

// --------------------------
// DELETE CONTEXT
// --------------------------

// msg types
const (
	TypeMsgDeleteDidContext = "delete-context"
)

func NewMsgDeleteDidContext(
	id string,
	key string,
	signerAccount string,
) *MsgDeleteIidContext {
	return &MsgDeleteIidContext{
		Id:         id,
		ContextKey: key,
		Signer:     signerAccount,
	}
}

// Route implements sdk.Msg
func (MsgDeleteIidContext) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (MsgDeleteIidContext) Type() string {
	return TypeMsgDeleteService
}

func (msg MsgDeleteIidContext) GetSignBytes() []byte {
	panic("IBC messages do not support amino")
}

// GetSigners implements sdk.Msg
func (msg MsgDeleteIidContext) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

// --------------------------
// ADD META
// --------------------------

// msg types
const (
	TypeMsgAddDidMeta = "add-did-meta"
)

var _ sdk.Msg = &MsgUpdateIidMeta{}

// NewMsgAddService creates a new MsgAddService instance
func NewMsgUpdateDidMetaData(
	id string,
	meta *IidMetadata,
	signerAccount string,
) *MsgUpdateIidMeta {
	return &MsgUpdateIidMeta{
		Id:     id,
		Meta:   meta,
		Signer: signerAccount,
	}
}

// Type implements sdk.Msg
func (MsgUpdateIidMeta) Type() string {
	return TypeMsgAddDidMeta
}

func (msg MsgUpdateIidMeta) GetSignBytes() []byte {
	panic("IBC messages do not support amino")
}

// GetSigners implements sdk.Msg
func (msg MsgUpdateIidMeta) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}

func NewMsgDeactivateIID(
	id string,
	state bool,
	signerAccount string,
) *MsgDeactivateIID {
	return &MsgDeactivateIID{
		Id:     id,
		State:  state,
		Signer: signerAccount,
	}
}

// Route implements sdk.Msg
func (MsgDeactivateIID) Route() string {
	return RouterKey
}

func (MsgDeactivateIID) Type() string {
	return TypeMsgAddDidMeta
}

func (msg MsgDeactivateIID) GetSignBytes() []byte {
	panic("IBC messages do not support amino")
}

// GetSigners implements sdk.Msg
func (msg MsgDeactivateIID) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(msg.Signer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{accAddr}
}
