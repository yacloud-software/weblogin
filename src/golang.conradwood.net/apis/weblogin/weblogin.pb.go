// Code generated by protoc-gen-go.
// source: protos/golang.conradwood.net/apis/weblogin/weblogin.proto
// DO NOT EDIT!

/*
Package weblogin is a generated protocol buffer package.

It is generated from these files:
	protos/golang.conradwood.net/apis/weblogin/weblogin.proto

It has these top-level messages:
	BasicAuthRequest
	AuthResponse
	WebloginRequest
	EmailPageResponse
	WebloginResponse
	State
	StateResponse
	RegisterState
	RegisterProto
	Email
	ActivityLog
	V3State
*/
package weblogin

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import auth "golang.conradwood.net/apis/auth"
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

type BasicAuthRequest struct {
	// the stuff extracted from "Authorization" Header
	Username string `protobuf:"bytes,1,opt,name=Username" json:"Username,omitempty"`
	Password string `protobuf:"bytes,2,opt,name=Password" json:"Password,omitempty"`
}

func (m *BasicAuthRequest) Reset()                    { *m = BasicAuthRequest{} }
func (m *BasicAuthRequest) String() string            { return proto.CompactTextString(m) }
func (*BasicAuthRequest) ProtoMessage()               {}
func (*BasicAuthRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *BasicAuthRequest) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *BasicAuthRequest) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type AuthResponse struct {
	IsValid  bool       `protobuf:"varint,1,opt,name=IsValid" json:"IsValid,omitempty"`
	Response *auth.User `protobuf:"bytes,2,opt,name=Response" json:"Response,omitempty"`
}

func (m *AuthResponse) Reset()                    { *m = AuthResponse{} }
func (m *AuthResponse) String() string            { return proto.CompactTextString(m) }
func (*AuthResponse) ProtoMessage()               {}
func (*AuthResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *AuthResponse) GetIsValid() bool {
	if m != nil {
		return m.IsValid
	}
	return false
}

func (m *AuthResponse) GetResponse() *auth.User {
	if m != nil {
		return m.Response
	}
	return nil
}

type WebloginRequest struct {
	Method    string             `protobuf:"bytes,1,opt,name=Method" json:"Method,omitempty"`
	Scheme    string             `protobuf:"bytes,2,opt,name=Scheme" json:"Scheme,omitempty"`
	Host      string             `protobuf:"bytes,3,opt,name=Host" json:"Host,omitempty"`
	Path      string             `protobuf:"bytes,4,opt,name=Path" json:"Path,omitempty"`
	Query     string             `protobuf:"bytes,5,opt,name=Query" json:"Query,omitempty"`
	Body      string             `protobuf:"bytes,6,opt,name=Body" json:"Body,omitempty"`
	Submitted map[string]string  `protobuf:"bytes,7,rep,name=Submitted" json:"Submitted,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Peer      string             `protobuf:"bytes,8,opt,name=Peer" json:"Peer,omitempty"`
	Cookies   []*h2gproxy.Cookie `protobuf:"bytes,9,rep,name=Cookies" json:"Cookies,omitempty"`
	UserAgent string             `protobuf:"bytes,10,opt,name=UserAgent" json:"UserAgent,omitempty"`
}

func (m *WebloginRequest) Reset()                    { *m = WebloginRequest{} }
func (m *WebloginRequest) String() string            { return proto.CompactTextString(m) }
func (*WebloginRequest) ProtoMessage()               {}
func (*WebloginRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *WebloginRequest) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

func (m *WebloginRequest) GetScheme() string {
	if m != nil {
		return m.Scheme
	}
	return ""
}

func (m *WebloginRequest) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *WebloginRequest) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *WebloginRequest) GetQuery() string {
	if m != nil {
		return m.Query
	}
	return ""
}

func (m *WebloginRequest) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *WebloginRequest) GetSubmitted() map[string]string {
	if m != nil {
		return m.Submitted
	}
	return nil
}

func (m *WebloginRequest) GetPeer() string {
	if m != nil {
		return m.Peer
	}
	return ""
}

func (m *WebloginRequest) GetCookies() []*h2gproxy.Cookie {
	if m != nil {
		return m.Cookies
	}
	return nil
}

func (m *WebloginRequest) GetUserAgent() string {
	if m != nil {
		return m.UserAgent
	}
	return ""
}

type EmailPageResponse struct {
	HTML     string            `protobuf:"bytes,1,opt,name=HTML" json:"HTML,omitempty"`
	Verified bool              `protobuf:"varint,2,opt,name=Verified" json:"Verified,omitempty"`
	User     *auth.User        `protobuf:"bytes,3,opt,name=User" json:"User,omitempty"`
	Headers  map[string]string `protobuf:"bytes,4,rep,name=Headers" json:"Headers,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *EmailPageResponse) Reset()                    { *m = EmailPageResponse{} }
func (m *EmailPageResponse) String() string            { return proto.CompactTextString(m) }
func (*EmailPageResponse) ProtoMessage()               {}
func (*EmailPageResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *EmailPageResponse) GetHTML() string {
	if m != nil {
		return m.HTML
	}
	return ""
}

func (m *EmailPageResponse) GetVerified() bool {
	if m != nil {
		return m.Verified
	}
	return false
}

func (m *EmailPageResponse) GetUser() *auth.User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *EmailPageResponse) GetHeaders() map[string]string {
	if m != nil {
		return m.Headers
	}
	return nil
}

