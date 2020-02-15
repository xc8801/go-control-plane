// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/config/filter/http/gzip/v3alpha/gzip.proto

package envoy_config_filter_http_gzip_v3alpha

import (
	fmt "fmt"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	wrappers "github.com/golang/protobuf/ptypes/wrappers"
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

type Gzip_CompressionStrategy int32

const (
	Gzip_DEFAULT  Gzip_CompressionStrategy = 0
	Gzip_FILTERED Gzip_CompressionStrategy = 1
	Gzip_HUFFMAN  Gzip_CompressionStrategy = 2
	Gzip_RLE      Gzip_CompressionStrategy = 3
)

var Gzip_CompressionStrategy_name = map[int32]string{
	0: "DEFAULT",
	1: "FILTERED",
	2: "HUFFMAN",
	3: "RLE",
}

var Gzip_CompressionStrategy_value = map[string]int32{
	"DEFAULT":  0,
	"FILTERED": 1,
	"HUFFMAN":  2,
	"RLE":      3,
}

func (x Gzip_CompressionStrategy) String() string {
	return proto.EnumName(Gzip_CompressionStrategy_name, int32(x))
}

func (Gzip_CompressionStrategy) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_10abb2e2f45ef3b6, []int{0, 0}
}

type Gzip_CompressionLevel_Enum int32

const (
	Gzip_CompressionLevel_DEFAULT Gzip_CompressionLevel_Enum = 0
	Gzip_CompressionLevel_BEST    Gzip_CompressionLevel_Enum = 1
	Gzip_CompressionLevel_SPEED   Gzip_CompressionLevel_Enum = 2
)

var Gzip_CompressionLevel_Enum_name = map[int32]string{
	0: "DEFAULT",
	1: "BEST",
	2: "SPEED",
}

var Gzip_CompressionLevel_Enum_value = map[string]int32{
	"DEFAULT": 0,
	"BEST":    1,
	"SPEED":   2,
}

func (x Gzip_CompressionLevel_Enum) String() string {
	return proto.EnumName(Gzip_CompressionLevel_Enum_name, int32(x))
}

func (Gzip_CompressionLevel_Enum) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_10abb2e2f45ef3b6, []int{0, 0, 0}
}

