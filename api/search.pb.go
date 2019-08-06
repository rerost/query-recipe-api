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
	Metadata             *_type.GithubMetadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
	Keyword              string                `protobuf:"bytes,2,opt,name=keyword,proto3" json:"keyword,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
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

func (m *SearchRequest) GetMetadata() *_type.GithubMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

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
	// 325 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x51, 0xbd, 0x4e, 0xc3, 0x30,
	0x10, 0x56, 0x0a, 0xea, 0x8f, 0x5b, 0x16, 0x33, 0x10, 0xb5, 0x0c, 0x55, 0xa4, 0xa2, 0x0a, 0xd1,
	0x58, 0x2a, 0x13, 0xb0, 0xb1, 0x40, 0x07, 0x96, 0x74, 0x63, 0x89, 0x9c, 0xe4, 0x68, 0xad, 0xb6,
	0xb1, 0x6b, 0x5f, 0x40, 0x91, 0x98, 0x78, 0x05, 0x1e, 0x8d, 0x57, 0xe0, 0x41, 0x50, 0xed, 0xa4,
	0x12, 0x43, 0x61, 0xf4, 0x77, 0xf7, 0xfd, 0xf9, 0x48, 0xcf, 0x00, 0xd7, 0xe9, 0x32, 0x54, 0x5a,
	0xa2, 0xa4, 0x67, 0x1a, 0xb4, 0x34, 0x18, 0x6e, 0x0b, 0xd0, 0x65, 0xac, 0x21, 0x15, 0x0a, 0x62,
	0xae, 0x44, 0xff, 0x7c, 0x21, 0xe5, 0x62, 0x0d, 0x8c, 0x2b, 0xc1, 0x78, 0x9e, 0x4b, 0xe4, 0x28,
	0x64, 0x6e, 0x1c, 0xad, 0x3f, 0xa8, 0xa6, 0xf6, 0x95, 0x14, 0x2f, 0x0c, 0x36, 0x0a, 0xcb, 0x6a,
	0x48, 0xb1, 0x54, 0xc0, 0x4c, 0x2e, 0x94, 0x02, 0xac, 0xb0, 0x53, 0x8b, 0x6d, 0x00, 0x79, 0xc6,
	0x91, 0x3b, 0x30, 0x18, 0x91, 0xe6, 0xdc, 0x86, 0xa1, 0x03, 0xd2, 0x71, 0xb1, 0x62, 0x91, 0xf9,
	0xde, 0xd0, 0x1b, 0x77, 0xa2, 0xb6, 0x03, 0x66, 0x59, 0x60, 0xc8, 0x89, 0x5b, 0x8b, 0x60, 0x5b,
	0x80, 0x41, 0xfa, 0x48, 0xda, 0xb5, 0x92, 0x5d, 0xee, 0x4e, 0xaf, 0xc2, 0x03, 0x3d, 0xc2, 0x9d,
	0x6f, 0xf8, 0x20, 0x70, 0x59, 0x24, 0x4f, 0x15, 0x27, 0xda, 0xb3, 0xa9, 0x4f, 0x5a, 0x2b, 0x28,
	0xdf, 0xa4, 0xce, 0xfc, 0x86, 0x75, 0xad, 0x9f, 0xc1, 0x8c, 0xf4, 0x6a, 0x53, 0x53, 0xac, 0x91,
	0xde, 0x90, 0xe3, 0xa5, 0x40, 0xe3, 0x7b, 0xc3, 0xa3, 0x71, 0x77, 0x3a, 0xfa, 0xdb, 0x6f, 0xee,
	0xba, 0x47, 0x96, 0x32, 0x7d, 0xaf, 0xf3, 0xcf, 0x41, 0xbf, 0x8a, 0x14, 0xe8, 0x6a, 0xdf, 0xfb,
	0xe2, 0xa0, 0xce, 0xaf, 0xc6, 0xfd, 0xd1, 0xbf, 0x7b, 0xbb, 0x90, 0x01, 0xfd, 0xf8, 0xfa, 0xfe,
	0x6c, 0xf4, 0x82, 0x16, 0x73, 0x9f, 0x77, 0xeb, 0x5d, 0xde, 0xb3, 0xe7, 0xc9, 0xc2, 0xd6, 0x0f,
	0x53, 0xb9, 0x61, 0x4e, 0x86, 0x59, 0x99, 0x89, 0x93, 0x99, 0xd8, 0x13, 0x2b, 0x71, 0xc7, 0x95,
	0x88, 0x55, 0x92, 0x34, 0xed, 0x71, 0xae, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xb4, 0xb4, 0x4c,
	0xec, 0x29, 0x02, 0x00, 0x00,
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
