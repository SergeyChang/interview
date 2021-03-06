// Code generated by protoc-gen-go. DO NOT EDIT.
// source: app.proto

/*
Package rpc is a generated protocol buffer package.

It is generated from these files:
	app.proto

It has these top-level messages:
	GetByTokenReq
	GetByTokenResp
	Account
	StoreReq
	StoreResp
	GetReq
	GetResp
*/
package rpc

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

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

type GetByTokenReq struct {
	Token string `protobuf:"bytes,1,opt,name=token" json:"token,omitempty"`
}

func (m *GetByTokenReq) Reset()                    { *m = GetByTokenReq{} }
func (m *GetByTokenReq) String() string            { return proto.CompactTextString(m) }
func (*GetByTokenReq) ProtoMessage()               {}
func (*GetByTokenReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *GetByTokenReq) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

type GetByTokenResp struct {
	Account *Account `protobuf:"bytes,1,opt,name=account" json:"account,omitempty"`
}

func (m *GetByTokenResp) Reset()                    { *m = GetByTokenResp{} }
func (m *GetByTokenResp) String() string            { return proto.CompactTextString(m) }
func (*GetByTokenResp) ProtoMessage()               {}
func (*GetByTokenResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetByTokenResp) GetAccount() *Account {
	if m != nil {
		return m.Account
	}
	return nil
}

type Account struct {
	MaxCacheKeys int64 `protobuf:"varint,1,opt,name=max_cache_keys,json=maxCacheKeys" json:"max_cache_keys,omitempty"`
}

func (m *Account) Reset()                    { *m = Account{} }
func (m *Account) String() string            { return proto.CompactTextString(m) }
func (*Account) ProtoMessage()               {}
func (*Account) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Account) GetMaxCacheKeys() int64 {
	if m != nil {
		return m.MaxCacheKeys
	}
	return 0
}

type StoreReq struct {
	Key          string `protobuf:"bytes,1,opt,name=Key" json:"Key,omitempty"`
	Val          []byte `protobuf:"bytes,2,opt,name=Val,proto3" json:"Val,omitempty"`
	AccountToken string `protobuf:"bytes,3,opt,name=account_token,json=accountToken" json:"account_token,omitempty"`
}

func (m *StoreReq) Reset()                    { *m = StoreReq{} }
func (m *StoreReq) String() string            { return proto.CompactTextString(m) }
func (*StoreReq) ProtoMessage()               {}
func (*StoreReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *StoreReq) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

func (m *StoreReq) GetVal() []byte {
	if m != nil {
		return m.Val
	}
	return nil
}

func (m *StoreReq) GetAccountToken() string {
	if m != nil {
		return m.AccountToken
	}
	return ""
}

type StoreResp struct {
}

func (m *StoreResp) Reset()                    { *m = StoreResp{} }
func (m *StoreResp) String() string            { return proto.CompactTextString(m) }
func (*StoreResp) ProtoMessage()               {}
func (*StoreResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type GetReq struct {
	Key string `protobuf:"bytes,1,opt,name=key" json:"key,omitempty"`
}

func (m *GetReq) Reset()                    { *m = GetReq{} }
func (m *GetReq) String() string            { return proto.CompactTextString(m) }
func (*GetReq) ProtoMessage()               {}
func (*GetReq) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *GetReq) GetKey() string {
	if m != nil {
		return m.Key
	}
	return ""
}

type GetResp struct {
	Val []byte `protobuf:"bytes,1,opt,name=val,proto3" json:"val,omitempty"`
}

func (m *GetResp) Reset()                    { *m = GetResp{} }
func (m *GetResp) String() string            { return proto.CompactTextString(m) }
func (*GetResp) ProtoMessage()               {}
func (*GetResp) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *GetResp) GetVal() []byte {
	if m != nil {
		return m.Val
	}
	return nil
}

