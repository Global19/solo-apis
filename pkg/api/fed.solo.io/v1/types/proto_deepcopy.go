// Code generated by skv2. DO NOT EDIT.

// This file contains generated Deepcopy methods for proto-based Spec and Status fields

package types

import (
	proto "github.com/gogo/protobuf/proto"
)

// DeepCopyInto for the GlooInstance.Spec
func (in *GlooInstanceSpec) DeepCopyInto(out *GlooInstanceSpec) {
	p := proto.Clone(in).(*GlooInstanceSpec)
	*out = *p
}

// DeepCopyInto for the GlooInstance.Status
func (in *GlooInstanceStatus) DeepCopyInto(out *GlooInstanceStatus) {
	p := proto.Clone(in).(*GlooInstanceStatus)
	*out = *p
}