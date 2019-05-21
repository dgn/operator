// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/api/source_info.proto

package serviceconfig // import "google.golang.org/genproto/googleapis/api/serviceconfig"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import any "github.com/golang/protobuf/ptypes/any"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// Source information used to create a Service Config
type SourceInfo struct {
	// All files used during config generation.
	SourceFiles          []*any.Any `protobuf:"bytes,1,rep,name=source_files,json=sourceFiles,proto3" json:"source_files,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *SourceInfo) Reset()         { *m = SourceInfo{} }
func (m *SourceInfo) String() string { return proto.CompactTextString(m) }
func (*SourceInfo) ProtoMessage()    {}
func (*SourceInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_source_info_a3e4837d95d932be, []int{0}
}
func (m *SourceInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SourceInfo.Unmarshal(m, b)
}
func (m *SourceInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SourceInfo.Marshal(b, m, deterministic)
}
func (dst *SourceInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SourceInfo.Merge(dst, src)
}
func (m *SourceInfo) XXX_Size() int {
	return xxx_messageInfo_SourceInfo.Size(m)
}
func (m *SourceInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_SourceInfo.DiscardUnknown(m)
}

var xxx_messageInfo_SourceInfo proto.InternalMessageInfo

func (m *SourceInfo) GetSourceFiles() []*any.Any {
	if m != nil {
		return m.SourceFiles
	}
	return nil
}

func init() {
	proto.RegisterType((*SourceInfo)(nil), "google.api.SourceInfo")
}

func init() {
	proto.RegisterFile("google/api/source_info.proto", fileDescriptor_source_info_a3e4837d95d932be)
}

var fileDescriptor_source_info_a3e4837d95d932be = []byte{
	// 198 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x49, 0xcf, 0xcf, 0x4f,
	0xcf, 0x49, 0xd5, 0x4f, 0x2c, 0xc8, 0xd4, 0x2f, 0xce, 0x2f, 0x2d, 0x4a, 0x4e, 0x8d, 0xcf, 0xcc,
	0x4b, 0xcb, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0xe2, 0x82, 0xc8, 0xea, 0x25, 0x16, 0x64,
	0x4a, 0x49, 0x42, 0x55, 0x82, 0x65, 0x92, 0x4a, 0xd3, 0xf4, 0x13, 0xf3, 0x2a, 0x21, 0xca, 0x94,
	0x5c, 0xb9, 0xb8, 0x82, 0xc1, 0x7a, 0x3d, 0xf3, 0xd2, 0xf2, 0x85, 0xcc, 0xb9, 0x78, 0xa0, 0x26,
	0xa5, 0x65, 0xe6, 0xa4, 0x16, 0x4b, 0x30, 0x2a, 0x30, 0x6b, 0x70, 0x1b, 0x89, 0xe8, 0x41, 0xcd,
	0x82, 0xe9, 0xd7, 0x73, 0xcc, 0xab, 0x0c, 0xe2, 0x86, 0xa8, 0x74, 0x03, 0x29, 0x74, 0x2a, 0xe4,
	0xe2, 0x4b, 0xce, 0xcf, 0xd5, 0x43, 0xd8, 0xe9, 0xc4, 0x8f, 0x30, 0x36, 0x00, 0xa4, 0x2d, 0x80,
	0x31, 0xca, 0x15, 0x2a, 0x9d, 0x9e, 0x9f, 0x93, 0x98, 0x97, 0xae, 0x97, 0x5f, 0x94, 0xae, 0x9f,
	0x9e, 0x9a, 0x07, 0x36, 0x54, 0x1f, 0x22, 0x95, 0x58, 0x90, 0x59, 0x0c, 0xf1, 0x4f, 0x6a, 0x51,
	0x59, 0x66, 0x72, 0x6a, 0x72, 0x7e, 0x5e, 0x5a, 0x66, 0xba, 0x35, 0x0a, 0x6f, 0x11, 0x13, 0x8b,
	0xbb, 0x63, 0x80, 0x67, 0x12, 0x1b, 0x58, 0xa3, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff, 0xd4, 0x78,
	0x5d, 0xab, 0x07, 0x01, 0x00, 0x00,
}