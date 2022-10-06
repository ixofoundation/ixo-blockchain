// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: iid/event.proto

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

// DidDocumentCreatedEvent is an event triggered on a DID document creation
type IidDocumentCreatedEvent struct {
	// the did being created
	Did string `protobuf:"bytes,1,opt,name=did,proto3" json:"did,omitempty"`
	// the signer account creating the did
	Signer string `protobuf:"bytes,2,opt,name=signer,proto3" json:"signer,omitempty"`
}

func (m *IidDocumentCreatedEvent) Reset()         { *m = IidDocumentCreatedEvent{} }
func (m *IidDocumentCreatedEvent) String() string { return proto.CompactTextString(m) }
func (*IidDocumentCreatedEvent) ProtoMessage()    {}
func (*IidDocumentCreatedEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d285af057a58eb1, []int{0}
}
func (m *IidDocumentCreatedEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IidDocumentCreatedEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IidDocumentCreatedEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IidDocumentCreatedEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IidDocumentCreatedEvent.Merge(m, src)
}
func (m *IidDocumentCreatedEvent) XXX_Size() int {
	return m.Size()
}
func (m *IidDocumentCreatedEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_IidDocumentCreatedEvent.DiscardUnknown(m)
}

var xxx_messageInfo_IidDocumentCreatedEvent proto.InternalMessageInfo

// DidDocumentUpdatedEvent is an event triggered on a DID document update
type IidDocumentUpdatedEvent struct {
	// the did being updated
	Did string `protobuf:"bytes,1,opt,name=did,proto3" json:"did,omitempty"`
	// the signer account of the change
	Signer string `protobuf:"bytes,2,opt,name=signer,proto3" json:"signer,omitempty"`
}

func (m *IidDocumentUpdatedEvent) Reset()         { *m = IidDocumentUpdatedEvent{} }
func (m *IidDocumentUpdatedEvent) String() string { return proto.CompactTextString(m) }
func (*IidDocumentUpdatedEvent) ProtoMessage()    {}
func (*IidDocumentUpdatedEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d285af057a58eb1, []int{1}
}
func (m *IidDocumentUpdatedEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *IidDocumentUpdatedEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_IidDocumentUpdatedEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *IidDocumentUpdatedEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IidDocumentUpdatedEvent.Merge(m, src)
}
func (m *IidDocumentUpdatedEvent) XXX_Size() int {
	return m.Size()
}
func (m *IidDocumentUpdatedEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_IidDocumentUpdatedEvent.DiscardUnknown(m)
}

var xxx_messageInfo_IidDocumentUpdatedEvent proto.InternalMessageInfo

func init() {
	proto.RegisterType((*IidDocumentCreatedEvent)(nil), "iid.IidDocumentCreatedEvent")
	proto.RegisterType((*IidDocumentUpdatedEvent)(nil), "iid.IidDocumentUpdatedEvent")
}

func init() { proto.RegisterFile("iid/event.proto", fileDescriptor_0d285af057a58eb1) }

var fileDescriptor_0d285af057a58eb1 = []byte{
	// 215 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0xcf, 0xcc, 0x4c, 0xd1,
	0x4f, 0x2d, 0x4b, 0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0xce, 0xcc, 0x4c,
	0x91, 0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0xf3, 0xf5, 0x41, 0x2c, 0x88, 0x94, 0x92, 0x2f, 0x97,
	0xb8, 0x67, 0x66, 0x8a, 0x4b, 0x7e, 0x72, 0x69, 0x6e, 0x6a, 0x5e, 0x89, 0x73, 0x51, 0x6a, 0x62,
	0x49, 0x6a, 0x8a, 0x2b, 0x48, 0xaf, 0x90, 0x00, 0x17, 0x73, 0x4a, 0x66, 0x8a, 0x04, 0xa3, 0x02,
	0xa3, 0x06, 0x67, 0x10, 0x88, 0x29, 0x24, 0xc6, 0xc5, 0x56, 0x9c, 0x99, 0x9e, 0x97, 0x5a, 0x24,
	0xc1, 0x04, 0x16, 0x84, 0xf2, 0xac, 0x38, 0x3a, 0x16, 0xc8, 0x33, 0xbc, 0x58, 0x20, 0xcf, 0x88,
	0x66, 0x5c, 0x68, 0x41, 0x0a, 0x45, 0xc6, 0x39, 0xf9, 0x9e, 0x78, 0x24, 0xc7, 0x78, 0xe1, 0x91,
	0x1c, 0xe3, 0x83, 0x47, 0x72, 0x8c, 0x13, 0x1e, 0xcb, 0x31, 0x5c, 0x78, 0x2c, 0xc7, 0x70, 0xe3,
	0xb1, 0x1c, 0x43, 0x94, 0x71, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae, 0x7e,
	0x66, 0x45, 0x7e, 0x5a, 0x7e, 0x69, 0x5e, 0x4a, 0x62, 0x49, 0x66, 0x7e, 0x1e, 0x88, 0xa7, 0x9b,
	0x94, 0x93, 0x9f, 0x9c, 0x9d, 0x9c, 0x91, 0x98, 0x99, 0xa7, 0x5f, 0xa1, 0x0f, 0x0a, 0x8d, 0x92,
	0xca, 0x82, 0xd4, 0xe2, 0x24, 0x36, 0xb0, 0x9f, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0x4d,
	0x4e, 0xd7, 0xdd, 0x21, 0x01, 0x00, 0x00,
}

func (this *IidDocumentCreatedEvent) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*IidDocumentCreatedEvent)
	if !ok {
		that2, ok := that.(IidDocumentCreatedEvent)
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
	if this.Did != that1.Did {
		return false
	}
	if this.Signer != that1.Signer {
		return false
	}
	return true
}
func (this *IidDocumentUpdatedEvent) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*IidDocumentUpdatedEvent)
	if !ok {
		that2, ok := that.(IidDocumentUpdatedEvent)
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
	if this.Did != that1.Did {
		return false
	}
	if this.Signer != that1.Signer {
		return false
	}
	return true
}
func (m *IidDocumentCreatedEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IidDocumentCreatedEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IidDocumentCreatedEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
	if len(m.Did) > 0 {
		i -= len(m.Did)
		copy(dAtA[i:], m.Did)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Did)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *IidDocumentUpdatedEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *IidDocumentUpdatedEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *IidDocumentUpdatedEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
	if len(m.Did) > 0 {
		i -= len(m.Did)
		copy(dAtA[i:], m.Did)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Did)))
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
func (m *IidDocumentCreatedEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Did)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.Signer)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	return n
}

func (m *IidDocumentUpdatedEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Did)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.Signer)
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
func (m *IidDocumentCreatedEvent) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: IidDocumentCreatedEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IidDocumentCreatedEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Did", wireType)
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
			m.Did = string(dAtA[iNdEx:postIndex])
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
func (m *IidDocumentUpdatedEvent) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: IidDocumentUpdatedEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: IidDocumentUpdatedEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Did", wireType)
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
			m.Did = string(dAtA[iNdEx:postIndex])
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
