// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ixo/claims/v1beta1/genesis.proto

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

// GenesisState defines the claims module's genesis state.
type GenesisState struct {
	Params      Params       `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	Collections []Collection `protobuf:"bytes,2,rep,name=collections,proto3" json:"collections"`
	Claims      []Claim      `protobuf:"bytes,3,rep,name=claims,proto3" json:"claims"`
	Disputes    []Dispute    `protobuf:"bytes,4,rep,name=disputes,proto3" json:"disputes"`
	Intents     []Intent     `protobuf:"bytes,5,rep,name=intents,proto3" json:"intents"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_43a260806dd06159, []int{0}
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

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetCollections() []Collection {
	if m != nil {
		return m.Collections
	}
	return nil
}

func (m *GenesisState) GetClaims() []Claim {
	if m != nil {
		return m.Claims
	}
	return nil
}

func (m *GenesisState) GetDisputes() []Dispute {
	if m != nil {
		return m.Disputes
	}
	return nil
}

func (m *GenesisState) GetIntents() []Intent {
	if m != nil {
		return m.Intents
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "ixo.claims.v1beta1.GenesisState")
}

func init() { proto.RegisterFile("ixo/claims/v1beta1/genesis.proto", fileDescriptor_43a260806dd06159) }

var fileDescriptor_43a260806dd06159 = []byte{
	// 317 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x4f, 0x4f, 0x32, 0x31,
	0x10, 0x87, 0x77, 0x81, 0x97, 0xd7, 0x14, 0x4f, 0x8d, 0x87, 0x15, 0x93, 0x42, 0x3c, 0x71, 0x71,
	0x1b, 0xe0, 0xe0, 0x9f, 0xc4, 0x0b, 0x1a, 0x8d, 0x37, 0x23, 0x37, 0x6f, 0xdd, 0x52, 0x97, 0xc6,
	0xa5, 0xb3, 0xa1, 0x85, 0xac, 0xdf, 0xc2, 0x0f, 0xe5, 0x81, 0x23, 0x47, 0x4f, 0xc6, 0xc0, 0x17,
	0x31, 0xdb, 0x16, 0x62, 0xe2, 0xde, 0x3a, 0x99, 0xe7, 0xf9, 0x75, 0x26, 0x83, 0xba, 0xb2, 0x00,
	0xca, 0x33, 0x26, 0x67, 0x9a, 0x2e, 0xfb, 0x89, 0x30, 0xac, 0x4f, 0x53, 0xa1, 0x84, 0x96, 0x3a,
	0xce, 0xe7, 0x60, 0x00, 0x63, 0x59, 0x40, 0xec, 0x88, 0xd8, 0x13, 0xed, 0xa3, 0x14, 0x52, 0xb0,
	0x6d, 0x5a, 0xbe, 0x1c, 0xd9, 0xee, 0x54, 0x64, 0x79, 0xd1, 0x02, 0xa7, 0x1f, 0x35, 0x74, 0x78,
	0xef, 0xc2, 0xc7, 0x86, 0x19, 0x81, 0x2f, 0x50, 0x33, 0x67, 0x73, 0x36, 0xd3, 0x51, 0xd8, 0x0d,
	0x7b, 0xad, 0x41, 0x3b, 0xfe, 0xfb, 0x59, 0xfc, 0x68, 0x89, 0x51, 0x63, 0xf5, 0xd5, 0x09, 0x9e,
	0x3c, 0x8f, 0xef, 0x50, 0x8b, 0x43, 0x96, 0x09, 0x6e, 0x24, 0x28, 0x1d, 0xd5, 0xba, 0xf5, 0x5e,
	0x6b, 0x40, 0xaa, 0xf4, 0x9b, 0x3d, 0xe6, 0x23, 0x7e, 0x8b, 0xf8, 0x1c, 0x35, 0x1d, 0x1f, 0xd5,
	0x6d, 0xc4, 0x71, 0x65, 0x44, 0x59, 0xee, 0x06, 0x70, 0x3d, 0x7c, 0x8d, 0x0e, 0x26, 0x52, 0xe7,
	0x0b, 0x23, 0x74, 0xd4, 0xb0, 0xea, 0x49, 0x95, 0x7a, 0xeb, 0x18, 0x2f, 0xef, 0x15, 0x7c, 0x85,
	0xfe, 0x4b, 0x65, 0x84, 0x32, 0x3a, 0xfa, 0x67, 0xed, 0xca, 0xd5, 0x1f, 0x2c, 0xe2, 0xe5, 0x9d,
	0x30, 0x1a, 0xaf, 0x36, 0x24, 0x5c, 0x6f, 0x48, 0xf8, 0xbd, 0x21, 0xe1, 0xfb, 0x96, 0x04, 0xeb,
	0x2d, 0x09, 0x3e, 0xb7, 0x24, 0x78, 0xbe, 0x4c, 0xa5, 0x99, 0x2e, 0x92, 0x98, 0xc3, 0x8c, 0xca,
	0x02, 0x5e, 0x60, 0xa1, 0x26, 0xac, 0x5c, 0xb4, 0xac, 0xce, 0x92, 0x0c, 0xf8, 0x2b, 0x9f, 0x32,
	0xa9, 0xe8, 0x72, 0x48, 0x8b, 0xdd, 0xa9, 0xcc, 0x5b, 0x2e, 0x74, 0xd2, 0xb4, 0x27, 0x1a, 0xfe,
	0x04, 0x00, 0x00, 0xff, 0xff, 0x37, 0x3a, 0x81, 0xa5, 0x11, 0x02, 0x00, 0x00,
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
	if len(m.Intents) > 0 {
		for iNdEx := len(m.Intents) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Intents[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.Disputes) > 0 {
		for iNdEx := len(m.Disputes) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Disputes[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Claims) > 0 {
		for iNdEx := len(m.Claims) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Claims[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Collections) > 0 {
		for iNdEx := len(m.Collections) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Collections[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
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
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.Collections) > 0 {
		for _, e := range m.Collections {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Claims) > 0 {
		for _, e := range m.Claims {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Disputes) > 0 {
		for _, e := range m.Disputes {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Intents) > 0 {
		for _, e := range m.Intents {
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
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
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
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Collections", wireType)
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
			m.Collections = append(m.Collections, Collection{})
			if err := m.Collections[len(m.Collections)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Claims", wireType)
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
			m.Claims = append(m.Claims, Claim{})
			if err := m.Claims[len(m.Claims)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Disputes", wireType)
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
			m.Disputes = append(m.Disputes, Dispute{})
			if err := m.Disputes[len(m.Disputes)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Intents", wireType)
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
			m.Intents = append(m.Intents, Intent{})
			if err := m.Intents[len(m.Intents)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
