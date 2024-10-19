//nolint:all
package v2claims

import (
	fmt "fmt"
	io "io"
	math_bits "math/bits"
	time "time"

	"cosmossdk.io/math"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	github_com_gogo_protobuf_types "github.com/gogo/protobuf/types"
	_ "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
)

type CollectionState int32

const (
	CollectionState_open   CollectionState = 0
	CollectionState_paused CollectionState = 1
	CollectionState_closed CollectionState = 2
)

var CollectionState_name = map[int32]string{
	0: "OPEN",
	1: "PAUSED",
	2: "CLOSED",
}

var CollectionState_value = map[string]int32{
	"OPEN":   0,
	"PAUSED": 1,
	"CLOSED": 2,
}

func (x CollectionState) String() string {
	return proto.EnumName(CollectionState_name, int32(x))
}

func (CollectionState) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_619c1a0876cd0592, []int{0}
}

type EvaluationStatus int32

const (
	EvaluationStatus_pending  EvaluationStatus = 0
	EvaluationStatus_approved EvaluationStatus = 1
	EvaluationStatus_rejected EvaluationStatus = 2
	EvaluationStatus_disputed EvaluationStatus = 3
)

var EvaluationStatus_name = map[int32]string{
	0: "PENDING",
	1: "APPROVED",
	2: "REJECTED",
	3: "DISPUTED",
}

var EvaluationStatus_value = map[string]int32{
	"PENDING":  0,
	"APPROVED": 1,
	"REJECTED": 2,
	"DISPUTED": 3,
}

func (x EvaluationStatus) String() string {
	return proto.EnumName(EvaluationStatus_name, int32(x))
}

func (EvaluationStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_619c1a0876cd0592, []int{1}
}

type PaymentType int32

const (
	PaymentType_submission PaymentType = 0
	PaymentType_approval   PaymentType = 1
	PaymentType_evaluation PaymentType = 2
	PaymentType_rejection  PaymentType = 3
)

var PaymentType_name = map[int32]string{
	0: "SUBMISSION",
	1: "APPROVAL",
	2: "EVALUATION",
	3: "REJECTION",
}

var PaymentType_value = map[string]int32{
	"SUBMISSION": 0,
	"APPROVAL":   1,
	"EVALUATION": 2,
	"REJECTION":  3,
}

func (x PaymentType) String() string {
	return proto.EnumName(PaymentType_name, int32(x))
}

func (PaymentType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_619c1a0876cd0592, []int{2}
}

type PaymentStatus int32

const (
	PaymentStatus_no_payment PaymentStatus = 0
	PaymentStatus_promised   PaymentStatus = 1
	PaymentStatus_authorized PaymentStatus = 2
	PaymentStatus_gauranteed PaymentStatus = 3
	PaymentStatus_paid       PaymentStatus = 4
	PaymentStatus_failed     PaymentStatus = 5
	PaymentStatus_disputed   PaymentStatus = 6
)

var PaymentStatus_name = map[int32]string{
	0: "NO_PAYMENT",
	1: "PROMISED",
	2: "AUTHORIZED",
	3: "GAURANTEED",
	4: "PAID",
	5: "FAILED",
	6: "DISPUTED",
}

var PaymentStatus_value = map[string]int32{
	"NO_PAYMENT": 0,
	"PROMISED":   1,
	"AUTHORIZED": 2,
	"GAURANTEED": 3,
	"PAID":       4,
	"FAILED":     5,
	"DISPUTED":   6,
}

func (x PaymentStatus) String() string {
	return proto.EnumName(PaymentStatus_name, int32(x))
}

func (PaymentStatus) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_619c1a0876cd0592, []int{3}
}

type Params struct {
	CollectionSequence   uint64         `protobuf:"varint,1,opt,name=collection_sequence,json=collectionSequence,proto3" json:"collection_sequence,omitempty"`
	IxoAccount           string         `protobuf:"bytes,2,opt,name=ixo_account,json=ixoAccount,proto3" json:"ixo_account,omitempty"`
	NetworkFeePercentage math.LegacyDec `protobuf:"bytes,3,opt,name=network_fee_percentage,json=networkFeePercentage,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"network_fee_percentage"`
	NodeFeePercentage    math.LegacyDec `protobuf:"bytes,4,opt,name=node_fee_percentage,json=nodeFeePercentage,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"node_fee_percentage"`
}

func (m *Params) Reset()         { *m = Params{} }
func (m *Params) String() string { return proto.CompactTextString(m) }
func (*Params) ProtoMessage()    {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_619c1a0876cd0592, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetCollectionSequence() uint64 {
	if m != nil {
		return m.CollectionSequence
	}
	return 0
}

func (m *Params) GetIxoAccount() string {
	if m != nil {
		return m.IxoAccount
	}
	return ""
}

type Collection struct {
	// collection id is the incremented internal id for the collection of claims
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// entity is the DID of the entity for which the claims are being created
	Entity string `protobuf:"bytes,2,opt,name=entity,proto3" json:"entity,omitempty"`
	// admin is the account address that will authorize or revoke agents and
	// payments (the grantor)
	Admin string `protobuf:"bytes,3,opt,name=admin,proto3" json:"admin,omitempty"`
	// protocol is the DID of the claim protocol
	Protocol string `protobuf:"bytes,4,opt,name=protocol,proto3" json:"protocol,omitempty"`
	// startDate is the date after which claims may be submitted
	StartDate *time.Time `protobuf:"bytes,5,opt,name=start_date,json=startDate,proto3,stdtime" json:"start_date,omitempty"`
	// endDate is the date after which no more claims may be submitted (no endDate
	// is allowed)
	EndDate *time.Time `protobuf:"bytes,6,opt,name=end_date,json=endDate,proto3,stdtime" json:"end_date,omitempty"`
	// quota is the maximum number of claims that may be submitted, 0 is unlimited
	Quota uint64 `protobuf:"varint,7,opt,name=quota,proto3" json:"quota,omitempty"`
	// count is the number of claims already submitted (internally calculated)
	Count uint64 `protobuf:"varint,8,opt,name=count,proto3" json:"count,omitempty"`
	// evaluated is the number of claims that have been evaluated (internally
	// calculated)
	Evaluated uint64 `protobuf:"varint,9,opt,name=evaluated,proto3" json:"evaluated,omitempty"`
	// approved is the number of claims that have been evaluated and approved
	// (internally calculated)
	Approved uint64 `protobuf:"varint,10,opt,name=approved,proto3" json:"approved,omitempty"`
	// rejected is the number of claims that have been evaluated and rejected
	// (internally calculated)
	Rejected uint64 `protobuf:"varint,11,opt,name=rejected,proto3" json:"rejected,omitempty"`
	// disputed is the number of claims that have disputed status (internally
	// calculated)
	Disputed uint64 `protobuf:"varint,12,opt,name=disputed,proto3" json:"disputed,omitempty"`
	// state is the current state of this Collection (open, paused, closed)
	State CollectionState `protobuf:"varint,13,opt,name=state,proto3,enum=ixo.claims.v1beta1.CollectionState" json:"state,omitempty"`
	// payments is the amount paid for claim submission, evaluation, approval, or
	// rejection
	Payments *Payments `protobuf:"bytes,14,opt,name=payments,proto3" json:"payments,omitempty"`
	// signer address
	Signer string `protobuf:"bytes,15,opt,name=signer,proto3" json:"signer,omitempty"`
}

func (m *Collection) Reset()         { *m = Collection{} }
func (m *Collection) String() string { return proto.CompactTextString(m) }
func (*Collection) ProtoMessage()    {}
func (*Collection) Descriptor() ([]byte, []int) {
	return fileDescriptor_619c1a0876cd0592, []int{1}
}
func (m *Collection) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Collection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Collection.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Collection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Collection.Merge(m, src)
}
func (m *Collection) XXX_Size() int {
	return m.Size()
}
func (m *Collection) XXX_DiscardUnknown() {
	xxx_messageInfo_Collection.DiscardUnknown(m)
}

var xxx_messageInfo_Collection proto.InternalMessageInfo

func (m *Collection) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Collection) GetEntity() string {
	if m != nil {
		return m.Entity
	}
	return ""
}

func (m *Collection) GetAdmin() string {
	if m != nil {
		return m.Admin
	}
	return ""
}

func (m *Collection) GetProtocol() string {
	if m != nil {
		return m.Protocol
	}
	return ""
}

func (m *Collection) GetStartDate() *time.Time {
	if m != nil {
		return m.StartDate
	}
	return nil
}

func (m *Collection) GetEndDate() *time.Time {
	if m != nil {
		return m.EndDate
	}
	return nil
}

func (m *Collection) GetQuota() uint64 {
	if m != nil {
		return m.Quota
	}
	return 0
}

func (m *Collection) GetCount() uint64 {
	if m != nil {
		return m.Count
	}
	return 0
}

func (m *Collection) GetEvaluated() uint64 {
	if m != nil {
		return m.Evaluated
	}
	return 0
}

