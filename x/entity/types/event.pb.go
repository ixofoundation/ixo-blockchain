// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ixo/entity/v1beta1/event.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// EntityCreatedEvent is an event triggered on a Entity creation
type EntityCreatedEvent struct {
	Entity *Entity `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Signer string  `protobuf:"bytes,2,opt,name=signer,proto3" json:"signer,omitempty"`
}

func (m *EntityCreatedEvent) Reset()         { *m = EntityCreatedEvent{} }
func (m *EntityCreatedEvent) String() string { return proto.CompactTextString(m) }
func (*EntityCreatedEvent) ProtoMessage()    {}
func (*EntityCreatedEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a73923d9783a1ef, []int{0}
}
func (m *EntityCreatedEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EntityCreatedEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EntityCreatedEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EntityCreatedEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntityCreatedEvent.Merge(m, src)
}
func (m *EntityCreatedEvent) XXX_Size() int {
	return m.Size()
}
func (m *EntityCreatedEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_EntityCreatedEvent.DiscardUnknown(m)
}

var xxx_messageInfo_EntityCreatedEvent proto.InternalMessageInfo

// EntityUpdatedEvent is an event triggered on a entity document update
type EntityUpdatedEvent struct {
	Entity *Entity `protobuf:"bytes,1,opt,name=entity,proto3" json:"entity,omitempty"`
	Signer string  `protobuf:"bytes,2,opt,name=signer,proto3" json:"signer,omitempty"`
}

func (m *EntityUpdatedEvent) Reset()         { *m = EntityUpdatedEvent{} }
func (m *EntityUpdatedEvent) String() string { return proto.CompactTextString(m) }
func (*EntityUpdatedEvent) ProtoMessage()    {}
func (*EntityUpdatedEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a73923d9783a1ef, []int{1}
}
func (m *EntityUpdatedEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EntityUpdatedEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EntityUpdatedEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EntityUpdatedEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntityUpdatedEvent.Merge(m, src)
}
func (m *EntityUpdatedEvent) XXX_Size() int {
	return m.Size()
}
func (m *EntityUpdatedEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_EntityUpdatedEvent.DiscardUnknown(m)
}

var xxx_messageInfo_EntityUpdatedEvent proto.InternalMessageInfo

// EntityVerifiedUpdatedEvent is an event triggered on a entity verified
// document update
type EntityVerifiedUpdatedEvent struct {
	Id             string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Signer         string `protobuf:"bytes,2,opt,name=signer,proto3" json:"signer,omitempty"`
	EntityVerified bool   `protobuf:"varint,3,opt,name=entity_verified,json=entityVerified,proto3" json:"entity_verified,omitempty"`
}

func (m *EntityVerifiedUpdatedEvent) Reset()         { *m = EntityVerifiedUpdatedEvent{} }
func (m *EntityVerifiedUpdatedEvent) String() string { return proto.CompactTextString(m) }
func (*EntityVerifiedUpdatedEvent) ProtoMessage()    {}
func (*EntityVerifiedUpdatedEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a73923d9783a1ef, []int{2}
}
func (m *EntityVerifiedUpdatedEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EntityVerifiedUpdatedEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EntityVerifiedUpdatedEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EntityVerifiedUpdatedEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntityVerifiedUpdatedEvent.Merge(m, src)
}
func (m *EntityVerifiedUpdatedEvent) XXX_Size() int {
	return m.Size()
}
func (m *EntityVerifiedUpdatedEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_EntityVerifiedUpdatedEvent.DiscardUnknown(m)
}

var xxx_messageInfo_EntityVerifiedUpdatedEvent proto.InternalMessageInfo

// EntityTransferredEvent is an event triggered on a entity transfer
type EntityTransferredEvent struct {
	Id   string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	From string `protobuf:"bytes,2,opt,name=from,proto3" json:"from,omitempty"`
	To   string `protobuf:"bytes,3,opt,name=to,proto3" json:"to,omitempty"`
}

func (m *EntityTransferredEvent) Reset()         { *m = EntityTransferredEvent{} }
func (m *EntityTransferredEvent) String() string { return proto.CompactTextString(m) }
func (*EntityTransferredEvent) ProtoMessage()    {}
func (*EntityTransferredEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_3a73923d9783a1ef, []int{3}
}
func (m *EntityTransferredEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *EntityTransferredEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_EntityTransferredEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *EntityTransferredEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntityTransferredEvent.Merge(m, src)
}
func (m *EntityTransferredEvent) XXX_Size() int {
	return m.Size()
}
func (m *EntityTransferredEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_EntityTransferredEvent.DiscardUnknown(m)
}

var xxx_messageInfo_EntityTransferredEvent proto.InternalMessageInfo

func init() {
	proto.RegisterType((*EntityCreatedEvent)(nil), "ixo.entity.v1beta1.EntityCreatedEvent")
	proto.RegisterType((*EntityUpdatedEvent)(nil), "ixo.entity.v1beta1.EntityUpdatedEvent")
	proto.RegisterType((*EntityVerifiedUpdatedEvent)(nil), "ixo.entity.v1beta1.EntityVerifiedUpdatedEvent")
	proto.RegisterType((*EntityTransferredEvent)(nil), "ixo.entity.v1beta1.EntityTransferredEvent")
}

func init() { proto.RegisterFile("ixo/entity/v1beta1/event.proto", fileDescriptor_3a73923d9783a1ef) }

var fileDescriptor_3a73923d9783a1ef = []byte{
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x92, 0x3f, 0x4f, 0xc2, 0x40,
	0x18, 0xc6, 0x7b, 0xd5, 0x10, 0x38, 0x13, 0x4c, 0x2e, 0x86, 0x90, 0x0e, 0x07, 0x71, 0x91, 0xc5,
	0x36, 0x60, 0xe2, 0xe0, 0xa8, 0x61, 0x37, 0x8d, 0x32, 0xb8, 0x98, 0x96, 0xbb, 0x96, 0x53, 0xb9,
	0x97, 0x1c, 0x07, 0x96, 0x6f, 0xe0, 0xe8, 0x47, 0xe0, 0xe3, 0x38, 0x32, 0x3a, 0x1a, 0xba, 0xf8,
	0x31, 0x4c, 0xef, 0x6a, 0x90, 0x28, 0xa3, 0xdb, 0xfb, 0xe4, 0xf9, 0xf3, 0x5b, 0x5e, 0x4c, 0x45,
	0x06, 0x01, 0x97, 0x5a, 0xe8, 0x45, 0x30, 0xef, 0xc6, 0x5c, 0x47, 0xdd, 0x80, 0xcf, 0xb9, 0xd4,
	0xfe, 0x44, 0x81, 0x06, 0x42, 0x44, 0x06, 0xbe, 0xf5, 0xfd, 0xd2, 0xf7, 0x8e, 0x52, 0x48, 0xc1,
	0xd8, 0x41, 0x71, 0xd9, 0xa4, 0xd7, 0xfa, 0x6b, 0xc9, 0x16, 0x4d, 0xe0, 0xf8, 0x01, 0x93, 0xbe,
	0xd1, 0x57, 0x8a, 0x47, 0x9a, 0xb3, 0x7e, 0x81, 0x21, 0x3d, 0x5c, 0xb1, 0xa9, 0x26, 0x6a, 0xa3,
	0xce, 0x41, 0xcf, 0xf3, 0x7f, 0x13, 0x7d, 0xdb, 0x0b, 0xcb, 0x24, 0x69, 0xe0, 0xca, 0x54, 0xa4,
	0x92, 0xab, 0xa6, 0xdb, 0x46, 0x9d, 0x5a, 0x58, 0xaa, 0x8b, 0xea, 0xcb, 0xb2, 0xe5, 0x7c, 0x2e,
	0x5b, 0x68, 0xc3, 0xba, 0x9d, 0xb0, 0xff, 0x66, 0x3d, 0x63, 0xcf, 0x76, 0x06, 0x5c, 0x89, 0x44,
	0x70, 0xb6, 0xc5, 0xac, 0x63, 0x57, 0x30, 0xc3, 0xab, 0x85, 0xae, 0x60, 0xbb, 0xf6, 0xc8, 0x09,
	0x3e, 0xb4, 0xc4, 0xfb, 0x79, 0x39, 0xd3, 0xdc, 0x6b, 0xa3, 0x4e, 0x35, 0xac, 0xf3, 0xad, 0xf1,
	0x1f, 0xe0, 0x01, 0x6e, 0x58, 0xf0, 0x8d, 0x8a, 0xe4, 0x34, 0xe1, 0x4a, 0xed, 0x82, 0x12, 0xbc,
	0x9f, 0x28, 0x18, 0x97, 0x48, 0x73, 0x17, 0x19, 0x0d, 0x86, 0x51, 0x0b, 0x5d, 0x0d, 0x9b, 0xdd,
	0xcb, 0xeb, 0xb7, 0x35, 0x45, 0xab, 0x35, 0x45, 0x1f, 0x6b, 0x8a, 0x5e, 0x73, 0xea, 0xac, 0x72,
	0xea, 0xbc, 0xe7, 0xd4, 0xb9, 0x3b, 0x4f, 0x85, 0x1e, 0xcd, 0x62, 0x7f, 0x08, 0xe3, 0x40, 0x64,
	0x90, 0xc0, 0x4c, 0xb2, 0x48, 0x0b, 0x90, 0x85, 0x3a, 0x8d, 0x9f, 0x60, 0xf8, 0x38, 0x1c, 0x45,
	0x42, 0x06, 0xd9, 0xf7, 0x27, 0xe8, 0xc5, 0x84, 0x4f, 0xe3, 0x8a, 0xf9, 0x80, 0xb3, 0xaf, 0x00,
	0x00, 0x00, 0xff, 0xff, 0x1f, 0xf2, 0x44, 0xec, 0x6e, 0x02, 0x00, 0x00,
}

func (this *EntityCreatedEvent) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*EntityCreatedEvent)
	if !ok {
		that2, ok := that.(EntityCreatedEvent)
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
	if !this.Entity.Equal(that1.Entity) {
		return false
	}
	if this.Signer != that1.Signer {
		return false
	}
	return true
}
func (this *EntityUpdatedEvent) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*EntityUpdatedEvent)
	if !ok {
		that2, ok := that.(EntityUpdatedEvent)
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
	if !this.Entity.Equal(that1.Entity) {
		return false
	}
	if this.Signer != that1.Signer {
		return false
	}
	return true
}
func (this *EntityVerifiedUpdatedEvent) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*EntityVerifiedUpdatedEvent)
	if !ok {
		that2, ok := that.(EntityVerifiedUpdatedEvent)
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
	if this.Id != that1.Id {
		return false
	}
	if this.Signer != that1.Signer {
		return false
	}
	if this.EntityVerified != that1.EntityVerified {
		return false
	}
	return true
}
func (this *EntityTransferredEvent) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*EntityTransferredEvent)
	if !ok {
		that2, ok := that.(EntityTransferredEvent)
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
	if this.Id != that1.Id {
		return false
	}
	if this.From != that1.From {
		return false
	}
	if this.To != that1.To {
		return false
	}
	return true
}
func (m *EntityCreatedEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EntityCreatedEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EntityCreatedEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signer) > 0 {
		i -= len(m.Signer)
		copy(dAtA[i:], m.Signer)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Signer)))
		i--
		dAtA[i] = 0x12
	}
	if m.Entity != nil {
		{
			size, err := m.Entity.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEvent(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EntityUpdatedEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EntityUpdatedEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EntityUpdatedEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signer) > 0 {
		i -= len(m.Signer)
		copy(dAtA[i:], m.Signer)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Signer)))
		i--
		dAtA[i] = 0x12
	}
	if m.Entity != nil {
		{
			size, err := m.Entity.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintEvent(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EntityVerifiedUpdatedEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EntityVerifiedUpdatedEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EntityVerifiedUpdatedEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.EntityVerified {
		i--
		if m.EntityVerified {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x18
	}
	if len(m.Signer) > 0 {
		i -= len(m.Signer)
		copy(dAtA[i:], m.Signer)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Signer)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *EntityTransferredEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *EntityTransferredEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *EntityTransferredEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.To) > 0 {
		i -= len(m.To)
		copy(dAtA[i:], m.To)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.To)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.From) > 0 {
		i -= len(m.From)
		copy(dAtA[i:], m.From)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.From)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Id)))
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
func (m *EntityCreatedEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Entity != nil {
		l = m.Entity.Size()
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.Signer)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	return n
}

func (m *EntityUpdatedEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Entity != nil {
		l = m.Entity.Size()
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.Signer)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	return n
}

func (m *EntityVerifiedUpdatedEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.Signer)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	if m.EntityVerified {
		n += 2
	}
	return n
}

func (m *EntityTransferredEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.From)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.To)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	return n
}

func sovEvent(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozEvent(x uint64) (n int) {
	return sovEvent(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *EntityCreatedEvent) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: EntityCreatedEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EntityCreatedEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Entity", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Entity == nil {
				m.Entity = &Entity{}
			}
			if err := m.Entity.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
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
			m.Signer = string(dAtA[iNdEx:postIndex])
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
func (m *EntityUpdatedEvent) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: EntityUpdatedEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EntityUpdatedEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Entity", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowEvent
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
				return ErrInvalidLengthEvent
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthEvent
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Entity == nil {
				m.Entity = &Entity{}
			}
			if err := m.Entity.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
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
			m.Signer = string(dAtA[iNdEx:postIndex])
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
func (m *EntityVerifiedUpdatedEvent) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: EntityVerifiedUpdatedEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EntityVerifiedUpdatedEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
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
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signer", wireType)
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
			m.Signer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field EntityVerified", wireType)
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
			m.EntityVerified = bool(v != 0)
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
func (m *EntityTransferredEvent) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: EntityTransferredEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: EntityTransferredEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
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
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field From", wireType)
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
			m.From = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field To", wireType)
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
			m.To = string(dAtA[iNdEx:postIndex])
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
