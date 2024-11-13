// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ixo/smartaccount/v1beta1/event.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
	_ "github.com/ixofoundation/ixo-blockchain/v3/x/token/types"
	io "io"
	math "math"
	math_bits "math/bits"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// AuthenticatorAddedEvent is an event triggered on Authenticator addition
type AuthenticatorAddedEvent struct {
	// sender is the address of the account that added the authenticator
	Sender string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	// authenticator_type is the type of the authenticator that was added
	AuthenticatorType string `protobuf:"bytes,2,opt,name=authenticator_type,json=authenticatorType,proto3" json:"authenticator_type,omitempty"`
	// authenticator_id is the id of the authenticator that was added
	AuthenticatorId string `protobuf:"bytes,3,opt,name=authenticator_id,json=authenticatorId,proto3" json:"authenticator_id,omitempty"`
}

func (m *AuthenticatorAddedEvent) Reset()         { *m = AuthenticatorAddedEvent{} }
func (m *AuthenticatorAddedEvent) String() string { return proto.CompactTextString(m) }
func (*AuthenticatorAddedEvent) ProtoMessage()    {}
func (*AuthenticatorAddedEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_052f74205b3b0272, []int{0}
}
func (m *AuthenticatorAddedEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AuthenticatorAddedEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuthenticatorAddedEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AuthenticatorAddedEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticatorAddedEvent.Merge(m, src)
}
func (m *AuthenticatorAddedEvent) XXX_Size() int {
	return m.Size()
}
func (m *AuthenticatorAddedEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticatorAddedEvent.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticatorAddedEvent proto.InternalMessageInfo

func (m *AuthenticatorAddedEvent) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *AuthenticatorAddedEvent) GetAuthenticatorType() string {
	if m != nil {
		return m.AuthenticatorType
	}
	return ""
}

func (m *AuthenticatorAddedEvent) GetAuthenticatorId() string {
	if m != nil {
		return m.AuthenticatorId
	}
	return ""
}

// AuthenticatorRemovedEvent is an event triggered on Authenticator removal
type AuthenticatorRemovedEvent struct {
	// sender is the address of the account that removed the authenticator
	Sender string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	// authenticator_id is the id of the authenticator that was removed
	AuthenticatorId string `protobuf:"bytes,2,opt,name=authenticator_id,json=authenticatorId,proto3" json:"authenticator_id,omitempty"`
}

func (m *AuthenticatorRemovedEvent) Reset()         { *m = AuthenticatorRemovedEvent{} }
func (m *AuthenticatorRemovedEvent) String() string { return proto.CompactTextString(m) }
func (*AuthenticatorRemovedEvent) ProtoMessage()    {}
func (*AuthenticatorRemovedEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_052f74205b3b0272, []int{1}
}
func (m *AuthenticatorRemovedEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AuthenticatorRemovedEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuthenticatorRemovedEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AuthenticatorRemovedEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticatorRemovedEvent.Merge(m, src)
}
func (m *AuthenticatorRemovedEvent) XXX_Size() int {
	return m.Size()
}
func (m *AuthenticatorRemovedEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticatorRemovedEvent.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticatorRemovedEvent proto.InternalMessageInfo

func (m *AuthenticatorRemovedEvent) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *AuthenticatorRemovedEvent) GetAuthenticatorId() string {
	if m != nil {
		return m.AuthenticatorId
	}
	return ""
}

// AuthenticatorSetActiveStateEvent is an event triggered on Authenticator
// active state change
type AuthenticatorSetActiveStateEvent struct {
	// sender is the address of the account that changed the active state
	Sender string `protobuf:"bytes,1,opt,name=sender,proto3" json:"sender,omitempty"`
	// active is the new active state
	IsSmartAccountActive bool `protobuf:"varint,2,opt,name=is_smart_account_active,json=isSmartAccountActive,proto3" json:"is_smart_account_active,omitempty"`
}

func (m *AuthenticatorSetActiveStateEvent) Reset()         { *m = AuthenticatorSetActiveStateEvent{} }
func (m *AuthenticatorSetActiveStateEvent) String() string { return proto.CompactTextString(m) }
func (*AuthenticatorSetActiveStateEvent) ProtoMessage()    {}
func (*AuthenticatorSetActiveStateEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_052f74205b3b0272, []int{2}
}
func (m *AuthenticatorSetActiveStateEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *AuthenticatorSetActiveStateEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_AuthenticatorSetActiveStateEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *AuthenticatorSetActiveStateEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AuthenticatorSetActiveStateEvent.Merge(m, src)
}
func (m *AuthenticatorSetActiveStateEvent) XXX_Size() int {
	return m.Size()
}
func (m *AuthenticatorSetActiveStateEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_AuthenticatorSetActiveStateEvent.DiscardUnknown(m)
}

