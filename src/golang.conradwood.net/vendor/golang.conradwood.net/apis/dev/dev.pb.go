// Code generated by protoc-gen-go.
// source: golang.conradwood.net/apis/dev/dev.proto
// DO NOT EDIT!

/*
Package dev is a generated protocol buffer package.

It is generated from these files:
	golang.conradwood.net/apis/dev/dev.proto

It has these top-level messages:
	Config
	ConfigResponse
*/
package dev

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

// configure devserver on-the-fly
type Config struct {
	// make all local users admin
	AuthAdmin bool `protobuf:"varint,1,opt,name=AuthAdmin" json:"AuthAdmin,omitempty"`
	// mock authentication service
	MockAuth bool `protobuf:"varint,2,opt,name=MockAuth" json:"MockAuth,omitempty"`
	// mock rpcinterceptor
	MockRPCInterceptor bool `protobuf:"varint,3,opt,name=MockRPCInterceptor" json:"MockRPCInterceptor,omitempty"`
	// debug (mocked) rpcinterceptor
	DebugRPCInterceptor bool `protobuf:"varint,4,opt,name=DebugRPCInterceptor" json:"DebugRPCInterceptor,omitempty"`
	// general debugging
	Debug bool `protobuf:"varint,5,opt,name=Debug" json:"Debug,omitempty"`
	// debug (mocked) auth
	DebugAuth bool `protobuf:"varint,6,opt,name=DebugAuth" json:"DebugAuth,omitempty"`
	// debug mocked registry
	DebugRegistry bool `protobuf:"varint,7,opt,name=DebugRegistry" json:"DebugRegistry,omitempty"`
}

func (m *Config) Reset()                    { *m = Config{} }
func (m *Config) String() string            { return proto.CompactTextString(m) }
func (*Config) ProtoMessage()               {}
func (*Config) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Config) GetAuthAdmin() bool {
	if m != nil {
		return m.AuthAdmin
	}
	return false
}

func (m *Config) GetMockAuth() bool {
	if m != nil {
		return m.MockAuth
	}
	return false
}

func (m *Config) GetMockRPCInterceptor() bool {
	if m != nil {
		return m.MockRPCInterceptor
	}
	return false
}

func (m *Config) GetDebugRPCInterceptor() bool {
	if m != nil {
		return m.DebugRPCInterceptor
	}
	return false
}

func (m *Config) GetDebug() bool {
	if m != nil {
		return m.Debug
	}
	return false
}

func (m *Config) GetDebugAuth() bool {
	if m != nil {
		return m.DebugAuth
	}
	return false
}

func (m *Config) GetDebugRegistry() bool {
	if m != nil {
		return m.DebugRegistry
	}
	return false
}

type ConfigResponse struct {
	Config  *Config `protobuf:"bytes,1,opt,name=Config" json:"Config,omitempty"`
	Display string  `protobuf:"bytes,2,opt,name=Display" json:"Display,omitempty"`
}

func (m *ConfigResponse) Reset()                    { *m = ConfigResponse{} }
func (m *ConfigResponse) String() string            { return proto.CompactTextString(m) }
func (*ConfigResponse) ProtoMessage()               {}
func (*ConfigResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *ConfigResponse) GetConfig() *Config {
	if m != nil {
		return m.Config
	}
	return nil
}

func (m *ConfigResponse) GetDisplay() string {
	if m != nil {
		return m.Display
	}
	return ""
}

func init() {
	proto.RegisterType((*Config)(nil), "dev.Config")
	proto.RegisterType((*ConfigResponse)(nil), "dev.ConfigResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for DevService service

type DevServiceClient interface {
	// Reconfigure devservice
	Configure(ctx context.Context, in *Config, opts ...grpc.CallOption) (*ConfigResponse, error)
	GetConfig(ctx context.Context, in *common.Void, opts ...grpc.CallOption) (*ConfigResponse, error)
}

type devServiceClient struct {
	cc *grpc.ClientConn
}

func NewDevServiceClient(cc *grpc.ClientConn) DevServiceClient {
	return &devServiceClient{cc}
}

func (c *devServiceClient) Configure(ctx context.Context, in *Config, opts ...grpc.CallOption) (*ConfigResponse, error) {
	out := new(ConfigResponse)
	err := grpc.Invoke(ctx, "/dev.DevService/Configure", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *devServiceClient) GetConfig(ctx context.Context, in *common.Void, opts ...grpc.CallOption) (*ConfigResponse, error) {
	out := new(ConfigResponse)
	err := grpc.Invoke(ctx, "/dev.DevService/GetConfig", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for DevService service

type DevServiceServer interface {
	// Reconfigure devservice
	Configure(context.Context, *Config) (*ConfigResponse, error)
	GetConfig(context.Context, *common.Void) (*ConfigResponse, error)
}

func RegisterDevServiceServer(s *grpc.Server, srv DevServiceServer) {
	s.RegisterService(&_DevService_serviceDesc, srv)
}

func _DevService_Configure_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Config)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevServiceServer).Configure(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dev.DevService/Configure",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevServiceServer).Configure(ctx, req.(*Config))
	}
	return interceptor(ctx, in, info, handler)
}

func _DevService_GetConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DevServiceServer).GetConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/dev.DevService/GetConfig",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DevServiceServer).GetConfig(ctx, req.(*common.Void))
	}
	return interceptor(ctx, in, info, handler)
}