type Gzip struct {
	MemoryLevel                *wrappers.UInt32Value      `protobuf:"bytes,1,opt,name=memory_level,json=memoryLevel,proto3" json:"memory_level,omitempty"`
	ContentLength              *wrappers.UInt32Value      `protobuf:"bytes,2,opt,name=content_length,json=contentLength,proto3" json:"content_length,omitempty"`
	CompressionLevel           Gzip_CompressionLevel_Enum `protobuf:"varint,3,opt,name=compression_level,json=compressionLevel,proto3,enum=envoy.config.filter.http.gzip.v3alpha.Gzip_CompressionLevel_Enum" json:"compression_level,omitempty"`
	CompressionStrategy        Gzip_CompressionStrategy   `protobuf:"varint,4,opt,name=compression_strategy,json=compressionStrategy,proto3,enum=envoy.config.filter.http.gzip.v3alpha.Gzip_CompressionStrategy" json:"compression_strategy,omitempty"`
	ContentType                []string                   `protobuf:"bytes,6,rep,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
	DisableOnEtagHeader        bool                       `protobuf:"varint,7,opt,name=disable_on_etag_header,json=disableOnEtagHeader,proto3" json:"disable_on_etag_header,omitempty"`
	RemoveAcceptEncodingHeader bool                       `protobuf:"varint,8,opt,name=remove_accept_encoding_header,json=removeAcceptEncodingHeader,proto3" json:"remove_accept_encoding_header,omitempty"`
	WindowBits                 *wrappers.UInt32Value      `protobuf:"bytes,9,opt,name=window_bits,json=windowBits,proto3" json:"window_bits,omitempty"`
	XXX_NoUnkeyedLiteral       struct{}                   `json:"-"`
	XXX_unrecognized           []byte                     `json:"-"`
	XXX_sizecache              int32                      `json:"-"`
}

func (m *Gzip) Reset()         { *m = Gzip{} }
func (m *Gzip) String() string { return proto.CompactTextString(m) }
func (*Gzip) ProtoMessage()    {}
func (*Gzip) Descriptor() ([]byte, []int) {
	return fileDescriptor_10abb2e2f45ef3b6, []int{0}
}

func (m *Gzip) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Gzip.Unmarshal(m, b)
}
func (m *Gzip) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Gzip.Marshal(b, m, deterministic)
}
func (m *Gzip) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Gzip.Merge(m, src)
}
func (m *Gzip) XXX_Size() int {
	return xxx_messageInfo_Gzip.Size(m)
}
func (m *Gzip) XXX_DiscardUnknown() {
	xxx_messageInfo_Gzip.DiscardUnknown(m)
}

var xxx_messageInfo_Gzip proto.InternalMessageInfo

func (m *Gzip) GetMemoryLevel() *wrappers.UInt32Value {
	if m != nil {
		return m.MemoryLevel
	}
	return nil
}

func (m *Gzip) GetContentLength() *wrappers.UInt32Value {
	if m != nil {
		return m.ContentLength
	}
	return nil
}

func (m *Gzip) GetCompressionLevel() Gzip_CompressionLevel_Enum {
	if m != nil {
		return m.CompressionLevel
	}
	return Gzip_CompressionLevel_DEFAULT
}

func (m *Gzip) GetCompressionStrategy() Gzip_CompressionStrategy {
	if m != nil {
		return m.CompressionStrategy
	}
	return Gzip_DEFAULT
}

func (m *Gzip) GetContentType() []string {
	if m != nil {
		return m.ContentType
	}
	return nil
}

func (m *Gzip) GetDisableOnEtagHeader() bool {
	if m != nil {
		return m.DisableOnEtagHeader
	}
	return false
}

func (m *Gzip) GetRemoveAcceptEncodingHeader() bool {
	if m != nil {
		return m.RemoveAcceptEncodingHeader
	}
	return false
}

func (m *Gzip) GetWindowBits() *wrappers.UInt32Value {
	if m != nil {
		return m.WindowBits
	}
	return nil
}

type Gzip_CompressionLevel struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Gzip_CompressionLevel) Reset()         { *m = Gzip_CompressionLevel{} }
func (m *Gzip_CompressionLevel) String() string { return proto.CompactTextString(m) }
func (*Gzip_CompressionLevel) ProtoMessage()    {}
func (*Gzip_CompressionLevel) Descriptor() ([]byte, []int) {
	return fileDescriptor_10abb2e2f45ef3b6, []int{0, 0}
}

func (m *Gzip_CompressionLevel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Gzip_CompressionLevel.Unmarshal(m, b)
}
func (m *Gzip_CompressionLevel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Gzip_CompressionLevel.Marshal(b, m, deterministic)
}
func (m *Gzip_CompressionLevel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Gzip_CompressionLevel.Merge(m, src)
}
func (m *Gzip_CompressionLevel) XXX_Size() int {
	return xxx_messageInfo_Gzip_CompressionLevel.Size(m)
}
func (m *Gzip_CompressionLevel) XXX_DiscardUnknown() {
	xxx_messageInfo_Gzip_CompressionLevel.DiscardUnknown(m)
}

var xxx_messageInfo_Gzip_CompressionLevel proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("envoy.config.filter.http.gzip.v3alpha.Gzip_CompressionStrategy", Gzip_CompressionStrategy_name, Gzip_CompressionStrategy_value)
	proto.RegisterEnum("envoy.config.filter.http.gzip.v3alpha.Gzip_CompressionLevel_Enum", Gzip_CompressionLevel_Enum_name, Gzip_CompressionLevel_Enum_value)
	proto.RegisterType((*Gzip)(nil), "envoy.config.filter.http.gzip.v3alpha.Gzip")
	proto.RegisterType((*Gzip_CompressionLevel)(nil), "envoy.config.filter.http.gzip.v3alpha.Gzip.CompressionLevel")
}

func init() {
	proto.RegisterFile("envoy/config/filter/http/gzip/v3alpha/gzip.proto", fileDescriptor_10abb2e2f45ef3b6)
}

var fileDescriptor_10abb2e2f45ef3b6 = []byte{
	// 541 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xc1, 0x6f, 0xd3, 0x30,
	0x18, 0xc5, 0x49, 0x9a, 0xb5, 0x8d, 0x3b, 0x46, 0xf0, 0x10, 0x44, 0x15, 0x4c, 0xd5, 0x24, 0xa4,
	0x68, 0x48, 0x0e, 0x6a, 0xaf, 0x48, 0xa8, 0x61, 0x29, 0x1b, 0x0a, 0x30, 0x65, 0x2d, 0xd7, 0xc8,
	0x4d, 0xdd, 0xd4, 0x52, 0x6a, 0x5b, 0x89, 0xdb, 0x2e, 0x13, 0x27, 0xee, 0x5c, 0xf8, 0x73, 0x7b,
	0x42, 0x71, 0x52, 0x18, 0x83, 0x43, 0xc5, 0xad, 0xd5, 0x7b, 0xbf, 0xf7, 0xf9, 0xfb, 0x5e, 0xc0,
	0x6b, 0xc2, 0xd6, 0xbc, 0x70, 0x63, 0xce, 0xe6, 0x34, 0x71, 0xe7, 0x34, 0x95, 0x24, 0x73, 0x17,
	0x52, 0x0a, 0x37, 0xb9, 0xa5, 0xc2, 0x5d, 0x0f, 0x70, 0x2a, 0x16, 0x58, 0xfd, 0x41, 0x22, 0xe3,
	0x92, 0xc3, 0x97, 0x8a, 0x40, 0x15, 0x81, 0x2a, 0x02, 0x95, 0x04, 0x52, 0xa6, 0x9a, 0xe8, 0x9e,
	0x24, 0x9c, 0x27, 0x29, 0x71, 0x15, 0x34, 0x5d, 0xcd, 0xdd, 0x4d, 0x86, 0x85, 0x20, 0x59, 0x5e,
	0xc5, 0x74, 0x9f, 0xad, 0x71, 0x4a, 0x67, 0x58, 0x12, 0x77, 0xf7, 0xa3, 0x12, 0x4e, 0xbf, 0x37,
	0x81, 0xf1, 0xfe, 0x96, 0x0a, 0xf8, 0x01, 0x1c, 0x2e, 0xc9, 0x92, 0x67, 0x45, 0x94, 0x92, 0x35,
	0x49, 0x6d, 0xad, 0xa7, 0x39, 0x9d, 0xfe, 0x73, 0x54, 0x05, 0xa3, 0x5d, 0x30, 0x9a, 0x5c, 0x32,
	0x39, 0xe8, 0x7f, 0xc1, 0xe9, 0x8a, 0x78, 0xe6, 0xd6, 0x6b, 0x9e, 0x19, 0xb6, 0xe9, 0x68, 0x61,
	0xa7, 0x82, 0x83, 0x92, 0x85, 0x01, 0x38, 0x8a, 0x39, 0x93, 0x84, 0xc9, 0x28, 0x25, 0x2c, 0x91,
	0x0b, 0x5b, 0xdf, 0x23, 0xad, 0xb5, 0xf5, 0x8c, 0x33, 0xdd, 0x39, 0x09, 0x1f, 0xd6, 0x70, 0xa0,
	0x58, 0x78, 0x03, 0x1e, 0xc7, 0x7c, 0x29, 0x32, 0x92, 0xe7, 0x94, 0xb3, 0xfa, 0x79, 0x8d, 0x9e,
	0xe6, 0x1c, 0xf5, 0x87, 0x68, 0xaf, 0xf3, 0xa0, 0x72, 0x43, 0xf4, 0xee, 0x77, 0x88, 0x7a, 0x26,
	0xf2, 0xd9, 0x6a, 0xe9, 0xb5, 0xb7, 0xde, 0xc1, 0x37, 0x4d, 0xb7, 0xb4, 0xd0, 0x8a, 0xef, 0x19,
	0xe0, 0x57, 0xf0, 0xe4, 0xee, 0xe4, 0x5c, 0x66, 0x58, 0x92, 0xa4, 0xb0, 0x0d, 0x35, 0xfc, 0xed,
	0x7f, 0x0e, 0xbf, 0xae, 0x63, 0xee, 0x8c, 0x3e, 0x8e, 0xff, 0x96, 0xe1, 0x2b, 0x70, 0xb8, 0xbb,
	0xa2, 0x2c, 0x04, 0xb1, 0x9b, 0xbd, 0x86, 0x63, 0x2a, 0xe8, 0x87, 0xa6, 0x5b, 0xfd, 0xb0, 0x53,
	0xab, 0xe3, 0x42, 0x10, 0x38, 0x00, 0x4f, 0x67, 0x34, 0xc7, 0xd3, 0x94, 0x44, 0x9c, 0x45, 0x44,
	0xe2, 0x24, 0x5a, 0x10, 0x3c, 0x23, 0x99, 0xdd, 0xea, 0x69, 0x4e, 0x3b, 0x3c, 0xae, 0xd5, 0xcf,
	0xcc, 0x97, 0x38, 0xb9, 0x50, 0x12, 0x1c, 0x82, 0x17, 0x19, 0x59, 0xf2, 0x35, 0x89, 0x70, 0x1c,
	0x13, 0x21, 0x23, 0xc2, 0x62, 0x3e, 0xa3, 0xec, 0x17, 0xdb, 0x56, 0x6c, 0xb7, 0x32, 0x0d, 0x95,
	0xc7, 0xaf, 0x2d, 0x75, 0xc4, 0x05, 0xe8, 0x6c, 0x28, 0x9b, 0xf1, 0x4d, 0x34, 0xa5, 0x32, 0xb7,
	0xcd, 0xfd, 0xbf, 0x9a, 0x47, 0x8e, 0x19, 0x82, 0x8a, 0xf5, 0xa8, 0xcc, 0xbb, 0x6f, 0x80, 0x75,
	0xbf, 0xa1, 0x53, 0x07, 0x18, 0x65, 0x49, 0xb0, 0x03, 0x5a, 0xe7, 0xfe, 0x68, 0x38, 0x09, 0xc6,
	0xd6, 0x03, 0xd8, 0x06, 0x86, 0xe7, 0x5f, 0x8f, 0x2d, 0x0d, 0x9a, 0xe0, 0xe0, 0xfa, 0xca, 0xf7,
	0xcf, 0x2d, 0xfd, 0x74, 0x04, 0x8e, 0xff, 0x71, 0xe2, 0x3f, 0xc1, 0x43, 0xd0, 0x1e, 0x5d, 0x06,
	0x63, 0x3f, 0xf4, 0xcf, 0x2d, 0xad, 0x94, 0x2e, 0x26, 0xa3, 0xd1, 0xc7, 0xe1, 0x27, 0x4b, 0x87,
	0x2d, 0xd0, 0x08, 0x03, 0xdf, 0x6a, 0x78, 0x1e, 0x18, 0x50, 0x5e, 0x15, 0x2b, 0x32, 0x7e, 0x53,
	0xec, 0xd7, 0xb1, 0x67, 0x96, 0x25, 0x5f, 0x95, 0xdb, 0x5e, 0x69, 0xd3, 0xa6, 0x5a, 0x7b, 0xf0,
	0x33, 0x00, 0x00, 0xff, 0xff, 0xe1, 0x93, 0xda, 0x12, 0xee, 0x03, 0x00, 0x00,
}