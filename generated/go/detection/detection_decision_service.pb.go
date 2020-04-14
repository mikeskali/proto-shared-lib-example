// Code generated by protoc-gen-go. DO NOT EDIT.
// source: detection/detection_decision_service.proto

package detection

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	mail "github.com/mikeskali/proto-shared-lib-example/generated/go/mail"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

type DetectionDecision_Severity int32

const (
	DetectionDecision_UNDEFINED_SEVERITY DetectionDecision_Severity = 0
	DetectionDecision_LOW                DetectionDecision_Severity = 1
	DetectionDecision_MEDIUM             DetectionDecision_Severity = 2
	DetectionDecision_HIGH               DetectionDecision_Severity = 3
)

var DetectionDecision_Severity_name = map[int32]string{
	0: "UNDEFINED_SEVERITY",
	1: "LOW",
	2: "MEDIUM",
	3: "HIGH",
}

var DetectionDecision_Severity_value = map[string]int32{
	"UNDEFINED_SEVERITY": 0,
	"LOW":                1,
	"MEDIUM":             2,
	"HIGH":               3,
}

func (x DetectionDecision_Severity) String() string {
	return proto.EnumName(DetectionDecision_Severity_name, int32(x))
}

func (DetectionDecision_Severity) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_251946e35ddc8419, []int{0, 0}
}

type DetectionDecision_Classification int32

const (
	DetectionDecision_UNDEFINED_CLASSIFICATION DetectionDecision_Classification = 0
	DetectionDecision_CLEAN                    DetectionDecision_Classification = 1
	DetectionDecision_MALICIOUS                DetectionDecision_Classification = 2
)

var DetectionDecision_Classification_name = map[int32]string{
	0: "UNDEFINED_CLASSIFICATION",
	1: "CLEAN",
	2: "MALICIOUS",
}

var DetectionDecision_Classification_value = map[string]int32{
	"UNDEFINED_CLASSIFICATION": 0,
	"CLEAN":                    1,
	"MALICIOUS":                2,
}

func (x DetectionDecision_Classification) String() string {
	return proto.EnumName(DetectionDecision_Classification_name, int32(x))
}

func (DetectionDecision_Classification) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_251946e35ddc8419, []int{0, 1}
}

