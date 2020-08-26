// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/v1/options/service_spec.proto

package options

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/solo-io/protoc-gen-ext/extproto"
	_ "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/options/grpc"
	_ "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/v1/options/rest"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for ServiceSpec
func (this *ServiceSpec) MarshalJSON() ([]byte, error) {
	str, err := ServiceSpecMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for ServiceSpec
func (this *ServiceSpec) UnmarshalJSON(b []byte) error {
	return ServiceSpecUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	ServiceSpecMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{EnumsAsInts: true}
	ServiceSpecUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)