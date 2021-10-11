// Code generated by protoc-gen-go.
// source: golang.conradwood.net/apis/csf/csf.proto
// DO NOT EDIT!

/*
Package csf is a generated protocol buffer package.

It is generated from these files:
	golang.conradwood.net/apis/csf/csf.proto

It has these top-level messages:
	PingRequest
	PingResponse
*/
package csf

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import h2gproxy "golang.conradwood.net/apis/h2gproxy"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// comment: message pingrequest
type PingRequest struct {
	// comment: payload
	Payload string `protobuf:"bytes,2,opt,name=Payload" json:"Payload,omitempty"`
	// comment: sequencenumber
	SequenceNumber      uint32 `protobuf:"varint,1,opt,name=SequenceNumber" json:"SequenceNumber,omitempty"`
	HasChristmasFeature bool   `protobuf:"varint,3,opt,name=HasChristmasFeature" json:"HasChristmasFeature,omitempty"`
}

func (m *PingRequest) Reset()                    { *m = PingRequest{} }
func (m *PingRequest) String() string            { return proto.CompactTextString(m) }
func (*PingRequest) ProtoMessage()               {}
func (*PingRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PingRequest) GetPayload() string {
	if m != nil {
		return m.Payload
	}
	return ""
}

func (m *PingRequest) GetSequenceNumber() uint32 {
	if m != nil {
		return m.SequenceNumber
	}
	return 0
}

func (m *PingRequest) GetHasChristmasFeature() bool {
	if m != nil {
		return m.HasChristmasFeature
	}
	return false
}

// comment: message pingresponse
type PingResponse struct {
	// comment: field pingresponse.response
	Response *PingRequest `protobuf:"bytes,1,opt,name=Response" json:"Response,omitempty"`
}

func (m *PingResponse) Reset()                    { *m = PingResponse{} }
func (m *PingResponse) String() string            { return proto.CompactTextString(m) }
func (*PingResponse) ProtoMessage()               {}
func (*PingResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *PingResponse) GetResponse() *PingRequest {
	if m != nil {
		return m.Response
	}
	return nil
}

