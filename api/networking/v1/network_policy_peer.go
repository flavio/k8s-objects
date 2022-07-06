// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	apimachinery_pkg_apis_meta_v1 "github.com/kubewarden/k8s-objects/apimachinery/pkg/apis/meta/v1"
)

// NetworkPolicyPeer NetworkPolicyPeer describes a peer to allow traffic from. Only certain combinations of fields are allowed
//
// swagger:model NetworkPolicyPeer
type NetworkPolicyPeer struct {

	// IPBlock defines policy on a particular IPBlock. If this field is set then neither of the other fields can be.
	IPBlock *IPBlock `json:"ipBlock,omitempty"`

	// Selects Namespaces using cluster-scoped labels. This field follows standard label selector semantics; if present but empty, it selects all namespaces.
	//
	// If PodSelector is also set, then the NetworkPolicyPeer as a whole selects the Pods matching PodSelector in the Namespaces selected by NamespaceSelector. Otherwise it selects all Pods in the Namespaces selected by NamespaceSelector.
	NamespaceSelector *apimachinery_pkg_apis_meta_v1.LabelSelector `json:"namespaceSelector,omitempty"`

	// This is a label selector which selects Pods. This field follows standard label selector semantics; if present but empty, it selects all pods.
	//
	// If NamespaceSelector is also set, then the NetworkPolicyPeer as a whole selects the Pods matching PodSelector in the Namespaces selected by NamespaceSelector. Otherwise it selects the Pods matching PodSelector in the policy's own Namespace.
	PodSelector *apimachinery_pkg_apis_meta_v1.LabelSelector `json:"podSelector,omitempty"`
}
