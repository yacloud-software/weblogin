// Code generated by protoc-gen-go.
// source: golang.conradwood.net/apis/common/common.proto
// DO NOT EDIT!

/*
Package common is a generated protocol buffer package.

It is generated from these files:
	golang.conradwood.net/apis/common/common.proto

It has these top-level messages:
	Void
	CPULoad
	Status
*/
package common

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/protoc-gen-go/descriptor"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ProgrammingLanguage int32

const (
	ProgrammingLanguage_INVALID ProgrammingLanguage = 0
	ProgrammingLanguage_GO      ProgrammingLanguage = 1
	ProgrammingLanguage_Java    ProgrammingLanguage = 2
)

var ProgrammingLanguage_name = map[int32]string{
	0: "INVALID",
	1: "GO",
	2: "Java",
}
var ProgrammingLanguage_value = map[string]int32{
	"INVALID": 0,
	"GO":      1,
	"Java":    2,
}

func (x ProgrammingLanguage) String() string {
	return proto.EnumName(ProgrammingLanguage_name, int32(x))
}
func (ProgrammingLanguage) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Language int32

const (
	Language_UNKNOWN    Language = 0
	Language_ENGLISH_UK Language = 1
	Language_GERMAN     Language = 2
)

var Language_name = map[int32]string{
	0: "UNKNOWN",
	1: "ENGLISH_UK",
	2: "GERMAN",
}
var Language_value = map[string]int32{
	"UNKNOWN":    0,
	"ENGLISH_UK": 1,
	"GERMAN":     2,
}

func (x Language) String() string {
	return proto.EnumName(Language_name, int32(x))
}
func (Language) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type Currency int32

const (
	Currency_UNKNOWN_CURRENCY Currency = 0
	Currency_GBP              Currency = 1
	Currency_EUR              Currency = 2
	Currency_USD              Currency = 3
)

var Currency_name = map[int32]string{
	0: "UNKNOWN_CURRENCY",
	1: "GBP",
	2: "EUR",
	3: "USD",
}
var Currency_value = map[string]int32{
	"UNKNOWN_CURRENCY": 0,
	"GBP":              1,
	"EUR":              2,
	"USD":              3,
}

func (x Currency) String() string {
	return proto.EnumName(Currency_name, int32(x))
}
func (Currency) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

// Void should be used when no explicit request or response is required.
type Void struct {
}

func (m *Void) Reset()                    { *m = Void{} }
func (m *Void) String() string            { return proto.CompactTextString(m) }
func (*Void) ProtoMessage()               {}
func (*Void) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

// some services report cpuload on where they are running
// usage:
// add this to your service:
// rpc GetCPULoad(common.Void) returns (common.CPULoad);
// add this to your source:
// import (
// "golang.conradwood.net/go-easyops/linux"
// )
// add this to your grpcServer struct:
// linux.Loadavg // embed GetCPULoad() RPC
type CPULoad struct {
	Avg1     float64 `protobuf:"fixed64,1,opt,name=Avg1" json:"Avg1,omitempty"`
	Avg5     float64 `protobuf:"fixed64,2,opt,name=Avg5" json:"Avg5,omitempty"`
	Avg15    float64 `protobuf:"fixed64,3,opt,name=Avg15" json:"Avg15,omitempty"`
	CPUCount uint32  `protobuf:"varint,4,opt,name=CPUCount" json:"CPUCount,omitempty"`
	PerCPU   float64 `protobuf:"fixed64,5,opt,name=PerCPU" json:"PerCPU,omitempty"`
	User     uint64  `protobuf:"varint,6,opt,name=User" json:"User,omitempty"`
	Nice     uint64  `protobuf:"varint,7,opt,name=Nice" json:"Nice,omitempty"`
	System   uint64  `protobuf:"varint,8,opt,name=System" json:"System,omitempty"`
	Idle     uint64  `protobuf:"varint,9,opt,name=Idle" json:"Idle,omitempty"`
	IOWait   uint64  `protobuf:"varint,10,opt,name=IOWait" json:"IOWait,omitempty"`
	IRQ      uint64  `protobuf:"varint,11,opt,name=IRQ" json:"IRQ,omitempty"`
	SoftIRQ  uint64  `protobuf:"varint,12,opt,name=SoftIRQ" json:"SoftIRQ,omitempty"`
	Sum      uint64  `protobuf:"varint,13,opt,name=Sum" json:"Sum,omitempty"`
	RawSum   uint64  `protobuf:"varint,14,opt,name=RawSum" json:"RawSum,omitempty"`
	IdleTime float64 `protobuf:"fixed64,15,opt,name=IdleTime" json:"IdleTime,omitempty"`
}

