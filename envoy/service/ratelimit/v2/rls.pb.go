// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/service/ratelimit/v2/rls.proto

package envoy_service_ratelimit_v2

import (
	context "context"
	fmt "fmt"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	ratelimit "github.com/envoyproxy/go-control-plane/envoy/api/v2/ratelimit"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
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

type RateLimitResponse_Code int32

const (
	RateLimitResponse_UNKNOWN    RateLimitResponse_Code = 0
	RateLimitResponse_OK         RateLimitResponse_Code = 1
	RateLimitResponse_OVER_LIMIT RateLimitResponse_Code = 2
)

var RateLimitResponse_Code_name = map[int32]string{
	0: "UNKNOWN",
	1: "OK",
	2: "OVER_LIMIT",
}

var RateLimitResponse_Code_value = map[string]int32{
	"UNKNOWN":    0,
	"OK":         1,
	"OVER_LIMIT": 2,
}

func (x RateLimitResponse_Code) String() string {
	return proto.EnumName(RateLimitResponse_Code_name, int32(x))
}

func (RateLimitResponse_Code) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1de95711edb19ee8, []int{1, 0}
}

type RateLimitResponse_RateLimit_Unit int32

const (
	RateLimitResponse_RateLimit_UNKNOWN RateLimitResponse_RateLimit_Unit = 0
	RateLimitResponse_RateLimit_SECOND  RateLimitResponse_RateLimit_Unit = 1
	RateLimitResponse_RateLimit_MINUTE  RateLimitResponse_RateLimit_Unit = 2
	RateLimitResponse_RateLimit_HOUR    RateLimitResponse_RateLimit_Unit = 3
	RateLimitResponse_RateLimit_DAY     RateLimitResponse_RateLimit_Unit = 4
)

var RateLimitResponse_RateLimit_Unit_name = map[int32]string{
	0: "UNKNOWN",
	1: "SECOND",
	2: "MINUTE",
	3: "HOUR",
	4: "DAY",
}

var RateLimitResponse_RateLimit_Unit_value = map[string]int32{
	"UNKNOWN": 0,
	"SECOND":  1,
	"MINUTE":  2,
	"HOUR":    3,
	"DAY":     4,
}

func (x RateLimitResponse_RateLimit_Unit) String() string {
	return proto.EnumName(RateLimitResponse_RateLimit_Unit_name, int32(x))
}

func (RateLimitResponse_RateLimit_Unit) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1de95711edb19ee8, []int{1, 0, 0}
}

