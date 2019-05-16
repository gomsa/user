// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/auth/auth.proto

package auth

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

import (
	client "github.com/micro/go-micro/client"
	server "github.com/micro/go-micro/server"
	context "golang.org/x/net/context"
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

type User struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
	Mobile               string   `protobuf:"bytes,3,opt,name=mobile,proto3" json:"mobile,omitempty"`
	Email                string   `protobuf:"bytes,4,opt,name=email,proto3" json:"email,omitempty"`
	Password             string   `protobuf:"bytes,5,opt,name=password,proto3" json:"password,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *User) Reset()         { *m = User{} }
func (m *User) String() string { return proto.CompactTextString(m) }
func (*User) ProtoMessage()    {}
func (*User) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{0}
}

func (m *User) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_User.Unmarshal(m, b)
}
func (m *User) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_User.Marshal(b, m, deterministic)
}
func (m *User) XXX_Merge(src proto.Message) {
	xxx_messageInfo_User.Merge(m, src)
}
func (m *User) XXX_Size() int {
	return xxx_messageInfo_User.Size(m)
}
func (m *User) XXX_DiscardUnknown() {
	xxx_messageInfo_User.DiscardUnknown(m)
}

var xxx_messageInfo_User proto.InternalMessageInfo

func (m *User) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *User) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func (m *User) GetMobile() string {
	if m != nil {
		return m.Mobile
	}
	return ""
}

func (m *User) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *User) GetPassword() string {
	if m != nil {
		return m.Password
	}
	return ""
}

type Request struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Service              string   `protobuf:"bytes,2,opt,name=service,proto3" json:"service,omitempty"`
	Method               string   `protobuf:"bytes,3,opt,name=method,proto3" json:"method,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{1}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Request) GetService() string {
	if m != nil {
		return m.Service
	}
	return ""
}

func (m *Request) GetMethod() string {
	if m != nil {
		return m.Method
	}
	return ""
}

type Token struct {
	Token                string   `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Valid                bool     `protobuf:"varint,2,opt,name=valid,proto3" json:"valid,omitempty"`
	Errors               []*Error `protobuf:"bytes,3,rep,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Token) Reset()         { *m = Token{} }
func (m *Token) String() string { return proto.CompactTextString(m) }
func (*Token) ProtoMessage()    {}
func (*Token) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{2}
}

func (m *Token) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Token.Unmarshal(m, b)
}
func (m *Token) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Token.Marshal(b, m, deterministic)
}
func (m *Token) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Token.Merge(m, src)
}
func (m *Token) XXX_Size() int {
	return xxx_messageInfo_Token.Size(m)
}
func (m *Token) XXX_DiscardUnknown() {
	xxx_messageInfo_Token.DiscardUnknown(m)
}

var xxx_messageInfo_Token proto.InternalMessageInfo

func (m *Token) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Token) GetValid() bool {
	if m != nil {
		return m.Valid
	}
	return false
}

func (m *Token) GetErrors() []*Error {
	if m != nil {
		return m.Errors
	}
	return nil
}

