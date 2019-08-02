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
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
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

func init() {
	proto.RegisterType((*Snippet)(nil), "rerost.query_recipe_api.type.Snippet")
}

func init() { proto.RegisterFile("snippet.proto", fileDescriptor_8ac3012f17c95bf2) }

var fileDescriptor_8ac3012f17c95bf2 = []byte{
	// 135 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0xce, 0xcb, 0x2c,
	0x28, 0x48, 0x2d, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x29, 0x4a, 0x2d, 0xca, 0x2f,
	0x2e, 0xd1, 0x2b, 0x2c, 0x4d, 0x2d, 0xaa, 0x8c, 0x2f, 0x4a, 0x4d, 0xce, 0x2c, 0x48, 0x8d, 0x4f,
	0x2c, 0xc8, 0xd4, 0x2b, 0xa9, 0x2c, 0x48, 0x55, 0x92, 0xe4, 0x62, 0x0f, 0x86, 0x28, 0x17, 0xe2,
	0xe3, 0x62, 0xca, 0x4c, 0x91, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x62, 0xca, 0x4c, 0x71, 0x32,
	0x8d, 0x32, 0x4e, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x87, 0x98, 0xa2,
	0x0f, 0x36, 0x45, 0x17, 0x62, 0x8a, 0x6e, 0x62, 0x41, 0xa6, 0x3e, 0x08, 0x83, 0x4c, 0xb2, 0x06,
	0x11, 0xf1, 0x05, 0x49, 0x49, 0x6c, 0x60, 0x6b, 0x8d, 0x01, 0x01, 0x00, 0x00, 0xff, 0xff, 0xbc,
	0xc8, 0x30, 0xf7, 0x87, 0x00, 0x00, 0x00,
}