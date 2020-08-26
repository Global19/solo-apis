// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/v1/options/healthcheck/healthcheck.proto

package healthcheck

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	_ "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

// Add this config to a Listener/Gateway to Enable Envoy Health Checks on that port
type HealthCheck struct {
	// match health check requests using this exact path
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HealthCheck) Reset()         { *m = HealthCheck{} }
func (m *HealthCheck) String() string { return proto.CompactTextString(m) }
func (*HealthCheck) ProtoMessage()    {}
func (*HealthCheck) Descriptor() ([]byte, []int) {
	return fileDescriptor_a7f57f647fce7138, []int{0}
}
func (m *HealthCheck) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HealthCheck.Unmarshal(m, b)
}
func (m *HealthCheck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HealthCheck.Marshal(b, m, deterministic)
}
func (m *HealthCheck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HealthCheck.Merge(m, src)
}
func (m *HealthCheck) XXX_Size() int {
	return xxx_messageInfo_HealthCheck.Size(m)
}
func (m *HealthCheck) XXX_DiscardUnknown() {
	xxx_messageInfo_HealthCheck.DiscardUnknown(m)
}

var xxx_messageInfo_HealthCheck proto.InternalMessageInfo

func (m *HealthCheck) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func init() {
	proto.RegisterType((*HealthCheck)(nil), "healthcheck.options.gloo.solo.io.HealthCheck")
}

func init() {
	proto.RegisterFile("github.com/solo-io/solo-apis/api/gloo/gloo/v1/options/healthcheck/healthcheck.proto", fileDescriptor_a7f57f647fce7138)
}

var fileDescriptor_a7f57f647fce7138 = []byte{
	// 227 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x50, 0x3b, 0x4e, 0x03, 0x31,
	0x10, 0x95, 0xa5, 0x08, 0x89, 0xa5, 0xb3, 0x28, 0x20, 0x45, 0x14, 0xa8, 0x68, 0xf0, 0x28, 0xe2,
	0x06, 0xd0, 0xa4, 0x0e, 0x1d, 0x9d, 0x77, 0xe5, 0xd8, 0xd6, 0x1a, 0x66, 0x64, 0xcf, 0x2e, 0x7b,
	0x24, 0x8e, 0xc0, 0x79, 0xb8, 0x03, 0x3d, 0xb2, 0xbd, 0xa0, 0x6d, 0x48, 0x63, 0x3f, 0xcf, 0xe7,
	0xbd, 0xe7, 0xd7, 0x3c, 0x5b, 0xcf, 0x6e, 0x68, 0x55, 0x87, 0xaf, 0x90, 0x30, 0xe0, 0xbd, 0xc7,
	0x7a, 0x6b, 0xf2, 0x09, 0x34, 0x79, 0xb0, 0x01, 0xb1, 0x1e, 0xe3, 0x0e, 0x90, 0xd8, 0xe3, 0x5b,
	0x02, 0x67, 0x74, 0x60, 0xd7, 0x39, 0xd3, 0xf5, 0x4b, 0xac, 0x28, 0x22, 0xa3, 0xdc, 0x2e, 0x4b,
	0xf3, 0x8a, 0xca, 0x14, 0x2a, 0x53, 0x2b, 0x8f, 0xeb, 0x4b, 0x8b, 0x16, 0xcb, 0x30, 0x64, 0x54,
	0xf7, 0xd6, 0xd7, 0x45, 0xb9, 0xf7, 0x5c, 0x84, 0xc7, 0x1d, 0x44, 0x73, 0x9c, 0x5b, 0xd2, 0x4c,
	0x5c, 0xe7, 0xcd, 0xc4, 0x73, 0x6d, 0x63, 0x11, 0x6d, 0x30, 0x50, 0x5e, 0xed, 0x70, 0x84, 0xf7,
	0xa8, 0x89, 0x4c, 0x4c, 0xb5, 0x7f, 0x7b, 0xd3, 0x5c, 0xec, 0x8b, 0x91, 0xa7, 0x6c, 0x44, 0xca,
	0x66, 0x45, 0x9a, 0xdd, 0x95, 0xd8, 0x8a, 0xbb, 0xf3, 0x43, 0xc1, 0x8f, 0x87, 0xcf, 0xef, 0x95,
	0xf8, 0xf8, 0xda, 0x88, 0x97, 0xfd, 0xc9, 0x20, 0xa8, 0xb7, 0x7f, 0x61, 0xfc, 0x7e, 0xe2, 0x9f,
	0x3c, 0xda, 0xb3, 0xa2, 0xfe, 0xf0, 0x13, 0x00, 0x00, 0xff, 0xff, 0xe6, 0x74, 0x8b, 0xfb, 0x5b,
	0x01, 0x00, 0x00,
}

func (this *HealthCheck) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*HealthCheck)
	if !ok {
		that2, ok := that.(HealthCheck)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Path != that1.Path {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}