type DetectionDecision struct {
	Classification       DetectionDecision_Classification `protobuf:"varint,1,opt,name=classification,proto3,enum=detection.DetectionDecision_Classification" json:"classification,omitempty"`
	Severity             DetectionDecision_Severity       `protobuf:"varint,2,opt,name=severity,proto3,enum=detection.DetectionDecision_Severity" json:"severity,omitempty"`
	DetectedIndicators   []*mail.Indicator                `protobuf:"bytes,3,rep,name=detectedIndicators,proto3" json:"detectedIndicators,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                         `json:"-"`
	XXX_unrecognized     []byte                           `json:"-"`
	XXX_sizecache        int32                            `json:"-"`
}

func (m *DetectionDecision) Reset()         { *m = DetectionDecision{} }
func (m *DetectionDecision) String() string { return proto.CompactTextString(m) }
func (*DetectionDecision) ProtoMessage()    {}
func (*DetectionDecision) Descriptor() ([]byte, []int) {
	return fileDescriptor_251946e35ddc8419, []int{0}
}

func (m *DetectionDecision) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DetectionDecision.Unmarshal(m, b)
}
func (m *DetectionDecision) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DetectionDecision.Marshal(b, m, deterministic)
}
func (m *DetectionDecision) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DetectionDecision.Merge(m, src)
}
func (m *DetectionDecision) XXX_Size() int {
	return xxx_messageInfo_DetectionDecision.Size(m)
}
func (m *DetectionDecision) XXX_DiscardUnknown() {
	xxx_messageInfo_DetectionDecision.DiscardUnknown(m)
}

var xxx_messageInfo_DetectionDecision proto.InternalMessageInfo

func (m *DetectionDecision) GetClassification() DetectionDecision_Classification {
	if m != nil {
		return m.Classification
	}
	return DetectionDecision_UNDEFINED_CLASSIFICATION
}

func (m *DetectionDecision) GetSeverity() DetectionDecision_Severity {
	if m != nil {
		return m.Severity
	}
	return DetectionDecision_UNDEFINED_SEVERITY
}

func (m *DetectionDecision) GetDetectedIndicators() []*mail.Indicator {
	if m != nil {
		return m.DetectedIndicators
	}
	return nil
}

type DetectionDecisionRequest struct {
	TxId                 string            `protobuf:"bytes,1,opt,name=tx_id,json=txId,proto3" json:"tx_id,omitempty"`
	Indicators           []*mail.Indicator `protobuf:"bytes,2,rep,name=indicators,proto3" json:"indicators,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *DetectionDecisionRequest) Reset()         { *m = DetectionDecisionRequest{} }
func (m *DetectionDecisionRequest) String() string { return proto.CompactTextString(m) }
func (*DetectionDecisionRequest) ProtoMessage()    {}
func (*DetectionDecisionRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_251946e35ddc8419, []int{1}
}

func (m *DetectionDecisionRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DetectionDecisionRequest.Unmarshal(m, b)
}
func (m *DetectionDecisionRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DetectionDecisionRequest.Marshal(b, m, deterministic)
}
func (m *DetectionDecisionRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DetectionDecisionRequest.Merge(m, src)
}
func (m *DetectionDecisionRequest) XXX_Size() int {
	return xxx_messageInfo_DetectionDecisionRequest.Size(m)
}
func (m *DetectionDecisionRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_DetectionDecisionRequest.DiscardUnknown(m)
}

var xxx_messageInfo_DetectionDecisionRequest proto.InternalMessageInfo

func (m *DetectionDecisionRequest) GetTxId() string {
	if m != nil {
		return m.TxId
	}
	return ""
}

func (m *DetectionDecisionRequest) GetIndicators() []*mail.Indicator {
	if m != nil {
		return m.Indicators
	}
	return nil
}

type DetectionDecisionResponse struct {
	DetectionDecision    *DetectionDecision `protobuf:"bytes,1,opt,name=detection_decision,json=detectionDecision,proto3" json:"detection_decision,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *DetectionDecisionResponse) Reset()         { *m = DetectionDecisionResponse{} }
func (m *DetectionDecisionResponse) String() string { return proto.CompactTextString(m) }
func (*DetectionDecisionResponse) ProtoMessage()    {}
func (*DetectionDecisionResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_251946e35ddc8419, []int{2}
}

func (m *DetectionDecisionResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DetectionDecisionResponse.Unmarshal(m, b)
}
func (m *DetectionDecisionResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DetectionDecisionResponse.Marshal(b, m, deterministic)
}
func (m *DetectionDecisionResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DetectionDecisionResponse.Merge(m, src)
}
func (m *DetectionDecisionResponse) XXX_Size() int {
	return xxx_messageInfo_DetectionDecisionResponse.Size(m)
}
func (m *DetectionDecisionResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_DetectionDecisionResponse.DiscardUnknown(m)
}

var xxx_messageInfo_DetectionDecisionResponse proto.InternalMessageInfo

func (m *DetectionDecisionResponse) GetDetectionDecision() *DetectionDecision {
	if m != nil {
		return m.DetectionDecision
	}
	return nil
}

func init() {
	proto.RegisterEnum("detection.DetectionDecision_Severity", DetectionDecision_Severity_name, DetectionDecision_Severity_value)
	proto.RegisterEnum("detection.DetectionDecision_Classification", DetectionDecision_Classification_name, DetectionDecision_Classification_value)
	proto.RegisterType((*DetectionDecision)(nil), "detection.DetectionDecision")
	proto.RegisterType((*DetectionDecisionRequest)(nil), "detection.DetectionDecisionRequest")
	proto.RegisterType((*DetectionDecisionResponse)(nil), "detection.DetectionDecisionResponse")
}

func init() {
	proto.RegisterFile("detection/detection_decision_service.proto", fileDescriptor_251946e35ddc8419)
}

var fileDescriptor_251946e35ddc8419 = []byte{
	// 447 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0x5f, 0x6f, 0xd2, 0x60,
	0x14, 0xc6, 0x47, 0xd9, 0x26, 0x9c, 0x45, 0xec, 0x8e, 0x8b, 0xa9, 0xcb, 0x2e, 0x96, 0xaa, 0xc9,
	0xa2, 0xa1, 0x4d, 0xf0, 0x03, 0x98, 0xda, 0x96, 0xf1, 0x46, 0xfe, 0x24, 0xed, 0xd0, 0xe8, 0x0d,
	0x96, 0xf6, 0x08, 0x6f, 0x56, 0x5a, 0xec, 0xfb, 0x6e, 0xc1, 0x0b, 0x3f, 0xb8, 0x77, 0x86, 0x76,
	0x94, 0x0d, 0x08, 0x77, 0x6f, 0x4e, 0x9e, 0xe7, 0x39, 0xf9, 0x9d, 0x3e, 0x85, 0xf7, 0x11, 0x49,
	0x0a, 0x25, 0x4f, 0x13, 0xb3, 0x7c, 0x8d, 0x22, 0x0a, 0xb9, 0x58, 0x3e, 0x04, 0x65, 0xf7, 0x3c,
	0x24, 0x63, 0x9e, 0xa5, 0x32, 0xc5, 0x7a, 0xa9, 0x38, 0x3f, 0x9b, 0x05, 0x3c, 0x36, 0x79, 0x12,
	0xf1, 0x30, 0x90, 0x69, 0x56, 0x08, 0xf4, 0x7f, 0x0a, 0x9c, 0x3a, 0x2b, 0x8d, 0xf3, 0x10, 0x82,
	0x3e, 0x34, 0xc2, 0x38, 0x10, 0x82, 0xff, 0x5a, 0x8a, 0x79, 0x9a, 0x68, 0x95, 0xcb, 0xca, 0x55,
	0xa3, 0xf5, 0xc1, 0x28, 0xf3, 0x8c, 0x2d, 0x97, 0x61, 0x3f, 0xb1, 0x78, 0x1b, 0x11, 0x68, 0x41,
	0x4d, 0xd0, 0x3d, 0x65, 0x5c, 0xfe, 0xd1, 0x94, 0x3c, 0xee, 0xdd, 0xde, 0x38, 0xff, 0x41, 0xec,
	0x95, 0x36, 0xfc, 0x04, 0x58, 0x38, 0x28, 0x62, 0x2b, 0x10, 0xa1, 0x55, 0x2f, 0xab, 0x57, 0x27,
	0xad, 0x17, 0xc6, 0x12, 0xd0, 0x28, 0xe7, 0xde, 0x0e, 0xa9, 0x6e, 0x41, 0x6d, 0x15, 0x8b, 0xaf,
	0x00, 0x87, 0x7d, 0xc7, 0x6d, 0xb3, 0xbe, 0xeb, 0x8c, 0x7c, 0xf7, 0xab, 0xeb, 0xb1, 0x9b, 0xef,
	0xea, 0x01, 0x3e, 0x83, 0x6a, 0x77, 0xf0, 0x4d, 0xad, 0x20, 0xc0, 0x71, 0xcf, 0x75, 0xd8, 0xb0,
	0xa7, 0x2a, 0x58, 0x83, 0xc3, 0x0e, 0xbb, 0xee, 0xa8, 0x55, 0xbd, 0x03, 0x8d, 0xa7, 0xa0, 0x78,
	0x01, 0xda, 0x3a, 0xc8, 0xee, 0x5a, 0xbe, 0xcf, 0xda, 0xcc, 0xb6, 0x6e, 0xd8, 0xa0, 0xaf, 0x1e,
	0x60, 0x1d, 0x8e, 0xec, 0xae, 0x6b, 0xf5, 0xd5, 0x0a, 0x3e, 0x87, 0x7a, 0xcf, 0xea, 0x32, 0x9b,
	0x0d, 0x86, 0xbe, 0xaa, 0xe8, 0x3f, 0x41, 0xdb, 0xa2, 0xf6, 0xe8, 0xf7, 0x1d, 0x09, 0x89, 0x2f,
	0xe1, 0x48, 0x2e, 0x46, 0x3c, 0xca, 0x0f, 0x5f, 0xf7, 0x0e, 0xe5, 0x82, 0x45, 0x68, 0x02, 0xf0,
	0x35, 0xb6, 0xb2, 0x1b, 0xfb, 0x91, 0x44, 0x9f, 0xc2, 0xeb, 0x1d, 0x1b, 0xc4, 0x3c, 0x4d, 0x04,
	0xe1, 0x97, 0xd5, 0x31, 0x1f, 0xf7, 0x27, 0xdf, 0x77, 0xd2, 0xba, 0xd8, 0xf7, 0x65, 0xbc, 0xd3,
	0x68, 0x73, 0xd4, 0xfa, 0xbb, 0x83, 0xc5, 0x2f, 0xaa, 0x88, 0x01, 0x9c, 0x5d, 0x93, 0xdc, 0x6e,
	0xd9, 0x9b, 0xbd, 0x4b, 0x8a, 0x43, 0x9c, 0xbf, 0xdd, 0x2f, 0x2a, 0x58, 0x3e, 0xb7, 0x7f, 0x38,
	0x13, 0x2e, 0xa7, 0x77, 0x63, 0x23, 0x4c, 0x67, 0xe6, 0x8c, 0xdf, 0x92, 0xb8, 0x0d, 0x62, 0x6e,
	0xe6, 0x25, 0x6f, 0x8a, 0x69, 0x90, 0x51, 0xd4, 0x8c, 0xf9, 0xb8, 0x49, 0x8b, 0x60, 0x36, 0x8f,
	0xc9, 0x9c, 0x50, 0x42, 0x59, 0x20, 0x29, 0x32, 0x27, 0xe9, 0xfa, 0x37, 0x1a, 0x1f, 0xe7, 0x86,
	0x8f, 0xff, 0x03, 0x00, 0x00, 0xff, 0xff, 0xb3, 0x21, 0xb6, 0x7b, 0x64, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// DetectionDecisionServiceClient is the client API for DetectionDecisionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type DetectionDecisionServiceClient interface {
	GetDetectionDecision(ctx context.Context, in *DetectionDecisionRequest, opts ...grpc.CallOption) (*DetectionDecisionResponse, error)
}

type detectionDecisionServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewDetectionDecisionServiceClient(cc grpc.ClientConnInterface) DetectionDecisionServiceClient {
	return &detectionDecisionServiceClient{cc}
}

func (c *detectionDecisionServiceClient) GetDetectionDecision(ctx context.Context, in *DetectionDecisionRequest, opts ...grpc.CallOption) (*DetectionDecisionResponse, error) {
	out := new(DetectionDecisionResponse)
	err := c.cc.Invoke(ctx, "/detection.DetectionDecisionService/GetDetectionDecision", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DetectionDecisionServiceServer is the server API for DetectionDecisionService service.
type DetectionDecisionServiceServer interface {
	GetDetectionDecision(context.Context, *DetectionDecisionRequest) (*DetectionDecisionResponse, error)
}

// UnimplementedDetectionDecisionServiceServer can be embedded to have forward compatible implementations.
type UnimplementedDetectionDecisionServiceServer struct {
}

func (*UnimplementedDetectionDecisionServiceServer) GetDetectionDecision(ctx context.Context, req *DetectionDecisionRequest) (*DetectionDecisionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetDetectionDecision not implemented")
}

func RegisterDetectionDecisionServiceServer(s *grpc.Server, srv DetectionDecisionServiceServer) {
	s.RegisterService(&_DetectionDecisionService_serviceDesc, srv)
}

func _DetectionDecisionService_GetDetectionDecision_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DetectionDecisionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DetectionDecisionServiceServer).GetDetectionDecision(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/detection.DetectionDecisionService/GetDetectionDecision",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DetectionDecisionServiceServer).GetDetectionDecision(ctx, req.(*DetectionDecisionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _DetectionDecisionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "detection.DetectionDecisionService",
	HandlerType: (*DetectionDecisionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetDetectionDecision",
			Handler:    _DetectionDecisionService_GetDetectionDecision_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "detection/detection_decision_service.proto",
}
