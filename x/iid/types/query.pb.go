// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ixo/iid/v1beta1/query.proto

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

type QueryIidDocumentsRequest struct {
	// pagination defines an optional pagination for the request.
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryIidDocumentsRequest) Reset()         { *m = QueryIidDocumentsRequest{} }
func (m *QueryIidDocumentsRequest) String() string { return proto.CompactTextString(m) }
func (*QueryIidDocumentsRequest) ProtoMessage()    {}
func (*QueryIidDocumentsRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4950ee37c7844ad4, []int{0}
}
func (m *QueryIidDocumentsRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryIidDocumentsRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryIidDocumentsRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryIidDocumentsRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryIidDocumentsRequest.Merge(m, src)
}
func (m *QueryIidDocumentsRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryIidDocumentsRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryIidDocumentsRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryIidDocumentsRequest proto.InternalMessageInfo

func (m *QueryIidDocumentsRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryIidDocumentsResponse struct {
	IidDocuments []IidDocument `protobuf:"bytes,1,rep,name=iidDocuments,proto3" json:"iidDocuments"`
	// pagination defines the pagination in the response.
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryIidDocumentsResponse) Reset()         { *m = QueryIidDocumentsResponse{} }
func (m *QueryIidDocumentsResponse) String() string { return proto.CompactTextString(m) }
func (*QueryIidDocumentsResponse) ProtoMessage()    {}
func (*QueryIidDocumentsResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4950ee37c7844ad4, []int{1}
}
func (m *QueryIidDocumentsResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryIidDocumentsResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryIidDocumentsResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryIidDocumentsResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryIidDocumentsResponse.Merge(m, src)
}
func (m *QueryIidDocumentsResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryIidDocumentsResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryIidDocumentsResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryIidDocumentsResponse proto.InternalMessageInfo

func (m *QueryIidDocumentsResponse) GetIidDocuments() []IidDocument {
	if m != nil {
		return m.IidDocuments
	}
	return nil
}

func (m *QueryIidDocumentsResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryIidDocumentRequest struct {
	// did id of iid document querying
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (m *QueryIidDocumentRequest) Reset()         { *m = QueryIidDocumentRequest{} }
func (m *QueryIidDocumentRequest) String() string { return proto.CompactTextString(m) }
func (*QueryIidDocumentRequest) ProtoMessage()    {}
func (*QueryIidDocumentRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_4950ee37c7844ad4, []int{2}
}
func (m *QueryIidDocumentRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryIidDocumentRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryIidDocumentRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryIidDocumentRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryIidDocumentRequest.Merge(m, src)
}
func (m *QueryIidDocumentRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryIidDocumentRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryIidDocumentRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryIidDocumentRequest proto.InternalMessageInfo

func (m *QueryIidDocumentRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type QueryIidDocumentResponse struct {
	IidDocument IidDocument `protobuf:"bytes,1,opt,name=iidDocument,proto3" json:"iidDocument"`
}

func (m *QueryIidDocumentResponse) Reset()         { *m = QueryIidDocumentResponse{} }
func (m *QueryIidDocumentResponse) String() string { return proto.CompactTextString(m) }
func (*QueryIidDocumentResponse) ProtoMessage()    {}
func (*QueryIidDocumentResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_4950ee37c7844ad4, []int{3}
}
func (m *QueryIidDocumentResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryIidDocumentResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryIidDocumentResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryIidDocumentResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryIidDocumentResponse.Merge(m, src)
}
func (m *QueryIidDocumentResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryIidDocumentResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryIidDocumentResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryIidDocumentResponse proto.InternalMessageInfo

func (m *QueryIidDocumentResponse) GetIidDocument() IidDocument {
	if m != nil {
		return m.IidDocument
	}
	return IidDocument{}
}

func init() {
	proto.RegisterType((*QueryIidDocumentsRequest)(nil), "ixo.iid.v1beta1.QueryIidDocumentsRequest")
	proto.RegisterType((*QueryIidDocumentsResponse)(nil), "ixo.iid.v1beta1.QueryIidDocumentsResponse")
	proto.RegisterType((*QueryIidDocumentRequest)(nil), "ixo.iid.v1beta1.QueryIidDocumentRequest")
	proto.RegisterType((*QueryIidDocumentResponse)(nil), "ixo.iid.v1beta1.QueryIidDocumentResponse")
}

func init() { proto.RegisterFile("ixo/iid/v1beta1/query.proto", fileDescriptor_4950ee37c7844ad4) }

var fileDescriptor_4950ee37c7844ad4 = []byte{
	// 445 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x53, 0x4d, 0x6b, 0x13, 0x41,
	0x18, 0xce, 0xae, 0x1f, 0xe0, 0xa4, 0x5a, 0x18, 0x2a, 0xa6, 0xb1, 0xac, 0x25, 0x07, 0x4d, 0x0a,
	0xce, 0xd0, 0x14, 0xfc, 0x01, 0xa5, 0x54, 0xbc, 0xd5, 0x1c, 0x3d, 0x39, 0xbb, 0x33, 0x6e, 0x5f,
	0x6c, 0xe6, 0xdd, 0x76, 0x66, 0x4b, 0x8a, 0x1f, 0x07, 0xc1, 0xbb, 0xe0, 0xdf, 0xf0, 0x87, 0xf4,
	0x58, 0xf0, 0xe2, 0x49, 0x24, 0xf1, 0x87, 0xc8, 0xcc, 0x6c, 0xeb, 0x6e, 0xaa, 0x24, 0x87, 0xc0,
	0x64, 0xe6, 0xf9, 0x7a, 0x9f, 0xd9, 0x21, 0x0f, 0x61, 0x82, 0x1c, 0x40, 0xf2, 0xd3, 0xed, 0x54,
	0x59, 0xb1, 0xcd, 0x8f, 0x4b, 0x75, 0x72, 0xc6, 0x8a, 0x13, 0xb4, 0x48, 0x57, 0x61, 0x82, 0x0c,
	0x40, 0xb2, 0xea, 0xb0, 0xbb, 0x91, 0x23, 0xe6, 0x47, 0x8a, 0x8b, 0x02, 0xb8, 0xd0, 0x1a, 0xad,
	0xb0, 0x80, 0xda, 0x04, 0x78, 0x77, 0x7d, 0x5e, 0xcb, 0x51, 0xc3, 0xd1, 0x5a, 0x8e, 0x39, 0xfa,
	0x25, 0x77, 0xab, 0x6a, 0x77, 0x2b, 0x43, 0x33, 0x46, 0xc3, 0x53, 0x61, 0x54, 0x30, 0xbe, 0xa2,
	0x16, 0x22, 0x07, 0xed, 0xd5, 0x03, 0xb6, 0x97, 0x92, 0xce, 0x4b, 0x87, 0x78, 0x01, 0x72, 0x0f,
	0xb3, 0x72, 0xac, 0xb4, 0x35, 0x23, 0x75, 0x5c, 0x2a, 0x63, 0xe9, 0x3e, 0x21, 0x7f, 0xf1, 0x9d,
	0x68, 0x33, 0xea, 0xb7, 0x87, 0x8f, 0x59, 0x10, 0x67, 0x4e, 0x9c, 0x85, 0xa9, 0x2a, 0x71, 0x76,
	0x20, 0x72, 0x55, 0x71, 0x47, 0x35, 0x66, 0xef, 0x5b, 0x44, 0xd6, 0xff, 0x61, 0x62, 0x0a, 0xd4,
	0x46, 0xd1, 0x7d, 0xb2, 0x02, 0xb5, 0xfd, 0x4e, 0xb4, 0x79, 0xa3, 0xdf, 0x1e, 0x6e, 0xb0, 0xb9,
	0x92, 0x58, 0x8d, 0xbc, 0x7b, 0xf3, 0xfc, 0xe7, 0xa3, 0xd6, 0xa8, 0xc1, 0xa3, 0xcf, 0x1b, 0x69,
	0x63, 0x9f, 0xf6, 0xc9, 0xc2, 0xb4, 0x21, 0x44, 0x23, 0xee, 0x80, 0x3c, 0x98, 0x4f, 0x7b, 0xd9,
	0xc8, 0x3d, 0x12, 0x83, 0xf4, 0x4d, 0xdc, 0x19, 0xc5, 0x20, 0x7b, 0xaf, 0xaf, 0xb7, 0x77, 0x35,
	0xd7, 0x1e, 0x69, 0xd7, 0xf2, 0x55, 0xf5, 0x2d, 0x33, 0x56, 0x9d, 0x36, 0xfc, 0x1c, 0x93, 0x5b,
	0xde, 0x82, 0xbe, 0x27, 0x2b, 0xf5, 0xfe, 0xe8, 0xe0, 0x9a, 0xd4, 0xff, 0x2e, 0xb2, 0xbb, 0xb5,
	0x0c, 0x34, 0xc4, 0xee, 0xdd, 0xff, 0xf4, 0xfd, 0xf7, 0xd7, 0x78, 0x95, 0xde, 0xe5, 0xee, 0xb3,
	0x93, 0x20, 0xdd, 0xcf, 0xd0, 0x8f, 0xa4, 0x5d, 0x83, 0xd3, 0xfe, 0x42, 0xc5, 0x4b, 0xef, 0xc1,
	0x12, 0xc8, 0xca, 0xba, 0xeb, 0xad, 0xd7, 0x28, 0x6d, 0x58, 0xf3, 0x77, 0x20, 0x3f, 0xec, 0x1e,
	0x9c, 0x4f, 0x93, 0xe8, 0x62, 0x9a, 0x44, 0xbf, 0xa6, 0x49, 0xf4, 0x65, 0x96, 0xb4, 0x2e, 0x66,
	0x49, 0xeb, 0xc7, 0x2c, 0x69, 0xbd, 0x7a, 0x96, 0x83, 0x3d, 0x2c, 0x53, 0x96, 0xe1, 0xd8, 0xf1,
	0xde, 0x60, 0xa9, 0xa5, 0xbf, 0x48, 0xf7, 0xef, 0x69, 0x7a, 0x84, 0xd9, 0xdb, 0xec, 0x50, 0x80,
	0xe6, 0xa7, 0x3b, 0x7c, 0xe2, 0xdf, 0x91, 0x3d, 0x2b, 0x94, 0x49, 0x6f, 0xfb, 0x07, 0xb0, 0xf3,
	0x27, 0x00, 0x00, 0xff, 0xff, 0x74, 0x76, 0x23, 0x90, 0xab, 0x03, 0x00, 0x00,
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
	// IidDocuments queries all iid documents that match the given status.
	IidDocuments(ctx context.Context, in *QueryIidDocumentsRequest, opts ...grpc.CallOption) (*QueryIidDocumentsResponse, error)
	// IidDocument queries a iid documents with an id.
	IidDocument(ctx context.Context, in *QueryIidDocumentRequest, opts ...grpc.CallOption) (*QueryIidDocumentResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) IidDocuments(ctx context.Context, in *QueryIidDocumentsRequest, opts ...grpc.CallOption) (*QueryIidDocumentsResponse, error) {
	out := new(QueryIidDocumentsResponse)
	err := c.cc.Invoke(ctx, "/ixo.iid.v1beta1.Query/IidDocuments", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) IidDocument(ctx context.Context, in *QueryIidDocumentRequest, opts ...grpc.CallOption) (*QueryIidDocumentResponse, error) {
	out := new(QueryIidDocumentResponse)
	err := c.cc.Invoke(ctx, "/ixo.iid.v1beta1.Query/IidDocument", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// IidDocuments queries all iid documents that match the given status.
	IidDocuments(context.Context, *QueryIidDocumentsRequest) (*QueryIidDocumentsResponse, error)
	// IidDocument queries a iid documents with an id.
	IidDocument(context.Context, *QueryIidDocumentRequest) (*QueryIidDocumentResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) IidDocuments(ctx context.Context, req *QueryIidDocumentsRequest) (*QueryIidDocumentsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IidDocuments not implemented")
}
func (*UnimplementedQueryServer) IidDocument(ctx context.Context, req *QueryIidDocumentRequest) (*QueryIidDocumentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method IidDocument not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_IidDocuments_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryIidDocumentsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).IidDocuments(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ixo.iid.v1beta1.Query/IidDocuments",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).IidDocuments(ctx, req.(*QueryIidDocumentsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_IidDocument_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryIidDocumentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).IidDocument(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/ixo.iid.v1beta1.Query/IidDocument",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).IidDocument(ctx, req.(*QueryIidDocumentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "ixo.iid.v1beta1.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "IidDocuments",
			Handler:    _Query_IidDocuments_Handler,
		},
		{
			MethodName: "IidDocument",
			Handler:    _Query_IidDocument_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ixo/iid/v1beta1/query.proto",
}

func (m *QueryIidDocumentsRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryIidDocumentsRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryIidDocumentsRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryIidDocumentsResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryIidDocumentsResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryIidDocumentsResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
	if len(m.IidDocuments) > 0 {
		for iNdEx := len(m.IidDocuments) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.IidDocuments[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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

func (m *QueryIidDocumentRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryIidDocumentRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryIidDocumentRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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

func (m *QueryIidDocumentResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryIidDocumentResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryIidDocumentResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.IidDocument.MarshalToSizedBuffer(dAtA[:i])
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
func (m *QueryIidDocumentsRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryIidDocumentsResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.IidDocuments) > 0 {
		for _, e := range m.IidDocuments {
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

func (m *QueryIidDocumentRequest) Size() (n int) {
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

func (m *QueryIidDocumentResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.IidDocument.Size()
	n += 1 + l + sovQuery(uint64(l))
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryIidDocumentsRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryIidDocumentsRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryIidDocumentsRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
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
func (m *QueryIidDocumentsResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryIidDocumentsResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryIidDocumentsResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IidDocuments", wireType)
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
			m.IidDocuments = append(m.IidDocuments, IidDocument{})
			if err := m.IidDocuments[len(m.IidDocuments)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
func (m *QueryIidDocumentRequest) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryIidDocumentRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryIidDocumentRequest: illegal tag %d (wire type %d)", fieldNum, wire)
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
func (m *QueryIidDocumentResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: QueryIidDocumentResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryIidDocumentResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IidDocument", wireType)
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
			if err := m.IidDocument.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
