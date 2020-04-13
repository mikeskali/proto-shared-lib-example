// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mail/mail_transaction.proto

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

type MailTransaction struct {
	Id                   string       `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	TxCreatedAt          int64        `protobuf:"varint,3,opt,name=tx_created_at,json=txCreatedAt,proto3" json:"tx_created_at,omitempty"`
	Indicators           []*Indicator `protobuf:"bytes,4,rep,name=indicators,proto3" json:"indicators,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *MailTransaction) Reset()         { *m = MailTransaction{} }
func (m *MailTransaction) String() string { return proto.CompactTextString(m) }
func (*MailTransaction) ProtoMessage()    {}
func (*MailTransaction) Descriptor() ([]byte, []int) {
	return fileDescriptor_dcb2371113915566, []int{0}
}

func (m *MailTransaction) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MailTransaction.Unmarshal(m, b)
}
func (m *MailTransaction) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MailTransaction.Marshal(b, m, deterministic)
}
func (m *MailTransaction) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MailTransaction.Merge(m, src)
}
func (m *MailTransaction) XXX_Size() int {
	return xxx_messageInfo_MailTransaction.Size(m)
}
func (m *MailTransaction) XXX_DiscardUnknown() {
	xxx_messageInfo_MailTransaction.DiscardUnknown(m)
}

var xxx_messageInfo_MailTransaction proto.InternalMessageInfo

func (m *MailTransaction) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *MailTransaction) GetTxCreatedAt() int64 {
	if m != nil {
		return m.TxCreatedAt
	}
	return 0
}

func (m *MailTransaction) GetIndicators() []*Indicator {
	if m != nil {
		return m.Indicators
	}
	return nil
}

func init() {
	proto.RegisterType((*MailTransaction)(nil), "mail.MailTransaction")
}

func init() {
	proto.RegisterFile("mail/mail_transaction.proto", fileDescriptor_dcb2371113915566)
}

var fileDescriptor_dcb2371113915566 = []byte{
	// 159 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0xce, 0x4d, 0xcc, 0xcc,
	0xd1, 0x07, 0x11, 0xf1, 0x25, 0x45, 0x89, 0x79, 0xc5, 0x89, 0xc9, 0x25, 0x99, 0xf9, 0x79, 0x7a,
	0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x2c, 0x20, 0x71, 0x29, 0x11, 0xb0, 0x92, 0xcc, 0xbc, 0x94,
	0xcc, 0xe4, 0xc4, 0x92, 0xfc, 0x22, 0x88, 0x9c, 0x52, 0x19, 0x17, 0xbf, 0x6f, 0x62, 0x66, 0x4e,
	0x08, 0x42, 0x93, 0x10, 0x1f, 0x17, 0x53, 0x66, 0x8a, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10,
	0x53, 0x66, 0x8a, 0x90, 0x12, 0x17, 0x6f, 0x49, 0x45, 0x7c, 0x72, 0x51, 0x6a, 0x62, 0x49, 0x6a,
	0x4a, 0x7c, 0x62, 0x89, 0x04, 0xb3, 0x02, 0xa3, 0x06, 0x73, 0x10, 0x77, 0x49, 0x85, 0x33, 0x44,
	0xcc, 0xb1, 0x44, 0x48, 0x9f, 0x8b, 0x0b, 0x6e, 0x72, 0xb1, 0x04, 0x8b, 0x02, 0xb3, 0x06, 0xb7,
	0x11, 0xbf, 0x1e, 0xc8, 0x46, 0x3d, 0x4f, 0x98, 0x78, 0x10, 0x92, 0x92, 0x24, 0x36, 0xb0, 0xf5,
	0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff, 0xef, 0x9b, 0xfc, 0x00, 0xb9, 0x00, 0x00, 0x00,
}
