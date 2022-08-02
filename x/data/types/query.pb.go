// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: data/query.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

// QueryDidDocumentsRequest is request type for Query/DidDocuments RPC method.
type QueryDidDocumentsRequest struct {
	// status enables to query for validators matching a given status.
	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
	// pagination defines an optional pagination for the request.
	Pagination *query.PageRequest `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryDidDocumentsRequest) Reset()         { *m = QueryDidDocumentsRequest{} }
func (m *QueryDidDocumentsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryDidDocumentsRequest) ProtoMessage()    {}
func (*QueryDidDocumentsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4887324672c65207, []int{0}
}
func (m *QueryDidDocumentsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDidDocumentsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDidDocumentsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDidDocumentsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDidDocumentsRequest.Merge(m, src)
}
func (m *QueryDidDocumentsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryDidDocumentsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDidDocumentsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDidDocumentsRequest proto.InternalMessageInfo

func (m *QueryDidDocumentsRequest) GetStatus() string {
	if m != nil {
		return m.Status
	}
	return ""
}

func (m *QueryDidDocumentsRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// QueryDidDocumentsResponse is response type for the Query/DidDocuments RPC method
type QueryDidDocumentsResponse struct {
	// validators contains all the queried validators.
	DidDocuments []DidDocument `protobuf:"bytes,1,rep,name=didDocuments,proto3" json:"didDocuments"`
	// pagination defines the pagination in the response.
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryDidDocumentsResponse) Reset()         { *m = QueryDidDocumentsResponse{} }
func (m *QueryDidDocumentsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryDidDocumentsResponse) ProtoMessage()    {}
func (*QueryDidDocumentsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4887324672c65207, []int{1}
}
func (m *QueryDidDocumentsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDidDocumentsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDidDocumentsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDidDocumentsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDidDocumentsResponse.Merge(m, src)
}
func (m *QueryDidDocumentsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryDidDocumentsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDidDocumentsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDidDocumentsResponse proto.InternalMessageInfo

func (m *QueryDidDocumentsResponse) GetDidDocuments() []DidDocument {
	if m != nil {
		return m.DidDocuments
	}
	return nil
}

func (m *QueryDidDocumentsResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

// QueryDidDocumentsRequest is request type for Query/DidDocuments RPC method.
type QueryDidDocumentRequest struct {
	// status enables to query for validators matching a given status.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *QueryDidDocumentRequest) Reset()         { *m = QueryDidDocumentRequest{} }
func (m *QueryDidDocumentRequest) String() string { return proto.CompactTextString(m) }
func (*QueryDidDocumentRequest) ProtoMessage()    {}
func (*QueryDidDocumentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4887324672c65207, []int{2}
}
func (m *QueryDidDocumentRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDidDocumentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDidDocumentRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDidDocumentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDidDocumentRequest.Merge(m, src)
}
func (m *QueryDidDocumentRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryDidDocumentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDidDocumentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDidDocumentRequest proto.InternalMessageInfo

func (m *QueryDidDocumentRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

// QueryDidDocumentsResponse is response type for the Query/DidDocuments RPC method
type QueryDidDocumentResponse struct {
	// validators contains all the queried validators.
	DidDocument DidDocument `protobuf:"bytes,1,opt,name=didDocument,proto3" json:"didDocument"`
}

func (m *QueryDidDocumentResponse) Reset()         { *m = QueryDidDocumentResponse{} }
func (m *QueryDidDocumentResponse) String() string { return proto.CompactTextString(m) }
func (*QueryDidDocumentResponse) ProtoMessage()    {}
func (*QueryDidDocumentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4887324672c65207, []int{3}
}
func (m *QueryDidDocumentResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryDidDocumentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryDidDocumentResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryDidDocumentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryDidDocumentResponse.Merge(m, src)
}
func (m *QueryDidDocumentResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryDidDocumentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryDidDocumentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryDidDocumentResponse proto.InternalMessageInfo

func (m *QueryDidDocumentResponse) GetDidDocument() DidDocument {
	if m != nil {
		return m.DidDocument
	}
	return DidDocument{}
}

func init() {
	proto.RegisterType((*QueryDidDocumentsRequest)(nil), "ixofoundation.ixo.data.QueryDidDocumentsRequest")
	proto.RegisterType((*QueryDidDocumentsResponse)(nil), "ixofoundation.ixo.data.QueryDidDocumentsResponse")
	proto.RegisterType((*QueryDidDocumentRequest)(nil), "ixofoundation.ixo.data.QueryDidDocumentRequest")
	proto.RegisterType((*QueryDidDocumentResponse)(nil), "ixofoundation.ixo.data.QueryDidDocumentResponse")
}

func init() { proto.RegisterFile("data/query.proto", fileDescriptor_4887324672c65207) }

var fileDescriptor_4887324672c65207 = []byte{
	// 460 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x93, 0xc1, 0x6e, 0xd3, 0x30,
	0x18, 0xc7, 0xeb, 0x00, 0x93, 0x70, 0xa7, 0x09, 0x19, 0x18, 0xa1, 0xa0, 0x50, 0x15, 0x09, 0x0a,
	0x12, 0xf6, 0x5a, 0x78, 0x82, 0x69, 0x82, 0x03, 0x02, 0x41, 0x8f, 0xdc, 0x9c, 0xd8, 0x78, 0x16,
	0xa9, 0xbf, 0x6c, 0x76, 0x50, 0x07, 0xe2, 0xc2, 0x13, 0x4c, 0xf0, 0x20, 0x1c, 0x78, 0x89, 0x1d,
	0x27, 0x71, 0xe1, 0x84, 0x50, 0xcb, 0x83, 0xa0, 0xd8, 0x19, 0xa4, 0x90, 0x69, 0xdb, 0x21, 0x92,
	0x3f, 0xc7, 0xff, 0xff, 0xf7, 0xfb, 0xfe, 0x89, 0xf1, 0x25, 0xc1, 0x1d, 0x67, 0x3b, 0xa5, 0xdc,
	0xdd, 0xa3, 0xc5, 0x2e, 0x38, 0x20, 0xeb, 0x7a, 0x06, 0xaf, 0xa1, 0x34, 0x82, 0x3b, 0x0d, 0x86,
	0xea, 0x19, 0xd0, 0xea, 0x4c, 0xef, 0xa6, 0x02, 0x50, 0xb9, 0x64, 0xbc, 0xd0, 0x8c, 0x1b, 0x03,
	0xce, 0xbf, 0xb7, 0x41, 0xd5, 0xbb, 0x9f, 0x81, 0x9d, 0x82, 0x65, 0x29, 0xb7, 0x32, 0xd8, 0xb1,
	0xb7, 0xa3, 0x54, 0x3a, 0x3e, 0x62, 0x05, 0x57, 0xda, 0x04, 0xb3, 0x70, 0x76, 0xcd, 0xf7, 0x14,
	0x5a, 0xd4, 0xf5, 0x15, 0x05, 0x0a, 0xfc, 0x92, 0x55, 0xab, 0xb0, 0x3b, 0x78, 0x87, 0xe3, 0x97,
	0x95, 0xcf, 0x96, 0x16, 0x5b, 0x90, 0x95, 0x53, 0x69, 0x9c, 0x9d, 0xc8, 0x9d, 0x52, 0x5a, 0x47,
	0xd6, 0xf1, 0x8a, 0x75, 0xdc, 0x95, 0x36, 0x46, 0x7d, 0x34, 0xbc, 0x38, 0xa9, 0x2b, 0xf2, 0x18,
	0xe3, 0xbf, 0xdd, 0xe2, 0xa8, 0x8f, 0x86, 0xdd, 0xf1, 0x1d, 0x1a, 0xd0, 0x68, 0x85, 0x46, 0xc3,
	0xa4, 0x35, 0x1a, 0x7d, 0xc1, 0x95, 0xac, 0x3d, 0x27, 0x0d, 0xe5, 0xe0, 0x2b, 0xc2, 0xd7, 0x5b,
	0x9a, 0xdb, 0x02, 0x8c, 0x95, 0xe4, 0x19, 0x5e, 0x15, 0x8d, 0xfd, 0x18, 0xf5, 0xcf, 0x0d, 0xbb,
	0xe3, 0xdb, 0xb4, 0x3d, 0x38, 0xda, 0xf0, 0xd8, 0x3c, 0x7f, 0xf0, 0xe3, 0x56, 0x67, 0xb2, 0x24,
	0x27, 0x4f, 0x5a, 0xa0, 0xef, 0x9e, 0x08, 0x1d, 0x58, 0x96, 0xa8, 0xef, 0xe1, 0x6b, 0xff, 0x42,
	0x1f, 0x05, 0xb6, 0x86, 0x23, 0x2d, 0xea, 0xb0, 0x22, 0x2d, 0x06, 0xea, 0xff, 0x70, 0xff, 0x8c,
	0xf7, 0x14, 0x77, 0x1b, 0x7c, 0x5e, 0x74, 0xa6, 0xe9, 0x9a, 0xea, 0xf1, 0x97, 0x08, 0x5f, 0xf0,
	0x9d, 0xc8, 0x3e, 0xc2, 0xab, 0xcd, 0x38, 0xc9, 0xc6, 0x71, 0x96, 0xc7, 0x7d, 0xf6, 0xde, 0xe8,
	0x0c, 0x8a, 0x30, 0xcc, 0xe0, 0xc6, 0xc7, 0x6f, 0xbf, 0x3e, 0x47, 0x57, 0xc9, 0x65, 0xc6, 0xf3,
	0x5c, 0x9b, 0x54, 0x3b, 0x5b, 0xfd, 0x79, 0xd5, 0x63, 0xc9, 0x27, 0x84, 0xbb, 0x0d, 0x15, 0x61,
	0xa7, 0xf5, 0x3f, 0x02, 0xda, 0x38, 0xbd, 0xa0, 0xe6, 0xe9, 0x7b, 0x9e, 0x1e, 0x89, 0x5b, 0x78,
	0xd8, 0x7b, 0x2d, 0x3e, 0x6c, 0x3e, 0x3f, 0x98, 0x27, 0xe8, 0x70, 0x9e, 0xa0, 0x9f, 0xf3, 0x04,
	0xed, 0x2f, 0x92, 0xce, 0xe1, 0x22, 0xe9, 0x7c, 0x5f, 0x24, 0x9d, 0x57, 0x8f, 0x94, 0x76, 0xdb,
	0x65, 0x4a, 0x33, 0x98, 0xb2, 0xa5, 0xbe, 0x55, 0xf5, 0x20, 0xcd, 0x21, 0x7b, 0x93, 0x6d, 0x73,
	0x6d, 0xd8, 0x8c, 0xf9, 0x1b, 0xe6, 0xf6, 0x0a, 0x69, 0xd3, 0x15, 0x7f, 0x9d, 0x1e, 0xfe, 0x0e,
	0x00, 0x00, 0xff, 0xff, 0x4c, 0x46, 0x6b, 0xc9, 0xea, 0x03, 0x00, 0x00,
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
	// DidDocuments queries all did documents that match the given status.
	DidDocuments(ctx context.Context, in *QueryDidDocumentsRequest, opts ...grpc.CallOption) (*QueryDidDocumentsResponse, error)
	// DidDocument queries a did documents with an id.
	DidDocument(ctx context.Context, in *QueryDidDocumentRequest, opts ...grpc.CallOption) (*QueryDidDocumentResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) DidDocuments(ctx context.Context, in *QueryDidDocumentsRequest, opts ...grpc.CallOption) (*QueryDidDocumentsResponse, error) {
	out := new(QueryDidDocumentsResponse)
	err := c.cc.Invoke(ctx, "/ixofoundation.ixo.data.Query/DidDocuments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) DidDocument(ctx context.Context, in *QueryDidDocumentRequest, opts ...grpc.CallOption) (*QueryDidDocumentResponse, error) {
	out := new(QueryDidDocumentResponse)
	err := c.cc.Invoke(ctx, "/ixofoundation.ixo.data.Query/DidDocument", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// DidDocuments queries all did documents that match the given status.
	DidDocuments(context.Context, *QueryDidDocumentsRequest) (*QueryDidDocumentsResponse, error)
	// DidDocument queries a did documents with an id.
	DidDocument(context.Context, *QueryDidDocumentRequest) (*QueryDidDocumentResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) DidDocuments(ctx context.Context, req *QueryDidDocumentsRequest) (*QueryDidDocumentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DidDocuments not implemented")
}
func (*UnimplementedQueryServer) DidDocument(ctx context.Context, req *QueryDidDocumentRequest) (*QueryDidDocumentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DidDocument not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_DidDocuments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDidDocumentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DidDocuments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ixofoundation.ixo.data.Query/DidDocuments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DidDocuments(ctx, req.(*QueryDidDocumentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_DidDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryDidDocumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).DidDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ixofoundation.ixo.data.Query/DidDocument",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).DidDocument(ctx, req.(*QueryDidDocumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ixofoundation.ixo.data.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "DidDocuments",
			Handler:    _Query_DidDocuments_Handler,
		},
		{
			MethodName: "DidDocument",
			Handler:    _Query_DidDocument_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "data/query.proto",
}

func (m *QueryDidDocumentsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDidDocumentsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDidDocumentsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Status) > 0 {
		i -= len(m.Status)
		copy(dAtA[i:], m.Status)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Status)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryDidDocumentsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDidDocumentsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDidDocumentsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.DidDocuments) > 0 {
		for iNdEx := len(m.DidDocuments) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DidDocuments[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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

func (m *QueryDidDocumentRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDidDocumentRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDidDocumentRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryDidDocumentResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryDidDocumentResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryDidDocumentResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.DidDocument.MarshalToSizedBuffer(dAtA[:i])
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
func (m *QueryDidDocumentsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Status)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryDidDocumentsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.DidDocuments) > 0 {
		for _, e := range m.DidDocuments {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryDidDocumentRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryDidDocumentResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.DidDocument.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryDidDocumentsRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryDidDocumentsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDidDocumentsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Status", wireType)
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
			m.Status = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
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
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryDidDocumentsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryDidDocumentsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDidDocumentsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DidDocuments", wireType)
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
			m.DidDocuments = append(m.DidDocuments, DidDocument{})
			if err := m.DidDocuments[len(m.DidDocuments)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
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
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryDidDocumentRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryDidDocumentRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDidDocumentRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
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
			m.Id = string(dAtA[iNdEx:postIndex])
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
func (m *QueryDidDocumentResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryDidDocumentResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryDidDocumentResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DidDocument", wireType)
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
			if err := m.DidDocument.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
