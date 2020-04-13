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
	// 189 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0xcf, 0x31, 0x4a, 0xc6, 0x40,
	0x10, 0x05, 0x60, 0x36, 0x7f, 0x12, 0xcd, 0x44, 0x2c, 0x86, 0x14, 0x0b, 0x36, 0xab, 0x55, 0x2a,
	0x2d, 0xbc, 0x82, 0x4d, 0xda, 0x5c, 0x20, 0xac, 0x3b, 0x53, 0x0c, 0x24, 0xbb, 0x61, 0x33, 0x0a,
	0x1e, 0xce, 0xbb, 0x89, 0x1b, 0x85, 0xbf, 0x9b, 0xf7, 0x3d, 0x78, 0x30, 0x30, 0x6c, 0x5e, 0xd6,
	0x17, 0x89, 0x24, 0xc1, 0x6b, 0xca, 0xcf, 0x7b, 0x4e, 0x9a, 0xb0, 0xfe, 0xd5, 0xa7, 0x6f, 0x03,
	0xdd, 0xf4, 0xdf, 0xe0, 0x3d, 0x54, 0x42, 0xd6, 0x38, 0x33, 0x76, 0x73, 0x25, 0x84, 0x08, 0x75,
	0xf4, 0x1b, 0xdb, 0xaa, 0x48, 0xb9, 0x71, 0x80, 0xe6, 0xd3, 0xaf, 0x1f, 0x6c, 0x2f, 0x05, 0xcf,
	0x80, 0x0e, 0x7a, 0xe2, 0x23, 0x64, 0xd9, 0x55, 0x52, 0xb4, 0x75, 0xe9, 0xae, 0x09, 0x1f, 0xe1,
	0x2e, 0x64, 0xf6, 0xca, 0xb4, 0x90, 0x57, 0xb6, 0x8d, 0x33, 0xe3, 0x65, 0xee, 0xff, 0xec, 0xcd,
	0x6b, 0x99, 0x3e, 0x42, 0xca, 0x6c, 0x5b, 0x67, 0x46, 0x33, 0x9f, 0x01, 0x1f, 0xa0, 0xe3, 0xa8,
	0xa2, 0x5f, 0x8b, 0x90, 0xbd, 0x29, 0xc3, 0xb7, 0x27, 0x4c, 0xf4, 0xde, 0x96, 0x67, 0x5e, 0x7f,
	0x02, 0x00, 0x00, 0xff, 0xff, 0xff, 0x97, 0xcb, 0x92, 0xe4, 0x00, 0x00, 0x00,
}
