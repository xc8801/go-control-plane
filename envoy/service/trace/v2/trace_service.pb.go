// Code generated by protoc-gen-go. DO NOT EDIT.
// source: envoy/service/trace/v2/trace_service.proto

package envoy_service_trace_v2

import (
	context "context"
	fmt "fmt"
	v1 "github.com/census-instrumentation/opencensus-proto/gen-go/trace/v1"
	core "github.com/envoyproxy/go-control-plane/envoy/api/v2/core"
	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

type StreamTracesResponse struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StreamTracesResponse) Reset()         { *m = StreamTracesResponse{} }
func (m *StreamTracesResponse) String() string { return proto.CompactTextString(m) }
func (*StreamTracesResponse) ProtoMessage()    {}
func (*StreamTracesResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_6feca8f22ae39b94, []int{0}
}

func (m *StreamTracesResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamTracesResponse.Unmarshal(m, b)
}
func (m *StreamTracesResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamTracesResponse.Marshal(b, m, deterministic)
}
func (m *StreamTracesResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamTracesResponse.Merge(m, src)
}
func (m *StreamTracesResponse) XXX_Size() int {
	return xxx_messageInfo_StreamTracesResponse.Size(m)
}
func (m *StreamTracesResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamTracesResponse.DiscardUnknown(m)
}

var xxx_messageInfo_StreamTracesResponse proto.InternalMessageInfo

type StreamTracesMessage struct {
	Identifier           *StreamTracesMessage_Identifier `protobuf:"bytes,1,opt,name=identifier,proto3" json:"identifier,omitempty"`
	Spans                []*v1.Span                      `protobuf:"bytes,2,rep,name=spans,proto3" json:"spans,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *StreamTracesMessage) Reset()         { *m = StreamTracesMessage{} }
func (m *StreamTracesMessage) String() string { return proto.CompactTextString(m) }
func (*StreamTracesMessage) ProtoMessage()    {}
func (*StreamTracesMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_6feca8f22ae39b94, []int{1}
}

func (m *StreamTracesMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamTracesMessage.Unmarshal(m, b)
}
func (m *StreamTracesMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamTracesMessage.Marshal(b, m, deterministic)
}
func (m *StreamTracesMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamTracesMessage.Merge(m, src)
}
func (m *StreamTracesMessage) XXX_Size() int {
	return xxx_messageInfo_StreamTracesMessage.Size(m)
}
func (m *StreamTracesMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamTracesMessage.DiscardUnknown(m)
}

var xxx_messageInfo_StreamTracesMessage proto.InternalMessageInfo

func (m *StreamTracesMessage) GetIdentifier() *StreamTracesMessage_Identifier {
	if m != nil {
		return m.Identifier
	}
	return nil
}

func (m *StreamTracesMessage) GetSpans() []*v1.Span {
	if m != nil {
		return m.Spans
	}
	return nil
}

type StreamTracesMessage_Identifier struct {
	Node                 *core.Node `protobuf:"bytes,1,opt,name=node,proto3" json:"node,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *StreamTracesMessage_Identifier) Reset()         { *m = StreamTracesMessage_Identifier{} }
func (m *StreamTracesMessage_Identifier) String() string { return proto.CompactTextString(m) }
func (*StreamTracesMessage_Identifier) ProtoMessage()    {}
func (*StreamTracesMessage_Identifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_6feca8f22ae39b94, []int{1, 0}
}

func (m *StreamTracesMessage_Identifier) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StreamTracesMessage_Identifier.Unmarshal(m, b)
}
func (m *StreamTracesMessage_Identifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StreamTracesMessage_Identifier.Marshal(b, m, deterministic)
}
func (m *StreamTracesMessage_Identifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StreamTracesMessage_Identifier.Merge(m, src)
}
func (m *StreamTracesMessage_Identifier) XXX_Size() int {
	return xxx_messageInfo_StreamTracesMessage_Identifier.Size(m)
}
func (m *StreamTracesMessage_Identifier) XXX_DiscardUnknown() {
	xxx_messageInfo_StreamTracesMessage_Identifier.DiscardUnknown(m)
}

