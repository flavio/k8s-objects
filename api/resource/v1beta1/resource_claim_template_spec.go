// Code generated by go-swagger; DO NOT EDIT.

package v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	apimachinery_pkg_apis_meta_v1 "github.com/kubewarden/k8s-objects/apimachinery/pkg/apis/meta/v1"
)

// ResourceClaimTemplateSpec ResourceClaimTemplateSpec contains the metadata and fields for a ResourceClaim.
//
// swagger:model ResourceClaimTemplateSpec
type ResourceClaimTemplateSpec struct {

	// ObjectMeta may contain labels and annotations that will be copied into the ResourceClaim when creating it. No other fields are allowed and will be rejected during validation.
	Metadata *apimachinery_pkg_apis_meta_v1.ObjectMeta `json:"metadata,omitempty"`

	// Spec for the ResourceClaim. The entire content is copied unchanged into the ResourceClaim that gets created from this template. The same fields as in a ResourceClaim are also valid here.
	// Required: true
	Spec *ResourceClaimSpec `json:"spec"`
}