func (m *Collection) GetApproved() uint64 {
	if m != nil {
		return m.Approved
	}
	return 0
}

func (m *Collection) GetRejected() uint64 {
	if m != nil {
		return m.Rejected
	}
	return 0
}

func (m *Collection) GetDisputed() uint64 {
	if m != nil {
		return m.Disputed
	}
	return 0
}

func (m *Collection) GetState() CollectionState {
	if m != nil {
		return m.State
	}
	return CollectionState_open
}

func (m *Collection) GetPayments() *Payments {
	if m != nil {
		return m.Payments
	}
	return nil
}

func (m *Collection) GetSigner() string {
	if m != nil {
		return m.Signer
	}
	return ""
}

type Payments struct {
	Submission *Payment `protobuf:"bytes,1,opt,name=submission,proto3" json:"submission,omitempty"`
	Evaluation *Payment `protobuf:"bytes,2,opt,name=evaluation,proto3" json:"evaluation,omitempty"`
	Approval   *Payment `protobuf:"bytes,3,opt,name=approval,proto3" json:"approval,omitempty"`
	Rejection  *Payment `protobuf:"bytes,4,opt,name=rejection,proto3" json:"rejection,omitempty"`
}

func (m *Payments) Reset()         { *m = Payments{} }
func (m *Payments) String() string { return proto.CompactTextString(m) }
func (*Payments) ProtoMessage()    {}
func (*Payments) Descriptor() ([]byte, []int) {
	return fileDescriptor_619c1a0876cd0592, []int{2}
}
func (m *Payments) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Payments) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Payments.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Payments) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Payments.Merge(m, src)
}
func (m *Payments) XXX_Size() int {
	return m.Size()
}
func (m *Payments) XXX_DiscardUnknown() {
	xxx_messageInfo_Payments.DiscardUnknown(m)
}

var xxx_messageInfo_Payments proto.InternalMessageInfo

func (m *Payments) GetSubmission() *Payment {
	if m != nil {
		return m.Submission
	}
	return nil
}

func (m *Payments) GetEvaluation() *Payment {
	if m != nil {
		return m.Evaluation
	}
	return nil
}

func (m *Payments) GetApproval() *Payment {
	if m != nil {
		return m.Approval
	}
	return nil
}

func (m *Payments) GetRejection() *Payment {
	if m != nil {
		return m.Rejection
	}
	return nil
}

type Payment struct {
	// account is the entity account address from which the payment will be made
	Account string      `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	Amount  types.Coins `protobuf:"bytes,2,rep,name=amount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"amount"`
	// if empty(nil) then no contract payment, not allowed for Evaluation Payment
	Contract_1155Payment *Contract1155Payment `protobuf:"bytes,3,opt,name=contract_1155_payment,json=contract1155Payment,proto3" json:"contract_1155_payment,omitempty"`
	// timeout after claim/evaluation to create authZ for payment, if 0 then
	// immidiate direct payment
	TimeoutNs time.Duration `protobuf:"bytes,4,opt,name=timeout_ns,json=timeoutNs,proto3,stdduration" json:"timeout_ns"`
}

func (m *Payment) Reset()         { *m = Payment{} }
func (m *Payment) String() string { return proto.CompactTextString(m) }
func (*Payment) ProtoMessage()    {}
func (*Payment) Descriptor() ([]byte, []int) {
	return fileDescriptor_619c1a0876cd0592, []int{3}
}
func (m *Payment) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Payment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Payment.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Payment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Payment.Merge(m, src)
}
func (m *Payment) XXX_Size() int {
	return m.Size()
}
func (m *Payment) XXX_DiscardUnknown() {
	xxx_messageInfo_Payment.DiscardUnknown(m)
}

var xxx_messageInfo_Payment proto.InternalMessageInfo

func (m *Payment) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *Payment) GetAmount() types.Coins {
	if m != nil {
		return m.Amount
	}
	return nil
}

func (m *Payment) GetContract_1155Payment() *Contract1155Payment {
	if m != nil {
		return m.Contract_1155Payment
	}
	return nil
}

func (m *Payment) GetTimeoutNs() time.Duration {
	if m != nil {
		return m.TimeoutNs
	}
	return 0
}

