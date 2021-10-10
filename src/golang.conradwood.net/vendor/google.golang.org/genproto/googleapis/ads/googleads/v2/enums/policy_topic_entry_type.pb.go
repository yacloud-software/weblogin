// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/ads/googleads/v2/enums/policy_topic_entry_type.proto

package enums

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
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

// The possible policy topic entry types.
type PolicyTopicEntryTypeEnum_PolicyTopicEntryType int32

const (
	// No value has been specified.
	PolicyTopicEntryTypeEnum_UNSPECIFIED PolicyTopicEntryTypeEnum_PolicyTopicEntryType = 0
	// The received value is not known in this version.
	//
	// This is a response-only value.
	PolicyTopicEntryTypeEnum_UNKNOWN PolicyTopicEntryTypeEnum_PolicyTopicEntryType = 1
	// The resource will not be served.
	PolicyTopicEntryTypeEnum_PROHIBITED PolicyTopicEntryTypeEnum_PolicyTopicEntryType = 2
	// The resource will not be served under some circumstances.
	PolicyTopicEntryTypeEnum_LIMITED PolicyTopicEntryTypeEnum_PolicyTopicEntryType = 4
	// The resource cannot serve at all because of the current targeting
	// criteria.
	PolicyTopicEntryTypeEnum_FULLY_LIMITED PolicyTopicEntryTypeEnum_PolicyTopicEntryType = 8
	// May be of interest, but does not limit how the resource is served.
	PolicyTopicEntryTypeEnum_DESCRIPTIVE PolicyTopicEntryTypeEnum_PolicyTopicEntryType = 5
	// Could increase coverage beyond normal.
	PolicyTopicEntryTypeEnum_BROADENING PolicyTopicEntryTypeEnum_PolicyTopicEntryType = 6
	// Constrained for all targeted countries, but may serve in other countries
	// through area of interest.
	PolicyTopicEntryTypeEnum_AREA_OF_INTEREST_ONLY PolicyTopicEntryTypeEnum_PolicyTopicEntryType = 7
)

var PolicyTopicEntryTypeEnum_PolicyTopicEntryType_name = map[int32]string{
	0: "UNSPECIFIED",
	1: "UNKNOWN",
	2: "PROHIBITED",
	4: "LIMITED",
	8: "FULLY_LIMITED",
	5: "DESCRIPTIVE",
	6: "BROADENING",
	7: "AREA_OF_INTEREST_ONLY",
}

var PolicyTopicEntryTypeEnum_PolicyTopicEntryType_value = map[string]int32{
	"UNSPECIFIED":           0,
	"UNKNOWN":               1,
	"PROHIBITED":            2,
	"LIMITED":               4,
	"FULLY_LIMITED":         8,
	"DESCRIPTIVE":           5,
	"BROADENING":            6,
	"AREA_OF_INTEREST_ONLY": 7,
}

func (x PolicyTopicEntryTypeEnum_PolicyTopicEntryType) String() string {
	return proto.EnumName(PolicyTopicEntryTypeEnum_PolicyTopicEntryType_name, int32(x))
}

func (PolicyTopicEntryTypeEnum_PolicyTopicEntryType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_59b40b5d8dcd4000, []int{0, 0}
}

// Container for enum describing possible policy topic entry types.
type PolicyTopicEntryTypeEnum struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PolicyTopicEntryTypeEnum) Reset()         { *m = PolicyTopicEntryTypeEnum{} }
func (m *PolicyTopicEntryTypeEnum) String() string { return proto.CompactTextString(m) }
func (*PolicyTopicEntryTypeEnum) ProtoMessage()    {}
func (*PolicyTopicEntryTypeEnum) Descriptor() ([]byte, []int) {
	return fileDescriptor_59b40b5d8dcd4000, []int{0}
}

func (m *PolicyTopicEntryTypeEnum) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PolicyTopicEntryTypeEnum.Unmarshal(m, b)
}
func (m *PolicyTopicEntryTypeEnum) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PolicyTopicEntryTypeEnum.Marshal(b, m, deterministic)
}
func (m *PolicyTopicEntryTypeEnum) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PolicyTopicEntryTypeEnum.Merge(m, src)
}
func (m *PolicyTopicEntryTypeEnum) XXX_Size() int {
	return xxx_messageInfo_PolicyTopicEntryTypeEnum.Size(m)
}
func (m *PolicyTopicEntryTypeEnum) XXX_DiscardUnknown() {
	xxx_messageInfo_PolicyTopicEntryTypeEnum.DiscardUnknown(m)
}

var xxx_messageInfo_PolicyTopicEntryTypeEnum proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("google.ads.googleads.v2.enums.PolicyTopicEntryTypeEnum_PolicyTopicEntryType", PolicyTopicEntryTypeEnum_PolicyTopicEntryType_name, PolicyTopicEntryTypeEnum_PolicyTopicEntryType_value)
	proto.RegisterType((*PolicyTopicEntryTypeEnum)(nil), "google.ads.googleads.v2.enums.PolicyTopicEntryTypeEnum")
}

func init() {
	proto.RegisterFile("google/ads/googleads/v2/enums/policy_topic_entry_type.proto", fileDescriptor_59b40b5d8dcd4000)
}