var xxx_messageInfo_AuthenticatorSetActiveStateEvent proto.InternalMessageInfo

func (m *AuthenticatorSetActiveStateEvent) GetSender() string {
	if m != nil {
		return m.Sender
	}
	return ""
}

func (m *AuthenticatorSetActiveStateEvent) GetIsSmartAccountActive() bool {
	if m != nil {
		return m.IsSmartAccountActive
	}
	return false
}

func init() {
	proto.RegisterType((*AuthenticatorAddedEvent)(nil), "ixo.smartaccount.v1beta1.AuthenticatorAddedEvent")
	proto.RegisterType((*AuthenticatorRemovedEvent)(nil), "ixo.smartaccount.v1beta1.AuthenticatorRemovedEvent")
	proto.RegisterType((*AuthenticatorSetActiveStateEvent)(nil), "ixo.smartaccount.v1beta1.AuthenticatorSetActiveStateEvent")
}

func init() {
	proto.RegisterFile("ixo/smartaccount/v1beta1/event.proto", fileDescriptor_052f74205b3b0272)
}

var fileDescriptor_052f74205b3b0272 = []byte{
	// 348 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xcf, 0x4e, 0xf2, 0x40,
	0x10, 0xa7, 0x7c, 0x09, 0xf9, 0xdc, 0x8b, 0xda, 0x10, 0x29, 0x24, 0x36, 0x84, 0x78, 0xd0, 0x03,
	0x6d, 0x08, 0xf1, 0x6e, 0x4d, 0x3c, 0x78, 0x05, 0x4f, 0x1c, 0x6c, 0xb6, 0xdb, 0x11, 0x36, 0xd8,
	0x1d, 0x6c, 0xa7, 0x4d, 0x79, 0x06, 0x2f, 0x3e, 0x96, 0x47, 0x8e, 0x1e, 0x0d, 0xbc, 0x88, 0xe9,
	0xb6, 0x1a, 0x9a, 0x10, 0xbd, 0xed, 0x6f, 0x7e, 0x7f, 0x66, 0xb2, 0x33, 0xec, 0x42, 0xe6, 0xe8,
	0x26, 0x11, 0x8f, 0x89, 0x0b, 0x81, 0xa9, 0x22, 0x37, 0x1b, 0x05, 0x40, 0x7c, 0xe4, 0x42, 0x06,
	0x8a, 0x9c, 0x55, 0x8c, 0x84, 0xa6, 0x25, 0x73, 0x74, 0xf6, 0x55, 0x4e, 0xa5, 0xea, 0x9d, 0x17,
	0x7e, 0xc2, 0x25, 0xa8, 0x1f, 0xa3, 0x46, 0xa5, 0xb1, 0xd7, 0x3b, 0x40, 0xe7, 0x15, 0xd7, 0x9e,
	0xe3, 0x1c, 0xf5, 0xd3, 0x2d, 0x5e, 0x55, 0xb5, 0x2b, 0x30, 0x89, 0x30, 0xf1, 0x4b, 0xa2, 0x04,
	0x25, 0x35, 0x78, 0x35, 0x58, 0xc7, 0x4b, 0x69, 0x01, 0x8a, 0xa4, 0xe0, 0x84, 0xb1, 0x17, 0x86,
	0x10, 0xde, 0x15, 0x73, 0x9a, 0x67, 0xac, 0x95, 0x80, 0x0a, 0x21, 0xb6, 0x8c, 0xbe, 0x71, 0x79,
	0x34, 0xa9, 0x90, 0x39, 0x64, 0x26, 0xdf, 0xb7, 0xf8, 0xb4, 0x5e, 0x81, 0xd5, 0xd4, 0x9a, 0xd3,
	0x1a, 0xf3, 0xb0, 0x5e, 0x81, 0x79, 0xc5, 0x4e, 0xea, 0x72, 0x19, 0x5a, 0xff, 0xb4, 0xf8, 0xb8,
	0x56, 0xbf, 0x0f, 0x07, 0x8f, 0xac, 0x5b, 0x1b, 0x66, 0x02, 0x11, 0x66, 0x7f, 0x8d, 0x73, 0x28,
	0xbf, 0x79, 0x38, 0xff, 0x85, 0xf5, 0x6b, 0xf9, 0x53, 0x20, 0x4f, 0x90, 0xcc, 0x60, 0x4a, 0x9c,
	0xe0, 0xf7, 0x36, 0xd7, 0xac, 0x23, 0x13, 0x5f, 0x2f, 0xcc, 0xaf, 0x36, 0xe6, 0x73, 0x6d, 0xd6,
	0xdd, 0xfe, 0x4f, 0xda, 0x32, 0x99, 0x16, 0xac, 0x57, 0x92, 0x65, 0xf0, 0xed, 0xec, 0x7d, 0x6b,
	0x1b, 0x9b, 0xad, 0x6d, 0x7c, 0x6e, 0x6d, 0xe3, 0x6d, 0x67, 0x37, 0x36, 0x3b, 0xbb, 0xf1, 0xb1,
	0xb3, 0x1b, 0xb3, 0x9b, 0xb9, 0xa4, 0x45, 0x1a, 0x38, 0x02, 0x23, 0x57, 0xe6, 0xf8, 0x84, 0xa9,
	0x0a, 0x39, 0x49, 0x54, 0x05, 0x1a, 0x06, 0xcf, 0x28, 0x96, 0x62, 0xc1, 0xa5, 0x72, 0xb3, 0xb1,
	0x9b, 0x97, 0xf7, 0x34, 0xfc, 0x3e, 0xa8, 0xe2, 0xc7, 0x93, 0xa0, 0xa5, 0x77, 0x38, 0xfe, 0x0a,
	0x00, 0x00, 0xff, 0xff, 0x1c, 0x2e, 0xc6, 0x97, 0x71, 0x02, 0x00, 0x00,
}