type Contract1155Payment struct {
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	TokenId string `protobuf:"bytes,2,opt,name=token_id,json=tokenId,proto3" json:"token_id,omitempty"`
	Amount  uint32 `protobuf:"varint,3,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (m *Contract1155Payment) Reset()         { *m = Contract1155Payment{} }
func (m *Contract1155Payment) String() string { return proto.CompactTextString(m) }
func (*Contract1155Payment) ProtoMessage()    {}
func (*Contract1155Payment) Descriptor() ([]byte, []int) {
	return fileDescriptor_619c1a0876cd0592, []int{4}
}
func (m *Contract1155Payment) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Contract1155Payment) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Contract1155Payment.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Contract1155Payment) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Contract1155Payment.Merge(m, src)
}
func (m *Contract1155Payment) XXX_Size() int {
	return m.Size()
}
func (m *Contract1155Payment) XXX_DiscardUnknown() {
	xxx_messageInfo_Contract1155Payment.DiscardUnknown(m)
}

var xxx_messageInfo_Contract1155Payment proto.InternalMessageInfo

func (m *Contract1155Payment) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *Contract1155Payment) GetTokenId() string {
	if m != nil {
		return m.TokenId
	}
	return ""
}

func (m *Contract1155Payment) GetAmount() uint32 {
	if m != nil {
		return m.Amount
	}
	return 0
}

type Claim struct {
	// collection_id indicates to which Collection this claim belongs
	CollectionId string `protobuf:"bytes,1,opt,name=collection_id,json=collectionId,proto3" json:"collection_id,omitempty"`
	// agent is the DID of the agent submitting the claim
	AgentDid     string `protobuf:"bytes,2,opt,name=agent_did,json=agentDid,proto3" json:"agent_did,omitempty"`
	AgentAddress string `protobuf:"bytes,3,opt,name=agent_address,json=agentAddress,proto3" json:"agent_address,omitempty"`
	// submissionDate is the date and time that the claim was submitted on-chain
	SubmissionDate *time.Time `protobuf:"bytes,4,opt,name=submission_date,json=submissionDate,proto3,stdtime" json:"submission_date,omitempty"`
	// claimID is the unique identifier of the claim in the cid hash format
	ClaimId string `protobuf:"bytes,5,opt,name=claim_id,json=claimId,proto3" json:"claim_id,omitempty"`
	// evaluation is the result of one or more claim evaluations
	Evaluation     *Evaluation    `protobuf:"bytes,6,opt,name=evaluation,proto3" json:"evaluation,omitempty"`
	PaymentsStatus *ClaimPayments `protobuf:"bytes,7,opt,name=payments_status,json=paymentsStatus,proto3" json:"payments_status,omitempty"`
}

func (m *Claim) Reset()         { *m = Claim{} }
func (m *Claim) String() string { return proto.CompactTextString(m) }
func (*Claim) ProtoMessage()    {}
func (*Claim) Descriptor() ([]byte, []int) {
	return fileDescriptor_619c1a0876cd0592, []int{5}
}
func (m *Claim) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Claim) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Claim.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Claim) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Claim.Merge(m, src)
}
func (m *Claim) XXX_Size() int {
	return m.Size()
}
func (m *Claim) XXX_DiscardUnknown() {
	xxx_messageInfo_Claim.DiscardUnknown(m)
}

var xxx_messageInfo_Claim proto.InternalMessageInfo

func (m *Claim) GetCollectionId() string {
	if m != nil {
		return m.CollectionId
	}
	return ""
}

func (m *Claim) GetAgentDid() string {
	if m != nil {
		return m.AgentDid
	}
	return ""
}

func (m *Claim) GetAgentAddress() string {
	if m != nil {
		return m.AgentAddress
	}
	return ""
}

func (m *Claim) GetSubmissionDate() *time.Time {
	if m != nil {
		return m.SubmissionDate
	}
	return nil
}

func (m *Claim) GetClaimId() string {
	if m != nil {
		return m.ClaimId
	}
	return ""
}

func (m *Claim) GetEvaluation() *Evaluation {
	if m != nil {
		return m.Evaluation
	}
	return nil
}

func (m *Claim) GetPaymentsStatus() *ClaimPayments {
	if m != nil {
		return m.PaymentsStatus
	}
	return nil
}

type ClaimPayments struct {
	Submission PaymentStatus `protobuf:"varint,1,opt,name=submission,proto3,enum=ixo.claims.v1beta1.PaymentStatus" json:"submission,omitempty"`
	Evaluation PaymentStatus `protobuf:"varint,2,opt,name=evaluation,proto3,enum=ixo.claims.v1beta1.PaymentStatus" json:"evaluation,omitempty"`
	Approval   PaymentStatus `protobuf:"varint,3,opt,name=approval,proto3,enum=ixo.claims.v1beta1.PaymentStatus" json:"approval,omitempty"`
	Rejection  PaymentStatus `protobuf:"varint,4,opt,name=rejection,proto3,enum=ixo.claims.v1beta1.PaymentStatus" json:"rejection,omitempty"`
}

func (m *ClaimPayments) Reset()         { *m = ClaimPayments{} }
func (m *ClaimPayments) String() string { return proto.CompactTextString(m) }
func (*ClaimPayments) ProtoMessage()    {}
func (*ClaimPayments) Descriptor() ([]byte, []int) {
	return fileDescriptor_619c1a0876cd0592, []int{6}
}
func (m *ClaimPayments) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ClaimPayments) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ClaimPayments.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ClaimPayments) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ClaimPayments.Merge(m, src)
}
func (m *ClaimPayments) XXX_Size() int {
	return m.Size()
}
func (m *ClaimPayments) XXX_DiscardUnknown() {
	xxx_messageInfo_ClaimPayments.DiscardUnknown(m)
}

var xxx_messageInfo_ClaimPayments proto.InternalMessageInfo

func (m *ClaimPayments) GetSubmission() PaymentStatus {
	if m != nil {
		return m.Submission
	}
	return PaymentStatus_no_payment
}

func (m *ClaimPayments) GetEvaluation() PaymentStatus {
	if m != nil {
		return m.Evaluation
	}
	return PaymentStatus_no_payment
}

func (m *ClaimPayments) GetApproval() PaymentStatus {
	if m != nil {
		return m.Approval
	}
	return PaymentStatus_no_payment
}

func (m *ClaimPayments) GetRejection() PaymentStatus {
	if m != nil {
		return m.Rejection
	}
	return PaymentStatus_no_payment
}

type Evaluation struct {
	// claim_id indicates which Claim this evaluation is for
	ClaimId string `protobuf:"bytes,1,opt,name=claim_id,json=claimId,proto3" json:"claim_id,omitempty"`
	// collection_id indicates to which Collection the claim being evaluated
	// belongs to
	CollectionId string `protobuf:"bytes,2,opt,name=collection_id,json=collectionId,proto3" json:"collection_id,omitempty"`
	// oracle is the DID of the Oracle entity that evaluates the claim
	Oracle string `protobuf:"bytes,3,opt,name=oracle,proto3" json:"oracle,omitempty"`
	// agent is the DID of the agent that submits the evaluation
	AgentDid     string `protobuf:"bytes,4,opt,name=agent_did,json=agentDid,proto3" json:"agent_did,omitempty"`
	AgentAddress string `protobuf:"bytes,5,opt,name=agent_address,json=agentAddress,proto3" json:"agent_address,omitempty"`
	// status is the evaluation status expressed as an integer (2=approved,
	// 3=rejected, ...)
	Status EvaluationStatus `protobuf:"varint,6,opt,name=status,proto3,enum=ixo.claims.v1beta1.EvaluationStatus" json:"status,omitempty"`
	// reason is the code expressed as an integer, for why the evaluation result
	// was given (codes defined by evaluator)
	Reason uint32 `protobuf:"varint,7,opt,name=reason,proto3" json:"reason,omitempty"`
	// verificationProof is the cid of the evaluation Verfiable Credential
	VerificationProof string `protobuf:"bytes,8,opt,name=verification_proof,json=verificationProof,proto3" json:"verification_proof,omitempty"`
	// evaluationDate is the date and time that the claim evaluation was submitted
	// on-chain
	EvaluationDate *time.Time `protobuf:"bytes,9,opt,name=evaluation_date,json=evaluationDate,proto3,stdtime" json:"evaluation_date,omitempty"`
	// custom amount specified by evaluator for claim approval, if empty list then
	// use default by Collection
	Amount types.Coins `protobuf:"bytes,10,rep,name=amount,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.Coins" json:"amount"`
}

func (m *Evaluation) Reset()         { *m = Evaluation{} }
func (m *Evaluation) String() string { return proto.CompactTextString(m) }
func (*Evaluation) ProtoMessage()    {}
func (*Evaluation) Descriptor() ([]byte, []int) {
	return fileDescriptor_619c1a0876cd0592, []int{7}
}
func (m *Evaluation) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Evaluation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Evaluation.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Evaluation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Evaluation.Merge(m, src)
}
func (m *Evaluation) XXX_Size() int {
	return m.Size()
}
func (m *Evaluation) XXX_DiscardUnknown() {
	xxx_messageInfo_Evaluation.DiscardUnknown(m)
}

var xxx_messageInfo_Evaluation proto.InternalMessageInfo

func (m *Evaluation) GetClaimId() string {
	if m != nil {
		return m.ClaimId
	}
	return ""
}

func (m *Evaluation) GetCollectionId() string {
	if m != nil {
		return m.CollectionId
	}
	return ""
}

func (m *Evaluation) GetOracle() string {
	if m != nil {
		return m.Oracle
	}
	return ""
}

func (m *Evaluation) GetAgentDid() string {
	if m != nil {
		return m.AgentDid
	}
	return ""
}

func (m *Evaluation) GetAgentAddress() string {
	if m != nil {
		return m.AgentAddress
	}
	return ""
}

func (m *Evaluation) GetStatus() EvaluationStatus {
	if m != nil {
		return m.Status
	}
	return EvaluationStatus_pending
}

func (m *Evaluation) GetReason() uint32 {
	if m != nil {
		return m.Reason
	}
	return 0
}

func (m *Evaluation) GetVerificationProof() string {
	if m != nil {
		return m.VerificationProof
	}
	return ""
}

func (m *Evaluation) GetEvaluationDate() *time.Time {
	if m != nil {
		return m.EvaluationDate
	}
	return nil
}

func (m *Evaluation) GetAmount() types.Coins {
	if m != nil {
		return m.Amount
	}
	return nil
}

type Dispute struct {
	SubjectId string `protobuf:"bytes,1,opt,name=subject_id,json=subjectId,proto3" json:"subject_id,omitempty"`
	// type is expressed as an integer, interpreted by the client
	Type int32        `protobuf:"varint,2,opt,name=type,proto3" json:"type,omitempty"`
	Data *DisputeData `protobuf:"bytes,3,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *Dispute) Reset()         { *m = Dispute{} }
func (m *Dispute) String() string { return proto.CompactTextString(m) }
func (*Dispute) ProtoMessage()    {}
func (*Dispute) Descriptor() ([]byte, []int) {
	return fileDescriptor_619c1a0876cd0592, []int{8}
}
func (m *Dispute) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Dispute) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Dispute.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Dispute) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Dispute.Merge(m, src)
}
func (m *Dispute) XXX_Size() int {
	return m.Size()
}
func (m *Dispute) XXX_DiscardUnknown() {
	xxx_messageInfo_Dispute.DiscardUnknown(m)
}

var xxx_messageInfo_Dispute proto.InternalMessageInfo

func (m *Dispute) GetSubjectId() string {
	if m != nil {
		return m.SubjectId
	}
	return ""
}

func (m *Dispute) GetType() int32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *Dispute) GetData() *DisputeData {
	if m != nil {
		return m.Data
	}
	return nil
}

type DisputeData struct {
	Uri       string `protobuf:"bytes,1,opt,name=uri,proto3" json:"uri,omitempty"`
	Type      string `protobuf:"bytes,2,opt,name=type,proto3" json:"type,omitempty"`
	Proof     string `protobuf:"bytes,3,opt,name=proof,proto3" json:"proof,omitempty"`
	Encrypted bool   `protobuf:"varint,4,opt,name=encrypted,proto3" json:"encrypted,omitempty"`
}

func (m *DisputeData) Reset()         { *m = DisputeData{} }
func (m *DisputeData) String() string { return proto.CompactTextString(m) }
func (*DisputeData) ProtoMessage()    {}
func (*DisputeData) Descriptor() ([]byte, []int) {
	return fileDescriptor_619c1a0876cd0592, []int{9}
}
func (m *DisputeData) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DisputeData) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DisputeData.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DisputeData) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DisputeData.Merge(m, src)
}
func (m *DisputeData) XXX_Size() int {
	return m.Size()
}
func (m *DisputeData) XXX_DiscardUnknown() {
	xxx_messageInfo_DisputeData.DiscardUnknown(m)
}

var xxx_messageInfo_DisputeData proto.InternalMessageInfo

func (m *DisputeData) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

func (m *DisputeData) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *DisputeData) GetProof() string {
	if m != nil {
		return m.Proof
	}
	return ""
}

func (m *DisputeData) GetEncrypted() bool {
	if m != nil {
		return m.Encrypted
	}
	return false
}

