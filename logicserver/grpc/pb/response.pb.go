// Code generated by protoc-gen-go.
// source: response.proto
// DO NOT EDIT!

package pb

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Response struct {
	Rid  uint64 `protobuf:"varint,1,opt,name=Rid,json=rid" json:"Rid,omitempty"`
	Code string `protobuf:"bytes,2,opt,name=Code,json=code" json:"Code,omitempty"`
	Desc string `protobuf:"bytes,3,opt,name=Desc,json=desc" json:"Desc,omitempty"`
}

func (m *Response) Reset()                    { *m = Response{} }
func (m *Response) String() string            { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()               {}
func (*Response) Descriptor() ([]byte, []int) { return fileDescriptor4, []int{0} }

func (m *Response) GetRid() uint64 {
	if m != nil {
		return m.Rid
	}
	return 0
}

func (m *Response) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Response) GetDesc() string {
	if m != nil {
		return m.Desc
	}
	return ""
}

func init() {
	proto.RegisterType((*Response)(nil), "pb.Response")
}

func init() { proto.RegisterFile("response.proto", fileDescriptor4) }

var fileDescriptor4 = []byte{
	// 108 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2b, 0x4a, 0x2d, 0x2e,
	0xc8, 0xcf, 0x2b, 0x4e, 0xd5, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x2a, 0x48, 0x52, 0x72,
	0xe1, 0xe2, 0x08, 0x82, 0x8a, 0x0a, 0x09, 0x70, 0x31, 0x07, 0x65, 0xa6, 0x48, 0x30, 0x2a, 0x30,
	0x6a, 0xb0, 0x04, 0x31, 0x17, 0x65, 0xa6, 0x08, 0x09, 0x71, 0xb1, 0x38, 0xe7, 0xa7, 0xa4, 0x4a,
	0x30, 0x29, 0x30, 0x6a, 0x70, 0x06, 0xb1, 0x24, 0xe7, 0xa7, 0xa4, 0x82, 0xc4, 0x5c, 0x52, 0x8b,
	0x93, 0x25, 0x98, 0x21, 0x62, 0x29, 0xa9, 0xc5, 0xc9, 0x49, 0x6c, 0x60, 0x03, 0x8d, 0x01, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x33, 0x28, 0x65, 0xc4, 0x62, 0x00, 0x00, 0x00,
}