// Code generated by protoc-gen-go.
// source: golang.conradwood.net/apis/starling/starling.proto
// DO NOT EDIT!

/*
Package starling is a generated protocol buffer package.

It is generated from these files:
	golang.conradwood.net/apis/starling/starling.proto

It has these top-level messages:
	Account
	Amount
	Transaction
	AccountList
	PayeeAccount
	Payee
	PayeeList
*/
package starling

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

type TransactionStatus int32

const (
	TransactionStatus_UNDEFINED     TransactionStatus = 0
	TransactionStatus_SETTLED       TransactionStatus = 1
	TransactionStatus_DECLINED      TransactionStatus = 2
	TransactionStatus_PENDING       TransactionStatus = 3
	TransactionStatus_ACCOUNT_CHECK TransactionStatus = 4
	TransactionStatus_REVERSED      TransactionStatus = 5
)

var TransactionStatus_name = map[int32]string{
	0: "UNDEFINED",
	1: "SETTLED",
	2: "DECLINED",
	3: "PENDING",
	4: "ACCOUNT_CHECK",
	5: "REVERSED",
}
var TransactionStatus_value = map[string]int32{
	"UNDEFINED":     0,
	"SETTLED":       1,
	"DECLINED":      2,
	"PENDING":       3,
	"ACCOUNT_CHECK": 4,
	"REVERSED":      5,
}

func (x TransactionStatus) String() string {
	return proto.EnumName(TransactionStatus_name, int32(x))
}
func (TransactionStatus) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type PayeeType int32

const (
	PayeeType_INDIVIDUAL PayeeType = 0
	PayeeType_BUSINESS   PayeeType = 1
)

var PayeeType_name = map[int32]string{
	0: "INDIVIDUAL",
	1: "BUSINESS",
}
var PayeeType_value = map[string]int32{
	"INDIVIDUAL": 0,
	"BUSINESS":   1,
}

func (x PayeeType) String() string {
	return proto.EnumName(PayeeType_name, int32(x))
}
func (PayeeType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

// comment: message pingrequest
type Account struct {
	AccountUID      string `protobuf:"bytes,1,opt,name=AccountUID" json:"AccountUID,omitempty"`
	DefaultCategory string `protobuf:"bytes,2,opt,name=DefaultCategory" json:"DefaultCategory,omitempty"`
	Currency        string `protobuf:"bytes,3,opt,name=Currency" json:"Currency,omitempty"`
	CreatedAt       uint32 `protobuf:"varint,4,opt,name=CreatedAt" json:"CreatedAt,omitempty"`
	AccountNumber   string `protobuf:"bytes,5,opt,name=AccountNumber" json:"AccountNumber,omitempty"`
	SortCode        string `protobuf:"bytes,6,opt,name=SortCode" json:"SortCode,omitempty"`
	IBAN            string `protobuf:"bytes,7,opt,name=IBAN" json:"IBAN,omitempty"`
	BIC             string `protobuf:"bytes,8,opt,name=BIC" json:"BIC,omitempty"`
}

func (m *Account) Reset()                    { *m = Account{} }
func (m *Account) String() string            { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()               {}
func (*Account) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *Account) GetAccountUID() string {
	if m != nil {
		return m.AccountUID
	}
	return ""
}

func (m *Account) GetDefaultCategory() string {
	if m != nil {
		return m.DefaultCategory
	}
	return ""
}

func (m *Account) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *Account) GetCreatedAt() uint32 {
	if m != nil {
		return m.CreatedAt
	}
	return 0
}

func (m *Account) GetAccountNumber() string {
	if m != nil {
		return m.AccountNumber
	}
	return ""
}

func (m *Account) GetSortCode() string {
	if m != nil {
		return m.SortCode
	}
	return ""
}

func (m *Account) GetIBAN() string {
	if m != nil {
		return m.IBAN
	}
	return ""
}

func (m *Account) GetBIC() string {
	if m != nil {
		return m.BIC
	}
	return ""
}

