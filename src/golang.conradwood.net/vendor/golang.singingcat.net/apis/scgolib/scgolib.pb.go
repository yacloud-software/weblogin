// Code generated by protoc-gen-go.
// source: golang.singingcat.net/apis/scgolib/scgolib.proto
// DO NOT EDIT!

/*
Package scgolib is a generated protocol buffer package.

It is generated from these files:
	golang.singingcat.net/apis/scgolib/scgolib.proto

It has these top-level messages:
	PingResponse
*/
package scgolib

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import common "golang.conradwood.net/apis/common"

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

// comment: message pingresponse
type PingResponse struct {
	// comment: field pingresponse.response
	Response string `protobuf:"bytes,1,opt,name=Response" json:"Response,omitempty"`
}

func (m *PingResponse) Reset()                    { *m = PingResponse{} }
func (m *PingResponse) String() string            { return proto.CompactTextString(m) }
func (*PingResponse) ProtoMessage()               {}
func (*PingResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *PingResponse) GetResponse() string {
	if m != nil {
		return m.Response
	}
	return ""
}

func init() {
	proto.RegisterType((*PingResponse)(nil), "scgolib.PingResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for SCGoLib service

type SCGoLibClient interface {
	// comment: rpc ping
	Ping(ctx context.Context, in *common.Void, opts ...grpc.CallOption) (*PingResponse, error)
}

type sCGoLibClient struct {
	cc *grpc.ClientConn
}

func NewSCGoLibClient(cc *grpc.ClientConn) SCGoLibClient {
	return &sCGoLibClient{cc}
}

func (c *sCGoLibClient) Ping(ctx context.Context, in *common.Void, opts ...grpc.CallOption) (*PingResponse, error) {
	out := new(PingResponse)
	err := grpc.Invoke(ctx, "/scgolib.SCGoLib/Ping", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for SCGoLib service

type SCGoLibServer interface {
	// comment: rpc ping
	Ping(context.Context, *common.Void) (*PingResponse, error)
}

func RegisterSCGoLibServer(s *grpc.Server, srv SCGoLibServer) {
	s.RegisterService(&_SCGoLib_serviceDesc, srv)
}

func _SCGoLib_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SCGoLibServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/scgolib.SCGoLib/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SCGoLibServer).Ping(ctx, req.(*common.Void))
	}
	return interceptor(ctx, in, info, handler)
}

var _SCGoLib_serviceDesc = grpc.ServiceDesc{
	ServiceName: "scgolib.SCGoLib",
	HandlerType: (*SCGoLibServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _SCGoLib_Ping_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "golang.singingcat.net/apis/scgolib/scgolib.proto",
}

func init() { proto.RegisterFile("golang.singingcat.net/apis/scgolib/scgolib.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 180 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0x32, 0x48, 0xcf, 0xcf, 0x49,
	0xcc, 0x4b, 0xd7, 0x2b, 0xce, 0xcc, 0x4b, 0xcf, 0xcc, 0x4b, 0x4f, 0x4e, 0x2c, 0xd1, 0xcb, 0x4b,
	0x2d, 0xd1, 0x4f, 0x2c, 0xc8, 0x2c, 0xd6, 0x2f, 0x4e, 0x4e, 0xcf, 0xcf, 0xc9, 0x4c, 0x82, 0xd1,
	0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xec, 0x50, 0xae, 0x94, 0x1e, 0x54, 0x6b, 0x72, 0x7e,
	0x5e, 0x51, 0x62, 0x4a, 0x79, 0x7e, 0x7e, 0x0a, 0x42, 0x6b, 0x72, 0x7e, 0x6e, 0x6e, 0x7e, 0x1e,
	0x94, 0x82, 0x68, 0x54, 0xd2, 0xe2, 0xe2, 0x09, 0xc8, 0xcc, 0x4b, 0x0f, 0x4a, 0x2d, 0x2e, 0xc8,
	0xcf, 0x2b, 0x4e, 0x15, 0x92, 0xe2, 0xe2, 0x80, 0xb1, 0x25, 0x18, 0x15, 0x18, 0x35, 0x38, 0x83,
	0xe0, 0x7c, 0x23, 0x33, 0x2e, 0xf6, 0x60, 0x67, 0xf7, 0x7c, 0x9f, 0xcc, 0x24, 0x21, 0x6d, 0x2e,
	0x16, 0x90, 0x36, 0x21, 0x1e, 0x3d, 0xa8, 0x69, 0x61, 0xf9, 0x99, 0x29, 0x52, 0xa2, 0x7a, 0x30,
	0x57, 0x21, 0x9b, 0xe9, 0x24, 0xcb, 0x25, 0x9d, 0x97, 0x5a, 0x82, 0xec, 0x1b, 0x90, 0x73, 0x60,
	0x6a, 0x93, 0xd8, 0xc0, 0x2e, 0x31, 0x06, 0x04, 0x00, 0x00, 0xff, 0xff, 0xa5, 0x05, 0x78, 0x76,
	0xf6, 0x00, 0x00, 0x00,
}