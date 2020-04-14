// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mail/indicator.proto

package mail

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

type Indicator struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Value                string   `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
	Description          string   `protobuf:"bytes,4,opt,name=description,proto3" json:"description,omitempty"`
	CreatedDate          int64    `protobuf:"varint,5,opt,name=created_date,json=createdDate,proto3" json:"created_date,omitempty"`
	Score                float64  `protobuf:"fixed64,6,opt,name=score,proto3" json:"score,omitempty"`
	EntityId             string   `protobuf:"bytes,7,opt,name=entity_id,json=entityId,proto3" json:"entity_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Indicator) Reset()         { *m = Indicator{} }
func (m *Indicator) String() string { return proto.CompactTextString(m) }
func (*Indicator) ProtoMessage()    {}
func (*Indicator) Descriptor() ([]byte, []int) {
	return fileDescriptor_97fd999573541b11, []int{0}
}

func (m *Indicator) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Indicator.Unmarshal(m, b)
}
func (m *Indicator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Indicator.Marshal(b, m, deterministic)
}
func (m *Indicator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Indicator.Merge(m, src)
}
func (m *Indicator) XXX_Size() int {
	return xxx_messageInfo_Indicator.Size(m)
}
func (m *Indicator) XXX_DiscardUnknown() {
	xxx_messageInfo_Indicator.DiscardUnknown(m)
}

var xxx_messageInfo_Indicator proto.InternalMessageInfo

func (m *Indicator) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Indicator) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Indicator) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *Indicator) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Indicator) GetCreatedDate() int64 {
	if m != nil {
		return m.CreatedDate
	}
	return 0
}

func (m *Indicator) GetScore() float64 {
	if m != nil {
		return m.Score
	}
	return 0
}

func (m *Indicator) GetEntityId() string {
	if m != nil {
		return m.EntityId
	}
	return ""
}

func init() {
	proto.RegisterType((*Indicator)(nil), "mail.Indicator")
}

func init() {
	proto.RegisterFile("mail/indicator.proto", fileDescriptor_97fd999573541b11)
}

var fileDescriptor_97fd999573541b11 = []byte{
	// 237 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x8f, 0x3f, 0x4f, 0xc3, 0x30,
	0x10, 0xc5, 0xe5, 0x34, 0x2d, 0xc4, 0x45, 0x0c, 0x56, 0x07, 0x4b, 0x2c, 0x81, 0x29, 0x4b, 0xea,
	0x81, 0x15, 0x31, 0x20, 0x96, 0xae, 0x19, 0x59, 0xaa, 0x8b, 0x7d, 0x6a, 0x4f, 0xf5, 0x9f, 0xc8,
	0x71, 0x10, 0x7c, 0x38, 0xbe, 0x1b, 0x8a, 0x03, 0x52, 0xb7, 0x7b, 0xbf, 0x77, 0x7a, 0x4f, 0x8f,
	0xef, 0x1c, 0x90, 0x55, 0xe4, 0x0d, 0x69, 0x48, 0x21, 0xee, 0x87, 0x18, 0x52, 0x10, 0xe5, 0x4c,
	0x9f, 0x7e, 0x18, 0xaf, 0x0e, 0xff, 0x8e, 0xb8, 0xe7, 0x05, 0x19, 0xc9, 0x6a, 0xd6, 0x54, 0x5d,
	0x41, 0x46, 0x08, 0x5e, 0x7a, 0x70, 0x28, 0x8b, 0x4c, 0xf2, 0x2d, 0x76, 0x7c, 0xfd, 0x09, 0x76,
	0x42, 0xb9, 0xca, 0x70, 0x11, 0xa2, 0xe6, 0x5b, 0x83, 0xa3, 0x8e, 0x34, 0x24, 0x0a, 0x5e, 0x96,
	0xd9, 0xbb, 0x46, 0xe2, 0x91, 0xdf, 0xe9, 0x88, 0x90, 0xd0, 0x1c, 0x0d, 0x24, 0x94, 0xeb, 0x9a,
	0x35, 0xab, 0x6e, 0xfb, 0xc7, 0xde, 0x21, 0xe5, 0xe8, 0x51, 0x87, 0x88, 0x72, 0x53, 0xb3, 0x86,
	0x75, 0x8b, 0x10, 0x0f, 0xbc, 0x42, 0x9f, 0x28, 0x7d, 0x1f, 0xc9, 0xc8, 0x9b, 0x1c, 0x7c, 0xbb,
	0x80, 0x83, 0x79, 0x7b, 0xfd, 0x78, 0x39, 0x51, 0x3a, 0x4f, 0xfd, 0x5e, 0x07, 0xa7, 0x1c, 0x5d,
	0x70, 0xbc, 0x80, 0x25, 0x95, 0x37, 0xb6, 0xe3, 0x19, 0x22, 0x9a, 0xd6, 0x52, 0xdf, 0xe2, 0x17,
	0xb8, 0xc1, 0xa2, 0x3a, 0xa1, 0xc7, 0x38, 0x37, 0xaa, 0x79, 0x7f, 0xbf, 0xc9, 0x8f, 0xcf, 0xbf,
	0x01, 0x00, 0x00, 0xff, 0xff, 0xb8, 0x6b, 0x9e, 0x1e, 0x24, 0x01, 0x00, 0x00,
}
