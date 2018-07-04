// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/functions/v1beta2/operations.proto

package functions // import "google.golang.org/genproto/googleapis/cloud/functions/v1beta2"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import any "github.com/golang/protobuf/ptypes/any"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// A type of an operation.
type OperationType int32

const (
	// Unknown operation type.
	OperationType_OPERATION_UNSPECIFIED OperationType = 0
	// Triggered by CreateFunction call
	OperationType_CREATE_FUNCTION OperationType = 1
	// Triggered by UpdateFunction call
	OperationType_UPDATE_FUNCTION OperationType = 2
	// Triggered by DeleteFunction call.
	OperationType_DELETE_FUNCTION OperationType = 3
)

var OperationType_name = map[int32]string{
	0: "OPERATION_UNSPECIFIED",
	1: "CREATE_FUNCTION",
	2: "UPDATE_FUNCTION",
	3: "DELETE_FUNCTION",
}
var OperationType_value = map[string]int32{
	"OPERATION_UNSPECIFIED": 0,
	"CREATE_FUNCTION":       1,
	"UPDATE_FUNCTION":       2,
	"DELETE_FUNCTION":       3,
}

func (x OperationType) String() string {
	return proto.EnumName(OperationType_name, int32(x))
}
func (OperationType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_operations_c731a618eb615b2f, []int{0}
}

