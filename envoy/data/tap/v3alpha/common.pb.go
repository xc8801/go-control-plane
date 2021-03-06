// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/data/tap/v3alpha/common.proto

package envoy_data_tap_v3alpha

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

type Body struct {
	// Types that are valid to be assigned to BodyType:
	//	*Body_AsBytes
	//	*Body_AsString
	BodyType             isBody_BodyType `protobuf_oneof:"body_type"`
	Truncated            bool            `protobuf:"varint,3,opt,name=truncated,proto3" json:"truncated,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Body) Reset()         { *m = Body{} }
func (m *Body) String() string { return proto.CompactTextString(m) }
func (*Body) ProtoMessage()    {}
func (*Body) Descriptor() ([]byte, []int) {
	return fileDescriptor_901f3a8c29139e93, []int{0}
}

func (m *Body) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Body.Unmarshal(m, b)
}
func (m *Body) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Body.Marshal(b, m, deterministic)
}
func (m *Body) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Body.Merge(m, src)
}
func (m *Body) XXX_Size() int {
	return xxx_messageInfo_Body.Size(m)
}
func (m *Body) XXX_DiscardUnknown() {
	xxx_messageInfo_Body.DiscardUnknown(m)
}

var xxx_messageInfo_Body proto.InternalMessageInfo

type isBody_BodyType interface {
	isBody_BodyType()
}

type Body_AsBytes struct {
	AsBytes []byte `protobuf:"bytes,1,opt,name=as_bytes,json=asBytes,proto3,oneof"`
}

type Body_AsString struct {
	AsString string `protobuf:"bytes,2,opt,name=as_string,json=asString,proto3,oneof"`
}

func (*Body_AsBytes) isBody_BodyType() {}

func (*Body_AsString) isBody_BodyType() {}

func (m *Body) GetBodyType() isBody_BodyType {
	if m != nil {
		return m.BodyType
	}
	return nil
}

func (m *Body) GetAsBytes() []byte {
	if x, ok := m.GetBodyType().(*Body_AsBytes); ok {
		return x.AsBytes
	}
	return nil
}

func (m *Body) GetAsString() string {
	if x, ok := m.GetBodyType().(*Body_AsString); ok {
		return x.AsString
	}
	return ""
}

func (m *Body) GetTruncated() bool {
	if m != nil {
		return m.Truncated
	}
	return false
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Body) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Body_AsBytes)(nil),
		(*Body_AsString)(nil),
	}
}

func init() {
	proto.RegisterType((*Body)(nil), "envoy.data.tap.v3alpha.Body")
}

func init() {
	proto.RegisterFile("envoy/data/tap/v3alpha/common.proto", fileDescriptor_901f3a8c29139e93)
}

var fileDescriptor_901f3a8c29139e93 = []byte{
	// 201 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0xcf, 0xb1, 0x4a, 0xc5, 0x30,
	0x14, 0xc6, 0xf1, 0x1b, 0x15, 0x6d, 0x72, 0x9d, 0x32, 0x48, 0x41, 0x85, 0xa2, 0x0e, 0x9d, 0x92,
	0xe1, 0xe2, 0x0b, 0xc4, 0xa5, 0x63, 0xa9, 0x0f, 0x50, 0x4e, 0x9b, 0xa0, 0x05, 0x9b, 0x13, 0x92,
	0x63, 0x31, 0x6f, 0x2f, 0x0d, 0x82, 0xcb, 0x5d, 0x3f, 0xfe, 0xc3, 0xef, 0x13, 0xcf, 0xce, 0x6f,
	0x98, 0xb5, 0x05, 0x02, 0x4d, 0x10, 0xf4, 0x76, 0x82, 0xaf, 0xf0, 0x09, 0x7a, 0xc6, 0x75, 0x45,
	0xaf, 0x42, 0x44, 0x42, 0x79, 0x57, 0x22, 0xb5, 0x47, 0x8a, 0x20, 0xa8, 0xbf, 0xe8, 0x69, 0x15,
	0x57, 0x06, 0x6d, 0x96, 0xf7, 0xa2, 0x82, 0x34, 0x4e, 0x99, 0x5c, 0xaa, 0x59, 0xc3, 0xda, 0xdb,
	0xee, 0x30, 0xdc, 0x40, 0x32, 0xfb, 0x20, 0x1f, 0x05, 0x87, 0x34, 0x26, 0x8a, 0x8b, 0xff, 0xa8,
	0x2f, 0x1a, 0xd6, 0xf2, 0xee, 0x30, 0x54, 0x90, 0xde, 0xcb, 0x22, 0x1f, 0x04, 0xa7, 0xf8, 0xed,
	0x67, 0x20, 0x67, 0xeb, 0xcb, 0x86, 0xb5, 0xd5, 0xf0, 0x3f, 0x98, 0xa3, 0xe0, 0x13, 0xda, 0x3c,
	0x52, 0x0e, 0xce, 0xbc, 0x8a, 0x97, 0x05, 0x55, 0xb1, 0x84, 0x88, 0x3f, 0x59, 0x9d, 0x67, 0x99,
	0xe3, 0x5b, 0xc1, 0xf7, 0xbb, 0xbd, 0x67, 0xd3, 0x75, 0x39, 0x71, 0xfa, 0x0d, 0x00, 0x00, 0xff,
	0xff, 0x9f, 0xd9, 0xaa, 0x88, 0xeb, 0x00, 0x00, 0x00,
}
