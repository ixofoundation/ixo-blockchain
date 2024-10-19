// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ixo/smartaccount/v1beta1/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
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

// QueryParamsRequest is request type for the Query/Params RPC method.
type QueryParamsRequest struct {
}

func (m *QueryParamsRequest) Reset()         { *m = QueryParamsRequest{} }
func (m *QueryParamsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryParamsRequest) ProtoMessage()    {}
func (*QueryParamsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_20645c5a3a5aaf72, []int{0}
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

// QueryParamsResponse is response type for the Query/Params RPC method.
type QueryParamsResponse struct {
	// params holds all the parameters of this module.
	Params Params `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
}

func (m *QueryParamsResponse) Reset()         { *m = QueryParamsResponse{} }
func (m *QueryParamsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryParamsResponse) ProtoMessage()    {}
func (*QueryParamsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_20645c5a3a5aaf72, []int{1}
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

// MsgGetAuthenticatorsRequest defines the Msg/GetAuthenticators request type.
type GetAuthenticatorsRequest struct {
	Account string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
}

func (m *GetAuthenticatorsRequest) Reset()         { *m = GetAuthenticatorsRequest{} }
func (m *GetAuthenticatorsRequest) String() string { return proto.CompactTextString(m) }
func (*GetAuthenticatorsRequest) ProtoMessage()    {}
func (*GetAuthenticatorsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_20645c5a3a5aaf72, []int{2}
}
func (m *GetAuthenticatorsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetAuthenticatorsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetAuthenticatorsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetAuthenticatorsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAuthenticatorsRequest.Merge(m, src)
}
func (m *GetAuthenticatorsRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetAuthenticatorsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAuthenticatorsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAuthenticatorsRequest proto.InternalMessageInfo

func (m *GetAuthenticatorsRequest) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

// MsgGetAuthenticatorsResponse defines the Msg/GetAuthenticators response type.
type GetAuthenticatorsResponse struct {
	AccountAuthenticators []*AccountAuthenticator `protobuf:"bytes,1,rep,name=account_authenticators,json=accountAuthenticators,proto3" json:"account_authenticators,omitempty"`
}

func (m *GetAuthenticatorsResponse) Reset()         { *m = GetAuthenticatorsResponse{} }
func (m *GetAuthenticatorsResponse) String() string { return proto.CompactTextString(m) }
func (*GetAuthenticatorsResponse) ProtoMessage()    {}
func (*GetAuthenticatorsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_20645c5a3a5aaf72, []int{3}
}
func (m *GetAuthenticatorsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetAuthenticatorsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetAuthenticatorsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetAuthenticatorsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAuthenticatorsResponse.Merge(m, src)
}
func (m *GetAuthenticatorsResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetAuthenticatorsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAuthenticatorsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAuthenticatorsResponse proto.InternalMessageInfo

func (m *GetAuthenticatorsResponse) GetAccountAuthenticators() []*AccountAuthenticator {
	if m != nil {
		return m.AccountAuthenticators
	}
	return nil
}

// MsgGetAuthenticatorRequest defines the Msg/GetAuthenticator request type.
type GetAuthenticatorRequest struct {
	Account         string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
	AuthenticatorId uint64 `protobuf:"varint,2,opt,name=authenticator_id,json=authenticatorId,proto3" json:"authenticator_id,omitempty"`
}

func (m *GetAuthenticatorRequest) Reset()         { *m = GetAuthenticatorRequest{} }
func (m *GetAuthenticatorRequest) String() string { return proto.CompactTextString(m) }
func (*GetAuthenticatorRequest) ProtoMessage()    {}
func (*GetAuthenticatorRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_20645c5a3a5aaf72, []int{4}
}
func (m *GetAuthenticatorRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetAuthenticatorRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetAuthenticatorRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetAuthenticatorRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAuthenticatorRequest.Merge(m, src)
}
func (m *GetAuthenticatorRequest) XXX_Size() int {
	return m.Size()
}
func (m *GetAuthenticatorRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAuthenticatorRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetAuthenticatorRequest proto.InternalMessageInfo

func (m *GetAuthenticatorRequest) GetAccount() string {
	if m != nil {
		return m.Account
	}
	return ""
}

func (m *GetAuthenticatorRequest) GetAuthenticatorId() uint64 {
	if m != nil {
		return m.AuthenticatorId
	}
	return 0
}

// MsgGetAuthenticatorResponse defines the Msg/GetAuthenticator response type.
type GetAuthenticatorResponse struct {
	AccountAuthenticator *AccountAuthenticator `protobuf:"bytes,1,opt,name=account_authenticator,json=accountAuthenticator,proto3" json:"account_authenticator,omitempty"`
}

func (m *GetAuthenticatorResponse) Reset()         { *m = GetAuthenticatorResponse{} }
func (m *GetAuthenticatorResponse) String() string { return proto.CompactTextString(m) }
func (*GetAuthenticatorResponse) ProtoMessage()    {}
func (*GetAuthenticatorResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_20645c5a3a5aaf72, []int{5}
}
func (m *GetAuthenticatorResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GetAuthenticatorResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GetAuthenticatorResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GetAuthenticatorResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetAuthenticatorResponse.Merge(m, src)
}
func (m *GetAuthenticatorResponse) XXX_Size() int {
	return m.Size()
}
func (m *GetAuthenticatorResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetAuthenticatorResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetAuthenticatorResponse proto.InternalMessageInfo

func (m *GetAuthenticatorResponse) GetAccountAuthenticator() *AccountAuthenticator {
	if m != nil {
		return m.AccountAuthenticator
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryParamsRequest)(nil), "ixo.smartaccount.v1beta1.QueryParamsRequest")
	proto.RegisterType((*QueryParamsResponse)(nil), "ixo.smartaccount.v1beta1.QueryParamsResponse")
	proto.RegisterType((*GetAuthenticatorsRequest)(nil), "ixo.smartaccount.v1beta1.GetAuthenticatorsRequest")
	proto.RegisterType((*GetAuthenticatorsResponse)(nil), "ixo.smartaccount.v1beta1.GetAuthenticatorsResponse")
	proto.RegisterType((*GetAuthenticatorRequest)(nil), "ixo.smartaccount.v1beta1.GetAuthenticatorRequest")
	proto.RegisterType((*GetAuthenticatorResponse)(nil), "ixo.smartaccount.v1beta1.GetAuthenticatorResponse")
}

func init() {
	proto.RegisterFile("ixo/smartaccount/v1beta1/query.proto", fileDescriptor_20645c5a3a5aaf72)
}

var fileDescriptor_20645c5a3a5aaf72 = []byte{
	// 527 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xcd, 0x96, 0x10, 0xc4, 0xf6, 0x40, 0x59, 0x52, 0x30, 0x11, 0x32, 0x96, 0x05, 0x52, 0xa8,
	0x1a, 0xaf, 0xe2, 0xf6, 0x88, 0x2a, 0x5a, 0x21, 0x21, 0x6e, 0x60, 0x89, 0x4b, 0x0f, 0x54, 0x6b,
	0x67, 0x71, 0x2c, 0x62, 0x8f, 0xeb, 0x5d, 0x57, 0xa9, 0xaa, 0x0a, 0xa9, 0x07, 0xb8, 0x22, 0xf1,
	0x23, 0xfc, 0x01, 0xd7, 0x1e, 0x2b, 0x71, 0xe1, 0x84, 0x50, 0xc2, 0x87, 0xa0, 0xd8, 0x9b, 0x80,
	0xe3, 0x58, 0x21, 0x37, 0xcf, 0xf8, 0xcd, 0xbc, 0xf7, 0x66, 0x46, 0x8b, 0x1f, 0x05, 0x43, 0xa0,
	0x22, 0x64, 0x89, 0x64, 0x9e, 0x07, 0x69, 0x24, 0xe9, 0x49, 0xd7, 0xe5, 0x92, 0x75, 0xe9, 0x71,
	0xca, 0x93, 0x53, 0x2b, 0x4e, 0x40, 0x02, 0xd1, 0x82, 0x21, 0x58, 0xff, 0xa2, 0x2c, 0x85, 0x6a,
	0x35, 0x7d, 0xf0, 0x21, 0x03, 0xd1, 0xc9, 0x57, 0x8e, 0x6f, 0x3d, 0xf0, 0x01, 0xfc, 0x01, 0xa7,
	0x2c, 0x0e, 0x28, 0x8b, 0x22, 0x90, 0x4c, 0x06, 0x10, 0x09, 0xf5, 0x77, 0xcb, 0x03, 0x11, 0x82,
	0xa0, 0x2e, 0x13, 0x3c, 0xa7, 0x99, 0x91, 0xc6, 0xcc, 0x0f, 0xa2, 0x0c, 0xac, 0xb0, 0x8f, 0x2b,
	0xf5, 0xc5, 0x2c, 0x61, 0xa1, 0x58, 0x0a, 0x0b, 0xa1, 0xc7, 0x07, 0x0a, 0x66, 0x36, 0x31, 0x79,
	0x3d, 0xe1, 0x7b, 0x95, 0xd5, 0x3a, 0xfc, 0x38, 0xe5, 0x42, 0x9a, 0x6f, 0xf0, 0x9d, 0x42, 0x56,
	0xc4, 0x10, 0x09, 0x4e, 0xf6, 0x70, 0x23, 0xe7, 0xd0, 0x90, 0x81, 0xda, 0xeb, 0xb6, 0x61, 0x55,
	0x4d, 0xc1, 0xca, 0x2b, 0x0f, 0xea, 0x97, 0x3f, 0x1f, 0xd6, 0x1c, 0x55, 0x65, 0xee, 0x62, 0xed,
	0x05, 0x97, 0xfb, 0xa9, 0xec, 0xf3, 0x48, 0x06, 0x1e, 0x93, 0x90, 0x4c, 0x29, 0x89, 0x86, 0x6f,
	0xa8, 0x1e, 0x59, 0xf3, 0x9b, 0xce, 0x34, 0x34, 0x2f, 0x10, 0xbe, 0xbf, 0xa0, 0x4c, 0x69, 0xe2,
	0xf8, 0xae, 0x02, 0x1e, 0xb1, 0x02, 0x42, 0x43, 0xc6, 0xb5, 0xf6, 0xba, 0x6d, 0x55, 0x6b, 0xdc,
	0xcf, 0xe3, 0x42, 0x63, 0x67, 0x93, 0x2d, 0xc8, 0x0a, 0xf3, 0x2d, 0xbe, 0x37, 0xaf, 0x61, 0xa9,
	0x72, 0xf2, 0x04, 0x6f, 0x14, 0x34, 0x1d, 0x05, 0x3d, 0x6d, 0xcd, 0x40, 0xed, 0xba, 0x73, 0xab,
	0x90, 0x7f, 0xd9, 0x33, 0x3f, 0x94, 0x47, 0x33, 0xb3, 0xe8, 0xe1, 0xcd, 0x85, 0x16, 0xd5, 0x16,
	0x56, 0x75, 0xd8, 0x5c, 0xe4, 0xd0, 0xfe, 0x58, 0xc7, 0xd7, 0xb3, 0x9d, 0x93, 0x4f, 0x08, 0x37,
	0xf2, 0xf5, 0x91, 0xed, 0xea, 0xd6, 0xe5, 0xab, 0x69, 0x75, 0xfe, 0x13, 0x9d, 0xdb, 0x32, 0x8d,
	0x8b, 0xef, 0xbf, 0xbf, 0xac, 0xb5, 0x88, 0x46, 0x4b, 0xa7, 0x9a, 0xdf, 0x0b, 0xf9, 0x86, 0xf0,
	0xc6, 0xfc, 0x54, 0x48, 0xb7, 0x9a, 0xa5, 0x62, 0x43, 0x2d, 0x7b, 0x95, 0x12, 0xa5, 0xee, 0x79,
	0xa6, 0x6e, 0x8f, 0x3c, 0x2d, 0xab, 0x2b, 0x2c, 0x81, 0x9e, 0xa9, 0xf4, 0x39, 0x3d, 0x9b, 0x5f,
	0xf6, 0x39, 0xf9, 0x8a, 0xf0, 0xed, 0xd2, 0xed, 0x92, 0x15, 0xf4, 0xcc, 0x86, 0xbb, 0xb3, 0x52,
	0x8d, 0x32, 0x61, 0x67, 0x26, 0xb6, 0xc9, 0xd6, 0x12, 0x13, 0xe2, 0xaf, 0x8b, 0x83, 0xc3, 0xcb,
	0x91, 0x8e, 0xae, 0x46, 0x3a, 0xfa, 0x35, 0xd2, 0xd1, 0xe7, 0xb1, 0x5e, 0xbb, 0x1a, 0xeb, 0xb5,
	0x1f, 0x63, 0xbd, 0x76, 0xf8, 0xcc, 0x0f, 0x64, 0x3f, 0x75, 0x2d, 0x0f, 0xc2, 0x49, 0xbf, 0x77,
	0x90, 0x46, 0xbd, 0xec, 0x65, 0x9a, 0x44, 0x1d, 0x77, 0x00, 0xde, 0x7b, 0xaf, 0xcf, 0x82, 0x88,
	0x9e, 0xec, 0xd2, 0x61, 0xce, 0xd6, 0x99, 0xd2, 0xc9, 0xd3, 0x98, 0x0b, 0xb7, 0x91, 0x3d, 0x3a,
	0x3b, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0x0d, 0xc5, 0x16, 0x5d, 0x64, 0x05, 0x00, 0x00,
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
	// Parameters queries the parameters of the module.
	Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error)
	GetAuthenticator(ctx context.Context, in *GetAuthenticatorRequest, opts ...grpc.CallOption) (*GetAuthenticatorResponse, error)
	GetAuthenticators(ctx context.Context, in *GetAuthenticatorsRequest, opts ...grpc.CallOption) (*GetAuthenticatorsResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Params(ctx context.Context, in *QueryParamsRequest, opts ...grpc.CallOption) (*QueryParamsResponse, error) {
	out := new(QueryParamsResponse)
	err := c.cc.Invoke(ctx, "/ixo.smartaccount.v1beta1.Query/Params", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetAuthenticator(ctx context.Context, in *GetAuthenticatorRequest, opts ...grpc.CallOption) (*GetAuthenticatorResponse, error) {
	out := new(GetAuthenticatorResponse)
	err := c.cc.Invoke(ctx, "/ixo.smartaccount.v1beta1.Query/GetAuthenticator", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) GetAuthenticators(ctx context.Context, in *GetAuthenticatorsRequest, opts ...grpc.CallOption) (*GetAuthenticatorsResponse, error) {
	out := new(GetAuthenticatorsResponse)
	err := c.cc.Invoke(ctx, "/ixo.smartaccount.v1beta1.Query/GetAuthenticators", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Parameters queries the parameters of the module.
	Params(context.Context, *QueryParamsRequest) (*QueryParamsResponse, error)
	GetAuthenticator(context.Context, *GetAuthenticatorRequest) (*GetAuthenticatorResponse, error)
	GetAuthenticators(context.Context, *GetAuthenticatorsRequest) (*GetAuthenticatorsResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Params(ctx context.Context, req *QueryParamsRequest) (*QueryParamsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Params not implemented")
}
func (*UnimplementedQueryServer) GetAuthenticator(ctx context.Context, req *GetAuthenticatorRequest) (*GetAuthenticatorResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuthenticator not implemented")
}
func (*UnimplementedQueryServer) GetAuthenticators(ctx context.Context, req *GetAuthenticatorsRequest) (*GetAuthenticatorsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAuthenticators not implemented")
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
		FullMethod: "/ixo.smartaccount.v1beta1.Query/Params",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Params(ctx, req.(*QueryParamsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetAuthenticator_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAuthenticatorRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetAuthenticator(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ixo.smartaccount.v1beta1.Query/GetAuthenticator",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetAuthenticator(ctx, req.(*GetAuthenticatorRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_GetAuthenticators_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetAuthenticatorsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).GetAuthenticators(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ixo.smartaccount.v1beta1.Query/GetAuthenticators",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).GetAuthenticators(ctx, req.(*GetAuthenticatorsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ixo.smartaccount.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Params",
			Handler:    _Query_Params_Handler,
		},
		{
			MethodName: "GetAuthenticator",
			Handler:    _Query_GetAuthenticator_Handler,
		},
		{
			MethodName: "GetAuthenticators",
			Handler:    _Query_GetAuthenticators_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ixo/smartaccount/v1beta1/query.proto",
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

func (m *GetAuthenticatorsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetAuthenticatorsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetAuthenticatorsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Account) > 0 {
		i -= len(m.Account)
		copy(dAtA[i:], m.Account)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Account)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GetAuthenticatorsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetAuthenticatorsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetAuthenticatorsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.AccountAuthenticators) > 0 {
		for iNdEx := len(m.AccountAuthenticators) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AccountAuthenticators[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *GetAuthenticatorRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetAuthenticatorRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetAuthenticatorRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.AuthenticatorId != 0 {
		i = encodeVarintQuery(dAtA, i, uint64(m.AuthenticatorId))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Account) > 0 {
		i -= len(m.Account)
		copy(dAtA[i:], m.Account)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Account)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GetAuthenticatorResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GetAuthenticatorResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GetAuthenticatorResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.AccountAuthenticator != nil {
		{
			size, err := m.AccountAuthenticator.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
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

func (m *GetAuthenticatorsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Account)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *GetAuthenticatorsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.AccountAuthenticators) > 0 {
		for _, e := range m.AccountAuthenticators {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	return n
}

func (m *GetAuthenticatorRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Account)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.AuthenticatorId != 0 {
		n += 1 + sovQuery(uint64(m.AuthenticatorId))
	}
	return n
}

func (m *GetAuthenticatorResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AccountAuthenticator != nil {
		l = m.AccountAuthenticator.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
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
func (m *GetAuthenticatorsRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: GetAuthenticatorsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetAuthenticatorsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Account", wireType)
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
			m.Account = string(dAtA[iNdEx:postIndex])
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
func (m *GetAuthenticatorsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: GetAuthenticatorsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetAuthenticatorsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccountAuthenticators", wireType)
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
			m.AccountAuthenticators = append(m.AccountAuthenticators, &AccountAuthenticator{})
			if err := m.AccountAuthenticators[len(m.AccountAuthenticators)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *GetAuthenticatorRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: GetAuthenticatorRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetAuthenticatorRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Account", wireType)
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
			m.Account = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuthenticatorId", wireType)
			}
			m.AuthenticatorId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AuthenticatorId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
func (m *GetAuthenticatorResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: GetAuthenticatorResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GetAuthenticatorResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AccountAuthenticator", wireType)
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
			if m.AccountAuthenticator == nil {
				m.AccountAuthenticator = &AccountAuthenticator{}
			}
			if err := m.AccountAuthenticator.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
