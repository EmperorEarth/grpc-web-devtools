// Code generated by protoc-gen-go.
// source: example2.proto
// DO NOT EDIT!

package main

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type ExampleTwoRequest struct {
	Id int32 `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
}

func (m *ExampleTwoRequest) Reset()                    { *m = ExampleTwoRequest{} }
func (m *ExampleTwoRequest) String() string            { return proto.CompactTextString(m) }
func (*ExampleTwoRequest) ProtoMessage()               {}
func (*ExampleTwoRequest) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *ExampleTwoRequest) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

type ExampleTwoResponse struct {
	Id  int32  `protobuf:"varint,1,opt,name=id" json:"id,omitempty"`
	Msg string `protobuf:"bytes,2,opt,name=msg" json:"msg,omitempty"`
}

func (m *ExampleTwoResponse) Reset()                    { *m = ExampleTwoResponse{} }
func (m *ExampleTwoResponse) String() string            { return proto.CompactTextString(m) }
func (*ExampleTwoResponse) ProtoMessage()               {}
func (*ExampleTwoResponse) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *ExampleTwoResponse) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ExampleTwoResponse) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func init() {
	proto.RegisterType((*ExampleTwoRequest)(nil), "s12.example.ExampleTwoRequest")
	proto.RegisterType((*ExampleTwoResponse)(nil), "s12.example.ExampleTwoResponse")
}

func init() { proto.RegisterFile("example2.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 124 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x4b, 0xad, 0x48, 0xcc,
	0x2d, 0xc8, 0x49, 0x35, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x2e, 0x36, 0x34, 0xd2,
	0x83, 0x8a, 0x29, 0x29, 0x73, 0x09, 0xba, 0x42, 0x98, 0x21, 0xe5, 0xf9, 0x41, 0xa9, 0x85, 0xa5,
	0xa9, 0xc5, 0x25, 0x42, 0x7c, 0x5c, 0x4c, 0x99, 0x29, 0x12, 0x8c, 0x0a, 0x8c, 0x1a, 0xac, 0x41,
	0x4c, 0x99, 0x29, 0x4a, 0x66, 0x5c, 0x42, 0xc8, 0x8a, 0x8a, 0x0b, 0xf2, 0xf3, 0x8a, 0x53, 0xd1,
	0x55, 0x09, 0x09, 0x70, 0x31, 0xe7, 0x16, 0xa7, 0x4b, 0x30, 0x29, 0x30, 0x6a, 0x70, 0x06, 0x81,
	0x98, 0x4e, 0x6c, 0x51, 0x2c, 0xb9, 0x89, 0x99, 0x79, 0x49, 0x6c, 0x60, 0x8b, 0x8d, 0x01, 0x01,
	0x00, 0x00, 0xff, 0xff, 0x32, 0xa6, 0xfd, 0x8a, 0x8a, 0x00, 0x00, 0x00,
}