var xxx_messageInfo_StreamTracesMessage_Identifier proto.InternalMessageInfo

func (m *StreamTracesMessage_Identifier) GetNode() *core.Node {
	if m != nil {
		return m.Node
	}
	return nil
}

func init() {
	proto.RegisterType((*StreamTracesResponse)(nil), "envoy.service.trace.v2.StreamTracesResponse")
	proto.RegisterType((*StreamTracesMessage)(nil), "envoy.service.trace.v2.StreamTracesMessage")
	proto.RegisterType((*StreamTracesMessage_Identifier)(nil), "envoy.service.trace.v2.StreamTracesMessage.Identifier")
}

func init() {
	proto.RegisterFile("envoy/service/trace/v2/trace_service.proto", fileDescriptor_6feca8f22ae39b94)
}

var fileDescriptor_6feca8f22ae39b94 = []byte{
	// 344 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x90, 0x41, 0x4b, 0xeb, 0x40,
	0x10, 0xc7, 0xdf, 0xf6, 0xbd, 0x3e, 0x64, 0xdb, 0x83, 0x46, 0x69, 0x4b, 0x28, 0x58, 0x8a, 0x42,
	0x51, 0xd9, 0xd0, 0x88, 0x1e, 0x3d, 0xc4, 0x53, 0x0f, 0x4a, 0x49, 0xc5, 0xab, 0x6c, 0x93, 0xb1,
	0x2c, 0xb4, 0x3b, 0x4b, 0x76, 0x0d, 0xf6, 0xe0, 0x5d, 0xfc, 0xa6, 0x7e, 0x05, 0x4f, 0x92, 0xdd,
	0xa4, 0xe6, 0x50, 0x41, 0x6f, 0xcb, 0xcc, 0x7f, 0x7e, 0xfb, 0x9b, 0xa1, 0x27, 0x20, 0x73, 0x5c,
	0x07, 0x1a, 0xb2, 0x5c, 0x24, 0x10, 0x98, 0x8c, 0x27, 0x10, 0xe4, 0xa1, 0x7b, 0x3c, 0x94, 0x65,
	0xa6, 0x32, 0x34, 0xe8, 0x75, 0x6c, 0x96, 0x55, 0x45, 0x1b, 0x61, 0x79, 0xe8, 0xf7, 0x1d, 0x83,
	0x2b, 0x51, 0x4c, 0x26, 0x98, 0x41, 0x30, 0xe7, 0xba, 0x9c, 0xf2, 0x8f, 0x51, 0x81, 0x4c, 0x40,
	0xea, 0x27, 0x1d, 0xd8, 0x4a, 0xf5, 0xc9, 0xd8, 0x3d, 0xca, 0x58, 0x7f, 0x81, 0xb8, 0x58, 0x82,
	0xa5, 0x70, 0x29, 0xd1, 0x70, 0x23, 0x50, 0xea, 0xb2, 0xdb, 0xcd, 0xf9, 0x52, 0xa4, 0xdc, 0x40,
	0x50, 0x3d, 0x5c, 0x63, 0xd8, 0xa1, 0x07, 0x33, 0x93, 0x01, 0x5f, 0xdd, 0x15, 0x2c, 0x1d, 0x83,
	0x56, 0x28, 0x35, 0x0c, 0xdf, 0x09, 0xdd, 0xaf, 0x37, 0x6e, 0x40, 0x6b, 0xbe, 0x00, 0xef, 0x9e,
	0x52, 0x91, 0x82, 0x34, 0xe2, 0x51, 0x40, 0xd6, 0x23, 0x03, 0x32, 0x6a, 0x85, 0x97, 0x6c, 0xfb,
	0x62, 0x6c, 0x0b, 0x80, 0x4d, 0x36, 0xd3, 0x71, 0x8d, 0xe4, 0x5d, 0xd0, 0xa6, 0x56, 0x5c, 0xea,
	0x5e, 0x63, 0xf0, 0x77, 0xd4, 0x0a, 0x0f, 0xd9, 0xd7, 0xd6, 0xce, 0xb4, 0xa2, 0x8e, 0xd9, 0x4c,
	0x71, 0x19, 0xbb, 0xb4, 0x7f, 0x4d, 0xe9, 0xa4, 0x0e, 0xf9, 0x27, 0x31, 0x85, 0x52, 0xab, 0x5b,
	0x6a, 0x71, 0x25, 0x0a, 0x99, 0xe2, 0xae, 0xec, 0x16, 0x53, 0x88, 0x76, 0x3e, 0xa2, 0xe6, 0x1b,
	0x69, 0xec, 0x92, 0xd8, 0xc6, 0xc3, 0x17, 0xda, 0xb6, 0x8e, 0x33, 0xe7, 0xef, 0xad, 0x68, 0xbb,
	0x6e, 0xee, 0x9d, 0xfe, 0x62, 0x3f, 0xff, 0xec, 0x27, 0xe1, 0xcd, 0x99, 0xff, 0x8c, 0x48, 0x74,
	0x45, 0x8f, 0x04, 0xba, 0x29, 0x95, 0xe1, 0xf3, 0xfa, 0x1b, 0x40, 0xb4, 0x57, 0x97, 0x9c, 0x16,
	0x37, 0x99, 0x92, 0x57, 0x42, 0xe6, 0xff, 0xed, 0x7d, 0xce, 0x3f, 0x03, 0x00, 0x00, 0xff, 0xff,
	0x86, 0x86, 0x56, 0x4a, 0x8b, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// TraceServiceClient is the client API for TraceService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TraceServiceClient interface {
	StreamTraces(ctx context.Context, opts ...grpc.CallOption) (TraceService_StreamTracesClient, error)
}

