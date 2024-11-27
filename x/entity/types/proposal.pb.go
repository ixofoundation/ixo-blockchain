// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ixo/entity/v1beta1/proposal.proto

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

type InitializeNftContract struct {
	NftContractCodeId uint64 `protobuf:"varint,1,opt,name=NftContractCodeId,proto3" json:"NftContractCodeId,omitempty"`
	NftMinterAddress  string `protobuf:"bytes,2,opt,name=NftMinterAddress,proto3" json:"NftMinterAddress,omitempty"`
}

func (m *InitializeNftContract) Reset()         { *m = InitializeNftContract{} }
func (m *InitializeNftContract) String() string { return proto.CompactTextString(m) }
func (*InitializeNftContract) ProtoMessage()    {}
func (*InitializeNftContract) Descriptor() ([]byte, []int) {
	return fileDescriptor_460dff4269aaf83b, []int{0}
}
func (m *InitializeNftContract) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *InitializeNftContract) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_InitializeNftContract.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *InitializeNftContract) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InitializeNftContract.Merge(m, src)
}
func (m *InitializeNftContract) XXX_Size() int {
	return m.Size()
}
func (m *InitializeNftContract) XXX_DiscardUnknown() {
	xxx_messageInfo_InitializeNftContract.DiscardUnknown(m)
}

var xxx_messageInfo_InitializeNftContract proto.InternalMessageInfo

func (m *InitializeNftContract) GetNftContractCodeId() uint64 {
	if m != nil {
		return m.NftContractCodeId
	}
	return 0
}

func (m *InitializeNftContract) GetNftMinterAddress() string {
	if m != nil {
		return m.NftMinterAddress
	}
	return ""
}

func init() {
	proto.RegisterType((*InitializeNftContract)(nil), "ixo.entity.v1beta1.InitializeNftContract")
}

func init() { proto.RegisterFile("ixo/entity/v1beta1/proposal.proto", fileDescriptor_460dff4269aaf83b) }

var fileDescriptor_460dff4269aaf83b = []byte{
	// 241 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x90, 0xbd, 0x4a, 0x03, 0x41,
	0x10, 0x80, 0x6f, 0x45, 0x04, 0xaf, 0xd2, 0x43, 0x21, 0x58, 0x2c, 0xd1, 0x2a, 0x88, 0xde, 0x12,
	0xb4, 0xb1, 0xd4, 0x54, 0x29, 0x4c, 0x11, 0x3b, 0xbb, 0xfd, 0xbb, 0xcb, 0xe0, 0xb9, 0xb3, 0xee,
	0x4e, 0xc2, 0xc6, 0xa7, 0xf0, 0xb1, 0x2c, 0x53, 0x5a, 0xca, 0xdd, 0x8b, 0x48, 0x12, 0x05, 0x21,
	0xdd, 0x7c, 0xf3, 0x7d, 0xcd, 0x4c, 0x7e, 0x0e, 0x09, 0x85, 0x75, 0x04, 0xb4, 0x14, 0x8b, 0xa1,
	0xb2, 0x24, 0x87, 0xc2, 0x07, 0xf4, 0x18, 0x65, 0x53, 0xfa, 0x80, 0x84, 0x45, 0x01, 0x09, 0xcb,
	0x6d, 0x52, 0xfe, 0x26, 0x67, 0x27, 0x35, 0xd6, 0xb8, 0xd1, 0x62, 0x3d, 0x6d, 0xcb, 0x8b, 0xb7,
	0xfc, 0x74, 0xec, 0x80, 0x40, 0x36, 0xf0, 0x6e, 0x27, 0x15, 0x8d, 0xd0, 0x51, 0x90, 0x9a, 0x8a,
	0xab, 0xfc, 0xf8, 0x1f, 0x8e, 0xd0, 0xd8, 0xb1, 0xe9, 0xb1, 0x3e, 0x1b, 0xec, 0x4f, 0x77, 0x45,
	0x71, 0x99, 0x1f, 0x4d, 0x2a, 0x7a, 0x04, 0x47, 0x36, 0xdc, 0x1b, 0x13, 0x6c, 0x8c, 0xbd, 0xbd,
	0x3e, 0x1b, 0x1c, 0x4e, 0x77, 0xf6, 0x0f, 0x4f, 0x9f, 0x2d, 0x67, 0xab, 0x96, 0xb3, 0xef, 0x96,
	0xb3, 0x8f, 0x8e, 0x67, 0xab, 0x8e, 0x67, 0x5f, 0x1d, 0xcf, 0x9e, 0xef, 0x6a, 0xa0, 0xd9, 0x5c,
	0x95, 0x1a, 0x5f, 0x05, 0x24, 0xac, 0x70, 0xee, 0x8c, 0x24, 0x40, 0xb7, 0xa6, 0x6b, 0xd5, 0xa0,
	0x7e, 0xd1, 0x33, 0x09, 0x4e, 0x2c, 0x6e, 0x45, 0xfa, 0x7b, 0x01, 0x2d, 0xbd, 0x8d, 0xea, 0x60,
	0x73, 0xce, 0xcd, 0x4f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf1, 0x5f, 0x98, 0x3d, 0x1d, 0x01, 0x00,
	0x00,
}

func (m *InitializeNftContract) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *InitializeNftContract) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *InitializeNftContract) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.NftMinterAddress) > 0 {
		i -= len(m.NftMinterAddress)
		copy(dAtA[i:], m.NftMinterAddress)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.NftMinterAddress)))
		i--
		dAtA[i] = 0x12
	}
	if m.NftContractCodeId != 0 {
		i = encodeVarintProposal(dAtA, i, uint64(m.NftContractCodeId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintProposal(dAtA []byte, offset int, v uint64) int {
	offset -= sovProposal(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *InitializeNftContract) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.NftContractCodeId != 0 {
		n += 1 + sovProposal(uint64(m.NftContractCodeId))
	}
	l = len(m.NftMinterAddress)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	return n
}

func sovProposal(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProposal(x uint64) (n int) {
	return sovProposal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *InitializeNftContract) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposal
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
			return fmt.Errorf("proto: InitializeNftContract: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: InitializeNftContract: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NftContractCodeId", wireType)
			}
			m.NftContractCodeId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NftContractCodeId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field NftMinterAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.NftMinterAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposal
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
func skipProposal(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProposal
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
					return 0, ErrIntOverflowProposal
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
					return 0, ErrIntOverflowProposal
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
				return 0, ErrInvalidLengthProposal
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProposal
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProposal
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProposal        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProposal          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProposal = fmt.Errorf("proto: unexpected end of group")
)