type Amount struct {
	Currency      string  `protobuf:"bytes,1,opt,name=Currency" json:"Currency,omitempty"`
	Amount        int64   `protobuf:"varint,2,opt,name=Amount" json:"Amount,omitempty"`
	DisplayAmount float64 `protobuf:"fixed64,3,opt,name=DisplayAmount" json:"DisplayAmount,omitempty"`
	Divisor       uint32  `protobuf:"varint,4,opt,name=Divisor" json:"Divisor,omitempty"`
}

func (m *Amount) Reset()                    { *m = Amount{} }
func (m *Amount) String() string            { return proto.CompactTextString(m) }
func (*Amount) ProtoMessage()               {}
func (*Amount) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Amount) GetCurrency() string {
	if m != nil {
		return m.Currency
	}
	return ""
}

func (m *Amount) GetAmount() int64 {
	if m != nil {
		return m.Amount
	}
	return 0
}

func (m *Amount) GetDisplayAmount() float64 {
	if m != nil {
		return m.DisplayAmount
	}
	return 0
}

func (m *Amount) GetDivisor() uint32 {
	if m != nil {
		return m.Divisor
	}
	return 0
}

type Transaction struct {
	UID                 string            `protobuf:"bytes,1,opt,name=UID" json:"UID,omitempty"`
	Outbound            bool              `protobuf:"varint,2,opt,name=Outbound" json:"Outbound,omitempty"`
	Amount              *Amount           `protobuf:"bytes,3,opt,name=Amount" json:"Amount,omitempty"`
	UpdatedAt           uint32            `protobuf:"varint,4,opt,name=UpdatedAt" json:"UpdatedAt,omitempty"`
	TransactionTime     uint32            `protobuf:"varint,5,opt,name=TransactionTime" json:"TransactionTime,omitempty"`
	CounterPartyName    string            `protobuf:"bytes,6,opt,name=CounterPartyName" json:"CounterPartyName,omitempty"`
	Reference           string            `protobuf:"bytes,7,opt,name=Reference" json:"Reference,omitempty"`
	Country             string            `protobuf:"bytes,8,opt,name=Country" json:"Country,omitempty"`
	Account             *Account          `protobuf:"bytes,9,opt,name=Account" json:"Account,omitempty"`
	Status              TransactionStatus `protobuf:"varint,10,opt,name=Status,enum=starling.TransactionStatus" json:"Status,omitempty"`
	CounterPartySubName string            `protobuf:"bytes,11,opt,name=CounterPartySubName" json:"CounterPartySubName,omitempty"`
}

func (m *Transaction) Reset()                    { *m = Transaction{} }
func (m *Transaction) String() string            { return proto.CompactTextString(m) }
func (*Transaction) ProtoMessage()               {}
func (*Transaction) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Transaction) GetUID() string {
	if m != nil {
		return m.UID
	}
	return ""
}

func (m *Transaction) GetOutbound() bool {
	if m != nil {
		return m.Outbound
	}
	return false
}

func (m *Transaction) GetAmount() *Amount {
	if m != nil {
		return m.Amount
	}
	return nil
}

func (m *Transaction) GetUpdatedAt() uint32 {
	if m != nil {
		return m.UpdatedAt
	}
	return 0
}

func (m *Transaction) GetTransactionTime() uint32 {
	if m != nil {
		return m.TransactionTime
	}
	return 0
}

func (m *Transaction) GetCounterPartyName() string {
	if m != nil {
		return m.CounterPartyName
	}
	return ""
}

func (m *Transaction) GetReference() string {
	if m != nil {
		return m.Reference
	}
	return ""
}

func (m *Transaction) GetCountry() string {
	if m != nil {
		return m.Country
	}
	return ""
}

func (m *Transaction) GetAccount() *Account {
	if m != nil {
		return m.Account
	}
	return nil
}

func (m *Transaction) GetStatus() TransactionStatus {
	if m != nil {
		return m.Status
	}
	return TransactionStatus_UNDEFINED
}

func (m *Transaction) GetCounterPartySubName() string {
	if m != nil {
		return m.CounterPartySubName
	}
	return ""
}

type AccountList struct {
	Accounts []*Account `protobuf:"bytes,1,rep,name=Accounts" json:"Accounts,omitempty"`
}

