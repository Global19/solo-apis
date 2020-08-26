// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo-fed/fed/core/v1/placement.proto

package v1

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
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

type PlacementStatus_State int32

const (
	// UNKNOWN indicates that the system does not know the placement status of the resource.
	PlacementStatus_UNKNOWN PlacementStatus_State = 0
	// PLACED indicates that the resource has been placed as specified.
	PlacementStatus_PLACED PlacementStatus_State = 1
	// FAILED indicates that the resource could not be placed in a specified destination.
	PlacementStatus_FAILED PlacementStatus_State = 2
	// STALE indicates that the resource continues to be present in a destination that is no longer specified.
	PlacementStatus_STALE PlacementStatus_State = 3
	// INVALID indicates that the resource cannot be placed as specified because one or more destinations do not exist.
	PlacementStatus_INVALID PlacementStatus_State = 4
	// PENDING indicates that the resource is waiting to be processed.
	PlacementStatus_PENDING PlacementStatus_State = 5
)

var PlacementStatus_State_name = map[int32]string{
	0: "UNKNOWN",
	1: "PLACED",
	2: "FAILED",
	3: "STALE",
	4: "INVALID",
	5: "PENDING",
}

var PlacementStatus_State_value = map[string]int32{
	"UNKNOWN": 0,
	"PLACED":  1,
	"FAILED":  2,
	"STALE":   3,
	"INVALID": 4,
	"PENDING": 5,
}

func (x PlacementStatus_State) String() string {
	return proto.EnumName(PlacementStatus_State_name, int32(x))
}

func (PlacementStatus_State) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_09a22494930a46b3, []int{1, 0}
}

