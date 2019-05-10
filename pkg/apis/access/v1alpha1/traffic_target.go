package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TrafficTarget associates a set of traffic definitions (rules) with a set of
// pods. With an accompanying IdentityBinding, access can be granted to traffic
// that matches the rules.
type TrafficTarget struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Most recently observed status of the object.
	// This data may not be up to date.
	// Populated by the system.
	// Read-only.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status
	// +optional
	Status Status `json:"status,omitempty" protobuf:"bytes,2,opt,name=status"`

	// Selector is the pod or group of pods to allow ingress traffic
	Selector TrafficTargetSelector `json:"selector,omitempty" protobuf:"bytes,3,opt,name=selector"`

	// Specs are the traffic specs to allow (HTTPRouteGroup | TCPRoute),
	Specs []TrafficTargetSpec `json:"specs,omitempty" protobuf:"bytes,4,opt,name=specs"`
}

// TrafficTargetSelector defines the pods to select for inbound traffic
type TrafficTargetSelector struct {
	// MatchLabels is a map of labels
	MatchLabels map[string]string `json:"matchLabels,omitempty" protobuf:"bytes,1,opt,name=matchLabels"`
}

// +genclient
// +genclient:noStatus
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ClusterTrafficTarget allows policy to be applied to targets that live in multiple namespaces.
// This is primarily useful to system wide endpoints such as metrics or health checks.
type ClusterTrafficTarget struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Most recently observed status of the object.
	// This data may not be up to date.
	// Populated by the system.
	// Read-only.
	// More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#spec-and-status
	// +optional
	Status Status `json:"status,omitempty" protobuf:"bytes,2,opt,name=status"`

	// Selector is the pod or group of pods to allow ingress traffic
	Selector []ClusterTargetSelector `json:"selector,omitempty" protobuf:"bytes,3,opt,name=selector"`

	// Specs are the traffic specs to allow (HTTPRouteGroup | TCPRoute),
	Specs []TrafficTargetSpec `json:"specs,omitempty" protobuf:"bytes,4,opt,name=specs"`
}

// ClusterTargetSelector defines a base type
type ClusterTargetSelector string

const (
	// ClusterTargetSelectorNotProtected is a target selector for pods which
	// are marketd with the !protected label
	ClusterTargetSelectorNotProtected ClusterTargetSelector = "!protected"
)

// TrafficTargetSpec is the traffic spec to allow for a TrafficTarget
type TrafficTargetSpec struct {
	// Kind is the kind of TrafficSpec to allow
	Kind string `json:"kind,omitempty" protobuf:"bytes,1,opt,name=kind"`
	// Name of the TrafficSpec to use
	Name string `json:"name,omitempty" protobuf:"bytes,2,opt,name=name"`
	// Namespace where the Spec is deployed
	Namespace string `json:"namespace,omitempty" protobuf:"bytes,3,opt,name=namespace"`
	// Matches is a list of TrafficSpec routes to allow traffic for
	Matches []string `json:"matches,omitempty" protobuf:"bytes,4,opt,name=matches"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// TrafficTargetList satisfy K8s code gen requirements
type TrafficTargetList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata"`

	Items []TrafficTarget `json:"items"`
}