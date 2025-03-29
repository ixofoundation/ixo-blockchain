// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ixo/mint/v1beta1/event.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// MintEpochProvisionsMintedEvent is triggered after a epoch is triggered
// minting module for inflation.
type MintEpochProvisionsMintedEvent struct {
	EpochNumber     string `protobuf:"bytes,1,opt,name=epoch_number,json=epochNumber,proto3" json:"epoch_number,omitempty"`
	EpochProvisions string `protobuf:"bytes,2,opt,name=epoch_provisions,json=epochProvisions,proto3" json:"epoch_provisions,omitempty"`
	Amount          string `protobuf:"bytes,3,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (m *MintEpochProvisionsMintedEvent) Reset()         { *m = MintEpochProvisionsMintedEvent{} }
func (m *MintEpochProvisionsMintedEvent) String() string { return proto.CompactTextString(m) }
func (*MintEpochProvisionsMintedEvent) ProtoMessage()    {}
func (*MintEpochProvisionsMintedEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_e43ebec99bf5eeaf, []int{0}
}
func (m *MintEpochProvisionsMintedEvent) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MintEpochProvisionsMintedEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MintEpochProvisionsMintedEvent.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MintEpochProvisionsMintedEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MintEpochProvisionsMintedEvent.Merge(m, src)
}
func (m *MintEpochProvisionsMintedEvent) XXX_Size() int {
	return m.Size()
}
func (m *MintEpochProvisionsMintedEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_MintEpochProvisionsMintedEvent.DiscardUnknown(m)
}

var xxx_messageInfo_MintEpochProvisionsMintedEvent proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MintEpochProvisionsMintedEvent)(nil), "ixo.mint.v1beta1.MintEpochProvisionsMintedEvent")
}

func init() { proto.RegisterFile("ixo/mint/v1beta1/event.proto", fileDescriptor_e43ebec99bf5eeaf) }

var fileDescriptor_e43ebec99bf5eeaf = []byte{
	// 263 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xc9, 0xac, 0xc8, 0xd7,
	0xcf, 0xcd, 0xcc, 0x2b, 0xd1, 0x2f, 0x33, 0x4c, 0x4a, 0x2d, 0x49, 0x34, 0xd4, 0x4f, 0x2d, 0x4b,
	0xcd, 0x2b, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x12, 0xc8, 0xac, 0xc8, 0xd7, 0x03, 0xc9,
	0xea, 0x41, 0x65, 0xa5, 0x44, 0xd2, 0xf3, 0xd3, 0xf3, 0xc1, 0x92, 0xfa, 0x20, 0x16, 0x44, 0x9d,
	0xd2, 0x04, 0x46, 0x2e, 0x39, 0xdf, 0xcc, 0xbc, 0x12, 0xd7, 0x82, 0xfc, 0xe4, 0x8c, 0x80, 0xa2,
	0xfc, 0xb2, 0xcc, 0xe2, 0xcc, 0xfc, 0xbc, 0x62, 0x90, 0x50, 0x6a, 0x8a, 0x2b, 0xc8, 0x40, 0x21,
	0x45, 0x2e, 0x9e, 0x54, 0x90, 0x6c, 0x7c, 0x5e, 0x69, 0x6e, 0x52, 0x6a, 0x91, 0x04, 0xa3, 0x02,
	0xa3, 0x06, 0x67, 0x10, 0x37, 0x58, 0xcc, 0x0f, 0x2c, 0x24, 0xa4, 0xc9, 0x25, 0x00, 0x51, 0x52,
	0x00, 0x37, 0x41, 0x82, 0x09, 0xac, 0x8c, 0x3f, 0x15, 0xd5, 0x60, 0x21, 0x31, 0x2e, 0xb6, 0xc4,
	0xdc, 0xfc, 0xd2, 0xbc, 0x12, 0x09, 0x66, 0xb0, 0x02, 0x28, 0xcf, 0x8a, 0xa3, 0x63, 0x81, 0x3c,
	0xc3, 0x8b, 0x05, 0xf2, 0x0c, 0x4e, 0x81, 0x27, 0x1e, 0xc9, 0x31, 0x5e, 0x78, 0x24, 0xc7, 0xf8,
	0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31, 0xdc, 0x78, 0x2c, 0xc7,
	0x10, 0x65, 0x9e, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f, 0xab, 0x9f, 0x59, 0x91,
	0x9f, 0x96, 0x5f, 0x9a, 0x97, 0x92, 0x58, 0x92, 0x99, 0x9f, 0x07, 0xe2, 0xe9, 0x26, 0xe5, 0xe4,
	0x27, 0x67, 0x27, 0x67, 0x24, 0x66, 0xe6, 0xe9, 0x97, 0x99, 0xea, 0x57, 0x40, 0xc2, 0xa6, 0xa4,
	0xb2, 0x20, 0xb5, 0x38, 0x89, 0x0d, 0xec, 0x59, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0x99,
	0x53, 0x0c, 0x92, 0x34, 0x01, 0x00, 0x00,
}

func (m *MintEpochProvisionsMintedEvent) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MintEpochProvisionsMintedEvent) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MintEpochProvisionsMintedEvent) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Amount) > 0 {
		i -= len(m.Amount)
		copy(dAtA[i:], m.Amount)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.Amount)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.EpochProvisions) > 0 {
		i -= len(m.EpochProvisions)
		copy(dAtA[i:], m.EpochProvisions)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.EpochProvisions)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.EpochNumber) > 0 {
		i -= len(m.EpochNumber)
		copy(dAtA[i:], m.EpochNumber)
		i = encodeVarintEvent(dAtA, i, uint64(len(m.EpochNumber)))
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
func (m *MintEpochProvisionsMintedEvent) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.EpochNumber)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.EpochProvisions)
	if l > 0 {
		n += 1 + l + sovEvent(uint64(l))
	}
	l = len(m.Amount)
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
func (m *MintEpochProvisionsMintedEvent) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MintEpochProvisionsMintedEvent: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MintEpochProvisionsMintedEvent: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochNumber", wireType)
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
			m.EpochNumber = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochProvisions", wireType)
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
			m.EpochProvisions = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
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
			m.Amount = string(dAtA[iNdEx:postIndex])
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
