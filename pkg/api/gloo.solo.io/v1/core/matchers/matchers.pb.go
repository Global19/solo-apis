// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/solo-io/solo-apis/api/gloo/gloo/v1/core/matchers/matchers.proto

package matchers

import (
	bytes "bytes"
	fmt "fmt"
	math "math"

	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// Parameters for matching routes to requests received by a Gloo-managed proxy
type Matcher struct {
	// Types that are valid to be assigned to PathSpecifier:
	//	*Matcher_Prefix
	//	*Matcher_Exact
	//	*Matcher_Regex
	PathSpecifier isMatcher_PathSpecifier `protobuf_oneof:"path_specifier"`
	// Specifies a set of headers that the route should match on. The router will
	// check the request’s headers against all the specified headers in the route
	// config. A match will happen if all the headers in the route are present in
	// the request with the same values (or based on presence if the value field
	// is not in the config).
	Headers []*HeaderMatcher `protobuf:"bytes,6,rep,name=headers,proto3" json:"headers,omitempty"`
	// Specifies a set of URL query parameters on which the route should
	// match. The router will check the query string from the *path* header
	// against all the specified query parameters. If the number of specified
	// query parameters is nonzero, they all must match the *path* header's
	// query string for a match to occur.
	QueryParameters []*QueryParameterMatcher `protobuf:"bytes,7,rep,name=query_parameters,json=queryParameters,proto3" json:"query_parameters,omitempty"`
	// HTTP Method/Verb(s) to match on. If none specified, the matcher will ignore the HTTP Method
	Methods              []string `protobuf:"bytes,8,rep,name=methods,proto3" json:"methods,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Matcher) Reset()         { *m = Matcher{} }
func (m *Matcher) String() string { return proto.CompactTextString(m) }
func (*Matcher) ProtoMessage()    {}
func (*Matcher) Descriptor() ([]byte, []int) {
	return fileDescriptor_eeb017f5b1d6b8a7, []int{0}
}
func (m *Matcher) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Matcher.Unmarshal(m, b)
}
func (m *Matcher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Matcher.Marshal(b, m, deterministic)
}
func (m *Matcher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Matcher.Merge(m, src)
}
func (m *Matcher) XXX_Size() int {
	return xxx_messageInfo_Matcher.Size(m)
}
func (m *Matcher) XXX_DiscardUnknown() {
	xxx_messageInfo_Matcher.DiscardUnknown(m)
}

var xxx_messageInfo_Matcher proto.InternalMessageInfo

type isMatcher_PathSpecifier interface {
	isMatcher_PathSpecifier()
	Equal(interface{}) bool
}

type Matcher_Prefix struct {
	Prefix string `protobuf:"bytes,1,opt,name=prefix,proto3,oneof" json:"prefix,omitempty"`
}
type Matcher_Exact struct {
	Exact string `protobuf:"bytes,2,opt,name=exact,proto3,oneof" json:"exact,omitempty"`
}
type Matcher_Regex struct {
	Regex string `protobuf:"bytes,3,opt,name=regex,proto3,oneof" json:"regex,omitempty"`
}

func (*Matcher_Prefix) isMatcher_PathSpecifier() {}
func (*Matcher_Exact) isMatcher_PathSpecifier()  {}
func (*Matcher_Regex) isMatcher_PathSpecifier()  {}

func (m *Matcher) GetPathSpecifier() isMatcher_PathSpecifier {
	if m != nil {
		return m.PathSpecifier
	}
	return nil
}

func (m *Matcher) GetPrefix() string {
	if x, ok := m.GetPathSpecifier().(*Matcher_Prefix); ok {
		return x.Prefix
	}
	return ""
}

func (m *Matcher) GetExact() string {
	if x, ok := m.GetPathSpecifier().(*Matcher_Exact); ok {
		return x.Exact
	}
	return ""
}

func (m *Matcher) GetRegex() string {
	if x, ok := m.GetPathSpecifier().(*Matcher_Regex); ok {
		return x.Regex
	}
	return ""
}

func (m *Matcher) GetHeaders() []*HeaderMatcher {
	if m != nil {
		return m.Headers
	}
	return nil
}

func (m *Matcher) GetQueryParameters() []*QueryParameterMatcher {
	if m != nil {
		return m.QueryParameters
	}
	return nil
}

func (m *Matcher) GetMethods() []string {
	if m != nil {
		return m.Methods
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Matcher) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Matcher_Prefix)(nil),
		(*Matcher_Exact)(nil),
		(*Matcher_Regex)(nil),
	}
}

// Internally, Gloo always uses the HTTP/2 *:authority* header to represent the HTTP/1 *Host* header.
// Thus, if attempting to match on *Host*, match on *:authority* instead.
type HeaderMatcher struct {
	// Specifies the name of the header in the request.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Specifies the value of the header. If the value is absent a request that
	// has the name header will match, regardless of the header’s value.
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	// Specifies whether the header value should be treated as regex or not.
	Regex bool `protobuf:"varint,3,opt,name=regex,proto3" json:"regex,omitempty"`
	// If set to true, the result of the match will be inverted. Defaults to false.
	//
	// Examples:
	// * name=foo, invert_match=true: matches if no header named `foo` is present
	// * name=foo, value=bar, invert_match=true: matches if no header named `foo` with value `bar` is present
	// * name=foo, value=``\d{3}``, regex=true, invert_match=true: matches if no header named `foo` with a value consisting of three integers is present
	InvertMatch          bool     `protobuf:"varint,4,opt,name=invert_match,json=invertMatch,proto3" json:"invert_match,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HeaderMatcher) Reset()         { *m = HeaderMatcher{} }
func (m *HeaderMatcher) String() string { return proto.CompactTextString(m) }
func (*HeaderMatcher) ProtoMessage()    {}
func (*HeaderMatcher) Descriptor() ([]byte, []int) {
	return fileDescriptor_eeb017f5b1d6b8a7, []int{1}
}
func (m *HeaderMatcher) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HeaderMatcher.Unmarshal(m, b)
}
func (m *HeaderMatcher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HeaderMatcher.Marshal(b, m, deterministic)
}
func (m *HeaderMatcher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HeaderMatcher.Merge(m, src)
}
func (m *HeaderMatcher) XXX_Size() int {
	return xxx_messageInfo_HeaderMatcher.Size(m)
}
func (m *HeaderMatcher) XXX_DiscardUnknown() {
	xxx_messageInfo_HeaderMatcher.DiscardUnknown(m)
}

var xxx_messageInfo_HeaderMatcher proto.InternalMessageInfo

func (m *HeaderMatcher) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *HeaderMatcher) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *HeaderMatcher) GetRegex() bool {
	if m != nil {
		return m.Regex
	}
	return false
}

