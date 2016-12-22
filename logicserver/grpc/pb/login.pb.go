// Code generated by protoc-gen-go.
// source: login.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	login.proto
	message.proto
	registe.proto
	request.proto
	response.proto
	rpc.proto
	session.proto

It has these top-level messages:
	DeviceLoginRequest
	DeviceLoginResponse
	DeviceOfflineRequest
	CreateMessageRequest
	CreateMessageResponse
	DeviceRegisteRequest
	DeviceRegisteResponse
	Request
	Response
	RpcInfo
	RpcRequest
	RpcResponse
	CreateSessionRequest
	CreateSessionResponse
	DeleteSessionUsersRequest
	AddSessionUsersRequest
*/
package pb

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

type DeviceLoginRequest struct {
	Rid      uint64 `protobuf:"varint,1,opt,name=Rid" json:"Rid,omitempty"`
	Token    string `protobuf:"bytes,11,opt,name=Token" json:"Token,omitempty"`
	AppId    string `protobuf:"bytes,12,opt,name=AppId" json:"AppId,omitempty"`
	DeviceId string `protobuf:"bytes,13,opt,name=DeviceId" json:"DeviceId,omitempty"`
	Platform string `protobuf:"bytes,14,opt,name=Platform" json:"Platform,omitempty"`
	UserId   string `protobuf:"bytes,15,opt,name=UserId" json:"UserId,omitempty"`
}

func (m *DeviceLoginRequest) Reset()                    { *m = DeviceLoginRequest{} }
func (m *DeviceLoginRequest) String() string            { return proto.CompactTextString(m) }
func (*DeviceLoginRequest) ProtoMessage()               {}
func (*DeviceLoginRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *DeviceLoginRequest) GetRid() uint64 {
	if m != nil {
		return m.Rid
	}
	return 0
}

func (m *DeviceLoginRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *DeviceLoginRequest) GetAppId() string {
	if m != nil {
		return m.AppId
	}
	return ""
}

func (m *DeviceLoginRequest) GetDeviceId() string {
	if m != nil {
		return m.DeviceId
	}
	return ""
}

func (m *DeviceLoginRequest) GetPlatform() string {
	if m != nil {
		return m.Platform
	}
	return ""
}

func (m *DeviceLoginRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

type DeviceLoginResponse struct {
	Rid  uint64 `protobuf:"varint,1,opt,name=rid" json:"rid,omitempty"`
	Code string `protobuf:"bytes,2,opt,name=code" json:"code,omitempty"`
	Desc string `protobuf:"bytes,3,opt,name=desc" json:"desc,omitempty"`
}

func (m *DeviceLoginResponse) Reset()                    { *m = DeviceLoginResponse{} }
func (m *DeviceLoginResponse) String() string            { return proto.CompactTextString(m) }
func (*DeviceLoginResponse) ProtoMessage()               {}
func (*DeviceLoginResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *DeviceLoginResponse) GetRid() uint64 {
	if m != nil {
		return m.Rid
	}
	return 0
}

func (m *DeviceLoginResponse) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *DeviceLoginResponse) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

type DeviceOfflineRequest struct {
	Token  string `protobuf:"bytes,1,opt,name=Token" json:"Token,omitempty"`
	UserId string `protobuf:"bytes,2,opt,name=UserId" json:"UserId,omitempty"`
}