var fileDescriptor_619c1a0876cd0592 = []byte{
	// 1507 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x57, 0xcd, 0x6f, 0x1b, 0x45,
	0x1b, 0xcf, 0x3a, 0xb6, 0x63, 0x3f, 0xce, 0x57, 0x27, 0x79, 0x2b, 0xc7, 0xed, 0x6b, 0xfb, 0x75,
	0x5f, 0x41, 0x14, 0xa9, 0x36, 0x49, 0x55, 0x41, 0x29, 0x50, 0x39, 0xb1, 0x5b, 0x5c, 0xa5, 0x89,
	0xb5, 0x4e, 0x2a, 0x51, 0x24, 0xac, 0xc9, 0xee, 0xc4, 0x19, 0x62, 0xef, 0x6c, 0x77, 0x67, 0x43,
	0xc2, 0x89, 0x1b, 0x52, 0x2e, 0xf4, 0x82, 0xc4, 0x25, 0x52, 0xaf, 0xf0, 0x97, 0xf4, 0x58, 0x89,
	0x0b, 0x02, 0xa9, 0xad, 0xda, 0x0b, 0x7f, 0x02, 0x47, 0x34, 0x1f, 0xeb, 0xb5, 0x13, 0xab, 0xb5,
	0x90, 0x38, 0xc5, 0xcf, 0xc7, 0xef, 0x99, 0x79, 0xbe, 0x7e, 0x3b, 0x81, 0x02, 0x3d, 0x66, 0x15,
	0xab, 0x8b, 0x69, 0xcf, 0xaf, 0x1c, 0xad, 0xee, 0x11, 0x8e, 0x57, 0xb5, 0x58, 0x76, 0x3d, 0xc6,
	0x19, 0x42, 0xf4, 0x98, 0x95, 0xb5, 0x46, 0x3b, 0xe4, 0x16, 0x3b, 0xac, 0xc3, 0xa4, 0xb9, 0x22,
	0x7e, 0x29, 0xcf, 0x5c, 0xa1, 0xc3, 0x58, 0xa7, 0x4b, 0x2a, 0x52, 0xda, 0x0b, 0xf6, 0x2b, 0x9c,
	0xf6, 0x88, 0xcf, 0x71, 0xcf, 0xd5, 0x0e, 0x79, 0x8b, 0xf9, 0x3d, 0xe6, 0x57, 0xf6, 0xb0, 0x4f,
	0xa2, 0xc3, 0x18, 0x75, 0x42, 0xfb, 0xf9, 0x00, 0x76, 0xe0, 0x61, 0x4e, 0x99, 0xb6, 0x97, 0x9e,
	0xc6, 0x20, 0xd9, 0xc4, 0x1e, 0xee, 0xf9, 0xa8, 0x02, 0x0b, 0x16, 0xeb, 0x76, 0x89, 0x25, 0xcc,
	0x6d, 0x9f, 0x3c, 0x0e, 0x88, 0x63, 0x91, 0xac, 0x51, 0x34, 0x96, 0xe3, 0x26, 0x8a, 0x4c, 0x2d,
	0x6d, 0x41, 0x05, 0xc8, 0xd0, 0x63, 0xd6, 0xc6, 0x96, 0xc5, 0x02, 0x87, 0x67, 0x63, 0x45, 0x63,
	0x39, 0x6d, 0x02, 0x3d, 0x66, 0x55, 0xa5, 0x41, 0x36, 0x5c, 0x76, 0x08, 0xff, 0x86, 0x79, 0x87,
	0xed, 0x7d, 0x42, 0xda, 0x2e, 0xf1, 0x2c, 0xe2, 0x70, 0xdc, 0x21, 0xd9, 0x49, 0xe1, 0xbb, 0x5e,
	0x7e, 0xf6, 0xa2, 0x30, 0xf1, 0xfb, 0x8b, 0xc2, 0x7b, 0x1d, 0xca, 0x0f, 0x82, 0xbd, 0xb2, 0xc5,
	0x7a, 0x15, 0x9d, 0x8f, 0xfa, 0x73, 0xdd, 0xb7, 0x0f, 0x2b, 0xfc, 0xc4, 0x25, 0x7e, 0xb9, 0x46,
	0x2c, 0x73, 0x51, 0x47, 0xbb, 0x4b, 0x48, 0xb3, 0x1f, 0x0b, 0x7d, 0x05, 0x0b, 0x0e, 0xb3, 0xc9,
	0xf9, 0x23, 0xe2, 0xff, 0xe8, 0x88, 0x4b, 0x22, 0xd4, 0x50, 0xfc, 0xd2, 0x0f, 0x71, 0x80, 0x8d,
	0x7e, 0xf6, 0x68, 0x16, 0x62, 0xd4, 0x96, 0x55, 0x49, 0x9b, 0x31, 0x6a, 0xa3, 0xcb, 0x90, 0x24,
	0x0e, 0xa7, 0xfc, 0x44, 0x17, 0x40, 0x4b, 0x68, 0x11, 0x12, 0xd8, 0xee, 0x51, 0x47, 0xe5, 0x6a,
	0x2a, 0x01, 0xe5, 0x20, 0x25, 0x0b, 0x6f, 0xb1, 0xae, 0xba, 0xa1, 0xd9, 0x97, 0xd1, 0x1d, 0x00,
	0x9f, 0x63, 0x8f, 0xb7, 0x6d, 0xcc, 0x49, 0x36, 0x51, 0x34, 0x96, 0x33, 0x6b, 0xb9, 0xb2, 0x6a,
	0x60, 0x39, 0x6c, 0x60, 0x79, 0x27, 0x9c, 0x80, 0xf5, 0xf8, 0x93, 0x97, 0x05, 0xc3, 0x4c, 0x4b,
	0x4c, 0x0d, 0x73, 0x82, 0x6e, 0x43, 0x8a, 0x38, 0xb6, 0x82, 0x27, 0xc7, 0x84, 0x4f, 0x11, 0xc7,
	0x96, 0xe0, 0x45, 0x48, 0x3c, 0x0e, 0x18, 0xc7, 0xd9, 0x29, 0xd9, 0x70, 0x25, 0x08, 0xad, 0xea,
	0x6e, 0x4a, 0x69, 0x55, 0x63, 0xaf, 0x42, 0x9a, 0x1c, 0xe1, 0x6e, 0x80, 0x39, 0xb1, 0xb3, 0x69,
	0x69, 0x89, 0x14, 0x22, 0x47, 0xec, 0xba, 0x1e, 0x3b, 0x22, 0x76, 0x16, 0xa4, 0xb1, 0x2f, 0x0b,
	0x9b, 0x47, 0xbe, 0x26, 0x96, 0x00, 0x66, 0x94, 0x2d, 0x94, 0x85, 0xcd, 0xa6, 0xbe, 0x1b, 0x08,
	0xdb, 0xb4, 0xb2, 0x85, 0x32, 0xba, 0x05, 0x09, 0x9f, 0x8b, 0xbc, 0x66, 0x8a, 0xc6, 0xf2, 0xec,
	0xda, 0xb5, 0xf2, 0xc5, 0x15, 0x2a, 0x47, 0x4d, 0x6a, 0x09, 0x57, 0x53, 0x21, 0xd0, 0x47, 0x90,
	0x72, 0xf1, 0x49, 0x8f, 0x38, 0xdc, 0xcf, 0xce, 0xca, 0xaa, 0x5c, 0x1d, 0x85, 0x6e, 0x6a, 0x1f,
	0xb3, 0xef, 0x2d, 0x5a, 0xeb, 0xd3, 0x8e, 0x43, 0xbc, 0xec, 0x9c, 0x6a, 0xad, 0x92, 0x4a, 0x7f,
	0x19, 0x90, 0x0a, 0xdd, 0xd1, 0x6d, 0x00, 0x3f, 0xd8, 0xeb, 0x51, 0xdf, 0xa7, 0xcc, 0x91, 0x73,
	0x91, 0x59, 0xbb, 0xf2, 0x96, 0x03, 0xcc, 0x01, 0x77, 0x01, 0xd6, 0x75, 0x13, 0xe0, 0xd8, 0x18,
	0xe0, 0xc8, 0x1d, 0x7d, 0x18, 0xd6, 0x19, 0x77, 0xe5, 0x90, 0xbd, 0x03, 0xda, 0x77, 0x46, 0xb7,
	0x20, 0xad, 0x8a, 0x2e, 0x0e, 0x8d, 0xbf, 0x1b, 0x19, 0x79, 0x97, 0x7e, 0x8e, 0xc1, 0x94, 0x56,
	0xa3, 0x2c, 0x4c, 0x85, 0xbb, 0xaf, 0xd6, 0x21, 0x14, 0x91, 0x05, 0x49, 0xdc, 0xd3, 0xa4, 0x30,
	0xb9, 0x9c, 0x59, 0x5b, 0x2a, 0xab, 0x65, 0x2b, 0x0b, 0x9a, 0x1a, 0xe8, 0x17, 0x75, 0xd6, 0x3f,
	0x10, 0x0b, 0xfa, 0xcb, 0xcb, 0xc2, 0xf2, 0x18, 0x0b, 0x2a, 0x00, 0xbe, 0xa9, 0x43, 0xa3, 0x2f,
	0xe1, 0x3f, 0x16, 0x73, 0xb8, 0x87, 0x2d, 0xde, 0x5e, 0x5d, 0xbd, 0x79, 0xb3, 0xad, 0xfb, 0xa6,
	0x6b, 0xf1, 0xfe, 0xe8, 0x11, 0x51, 0x00, 0xe1, 0x1f, 0x66, 0xb7, 0x60, 0x5d, 0x54, 0xa2, 0x75,
	0x00, 0x41, 0xb5, 0x2c, 0xe0, 0x6d, 0xc7, 0xd7, 0x35, 0x5a, 0xba, 0xb0, 0x4c, 0x35, 0x4d, 0xa6,
	0xeb, 0x29, 0x91, 0xc5, 0x4f, 0x72, 0x1d, 0x35, 0x6c, 0xcb, 0x2f, 0x1d, 0xc0, 0xc2, 0x88, 0xf3,
	0x64, 0xd9, 0x6c, 0xdb, 0x23, 0xbe, 0xdf, 0x2f, 0x9b, 0x12, 0xd1, 0x12, 0xa4, 0x38, 0x3b, 0x24,
	0x4e, 0x9b, 0xda, 0x9a, 0x4c, 0xa6, 0xa4, 0xdc, 0x90, 0x2c, 0xa3, 0x2b, 0x2a, 0xb2, 0x9b, 0x09,
	0x8b, 0xf0, 0x71, 0xfc, 0xcf, 0xa7, 0x05, 0xa3, 0xf4, 0x2a, 0x06, 0x89, 0x0d, 0x91, 0x29, 0xba,
	0x06, 0x33, 0x03, 0x24, 0xde, 0x27, 0xaa, 0xe9, 0x48, 0xd9, 0xb0, 0xd1, 0x15, 0x48, 0xe3, 0x0e,
	0x71, 0x78, 0xdb, 0xee, 0x1f, 0x94, 0x92, 0x8a, 0x1a, 0xb5, 0x45, 0x04, 0x65, 0x0c, 0x2f, 0xa9,
	0xf8, 0x6b, 0x5a, 0x2a, 0xab, 0xfa, 0xa6, 0x0d, 0x98, 0x8b, 0xa6, 0x58, 0x11, 0x4e, 0x7c, 0x4c,
	0xc2, 0x99, 0x8d, 0x80, 0x92, 0x77, 0x96, 0x20, 0x25, 0x9b, 0x24, 0x2e, 0x9b, 0x50, 0x49, 0x4b,
	0xb9, 0x61, 0xa3, 0xcf, 0x86, 0xb6, 0x43, 0x31, 0x5a, 0x7e, 0x54, 0x5b, 0xeb, 0x7d, 0xaf, 0xa1,
	0x05, 0xb9, 0x0f, 0x73, 0xe1, 0x2e, 0xb7, 0x05, 0x17, 0x04, 0xbe, 0x24, 0xb7, 0xcc, 0xda, 0xff,
	0x46, 0xce, 0x86, 0x10, 0xfb, 0x2c, 0x30, 0x1b, 0x22, 0x5b, 0x12, 0x58, 0xfa, 0x31, 0x06, 0x33,
	0x43, 0x1e, 0xa8, 0x7a, 0x61, 0xf1, 0x67, 0x47, 0x07, 0xd6, 0x08, 0x15, 0x68, 0x68, 0xfd, 0xab,
	0x17, 0xd6, 0x7f, 0xbc, 0x10, 0x03, 0x39, 0x7e, 0x7a, 0x8e, 0x04, 0xc6, 0x0a, 0x10, 0x51, 0xc1,
	0x9d, 0xf3, 0x54, 0x30, 0x16, 0x7e, 0x80, 0x10, 0xfe, 0x98, 0x04, 0x88, 0xca, 0x3f, 0xd4, 0x4d,
	0x63, 0xb8, 0x9b, 0x17, 0x46, 0x33, 0x36, 0x62, 0x34, 0x2f, 0x43, 0x92, 0x79, 0xd8, 0xea, 0xea,
	0x27, 0x82, 0xa9, 0xa5, 0xe1, 0x91, 0x8d, 0xbf, 0x6b, 0x64, 0x13, 0x23, 0x46, 0xf6, 0x13, 0x48,
	0xea, 0x19, 0x48, 0xca, 0x34, 0xff, 0xff, 0xf6, 0x41, 0xd2, 0x99, 0x6a, 0x8c, 0xb8, 0x97, 0x47,
	0xb0, 0xcf, 0x1c, 0x39, 0x41, 0x33, 0xa6, 0x96, 0xd0, 0x75, 0x40, 0x47, 0xc4, 0xa3, 0xfb, 0xd4,
	0x92, 0xa8, 0xb6, 0xeb, 0x31, 0xb6, 0x2f, 0x3f, 0x96, 0x69, 0xf3, 0xd2, 0xa0, 0xa5, 0x29, 0x0c,
	0x62, 0x6f, 0xa2, 0xde, 0xa9, 0xbd, 0x49, 0x8f, 0xbb, 0x37, 0x11, 0x50, 0xee, 0x4d, 0xc4, 0xb1,
	0xf0, 0xaf, 0x71, 0x6c, 0xe9, 0x31, 0x4c, 0xd5, 0xd4, 0x27, 0x18, 0xfd, 0x57, 0x8e, 0xbb, 0x68,
	0x7b, 0xd4, 0xdb, 0xb4, 0xd6, 0x34, 0x6c, 0x84, 0x20, 0x2e, 0x02, 0xc8, 0xa6, 0x26, 0x4c, 0xf9,
	0x1b, 0xdd, 0x80, 0xb8, 0x8d, 0x39, 0xd6, 0x84, 0x5c, 0x18, 0x55, 0x70, 0x1d, 0xbd, 0x86, 0x39,
	0x36, 0xa5, 0x73, 0xa9, 0x03, 0x99, 0x01, 0x25, 0x9a, 0x87, 0xc9, 0xc0, 0xa3, 0xfa, 0x3c, 0xf1,
	0x73, 0xe8, 0xa4, 0xb4, 0x3e, 0x69, 0x11, 0x12, 0xaa, 0xf2, 0xfa, 0xb1, 0x25, 0x05, 0xf9, 0x4c,
	0x71, 0x2c, 0xef, 0xc4, 0x15, 0x2f, 0x0a, 0x31, 0x34, 0x29, 0x33, 0x52, 0xac, 0xec, 0xc2, 0xdc,
	0xb9, 0x17, 0x83, 0x08, 0xbd, 0xdd, 0xac, 0x6f, 0xcd, 0x4f, 0xe4, 0x52, 0xa7, 0x67, 0xc5, 0x38,
	0x73, 0x89, 0x23, 0x3a, 0xdf, 0xac, 0xee, 0xb6, 0xea, 0xb5, 0x79, 0x23, 0x07, 0xa7, 0x67, 0xc5,
	0xa4, 0x8b, 0x03, 0x9f, 0xc8, 0x49, 0xdd, 0xd8, 0xdc, 0x16, 0xfa, 0x98, 0xd2, 0x5b, 0x5d, 0xe6,
	0x13, 0x7b, 0xe5, 0x3b, 0x03, 0xe6, 0xcf, 0x8f, 0x91, 0xe0, 0xfc, 0x66, 0x7d, 0xab, 0xd6, 0xd8,
	0xba, 0x37, 0x3f, 0x91, 0xcb, 0x9c, 0x9e, 0x15, 0xa7, 0x5c, 0xe2, 0xd8, 0xd4, 0xe9, 0x88, 0x47,
	0x4f, 0xb5, 0xd9, 0x34, 0xb7, 0x1f, 0xca, 0x03, 0xa6, 0x4f, 0xcf, 0x8a, 0x43, 0x8f, 0x25, 0xb3,
	0x7e, 0xbf, 0xbe, 0xb1, 0x23, 0x0f, 0x91, 0xb6, 0xc1, 0xc7, 0x52, 0xad, 0xd1, 0x6a, 0xee, 0x0a,
	0xdb, 0xa4, 0xb2, 0x85, 0x8f, 0xa5, 0x95, 0xef, 0x0d, 0xc8, 0xe8, 0x85, 0xdd, 0x11, 0xd5, 0xc9,
	0x03, 0xb4, 0x76, 0xd7, 0x1f, 0x34, 0x5a, 0xad, 0xc6, 0xb6, 0x48, 0x6e, 0xf6, 0xf4, 0xac, 0x38,
	0x48, 0x43, 0xfd, 0x3b, 0x54, 0x37, 0x87, 0xef, 0x80, 0xbb, 0x02, 0x5b, 0x7f, 0x58, 0xdd, 0xdc,
	0xad, 0xee, 0x08, 0x6c, 0x4c, 0x61, 0x07, 0xf8, 0xe7, 0x2a, 0xa4, 0xd5, 0x1d, 0x85, 0x79, 0x32,
	0x37, 0x73, 0x7a, 0x56, 0x8c, 0xd8, 0x61, 0xe5, 0x57, 0x03, 0x66, 0x86, 0xa8, 0x43, 0xc4, 0xdb,
	0xda, 0x6e, 0x37, 0xab, 0x5f, 0x3c, 0xa8, 0x6f, 0xed, 0x84, 0x77, 0x71, 0x58, 0xf8, 0xf1, 0x16,
	0x77, 0x69, 0x9a, 0xdb, 0x0f, 0x1a, 0xad, 0xa8, 0x1e, 0xae, 0xc7, 0x7a, 0x54, 0x94, 0x3c, 0x0f,
	0x50, 0xdd, 0xdd, 0xf9, 0x7c, 0xdb, 0x6c, 0x3c, 0x92, 0x15, 0x91, 0x58, 0x1c, 0xf0, 0x03, 0xe6,
	0xd1, 0x6f, 0x95, 0xfd, 0x5e, 0x75, 0xd7, 0xac, 0x6e, 0xed, 0xd4, 0x65, 0x55, 0xa4, 0xbd, 0x83,
	0x03, 0x0f, 0x3b, 0x9c, 0x10, 0x39, 0xa3, 0xcd, 0x6a, 0xa3, 0x36, 0x1f, 0x57, 0xed, 0x75, 0xb1,
	0x7a, 0xbe, 0xdf, 0xad, 0x36, 0x36, 0xeb, 0xb5, 0xf9, 0x84, 0x6a, 0xe3, 0x3e, 0xa6, 0xdd, 0x73,
	0xf5, 0x4d, 0x0e, 0xd7, 0x77, 0xbd, 0xf5, 0xec, 0x75, 0xde, 0x78, 0xfe, 0x3a, 0x6f, 0xbc, 0x7a,
	0x9d, 0x37, 0x9e, 0xbc, 0xc9, 0x4f, 0x3c, 0x7f, 0x93, 0x9f, 0xf8, 0xed, 0x4d, 0x7e, 0xe2, 0xd1,
	0xad, 0x81, 0x0d, 0xa3, 0xc7, 0x6c, 0x9f, 0x05, 0x8e, 0x2d, 0xeb, 0x24, 0xa4, 0xeb, 0x7b, 0x5d,
	0x66, 0x1d, 0x5a, 0x07, 0x98, 0x3a, 0x95, 0xa3, 0xb5, 0xca, 0x71, 0xf8, 0x3f, 0xa2, 0x5c, 0xbc,
	0xbd, 0xa4, 0xdc, 0xfc, 0x1b, 0x7f, 0x07, 0x00, 0x00, 0xff, 0xff, 0xef, 0x10, 0xad, 0xdc, 0x3e,
	0x0e, 0x00, 0x00,
}