// Metadata describing an [Operation][google.longrunning.Operation]
type OperationMetadataV1Beta2 struct {
	// Target of the operation - for example
	// projects/project-1/locations/region-1/functions/function-1
	Target string `protobuf:"bytes,1,opt,name=target,proto3" json:"target,omitempty"`
	// Type of operation.
	Type OperationType `protobuf:"varint,2,opt,name=type,proto3,enum=google.cloud.functions.v1beta2.OperationType" json:"type,omitempty"`
	// The original request that started the operation.
	Request              *any.Any `protobuf:"bytes,3,opt,name=request,proto3" json:"request,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OperationMetadataV1Beta2) Reset()         { *m = OperationMetadataV1Beta2{} }
func (m *OperationMetadataV1Beta2) String() string { return proto.CompactTextString(m) }
func (*OperationMetadataV1Beta2) ProtoMessage()    {}
func (*OperationMetadataV1Beta2) Descriptor() ([]byte, []int) {
	return fileDescriptor_operations_c731a618eb615b2f, []int{0}
}
func (m *OperationMetadataV1Beta2) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OperationMetadataV1Beta2.Unmarshal(m, b)
}
func (m *OperationMetadataV1Beta2) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OperationMetadataV1Beta2.Marshal(b, m, deterministic)
}
func (dst *OperationMetadataV1Beta2) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OperationMetadataV1Beta2.Merge(dst, src)
}
func (m *OperationMetadataV1Beta2) XXX_Size() int {
	return xxx_messageInfo_OperationMetadataV1Beta2.Size(m)
}
func (m *OperationMetadataV1Beta2) XXX_DiscardUnknown() {
	xxx_messageInfo_OperationMetadataV1Beta2.DiscardUnknown(m)
}

var xxx_messageInfo_OperationMetadataV1Beta2 proto.InternalMessageInfo

func (m *OperationMetadataV1Beta2) GetTarget() string {
	if m != nil {
		return m.Target
	}
	return ""
}

func (m *OperationMetadataV1Beta2) GetType() OperationType {
	if m != nil {
		return m.Type
	}
	return OperationType_OPERATION_UNSPECIFIED
}

func (m *OperationMetadataV1Beta2) GetRequest() *any.Any {
	if m != nil {
		return m.Request
	}
	return nil
}

func init() {
	proto.RegisterType((*OperationMetadataV1Beta2)(nil), "google.cloud.functions.v1beta2.OperationMetadataV1Beta2")
	proto.RegisterEnum("google.cloud.functions.v1beta2.OperationType", OperationType_name, OperationType_value)
}

func init() {
	proto.RegisterFile("google/cloud/functions/v1beta2/operations.proto", fileDescriptor_operations_c731a618eb615b2f)
}

var fileDescriptor_operations_c731a618eb615b2f = []byte{
	// 333 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x91, 0x4f, 0x4f, 0xf2, 0x30,
	0x1c, 0xc7, 0x9f, 0xc2, 0x13, 0x8c, 0x35, 0x2a, 0x99, 0x7f, 0x32, 0x88, 0x31, 0x84, 0x13, 0x31,
	0xb1, 0x0d, 0x78, 0xf4, 0x34, 0xa0, 0x18, 0x12, 0x85, 0x65, 0x82, 0x07, 0x2f, 0xa4, 0x40, 0x69,
	0x96, 0xcc, 0xb6, 0x6e, 0x9d, 0xc9, 0x5e, 0x82, 0x2f, 0xc4, 0xf7, 0x69, 0x56, 0xba, 0x05, 0x0e,
	0xea, 0xb1, 0x9f, 0xf6, 0xf3, 0xed, 0xf7, 0x97, 0x1f, 0xc4, 0x5c, 0x4a, 0x1e, 0x31, 0xbc, 0x8a,
	0x64, 0xba, 0xc6, 0x9b, 0x54, 0xac, 0x74, 0x28, 0x45, 0x82, 0x3f, 0xba, 0x4b, 0xa6, 0x69, 0x0f,
	0x4b, 0xc5, 0x62, 0x6a, 0x10, 0x52, 0xb1, 0xd4, 0xd2, 0xb9, 0xde, 0x0a, 0xc8, 0x08, 0xa8, 0x14,
	0x90, 0x15, 0x9a, 0x57, 0x36, 0x90, 0xaa, 0x10, 0x53, 0x21, 0xa4, 0xde, 0xb5, 0x9b, 0x0d, 0x7b,
	0x6b, 0x4e, 0xcb, 0x74, 0x83, 0xa9, 0xc8, 0xb6, 0x57, 0xed, 0x2f, 0x00, 0xdd, 0x69, 0xf1, 0xdb,
	0x13, 0xd3, 0x74, 0x4d, 0x35, 0x7d, 0xe9, 0xf6, 0xf3, 0x54, 0xe7, 0x12, 0xd6, 0x34, 0x8d, 0x39,
	0xd3, 0x2e, 0x68, 0x81, 0xce, 0x61, 0x60, 0x4f, 0x8e, 0x07, 0xff, 0xeb, 0x4c, 0x31, 0xb7, 0xd2,
	0x02, 0x9d, 0x93, 0xde, 0x2d, 0xfa, 0xbd, 0x1c, 0x2a, 0xf3, 0x67, 0x99, 0x62, 0x81, 0x51, 0x1d,
	0x04, 0x0f, 0x62, 0xf6, 0x9e, 0xb2, 0x44, 0xbb, 0xd5, 0x16, 0xe8, 0x1c, 0xf5, 0xce, 0x8b, 0x94,
	0xa2, 0x24, 0xf2, 0x44, 0x16, 0x14, 0x8f, 0x6e, 0x42, 0x78, 0xbc, 0x17, 0xe3, 0x34, 0xe0, 0xc5,
	0xd4, 0x27, 0x81, 0x37, 0x1b, 0x4f, 0x27, 0x8b, 0xf9, 0xe4, 0xd9, 0x27, 0x83, 0xf1, 0x68, 0x4c,
	0x86, 0xf5, 0x7f, 0xce, 0x19, 0x3c, 0x1d, 0x04, 0xc4, 0x9b, 0x91, 0xc5, 0x68, 0x3e, 0x19, 0xe4,
	0x0f, 0xea, 0x20, 0x87, 0x73, 0x7f, 0xb8, 0x07, 0x2b, 0x39, 0x1c, 0x92, 0x47, 0xb2, 0x0b, 0xab,
	0xfd, 0x4f, 0x00, 0xdb, 0x2b, 0xf9, 0xf6, 0xc7, 0x54, 0x7d, 0x77, 0x54, 0xa0, 0xb2, 0x58, 0xe2,
	0xe7, 0xdd, 0x7d, 0xf0, 0xfa, 0x60, 0x5d, 0x2e, 0x23, 0x2a, 0x38, 0x92, 0x31, 0xc7, 0x9c, 0x09,
	0x33, 0x99, 0x5d, 0x3d, 0x55, 0x61, 0xf2, 0xd3, 0xfa, 0xef, 0x4b, 0xb2, 0xac, 0x19, 0xe7, 0xee,
	0x3b, 0x00, 0x00, 0xff, 0xff, 0x3c, 0x8a, 0xb1, 0x83, 0x31, 0x02, 0x00, 0x00,
}
