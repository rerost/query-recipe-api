// Code generated by protoc-gen-go. DO NOT EDIT.
// source: search.proto

package api_pb

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/empty"
	_type "github.com/rerost/query-recipe-api/api/type"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Search struct {
	SearchId             string   `protobuf:"bytes,1,opt,name=search_id,json=searchId,proto3" json:"search_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Search) Reset()         { *m = Search{} }
func (m *Search) String() string { return proto.CompactTextString(m) }
func (*Search) ProtoMessage()    {}
func (*Search) Descriptor() ([]byte, []int) {
	return fileDescriptor_453745cff914010e, []int{0}
}

func (m *Search) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Search.Unmarshal(m, b)
}
func (m *Search) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Search.Marshal(b, m, deterministic)
}
func (m *Search) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Search.Merge(m, src)
}
func (m *Search) XXX_Size() int {
	return xxx_messageInfo_Search.Size(m)
}
func (m *Search) XXX_DiscardUnknown() {
	xxx_messageInfo_Search.DiscardUnknown(m)
}

var xxx_messageInfo_Search proto.InternalMessageInfo

func (m *Search) GetSearchId() string {
	if m != nil {
		return m.SearchId
	}
	return ""
}

type SearchRequest struct {
	Keyword              string   `protobuf:"bytes,1,opt,name=keyword,proto3" json:"keyword,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SearchRequest) Reset()         { *m = SearchRequest{} }
func (m *SearchRequest) String() string { return proto.CompactTextString(m) }
func (*SearchRequest) ProtoMessage()    {}
func (*SearchRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_453745cff914010e, []int{1}
}

func (m *SearchRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchRequest.Unmarshal(m, b)
}
func (m *SearchRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchRequest.Marshal(b, m, deterministic)
}
func (m *SearchRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchRequest.Merge(m, src)
}
func (m *SearchRequest) XXX_Size() int {
	return xxx_messageInfo_SearchRequest.Size(m)
}
func (m *SearchRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchRequest.DiscardUnknown(m)
}

var xxx_messageInfo_SearchRequest proto.InternalMessageInfo

func (m *SearchRequest) GetKeyword() string {
	if m != nil {
		return m.Keyword
	}
	return ""
}

type SearchResult struct {
	Hits                 []*_type.Snippet `protobuf:"bytes,1,rep,name=hits,proto3" json:"hits,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *SearchResult) Reset()         { *m = SearchResult{} }
func (m *SearchResult) String() string { return proto.CompactTextString(m) }
func (*SearchResult) ProtoMessage()    {}
func (*SearchResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_453745cff914010e, []int{2}
}

func (m *SearchResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SearchResult.Unmarshal(m, b)
}
func (m *SearchResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SearchResult.Marshal(b, m, deterministic)
}
func (m *SearchResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SearchResult.Merge(m, src)
}
func (m *SearchResult) XXX_Size() int {
	return xxx_messageInfo_SearchResult.Size(m)
}
func (m *SearchResult) XXX_DiscardUnknown() {
	xxx_messageInfo_SearchResult.DiscardUnknown(m)
}

var xxx_messageInfo_SearchResult proto.InternalMessageInfo

func (m *SearchResult) GetHits() []*_type.Snippet {
	if m != nil {
		return m.Hits
	}
	return nil
}

func init() {
	proto.RegisterType((*Search)(nil), "rerost.query_recipe_api.Search")
	proto.RegisterType((*SearchRequest)(nil), "rerost.query_recipe_api.SearchRequest")
	proto.RegisterType((*SearchResult)(nil), "rerost.query_recipe_api.SearchResult")
}

func init() { proto.RegisterFile("search.proto", fileDescriptor_453745cff914010e) }

var fileDescriptor_453745cff914010e = []byte{
	// 289 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x55, 0x40, 0x2d, 0x31, 0x45, 0x48, 0x5e, 0x88, 0x52, 0x86, 0x2a, 0x52, 0x50, 0x19,
	0x62, 0x4b, 0x65, 0x42, 0x6c, 0x6c, 0x5d, 0x93, 0x8d, 0x25, 0x72, 0xd2, 0x23, 0xb1, 0x68, 0x63,
	0xd7, 0xbe, 0x80, 0xb2, 0xf2, 0x0a, 0x3c, 0x1a, 0xaf, 0xc0, 0x83, 0xa0, 0xda, 0x0d, 0x12, 0x43,
	0xc5, 0x78, 0x77, 0xff, 0xfd, 0xf7, 0xfd, 0x36, 0x99, 0x5a, 0x10, 0xa6, 0x6a, 0x98, 0x36, 0x0a,
	0x15, 0xbd, 0x36, 0x60, 0x94, 0x45, 0xb6, 0xeb, 0xc0, 0xf4, 0x85, 0x81, 0x4a, 0x6a, 0x28, 0x84,
	0x96, 0xd1, 0x4d, 0xad, 0x54, 0xbd, 0x01, 0x2e, 0xb4, 0xe4, 0xa2, 0x6d, 0x15, 0x0a, 0x94, 0xaa,
	0xb5, 0x7e, 0x2d, 0x9a, 0x1d, 0xa6, 0xae, 0x2a, 0xbb, 0x17, 0x0e, 0x5b, 0x8d, 0xfd, 0x61, 0x48,
	0xb1, 0xd7, 0xc0, 0x6d, 0x2b, 0xb5, 0x06, 0xf4, 0xbd, 0x38, 0x21, 0xe3, 0xdc, 0xdd, 0xa5, 0x33,
	0x12, 0x78, 0x82, 0x42, 0xae, 0xc3, 0xd1, 0x7c, 0xb4, 0x08, 0xb2, 0x73, 0xdf, 0x58, 0xad, 0xe3,
	0x3b, 0x72, 0xe9, 0x65, 0x19, 0xec, 0x3a, 0xb0, 0x48, 0x43, 0x32, 0x79, 0x85, 0xfe, 0x5d, 0x99,
	0x41, 0x3b, 0x94, 0xf1, 0x8a, 0x4c, 0x07, 0xa9, 0xed, 0x36, 0x48, 0x1f, 0xc8, 0x59, 0x23, 0xd1,
	0x86, 0xa3, 0xf9, 0xe9, 0xe2, 0x62, 0x99, 0xb0, 0x23, 0xc1, 0xd8, 0x1e, 0x8e, 0xe5, 0x1e, 0x2e,
	0x73, 0x2b, 0xcb, 0x7e, 0xb8, 0x9a, 0x83, 0x79, 0x93, 0x15, 0xd0, 0xe6, 0x97, 0xf6, 0xf6, 0xa8,
	0xcf, 0x1f, 0xce, 0x28, 0xf9, 0x57, 0xb7, 0x87, 0x8c, 0xaf, 0x3e, 0xbe, 0xbe, 0x3f, 0x4f, 0x02,
	0x3a, 0xe1, 0x3e, 0xf2, 0x13, 0x7f, 0x4e, 0x6b, 0x89, 0x4d, 0x57, 0xb2, 0x4a, 0x6d, 0xb9, 0xf7,
	0xe0, 0xce, 0x23, 0xf5, 0x1e, 0xa9, 0xfb, 0x00, 0x2d, 0x1f, 0x85, 0x96, 0x85, 0x2e, 0xcb, 0xb1,
	0x7b, 0xcf, 0xfb, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0xad, 0xbf, 0xa2, 0x20, 0xc7, 0x01, 0x00,
	0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// SearchServiceClient is the client API for SearchService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SearchServiceClient interface {
	Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResult, error)
}

type searchServiceClient struct {
	cc *grpc.ClientConn
}

func NewSearchServiceClient(cc *grpc.ClientConn) SearchServiceClient {
	return &searchServiceClient{cc}
}

func (c *searchServiceClient) Search(ctx context.Context, in *SearchRequest, opts ...grpc.CallOption) (*SearchResult, error) {
	out := new(SearchResult)
	err := c.cc.Invoke(ctx, "/rerost.query_recipe_api.SearchService/Search", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SearchServiceServer is the server API for SearchService service.
type SearchServiceServer interface {
	Search(context.Context, *SearchRequest) (*SearchResult, error)
}

// UnimplementedSearchServiceServer can be embedded to have forward compatible implementations.
type UnimplementedSearchServiceServer struct {
}

func (*UnimplementedSearchServiceServer) Search(ctx context.Context, req *SearchRequest) (*SearchResult, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Search not implemented")
}

func RegisterSearchServiceServer(s *grpc.Server, srv SearchServiceServer) {
	s.RegisterService(&_SearchService_serviceDesc, srv)
}

func _SearchService_Search_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SearchRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SearchServiceServer).Search(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rerost.query_recipe_api.SearchService/Search",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SearchServiceServer).Search(ctx, req.(*SearchRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SearchService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rerost.query_recipe_api.SearchService",
	HandlerType: (*SearchServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Search",
			Handler:    _SearchService_Search_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "search.proto",
}