func (this *Contract1155Payment) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Contract1155Payment)
	if !ok {
		that2, ok := that.(Contract1155Payment)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Address != that1.Address {
		return false
	}
	if this.TokenId != that1.TokenId {
		return false
	}
	if this.Amount != that1.Amount {
		return false
	}
	return true
}
func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.NodeFeePercentage.Size()
		i -= size
		if _, err := m.NodeFeePercentage.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintClaims(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	{
		size := m.NetworkFeePercentage.Size()
		i -= size
		if _, err := m.NetworkFeePercentage.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintClaims(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x1a
	if len(m.IxoAccount) > 0 {
		i -= len(m.IxoAccount)
		copy(dAtA[i:], m.IxoAccount)
		i = encodeVarintClaims(dAtA, i, uint64(len(m.IxoAccount)))
		i--
		dAtA[i] = 0x12
	}
	if m.CollectionSequence != 0 {
		i = encodeVarintClaims(dAtA, i, uint64(m.CollectionSequence))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Collection) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Collection) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Collection) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signer) > 0 {
		i -= len(m.Signer)
		copy(dAtA[i:], m.Signer)
		i = encodeVarintClaims(dAtA, i, uint64(len(m.Signer)))
		i--
		dAtA[i] = 0x7a
	}
	if m.Payments != nil {
		{
			size, err := m.Payments.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintClaims(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x72
	}
	if m.State != 0 {
		i = encodeVarintClaims(dAtA, i, uint64(m.State))
		i--
		dAtA[i] = 0x68
	}
	if m.Disputed != 0 {
		i = encodeVarintClaims(dAtA, i, uint64(m.Disputed))
		i--
		dAtA[i] = 0x60
	}
	if m.Rejected != 0 {
		i = encodeVarintClaims(dAtA, i, uint64(m.Rejected))
		i--
		dAtA[i] = 0x58
	}
	if m.Approved != 0 {
		i = encodeVarintClaims(dAtA, i, uint64(m.Approved))
		i--
		dAtA[i] = 0x50
	}
	if m.Evaluated != 0 {
		i = encodeVarintClaims(dAtA, i, uint64(m.Evaluated))
		i--
		dAtA[i] = 0x48
	}
	if m.Count != 0 {
		i = encodeVarintClaims(dAtA, i, uint64(m.Count))
		i--
		dAtA[i] = 0x40
	}
	if m.Quota != 0 {
		i = encodeVarintClaims(dAtA, i, uint64(m.Quota))
		i--
		dAtA[i] = 0x38
	}
	if m.EndDate != nil {
		n2, err2 := github_com_gogo_protobuf_types.StdTimeMarshalTo(*m.EndDate, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(*m.EndDate):])
		if err2 != nil {
			return 0, err2
		}
		i -= n2
		i = encodeVarintClaims(dAtA, i, uint64(n2))
		i--
		dAtA[i] = 0x32
	}
	if m.StartDate != nil {
		n3, err3 := github_com_gogo_protobuf_types.StdTimeMarshalTo(*m.StartDate, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(*m.StartDate):])
		if err3 != nil {
			return 0, err3
		}
		i -= n3
		i = encodeVarintClaims(dAtA, i, uint64(n3))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Protocol) > 0 {
		i -= len(m.Protocol)
		copy(dAtA[i:], m.Protocol)
		i = encodeVarintClaims(dAtA, i, uint64(len(m.Protocol)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Admin) > 0 {
		i -= len(m.Admin)
		copy(dAtA[i:], m.Admin)
		i = encodeVarintClaims(dAtA, i, uint64(len(m.Admin)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Entity) > 0 {
		i -= len(m.Entity)
		copy(dAtA[i:], m.Entity)
		i = encodeVarintClaims(dAtA, i, uint64(len(m.Entity)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintClaims(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Payments) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Payments) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Payments) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Rejection != nil {
		{
			size, err := m.Rejection.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintClaims(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.Approval != nil {
		{
			size, err := m.Approval.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintClaims(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.Evaluation != nil {
		{
			size, err := m.Evaluation.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintClaims(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Submission != nil {
		{
			size, err := m.Submission.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintClaims(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Payment) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Payment) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Payment) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	n8, err8 := github_com_gogo_protobuf_types.StdDurationMarshalTo(m.TimeoutNs, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdDuration(m.TimeoutNs):])
	if err8 != nil {
		return 0, err8
	}
	i -= n8
	i = encodeVarintClaims(dAtA, i, uint64(n8))
	i--
	dAtA[i] = 0x22
	if m.Contract_1155Payment != nil {
		{
			size, err := m.Contract_1155Payment.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintClaims(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Amount) > 0 {
		for iNdEx := len(m.Amount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Amount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintClaims(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Account) > 0 {
		i -= len(m.Account)
		copy(dAtA[i:], m.Account)
		i = encodeVarintClaims(dAtA, i, uint64(len(m.Account)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Contract1155Payment) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Contract1155Payment) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Contract1155Payment) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Amount != 0 {
		i = encodeVarintClaims(dAtA, i, uint64(m.Amount))
		i--
		dAtA[i] = 0x18
	}
	if len(m.TokenId) > 0 {
		i -= len(m.TokenId)
		copy(dAtA[i:], m.TokenId)
		i = encodeVarintClaims(dAtA, i, uint64(len(m.TokenId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintClaims(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Claim) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Claim) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Claim) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.PaymentsStatus != nil {
		{
			size, err := m.PaymentsStatus.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintClaims(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if m.Evaluation != nil {
		{
			size, err := m.Evaluation.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintClaims(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if len(m.ClaimId) > 0 {
		i -= len(m.ClaimId)
		copy(dAtA[i:], m.ClaimId)
		i = encodeVarintClaims(dAtA, i, uint64(len(m.ClaimId)))
		i--
		dAtA[i] = 0x2a
	}
	if m.SubmissionDate != nil {
		n12, err12 := github_com_gogo_protobuf_types.StdTimeMarshalTo(*m.SubmissionDate, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(*m.SubmissionDate):])
		if err12 != nil {
			return 0, err12
		}
		i -= n12
		i = encodeVarintClaims(dAtA, i, uint64(n12))
		i--
		dAtA[i] = 0x22
	}
	if len(m.AgentAddress) > 0 {
		i -= len(m.AgentAddress)
		copy(dAtA[i:], m.AgentAddress)
		i = encodeVarintClaims(dAtA, i, uint64(len(m.AgentAddress)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.AgentDid) > 0 {
		i -= len(m.AgentDid)
		copy(dAtA[i:], m.AgentDid)
		i = encodeVarintClaims(dAtA, i, uint64(len(m.AgentDid)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.CollectionId) > 0 {
		i -= len(m.CollectionId)
		copy(dAtA[i:], m.CollectionId)
		i = encodeVarintClaims(dAtA, i, uint64(len(m.CollectionId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ClaimPayments) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ClaimPayments) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ClaimPayments) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Rejection != 0 {
		i = encodeVarintClaims(dAtA, i, uint64(m.Rejection))
		i--
		dAtA[i] = 0x20
	}
	if m.Approval != 0 {
		i = encodeVarintClaims(dAtA, i, uint64(m.Approval))
		i--
		dAtA[i] = 0x18
	}
	if m.Evaluation != 0 {
		i = encodeVarintClaims(dAtA, i, uint64(m.Evaluation))
		i--
		dAtA[i] = 0x10
	}
	if m.Submission != 0 {
		i = encodeVarintClaims(dAtA, i, uint64(m.Submission))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *Evaluation) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Evaluation) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Evaluation) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Amount) > 0 {
		for iNdEx := len(m.Amount) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Amount[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintClaims(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x52
		}
	}
	if m.EvaluationDate != nil {
		n13, err13 := github_com_gogo_protobuf_types.StdTimeMarshalTo(*m.EvaluationDate, dAtA[i-github_com_gogo_protobuf_types.SizeOfStdTime(*m.EvaluationDate):])
		if err13 != nil {
			return 0, err13
		}
		i -= n13
		i = encodeVarintClaims(dAtA, i, uint64(n13))
		i--
		dAtA[i] = 0x4a
	}
	if len(m.VerificationProof) > 0 {
		i -= len(m.VerificationProof)
		copy(dAtA[i:], m.VerificationProof)
		i = encodeVarintClaims(dAtA, i, uint64(len(m.VerificationProof)))
		i--
		dAtA[i] = 0x42
	}
	if m.Reason != 0 {
		i = encodeVarintClaims(dAtA, i, uint64(m.Reason))
		i--
		dAtA[i] = 0x38
	}
	if m.Status != 0 {
		i = encodeVarintClaims(dAtA, i, uint64(m.Status))
		i--
		dAtA[i] = 0x30
	}
	if len(m.AgentAddress) > 0 {
		i -= len(m.AgentAddress)
		copy(dAtA[i:], m.AgentAddress)
		i = encodeVarintClaims(dAtA, i, uint64(len(m.AgentAddress)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.AgentDid) > 0 {
		i -= len(m.AgentDid)
		copy(dAtA[i:], m.AgentDid)
		i = encodeVarintClaims(dAtA, i, uint64(len(m.AgentDid)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Oracle) > 0 {
		i -= len(m.Oracle)
		copy(dAtA[i:], m.Oracle)
		i = encodeVarintClaims(dAtA, i, uint64(len(m.Oracle)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.CollectionId) > 0 {
		i -= len(m.CollectionId)
		copy(dAtA[i:], m.CollectionId)
		i = encodeVarintClaims(dAtA, i, uint64(len(m.CollectionId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ClaimId) > 0 {
		i -= len(m.ClaimId)
		copy(dAtA[i:], m.ClaimId)
		i = encodeVarintClaims(dAtA, i, uint64(len(m.ClaimId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Dispute) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Dispute) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Dispute) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Data != nil {
		{
			size, err := m.Data.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintClaims(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x1a
	}
	if m.Type != 0 {
		i = encodeVarintClaims(dAtA, i, uint64(m.Type))
		i--
		dAtA[i] = 0x10
	}
	if len(m.SubjectId) > 0 {
		i -= len(m.SubjectId)
		copy(dAtA[i:], m.SubjectId)
		i = encodeVarintClaims(dAtA, i, uint64(len(m.SubjectId)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DisputeData) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DisputeData) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DisputeData) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Encrypted {
		i--
		if m.Encrypted {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if len(m.Proof) > 0 {
		i -= len(m.Proof)
		copy(dAtA[i:], m.Proof)
		i = encodeVarintClaims(dAtA, i, uint64(len(m.Proof)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Type) > 0 {
		i -= len(m.Type)
		copy(dAtA[i:], m.Type)
		i = encodeVarintClaims(dAtA, i, uint64(len(m.Type)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Uri) > 0 {
		i -= len(m.Uri)
		copy(dAtA[i:], m.Uri)
		i = encodeVarintClaims(dAtA, i, uint64(len(m.Uri)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintClaims(dAtA []byte, offset int, v uint64) int {
	offset -= sovClaims(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CollectionSequence != 0 {
		n += 1 + sovClaims(uint64(m.CollectionSequence))
	}
	l = len(m.IxoAccount)
	if l > 0 {
		n += 1 + l + sovClaims(uint64(l))
	}
	l = m.NetworkFeePercentage.Size()
	n += 1 + l + sovClaims(uint64(l))
	l = m.NodeFeePercentage.Size()
	n += 1 + l + sovClaims(uint64(l))
	return n
}

func (m *Collection) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovClaims(uint64(l))
	}
	l = len(m.Entity)
	if l > 0 {
		n += 1 + l + sovClaims(uint64(l))
	}
	l = len(m.Admin)
	if l > 0 {
		n += 1 + l + sovClaims(uint64(l))
	}
	l = len(m.Protocol)
	if l > 0 {
		n += 1 + l + sovClaims(uint64(l))
	}
	if m.StartDate != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.StartDate)
		n += 1 + l + sovClaims(uint64(l))
	}
	if m.EndDate != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.EndDate)
		n += 1 + l + sovClaims(uint64(l))
	}
	if m.Quota != 0 {
		n += 1 + sovClaims(uint64(m.Quota))
	}
	if m.Count != 0 {
		n += 1 + sovClaims(uint64(m.Count))
	}
	if m.Evaluated != 0 {
		n += 1 + sovClaims(uint64(m.Evaluated))
	}
	if m.Approved != 0 {
		n += 1 + sovClaims(uint64(m.Approved))
	}
	if m.Rejected != 0 {
		n += 1 + sovClaims(uint64(m.Rejected))
	}
	if m.Disputed != 0 {
		n += 1 + sovClaims(uint64(m.Disputed))
	}
	if m.State != 0 {
		n += 1 + sovClaims(uint64(m.State))
	}
	if m.Payments != nil {
		l = m.Payments.Size()
		n += 1 + l + sovClaims(uint64(l))
	}
	l = len(m.Signer)
	if l > 0 {
		n += 1 + l + sovClaims(uint64(l))
	}
	return n
}

func (m *Payments) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Submission != nil {
		l = m.Submission.Size()
		n += 1 + l + sovClaims(uint64(l))
	}
	if m.Evaluation != nil {
		l = m.Evaluation.Size()
		n += 1 + l + sovClaims(uint64(l))
	}
	if m.Approval != nil {
		l = m.Approval.Size()
		n += 1 + l + sovClaims(uint64(l))
	}
	if m.Rejection != nil {
		l = m.Rejection.Size()
		n += 1 + l + sovClaims(uint64(l))
	}
	return n
}

func (m *Payment) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Account)
	if l > 0 {
		n += 1 + l + sovClaims(uint64(l))
	}
	if len(m.Amount) > 0 {
		for _, e := range m.Amount {
			l = e.Size()
			n += 1 + l + sovClaims(uint64(l))
		}
	}
	if m.Contract_1155Payment != nil {
		l = m.Contract_1155Payment.Size()
		n += 1 + l + sovClaims(uint64(l))
	}
	l = github_com_gogo_protobuf_types.SizeOfStdDuration(m.TimeoutNs)
	n += 1 + l + sovClaims(uint64(l))
	return n
}

func (m *Contract1155Payment) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovClaims(uint64(l))
	}
	l = len(m.TokenId)
	if l > 0 {
		n += 1 + l + sovClaims(uint64(l))
	}
	if m.Amount != 0 {
		n += 1 + sovClaims(uint64(m.Amount))
	}
	return n
}

func (m *Claim) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.CollectionId)
	if l > 0 {
		n += 1 + l + sovClaims(uint64(l))
	}
	l = len(m.AgentDid)
	if l > 0 {
		n += 1 + l + sovClaims(uint64(l))
	}
	l = len(m.AgentAddress)
	if l > 0 {
		n += 1 + l + sovClaims(uint64(l))
	}
	if m.SubmissionDate != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.SubmissionDate)
		n += 1 + l + sovClaims(uint64(l))
	}
	l = len(m.ClaimId)
	if l > 0 {
		n += 1 + l + sovClaims(uint64(l))
	}
	if m.Evaluation != nil {
		l = m.Evaluation.Size()
		n += 1 + l + sovClaims(uint64(l))
	}
	if m.PaymentsStatus != nil {
		l = m.PaymentsStatus.Size()
		n += 1 + l + sovClaims(uint64(l))
	}
	return n
}

func (m *ClaimPayments) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Submission != 0 {
		n += 1 + sovClaims(uint64(m.Submission))
	}
	if m.Evaluation != 0 {
		n += 1 + sovClaims(uint64(m.Evaluation))
	}
	if m.Approval != 0 {
		n += 1 + sovClaims(uint64(m.Approval))
	}
	if m.Rejection != 0 {
		n += 1 + sovClaims(uint64(m.Rejection))
	}
	return n
}

func (m *Evaluation) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ClaimId)
	if l > 0 {
		n += 1 + l + sovClaims(uint64(l))
	}
	l = len(m.CollectionId)
	if l > 0 {
		n += 1 + l + sovClaims(uint64(l))
	}
	l = len(m.Oracle)
	if l > 0 {
		n += 1 + l + sovClaims(uint64(l))
	}
	l = len(m.AgentDid)
	if l > 0 {
		n += 1 + l + sovClaims(uint64(l))
	}
	l = len(m.AgentAddress)
	if l > 0 {
		n += 1 + l + sovClaims(uint64(l))
	}
	if m.Status != 0 {
		n += 1 + sovClaims(uint64(m.Status))
	}
	if m.Reason != 0 {
		n += 1 + sovClaims(uint64(m.Reason))
	}
	l = len(m.VerificationProof)
	if l > 0 {
		n += 1 + l + sovClaims(uint64(l))
	}
	if m.EvaluationDate != nil {
		l = github_com_gogo_protobuf_types.SizeOfStdTime(*m.EvaluationDate)
		n += 1 + l + sovClaims(uint64(l))
	}
	if len(m.Amount) > 0 {
		for _, e := range m.Amount {
			l = e.Size()
			n += 1 + l + sovClaims(uint64(l))
		}
	}
	return n
}

func (m *Dispute) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.SubjectId)
	if l > 0 {
		n += 1 + l + sovClaims(uint64(l))
	}
	if m.Type != 0 {
		n += 1 + sovClaims(uint64(m.Type))
	}
	if m.Data != nil {
		l = m.Data.Size()
		n += 1 + l + sovClaims(uint64(l))
	}
	return n
}

func (m *DisputeData) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Uri)
	if l > 0 {
		n += 1 + l + sovClaims(uint64(l))
	}
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovClaims(uint64(l))
	}
	l = len(m.Proof)
	if l > 0 {
		n += 1 + l + sovClaims(uint64(l))
	}
	if m.Encrypted {
		n += 2
	}
	return n
}

func sovClaims(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozClaims(x uint64) (n int) {
	return sovClaims(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClaims
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CollectionSequence", wireType)
			}
			m.CollectionSequence = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaims
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CollectionSequence |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IxoAccount", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaims
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthClaims
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClaims
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IxoAccount = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NetworkFeePercentage", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaims
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthClaims
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClaims
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.NetworkFeePercentage.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NodeFeePercentage", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaims
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthClaims
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClaims
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.NodeFeePercentage.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipClaims(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthClaims
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Collection) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClaims
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Collection: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Collection: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaims
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthClaims
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClaims
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Entity", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaims
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthClaims
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClaims
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Entity = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Admin", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaims
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthClaims
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClaims
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Admin = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Protocol", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaims
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthClaims
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClaims
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Protocol = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field StartDate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaims
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthClaims
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClaims
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.StartDate == nil {
				m.StartDate = new(time.Time)
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(m.StartDate, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EndDate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaims
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthClaims
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClaims
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.EndDate == nil {
				m.EndDate = new(time.Time)
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(m.EndDate, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Quota", wireType)
			}
			m.Quota = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaims
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Quota |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Count", wireType)
			}
			m.Count = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaims
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Count |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 9:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Evaluated", wireType)
			}
			m.Evaluated = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaims
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Evaluated |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 10:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Approved", wireType)
			}
			m.Approved = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaims
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Approved |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 11:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rejected", wireType)
			}
			m.Rejected = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaims
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Rejected |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 12:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Disputed", wireType)
			}
			m.Disputed = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaims
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Disputed |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 13:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
			}
			m.State = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaims
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.State |= CollectionState(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 14:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Payments", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaims
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthClaims
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClaims
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Payments == nil {
				m.Payments = &Payments{}
			}
			if err := m.Payments.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 15:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaims
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthClaims
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClaims
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipClaims(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthClaims
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Payments) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClaims
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Payments: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Payments: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Submission", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaims
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthClaims
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClaims
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Submission == nil {
				m.Submission = &Payment{}
			}
			if err := m.Submission.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Evaluation", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaims
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthClaims
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClaims
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Evaluation == nil {
				m.Evaluation = &Payment{}
			}
			if err := m.Evaluation.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Approval", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaims
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthClaims
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClaims
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Approval == nil {
				m.Approval = &Payment{}
			}
			if err := m.Approval.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rejection", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaims
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthClaims
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClaims
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Rejection == nil {
				m.Rejection = &Payment{}
			}
			if err := m.Rejection.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipClaims(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthClaims
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Payment) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClaims
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Payment: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Payment: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Account", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaims
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthClaims
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClaims
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Account = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaims
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthClaims
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClaims
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = append(m.Amount, types.Coin{})
			if err := m.Amount[len(m.Amount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contract_1155Payment", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaims
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthClaims
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClaims
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Contract_1155Payment == nil {
				m.Contract_1155Payment = &Contract1155Payment{}
			}
			if err := m.Contract_1155Payment.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutNs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaims
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthClaims
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClaims
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := github_com_gogo_protobuf_types.StdDurationUnmarshal(&m.TimeoutNs, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipClaims(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthClaims
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Contract1155Payment) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClaims
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Contract1155Payment: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Contract1155Payment: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaims
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthClaims
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClaims
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TokenId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaims
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthClaims
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClaims
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TokenId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			m.Amount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaims
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Amount |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipClaims(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthClaims
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Claim) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClaims
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Claim: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Claim: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CollectionId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaims
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthClaims
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClaims
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CollectionId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AgentDid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaims
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthClaims
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClaims
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AgentDid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AgentAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaims
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthClaims
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClaims
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AgentAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubmissionDate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaims
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthClaims
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClaims
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.SubmissionDate == nil {
				m.SubmissionDate = new(time.Time)
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(m.SubmissionDate, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaims
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthClaims
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClaims
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClaimId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Evaluation", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaims
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthClaims
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClaims
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Evaluation == nil {
				m.Evaluation = &Evaluation{}
			}
			if err := m.Evaluation.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PaymentsStatus", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaims
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthClaims
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClaims
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.PaymentsStatus == nil {
				m.PaymentsStatus = &ClaimPayments{}
			}
			if err := m.PaymentsStatus.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipClaims(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthClaims
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *ClaimPayments) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClaims
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: ClaimPayments: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ClaimPayments: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Submission", wireType)
			}
			m.Submission = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaims
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Submission |= PaymentStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Evaluation", wireType)
			}
			m.Evaluation = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaims
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Evaluation |= PaymentStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Approval", wireType)
			}
			m.Approval = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaims
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Approval |= PaymentStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Rejection", wireType)
			}
			m.Rejection = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaims
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Rejection |= PaymentStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipClaims(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthClaims
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Evaluation) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClaims
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Evaluation: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Evaluation: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ClaimId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaims
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthClaims
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClaims
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ClaimId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CollectionId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaims
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthClaims
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClaims
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.CollectionId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Oracle", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaims
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthClaims
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClaims
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Oracle = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AgentDid", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaims
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthClaims
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClaims
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AgentDid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AgentAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaims
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthClaims
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClaims
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AgentAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
			}
			m.Status = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaims
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Status |= EvaluationStatus(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Reason", wireType)
			}
			m.Reason = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaims
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Reason |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VerificationProof", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaims
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthClaims
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClaims
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VerificationProof = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EvaluationDate", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaims
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthClaims
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClaims
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.EvaluationDate == nil {
				m.EvaluationDate = new(time.Time)
			}
			if err := github_com_gogo_protobuf_types.StdTimeUnmarshal(m.EvaluationDate, dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 10:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaims
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthClaims
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClaims
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Amount = append(m.Amount, types.Coin{})
			if err := m.Amount[len(m.Amount)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipClaims(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthClaims
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Dispute) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClaims
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: Dispute: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Dispute: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SubjectId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaims
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthClaims
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClaims
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.SubjectId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			m.Type = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaims
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Type |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Data", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaims
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthClaims
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthClaims
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Data == nil {
				m.Data = &DisputeData{}
			}
			if err := m.Data.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipClaims(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthClaims
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *DisputeData) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowClaims
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: DisputeData: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DisputeData: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Uri", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaims
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthClaims
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClaims
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Uri = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaims
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthClaims
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClaims
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proof", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaims
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthClaims
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthClaims
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Proof = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Encrypted", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowClaims
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Encrypted = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipClaims(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthClaims
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipClaims(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowClaims
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowClaims
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowClaims
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthClaims
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupClaims
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthClaims
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthClaims        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowClaims          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupClaims = fmt.Errorf("proto: unexpected end of group")
)
