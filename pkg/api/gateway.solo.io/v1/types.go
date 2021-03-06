// Code generated by skv2. DO NOT EDIT.

// Definitions for the Kubernetes types
package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status

// GroupVersionKind for Gateway
var GatewayGVK = schema.GroupVersionKind{
	Group:   "gateway.solo.io",
	Version: "v1",
	Kind:    "Gateway",
}

// Gateway is the Schema for the gateway API
type Gateway struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   GatewaySpec   `json:"spec,omitempty"`
	Status GatewayStatus `json:"status,omitempty"`
}

// GVK returns the GroupVersionKind associated with the resource type.
func (Gateway) GVK() schema.GroupVersionKind {
	return GatewayGVK
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// GatewayList contains a list of Gateway
type GatewayList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Gateway `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status

// GroupVersionKind for RouteTable
var RouteTableGVK = schema.GroupVersionKind{
	Group:   "gateway.solo.io",
	Version: "v1",
	Kind:    "RouteTable",
}

// RouteTable is the Schema for the routeTable API
type RouteTable struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   RouteTableSpec   `json:"spec,omitempty"`
	Status RouteTableStatus `json:"status,omitempty"`
}

// GVK returns the GroupVersionKind associated with the resource type.
func (RouteTable) GVK() schema.GroupVersionKind {
	return RouteTableGVK
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// RouteTableList contains a list of RouteTable
type RouteTableList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RouteTable `json:"items"`
}

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:openapi-gen=true
// +kubebuilder:subresource:status

// GroupVersionKind for VirtualService
var VirtualServiceGVK = schema.GroupVersionKind{
	Group:   "gateway.solo.io",
	Version: "v1",
	Kind:    "VirtualService",
}

// VirtualService is the Schema for the virtualService API
type VirtualService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VirtualServiceSpec   `json:"spec,omitempty"`
	Status VirtualServiceStatus `json:"status,omitempty"`
}

// GVK returns the GroupVersionKind associated with the resource type.
func (VirtualService) GVK() schema.GroupVersionKind {
	return VirtualServiceGVK
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// VirtualServiceList contains a list of VirtualService
type VirtualServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualService `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Gateway{}, &GatewayList{})
	SchemeBuilder.Register(&RouteTable{}, &RouteTableList{})
	SchemeBuilder.Register(&VirtualService{}, &VirtualServiceList{})
}
