// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ixo/iid/v1beta1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
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

// GenesisState defines the did module's genesis state.
type GenesisState struct {
	IidDocs []IidDocument `protobuf:"bytes,1,rep,name=iid_docs,json=iidDocs,proto3" json:"iid_docs" yaml:"iid_docs"`
	IidMeta []IidMetadata `protobuf:"bytes,2,rep,name=iid_meta,json=iidMeta,proto3" json:"iid_meta" yaml:"iid_meta"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_da4523a0526d3e8e, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetIidDocs() []IidDocument {
	if m != nil {
		return m.IidDocs
	}
	return nil
}

func (m *GenesisState) GetIidMeta() []IidMetadata {
	if m != nil {
		return m.IidMeta
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "ixo.iid.v1beta1.GenesisState")
}

func init() { proto.RegisterFile("ixo/iid/v1beta1/genesis.proto", fileDescriptor_da4523a0526d3e8e) }

var fileDescriptor_da4523a0526d3e8e = []byte{
	// 286 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x74, 0x90, 0xc1, 0x4a, 0xf3, 0x40,
	0x10, 0xc7, 0x93, 0xef, 0x03, 0x95, 0x28, 0x14, 0x8a, 0x60, 0x2d, 0xba, 0x95, 0x9c, 0xbc, 0xb8,
	0x4b, 0xed, 0xcd, 0x63, 0x11, 0xc4, 0x43, 0x41, 0xf4, 0xe6, 0x45, 0x36, 0xd9, 0xed, 0x76, 0x30,
	0xd9, 0x29, 0x66, 0x22, 0xc9, 0x5b, 0xf8, 0x2c, 0x3e, 0x45, 0x8f, 0x3d, 0x7a, 0x2a, 0x92, 0xbc,
	0x81, 0x4f, 0x20, 0xd9, 0x98, 0x4b, 0xd1, 0xdb, 0x0e, 0xbf, 0xff, 0xfc, 0x66, 0xf9, 0x07, 0xa7,
	0x50, 0xa0, 0x00, 0x50, 0xe2, 0x75, 0x1c, 0x69, 0x92, 0x63, 0x61, 0xb4, 0xd5, 0x19, 0x64, 0x7c,
	0xf9, 0x82, 0x84, 0xfd, 0x1e, 0x14, 0xc8, 0x01, 0x14, 0xff, 0xc1, 0xc3, 0x43, 0x83, 0x06, 0x1d,
	0x13, 0xcd, 0xab, 0x8d, 0x0d, 0x8f, 0x0d, 0xa2, 0x49, 0xb4, 0x70, 0x53, 0x94, 0xcf, 0x85, 0xb4,
	0x65, 0x87, 0xb6, 0x0f, 0x34, 0x36, 0x87, 0xc2, 0x77, 0x3f, 0x38, 0xb8, 0x69, 0xcf, 0x3d, 0x90,
	0x24, 0xdd, 0xbf, 0x0b, 0xf6, 0x00, 0xd4, 0x93, 0xc2, 0x38, 0x1b, 0xf8, 0x67, 0xff, 0xcf, 0xf7,
	0x2f, 0x4f, 0xf8, 0xd6, 0x07, 0xf8, 0x2d, 0xa8, 0x6b, 0x8c, 0xf3, 0x54, 0x5b, 0x9a, 0x1e, 0xad,
	0x36, 0x23, 0xef, 0x6b, 0x33, 0xea, 0x95, 0x32, 0x4d, 0xae, 0xc2, 0x6e, 0x37, 0xbc, 0xdf, 0x05,
	0x97, 0xca, 0x3a, 0x63, 0xaa, 0x49, 0x0e, 0xfe, 0xfd, 0x6d, 0x9c, 0x69, 0x92, 0x4a, 0x92, 0xfc,
	0xcd, 0xd8, 0xec, 0xb6, 0xc6, 0x26, 0x35, 0x9d, 0xad, 0x2a, 0xe6, 0xaf, 0x2b, 0xe6, 0x7f, 0x56,
	0xcc, 0x7f, 0xab, 0x99, 0xb7, 0xae, 0x99, 0xf7, 0x51, 0x33, 0xef, 0x71, 0x62, 0x80, 0x16, 0x79,
	0xc4, 0x63, 0x4c, 0x05, 0x14, 0x38, 0xc7, 0xdc, 0x2a, 0x49, 0x80, 0xb6, 0x99, 0x2e, 0xa2, 0x04,
	0xe3, 0xe7, 0x78, 0x21, 0xc1, 0x8a, 0xc2, 0xf5, 0x41, 0xe5, 0x52, 0x67, 0xd1, 0x8e, 0xab, 0x62,
	0xf2, 0x1d, 0x00, 0x00, 0xff, 0xff, 0x38, 0x9a, 0xd7, 0x11, 0x88, 0x01, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.IidMeta) > 0 {
		for iNdEx := len(m.IidMeta) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.IidMeta[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.IidDocs) > 0 {
		for iNdEx := len(m.IidDocs) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.IidDocs[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.IidDocs) > 0 {
		for _, e := range m.IidDocs {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.IidMeta) > 0 {
		for _, e := range m.IidMeta {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IidDocs", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IidDocs = append(m.IidDocs, IidDocument{})
			if err := m.IidDocs[len(m.IidDocs)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IidMeta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IidMeta = append(m.IidMeta, IidMetadata{})
			if err := m.IidMeta[len(m.IidMeta)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