func (m *AccountList) Reset()                    { *m = AccountList{} }
func (m *AccountList) String() string            { return proto.CompactTextString(m) }
func (*AccountList) ProtoMessage()               {}
func (*AccountList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *AccountList) GetAccounts() []*Account {
	if m != nil {
		return m.Accounts
	}
	return nil
}

type PayeeAccount struct {
	Account     *Account `protobuf:"bytes,1,opt,name=Account" json:"Account,omitempty"`
	CountryCode string   `protobuf:"bytes,2,opt,name=CountryCode" json:"CountryCode,omitempty"`
}

func (m *PayeeAccount) Reset()                    { *m = PayeeAccount{} }
func (m *PayeeAccount) String() string            { return proto.CompactTextString(m) }
func (*PayeeAccount) ProtoMessage()               {}
func (*PayeeAccount) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *PayeeAccount) GetAccount() *Account {
	if m != nil {
		return m.Account
	}
	return nil
}

func (m *PayeeAccount) GetCountryCode() string {
	if m != nil {
		return m.CountryCode
	}
	return ""
}

type Payee struct {
	UID      string          `protobuf:"bytes,1,opt,name=UID" json:"UID,omitempty"`
	Name     string          `protobuf:"bytes,2,opt,name=Name" json:"Name,omitempty"`
	Type     PayeeType       `protobuf:"varint,3,opt,name=Type,enum=starling.PayeeType" json:"Type,omitempty"`
	Accounts []*PayeeAccount `protobuf:"bytes,4,rep,name=Accounts" json:"Accounts,omitempty"`
}

func (m *Payee) Reset()                    { *m = Payee{} }
func (m *Payee) String() string            { return proto.CompactTextString(m) }
func (*Payee) ProtoMessage()               {}
func (*Payee) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Payee) GetUID() string {
	if m != nil {
		return m.UID
	}
	return ""
}

func (m *Payee) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Payee) GetType() PayeeType {
	if m != nil {
		return m.Type
	}
	return PayeeType_INDIVIDUAL
}

func (m *Payee) GetAccounts() []*PayeeAccount {
	if m != nil {
		return m.Accounts
	}
	return nil
}

type PayeeList struct {
	Payees []*Payee `protobuf:"bytes,1,rep,name=Payees" json:"Payees,omitempty"`
}

func (m *PayeeList) Reset()                    { *m = PayeeList{} }
func (m *PayeeList) String() string            { return proto.CompactTextString(m) }
func (*PayeeList) ProtoMessage()               {}
func (*PayeeList) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *PayeeList) GetPayees() []*Payee {
	if m != nil {
		return m.Payees
	}
	return nil
}