var fileDescriptor_59b40b5d8dcd4000 = []byte{
	// 379 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xc1, 0xae, 0x93, 0x40,
	0x18, 0x85, 0x05, 0xf5, 0x5e, 0x33, 0x37, 0x2a, 0x12, 0x4d, 0xbc, 0xc6, 0x2e, 0xda, 0x07, 0x18,
	0x12, 0xdc, 0x4d, 0x57, 0x43, 0x99, 0xd6, 0x89, 0x38, 0x10, 0x4a, 0x31, 0x35, 0x24, 0x04, 0x0b,
	0x21, 0x24, 0xed, 0x0c, 0xe9, 0xd0, 0x26, 0xbc, 0x8d, 0x71, 0xe9, 0x3b, 0xf8, 0x02, 0x3e, 0x8a,
	0x0b, 0x9f, 0xc1, 0xcc, 0x20, 0x5d, 0x55, 0x37, 0xe4, 0xf0, 0x9f, 0xff, 0x7c, 0x81, 0xf3, 0x83,
	0x79, 0x2d, 0x44, 0xbd, 0xaf, 0x9c, 0xa2, 0x94, 0xce, 0x20, 0x95, 0x3a, 0xbb, 0x4e, 0xc5, 0x4f,
	0x07, 0xe9, 0xb4, 0x62, 0xdf, 0xec, 0xfa, 0xbc, 0x13, 0x6d, 0xb3, 0xcb, 0x2b, 0xde, 0x1d, 0xfb,
	0xbc, 0xeb, 0xdb, 0x0a, 0xb6, 0x47, 0xd1, 0x09, 0x7b, 0x32, 0x24, 0x60, 0x51, 0x4a, 0x78, 0x09,
	0xc3, 0xb3, 0x0b, 0x75, 0xf8, 0xcd, 0xdb, 0x91, 0xdd, 0x36, 0x4e, 0xc1, 0xb9, 0xe8, 0x8a, 0xae,
	0x11, 0x5c, 0x0e, 0xe1, 0xd9, 0x0f, 0x03, 0xbc, 0x8e, 0x34, 0x3e, 0x51, 0x74, 0xa2, 0xe0, 0x49,
	0xdf, 0x56, 0x84, 0x9f, 0x0e, 0xb3, 0xaf, 0x06, 0x78, 0x79, 0xcd, 0xb4, 0x9f, 0x83, 0xbb, 0x0d,
	0x5b, 0x47, 0x64, 0x41, 0x97, 0x94, 0xf8, 0xd6, 0x03, 0xfb, 0x0e, 0xdc, 0x6e, 0xd8, 0x07, 0x16,
	0x7e, 0x62, 0x96, 0x61, 0x3f, 0x03, 0x20, 0x8a, 0xc3, 0xf7, 0xd4, 0xa3, 0x09, 0xf1, 0x2d, 0x53,
	0x99, 0x01, 0xfd, 0xa8, 0x5f, 0x1e, 0xd9, 0x2f, 0xc0, 0xd3, 0xe5, 0x26, 0x08, 0xb6, 0xf9, 0x38,
	0x7a, 0xa2, 0x68, 0x3e, 0x59, 0x2f, 0x62, 0x1a, 0x25, 0x34, 0x25, 0xd6, 0x63, 0x05, 0xf0, 0xe2,
	0x10, 0xfb, 0x84, 0x51, 0xb6, 0xb2, 0x6e, 0xec, 0x7b, 0xf0, 0x0a, 0xc7, 0x04, 0xe7, 0xe1, 0x32,
	0xa7, 0x2c, 0x21, 0x31, 0x59, 0x27, 0x79, 0xc8, 0x82, 0xad, 0x75, 0xeb, 0xfd, 0x36, 0xc0, 0x74,
	0x27, 0x0e, 0xf0, 0xbf, 0x1d, 0x78, 0xf7, 0xd7, 0xfe, 0x22, 0x52, 0x05, 0x44, 0xc6, 0x67, 0xef,
	0x6f, 0xb6, 0x16, 0xfb, 0x82, 0xd7, 0x50, 0x1c, 0x6b, 0xa7, 0xae, 0xb8, 0xae, 0x67, 0x3c, 0x46,
	0xdb, 0xc8, 0x7f, 0xdc, 0x66, 0xae, 0x9f, 0xdf, 0xcc, 0x87, 0x2b, 0x8c, 0xbf, 0x9b, 0x93, 0xd5,
	0x80, 0xc2, 0xa5, 0x84, 0x83, 0x54, 0x2a, 0x75, 0xa1, 0xaa, 0x53, 0xfe, 0x1c, 0xfd, 0x0c, 0x97,
	0x32, 0xbb, 0xf8, 0x59, 0xea, 0x66, 0xda, 0xff, 0x65, 0x4e, 0x87, 0x21, 0x42, 0xb8, 0x94, 0x08,
	0x5d, 0x36, 0x10, 0x4a, 0x5d, 0x84, 0xf4, 0xce, 0x97, 0x1b, 0xfd, 0x61, 0xef, 0xfe, 0x04, 0x00,
	0x00, 0xff, 0xff, 0xb3, 0x9e, 0x64, 0x83, 0x33, 0x02, 0x00, 0x00,
}