func init() {
	proto.RegisterType((*GetByTokenReq)(nil), "rpc.GetByTokenReq")
	proto.RegisterType((*GetByTokenResp)(nil), "rpc.GetByTokenResp")
	proto.RegisterType((*Account)(nil), "rpc.Account")
	proto.RegisterType((*StoreReq)(nil), "rpc.StoreReq")
	proto.RegisterType((*StoreResp)(nil), "rpc.StoreResp")
	proto.RegisterType((*GetReq)(nil), "rpc.GetReq")
	proto.RegisterType((*GetResp)(nil), "rpc.GetResp")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Cache service

type CacheClient interface {
	Store(ctx context.Context, in *StoreReq, opts ...grpc.CallOption) (*StoreResp, error)
	Get(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*GetResp, error)
}

type cacheClient struct {
	cc *grpc.ClientConn
}

func NewCacheClient(cc *grpc.ClientConn) CacheClient {
	return &cacheClient{cc}
}

func (c *cacheClient) Store(ctx context.Context, in *StoreReq, opts ...grpc.CallOption) (*StoreResp, error) {
	out := new(StoreResp)
	err := grpc.Invoke(ctx, "/rpc.Cache/Store", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *cacheClient) Get(ctx context.Context, in *GetReq, opts ...grpc.CallOption) (*GetResp, error) {
	out := new(GetResp)
	err := grpc.Invoke(ctx, "/rpc.Cache/Get", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Cache service

type CacheServer interface {
	Store(context.Context, *StoreReq) (*StoreResp, error)
	Get(context.Context, *GetReq) (*GetResp, error)
}

func RegisterCacheServer(s *grpc.Server, srv CacheServer) {
	s.RegisterService(&_Cache_serviceDesc, srv)
}

func _Cache_Store_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServer).Store(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Cache/Store",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServer).Store(ctx, req.(*StoreReq))
	}
	return interceptor(ctx, in, info, handler)
}

func _Cache_Get_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServer).Get(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Cache/Get",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServer).Get(ctx, req.(*GetReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Cache_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.Cache",
	HandlerType: (*CacheServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Store",
			Handler:    _Cache_Store_Handler,
		},
		{
			MethodName: "Get",
			Handler:    _Cache_Get_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app.proto",
}

// Client API for Accounts service

type AccountsClient interface {
	GetByToken(ctx context.Context, in *GetByTokenReq, opts ...grpc.CallOption) (*GetByTokenResp, error)
}

type accountsClient struct {
	cc *grpc.ClientConn
}

func NewAccountsClient(cc *grpc.ClientConn) AccountsClient {
	return &accountsClient{cc}
}

func (c *accountsClient) GetByToken(ctx context.Context, in *GetByTokenReq, opts ...grpc.CallOption) (*GetByTokenResp, error) {
	out := new(GetByTokenResp)
	err := grpc.Invoke(ctx, "/rpc.Accounts/GetByToken", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Accounts service

type AccountsServer interface {
	GetByToken(context.Context, *GetByTokenReq) (*GetByTokenResp, error)
}

func RegisterAccountsServer(s *grpc.Server, srv AccountsServer) {
	s.RegisterService(&_Accounts_serviceDesc, srv)
}

func _Accounts_GetByToken_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetByTokenReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(AccountsServer).GetByToken(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/rpc.Accounts/GetByToken",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(AccountsServer).GetByToken(ctx, req.(*GetByTokenReq))
	}
	return interceptor(ctx, in, info, handler)
}

var _Accounts_serviceDesc = grpc.ServiceDesc{
	ServiceName: "rpc.Accounts",
	HandlerType: (*AccountsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetByToken",
			Handler:    _Accounts_GetByToken_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "app.proto",
}

func init() { proto.RegisterFile("app.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 297 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x64, 0x51, 0x3d, 0x4f, 0xc3, 0x30,
	0x14, 0xa4, 0x44, 0xfd, 0xc8, 0x8b, 0x13, 0xa1, 0x07, 0x43, 0x15, 0x96, 0xca, 0x7c, 0x28, 0x53,
	0x90, 0xc2, 0x00, 0x2b, 0x74, 0xc8, 0xd0, 0xcd, 0x05, 0xd6, 0xc8, 0x58, 0x96, 0x90, 0xd2, 0x36,
	0x26, 0x36, 0xa8, 0xf9, 0xf7, 0xc8, 0x8e, 0xa3, 0x54, 0x62, 0xbb, 0x77, 0x77, 0x39, 0x5f, 0xde,
	0x83, 0x90, 0x2b, 0x95, 0xab, 0xb6, 0x31, 0x0d, 0x06, 0xad, 0x12, 0xf4, 0x0e, 0xe2, 0x52, 0x9a,
	0xd7, 0xee, 0xad, 0xa9, 0xe5, 0x81, 0xc9, 0x6f, 0xbc, 0x82, 0xa9, 0xb1, 0x78, 0x39, 0x59, 0x4d,
	0xb2, 0x90, 0xf5, 0x03, 0x7d, 0x86, 0xe4, 0xd4, 0xa6, 0x15, 0xde, 0xc3, 0x9c, 0x0b, 0xd1, 0xfc,
	0x1c, 0x8c, 0x73, 0x46, 0x05, 0xc9, 0x5b, 0x25, 0xf2, 0x97, 0x9e, 0x63, 0x83, 0x48, 0x1f, 0x60,
	0xee, 0x39, 0xbc, 0x85, 0x64, 0xcf, 0x8f, 0x95, 0xe0, 0xe2, 0x4b, 0x56, 0xb5, 0xec, 0xb4, 0xfb,
	0x32, 0x60, 0x64, 0xcf, 0x8f, 0x6b, 0x4b, 0x6e, 0x64, 0xa7, 0xe9, 0x16, 0x16, 0x5b, 0xd3, 0xb4,
	0xd2, 0x96, 0xb9, 0x80, 0x60, 0x23, 0x3b, 0x5f, 0xc5, 0x42, 0xcb, 0x7c, 0xf0, 0xdd, 0xf2, 0x7c,
	0x35, 0xc9, 0x08, 0xb3, 0x10, 0x6f, 0x20, 0xf6, 0x6f, 0x55, 0x7d, 0xf1, 0xc0, 0xb9, 0x89, 0x27,
	0x5d, 0x63, 0x1a, 0x41, 0xe8, 0x43, 0xb5, 0xa2, 0x29, 0xcc, 0x4a, 0x69, 0x7c, 0x7e, 0x3d, 0xe6,
	0xd7, 0xb2, 0xa3, 0xd7, 0x30, 0x77, 0x9a, 0x56, 0x56, 0xfc, 0xe5, 0x3b, 0x27, 0x12, 0x66, 0x61,
	0xf1, 0x0e, 0x53, 0xd7, 0x13, 0x33, 0x98, 0xba, 0x38, 0x8c, 0xdd, 0x4f, 0x0f, 0x7d, 0xd3, 0xe4,
	0x74, 0xd4, 0x8a, 0x9e, 0x21, 0x85, 0xa0, 0x94, 0x06, 0x23, 0x27, 0xf4, 0xaf, 0xa6, 0x64, 0x1c,
	0xac, 0xa7, 0x58, 0xc3, 0xc2, 0xaf, 0x48, 0xe3, 0x13, 0xc0, 0xb8, 0x68, 0xc4, 0xc1, 0x39, 0x1e,
	0x28, 0xbd, 0xfc, 0xc7, 0xd9, 0x90, 0xcf, 0x99, 0x3b, 0xea, 0xe3, 0x5f, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x24, 0xeb, 0x65, 0x53, 0xe1, 0x01, 0x00, 0x00,
}
