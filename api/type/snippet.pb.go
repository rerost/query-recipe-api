// Code generated by protoc-gen-go. DO NOT EDIT.
// source: snippet.proto

package type_pb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type Snippet struct {
	Id  string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Sql string `protobuf:"bytes,2,opt,name=sql,proto3" json:"sql,omitempty"`
	// GFM
	Document             string   `protobuf:"bytes,3,opt,name=document,proto3" json:"document,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Snippet) Reset()         { *m = Snippet{} }
func (m *Snippet) String() string { return proto.CompactTextString(m) }
func (*Snippet) ProtoMessage()    {}
func (*Snippet) Descriptor() ([]byte, []int) {
	return fileDescriptor_8ac3012f17c95bf2, []int{0}
}

func (m *Snippet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Snippet.Unmarshal(m, b)
}
func (m *Snippet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Snippet.Marshal(b, m, deterministic)
}
func (m *Snippet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Snippet.Merge(m, src)
}
func (m *Snippet) XXX_Size() int {
	return xxx_messageInfo_Snippet.Size(m)
}
func (m *Snippet) XXX_DiscardUnknown() {
	xxx_messageInfo_Snippet.DiscardUnknown(m)
}

var xxx_messageInfo_Snippet proto.InternalMessageInfo

func (m *Snippet) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Snippet) GetSql() string {
	if m != nil {
		return m.Sql
	}
	return ""
}

func (m *Snippet) GetDocument() string {
	if m != nil {
		return m.Document
	}
	return ""
}

func init() {
	proto.RegisterType((*Snippet)(nil), "rerost.query_recipe_api.type.Snippet")
}

func init() { proto.RegisterFile("snippet.proto", fileDescriptor_8ac3012f17c95bf2) }

var fileDescriptor_8ac3012f17c95bf2 = []byte{
	// 165 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0xce, 0xcb, 0x2c,
	0x28, 0x48, 0x2d, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x29, 0x4a, 0x2d, 0xca, 0x2f,
	0x2e, 0xd1, 0x2b, 0x2c, 0x4d, 0x2d, 0xaa, 0x8c, 0x2f, 0x4a, 0x4d, 0xce, 0x2c, 0x48, 0x8d, 0x4f,
	0x2c, 0xc8, 0xd4, 0x2b, 0xa9, 0x2c, 0x48, 0x55, 0x72, 0xe7, 0x62, 0x0f, 0x86, 0x28, 0x17, 0xe2,
	0xe3, 0x62, 0xca, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x62, 0xca, 0x4c, 0x11, 0x12,
	0xe0, 0x62, 0x2e, 0x2e, 0xcc, 0x91, 0x60, 0x02, 0x0b, 0x80, 0x98, 0x42, 0x52, 0x5c, 0x1c, 0x29,
	0xf9, 0xc9, 0xa5, 0xb9, 0xa9, 0x79, 0x25, 0x12, 0xcc, 0x60, 0x61, 0x38, 0xdf, 0xc9, 0x34, 0xca,
	0x38, 0x3d, 0xb3, 0x24, 0xa3, 0x34, 0x49, 0x2f, 0x39, 0x3f, 0x57, 0x1f, 0x62, 0xa7, 0x3e, 0xd8,
	0x4e, 0x5d, 0x88, 0x9d, 0xba, 0x89, 0x05, 0x99, 0xfa, 0x20, 0x0c, 0xb2, 0xd7, 0x1a, 0x44, 0xc4,
	0x17, 0x24, 0x25, 0xb1, 0x81, 0x1d, 0x69, 0x0c, 0x08, 0x00, 0x00, 0xff, 0xff, 0x11, 0xa9, 0x25,
	0x8a, 0xb5, 0x00, 0x00, 0x00,
}