func (m *DeviceOfflineRequest) Reset()                    { *m = DeviceOfflineRequest{} }
func (m *DeviceOfflineRequest) String() string            { return proto.CompactTextString(m) }
func (*DeviceOfflineRequest) ProtoMessage()               {}
func (*DeviceOfflineRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *DeviceOfflineRequest) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *DeviceOfflineRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func init() {
	proto.RegisterType((*DeviceLoginRequest)(nil), "pb.DeviceLoginRequest")
	proto.RegisterType((*DeviceLoginResponse)(nil), "pb.DeviceLoginResponse")
	proto.RegisterType((*DeviceOfflineRequest)(nil), "pb.DeviceOfflineRequest")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Login service

type LoginClient interface {
	// Sends a greeting
	Login(ctx context.Context, in *DeviceLoginRequest, opts ...grpc.CallOption) (*DeviceLoginResponse, error)
	Offline(ctx context.Context, in *DeviceOfflineRequest, opts ...grpc.CallOption) (*Response, error)
}

type loginClient struct {
	cc *grpc.ClientConn
}

func NewLoginClient(cc *grpc.ClientConn) LoginClient {
	return &loginClient{cc}
}

func (c *loginClient) Login(ctx context.Context, in *DeviceLoginRequest, opts ...grpc.CallOption) (*DeviceLoginResponse, error) {
	out := new(DeviceLoginResponse)
	err := grpc.Invoke(ctx, "/pb.Login/Login", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *loginClient) Offline(ctx context.Context, in *DeviceOfflineRequest, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := grpc.Invoke(ctx, "/pb.Login/Offline", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Login service

type LoginServer interface {
	// Sends a greeting
	Login(context.Context, *DeviceLoginRequest) (*DeviceLoginResponse, error)
	Offline(context.Context, *DeviceOfflineRequest) (*Response, error)
}

func RegisterLoginServer(s *grpc.Server, srv LoginServer) {
	s.RegisterService(&_Login_serviceDesc, srv)
}

func _Login_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeviceLoginRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Login/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServer).Login(ctx, req.(*DeviceLoginRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Login_Offline_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeviceOfflineRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(LoginServer).Offline(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Login/Offline",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(LoginServer).Offline(ctx, req.(*DeviceOfflineRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Login_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Login",
	HandlerType: (*LoginServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _Login_Login_Handler,
		},
		{
			MethodName: "Offline",
			Handler:    _Login_Offline_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "login.proto",
}

func init() { proto.RegisterFile("login.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 274 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x64, 0x91, 0xbb, 0x4e, 0xc3, 0x30,
	0x14, 0x86, 0x71, 0x7a, 0x01, 0x4e, 0x4b, 0x41, 0x87, 0xaa, 0x58, 0x99, 0xaa, 0x4c, 0x9d, 0x32,
	0xd0, 0x8d, 0x0d, 0xa9, 0x4b, 0x24, 0xa4, 0xa2, 0x08, 0x1e, 0x80, 0xc4, 0x27, 0xc8, 0x22, 0xc4,
	0xc6, 0x0e, 0x88, 0xe7, 0xe1, 0x49, 0x91, 0xed, 0xa6, 0x69, 0xd5, 0xed, 0xbf, 0xd8, 0xd6, 0xf7,
	0x27, 0x30, 0xa9, 0xd5, 0xbb, 0x6c, 0x52, 0x6d, 0x54, 0xab, 0x30, 0xd2, 0x45, 0x3c, 0x33, 0x64,
	0xb5, 0x6a, 0x2c, 0x85, 0x2c, 0xf9, 0x63, 0x80, 0x1b, 0xfa, 0x91, 0x25, 0x3d, 0xb9, 0x93, 0x39,
	0x7d, 0x7d, 0x93, 0x6d, 0xf1, 0x06, 0x06, 0xb9, 0x14, 0x9c, 0x2d, 0xd9, 0x6a, 0x98, 0x3b, 0x89,
	0x73, 0x18, 0xbd, 0xa8, 0x0f, 0x6a, 0xf8, 0x64, 0xc9, 0x56, 0x97, 0x79, 0x30, 0x2e, 0x7d, 0xd4,
	0x3a, 0x13, 0x7c, 0x1a, 0x52, 0x6f, 0x30, 0x86, 0x8b, 0xf0, 0x66, 0x26, 0xf8, 0x95, 0x2f, 0xf6,
	0xde, 0x75, 0xcf, 0xf5, 0x5b, 0x5b, 0x29, 0xf3, 0xc9, 0x67, 0xa1, 0xeb, 0x3c, 0x2e, 0x60, 0xfc,
	0x6a, 0xc9, 0x64, 0x82, 0x5f, 0xfb, 0x66, 0xe7, 0x92, 0x2d, 0xdc, 0x1e, 0x31, 0x86, 0x05, 0x0e,
	0xd2, 0xf4, 0x90, 0x46, 0x0a, 0x44, 0x18, 0x96, 0x4a, 0x10, 0x8f, 0xfc, 0x75, 0xaf, 0x5d, 0x26,
	0xc8, 0x96, 0x7c, 0x10, 0x32, 0xa7, 0x93, 0x0d, 0xcc, 0xc3, 0x83, 0xdb, 0xaa, 0xaa, 0x65, 0x43,
	0xdd, 0xec, 0xfd, 0x48, 0x76, 0x38, 0xb2, 0xc7, 0x8a, 0x0e, 0xb1, 0xee, 0x7f, 0x61, 0xe4, 0x81,
	0xf0, 0xa1, 0x13, 0x8b, 0x54, 0x17, 0xe9, 0xe9, 0xe7, 0x8c, 0xef, 0x4e, 0xf2, 0x30, 0x21, 0x39,
	0xc3, 0x35, 0x9c, 0xef, 0x20, 0x90, 0xf7, 0xa7, 0x8e, 0xb9, 0xe2, 0xa9, 0x6b, 0xfa, 0x4b, 0xc5,
	0xd8, 0xff, 0xbc, 0xf5, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x3a, 0x8d, 0x1b, 0x84, 0xdf, 0x01,
	0x00, 0x00,
}