func init() {
	proto.RegisterType((*Account)(nil), "starling.Account")
	proto.RegisterType((*Amount)(nil), "starling.Amount")
	proto.RegisterType((*Transaction)(nil), "starling.Transaction")
	proto.RegisterType((*AccountList)(nil), "starling.AccountList")
	proto.RegisterType((*PayeeAccount)(nil), "starling.PayeeAccount")
	proto.RegisterType((*Payee)(nil), "starling.Payee")
	proto.RegisterType((*PayeeList)(nil), "starling.PayeeList")
	proto.RegisterEnum("starling.TransactionStatus", TransactionStatus_name, TransactionStatus_value)
	proto.RegisterEnum("starling.PayeeType", PayeeType_name, PayeeType_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for StarlingService service

type StarlingServiceClient interface {
	// comment: rpc ping
	GetAccounts(ctx context.Context, in *common.Void, opts ...grpc.CallOption) (*AccountList, error)
	// get all payees
	GetPayees(ctx context.Context, in *common.Void, opts ...grpc.CallOption) (*PayeeList, error)
	// create a payee
	CreatePayee(ctx context.Context, in *Payee, opts ...grpc.CallOption) (*Payee, error)
}

type starlingServiceClient struct {
	cc *grpc.ClientConn
}

func NewStarlingServiceClient(cc *grpc.ClientConn) StarlingServiceClient {
	return &starlingServiceClient{cc}
}

func (c *starlingServiceClient) GetAccounts(ctx context.Context, in *common.Void, opts ...grpc.CallOption) (*AccountList, error) {
	out := new(AccountList)
	err := grpc.Invoke(ctx, "/starling.StarlingService/GetAccounts", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *starlingServiceClient) GetPayees(ctx context.Context, in *common.Void, opts ...grpc.CallOption) (*PayeeList, error) {
	out := new(PayeeList)
	err := grpc.Invoke(ctx, "/starling.StarlingService/GetPayees", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *starlingServiceClient) CreatePayee(ctx context.Context, in *Payee, opts ...grpc.CallOption) (*Payee, error) {
	out := new(Payee)
	err := grpc.Invoke(ctx, "/starling.StarlingService/CreatePayee", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for StarlingService service

type StarlingServiceServer interface {
	// comment: rpc ping
	GetAccounts(context.Context, *common.Void) (*AccountList, error)
	// get all payees
	GetPayees(context.Context, *common.Void) (*PayeeList, error)
	// create a payee
	CreatePayee(context.Context, *Payee) (*Payee, error)
}

func RegisterStarlingServiceServer(s *grpc.Server, srv StarlingServiceServer) {
	s.RegisterService(&_StarlingService_serviceDesc, srv)
}

func _StarlingService_GetAccounts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StarlingServiceServer).GetAccounts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/starling.StarlingService/GetAccounts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StarlingServiceServer).GetAccounts(ctx, req.(*common.Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _StarlingService_GetPayees_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(common.Void)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StarlingServiceServer).GetPayees(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/starling.StarlingService/GetPayees",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StarlingServiceServer).GetPayees(ctx, req.(*common.Void))
	}
	return interceptor(ctx, in, info, handler)
}

func _StarlingService_CreatePayee_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Payee)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(StarlingServiceServer).CreatePayee(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/starling.StarlingService/CreatePayee",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(StarlingServiceServer).CreatePayee(ctx, req.(*Payee))
	}
	return interceptor(ctx, in, info, handler)
}

var _StarlingService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "starling.StarlingService",
	HandlerType: (*StarlingServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetAccounts",
			Handler:    _StarlingService_GetAccounts_Handler,
		},
		{
			MethodName: "GetPayees",
			Handler:    _StarlingService_GetPayees_Handler,
		},
		{
			MethodName: "CreatePayee",
			Handler:    _StarlingService_CreatePayee_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "golang.conradwood.net/apis/starling/starling.proto",
}

func init() { proto.RegisterFile("golang.conradwood.net/apis/starling/starling.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 745 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x54, 0x5d, 0x6f, 0xda, 0x30,
	0x14, 0x6d, 0x0a, 0xe5, 0xe3, 0xa6, 0xb4, 0xa9, 0xab, 0x55, 0x11, 0xab, 0x26, 0xc4, 0x3a, 0x95,
	0x75, 0x1a, 0x9d, 0xe8, 0x1e, 0xf7, 0x02, 0x49, 0xd6, 0x45, 0x43, 0x69, 0x95, 0x40, 0xdf, 0xa6,
	0x29, 0x80, 0x8b, 0xa2, 0x41, 0x8c, 0x1c, 0xa7, 0x13, 0xd2, 0x7e, 0xc1, 0xfe, 0xc7, 0xf6, 0x0f,
	0xf7, 0x3e, 0xd9, 0x71, 0x42, 0x1a, 0xaa, 0xed, 0x09, 0xdf, 0x73, 0xcf, 0x8d, 0xcf, 0x39, 0x36,
	0x86, 0xde, 0x9c, 0x2c, 0xfc, 0x70, 0xde, 0x9d, 0x92, 0x90, 0xfa, 0xb3, 0xef, 0x84, 0xcc, 0xba,
	0x21, 0x66, 0x97, 0xfe, 0x2a, 0x88, 0x2e, 0x23, 0xe6, 0xd3, 0x45, 0x10, 0xce, 0xb3, 0x45, 0x77,
	0x45, 0x09, 0x23, 0xa8, 0x96, 0xd6, 0xcd, 0xee, 0x3f, 0xa6, 0xa7, 0x64, 0xb9, 0x24, 0xa1, 0xfc,
	0x49, 0x26, 0xdb, 0x7f, 0x14, 0xa8, 0xf6, 0xa7, 0x53, 0x12, 0x87, 0x0c, 0xbd, 0x00, 0x90, 0xcb,
	0xb1, 0x6d, 0xea, 0x4a, 0x4b, 0xe9, 0xd4, 0xdd, 0x1c, 0x82, 0x3a, 0x70, 0x68, 0xe2, 0x7b, 0x3f,
	0x5e, 0x30, 0xc3, 0x67, 0x78, 0x4e, 0xe8, 0x5a, 0xdf, 0x15, 0xa4, 0x22, 0x8c, 0x9a, 0x50, 0x33,
	0x62, 0x4a, 0x71, 0x38, 0x5d, 0xeb, 0x25, 0x41, 0xc9, 0x6a, 0x74, 0x0a, 0x75, 0x83, 0x62, 0x9f,
	0xe1, 0x59, 0x9f, 0xe9, 0xe5, 0x96, 0xd2, 0x69, 0xb8, 0x1b, 0x00, 0x9d, 0x41, 0x43, 0xee, 0xe8,
	0xc4, 0xcb, 0x09, 0xa6, 0xfa, 0x9e, 0x18, 0x7f, 0x0c, 0xf2, 0xef, 0x7b, 0x84, 0x32, 0x83, 0xcc,
	0xb0, 0x5e, 0x49, 0xbe, 0x9f, 0xd6, 0x08, 0x41, 0xd9, 0x1e, 0xf4, 0x1d, 0xbd, 0x2a, 0x70, 0xb1,
	0x46, 0x1a, 0x94, 0x06, 0xb6, 0xa1, 0xd7, 0x04, 0xc4, 0x97, 0xed, 0x1f, 0x50, 0xe9, 0x2f, 0x85,
	0xeb, 0xbc, 0x56, 0xa5, 0xa0, 0xf5, 0x24, 0x65, 0x09, 0xa3, 0x25, 0x37, 0x9d, 0x39, 0x83, 0x86,
	0x19, 0x44, 0xab, 0x85, 0xbf, 0x96, 0x6d, 0x6e, 0x52, 0x71, 0x1f, 0x83, 0x48, 0x87, 0xaa, 0x19,
	0x3c, 0x04, 0x11, 0xa1, 0xd2, 0x67, 0x5a, 0xb6, 0x7f, 0x95, 0x40, 0x1d, 0x51, 0x3f, 0x8c, 0xfc,
	0x29, 0x0b, 0x48, 0xc8, 0xf5, 0x6d, 0x22, 0xe7, 0x4b, 0xae, 0xea, 0x26, 0x66, 0x13, 0x12, 0x87,
	0x33, 0xb1, 0x77, 0xcd, 0xcd, 0x6a, 0xd4, 0xc9, 0x54, 0xf1, 0x6d, 0xd5, 0x9e, 0xd6, 0xcd, 0xae,
	0x43, 0x82, 0x67, 0x3a, 0x4f, 0xa1, 0x3e, 0x5e, 0xcd, 0x1e, 0x67, 0x9d, 0x01, 0xfc, 0x3c, 0x73,
	0x22, 0x46, 0xc1, 0x12, 0x8b, 0xb4, 0x1b, 0x6e, 0x11, 0x46, 0x17, 0xa0, 0x19, 0xfc, 0x83, 0x98,
	0xde, 0xfa, 0x94, 0xad, 0x1d, 0x7f, 0x99, 0xe6, 0xbe, 0x85, 0xf3, 0x3d, 0x5d, 0x7c, 0x8f, 0x79,
	0x80, 0x58, 0x1e, 0xc2, 0x06, 0xe0, 0x99, 0x88, 0x09, 0xba, 0x96, 0xa7, 0x91, 0x96, 0xe8, 0x4d,
	0x76, 0x11, 0xf5, 0xba, 0xb0, 0x75, 0x94, 0xb3, 0x95, 0x34, 0xdc, 0xec, 0xaa, 0x5e, 0x41, 0xc5,
	0x63, 0x3e, 0x8b, 0x23, 0x1d, 0x5a, 0x4a, 0xe7, 0xa0, 0xf7, 0x7c, 0xc3, 0xcd, 0x69, 0x4f, 0x28,
	0xae, 0xa4, 0xa2, 0x77, 0x70, 0x9c, 0x57, 0xeb, 0xc5, 0x13, 0x61, 0x44, 0x15, 0x3a, 0x9e, 0x6a,
	0xb5, 0x3f, 0x80, 0x2a, 0x77, 0x1c, 0x06, 0x11, 0x43, 0x6f, 0xa1, 0x26, 0xcb, 0x48, 0x57, 0x5a,
	0xa5, 0xa7, 0x35, 0x66, 0x94, 0xf6, 0x17, 0xd8, 0xbf, 0xf5, 0xd7, 0x18, 0xa7, 0xa2, 0x73, 0x0e,
	0x95, 0xff, 0x3a, 0x6c, 0x81, 0x2a, 0x93, 0x11, 0xb7, 0x3c, 0xf9, 0xa3, 0xe5, 0xa1, 0xf6, 0x4f,
	0x05, 0xf6, 0xc4, 0xf7, 0x9f, 0xb8, 0x3e, 0x08, 0xca, 0xc2, 0x5b, 0x32, 0x26, 0xd6, 0xe8, 0x1c,
	0xca, 0xa3, 0xf5, 0x0a, 0x8b, 0x4b, 0x73, 0xd0, 0x3b, 0xde, 0xec, 0x2d, 0x3e, 0xc2, 0x5b, 0xae,
	0x20, 0xa0, 0x5e, 0xce, 0x66, 0x59, 0xd8, 0x3c, 0x29, 0x90, 0xb7, 0xbd, 0xbe, 0x87, 0xba, 0xe8,
	0x88, 0x9c, 0xce, 0xa1, 0x22, 0x8a, 0x34, 0xa5, 0xc3, 0xc2, 0xb8, 0x2b, 0xdb, 0x17, 0xdf, 0xe0,
	0x68, 0xeb, 0xb8, 0x50, 0x03, 0xea, 0x63, 0xc7, 0xb4, 0x3e, 0xda, 0x8e, 0x65, 0x6a, 0x3b, 0x48,
	0x85, 0xaa, 0x67, 0x8d, 0x46, 0x43, 0xcb, 0xd4, 0x14, 0xb4, 0x0f, 0x35, 0xd3, 0x32, 0x86, 0xa2,
	0xb5, 0xcb, 0x5b, 0xb7, 0x96, 0x63, 0xda, 0xce, 0xb5, 0x56, 0x42, 0x47, 0xd0, 0xe8, 0x1b, 0xc6,
	0xcd, 0xd8, 0x19, 0x7d, 0x35, 0x3e, 0x59, 0xc6, 0x67, 0xad, 0xcc, 0xd9, 0xae, 0x75, 0x67, 0xb9,
	0x9e, 0x65, 0x6a, 0x7b, 0x17, 0xaf, 0xa5, 0x44, 0xe1, 0xf1, 0x00, 0xc0, 0x76, 0x4c, 0xfb, 0xce,
	0x36, 0xc7, 0xfd, 0xa1, 0xb6, 0xc3, 0xa9, 0x83, 0xb1, 0x67, 0x3b, 0x96, 0xe7, 0x69, 0x4a, 0xef,
	0xb7, 0x02, 0x87, 0x9e, 0x94, 0xec, 0x61, 0xfa, 0x10, 0x4c, 0x79, 0x2a, 0xea, 0x35, 0x66, 0xa9,
	0x61, 0xb4, 0xdf, 0x95, 0xef, 0xe8, 0x1d, 0x09, 0x66, 0xcd, 0x67, 0x5b, 0x27, 0x29, 0x82, 0xe8,
	0x42, 0xfd, 0x1a, 0xb3, 0xc4, 0x6c, 0x61, 0xa2, 0x98, 0xbf, 0xe0, 0x5f, 0x82, 0x9a, 0x3c, 0x85,
	0xc9, 0xb9, 0x16, 0x73, 0x6b, 0x16, 0x81, 0xc1, 0x2b, 0x78, 0x19, 0x62, 0x96, 0x7f, 0xed, 0xe5,
	0xfb, 0xcf, 0x1f, 0xfc, 0x8c, 0x3c, 0xa9, 0x88, 0xc7, 0xfe, 0xea, 0x6f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x1b, 0x50, 0xc6, 0xe2, 0x5c, 0x06, 0x00, 0x00,
}