func (m *CPULoad) Reset()                    { *m = CPULoad{} }
func (m *CPULoad) String() string            { return proto.CompactTextString(m) }
func (*CPULoad) ProtoMessage()               {}
func (*CPULoad) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CPULoad) GetAvg1() float64 {
	if m != nil {
		return m.Avg1
	}
	return 0
}

func (m *CPULoad) GetAvg5() float64 {
	if m != nil {
		return m.Avg5
	}
	return 0
}

func (m *CPULoad) GetAvg15() float64 {
	if m != nil {
		return m.Avg15
	}
	return 0
}

func (m *CPULoad) GetCPUCount() uint32 {
	if m != nil {
		return m.CPUCount
	}
	return 0
}

func (m *CPULoad) GetPerCPU() float64 {
	if m != nil {
		return m.PerCPU
	}
	return 0
}

func (m *CPULoad) GetUser() uint64 {
	if m != nil {
		return m.User
	}
	return 0
}

func (m *CPULoad) GetNice() uint64 {
	if m != nil {
		return m.Nice
	}
	return 0
}

func (m *CPULoad) GetSystem() uint64 {
	if m != nil {
		return m.System
	}
	return 0
}

func (m *CPULoad) GetIdle() uint64 {
	if m != nil {
		return m.Idle
	}
	return 0
}

func (m *CPULoad) GetIOWait() uint64 {
	if m != nil {
		return m.IOWait
	}
	return 0
}

func (m *CPULoad) GetIRQ() uint64 {
	if m != nil {
		return m.IRQ
	}
	return 0
}

func (m *CPULoad) GetSoftIRQ() uint64 {
	if m != nil {
		return m.SoftIRQ
	}
	return 0
}

func (m *CPULoad) GetSum() uint64 {
	if m != nil {
		return m.Sum
	}
	return 0
}

func (m *CPULoad) GetRawSum() uint64 {
	if m != nil {
		return m.RawSum
	}
	return 0
}

func (m *CPULoad) GetIdleTime() float64 {
	if m != nil {
		return m.IdleTime
	}
	return 0
}

// Status represents whether or not the specified operation
// was successful, and what error occurred if it was not.
type Status struct {
	// Success is set to true or false depending on whether or
	// not the operation was successful.
	// In the event of an error, further details can be found
	// in the `ErrorCode` and `ErrorDescription` fields.
	Success bool `protobuf:"varint,1,opt,name=Success" json:"Success,omitempty"`
	// ErrorCode is present if an error has occurred during the
	// operation. ErrorCode mappings will be listed in our
	// documentation.
	ErrorCode int32 `protobuf:"varint,2,opt,name=ErrorCode" json:"ErrorCode,omitempty"`
	// ErrorDescription is present if an error has occurred during
	// the operation. This is intended to be human-readable (machines
	// should use the ErrorCode instead).
	ErrorDescription string `protobuf:"bytes,3,opt,name=ErrorDescription" json:"ErrorDescription,omitempty"`
}

func (m *Status) Reset()                    { *m = Status{} }
func (m *Status) String() string            { return proto.CompactTextString(m) }
func (*Status) ProtoMessage()               {}
func (*Status) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Status) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *Status) GetErrorCode() int32 {
	if m != nil {
		return m.ErrorCode
	}
	return 0
}

func (m *Status) GetErrorDescription() string {
	if m != nil {
		return m.ErrorDescription
	}
	return ""
}

var E_SqlType = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         51235,
	Name:          "common.sql_type",
	Tag:           "bytes,51235,opt,name=sql_type,json=sqlType",
	Filename:      "golang.conradwood.net/apis/common/common.proto",
}

var E_SqlReference = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         51236,
	Name:          "common.sql_reference",
	Tag:           "bytes,51236,opt,name=sql_reference,json=sqlReference",
	Filename:      "golang.conradwood.net/apis/common/common.proto",
}

