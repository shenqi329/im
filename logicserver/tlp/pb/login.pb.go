// Code generated by protoc-gen-go.
// source: login.proto
// DO NOT EDIT!

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type DeviceLoginRequest struct {
	Rid      uint64 `protobuf:"varint,1,opt,name=Rid,json=rid" json:"Rid,omitempty"`
	Token    string `protobuf:"bytes,11,opt,name=Token,json=token" json:"Token,omitempty"`
	AppId    string `protobuf:"bytes,12,opt,name=AppId,json=appId" json:"AppId,omitempty"`
	DeviceId string `protobuf:"bytes,13,opt,name=DeviceId,json=deviceId" json:"DeviceId,omitempty"`
	Platform string `protobuf:"bytes,14,opt,name=Platform,json=platform" json:"Platform,omitempty"`
	UserId   string `protobuf:"bytes,15,opt,name=UserId,json=userId" json:"UserId,omitempty"`
}

func (m *DeviceLoginRequest) Reset()                    { *m = DeviceLoginRequest{} }
func (m *DeviceLoginRequest) String() string            { return proto.CompactTextString(m) }
func (*DeviceLoginRequest) ProtoMessage()               {}
func (*DeviceLoginRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

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
func (*DeviceLoginResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

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
	Token  string `protobuf:"bytes,1,opt,name=Token,json=token" json:"Token,omitempty"`
	UserId string `protobuf:"bytes,2,opt,name=UserId,json=userId" json:"UserId,omitempty"`
}

func (m *DeviceOfflineRequest) Reset()                    { *m = DeviceOfflineRequest{} }
func (m *DeviceOfflineRequest) String() string            { return proto.CompactTextString(m) }
func (*DeviceOfflineRequest) ProtoMessage()               {}
func (*DeviceOfflineRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

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

func init() { proto.RegisterFile("login.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 233 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x54, 0x90, 0xcf, 0x4a, 0x03, 0x31,
	0x10, 0xc6, 0xc9, 0x6e, 0xbb, 0xd4, 0xa9, 0xff, 0x88, 0x45, 0x06, 0x4f, 0x65, 0x4f, 0x3d, 0x79,
	0xf1, 0x09, 0x84, 0x5e, 0x16, 0x84, 0xca, 0xa2, 0x0f, 0xd0, 0xee, 0xcc, 0x4a, 0x70, 0xdd, 0xc4,
	0x24, 0xf5, 0x81, 0x7c, 0x52, 0x99, 0xa4, 0x6a, 0x7b, 0x9b, 0xef, 0xfb, 0x91, 0xe1, 0x37, 0x81,
	0xf9, 0x60, 0xdf, 0xcc, 0x78, 0xef, 0xbc, 0x8d, 0x56, 0x17, 0x6e, 0x57, 0x7f, 0x2b, 0xd0, 0x6b,
	0xfe, 0x32, 0x1d, 0x3f, 0x09, 0x69, 0xf9, 0x73, 0xcf, 0x21, 0xea, 0x6b, 0x28, 0x5b, 0x43, 0xa8,
	0x96, 0x6a, 0x35, 0x69, 0x4b, 0x6f, 0x48, 0x2f, 0x60, 0xfa, 0x62, 0xdf, 0x79, 0xc4, 0xf9, 0x52,
	0xad, 0xce, 0xda, 0x69, 0x94, 0x20, 0xed, 0xa3, 0x73, 0x0d, 0xe1, 0x79, 0x6e, 0xb7, 0x12, 0xf4,
	0x1d, 0xcc, 0xf2, 0xce, 0x86, 0xf0, 0x22, 0x81, 0x19, 0x1d, 0xb2, 0xb0, 0xe7, 0x61, 0x1b, 0x7b,
	0xeb, 0x3f, 0xf0, 0x32, 0x33, 0x77, 0xc8, 0xfa, 0x16, 0xaa, 0xd7, 0xc0, 0xbe, 0x21, 0xbc, 0x4a,
	0xa4, 0xda, 0xa7, 0x54, 0x6f, 0xe0, 0xe6, 0xc4, 0x31, 0x38, 0x3b, 0x06, 0x16, 0x49, 0x7f, 0x2a,
	0xa9, 0x61, 0xd2, 0x59, 0x62, 0x2c, 0xd2, 0xf3, 0x34, 0x4b, 0x47, 0x1c, 0x3a, 0x2c, 0x73, 0x27,
	0x73, 0xbd, 0x86, 0x45, 0x5e, 0xb8, 0xe9, 0xfb, 0xc1, 0x8c, 0xfc, 0x7b, 0xf6, 0xdf, 0x91, 0xea,
	0xf8, 0xc8, 0x7f, 0xad, 0xe2, 0x58, 0x6b, 0x57, 0xa5, 0x6f, 0x7c, 0xf8, 0x09, 0x00, 0x00, 0xff,
	0xff, 0x53, 0x08, 0x05, 0xfd, 0x55, 0x01, 0x00, 0x00,
}
