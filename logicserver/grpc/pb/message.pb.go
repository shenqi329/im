// Code generated by protoc-gen-go.
// source: message.proto
// DO NOT EDIT!

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

type CreateMessageRequest struct {
	Rid       uint64 `protobuf:"varint,1,opt,name=Rid" json:"Rid,omitempty"`
	UserId    string `protobuf:"bytes,2,opt,name=UserId" json:"UserId,omitempty"`
	SessionId uint64 `protobuf:"varint,11,opt,name=SessionId" json:"SessionId,omitempty"`
	Type      uint32 `protobuf:"varint,12,opt,name=Type" json:"Type,omitempty"`
	Content   string `protobuf:"bytes,13,opt,name=Content" json:"Content,omitempty"`
	Id        string `protobuf:"bytes,14,opt,name=Id" json:"Id,omitempty"`
}

func (m *CreateMessageRequest) Reset()                    { *m = CreateMessageRequest{} }
func (m *CreateMessageRequest) String() string            { return proto.CompactTextString(m) }
func (*CreateMessageRequest) ProtoMessage()               {}
func (*CreateMessageRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *CreateMessageRequest) GetRid() uint64 {
	if m != nil {
		return m.Rid
	}
	return 0
}

func (m *CreateMessageRequest) GetUserId() string {
	if m != nil {
		return m.UserId
	}
	return ""
}

func (m *CreateMessageRequest) GetSessionId() uint64 {
	if m != nil {
		return m.SessionId
	}
	return 0
}

func (m *CreateMessageRequest) GetType() uint32 {
	if m != nil {
		return m.Type
	}
	return 0
}

func (m *CreateMessageRequest) GetContent() string {
	if m != nil {
		return m.Content
	}
	return ""
}

func (m *CreateMessageRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type CreateMessageResponse struct {
	Rid  uint64 `protobuf:"varint,1,opt,name=rid" json:"rid,omitempty"`
	Code string `protobuf:"bytes,2,opt,name=code" json:"code,omitempty"`
	Desc string `protobuf:"bytes,3,opt,name=desc" json:"desc,omitempty"`
}

func (m *CreateMessageResponse) Reset()                    { *m = CreateMessageResponse{} }
func (m *CreateMessageResponse) String() string            { return proto.CompactTextString(m) }
func (*CreateMessageResponse) ProtoMessage()               {}
func (*CreateMessageResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *CreateMessageResponse) GetRid() uint64 {
	if m != nil {
		return m.Rid
	}
	return 0
}

func (m *CreateMessageResponse) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *CreateMessageResponse) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func init() {
	proto.RegisterType((*CreateMessageRequest)(nil), "pb.CreateMessageRequest")
	proto.RegisterType((*CreateMessageResponse)(nil), "pb.CreateMessageResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for Message service

type MessageClient interface {
	// Sends a greeting
	CreateMessage(ctx context.Context, in *CreateMessageRequest, opts ...grpc.CallOption) (*CreateMessageResponse, error)
}

type messageClient struct {
	cc *grpc.ClientConn
}

func NewMessageClient(cc *grpc.ClientConn) MessageClient {
	return &messageClient{cc}
}

func (c *messageClient) CreateMessage(ctx context.Context, in *CreateMessageRequest, opts ...grpc.CallOption) (*CreateMessageResponse, error) {
	out := new(CreateMessageResponse)
	err := grpc.Invoke(ctx, "/pb.Message/CreateMessage", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Message service

type MessageServer interface {
	// Sends a greeting
	CreateMessage(context.Context, *CreateMessageRequest) (*CreateMessageResponse, error)
}

func RegisterMessageServer(s *grpc.Server, srv MessageServer) {
	s.RegisterService(&_Message_serviceDesc, srv)
}

func _Message_CreateMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MessageServer).CreateMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pb.Message/CreateMessage",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MessageServer).CreateMessage(ctx, req.(*CreateMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Message_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pb.Message",
	HandlerType: (*MessageServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateMessage",
			Handler:    _Message_CreateMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "message.proto",
}

func init() { proto.RegisterFile("message.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 234 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x6c, 0x90, 0xc1, 0x4a, 0x03, 0x31,
	0x10, 0x86, 0xcd, 0xb6, 0xb4, 0x74, 0x74, 0x8b, 0x0c, 0x2a, 0xa3, 0x78, 0x58, 0xf6, 0xb4, 0xa7,
	0x3d, 0xe8, 0x23, 0x14, 0x84, 0x3d, 0x78, 0x68, 0xd4, 0x07, 0xe8, 0x36, 0x83, 0xec, 0xc1, 0x24,
	0x66, 0xe2, 0xc1, 0x77, 0xf1, 0x61, 0x25, 0xd9, 0x95, 0x62, 0xe9, 0xed, 0x9f, 0x6f, 0xc2, 0x9f,
	0x8f, 0x81, 0xf2, 0x83, 0x45, 0x76, 0xef, 0xdc, 0xfa, 0xe0, 0xa2, 0xc3, 0xc2, 0xf7, 0xf5, 0x8f,
	0x82, 0xab, 0x4d, 0xe0, 0x5d, 0xe4, 0xe7, 0x71, 0xa7, 0xf9, 0xf3, 0x8b, 0x25, 0xe2, 0x25, 0xcc,
	0xf4, 0x60, 0x48, 0x55, 0xaa, 0x99, 0xeb, 0x14, 0xf1, 0x06, 0x16, 0x6f, 0xc2, 0xa1, 0x33, 0x54,
	0x54, 0xaa, 0x59, 0xe9, 0x69, 0xc2, 0x7b, 0x58, 0xbd, 0xb0, 0xc8, 0xe0, 0x6c, 0x67, 0xe8, 0x3c,
	0xbf, 0x3f, 0x00, 0x44, 0x98, 0xbf, 0x7e, 0x7b, 0xa6, 0x8b, 0x4a, 0x35, 0xa5, 0xce, 0x19, 0x09,
	0x96, 0x1b, 0x67, 0x23, 0xdb, 0x48, 0x65, 0xae, 0xfa, 0x1b, 0x71, 0x0d, 0x45, 0x67, 0x68, 0x9d,
	0x61, 0xd1, 0x99, 0x7a, 0x0b, 0xd7, 0x47, 0x76, 0xe2, 0x9d, 0x15, 0x4e, 0x7a, 0xe1, 0xa0, 0x17,
	0x86, 0xfc, 0xd1, 0xde, 0x19, 0x9e, 0xe4, 0x72, 0x4e, 0xcc, 0xb0, 0xec, 0x69, 0x36, 0xb2, 0x94,
	0x1f, 0xb6, 0xb0, 0x9c, 0xca, 0xf0, 0x09, 0xca, 0x7f, 0xed, 0x48, 0xad, 0xef, 0xdb, 0x53, 0xe7,
	0xb8, 0xbb, 0x3d, 0xb1, 0x19, 0x55, 0xea, 0xb3, 0x7e, 0x91, 0xef, 0xf9, 0xf8, 0x1b, 0x00, 0x00,
	0xff, 0xff, 0xa1, 0x53, 0x7a, 0x82, 0x60, 0x01, 0x00, 0x00,
}