type Error struct {
	Code                 string   `protobuf:"bytes,1,opt,name=code,proto3" json:"code,omitempty"`
	Description          string   `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_82b5829f48cfb8e5, []int{3}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Error) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func init() {
	proto.RegisterType((*User)(nil), "auth.User")
	proto.RegisterType((*Request)(nil), "auth.Request")
	proto.RegisterType((*Token)(nil), "auth.Token")
	proto.RegisterType((*Error)(nil), "auth.Error")
}

func init() { proto.RegisterFile("proto/auth/auth.proto", fileDescriptor_82b5829f48cfb8e5) }

var fileDescriptor_82b5829f48cfb8e5 = []byte{
	// 307 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0x41, 0x4b, 0x03, 0x31,
	0x10, 0x85, 0x6d, 0xbb, 0xdb, 0xd6, 0x29, 0xf5, 0x30, 0x54, 0x09, 0x3d, 0xd5, 0x15, 0xa1, 0x17,
	0x2b, 0xd4, 0xb3, 0x07, 0x05, 0x0f, 0x1e, 0x5d, 0x54, 0xbc, 0xa6, 0xcd, 0x40, 0x83, 0x6d, 0xb3,
	0x26, 0xd9, 0x8a, 0x20, 0xf8, 0xd7, 0x25, 0x93, 0x6c, 0x11, 0x11, 0x2f, 0xbb, 0xf9, 0xde, 0x4c,
	0x5e, 0x5e, 0x32, 0x70, 0x5c, 0x59, 0xe3, 0xcd, 0xa5, 0xac, 0xfd, 0x8a, 0x3f, 0x33, 0x66, 0xcc,
	0xc2, 0xba, 0xf8, 0x84, 0xec, 0xc9, 0x91, 0xc5, 0x23, 0x68, 0x6b, 0x25, 0x5a, 0x93, 0xd6, 0xf4,
	0xb0, 0x6c, 0x6b, 0x85, 0x63, 0xe8, 0xd7, 0x8e, 0xec, 0x56, 0x6e, 0x48, 0xb4, 0x59, 0xdd, 0x33,
	0x9e, 0x40, 0x77, 0x63, 0x16, 0x7a, 0x4d, 0xa2, 0xc3, 0x95, 0x44, 0x38, 0x82, 0x9c, 0x36, 0x52,
	0xaf, 0x45, 0xc6, 0x72, 0x84, 0xe0, 0x54, 0x49, 0xe7, 0xde, 0x8d, 0x55, 0x22, 0x8f, 0x4e, 0x0d,
	0x17, 0x0f, 0xd0, 0x2b, 0xe9, 0xad, 0x26, 0xe7, 0xc3, 0x66, 0x6f, 0x5e, 0x69, 0x9b, 0x32, 0x44,
	0x40, 0x01, 0x3d, 0x47, 0x76, 0xa7, 0x97, 0x4d, 0x8a, 0x06, 0x39, 0x04, 0xf9, 0x95, 0x51, 0xfb,
	0x10, 0x4c, 0xc5, 0x0b, 0xe4, 0x8f, 0xbc, 0xf5, 0x6f, 0xc3, 0x11, 0xe4, 0x3b, 0xb9, 0xd6, 0x8a,
	0xed, 0xfa, 0x65, 0x04, 0x3c, 0x83, 0x2e, 0x59, 0x6b, 0xac, 0x13, 0x9d, 0x49, 0x67, 0x3a, 0x98,
	0x0f, 0x66, 0xfc, 0x50, 0x77, 0x41, 0x2b, 0x53, 0xa9, 0xb8, 0x86, 0x9c, 0x05, 0x44, 0xc8, 0x96,
	0x46, 0x51, 0x32, 0xe6, 0x35, 0x4e, 0x60, 0xa0, 0xc8, 0x2d, 0xad, 0xae, 0xbc, 0x36, 0xdb, 0x14,
	0xf6, 0xa7, 0x34, 0xff, 0x82, 0xec, 0xa6, 0xf6, 0x2b, 0x3c, 0x87, 0x7e, 0xf8, 0xdf, 0x7e, 0xdc,
	0x2b, 0x84, 0x78, 0x4e, 0x98, 0xc0, 0x38, 0x9d, 0xc9, 0xe1, 0x8b, 0x03, 0x3c, 0x4d, 0xed, 0xff,
	0xb4, 0x5c, 0xc0, 0xf0, 0x39, 0xc4, 0x97, 0x9e, 0xe2, 0x95, 0x87, 0xb1, 0x9e, 0x9e, 0xf4, 0x57,
	0xfb, 0xa2, 0xcb, 0x73, 0xbf, 0xfa, 0x0e, 0x00, 0x00, 0xff, 0xff, 0x4f, 0xc1, 0x96, 0x69, 0x10,
	0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ client.Option
var _ server.Option

// Client API for Auth service

type AuthClient interface {
	// 根据用户ID获取授权
	AuthById(ctx context.Context, in *User, opts ...client.CallOption) (*Token, error)
	// 用户验证授权
	Auth(ctx context.Context, in *User, opts ...client.CallOption) (*Token, error)
	// token 验证
	ValidateToken(ctx context.Context, in *Request, opts ...client.CallOption) (*Token, error)
}

type authClient struct {
	c           client.Client
	serviceName string
}

func NewAuthClient(serviceName string, c client.Client) AuthClient {
	if c == nil {
		c = client.NewClient()
	}
	if len(serviceName) == 0 {
		serviceName = "auth"
	}
	return &authClient{
		c:           c,
		serviceName: serviceName,
	}
}

func (c *authClient) AuthById(ctx context.Context, in *User, opts ...client.CallOption) (*Token, error) {
	req := c.c.NewRequest(c.serviceName, "Auth.AuthById", in)
	out := new(Token)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) Auth(ctx context.Context, in *User, opts ...client.CallOption) (*Token, error) {
	req := c.c.NewRequest(c.serviceName, "Auth.Auth", in)
	out := new(Token)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *authClient) ValidateToken(ctx context.Context, in *Request, opts ...client.CallOption) (*Token, error) {
	req := c.c.NewRequest(c.serviceName, "Auth.ValidateToken", in)
	out := new(Token)
	err := c.c.Call(ctx, req, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Auth service

type AuthHandler interface {
	// 根据用户ID获取授权
	AuthById(context.Context, *User, *Token) error
	// 用户验证授权
	Auth(context.Context, *User, *Token) error
	// token 验证
	ValidateToken(context.Context, *Request, *Token) error
}

func RegisterAuthHandler(s server.Server, hdlr AuthHandler, opts ...server.HandlerOption) {
	s.Handle(s.NewHandler(&Auth{hdlr}, opts...))
}

type Auth struct {
	AuthHandler
}

func (h *Auth) AuthById(ctx context.Context, in *User, out *Token) error {
	return h.AuthHandler.AuthById(ctx, in, out)
}

func (h *Auth) Auth(ctx context.Context, in *User, out *Token) error {
	return h.AuthHandler.Auth(ctx, in, out)
}

func (h *Auth) ValidateToken(ctx context.Context, in *Request, out *Token) error {
	return h.AuthHandler.ValidateToken(ctx, in, out)
}
