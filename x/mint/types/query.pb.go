// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ixo/mint/v1beta1/query.proto

package types

import (
	context "context"
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// QueryParamsRequest is the request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_249feb5acefc22ab, []int{0}
}
func (m *QueryParamsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsRequest.Merge(m, src)
}
func (m *QueryParamsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsRequest proto.InternalMessageInfo

// QueryParamsResponse is the response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// params defines the parameters of the module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_249feb5acefc22ab, []int{1}
}
func (m *QueryParamsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryParamsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryParamsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryParamsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParamsResponse.Merge(m, src)
}
func (m *QueryParamsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryParamsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParamsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParamsResponse proto.InternalMessageInfo

func (m *QueryParamsResponse) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

// QueryEpochProvisionsRequest is the request type for the
// Query/EpochProvisions RPC method.
type QueryEpochProvisionsRequest struct {
}

func (m *QueryEpochProvisionsRequest) Reset()         { *m = QueryEpochProvisionsRequest{} }
func (m *QueryEpochProvisionsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryEpochProvisionsRequest) ProtoMessage()    {}
func (*QueryEpochProvisionsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_249feb5acefc22ab, []int{2}
}
func (m *QueryEpochProvisionsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryEpochProvisionsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryEpochProvisionsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryEpochProvisionsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryEpochProvisionsRequest.Merge(m, src)
}
func (m *QueryEpochProvisionsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryEpochProvisionsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryEpochProvisionsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryEpochProvisionsRequest proto.InternalMessageInfo

// QueryEpochProvisionsResponse is the response type for the
// Query/EpochProvisions RPC method.
type QueryEpochProvisionsResponse struct {
	// epoch_provisions is the current minting per epoch provisions value.
	EpochProvisions cosmossdk_io_math.LegacyDec `protobuf:"bytes,1,opt,name=epoch_provisions,json=epochProvisions,proto3,customtype=cosmossdk.io/math.LegacyDec" json:"epoch_provisions"`
}

func (m *QueryEpochProvisionsResponse) Reset()         { *m = QueryEpochProvisionsResponse{} }
func (m *QueryEpochProvisionsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryEpochProvisionsResponse) ProtoMessage()    {}
func (*QueryEpochProvisionsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_249feb5acefc22ab, []int{3}
}
func (m *QueryEpochProvisionsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryEpochProvisionsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryEpochProvisionsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryEpochProvisionsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryEpochProvisionsResponse.Merge(m, src)
}
func (m *QueryEpochProvisionsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryEpochProvisionsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryEpochProvisionsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryEpochProvisionsResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "ixo.mint.v1beta1.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "ixo.mint.v1beta1.QueryParamsResponse")
	proto.RegisterType((*QueryEpochProvisionsRequest)(nil), "ixo.mint.v1beta1.QueryEpochProvisionsRequest")
	proto.RegisterType((*QueryEpochProvisionsResponse)(nil), "ixo.mint.v1beta1.QueryEpochProvisionsResponse")
}

func init() { proto.RegisterFile("ixo/mint/v1beta1/query.proto", fileDescriptor_249feb5acefc22ab) }

var fileDescriptor_249feb5acefc22ab = []byte{
	// 402 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0x41, 0x8b, 0xd3, 0x40,
	0x1c, 0xc5, 0x93, 0x45, 0x0b, 0x8e, 0x87, 0x5d, 0xc6, 0x3d, 0x94, 0xb4, 0x66, 0x97, 0x58, 0x41,
	0x84, 0xce, 0xd0, 0x8a, 0x7a, 0x2f, 0x7a, 0x53, 0x69, 0x7b, 0xf4, 0x22, 0x93, 0x74, 0x4c, 0x86,
	0x36, 0xf3, 0x4f, 0x33, 0x93, 0xda, 0x5e, 0xfd, 0x04, 0x82, 0x57, 0xc1, 0xaf, 0xd3, 0x63, 0xc1,
	0x8b, 0x78, 0x28, 0xd2, 0xfa, 0x41, 0x64, 0x26, 0x51, 0x68, 0x53, 0x71, 0x6f, 0xc9, 0xbc, 0xff,
	0xbc, 0xf7, 0x9b, 0xc7, 0x1f, 0xb5, 0xc5, 0x12, 0x68, 0x2a, 0xa4, 0xa6, 0x8b, 0x5e, 0xc8, 0x35,
	0xeb, 0xd1, 0x79, 0xc1, 0xf3, 0x15, 0xc9, 0x72, 0xd0, 0x80, 0x2f, 0xc4, 0x12, 0x88, 0x51, 0x49,
	0xa5, 0x7a, 0x97, 0x31, 0xc4, 0x60, 0x45, 0x6a, 0xbe, 0xca, 0x39, 0xaf, 0x1d, 0x03, 0xc4, 0x33,
	0x4e, 0x59, 0x26, 0x28, 0x93, 0x12, 0x34, 0xd3, 0x02, 0xa4, 0xaa, 0xd4, 0x56, 0x2d, 0xc3, 0x5a,
	0x5a, 0x31, 0xb8, 0x44, 0x78, 0x64, 0x12, 0x87, 0x2c, 0x67, 0xa9, 0x1a, 0xf3, 0x79, 0xc1, 0x95,
	0x0e, 0x5e, 0xa3, 0x7b, 0x07, 0xa7, 0x2a, 0x03, 0xa9, 0x38, 0x7e, 0x86, 0x1a, 0x99, 0x3d, 0x69,
	0xba, 0xd7, 0xee, 0xa3, 0xbb, 0xfd, 0x26, 0x39, 0x06, 0x24, 0xe5, 0x8d, 0xc1, 0xad, 0xf5, 0xf6,
	0xca, 0x19, 0x57, 0xd3, 0xc1, 0x7d, 0xd4, 0xb2, 0x76, 0x2f, 0x33, 0x88, 0x92, 0x61, 0x0e, 0x0b,
	0xa1, 0x0c, 0xdf, 0x9f, 0x34, 0x89, 0xda, 0xa7, 0xe5, 0x2a, 0xf6, 0x0d, 0xba, 0xe0, 0x46, 0x7a,
	0x97, 0xfd, 0xd5, 0x2c, 0xc0, 0x9d, 0xc1, 0x03, 0x13, 0xf3, 0x63, 0x7b, 0xd5, 0x8a, 0x40, 0xa5,
	0xa0, 0xd4, 0x64, 0x4a, 0x04, 0xd0, 0x94, 0xe9, 0x84, 0xbc, 0xe2, 0x31, 0x8b, 0x56, 0x2f, 0x78,
	0x34, 0x3e, 0xe7, 0x87, 0xbe, 0xfd, 0xaf, 0x67, 0xe8, 0xb6, 0x0d, 0xc4, 0x1f, 0x50, 0xa3, 0x04,
	0xc6, 0x9d, 0xfa, 0x53, 0xea, 0xbd, 0x78, 0x0f, 0xff, 0x33, 0x55, 0x02, 0x07, 0xd7, 0x1f, 0xbf,
	0xfd, 0xfa, 0x7c, 0xe6, 0xe1, 0x26, 0xad, 0x55, 0x5f, 0x36, 0x82, 0xbf, 0xb8, 0xe8, 0xfc, 0xe8,
	0xb9, 0xb8, 0xfb, 0x0f, 0xf3, 0xd3, 0xad, 0x79, 0xe4, 0xa6, 0xe3, 0x15, 0xd4, 0x63, 0x0b, 0xd5,
	0xc1, 0x41, 0x1d, 0xea, 0xb8, 0xdd, 0xc1, 0x68, 0xbd, 0xf3, 0xdd, 0xcd, 0xce, 0x77, 0x7f, 0xee,
	0x7c, 0xf7, 0xd3, 0xde, 0x77, 0x36, 0x7b, 0xdf, 0xf9, 0xbe, 0xf7, 0x9d, 0xb7, 0xcf, 0x63, 0xa1,
	0x93, 0x22, 0x24, 0x11, 0xa4, 0xc6, 0xe7, 0x3d, 0x14, 0x72, 0x62, 0xb7, 0xcd, 0xfc, 0x75, 0xc3,
	0x19, 0x44, 0xd3, 0x28, 0x61, 0x42, 0xd2, 0xc5, 0x53, 0xba, 0x2c, 0x53, 0xf4, 0x2a, 0xe3, 0x2a,
	0x6c, 0xd8, 0x7d, 0x7b, 0xf2, 0x3b, 0x00, 0x00, 0xff, 0xff, 0x95, 0xae, 0x51, 0x32, 0xf2, 0x02,
	0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
	// Params returns the total set of minting parameters.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	// EpochProvisions returns the current minting epoch provisions value.
	EpochProvisions(ctx context.Context, in *QueryEpochProvisionsRequest, opts ...grpc.CallOption) (*QueryEpochProvisionsResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/ixo.mint.v1beta1.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) EpochProvisions(ctx context.Context, in *QueryEpochProvisionsRequest, opts ...grpc.CallOption) (*QueryEpochProvisionsResponse, error) {
	out := new(QueryEpochProvisionsResponse)
	err := c.cc.Invoke(ctx, "/ixo.mint.v1beta1.Query/EpochProvisions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Params returns the total set of minting parameters.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	// EpochProvisions returns the current minting epoch provisions value.
	EpochProvisions(context.Context, *QueryEpochProvisionsRequest) (*QueryEpochProvisionsResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) EpochProvisions(ctx context.Context, req *QueryEpochProvisionsRequest) (*QueryEpochProvisionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EpochProvisions not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Params_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryParamsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Params(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ixo.mint.v1beta1.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_EpochProvisions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryEpochProvisionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).EpochProvisions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ixo.mint.v1beta1.Query/EpochProvisions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).EpochProvisions(ctx, req.(*QueryEpochProvisionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ixo.mint.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "EpochProvisions",
			Handler:    _Query_EpochProvisions_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ixo/mint/v1beta1/query.proto",
}

func (m *QueryParamsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryParamsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryParamsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryParamsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *QueryEpochProvisionsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryEpochProvisionsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryEpochProvisionsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *QueryEpochProvisionsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryEpochProvisionsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryEpochProvisionsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.EpochProvisions.Size()
		i -= size
		if _, err := m.EpochProvisions.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintQuery(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryParamsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryParamsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Params.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func (m *QueryEpochProvisionsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *QueryEpochProvisionsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.EpochProvisions.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryParamsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryParamsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryParamsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryParamsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryParamsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryEpochProvisionsRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryEpochProvisionsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryEpochProvisionsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func (m *QueryEpochProvisionsResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
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
			return fmt.Errorf("proto: QueryEpochProvisionsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryEpochProvisionsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field EpochProvisions", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
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
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.EpochProvisions.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
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
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
					return 0, ErrIntOverflowQuery
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
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)