type TemplateMetadata struct {
	Annotations          map[string]string `protobuf:"bytes,1,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Labels               map[string]string `protobuf:"bytes,2,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Name                 string            `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *TemplateMetadata) Reset()         { *m = TemplateMetadata{} }
func (m *TemplateMetadata) String() string { return proto.CompactTextString(m) }
func (*TemplateMetadata) ProtoMessage()    {}
func (*TemplateMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_09a22494930a46b3, []int{0}
}
func (m *TemplateMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TemplateMetadata.Unmarshal(m, b)
}
func (m *TemplateMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TemplateMetadata.Marshal(b, m, deterministic)
}
func (m *TemplateMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TemplateMetadata.Merge(m, src)
}
func (m *TemplateMetadata) XXX_Size() int {
	return xxx_messageInfo_TemplateMetadata.Size(m)
}
func (m *TemplateMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_TemplateMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_TemplateMetadata proto.InternalMessageInfo

func (m *TemplateMetadata) GetAnnotations() map[string]string {
	if m != nil {
		return m.Annotations
	}
	return nil
}

func (m *TemplateMetadata) GetLabels() map[string]string {
	if m != nil {
		return m.Labels
	}
	return nil
}

func (m *TemplateMetadata) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type PlacementStatus struct {
	Clusters map[string]*PlacementStatus_Cluster `protobuf:"bytes,1,rep,name=clusters,proto3" json:"clusters,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	State    PlacementStatus_State               `protobuf:"varint,2,opt,name=state,proto3,enum=core.fed.solo.io.PlacementStatus_State" json:"state,omitempty"`
	Message  string                              `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	// metadata.Generation of the resource which has been processed
	ObservedGeneration int64 `protobuf:"varint,4,opt,name=observed_generation,json=observedGeneration,proto3" json:"observed_generation,omitempty"`
	//
	//A field indicating the entity responsible for writing this status.
	//This is useful for determining if the pod has been restarted since the resource was processed.
	//Typically this value will be set to metadata.name of the pod
	WrittenBy            string   `protobuf:"bytes,5,opt,name=written_by,json=writtenBy,proto3" json:"written_by,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PlacementStatus) Reset()         { *m = PlacementStatus{} }
func (m *PlacementStatus) String() string { return proto.CompactTextString(m) }
func (*PlacementStatus) ProtoMessage()    {}
func (*PlacementStatus) Descriptor() ([]byte, []int) {
	return fileDescriptor_09a22494930a46b3, []int{1}
}
func (m *PlacementStatus) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlacementStatus.Unmarshal(m, b)
}
func (m *PlacementStatus) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlacementStatus.Marshal(b, m, deterministic)
}
func (m *PlacementStatus) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlacementStatus.Merge(m, src)
}
func (m *PlacementStatus) XXX_Size() int {
	return xxx_messageInfo_PlacementStatus.Size(m)
}
func (m *PlacementStatus) XXX_DiscardUnknown() {
	xxx_messageInfo_PlacementStatus.DiscardUnknown(m)
}

var xxx_messageInfo_PlacementStatus proto.InternalMessageInfo

func (m *PlacementStatus) GetClusters() map[string]*PlacementStatus_Cluster {
	if m != nil {
		return m.Clusters
	}
	return nil
}

func (m *PlacementStatus) GetState() PlacementStatus_State {
	if m != nil {
		return m.State
	}
	return PlacementStatus_UNKNOWN
}

func (m *PlacementStatus) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func (m *PlacementStatus) GetObservedGeneration() int64 {
	if m != nil {
		return m.ObservedGeneration
	}
	return 0
}

func (m *PlacementStatus) GetWrittenBy() string {
	if m != nil {
		return m.WrittenBy
	}
	return ""
}

type PlacementStatus_Namespace struct {
	State                PlacementStatus_State `protobuf:"varint,2,opt,name=state,proto3,enum=core.fed.solo.io.PlacementStatus_State" json:"state,omitempty"`
	Message              string                `protobuf:"bytes,3,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{}              `json:"-"`
	XXX_unrecognized     []byte                `json:"-"`
	XXX_sizecache        int32                 `json:"-"`
}

func (m *PlacementStatus_Namespace) Reset()         { *m = PlacementStatus_Namespace{} }
func (m *PlacementStatus_Namespace) String() string { return proto.CompactTextString(m) }
func (*PlacementStatus_Namespace) ProtoMessage()    {}
func (*PlacementStatus_Namespace) Descriptor() ([]byte, []int) {
	return fileDescriptor_09a22494930a46b3, []int{1, 0}
}
func (m *PlacementStatus_Namespace) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlacementStatus_Namespace.Unmarshal(m, b)
}
func (m *PlacementStatus_Namespace) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlacementStatus_Namespace.Marshal(b, m, deterministic)
}
func (m *PlacementStatus_Namespace) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlacementStatus_Namespace.Merge(m, src)
}
func (m *PlacementStatus_Namespace) XXX_Size() int {
	return xxx_messageInfo_PlacementStatus_Namespace.Size(m)
}
func (m *PlacementStatus_Namespace) XXX_DiscardUnknown() {
	xxx_messageInfo_PlacementStatus_Namespace.DiscardUnknown(m)
}

var xxx_messageInfo_PlacementStatus_Namespace proto.InternalMessageInfo

func (m *PlacementStatus_Namespace) GetState() PlacementStatus_State {
	if m != nil {
		return m.State
	}
	return PlacementStatus_UNKNOWN
}

func (m *PlacementStatus_Namespace) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

