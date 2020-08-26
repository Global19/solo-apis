// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/v1/options/aws/ec2/aws_ec2.proto

package ec2

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	_ "github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for UpstreamSpec
func (this *UpstreamSpec) MarshalJSON() ([]byte, error) {
	str, err := AwsEc2Marshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for UpstreamSpec
func (this *UpstreamSpec) UnmarshalJSON(b []byte) error {
	return AwsEc2Unmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for TagFilter
func (this *TagFilter) MarshalJSON() ([]byte, error) {
	str, err := AwsEc2Marshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for TagFilter
func (this *TagFilter) UnmarshalJSON(b []byte) error {
	return AwsEc2Unmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for TagFilter_KvPair
func (this *TagFilter_KvPair) MarshalJSON() ([]byte, error) {
	str, err := AwsEc2Marshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for TagFilter_KvPair
func (this *TagFilter_KvPair) UnmarshalJSON(b []byte) error {
	return AwsEc2Unmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	AwsEc2Marshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{EnumsAsInts: true}
	AwsEc2Unmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)