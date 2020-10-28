// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/external/envoy/extensions/wasm/v3/wasm.proto

package v3

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/envoyproxy/protoc-gen-validate/validate"
	_ "github.com/gogo/protobuf/gogoproto"
	github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/gogo/protobuf/types"
	_ "github.com/solo-io/solo-apis/pkg/api/gloo.solo.io/external/envoy/config/core/v3"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler for VmConfig
func (this *VmConfig) MarshalJSON() ([]byte, error) {
	str, err := WasmMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for VmConfig
func (this *VmConfig) UnmarshalJSON(b []byte) error {
	return WasmUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for PluginConfig
func (this *PluginConfig) MarshalJSON() ([]byte, error) {
	str, err := WasmMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for PluginConfig
func (this *PluginConfig) UnmarshalJSON(b []byte) error {
	return WasmUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

// MarshalJSON is a custom marshaler for WasmService
func (this *WasmService) MarshalJSON() ([]byte, error) {
	str, err := WasmMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler for WasmService
func (this *WasmService) UnmarshalJSON(b []byte) error {
	return WasmUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	WasmMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	WasmUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)