type traceServiceClient struct {
	cc *grpc.ClientConn
}

func NewTraceServiceClient(cc *grpc.ClientConn) TraceServiceClient {
	return &traceServiceClient{cc}
}

func (c *traceServiceClient) StreamTraces(ctx context.Context, opts ...grpc.CallOption) (TraceService_StreamTracesClient, error) {
	stream, err := c.cc.NewStream(ctx, &_TraceService_serviceDesc.Streams[0], "/envoy.service.trace.v2.TraceService/StreamTraces", opts...)
	if err != nil {
		return nil, err
	}
	x := &traceServiceStreamTracesClient{stream}
	return x, nil
}

type TraceService_StreamTracesClient interface {
	Send(*StreamTracesMessage) error
	CloseAndRecv() (*StreamTracesResponse, error)
	grpc.ClientStream
}

type traceServiceStreamTracesClient struct {
	grpc.ClientStream
}

func (x *traceServiceStreamTracesClient) Send(m *StreamTracesMessage) error {
	return x.ClientStream.SendMsg(m)
}

func (x *traceServiceStreamTracesClient) CloseAndRecv() (*StreamTracesResponse, error) {
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	m := new(StreamTracesResponse)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

// TraceServiceServer is the server API for TraceService service.
type TraceServiceServer interface {
	StreamTraces(TraceService_StreamTracesServer) error
}

func RegisterTraceServiceServer(s *grpc.Server, srv TraceServiceServer) {
	s.RegisterService(&_TraceService_serviceDesc, srv)
}

func _TraceService_StreamTraces_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(TraceServiceServer).StreamTraces(&traceServiceStreamTracesServer{stream})
}

type TraceService_StreamTracesServer interface {
	SendAndClose(*StreamTracesResponse) error
	Recv() (*StreamTracesMessage, error)
	grpc.ServerStream
}

type traceServiceStreamTracesServer struct {
	grpc.ServerStream
}

func (x *traceServiceStreamTracesServer) SendAndClose(m *StreamTracesResponse) error {
	return x.ServerStream.SendMsg(m)
}

func (x *traceServiceStreamTracesServer) Recv() (*StreamTracesMessage, error) {
	m := new(StreamTracesMessage)
	if err := x.ServerStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

var _TraceService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "envoy.service.trace.v2.TraceService",
	HandlerType: (*TraceServiceServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "StreamTraces",
			Handler:       _TraceService_StreamTraces_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "envoy/service/trace/v2/trace_service.proto",
}