var E_SqlUnique = &proto.ExtensionDesc{
	ExtendedType:  (*google_protobuf.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         51237,
	Name:          "common.sql_unique",
	Tag:           "bytes,51237,opt,name=sql_unique,json=sqlUnique",
	Filename:      "golang.conradwood.net/apis/common/common.proto",
}

func init() {
	proto.RegisterType((*Void)(nil), "common.Void")
	proto.RegisterType((*CPULoad)(nil), "common.CPULoad")
	proto.RegisterType((*Status)(nil), "common.Status")
	proto.RegisterEnum("common.ProgrammingLanguage", ProgrammingLanguage_name, ProgrammingLanguage_value)
	proto.RegisterEnum("common.Language", Language_name, Language_value)
	proto.RegisterEnum("common.Currency", Currency_name, Currency_value)
	proto.RegisterExtension(E_SqlType)
	proto.RegisterExtension(E_SqlReference)
	proto.RegisterExtension(E_SqlUnique)
}

func init() { proto.RegisterFile("golang.conradwood.net/apis/common/common.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 587 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x93, 0x5f, 0x6b, 0xdb, 0x3c,
	0x14, 0xc6, 0xeb, 0x24, 0x75, 0x9c, 0xd3, 0x3f, 0xaf, 0xd0, 0x5b, 0x86, 0x28, 0x1b, 0x64, 0xdd,
	0x4d, 0xc9, 0x85, 0xc3, 0xe8, 0x7a, 0x93, 0xc1, 0xa0, 0x75, 0xb2, 0x2c, 0x6b, 0xe6, 0x78, 0x72,
	0xdd, 0xb2, 0xab, 0xe2, 0xda, 0xaa, 0x31, 0xd8, 0x56, 0x22, 0xdb, 0x2d, 0xf9, 0x16, 0xfb, 0x0e,
	0xdb, 0x3e, 0xd4, 0xbe, 0xcd, 0x90, 0xe4, 0x74, 0x85, 0x5d, 0xf4, 0x2a, 0xe7, 0xf9, 0xe5, 0x39,
	0x8f, 0x38, 0xb2, 0x0e, 0xd8, 0x09, 0xcf, 0xc2, 0x22, 0xb1, 0x23, 0x5e, 0x88, 0x30, 0x7e, 0xe0,
	0x3c, 0xb6, 0x0b, 0x56, 0x0d, 0xc3, 0x65, 0x5a, 0x0e, 0x23, 0x9e, 0xe7, 0xbc, 0x68, 0x7e, 0xec,
	0xa5, 0xe0, 0x15, 0xc7, 0xa6, 0x56, 0x87, 0xfd, 0x84, 0xf3, 0x24, 0x63, 0x43, 0x45, 0x6f, 0xeb,
	0xbb, 0x61, 0xcc, 0xca, 0x48, 0xa4, 0xcb, 0x8a, 0x0b, 0xed, 0x3c, 0x32, 0xa1, 0x73, 0xc5, 0xd3,
	0xf8, 0xe8, 0x77, 0x0b, 0xba, 0x8e, 0x17, 0xcc, 0x79, 0x18, 0x63, 0x0c, 0x9d, 0xb3, 0xfb, 0xe4,
	0x2d, 0x31, 0xfa, 0xc6, 0xb1, 0x41, 0x55, 0xdd, 0xb0, 0x53, 0xd2, 0x7a, 0x64, 0xa7, 0xf8, 0x00,
	0xb6, 0xe5, 0x7f, 0xa7, 0xa4, 0xad, 0xa0, 0x16, 0xf8, 0x10, 0x2c, 0xc7, 0x0b, 0x1c, 0x5e, 0x17,
	0x15, 0xe9, 0xf4, 0x8d, 0xe3, 0x3d, 0xfa, 0xa8, 0xf1, 0x0b, 0x30, 0x3d, 0x26, 0x1c, 0x2f, 0x20,
	0xdb, 0xaa, 0xa5, 0x51, 0x32, 0x3d, 0x28, 0x99, 0x20, 0x66, 0xdf, 0x38, 0xee, 0x50, 0x55, 0x4b,
	0xe6, 0xa6, 0x11, 0x23, 0x5d, 0xcd, 0x64, 0x2d, 0xfb, 0xfd, 0x75, 0x59, 0xb1, 0x9c, 0x58, 0x8a,
	0x36, 0x4a, 0x7a, 0x67, 0x71, 0xc6, 0x48, 0x4f, 0x7b, 0x65, 0x2d, 0xbd, 0xb3, 0xc5, 0x75, 0x98,
	0x56, 0x04, 0xb4, 0x57, 0x2b, 0x8c, 0xa0, 0x3d, 0xa3, 0x5f, 0xc9, 0x8e, 0x82, 0xb2, 0xc4, 0x04,
	0xba, 0x3e, 0xbf, 0xab, 0x24, 0xdd, 0x55, 0x74, 0x23, 0xa5, 0xd7, 0xaf, 0x73, 0xb2, 0xa7, 0xbd,
	0x7e, 0x9d, 0xcb, 0x54, 0x1a, 0x3e, 0x48, 0xb8, 0xaf, 0x53, 0xb5, 0x92, 0x53, 0xcb, 0x53, 0x2f,
	0xd3, 0x9c, 0x91, 0xff, 0xd4, 0x6c, 0x8f, 0xfa, 0x28, 0x03, 0xd3, 0xaf, 0xc2, 0xaa, 0x2e, 0xd5,
	0x49, 0x75, 0x14, 0xb1, 0xb2, 0x54, 0x97, 0x6b, 0xd1, 0x8d, 0xc4, 0x2f, 0xa1, 0x37, 0x11, 0x82,
	0x0b, 0x87, 0xc7, 0x4c, 0x5d, 0xf2, 0x36, 0xfd, 0x0b, 0xf0, 0x00, 0x90, 0x12, 0xe3, 0xe6, 0xf3,
	0xa5, 0xbc, 0x50, 0x97, 0xde, 0xa3, 0xff, 0xf0, 0xc1, 0x3b, 0xf8, 0xdf, 0x13, 0x3c, 0x11, 0x61,
	0x9e, 0xa7, 0x45, 0x32, 0x0f, 0x8b, 0xa4, 0x0e, 0x13, 0x86, 0x77, 0xa0, 0x3b, 0x73, 0xaf, 0xce,
	0xe6, 0xb3, 0x31, 0xda, 0xc2, 0x26, 0xb4, 0xa6, 0x0b, 0x64, 0x60, 0x0b, 0x3a, 0x9f, 0xc3, 0xfb,
	0x10, 0xb5, 0x06, 0x27, 0x60, 0x3d, 0xb5, 0x06, 0xee, 0x85, 0xbb, 0xb8, 0x76, 0xd1, 0x16, 0xde,
	0x07, 0x98, 0xb8, 0xd3, 0xf9, 0xcc, 0xff, 0x74, 0x13, 0x5c, 0x20, 0x03, 0x03, 0x98, 0xd3, 0x09,
	0xfd, 0x72, 0xe6, 0xa2, 0xd6, 0xe0, 0x3d, 0x58, 0x4e, 0x2d, 0x04, 0x2b, 0xa2, 0x35, 0x3e, 0x00,
	0xd4, 0x34, 0xdd, 0x38, 0x01, 0xa5, 0x13, 0xd7, 0xf9, 0x86, 0xb6, 0x70, 0x17, 0xda, 0xd3, 0x73,
	0x0f, 0x19, 0xb2, 0x98, 0x04, 0x14, 0xb5, 0x64, 0x11, 0xf8, 0x63, 0xd4, 0x1e, 0x8d, 0xc0, 0x2a,
	0x57, 0xd9, 0x4d, 0xb5, 0x5e, 0x32, 0xfc, 0xca, 0xd6, 0x0f, 0xd5, 0xde, 0x3c, 0x54, 0xfb, 0x63,
	0xca, 0xb2, 0x78, 0xa1, 0x26, 0x2a, 0xc9, 0x8f, 0xef, 0x7a, 0xd8, 0x6e, 0xb9, 0xca, 0x2e, 0xd7,
	0x4b, 0x36, 0x1a, 0xc3, 0x9e, 0xec, 0x15, 0xec, 0x8e, 0xc9, 0xd3, 0x9f, 0x0d, 0xf8, 0xd9, 0x04,
	0xec, 0x96, 0xab, 0x8c, 0x6e, 0x9a, 0x46, 0x1f, 0x00, 0x64, 0x4a, 0x5d, 0xa4, 0xab, 0xfa, 0xd9,
	0x88, 0x5f, 0x4d, 0x44, 0xaf, 0x5c, 0x65, 0x81, 0xea, 0x38, 0x7f, 0x03, 0xaf, 0x0b, 0x56, 0x3d,
	0x5d, 0xca, 0x66, 0x4d, 0xe5, 0x5e, 0xda, 0x7a, 0x05, 0x6f, 0x4d, 0x15, 0x77, 0xf2, 0x27, 0x00,
	0x00, 0xff, 0xff, 0x2d, 0x9a, 0xd8, 0x82, 0xc3, 0x03, 0x00, 0x00,
}