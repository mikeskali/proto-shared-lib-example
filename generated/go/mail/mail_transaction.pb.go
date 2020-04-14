// Code generated by protoc-gen-go. DO NOT EDIT.
// source: mail/mail_transaction.proto

package mail

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
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
	// 242 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xc1, 0x4a, 0xc3, 0x40,
	0x10, 0x86, 0xd9, 0xa4, 0x14, 0xdc, 0x22, 0x85, 0x20, 0x12, 0xeb, 0xc1, 0xe0, 0x29, 0x08, 0xc9,
	0x82, 0x5e, 0x45, 0xb0, 0x9e, 0x3c, 0x78, 0x09, 0x1e, 0xc4, 0x4b, 0x98, 0x64, 0x87, 0x76, 0xe8,
	0x26, 0x5b, 0x36, 0xa3, 0xe4, 0x09, 0x7c, 0x4d, 0xc1, 0x27, 0x91, 0xcd, 0x62, 0xed, 0x65, 0x59,
	0xfe, 0xff, 0x63, 0xe6, 0x63, 0xe4, 0x65, 0x07, 0x64, 0x94, 0x7f, 0x6a, 0x76, 0xd0, 0x0f, 0xd0,
	0x32, 0xd9, 0xbe, 0xdc, 0x3b, 0xcb, 0x36, 0x99, 0xf9, 0x7c, 0x75, 0x36, 0x21, 0xd4, 0x6b, 0x6a,
	0x81, 0xad, 0x0b, 0xdd, 0xea, 0xe2, 0x13, 0x0c, 0x69, 0x1f, 0xa8, 0xc3, 0x2f, 0x54, 0xd7, 0x5f,
	0x42, 0x2e, 0x5f, 0x80, 0xcc, 0xeb, 0xff, 0xc0, 0xe4, 0x5c, 0x46, 0xa4, 0x53, 0x91, 0x89, 0xfc,
	0x64, 0x3d, 0xff, 0xf9, 0xbe, 0x8a, 0xde, 0x44, 0x15, 0x91, 0x4e, 0x6e, 0xe4, 0x29, 0x8f, 0x75,
	0xeb, 0x10, 0x18, 0x75, 0x0d, 0x9c, 0xc6, 0x99, 0xc8, 0xe3, 0x80, 0x64, 0xa2, 0x5a, 0xf0, 0xf8,
	0x14, 0xba, 0x47, 0x4e, 0x94, 0x94, 0x07, 0x8b, 0x21, 0x9d, 0x65, 0x71, 0xbe, 0xb8, 0x5d, 0x96,
	0xde, 0xae, 0x7c, 0xfe, 0xcb, 0xab, 0x23, 0x64, 0xfd, 0xf0, 0x7e, 0xbf, 0x21, 0xde, 0x7e, 0x34,
	0x65, 0x6b, 0x3b, 0xd5, 0xd1, 0x0e, 0x87, 0x1d, 0x18, 0x52, 0x93, 0x66, 0x31, 0x6c, 0xc1, 0xa1,
	0x2e, 0x0c, 0x35, 0x05, 0x8e, 0xd0, 0xed, 0x0d, 0xaa, 0x0d, 0xf6, 0xe8, 0xfc, 0xb6, 0xe9, 0x22,
	0xcd, 0x7c, 0x02, 0xef, 0x7e, 0x03, 0x00, 0x00, 0xff, 0xff, 0x58, 0xef, 0x78, 0x6d, 0x25, 0x01,
	0x00, 0x00,
}