type WebloginResponse struct {
	Body                  []byte             `protobuf:"bytes,1,opt,name=Body,proto3" json:"Body,omitempty"`
	Authenticated         bool               `protobuf:"varint,2,opt,name=Authenticated" json:"Authenticated,omitempty"`
	User                  *auth.User         `protobuf:"bytes,3,opt,name=User" json:"User,omitempty"`
	Token                 string             `protobuf:"bytes,4,opt,name=Token" json:"Token,omitempty"`
	CookieLivetime        uint32             `protobuf:"varint,5,opt,name=CookieLivetime" json:"CookieLivetime,omitempty"`
	RedirectTo            string             `protobuf:"bytes,6,opt,name=RedirectTo" json:"RedirectTo,omitempty"`
	Cookies               []*h2gproxy.Cookie `protobuf:"bytes,7,rep,name=Cookies" json:"Cookies,omitempty"`
	Headers               map[string]string  `protobuf:"bytes,8,rep,name=Headers" json:"Headers,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	ForceGetAfterRedirect bool               `protobuf:"varint,9,opt,name=ForceGetAfterRedirect" json:"ForceGetAfterRedirect,omitempty"`
	PeerIsDosing          bool               `protobuf:"varint,10,opt,name=PeerIsDosing" json:"PeerIsDosing,omitempty"`
	PeerIP                string             `protobuf:"bytes,11,opt,name=PeerIP" json:"PeerIP,omitempty"`
	HTTPCode              uint32             `protobuf:"varint,12,opt,name=HTTPCode" json:"HTTPCode,omitempty"`
	MimeType              string             `protobuf:"bytes,13,opt,name=MimeType" json:"MimeType,omitempty"`
}

func (m *WebloginResponse) Reset()                    { *m = WebloginResponse{} }
func (m *WebloginResponse) String() string            { return proto.CompactTextString(m) }
func (*WebloginResponse) ProtoMessage()               {}
func (*WebloginResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *WebloginResponse) GetBody() []byte {
	if m != nil {
		return m.Body
	}
	return nil
}

func (m *WebloginResponse) GetAuthenticated() bool {
	if m != nil {
		return m.Authenticated
	}
	return false
}

func (m *WebloginResponse) GetUser() *auth.User {
	if m != nil {
		return m.User
	}
	return nil
}

func (m *WebloginResponse) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *WebloginResponse) GetCookieLivetime() uint32 {
	if m != nil {
		return m.CookieLivetime
	}
	return 0
}

func (m *WebloginResponse) GetRedirectTo() string {
	if m != nil {
		return m.RedirectTo
	}
	return ""
}

func (m *WebloginResponse) GetCookies() []*h2gproxy.Cookie {
	if m != nil {
		return m.Cookies
	}
	return nil
}

func (m *WebloginResponse) GetHeaders() map[string]string {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *WebloginResponse) GetForceGetAfterRedirect() bool {
	if m != nil {
		return m.ForceGetAfterRedirect
	}
	return false
}

func (m *WebloginResponse) GetPeerIsDosing() bool {
	if m != nil {
		return m.PeerIsDosing
	}
	return false
}

func (m *WebloginResponse) GetPeerIP() string {
	if m != nil {
		return m.PeerIP
	}
	return ""
}

func (m *WebloginResponse) GetHTTPCode() uint32 {
	if m != nil {
		return m.HTTPCode
	}
	return 0
}

func (m *WebloginResponse) GetMimeType() string {
	if m != nil {
		return m.MimeType
	}
	return ""
}

// this protobuf is never exposed to the user. it is held server-side and referred to by a shortlived magic
type State struct {
	TriggerHost  string `protobuf:"bytes,1,opt,name=TriggerHost" json:"TriggerHost,omitempty"`
	TriggerPath  string `protobuf:"bytes,2,opt,name=TriggerPath" json:"TriggerPath,omitempty"`
	TriggerQuery string `protobuf:"bytes,3,opt,name=TriggerQuery" json:"TriggerQuery,omitempty"`
	Token        string `protobuf:"bytes,4,opt,name=Token" json:"Token,omitempty"`
}

func (m *State) Reset()                    { *m = State{} }
func (m *State) String() string            { return proto.CompactTextString(m) }
func (*State) ProtoMessage()               {}
func (*State) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *State) GetTriggerHost() string {
	if m != nil {
		return m.TriggerHost
	}
	return ""
}

func (m *State) GetTriggerPath() string {
	if m != nil {
		return m.TriggerPath
	}
	return ""
}

func (m *State) GetTriggerQuery() string {
	if m != nil {
		return m.TriggerQuery
	}
	return ""
}

func (m *State) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type StateResponse struct {
	YacloudWebloginState string `protobuf:"bytes,1,opt,name=YacloudWebloginState" json:"YacloudWebloginState,omitempty"`
	URLStateName         string `protobuf:"bytes,2,opt,name=URLStateName" json:"URLStateName,omitempty"`
}

func (m *StateResponse) Reset()                    { *m = StateResponse{} }
func (m *StateResponse) String() string            { return proto.CompactTextString(m) }
func (*StateResponse) ProtoMessage()               {}
func (*StateResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *StateResponse) GetYacloudWebloginState() string {
	if m != nil {
		return m.YacloudWebloginState
	}
	return ""
}

func (m *StateResponse) GetURLStateName() string {
	if m != nil {
		return m.URLStateName
	}
	return ""
}

type RegisterState struct {
	Host    string `protobuf:"bytes,1,opt,name=Host" json:"Host,omitempty"`
	Email   string `protobuf:"bytes,2,opt,name=Email" json:"Email,omitempty"`
	Created uint32 `protobuf:"varint,3,opt,name=Created" json:"Created,omitempty"`
	Magic   string `protobuf:"bytes,4,opt,name=Magic" json:"Magic,omitempty"`
}

func (m *RegisterState) Reset()                    { *m = RegisterState{} }
func (m *RegisterState) String() string            { return proto.CompactTextString(m) }
func (*RegisterState) ProtoMessage()               {}
func (*RegisterState) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *RegisterState) GetHost() string {
	if m != nil {
		return m.Host
	}
	return ""
}

func (m *RegisterState) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *RegisterState) GetCreated() uint32 {
	if m != nil {
		return m.Created
	}
	return 0
}

func (m *RegisterState) GetMagic() string {
	if m != nil {
		return m.Magic
	}
	return ""
}

type RegisterProto struct {
	State     string `protobuf:"bytes,1,opt,name=State" json:"State,omitempty"`
	Signature []byte `protobuf:"bytes,2,opt,name=Signature,proto3" json:"Signature,omitempty"`
}

func (m *RegisterProto) Reset()                    { *m = RegisterProto{} }
func (m *RegisterProto) String() string            { return proto.CompactTextString(m) }
func (*RegisterProto) ProtoMessage()               {}
func (*RegisterProto) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *RegisterProto) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *RegisterProto) GetSignature() []byte {
	if m != nil {
		return m.Signature
	}
	return nil
}

type Email struct {
	Subject string `protobuf:"bytes,1,opt,name=Subject" json:"Subject,omitempty"`
	Body    string `protobuf:"bytes,2,opt,name=Body" json:"Body,omitempty"`
	Link    string `protobuf:"bytes,3,opt,name=Link" json:"Link,omitempty"`
}

func (m *Email) Reset()                    { *m = Email{} }
func (m *Email) String() string            { return proto.CompactTextString(m) }
func (*Email) ProtoMessage()               {}
func (*Email) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *Email) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *Email) GetBody() string {
	if m != nil {
		return m.Body
	}
	return ""
}

func (m *Email) GetLink() string {
	if m != nil {
		return m.Link
	}
	return ""
}

type ActivityLog struct {
	ID          uint64 `protobuf:"varint,1,opt,name=ID" json:"ID,omitempty"`
	IP          string `protobuf:"bytes,2,opt,name=IP" json:"IP,omitempty"`
	UserID      string `protobuf:"bytes,3,opt,name=UserID" json:"UserID,omitempty"`
	Email       string `protobuf:"bytes,4,opt,name=Email" json:"Email,omitempty"`
	TriggerHost string `protobuf:"bytes,5,opt,name=TriggerHost" json:"TriggerHost,omitempty"`
	Occured     uint32 `protobuf:"varint,6,opt,name=Occured" json:"Occured,omitempty"`
	LogMessage  string `protobuf:"bytes,7,opt,name=LogMessage" json:"LogMessage,omitempty"`
}

func (m *ActivityLog) Reset()                    { *m = ActivityLog{} }
func (m *ActivityLog) String() string            { return proto.CompactTextString(m) }
func (*ActivityLog) ProtoMessage()               {}
func (*ActivityLog) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *ActivityLog) GetID() uint64 {
	if m != nil {
		return m.ID
	}
	return 0
}

func (m *ActivityLog) GetIP() string {
	if m != nil {
		return m.IP
	}
	return ""
}

func (m *ActivityLog) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *ActivityLog) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *ActivityLog) GetTriggerHost() string {
	if m != nil {
		return m.TriggerHost
	}
	return ""
}

func (m *ActivityLog) GetOccured() uint32 {
	if m != nil {
		return m.Occured
	}
	return 0
}

func (m *ActivityLog) GetLogMessage() string {
	if m != nil {
		return m.LogMessage
	}
	return ""
}

// this is the state of a process from beginning to end.
type V3State struct {
	TriggerHost           string `protobuf:"bytes,1,opt,name=TriggerHost" json:"TriggerHost,omitempty"`
	TriggerPath           string `protobuf:"bytes,2,opt,name=TriggerPath" json:"TriggerPath,omitempty"`
	TriggerQuery          string `protobuf:"bytes,3,opt,name=TriggerQuery" json:"TriggerQuery,omitempty"`
	HaveTriedGlobalCookie bool   `protobuf:"varint,4,opt,name=HaveTriedGlobalCookie" json:"HaveTriedGlobalCookie,omitempty"`
	RequestID             uint64 `protobuf:"varint,5,opt,name=RequestID" json:"RequestID,omitempty"`
}

func (m *V3State) Reset()                    { *m = V3State{} }
func (m *V3State) String() string            { return proto.CompactTextString(m) }
func (*V3State) ProtoMessage()               {}
func (*V3State) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *V3State) GetTriggerHost() string {
	if m != nil {
		return m.TriggerHost
	}
	return ""
}

func (m *V3State) GetTriggerPath() string {
	if m != nil {
		return m.TriggerPath
	}
	return ""
}

func (m *V3State) GetTriggerQuery() string {
	if m != nil {
		return m.TriggerQuery
	}
	return ""
}

func (m *V3State) GetHaveTriedGlobalCookie() bool {
	if m != nil {
		return m.HaveTriedGlobalCookie
	}
	return false
}

func (m *V3State) GetRequestID() uint64 {
	if m != nil {
		return m.RequestID
	}
	return 0
}

func init() {
	proto.RegisterType((*BasicAuthRequest)(nil), "weblogin.BasicAuthRequest")
	proto.RegisterType((*AuthResponse)(nil), "weblogin.AuthResponse")
	proto.RegisterType((*WebloginRequest)(nil), "weblogin.WebloginRequest")
	proto.RegisterType((*EmailPageResponse)(nil), "weblogin.EmailPageResponse")
	proto.RegisterType((*WebloginResponse)(nil), "weblogin.WebloginResponse")
	proto.RegisterType((*State)(nil), "weblogin.State")
	proto.RegisterType((*StateResponse)(nil), "weblogin.StateResponse")
	proto.RegisterType((*RegisterState)(nil), "weblogin.RegisterState")
	proto.RegisterType((*RegisterProto)(nil), "weblogin.RegisterProto")
	proto.RegisterType((*Email)(nil), "weblogin.Email")
	proto.RegisterType((*ActivityLog)(nil), "weblogin.ActivityLog")
	proto.RegisterType((*V3State)(nil), "weblogin.V3State")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Weblogin service

type WebloginClient interface {
	CreateRegisterEmail(ctx context.Context, in *RegisterState, opts ...grpc.CallOption) (*Email, error)
	// sometimes we need basic auth if so, we get a request to same URL with a an
	// Authorization Header
	// if so, the h2gproxy needs to verify the header
	IsBasicAuthValid(ctx context.Context, in *BasicAuthRequest, opts ...grpc.CallOption) (*AuthResponse, error)
	// serve the login page
	GetLoginPage(ctx context.Context, in *WebloginRequest, opts ...grpc.CallOption) (*WebloginResponse, error)
	// generic html serve thing for h2gproxy
	ServeHTML(ctx context.Context, in *WebloginRequest, opts ...grpc.CallOption) (*WebloginResponse, error)
	// called if a user is authenticated but the users' email address is not yet verified
	// returns true if email is now verified
	GetVerifyEmail(ctx context.Context, in *WebloginRequest, opts ...grpc.CallOption) (*EmailPageResponse, error)
	// helper for h2gproxy. verify a "weblogin=foobar" url. Body is ignored in the response.
	// cookie is helpful though
	VerifyURL(ctx context.Context, in *WebloginRequest, opts ...grpc.CallOption) (*WebloginResponse, error)
	// helper for h2gproxy. create a state (before redirecting browser to weblogin)
	SaveState(ctx context.Context, in *WebloginRequest, opts ...grpc.CallOption) (*StateResponse, error)
}

type webloginClient struct {
	cc *grpc.ClientConn
}

func NewWebloginClient(cc *grpc.ClientConn) WebloginClient {
	return &webloginClient{cc}
}

func (c *webloginClient) CreateRegisterEmail(ctx context.Context, in *RegisterState, opts ...grpc.CallOption) (*Email, error) {
	out := new(Email)
	err := grpc.Invoke(ctx, "/weblogin.Weblogin/CreateRegisterEmail", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webloginClient) IsBasicAuthValid(ctx context.Context, in *BasicAuthRequest, opts ...grpc.CallOption) (*AuthResponse, error) {
	out := new(AuthResponse)
	err := grpc.Invoke(ctx, "/weblogin.Weblogin/IsBasicAuthValid", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webloginClient) GetLoginPage(ctx context.Context, in *WebloginRequest, opts ...grpc.CallOption) (*WebloginResponse, error) {
	out := new(WebloginResponse)
	err := grpc.Invoke(ctx, "/weblogin.Weblogin/GetLoginPage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webloginClient) ServeHTML(ctx context.Context, in *WebloginRequest, opts ...grpc.CallOption) (*WebloginResponse, error) {
	out := new(WebloginResponse)
	err := grpc.Invoke(ctx, "/weblogin.Weblogin/ServeHTML", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webloginClient) GetVerifyEmail(ctx context.Context, in *WebloginRequest, opts ...grpc.CallOption) (*EmailPageResponse, error) {
	out := new(EmailPageResponse)
	err := grpc.Invoke(ctx, "/weblogin.Weblogin/GetVerifyEmail", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webloginClient) VerifyURL(ctx context.Context, in *WebloginRequest, opts ...grpc.CallOption) (*WebloginResponse, error) {
	out := new(WebloginResponse)
	err := grpc.Invoke(ctx, "/weblogin.Weblogin/VerifyURL", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *webloginClient) SaveState(ctx context.Context, in *WebloginRequest, opts ...grpc.CallOption) (*StateResponse, error) {
	out := new(StateResponse)
	err := grpc.Invoke(ctx, "/weblogin.Weblogin/SaveState", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Weblogin service

type WebloginServer interface {
	CreateRegisterEmail(context.Context, *RegisterState) (*Email, error)
	// sometimes we need basic auth if so, we get a request to same URL with a an
	// Authorization Header
	// if so, the h2gproxy needs to verify the header
	IsBasicAuthValid(context.Context, *BasicAuthRequest) (*AuthResponse, error)
	// serve the login page
	GetLoginPage(context.Context, *WebloginRequest) (*WebloginResponse, error)
	// generic html serve thing for h2gproxy
	ServeHTML(context.Context, *WebloginRequest) (*WebloginResponse, error)
	// called if a user is authenticated but the users' email address is not yet verified
	// returns true if email is now verified
	GetVerifyEmail(context.Context, *WebloginRequest) (*EmailPageResponse, error)
	// helper for h2gproxy. verify a "weblogin=foobar" url. Body is ignored in the response.
	// cookie is helpful though
	VerifyURL(context.Context, *WebloginRequest) (*WebloginResponse, error)
	// helper for h2gproxy. create a state (before redirecting browser to weblogin)
	SaveState(context.Context, *WebloginRequest) (*StateResponse, error)
}

func RegisterWebloginServer(s *grpc.Server, srv WebloginServer) {
	s.RegisterService(&_Weblogin_serviceDesc, srv)
}

func _Weblogin_CreateRegisterEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterState)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebloginServer).CreateRegisterEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/weblogin.Weblogin/CreateRegisterEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebloginServer).CreateRegisterEmail(ctx, req.(*RegisterState))
	}
	return interceptor(ctx, in, info, handler)
}

func _Weblogin_IsBasicAuthValid_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(BasicAuthRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebloginServer).IsBasicAuthValid(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/weblogin.Weblogin/IsBasicAuthValid",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebloginServer).IsBasicAuthValid(ctx, req.(*BasicAuthRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Weblogin_GetLoginPage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WebloginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebloginServer).GetLoginPage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/weblogin.Weblogin/GetLoginPage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebloginServer).GetLoginPage(ctx, req.(*WebloginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Weblogin_ServeHTML_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WebloginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebloginServer).ServeHTML(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/weblogin.Weblogin/ServeHTML",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebloginServer).ServeHTML(ctx, req.(*WebloginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Weblogin_GetVerifyEmail_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WebloginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebloginServer).GetVerifyEmail(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/weblogin.Weblogin/GetVerifyEmail",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebloginServer).GetVerifyEmail(ctx, req.(*WebloginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Weblogin_VerifyURL_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WebloginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebloginServer).VerifyURL(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/weblogin.Weblogin/VerifyURL",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebloginServer).VerifyURL(ctx, req.(*WebloginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Weblogin_SaveState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WebloginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(WebloginServer).SaveState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/weblogin.Weblogin/SaveState",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(WebloginServer).SaveState(ctx, req.(*WebloginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Weblogin_serviceDesc = grpc.ServiceDesc{
	ServiceName: "weblogin.Weblogin",
	HandlerType: (*WebloginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateRegisterEmail",
			Handler:    _Weblogin_CreateRegisterEmail_Handler,
		},
		{
			MethodName: "IsBasicAuthValid",
			Handler:    _Weblogin_IsBasicAuthValid_Handler,
		},
		{
			MethodName: "GetLoginPage",
			Handler:    _Weblogin_GetLoginPage_Handler,
		},
		{
			MethodName: "ServeHTML",
			Handler:    _Weblogin_ServeHTML_Handler,
		},
		{
			MethodName: "GetVerifyEmail",
			Handler:    _Weblogin_GetVerifyEmail_Handler,
		},
		{
			MethodName: "VerifyURL",
			Handler:    _Weblogin_VerifyURL_Handler,
		},
		{
			MethodName: "SaveState",
			Handler:    _Weblogin_SaveState_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protos/golang.conradwood.net/apis/weblogin/weblogin.proto",
}

func init() {
	proto.RegisterFile("protos/golang.conradwood.net/apis/weblogin/weblogin.proto", fileDescriptor0)
}

var fileDescriptor0 = []byte{
	// 1092 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xbc, 0x56, 0xdd, 0x6e, 0x1b, 0x45,
	0x14, 0x96, 0x1d, 0x3b, 0xb6, 0x4f, 0xec, 0x34, 0x0c, 0x69, 0x59, 0x4c, 0x55, 0x45, 0x0b, 0x94,
	0xa8, 0x17, 0x8e, 0xe4, 0xf6, 0x02, 0x2a, 0xa4, 0x2a, 0xff, 0x31, 0x72, 0xc0, 0xac, 0x9d, 0x20,
	0x2e, 0x27, 0xbb, 0xa7, 0xeb, 0x21, 0xf6, 0x4e, 0x98, 0x1d, 0x3b, 0xf8, 0x01, 0x78, 0x0e, 0x1e,
	0x83, 0x27, 0xe0, 0x49, 0xb8, 0xe6, 0x9a, 0x5b, 0x34, 0x7f, 0xbb, 0xb6, 0xeb, 0xba, 0xa2, 0x48,
	0xdc, 0xac, 0xe6, 0x3b, 0x7f, 0x7b, 0xe6, 0x9c, 0x6f, 0xce, 0x0c, 0x7c, 0x75, 0x27, 0xb8, 0xe4,
	0xe9, 0x41, 0xcc, 0x47, 0x34, 0x89, 0x5b, 0x21, 0x4f, 0x04, 0x8d, 0xee, 0x39, 0x8f, 0x5a, 0x09,
	0xca, 0x03, 0x7a, 0xc7, 0xd2, 0x83, 0x7b, 0xbc, 0x19, 0xf1, 0x98, 0x25, 0xd9, 0xa2, 0xa5, 0x7d,
	0x48, 0xd5, 0xe1, 0xe6, 0xb3, 0x35, 0xde, 0x74, 0x22, 0x87, 0xfa, 0x63, 0xbc, 0x9a, 0xed, 0x35,
	0xb6, 0xc3, 0x76, 0x7c, 0x27, 0xf8, 0x2f, 0xb3, 0x6c, 0x61, 0x7c, 0xfc, 0x6f, 0x60, 0xe7, 0x88,
	0xa6, 0x2c, 0x3c, 0x9c, 0xc8, 0x61, 0x80, 0x3f, 0x4f, 0x30, 0x95, 0xa4, 0x09, 0xd5, 0xab, 0x14,
	0x45, 0x42, 0xc7, 0xe8, 0x15, 0xf6, 0x0a, 0xfb, 0xb5, 0x20, 0xc3, 0x4a, 0xd7, 0xa3, 0x69, 0x7a,
	0xcf, 0x45, 0xe4, 0x15, 0x8d, 0xce, 0x61, 0xbf, 0x07, 0x75, 0x13, 0x26, 0xbd, 0xe3, 0x49, 0x8a,
	0xc4, 0x83, 0x4a, 0x27, 0xbd, 0xa6, 0x23, 0x16, 0xe9, 0x30, 0xd5, 0xc0, 0x41, 0xf2, 0x14, 0xaa,
	0xce, 0x4a, 0x47, 0xd9, 0x6a, 0x43, 0x4b, 0x6f, 0x44, 0xfd, 0x27, 0xc8, 0x74, 0xfe, 0xdf, 0x45,
	0x78, 0xf0, 0x83, 0x2d, 0x85, 0xcb, 0xee, 0x11, 0x6c, 0x5e, 0xa2, 0x1c, 0xf2, 0xc8, 0xe6, 0x66,
	0x91, 0x92, 0xf7, 0xc3, 0x21, 0x8e, 0xd1, 0xe6, 0x65, 0x11, 0x21, 0x50, 0xba, 0xe0, 0xa9, 0xf4,
	0x36, 0xb4, 0x54, 0xaf, 0x95, 0xac, 0x47, 0xe5, 0xd0, 0x2b, 0x19, 0x99, 0x5a, 0x93, 0x5d, 0x28,
	0x7f, 0x3f, 0x41, 0x31, 0xf3, 0xca, 0x5a, 0x68, 0x80, 0xb2, 0x3c, 0xe2, 0xd1, 0xcc, 0xdb, 0x34,
	0x96, 0x6a, 0x4d, 0xce, 0xa0, 0xd6, 0x9f, 0xdc, 0x8c, 0x99, 0x94, 0x18, 0x79, 0x95, 0xbd, 0x8d,
	0xfd, 0xad, 0xf6, 0x7e, 0x2b, 0xeb, 0xe0, 0x52, 0xbe, 0xad, 0xcc, 0xf4, 0x34, 0x91, 0x62, 0x16,
	0xe4, 0xae, 0x3a, 0x0b, 0x44, 0xe1, 0x55, 0x6d, 0x16, 0x88, 0x82, 0x3c, 0x83, 0xca, 0x31, 0xe7,
	0xb7, 0x0c, 0x53, 0xaf, 0xa6, 0x23, 0xef, 0xb4, 0xb2, 0x8e, 0x19, 0x45, 0xe0, 0x0c, 0xc8, 0x63,
	0xa8, 0xa9, 0x7a, 0x1d, 0xc6, 0x98, 0x48, 0x0f, 0x74, 0x90, 0x5c, 0xd0, 0xfc, 0x1a, 0xb6, 0x17,
	0x7f, 0x4d, 0x76, 0x60, 0xe3, 0x16, 0x67, 0xb6, 0x6c, 0x6a, 0xa9, 0xf6, 0x3c, 0xa5, 0xa3, 0x89,
	0x2b, 0x99, 0x01, 0x2f, 0x8b, 0x5f, 0x16, 0xfc, 0x3f, 0x0b, 0xf0, 0xc1, 0xe9, 0x98, 0xb2, 0x51,
	0x8f, 0xc6, 0x98, 0x75, 0x54, 0xd5, 0x72, 0x70, 0xd9, 0xb5, 0x21, 0xf4, 0x5a, 0x31, 0xe2, 0x1a,
	0x05, 0x7b, 0xcd, 0xd0, 0x30, 0xa2, 0x1a, 0x64, 0x98, 0x3c, 0x81, 0x92, 0x4a, 0x48, 0xd7, 0x7e,
	0xb1, 0xc7, 0x5a, 0x4e, 0x8e, 0xa0, 0x72, 0x81, 0x34, 0x42, 0x91, 0x7a, 0xa5, 0xe5, 0x3a, 0xbe,
	0xf1, 0xf7, 0x96, 0x35, 0x35, 0x75, 0x74, 0x8e, 0xcd, 0x97, 0x50, 0x9f, 0x57, 0xfc, 0xab, 0x5d,
	0xfe, 0x56, 0x82, 0x9d, 0xbc, 0x5f, 0xf9, 0x26, 0x75, 0xcb, 0x55, 0x84, 0xba, 0x6d, 0xf9, 0x67,
	0xd0, 0x50, 0xd4, 0xc6, 0x44, 0xb2, 0x90, 0xca, 0x6c, 0xa7, 0x8b, 0xc2, 0x77, 0x6e, 0x77, 0x17,
	0xca, 0x03, 0x7e, 0x8b, 0x89, 0xe5, 0x9d, 0x01, 0xe4, 0x29, 0x6c, 0x9b, 0x8e, 0x76, 0xd9, 0x14,
	0x25, 0x1b, 0xa3, 0x66, 0x60, 0x23, 0x58, 0x92, 0x92, 0x27, 0x00, 0x01, 0x46, 0x4c, 0x60, 0x28,
	0x07, 0xdc, 0x12, 0x72, 0x4e, 0x32, 0x4f, 0x9d, 0xca, 0xbb, 0xa8, 0x73, 0x98, 0x17, 0xbe, 0xaa,
	0x6d, 0xbf, 0x58, 0x45, 0xe0, 0x75, 0x75, 0x27, 0x2f, 0xe0, 0xe1, 0x19, 0x17, 0x21, 0x9e, 0xa3,
	0x3c, 0x7c, 0x2d, 0x51, 0xb8, 0x4c, 0xbc, 0x9a, 0x2e, 0xcd, 0x6a, 0x25, 0xf1, 0xa1, 0xae, 0x78,
	0xde, 0x49, 0x4f, 0x78, 0xca, 0x92, 0x58, 0xd3, 0xb6, 0x1a, 0x2c, 0xc8, 0xd4, 0x49, 0xd6, 0xb8,
	0xe7, 0x6d, 0x99, 0x93, 0x6c, 0x90, 0x62, 0xda, 0xc5, 0x60, 0xd0, 0x3b, 0xe6, 0x11, 0x7a, 0x75,
	0x5d, 0xa2, 0x0c, 0x2b, 0xdd, 0x25, 0x1b, 0xe3, 0x60, 0x76, 0x87, 0x5e, 0xc3, 0xcc, 0x25, 0x87,
	0xff, 0x13, 0x43, 0x7e, 0x2d, 0x40, 0xb9, 0x2f, 0xa9, 0x44, 0xb2, 0x07, 0x5b, 0x03, 0xc1, 0xe2,
	0x18, 0x85, 0x1e, 0x27, 0xc6, 0x7b, 0x5e, 0x34, 0x67, 0xa1, 0x87, 0x4b, 0x71, 0xc1, 0x42, 0xcf,
	0x18, 0x1f, 0xea, 0x16, 0x9a, 0x51, 0x63, 0x66, 0xd2, 0x82, 0x6c, 0x35, 0x49, 0xfc, 0x18, 0x1a,
	0x3a, 0x8d, 0x8c, 0xa5, 0x6d, 0xd8, 0xfd, 0x91, 0x86, 0x23, 0x3e, 0x89, 0x5c, 0xbf, 0xb4, 0xde,
	0xe6, 0xb5, 0x52, 0xa7, 0x7e, 0x7f, 0x15, 0x74, 0xf5, 0xfa, 0x5b, 0x9a, 0x0d, 0xca, 0x05, 0x99,
	0xcf, 0xa0, 0x11, 0x60, 0xcc, 0x52, 0x89, 0xc2, 0x38, 0xb9, 0xf9, 0x59, 0x98, 0x9b, 0x9f, 0xbb,
	0x50, 0xd6, 0xc7, 0xd3, 0xd5, 0x4b, 0x03, 0x35, 0xef, 0x8f, 0x05, 0xea, 0xe3, 0xb1, 0xa1, 0xdb,
	0xe3, 0xa0, 0xb2, 0xbf, 0xa4, 0x31, 0x0b, 0xdd, 0x9e, 0x34, 0xf0, 0x8f, 0xf3, 0x5f, 0xf5, 0xf4,
	0xb5, 0xb7, 0x6b, 0x6b, 0x6d, 0xff, 0x65, 0x0b, 0xff, 0x18, 0x6a, 0x7d, 0x16, 0x27, 0x54, 0x4e,
	0x84, 0x49, 0xb9, 0x1e, 0xe4, 0x02, 0xbf, 0x03, 0xf9, 0xdf, 0xfb, 0x93, 0x9b, 0x9f, 0x14, 0x03,
	0x8d, 0xbb, 0x83, 0xd9, 0x81, 0x2e, 0xce, 0xcd, 0x70, 0x02, 0xa5, 0x2e, 0x4b, 0x6e, 0xdd, 0xad,
	0xa0, 0xd6, 0xfe, 0xef, 0x05, 0xd8, 0x3a, 0x0c, 0x25, 0x9b, 0x32, 0x39, 0xeb, 0xf2, 0x98, 0x6c,
	0x43, 0xb1, 0x73, 0xa2, 0x83, 0x95, 0x82, 0x62, 0xe7, 0x44, 0xe3, 0x9e, 0x8d, 0x52, 0xec, 0xf4,
	0x14, 0x4f, 0xd5, 0xb1, 0xee, 0x9c, 0xd8, 0x28, 0x16, 0xe5, 0xd5, 0x29, 0xcd, 0x57, 0x67, 0x89,
	0x3f, 0xe5, 0x37, 0xf9, 0xe3, 0x41, 0xe5, 0xbb, 0x30, 0x9c, 0x08, 0x8c, 0xf4, 0xe9, 0x6e, 0x04,
	0x0e, 0xaa, 0xa3, 0xdf, 0xe5, 0xf1, 0x25, 0xa6, 0x29, 0x8d, 0xd1, 0xab, 0x98, 0xa3, 0x9f, 0x4b,
	0xfc, 0x3f, 0x0a, 0x50, 0xb9, 0x7e, 0xfe, 0xff, 0xf2, 0xf4, 0x05, 0x3c, 0xbc, 0xa0, 0x53, 0x1c,
	0x08, 0x86, 0xd1, 0xf9, 0x88, 0xdf, 0xd0, 0x91, 0x19, 0x2e, 0x7a, 0xd7, 0xd5, 0x60, 0xb5, 0x52,
	0x35, 0xd3, 0x5e, 0x8c, 0x9d, 0x13, 0x5d, 0x83, 0x52, 0x90, 0x0b, 0xda, 0x7f, 0x6d, 0x40, 0xd5,
	0x51, 0x96, 0xbc, 0x82, 0x0f, 0x0d, 0x7f, 0x1c, 0x49, 0x4c, 0x1d, 0x3f, 0xca, 0x27, 0xd5, 0x02,
	0x51, 0x9b, 0x0f, 0x96, 0xee, 0x0e, 0x72, 0x06, 0x3b, 0x9d, 0x34, 0x7b, 0xdd, 0x98, 0x97, 0x47,
	0x33, 0x37, 0x5a, 0x7e, 0xf7, 0x34, 0x1f, 0xe5, 0xba, 0x85, 0x77, 0xcc, 0x29, 0xd4, 0xcf, 0x51,
	0x76, 0x95, 0x42, 0xdd, 0x47, 0xe4, 0xe3, 0xb7, 0x5e, 0xf6, 0xcd, 0xe6, 0xdb, 0xc7, 0x28, 0x39,
	0x82, 0x5a, 0x1f, 0xc5, 0x14, 0xf5, 0xad, 0xf9, 0x9e, 0x31, 0x2e, 0x60, 0xfb, 0x1c, 0xa5, 0xbe,
	0x5f, 0x67, 0x66, 0x93, 0x6b, 0x02, 0x7d, 0xb2, 0xe6, 0x32, 0x55, 0xd9, 0x98, 0x30, 0x57, 0xc1,
	0x7b, 0x67, 0xf3, 0x0a, 0x6a, 0x7d, 0x3a, 0x45, 0xc3, 0xbb, 0x35, 0x31, 0xe6, 0x5a, 0xb6, 0x30,
	0xc4, 0x8e, 0x3e, 0x87, 0x4f, 0x13, 0x94, 0xf3, 0x0f, 0x56, 0xfb, 0x84, 0x55, 0x6f, 0xd6, 0xcc,
	0xe9, 0x66, 0x53, 0xbf, 0x55, 0x9f, 0xff, 0x13, 0x00, 0x00, 0xff, 0xff, 0xf9, 0xc2, 0x94, 0xfb,
	0x52, 0x0b, 0x00, 0x00,
}
