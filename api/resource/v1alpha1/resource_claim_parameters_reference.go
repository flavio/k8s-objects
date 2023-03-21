// Code generated by go-swagger; DO NOT EDIT.

package v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// ResourceClaimParametersReference ResourceClaimParametersReference contains enough information to let you locate the parameters for a ResourceClaim. The object must be in the same namespace as the ResourceClaim.
//
// swagger:model ResourceClaimParametersReference
type ResourceClaimParametersReference struct {

	// APIGroup is the group for the resource being referenced. It is empty for the core API. This matches the group in the APIVersion that is used when creating the resources.
	APIGroup string `json:"apiGroup,omitempty"`

	// Kind is the type of resource being referenced. This is the same value as in the parameter object's metadata, for example "ConfigMap".
	// Required: true
	Kind *string `json:"kind"`

	// Name is the name of resource being referenced.
	// Required: true
	Name *string `json:"name"`
}
