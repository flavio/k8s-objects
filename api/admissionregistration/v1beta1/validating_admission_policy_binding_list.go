// Code generated by go-swagger; DO NOT EDIT.

package v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	apimachinery_pkg_apis_meta_v1 "github.com/kubewarden/k8s-objects/apimachinery/pkg/apis/meta/v1"
)

// ValidatingAdmissionPolicyBindingList ValidatingAdmissionPolicyBindingList is a list of ValidatingAdmissionPolicyBinding.
//
// swagger:model ValidatingAdmissionPolicyBindingList
type ValidatingAdmissionPolicyBindingList struct {

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	APIVersion string `json:"apiVersion,omitempty"`

	// List of PolicyBinding.
	// Required: true
	Items []*ValidatingAdmissionPolicyBinding `json:"items"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind string `json:"kind,omitempty"`

	// Standard list metadata. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Metadata *apimachinery_pkg_apis_meta_v1.ListMeta `json:"metadata,omitempty"`
}