var _DevService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "dev.DevService",
	HandlerType: (*DevServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Configure",
			Handler:    _DevService_Configure_Handler,
		},
		{
			MethodName: "GetConfig",
			Handler:    _DevService_GetConfig_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "golang.conradwood.net/apis/dev/dev.proto",
}

func init() { proto.RegisterFile("golang.conradwood.net/apis/dev/dev.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 319 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x51, 0x4d, 0x4f, 0xc2, 0x40,
	0x10, 0x0d, 0x22, 0x1f, 0x1d, 0xd4, 0xc3, 0xe2, 0xa1, 0x21, 0xc6, 0x10, 0xf4, 0xc0, 0xc5, 0xc5,
	0xe0, 0x2f, 0x40, 0x48, 0x8c, 0x07, 0xa3, 0xa9, 0x89, 0xf7, 0xd2, 0x1d, 0xeb, 0x06, 0xd8, 0x69,
	0xb6, 0x4b, 0x0d, 0x7f, 0xde, 0x98, 0xdd, 0x6d, 0x11, 0x4d, 0xe3, 0xa1, 0xe9, 0xce, 0x9b, 0xf7,
	0x66, 0x67, 0xdf, 0x83, 0x71, 0x4a, 0xeb, 0x58, 0xa5, 0x3c, 0x21, 0xa5, 0x63, 0xf1, 0x49, 0x24,
	0xb8, 0x42, 0x33, 0x89, 0x33, 0x99, 0x4f, 0x04, 0x16, 0xf6, 0xe3, 0x99, 0x26, 0x43, 0xac, 0x29,
	0xb0, 0x18, 0xf0, 0x7f, 0xe8, 0x09, 0x6d, 0x36, 0xa4, 0xca, 0x9f, 0x17, 0x8d, 0xbe, 0x1a, 0xd0,
	0x9e, 0x93, 0x7a, 0x97, 0x29, 0xbb, 0x80, 0x60, 0xb6, 0x35, 0x1f, 0x33, 0xb1, 0x91, 0x2a, 0x6c,
	0x0c, 0x1b, 0xe3, 0x6e, 0xf4, 0x03, 0xb0, 0x01, 0x74, 0x9f, 0x28, 0x59, 0x59, 0x20, 0x3c, 0x72,
	0xcd, 0x7d, 0xcd, 0x38, 0x30, 0x7b, 0x8e, 0x5e, 0xe6, 0x8f, 0xca, 0xa0, 0x4e, 0x30, 0x33, 0xa4,
	0xc3, 0xa6, 0x63, 0xd5, 0x74, 0xd8, 0x2d, 0xf4, 0x17, 0xb8, 0xdc, 0xa6, 0x7f, 0x04, 0xc7, 0x4e,
	0x50, 0xd7, 0x62, 0xe7, 0xd0, 0x72, 0x70, 0xd8, 0x72, 0x1c, 0x5f, 0xd8, 0x8d, 0xdd, 0xc1, 0x2d,
	0xd5, 0xf6, 0x1b, 0xef, 0x01, 0x76, 0x0d, 0xa7, 0x7e, 0x14, 0xa6, 0x32, 0x37, 0x7a, 0x17, 0x76,
	0x1c, 0xe3, 0x37, 0x38, 0x7a, 0x86, 0x33, 0xff, 0xfe, 0x08, 0xf3, 0x8c, 0x54, 0x8e, 0xec, 0xaa,
	0x72, 0xc4, 0x99, 0xd0, 0x9b, 0xf6, 0xb8, 0xf5, 0xb8, 0x24, 0x55, 0x66, 0x85, 0xd0, 0x59, 0xc8,
	0x3c, 0x5b, 0xc7, 0x3b, 0xe7, 0x46, 0x10, 0x55, 0xe5, 0x74, 0x05, 0xb0, 0xc0, 0xe2, 0x15, 0x75,
	0x21, 0x13, 0x64, 0x37, 0x10, 0x78, 0xc5, 0x56, 0x23, 0x3b, 0x9c, 0x34, 0xe8, 0x1f, 0x8e, 0xad,
	0xee, 0xe6, 0x10, 0x3c, 0xa0, 0x29, 0xef, 0x38, 0xe1, 0x65, 0x54, 0x6f, 0x24, 0x45, 0x2d, 0xff,
	0x7e, 0x08, 0x97, 0x0a, 0xcd, 0x61, 0xda, 0x65, 0xfe, 0x36, 0x70, 0x2b, 0x58, 0xb6, 0x5d, 0xce,
	0x77, 0xdf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x6f, 0x2a, 0xd1, 0x07, 0x48, 0x02, 0x00, 0x00,
}