func (m *HeaderMatcher) GetInvertMatch() bool {
	if m != nil {
		return m.InvertMatch
	}
	return false
}

// Query parameter matching treats the query string of a request's :path header
// as an ampersand-separated list of keys and/or key=value elements.
type QueryParameterMatcher struct {
	// Specifies the name of a key that must be present in the requested
	// *path*'s query string.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Specifies the value of the key. If the value is absent, a request
	// that contains the key in its query string will match, whether the
	// key appears with a value (e.g., "?debug=true") or not (e.g., "?debug")
	Value string `protobuf:"bytes,2,opt,name=value,proto3" json:"value,omitempty"`
	// Specifies whether the query parameter value is a regular expression.
	// Defaults to false. The entire query parameter value (i.e., the part to
	// the right of the equals sign in "key=value") must match the regex.
	// E.g., the regex "\d+$" will match "123" but not "a123" or "123a".
	Regex                bool     `protobuf:"varint,3,opt,name=regex,proto3" json:"regex,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *QueryParameterMatcher) Reset()         { *m = QueryParameterMatcher{} }
func (m *QueryParameterMatcher) String() string { return proto.CompactTextString(m) }
func (*QueryParameterMatcher) ProtoMessage()    {}
func (*QueryParameterMatcher) Descriptor() ([]byte, []int) {
	return fileDescriptor_eeb017f5b1d6b8a7, []int{2}
}
func (m *QueryParameterMatcher) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_QueryParameterMatcher.Unmarshal(m, b)
}
func (m *QueryParameterMatcher) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_QueryParameterMatcher.Marshal(b, m, deterministic)
}
func (m *QueryParameterMatcher) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryParameterMatcher.Merge(m, src)
}
func (m *QueryParameterMatcher) XXX_Size() int {
	return xxx_messageInfo_QueryParameterMatcher.Size(m)
}
func (m *QueryParameterMatcher) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryParameterMatcher.DiscardUnknown(m)
}

var xxx_messageInfo_QueryParameterMatcher proto.InternalMessageInfo

func (m *QueryParameterMatcher) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *QueryParameterMatcher) GetValue() string {
	if m != nil {
		return m.Value
	}
	return ""
}

func (m *QueryParameterMatcher) GetRegex() bool {
	if m != nil {
		return m.Regex
	}
	return false
}

func init() {
	proto.RegisterType((*Matcher)(nil), "matchers.core.gloo.solo.io.Matcher")
	proto.RegisterType((*HeaderMatcher)(nil), "matchers.core.gloo.solo.io.HeaderMatcher")
	proto.RegisterType((*QueryParameterMatcher)(nil), "matchers.core.gloo.solo.io.QueryParameterMatcher")
}

func init() {
	proto.RegisterFile("github.com/solo-io/solo-apis/api/gloo/gloo/v1/core/matchers/matchers.proto", fileDescriptor_eeb017f5b1d6b8a7)
}

var fileDescriptor_eeb017f5b1d6b8a7 = []byte{
	// 375 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x52, 0xcd, 0x4e, 0x32, 0x31,
	0x14, 0xfd, 0x86, 0xbf, 0x81, 0xf2, 0xfd, 0x90, 0x86, 0xcf, 0x34, 0x2c, 0x0c, 0xb2, 0xc2, 0x85,
	0x33, 0x41, 0xdf, 0x00, 0x37, 0xc4, 0x68, 0xa2, 0xb3, 0x31, 0x31, 0x26, 0xa4, 0x0c, 0x97, 0x99,
	0x46, 0x86, 0x96, 0x4e, 0x21, 0xe3, 0xeb, 0xb8, 0xf2, 0x11, 0x7c, 0x1e, 0xdf, 0xc1, 0xbd, 0x69,
	0x4b, 0x31, 0x44, 0x65, 0xe3, 0x66, 0xe6, 0x9e, 0xd3, 0x7b, 0x4e, 0x6f, 0x4f, 0x2e, 0xba, 0x48,
	0x98, 0x4a, 0x57, 0x93, 0x20, 0xe6, 0x59, 0x98, 0xf3, 0x39, 0x3f, 0x61, 0xdc, 0xfe, 0xa9, 0x60,
	0x79, 0x48, 0x05, 0x0b, 0x93, 0x39, 0xe7, 0xf6, 0xb3, 0x1e, 0x84, 0x31, 0x97, 0x10, 0x66, 0x54,
	0xc5, 0x29, 0xc8, 0x7c, 0x5b, 0x04, 0x42, 0x72, 0xc5, 0x71, 0x67, 0x8b, 0x75, 0x5b, 0xa0, 0x35,
	0x81, 0xf6, 0x0a, 0x18, 0xef, 0xb4, 0x13, 0x9e, 0x70, 0xd3, 0x16, 0xea, 0xca, 0x2a, 0x3a, 0x18,
	0x0a, 0x65, 0x49, 0x28, 0x94, 0xe5, 0x7a, 0x4f, 0x25, 0xe4, 0x5f, 0x59, 0x23, 0x4c, 0x50, 0x4d,
	0x48, 0x98, 0xb1, 0x82, 0x78, 0x5d, 0xaf, 0xdf, 0x18, 0xfd, 0x8a, 0x36, 0x18, 0x1f, 0xa0, 0x2a,
	0x14, 0x34, 0x56, 0xa4, 0xb4, 0x39, 0xb0, 0x50, 0xf3, 0x12, 0x12, 0x28, 0x48, 0xd9, 0xf1, 0x06,
	0xe2, 0x73, 0xe4, 0xa7, 0x40, 0xa7, 0x20, 0x73, 0x52, 0xeb, 0x96, 0xfb, 0xcd, 0xd3, 0xe3, 0xe0,
	0xfb, 0x69, 0x83, 0x91, 0x69, 0xdd, 0x4c, 0x11, 0x39, 0x25, 0xbe, 0x47, 0xad, 0xe5, 0x0a, 0xe4,
	0xe3, 0x58, 0x50, 0x49, 0x33, 0x50, 0xda, 0xcd, 0x37, 0x6e, 0x83, 0x7d, 0x6e, 0x37, 0x5a, 0x73,
	0xed, 0x24, 0xce, 0xf5, 0xdf, 0x72, 0x87, 0xce, 0x31, 0x41, 0x7e, 0x06, 0x2a, 0xe5, 0xd3, 0x9c,
	0xd4, 0xbb, 0xe5, 0x7e, 0x23, 0x72, 0x70, 0xd8, 0x42, 0x7f, 0x05, 0x55, 0xe9, 0x38, 0x17, 0x10,
	0xb3, 0x19, 0x03, 0xd9, 0x93, 0xe8, 0xcf, 0xce, 0x8c, 0x18, 0xa3, 0xca, 0x82, 0x66, 0x60, 0x73,
	0x8a, 0x4c, 0x8d, 0xdb, 0xa8, 0xba, 0xa6, 0xf3, 0x15, 0xd8, 0x8c, 0x22, 0x0b, 0x34, 0xfb, 0x91,
	0x50, 0xdd, 0xe5, 0x73, 0x84, 0x7e, 0xb3, 0xc5, 0x1a, 0xa4, 0x1a, 0x9b, 0x87, 0x90, 0x8a, 0x39,
	0x6c, 0x5a, 0xce, 0x5c, 0xd2, 0xbb, 0x45, 0xff, 0xbf, 0x7c, 0xc9, 0x4f, 0xef, 0x1e, 0x5e, 0xbe,
	0xbc, 0x55, 0xbc, 0xe7, 0xd7, 0x43, 0xef, 0x6e, 0xb8, 0x77, 0x1b, 0xc5, 0x43, 0xb2, 0xdd, 0x48,
	0x17, 0xee, 0xa7, 0xa5, 0x9c, 0xd4, 0xcc, 0x1a, 0x9d, 0xbd, 0x07, 0x00, 0x00, 0xff, 0xff, 0xcc,
	0x4d, 0x27, 0x26, 0xda, 0x02, 0x00, 0x00,
}

func (this *Matcher) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Matcher)
	if !ok {
		that2, ok := that.(Matcher)
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
	if that1.PathSpecifier == nil {
		if this.PathSpecifier != nil {
			return false
		}
	} else if this.PathSpecifier == nil {
		return false
	} else if !this.PathSpecifier.Equal(that1.PathSpecifier) {
		return false
	}
	if len(this.Headers) != len(that1.Headers) {
		return false
	}
	for i := range this.Headers {
		if !this.Headers[i].Equal(that1.Headers[i]) {
			return false
		}
	}
	if len(this.QueryParameters) != len(that1.QueryParameters) {
		return false
	}
	for i := range this.QueryParameters {
		if !this.QueryParameters[i].Equal(that1.QueryParameters[i]) {
			return false
		}
	}
	if len(this.Methods) != len(that1.Methods) {
		return false
	}
	for i := range this.Methods {
		if this.Methods[i] != that1.Methods[i] {
			return false
		}
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *Matcher_Prefix) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Matcher_Prefix)
	if !ok {
		that2, ok := that.(Matcher_Prefix)
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
	if this.Prefix != that1.Prefix {
		return false
	}
	return true
}
func (this *Matcher_Exact) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Matcher_Exact)
	if !ok {
		that2, ok := that.(Matcher_Exact)
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
	if this.Exact != that1.Exact {
		return false
	}
	return true
}
func (this *Matcher_Regex) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*Matcher_Regex)
	if !ok {
		that2, ok := that.(Matcher_Regex)
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
	if this.Regex != that1.Regex {
		return false
	}
	return true
}
func (this *HeaderMatcher) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*HeaderMatcher)
	if !ok {
		that2, ok := that.(HeaderMatcher)
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
	if this.Name != that1.Name {
		return false
	}
	if this.Value != that1.Value {
		return false
	}
	if this.Regex != that1.Regex {
		return false
	}
	if this.InvertMatch != that1.InvertMatch {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}
func (this *QueryParameterMatcher) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*QueryParameterMatcher)
	if !ok {
		that2, ok := that.(QueryParameterMatcher)
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
	if this.Name != that1.Name {
		return false
	}
	if this.Value != that1.Value {
		return false
	}
	if this.Regex != that1.Regex {
		return false
	}
	if !bytes.Equal(this.XXX_unrecognized, that1.XXX_unrecognized) {
		return false
	}
	return true
}