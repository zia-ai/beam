// Code generated by protoc-gen-go. DO NOT EDIT.
// source: v1.proto

package v1

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

type PubSubPayload_Op int32

const (
	PubSubPayload_INVALID PubSubPayload_Op = 0
	PubSubPayload_READ    PubSubPayload_Op = 1
	PubSubPayload_WRITE   PubSubPayload_Op = 2
)

var PubSubPayload_Op_name = map[int32]string{
	0: "INVALID",
	1: "READ",
	2: "WRITE",
}

var PubSubPayload_Op_value = map[string]int32{
	"INVALID": 0,
	"READ":    1,
	"WRITE":   2,
}

func (x PubSubPayload_Op) String() string {
	return proto.EnumName(PubSubPayload_Op_name, int32(x))
}

func (PubSubPayload_Op) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_2e4aa7d76fd7ee8a, []int{0, 0}
}

type PubSubPayload struct {
	Op                   PubSubPayload_Op `protobuf:"varint,1,opt,name=op,proto3,enum=org.apache.beam.io.pubsubio.v1.PubSubPayload_Op" json:"op,omitempty"`
	Topic                string           `protobuf:"bytes,2,opt,name=Topic,proto3" json:"Topic,omitempty"`
	Subscription         string           `protobuf:"bytes,3,opt,name=Subscription,proto3" json:"Subscription,omitempty"`
	IdAttribute          string           `protobuf:"bytes,4,opt,name=IdAttribute,proto3" json:"IdAttribute,omitempty"`
	TimestampAttribute   string           `protobuf:"bytes,5,opt,name=TimestampAttribute,proto3" json:"TimestampAttribute,omitempty"`
	WithAttributes       bool             `protobuf:"varint,6,opt,name=WithAttributes,proto3" json:"WithAttributes,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *PubSubPayload) Reset()         { *m = PubSubPayload{} }
func (m *PubSubPayload) String() string { return proto.CompactTextString(m) }
func (*PubSubPayload) ProtoMessage()    {}
func (*PubSubPayload) Descriptor() ([]byte, []int) {
	return fileDescriptor_2e4aa7d76fd7ee8a, []int{0}
}

func (m *PubSubPayload) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PubSubPayload.Unmarshal(m, b)
}
func (m *PubSubPayload) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PubSubPayload.Marshal(b, m, deterministic)
}
func (m *PubSubPayload) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PubSubPayload.Merge(m, src)
}
func (m *PubSubPayload) XXX_Size() int {
	return xxx_messageInfo_PubSubPayload.Size(m)
}
func (m *PubSubPayload) XXX_DiscardUnknown() {
	xxx_messageInfo_PubSubPayload.DiscardUnknown(m)
}

var xxx_messageInfo_PubSubPayload proto.InternalMessageInfo

func (m *PubSubPayload) GetOp() PubSubPayload_Op {
	if m != nil {
		return m.Op
	}
	return PubSubPayload_INVALID
}

func (m *PubSubPayload) GetTopic() string {
	if m != nil {
		return m.Topic
	}
	return ""
}

func (m *PubSubPayload) GetSubscription() string {
	if m != nil {
		return m.Subscription
	}
	return ""
}

func (m *PubSubPayload) GetIdAttribute() string {
	if m != nil {
		return m.IdAttribute
	}
	return ""
}

func (m *PubSubPayload) GetTimestampAttribute() string {
	if m != nil {
		return m.TimestampAttribute
	}
	return ""
}

func (m *PubSubPayload) GetWithAttributes() bool {
	if m != nil {
		return m.WithAttributes
	}
	return false
}

func init() {
	proto.RegisterEnum("org.apache.beam.io.pubsubio.v1.PubSubPayload_Op", PubSubPayload_Op_name, PubSubPayload_Op_value)
	proto.RegisterType((*PubSubPayload)(nil), "org.apache.beam.io.pubsubio.v1.PubSubPayload")
}

func init() {
	proto.RegisterFile("v1.proto", fileDescriptor_2e4aa7d76fd7ee8a)
}

var fileDescriptor_2e4aa7d76fd7ee8a = []byte{
	// 254 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x7c, 0x90, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0xcd, 0x9a, 0xd4, 0x74, 0xaa, 0x25, 0x0c, 0x1e, 0xf6, 0x24, 0x25, 0x87, 0xe2, 0x69,
	0x31, 0xfa, 0x02, 0x46, 0xda, 0x43, 0x40, 0x6c, 0x59, 0x83, 0x05, 0x6f, 0xbb, 0x69, 0xb0, 0x0b,
	0xc6, 0x5d, 0x92, 0xdd, 0x80, 0x0f, 0xe5, 0x3b, 0xba, 0x44, 0x50, 0x2b, 0xd2, 0xdb, 0xcc, 0xf7,
	0x7f, 0xff, 0x1c, 0x06, 0xe2, 0x3e, 0x63, 0xa6, 0xd5, 0x56, 0xe3, 0x85, 0x6e, 0x5f, 0x98, 0x30,
	0xa2, 0xda, 0xd5, 0x4c, 0xd6, 0xa2, 0x61, 0x4a, 0x33, 0xe3, 0x64, 0xe7, 0xa4, 0x1f, 0xfa, 0x2c,
	0xfd, 0x20, 0x70, 0xb6, 0x76, 0xf2, 0xd1, 0xc9, 0xb5, 0x78, 0x7f, 0xd5, 0x62, 0x8b, 0xb7, 0x40,
	0xb4, 0xa1, 0xc1, 0x2c, 0xb8, 0x9c, 0x5e, 0x5f, 0xb1, 0xc3, 0x75, 0xb6, 0x57, 0x65, 0x2b, 0xc3,
	0x7d, 0x17, 0xcf, 0x21, 0x2a, 0xb5, 0x51, 0x15, 0x25, 0xfe, 0xc8, 0x98, 0x7f, 0x2d, 0x98, 0xc2,
	0xa9, 0x57, 0xbb, 0xaa, 0x55, 0xc6, 0x2a, 0xfd, 0x46, 0x8f, 0x87, 0x70, 0x8f, 0xe1, 0x0c, 0x26,
	0xc5, 0x36, 0xb7, 0xb6, 0x55, 0xd2, 0xd9, 0x9a, 0x86, 0x83, 0xf2, 0x1b, 0x21, 0x03, 0x2c, 0x55,
	0x53, 0x77, 0x56, 0x34, 0xe6, 0x47, 0x8c, 0x06, 0xf1, 0x9f, 0x04, 0xe7, 0x30, 0xdd, 0x28, 0xbb,
	0xfb, 0x06, 0x1d, 0x1d, 0x79, 0x37, 0xe6, 0x7f, 0x68, 0x3a, 0x07, 0xb2, 0x32, 0x38, 0x81, 0x93,
	0xe2, 0xe1, 0x29, 0xbf, 0x2f, 0x16, 0xc9, 0x11, 0xc6, 0x10, 0xf2, 0x65, 0xbe, 0x48, 0x02, 0x1c,
	0x43, 0xb4, 0xe1, 0x45, 0xb9, 0x4c, 0xc8, 0x5d, 0xf8, 0x4c, 0xfa, 0x4c, 0x8e, 0x86, 0xe7, 0xde,
	0x7c, 0x06, 0x00, 0x00, 0xff, 0xff, 0xba, 0x3b, 0xdd, 0x0b, 0x68, 0x01, 0x00, 0x00,
}