type RateLimitRequest struct {
	Domain               string                           `protobuf:"bytes,1,opt,name=domain,proto3" json:"domain,omitempty"`
	Descriptors          []*ratelimit.RateLimitDescriptor `protobuf:"bytes,2,rep,name=descriptors,proto3" json:"descriptors,omitempty"`
	HitsAddend           uint32                           `protobuf:"varint,3,opt,name=hits_addend,json=hitsAddend,proto3" json:"hits_addend,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *RateLimitRequest) Reset()         { *m = RateLimitRequest{} }
func (m *RateLimitRequest) String() string { return proto.CompactTextString(m) }
func (*RateLimitRequest) ProtoMessage()    {}
func (*RateLimitRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1de95711edb19ee8, []int{0}
}

func (m *RateLimitRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimitRequest.Unmarshal(m, b)
}
func (m *RateLimitRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimitRequest.Marshal(b, m, deterministic)
}
func (m *RateLimitRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitRequest.Merge(m, src)
}
func (m *RateLimitRequest) XXX_Size() int {
	return xxx_messageInfo_RateLimitRequest.Size(m)
}
func (m *RateLimitRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitRequest.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitRequest proto.InternalMessageInfo

func (m *RateLimitRequest) GetDomain() string {
	if m != nil {
		return m.Domain
	}
	return ""
}

func (m *RateLimitRequest) GetDescriptors() []*ratelimit.RateLimitDescriptor {
	if m != nil {
		return m.Descriptors
	}
	return nil
}

func (m *RateLimitRequest) GetHitsAddend() uint32 {
	if m != nil {
		return m.HitsAddend
	}
	return 0
}

type RateLimitResponse struct {
	OverallCode          RateLimitResponse_Code                `protobuf:"varint,1,opt,name=overall_code,json=overallCode,proto3,enum=envoy.service.ratelimit.v2.RateLimitResponse_Code" json:"overall_code,omitempty"`
	Statuses             []*RateLimitResponse_DescriptorStatus `protobuf:"bytes,2,rep,name=statuses,proto3" json:"statuses,omitempty"`
	Headers              []*core.HeaderValue                   `protobuf:"bytes,3,rep,name=headers,proto3" json:"headers,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *RateLimitResponse) Reset()         { *m = RateLimitResponse{} }
func (m *RateLimitResponse) String() string { return proto.CompactTextString(m) }
func (*RateLimitResponse) ProtoMessage()    {}
func (*RateLimitResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1de95711edb19ee8, []int{1}
}

func (m *RateLimitResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimitResponse.Unmarshal(m, b)
}
func (m *RateLimitResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimitResponse.Marshal(b, m, deterministic)
}
func (m *RateLimitResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitResponse.Merge(m, src)
}
func (m *RateLimitResponse) XXX_Size() int {
	return xxx_messageInfo_RateLimitResponse.Size(m)
}
func (m *RateLimitResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitResponse.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitResponse proto.InternalMessageInfo

func (m *RateLimitResponse) GetOverallCode() RateLimitResponse_Code {
	if m != nil {
		return m.OverallCode
	}
	return RateLimitResponse_UNKNOWN
}

func (m *RateLimitResponse) GetStatuses() []*RateLimitResponse_DescriptorStatus {
	if m != nil {
		return m.Statuses
	}
	return nil
}

func (m *RateLimitResponse) GetHeaders() []*core.HeaderValue {
	if m != nil {
		return m.Headers
	}
	return nil
}

type RateLimitResponse_RateLimit struct {
	RequestsPerUnit      uint32                           `protobuf:"varint,1,opt,name=requests_per_unit,json=requestsPerUnit,proto3" json:"requests_per_unit,omitempty"`
	Unit                 RateLimitResponse_RateLimit_Unit `protobuf:"varint,2,opt,name=unit,proto3,enum=envoy.service.ratelimit.v2.RateLimitResponse_RateLimit_Unit" json:"unit,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *RateLimitResponse_RateLimit) Reset()         { *m = RateLimitResponse_RateLimit{} }
func (m *RateLimitResponse_RateLimit) String() string { return proto.CompactTextString(m) }
func (*RateLimitResponse_RateLimit) ProtoMessage()    {}
func (*RateLimitResponse_RateLimit) Descriptor() ([]byte, []int) {
	return fileDescriptor_1de95711edb19ee8, []int{1, 0}
}

func (m *RateLimitResponse_RateLimit) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimitResponse_RateLimit.Unmarshal(m, b)
}
func (m *RateLimitResponse_RateLimit) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimitResponse_RateLimit.Marshal(b, m, deterministic)
}
func (m *RateLimitResponse_RateLimit) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitResponse_RateLimit.Merge(m, src)
}
func (m *RateLimitResponse_RateLimit) XXX_Size() int {
	return xxx_messageInfo_RateLimitResponse_RateLimit.Size(m)
}
func (m *RateLimitResponse_RateLimit) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitResponse_RateLimit.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitResponse_RateLimit proto.InternalMessageInfo

func (m *RateLimitResponse_RateLimit) GetRequestsPerUnit() uint32 {
	if m != nil {
		return m.RequestsPerUnit
	}
	return 0
}

func (m *RateLimitResponse_RateLimit) GetUnit() RateLimitResponse_RateLimit_Unit {
	if m != nil {
		return m.Unit
	}
	return RateLimitResponse_RateLimit_UNKNOWN
}

type RateLimitResponse_DescriptorStatus struct {
	Code                 RateLimitResponse_Code       `protobuf:"varint,1,opt,name=code,proto3,enum=envoy.service.ratelimit.v2.RateLimitResponse_Code" json:"code,omitempty"`
	CurrentLimit         *RateLimitResponse_RateLimit `protobuf:"bytes,2,opt,name=current_limit,json=currentLimit,proto3" json:"current_limit,omitempty"`
	LimitRemaining       uint32                       `protobuf:"varint,3,opt,name=limit_remaining,json=limitRemaining,proto3" json:"limit_remaining,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *RateLimitResponse_DescriptorStatus) Reset()         { *m = RateLimitResponse_DescriptorStatus{} }
func (m *RateLimitResponse_DescriptorStatus) String() string { return proto.CompactTextString(m) }
func (*RateLimitResponse_DescriptorStatus) ProtoMessage()    {}
func (*RateLimitResponse_DescriptorStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_1de95711edb19ee8, []int{1, 1}
}

func (m *RateLimitResponse_DescriptorStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RateLimitResponse_DescriptorStatus.Unmarshal(m, b)
}
func (m *RateLimitResponse_DescriptorStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RateLimitResponse_DescriptorStatus.Marshal(b, m, deterministic)
}
func (m *RateLimitResponse_DescriptorStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RateLimitResponse_DescriptorStatus.Merge(m, src)
}
func (m *RateLimitResponse_DescriptorStatus) XXX_Size() int {
	return xxx_messageInfo_RateLimitResponse_DescriptorStatus.Size(m)
}
func (m *RateLimitResponse_DescriptorStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_RateLimitResponse_DescriptorStatus.DiscardUnknown(m)
}

var xxx_messageInfo_RateLimitResponse_DescriptorStatus proto.InternalMessageInfo

func (m *RateLimitResponse_DescriptorStatus) GetCode() RateLimitResponse_Code {
	if m != nil {
		return m.Code
	}
	return RateLimitResponse_UNKNOWN
}

func (m *RateLimitResponse_DescriptorStatus) GetCurrentLimit() *RateLimitResponse_RateLimit {
	if m != nil {
		return m.CurrentLimit
	}
	return nil
}

func (m *RateLimitResponse_DescriptorStatus) GetLimitRemaining() uint32 {
	if m != nil {
		return m.LimitRemaining
	}
	return 0
}

func init() {
	proto.RegisterEnum("envoy.service.ratelimit.v2.RateLimitResponse_Code", RateLimitResponse_Code_name, RateLimitResponse_Code_value)
	proto.RegisterEnum("envoy.service.ratelimit.v2.RateLimitResponse_RateLimit_Unit", RateLimitResponse_RateLimit_Unit_name, RateLimitResponse_RateLimit_Unit_value)
	proto.RegisterType((*RateLimitRequest)(nil), "envoy.service.ratelimit.v2.RateLimitRequest")
	proto.RegisterType((*RateLimitResponse)(nil), "envoy.service.ratelimit.v2.RateLimitResponse")
	proto.RegisterType((*RateLimitResponse_RateLimit)(nil), "envoy.service.ratelimit.v2.RateLimitResponse.RateLimit")
	proto.RegisterType((*RateLimitResponse_DescriptorStatus)(nil), "envoy.service.ratelimit.v2.RateLimitResponse.DescriptorStatus")
}

func init() {
	proto.RegisterFile("envoy/service/ratelimit/v2/rls.proto", fileDescriptor_1de95711edb19ee8)
}

var fileDescriptor_1de95711edb19ee8 = []byte{
	// 583 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x94, 0xdd, 0x6e, 0xd3, 0x4c,
	0x10, 0x86, 0xeb, 0x24, 0x5f, 0xda, 0x4e, 0xfa, 0xe3, 0xee, 0xc1, 0x47, 0x64, 0x21, 0xa8, 0x22,
	0x04, 0x11, 0x05, 0x5b, 0x32, 0x07, 0x80, 0x84, 0x2a, 0xf5, 0x0f, 0xb5, 0x6a, 0x9b, 0x44, 0x9b,
	0xa6, 0x08, 0x84, 0x64, 0x6d, 0xe3, 0x11, 0x5d, 0xc9, 0xf5, 0x9a, 0xdd, 0xb5, 0x45, 0xcf, 0x39,
	0xe0, 0x1a, 0xb8, 0x23, 0x2e, 0x82, 0x03, 0xee, 0x04, 0x79, 0xed, 0x38, 0x29, 0x08, 0x44, 0xe0,
	0xcc, 0x9e, 0x99, 0xf7, 0xf1, 0xbc, 0x33, 0xeb, 0x85, 0x7b, 0x18, 0x67, 0xe2, 0xda, 0x53, 0x28,
	0x33, 0x3e, 0x46, 0x4f, 0x32, 0x8d, 0x11, 0xbf, 0xe2, 0xda, 0xcb, 0x7c, 0x4f, 0x46, 0xca, 0x4d,
	0xa4, 0xd0, 0x82, 0x38, 0xa6, 0xca, 0x2d, 0xab, 0xdc, 0xaa, 0xca, 0xcd, 0x7c, 0xe7, 0x76, 0x41,
	0x60, 0x09, 0xcf, 0x35, 0x63, 0x21, 0xd1, 0xbb, 0x60, 0x0a, 0x0b, 0xa5, 0x73, 0xff, 0x46, 0x76,
	0x8a, 0x9f, 0x22, 0x8a, 0xba, 0x5b, 0x19, 0x8b, 0x78, 0xc8, 0x34, 0x7a, 0x93, 0x87, 0x22, 0xd1,
	0xf9, 0x6c, 0x81, 0x4d, 0x99, 0xc6, 0x93, 0xbc, 0x98, 0xe2, 0xfb, 0x14, 0x95, 0x26, 0xff, 0x43,
	0x33, 0x14, 0x57, 0x8c, 0xc7, 0x6d, 0x6b, 0xd3, 0xea, 0x2e, 0xd3, 0xf2, 0x8d, 0x9c, 0x42, 0x2b,
	0x44, 0x35, 0x96, 0x3c, 0xd1, 0x42, 0xaa, 0x76, 0x6d, 0xb3, 0xde, 0x6d, 0xf9, 0x5b, 0x6e, 0xd1,
	0x3d, 0x4b, 0xb8, 0x9b, 0xf9, 0x33, 0xcd, 0x57, 0xd8, 0xfd, 0x4a, 0x43, 0x67, 0xf5, 0xe4, 0x2e,
	0xb4, 0x2e, 0xb9, 0x56, 0x01, 0x0b, 0x43, 0x8c, 0xc3, 0x76, 0x7d, 0xd3, 0xea, 0xae, 0x52, 0xc8,
	0x43, 0x3b, 0x26, 0xd2, 0xf9, 0xfa, 0x1f, 0x6c, 0xcc, 0x34, 0xa7, 0x12, 0x11, 0x2b, 0x24, 0x23,
	0x58, 0x11, 0x19, 0x4a, 0x16, 0x45, 0xc1, 0x58, 0x84, 0x68, 0x7a, 0x5c, 0xf3, 0x7d, 0xf7, 0xd7,
	0x43, 0x74, 0x7f, 0x82, 0xb8, 0x7b, 0x22, 0x44, 0xda, 0x2a, 0x39, 0xf9, 0x0b, 0x79, 0x03, 0x4b,
	0x4a, 0x33, 0x9d, 0x2a, 0x9c, 0x38, 0xdb, 0x9e, 0x0f, 0x39, 0xb5, 0x39, 0x34, 0x1c, 0x5a, 0xf1,
	0xc8, 0x33, 0x58, 0xbc, 0x44, 0x16, 0xa2, 0x54, 0xed, 0xba, 0x41, 0xdf, 0xb9, 0x39, 0xb4, 0x7c,
	0xad, 0xee, 0xa1, 0xa9, 0x38, 0x67, 0x51, 0x8a, 0x74, 0x52, 0xee, 0x7c, 0xb1, 0x60, 0xb9, 0xfa,
	0x14, 0x79, 0x08, 0x1b, 0xb2, 0xd8, 0x91, 0x0a, 0x12, 0x94, 0x41, 0x1a, 0x73, 0x6d, 0xfc, 0xaf,
	0xd2, 0xf5, 0x49, 0x62, 0x80, 0x72, 0x14, 0x73, 0x4d, 0x06, 0xd0, 0x30, 0xe9, 0x9a, 0x19, 0xcf,
	0x8b, 0xf9, 0xbc, 0x54, 0x11, 0x37, 0x67, 0x51, 0x43, 0xea, 0x6c, 0x43, 0xc3, 0x90, 0x5b, 0xb0,
	0x38, 0xea, 0x1d, 0xf7, 0xfa, 0xaf, 0x7a, 0xf6, 0x02, 0x01, 0x68, 0x0e, 0x0f, 0xf6, 0xfa, 0xbd,
	0x7d, 0xdb, 0xca, 0x9f, 0x4f, 0x8f, 0x7a, 0xa3, 0xb3, 0x03, 0xbb, 0x46, 0x96, 0xa0, 0x71, 0xd8,
	0x1f, 0x51, 0xbb, 0x4e, 0x16, 0xa1, 0xbe, 0xbf, 0xf3, 0xda, 0x6e, 0x38, 0xdf, 0x2c, 0xb0, 0x7f,
	0x1c, 0x12, 0x79, 0x09, 0x8d, 0x7f, 0xdc, 0xa2, 0xd1, 0x93, 0xb7, 0xb0, 0x3a, 0x4e, 0xa5, 0xc4,
	0x58, 0x07, 0x46, 0x60, 0x7c, 0xb7, 0xfc, 0xa7, 0x7f, 0xe9, 0x9b, 0xae, 0x94, 0xb4, 0x62, 0xf0,
	0x0f, 0x60, 0xdd, 0xa8, 0x02, 0x89, 0xf9, 0x9f, 0xc0, 0xe3, 0x77, 0xe5, 0x71, 0x5d, 0x8b, 0x0a,
	0x7d, 0x19, 0xed, 0x6c, 0x41, 0xc3, 0x9c, 0xa6, 0x1b, 0x33, 0x6a, 0x42, 0xad, 0x7f, 0x6c, 0x5b,
	0x64, 0x0d, 0xa0, 0x7f, 0x7e, 0x40, 0x83, 0x93, 0xa3, 0xd3, 0xa3, 0x33, 0xbb, 0xe6, 0x7f, 0x9c,
	0xfd, 0xf9, 0x86, 0x45, 0x87, 0x24, 0x81, 0xf5, 0xe1, 0xa5, 0x48, 0xa3, 0x70, 0xba, 0xf6, 0x47,
	0x7f, 0x68, 0xc2, 0x1c, 0x00, 0xe7, 0xf1, 0x5c, 0x96, 0x3b, 0x0b, 0xbb, 0xcf, 0xa1, 0xcb, 0x45,
	0x21, 0x4a, 0xa4, 0xf8, 0x70, 0xfd, 0x1b, 0xfd, 0xee, 0x12, 0x8d, 0xd4, 0x20, 0xbf, 0x39, 0x06,
	0xd6, 0x27, 0xcb, 0xba, 0x68, 0x9a, 0x5b, 0xe4, 0xc9, 0xf7, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe9,
	0xb0, 0x2f, 0x86, 0xe8, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// RateLimitServiceClient is the client API for RateLimitService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type RateLimitServiceClient interface {
	ShouldRateLimit(ctx context.Context, in *RateLimitRequest, opts ...grpc.CallOption) (*RateLimitResponse, error)
}

type rateLimitServiceClient struct {
	cc *grpc.ClientConn
}

func NewRateLimitServiceClient(cc *grpc.ClientConn) RateLimitServiceClient {
	return &rateLimitServiceClient{cc}
}

func (c *rateLimitServiceClient) ShouldRateLimit(ctx context.Context, in *RateLimitRequest, opts ...grpc.CallOption) (*RateLimitResponse, error) {
	out := new(RateLimitResponse)
	err := c.cc.Invoke(ctx, "/envoy.service.ratelimit.v2.RateLimitService/ShouldRateLimit", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RateLimitServiceServer is the server API for RateLimitService service.
type RateLimitServiceServer interface {
	ShouldRateLimit(context.Context, *RateLimitRequest) (*RateLimitResponse, error)
}

func RegisterRateLimitServiceServer(s *grpc.Server, srv RateLimitServiceServer) {
	s.RegisterService(&_RateLimitService_serviceDesc, srv)
}

func _RateLimitService_ShouldRateLimit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RateLimitRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RateLimitServiceServer).ShouldRateLimit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/envoy.service.ratelimit.v2.RateLimitService/ShouldRateLimit",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RateLimitServiceServer).ShouldRateLimit(ctx, req.(*RateLimitRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _RateLimitService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.ratelimit.v2.RateLimitService",
	HandlerType: (*RateLimitServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ShouldRateLimit",
			Handler:    _RateLimitService_ShouldRateLimit_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "envoy/service/ratelimit/v2/rls.proto",
}