func init() {
	proto.RegisterType((*PingRequest)(nil), "csf.PingRequest")
	proto.RegisterType((*PingResponse)(nil), "csf.PingResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for CSFService service

type CSFServiceClient interface {
	// comment: rpc ping
	Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error)
	// this rpc is called by the html service
	ServeHTML(ctx context.Context, in *h2gproxy.ServeRequest, opts ...grpc.CallOption) (*h2gproxy.ServeResponse, error)
}

type cSFServiceClient struct {
	cc *grpc.ClientConn
}

func NewCSFServiceClient(cc *grpc.ClientConn) CSFServiceClient {
	return &cSFServiceClient{cc}
}

func (c *cSFServiceClient) Ping(ctx context.Context, in *PingRequest, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := grpc.Invoke(ctx, "/csf.CSFService/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cSFServiceClient) ServeHTML(ctx context.Context, in *h2gproxy.ServeRequest, opts ...grpc.CallOption) (*h2gproxy.ServeResponse, error) {
	out := new(h2gproxy.ServeResponse)
	err := grpc.Invoke(ctx, "/csf.CSFService/ServeHTML", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for CSFService service

type CSFServiceServer interface {
	// comment: rpc ping
	Ping(context.Context, *PingRequest) (*PingResponse, error)
	// this rpc is called by the html service
	ServeHTML(context.Context, *h2gproxy.ServeRequest) (*h2gproxy.ServeResponse, error)
}

func RegisterCSFServiceServer(s *grpc.Server, srv CSFServiceServer) {
	s.RegisterService(&_CSFService_serviceDesc, srv)
}

func _CSFService_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CSFServiceServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/csf.CSFService/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CSFServiceServer).Ping(ctx, req.(*PingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _CSFService_ServeHTML_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(h2gproxy.ServeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CSFServiceServer).ServeHTML(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/csf.CSFService/ServeHTML",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CSFServiceServer).ServeHTML(ctx, req.(*h2gproxy.ServeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CSFService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "csf.CSFService",
	HandlerType: (*CSFServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _CSFService_Ping_Handler,
		},
		{
			MethodName: "ServeHTML",
			Handler:    _CSFService_ServeHTML_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "golang.conradwood.net/apis/csf/csf.proto",
}

func init() { proto.RegisterFile("golang.conradwood.net/apis/csf/csf.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 281 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x90, 0x51, 0x4b, 0xfb, 0x30,
	0x14, 0xc5, 0xe9, 0x7f, 0x7f, 0x74, 0xbb, 0x53, 0xd1, 0x08, 0x5a, 0xf6, 0x20, 0x65, 0x0f, 0x52,
	0x50, 0x32, 0xa9, 0xaf, 0x7b, 0x72, 0x30, 0xf6, 0xa0, 0x32, 0x52, 0xbf, 0x40, 0x96, 0xde, 0x75,
	0x85, 0x2d, 0xa9, 0x49, 0xea, 0xdc, 0xa3, 0xdf, 0x5c, 0xd2, 0xae, 0xa5, 0x38, 0xf1, 0x21, 0x70,
	0x73, 0x72, 0x0e, 0xe7, 0xe6, 0x07, 0x61, 0xaa, 0xd6, 0x5c, 0xa6, 0x54, 0x28, 0xa9, 0x79, 0xb2,
	0x55, 0x2a, 0xa1, 0x12, 0xed, 0x88, 0xe7, 0x99, 0x19, 0x09, 0xb3, 0x74, 0x87, 0xe6, 0x5a, 0x59,
	0x45, 0x3a, 0xc2, 0x2c, 0x07, 0xd1, 0x1f, 0xf6, 0x55, 0x94, 0xe6, 0x5a, 0x7d, 0xee, 0x9a, 0xa1,
	0x0a, 0x0e, 0xbf, 0x3c, 0xe8, 0xcf, 0x33, 0x99, 0x32, 0x7c, 0x2f, 0xd0, 0x58, 0xe2, 0xc3, 0xf1,
	0x9c, 0xef, 0xd6, 0x8a, 0x27, 0xfe, 0xbf, 0xc0, 0x0b, 0x7b, 0xac, 0xbe, 0x92, 0x5b, 0x38, 0x8b,
	0x9d, 0x49, 0x0a, 0x7c, 0x2d, 0x36, 0x0b, 0xd4, 0xbe, 0x17, 0x78, 0xe1, 0x29, 0xfb, 0xa1, 0x92,
	0x07, 0xb8, 0x9c, 0x71, 0x33, 0x59, 0xe9, 0xcc, 0xd8, 0x0d, 0x37, 0x53, 0xe4, 0xb6, 0xd0, 0xe8,
	0x77, 0x02, 0x2f, 0xec, 0xb2, 0xdf, 0x9e, 0x86, 0x63, 0x38, 0xa9, 0x56, 0x30, 0xb9, 0x92, 0x06,
	0xc9, 0x3d, 0x74, 0xeb, 0xb9, 0xec, 0xe8, 0x47, 0xe7, 0xd4, 0x7d, 0xb5, 0xb5, 0x27, 0x6b, 0x1c,
	0xd1, 0x16, 0x60, 0x12, 0x4f, 0x63, 0xd4, 0x1f, 0x99, 0x40, 0x72, 0x07, 0xff, 0x9d, 0x8d, 0x1c,
	0x24, 0x06, 0x17, 0x2d, 0x65, 0x5f, 0x34, 0x86, 0x9e, 0xcb, 0xe1, 0xec, 0xed, 0xe5, 0x99, 0x5c,
	0xd1, 0x06, 0x4d, 0x29, 0xd6, 0xb9, 0xeb, 0x03, 0xbd, 0x4a, 0x3f, 0x05, 0x70, 0x23, 0xd1, 0xb6,
	0x69, 0xef, 0xf9, 0x3b, 0xe0, 0xae, 0x6c, 0x71, 0x54, 0x32, 0x7e, 0xfc, 0x0e, 0x00, 0x00, 0xff,
	0xff, 0xb3, 0xb9, 0x48, 0x09, 0xc8, 0x01, 0x00, 0x00,
}