func (m *AuthenticatorAddedEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthenticatorAddedEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AuthenticatorAddedEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AuthenticatorId) > 0 {
		i -= len(m.AuthenticatorId)
		copy(dAtA[i:], m.AuthenticatorId)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.AuthenticatorId)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.AuthenticatorType) > 0 {
		i -= len(m.AuthenticatorType)
		copy(dAtA[i:], m.AuthenticatorType)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.AuthenticatorType)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AuthenticatorRemovedEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthenticatorRemovedEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AuthenticatorRemovedEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AuthenticatorId) > 0 {
		i -= len(m.AuthenticatorId)
		copy(dAtA[i:], m.AuthenticatorId)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.AuthenticatorId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *AuthenticatorSetActiveStateEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *AuthenticatorSetActiveStateEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *AuthenticatorSetActiveStateEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.IsSmartAccountActive {
		i--
		if m.IsSmartAccountActive {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if len(m.Sender) > 0 {
		i -= len(m.Sender)
		copy(dAtA[i:], m.Sender)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Sender)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintEvent(dAtA []byte, offset int, v uint64) int {
	offset -= sovEvent(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *AuthenticatorAddedEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.AuthenticatorType)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.AuthenticatorId)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	return n
}

func (m *AuthenticatorRemovedEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.AuthenticatorId)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	return n
}

func (m *AuthenticatorSetActiveStateEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Sender)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	if m.IsSmartAccountActive {
		n += 2
	}
	return n
}

func sovEvent(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvent(x uint64) (n int) {
	return sovEvent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *AuthenticatorAddedEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: AuthenticatorAddedEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthenticatorAddedEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthenticatorType", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AuthenticatorType = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthenticatorId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AuthenticatorId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func (m *AuthenticatorRemovedEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: AuthenticatorRemovedEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthenticatorRemovedEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthenticatorId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.AuthenticatorId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func (m *AuthenticatorSetActiveStateEvent) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowEvent
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
			return fmt.Errorf("proto: AuthenticatorSetActiveStateEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: AuthenticatorSetActiveStateEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Sender", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Sender = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IsSmartAccountActive", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
			m.IsSmartAccountActive = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipEvent(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthEvent
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
func skipEvent(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowEvent
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
					return 0, ErrIntOverflowEvent
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
					return 0, ErrIntOverflowEvent
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
				return 0, ErrInvalidLengthEvent
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupEvent
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthEvent
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthEvent        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowEvent          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupEvent = fmt.Errorf("proto: unexpected end of group")
)