type PlacementStatus_Cluster struct {
	// map containing the name of the namespace, with the associated status.
	Namespaces           map[string]*PlacementStatus_Namespace `protobuf:"bytes,1,rep,name=namespaces,proto3" json:"namespaces,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}                              `json:"-"`
	XXX_unrecognized     []byte                                `json:"-"`
	XXX_sizecache        int32                                 `json:"-"`
}

func (m *PlacementStatus_Cluster) Reset()         { *m = PlacementStatus_Cluster{} }
func (m *PlacementStatus_Cluster) String() string { return proto.CompactTextString(m) }
func (*PlacementStatus_Cluster) ProtoMessage()    {}
func (*PlacementStatus_Cluster) Descriptor() ([]byte, []int) {
	return fileDescriptor_09a22494930a46b3, []int{1, 1}
}
func (m *PlacementStatus_Cluster) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PlacementStatus_Cluster.Unmarshal(m, b)
}
func (m *PlacementStatus_Cluster) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PlacementStatus_Cluster.Marshal(b, m, deterministic)
}
func (m *PlacementStatus_Cluster) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PlacementStatus_Cluster.Merge(m, src)
}
func (m *PlacementStatus_Cluster) XXX_Size() int {
	return xxx_messageInfo_PlacementStatus_Cluster.Size(m)
}
func (m *PlacementStatus_Cluster) XXX_DiscardUnknown() {
	xxx_messageInfo_PlacementStatus_Cluster.DiscardUnknown(m)
}

var xxx_messageInfo_PlacementStatus_Cluster proto.InternalMessageInfo

func (m *PlacementStatus_Cluster) GetNamespaces() map[string]*PlacementStatus_Namespace {
	if m != nil {
		return m.Namespaces
	}
	return nil
}

func init() {
	proto.RegisterEnum("core.fed.solo.io.PlacementStatus_State", PlacementStatus_State_name, PlacementStatus_State_value)
	proto.RegisterType((*TemplateMetadata)(nil), "core.fed.solo.io.TemplateMetadata")
	proto.RegisterMapType((map[string]string)(nil), "core.fed.solo.io.TemplateMetadata.AnnotationsEntry")
	proto.RegisterMapType((map[string]string)(nil), "core.fed.solo.io.TemplateMetadata.LabelsEntry")
	proto.RegisterType((*PlacementStatus)(nil), "core.fed.solo.io.PlacementStatus")
	proto.RegisterMapType((map[string]*PlacementStatus_Cluster)(nil), "core.fed.solo.io.PlacementStatus.ClustersEntry")
	proto.RegisterType((*PlacementStatus_Namespace)(nil), "core.fed.solo.io.PlacementStatus.Namespace")
	proto.RegisterType((*PlacementStatus_Cluster)(nil), "core.fed.solo.io.PlacementStatus.Cluster")
	proto.RegisterMapType((map[string]*PlacementStatus_Namespace)(nil), "core.fed.solo.io.PlacementStatus.Cluster.NamespacesEntry")
}

func init() {
	proto.RegisterFile("github.com/solo-io/solo-apis/api/gloo-fed/fed/core/v1/placement.proto", fileDescriptor_09a22494930a46b3)
}

var fileDescriptor_09a22494930a46b3 = []byte{
	// 585 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x54, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xc5, 0x49, 0xdc, 0x92, 0x89, 0xa0, 0xd6, 0xd2, 0x83, 0x65, 0x89, 0xaa, 0xea, 0x85, 0x22,
	0x54, 0x5b, 0xb4, 0x97, 0x16, 0x09, 0x90, 0xdb, 0xb8, 0x55, 0xd4, 0x60, 0x42, 0xda, 0x82, 0xe0,
	0x52, 0xad, 0x93, 0x89, 0x31, 0xb5, 0xbd, 0x96, 0x77, 0x93, 0x36, 0x7f, 0xc4, 0x27, 0x70, 0xe4,
	0x13, 0xf8, 0x06, 0xfe, 0x81, 0x13, 0x17, 0xe4, 0xb5, 0x1d, 0xdc, 0x50, 0xd1, 0x70, 0xe0, 0x60,
	0x79, 0x67, 0x67, 0xdf, 0x7b, 0xb3, 0x6f, 0x47, 0x03, 0x8e, 0x1f, 0x88, 0x8f, 0x63, 0xcf, 0x1c,
	0xb0, 0xc8, 0xe2, 0x2c, 0x64, 0x5b, 0x01, 0xcb, 0xff, 0x34, 0x09, 0xb8, 0x45, 0x93, 0xc0, 0xf2,
	0x43, 0xc6, 0xb6, 0x46, 0x38, 0xb4, 0xb2, 0x6f, 0xc0, 0x52, 0xb4, 0x26, 0x4f, 0xad, 0x24, 0xa4,
	0x03, 0x8c, 0x30, 0x16, 0x66, 0x92, 0x32, 0xc1, 0x88, 0x96, 0x25, 0xcc, 0x11, 0x0e, 0xcd, 0x0c,
	0x6c, 0x06, 0xcc, 0x58, 0xf5, 0x99, 0xcf, 0x64, 0xd2, 0xca, 0x56, 0xf9, 0x39, 0x83, 0xe0, 0x95,
	0xc8, 0x37, 0xf1, 0xaa, 0xc0, 0x1a, 0x6b, 0x3e, 0x63, 0x7e, 0x88, 0x96, 0x8c, 0xbc, 0xf1, 0xc8,
	0xba, 0x4c, 0x69, 0x92, 0x60, 0xca, 0xf3, 0xfc, 0xc6, 0xd7, 0x1a, 0x68, 0xa7, 0x18, 0x25, 0x21,
	0x15, 0xf8, 0x0a, 0x05, 0x1d, 0x52, 0x41, 0xc9, 0x19, 0xb4, 0x68, 0x1c, 0x33, 0x41, 0x45, 0xc0,
	0x62, 0xae, 0x2b, 0xeb, 0xf5, 0xcd, 0xd6, 0xf6, 0x8e, 0x39, 0x5f, 0x86, 0x39, 0x0f, 0x34, 0xed,
	0xdf, 0x28, 0x27, 0x16, 0xe9, 0xb4, 0x5f, 0xe5, 0x21, 0x87, 0xb0, 0x14, 0x52, 0x0f, 0x43, 0xae,
	0xd7, 0x24, 0xa3, 0xb9, 0x00, 0x63, 0x57, 0x02, 0x72, 0xb2, 0x02, 0x4d, 0x08, 0x34, 0x62, 0x1a,
	0xa1, 0x5e, 0x5f, 0x57, 0x36, 0x9b, 0x7d, 0xb9, 0x36, 0x5e, 0x80, 0x36, 0x2f, 0x4e, 0x34, 0xa8,
	0x5f, 0xe0, 0x54, 0x57, 0xe4, 0xb1, 0x6c, 0x49, 0x56, 0x41, 0x9d, 0xd0, 0x70, 0x8c, 0x7a, 0x4d,
	0xee, 0xe5, 0xc1, 0xb3, 0xda, 0xae, 0x62, 0xec, 0x41, 0xab, 0x22, 0xf5, 0x2f, 0xd0, 0x8d, 0x9f,
	0x2a, 0xac, 0xf4, 0xca, 0x27, 0x3b, 0x11, 0x54, 0x8c, 0x39, 0x39, 0x86, 0xbb, 0x83, 0x70, 0xcc,
	0x05, 0xa6, 0xa5, 0x7d, 0xd6, 0x9f, 0x97, 0x9d, 0x03, 0x99, 0x07, 0x05, 0x22, 0xbf, 0xed, 0x8c,
	0x80, 0x3c, 0x07, 0x95, 0x0b, 0x2a, 0x72, 0xe9, 0xfb, 0xdb, 0x8f, 0x6e, 0x67, 0xca, 0x7e, 0xd8,
	0xcf, 0x51, 0x44, 0x87, 0xe5, 0x08, 0x39, 0xa7, 0x7e, 0xe9, 0x58, 0x19, 0x12, 0x0b, 0x1e, 0x30,
	0x8f, 0x63, 0x3a, 0xc1, 0xe1, 0xb9, 0x8f, 0x31, 0xa6, 0xd2, 0x3d, 0xbd, 0xb1, 0xae, 0x6c, 0xd6,
	0xfb, 0xa4, 0x4c, 0x1d, 0xcd, 0x32, 0xe4, 0x21, 0xc0, 0x65, 0x1a, 0x08, 0x81, 0xf1, 0xb9, 0x37,
	0xd5, 0x55, 0xc9, 0xd6, 0x2c, 0x76, 0xf6, 0xa7, 0xc6, 0x10, 0x9a, 0x2e, 0x8d, 0x90, 0x27, 0x74,
	0x80, 0xff, 0xad, 0x6a, 0xe3, 0x9b, 0x02, 0xcb, 0x85, 0x55, 0xe4, 0x3d, 0x40, 0x5c, 0x2a, 0x96,
	0x4e, 0xef, 0x2d, 0xec, 0xb4, 0x39, 0xab, 0xb6, 0xf0, 0xbc, 0x42, 0x66, 0x7c, 0x82, 0x95, 0xb9,
	0xf4, 0x0d, 0x5d, 0x61, 0x57, 0xbb, 0xa2, 0xb5, 0xfd, 0xe4, 0x76, 0xe9, 0x19, 0x67, 0xb5, 0xfb,
	0x46, 0x70, 0xef, 0xda, 0xe3, 0xdf, 0xa0, 0xf4, 0xf2, 0xba, 0xd2, 0xe3, 0x85, 0x2f, 0x59, 0x6d,
	0xd5, 0x37, 0xa0, 0x4a, 0x93, 0x49, 0x0b, 0x96, 0xcf, 0xdc, 0x63, 0xf7, 0xf5, 0x3b, 0x57, 0xbb,
	0x43, 0x00, 0x96, 0x7a, 0x5d, 0xfb, 0xc0, 0x69, 0x6b, 0x4a, 0xb6, 0x3e, 0xb4, 0x3b, 0x5d, 0xa7,
	0xad, 0xd5, 0x48, 0x13, 0xd4, 0x93, 0x53, 0xbb, 0xeb, 0x68, 0xf5, 0xec, 0x7c, 0xc7, 0x7d, 0x6b,
	0x77, 0x3b, 0x6d, 0xad, 0x91, 0x05, 0x3d, 0xc7, 0x6d, 0x77, 0xdc, 0x23, 0x4d, 0xdd, 0xdf, 0xff,
	0xf2, 0xa3, 0xa1, 0x7c, 0xfe, 0xbe, 0xa6, 0x7c, 0xd8, 0xfd, 0xeb, 0xb4, 0x4b, 0x2e, 0x7c, 0x39,
	0xf1, 0x2a, 0x15, 0x97, 0x03, 0xcf, 0x5b, 0x92, 0xb3, 0x68, 0xe7, 0x57, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x4b, 0x49, 0xbe, 0x3b, 0x30, 0x05, 0x00, 0x00,
}

func (this *TemplateMetadata) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*TemplateMetadata)
	if !ok {
		that2, ok := that.(TemplateMetadata)
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
	if len(this.Annotations) != len(that1.Annotations) {
		return false
	}
	for i := range this.Annotations {
		if this.Annotations[i] != that1.Annotations[i] {
			return false
		}
	}
	if len(this.Labels) != len(that1.Labels) {
		return false
	}
	for i := range this.Labels {
		if this.Labels[i] != that1.Labels[i] {
			return false
		}
	}
	if this.Name != that1.Name {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *PlacementStatus) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*PlacementStatus)
	if !ok {
		that2, ok := that.(PlacementStatus)
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
	if len(this.Clusters) != len(that1.Clusters) {
		return false
	}
	for i := range this.Clusters {
		if !this.Clusters[i].Equal(that1.Clusters[i]) {
			return false
		}
	}
	if this.State != that1.State {
		return false
	}
	if this.Message != that1.Message {
		return false
	}
	if this.ObservedGeneration != that1.ObservedGeneration {
		return false
	}
	if this.WrittenBy != that1.WrittenBy {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *PlacementStatus_Namespace) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*PlacementStatus_Namespace)
	if !ok {
		that2, ok := that.(PlacementStatus_Namespace)
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
	if this.State != that1.State {
		return false
	}
	if this.Message != that1.Message {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *PlacementStatus_Cluster) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*PlacementStatus_Cluster)
	if !ok {
		that2, ok := that.(PlacementStatus_Cluster)
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
	if len(this.Namespaces) != len(that1.Namespaces) {
		return false
	}
	for i := range this.Namespaces {
		if !this.Namespaces[i].Equal(that1.Namespaces[i]) {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}