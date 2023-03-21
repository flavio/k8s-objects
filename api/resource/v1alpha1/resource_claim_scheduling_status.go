// Code generated by go-swagger; DO NOT EDIT.

package v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// ResourceClaimSchedulingStatus ResourceClaimSchedulingStatus contains information about one particular ResourceClaim with "WaitForFirstConsumer" allocation mode.
//
// swagger:model ResourceClaimSchedulingStatus
type ResourceClaimSchedulingStatus struct {

	// Name matches the pod.spec.resourceClaims[*].Name field.
	Name string `json:"name,omitempty"`

	// UnsuitableNodes lists nodes that the ResourceClaim cannot be allocated for.
	//
	// The size of this field is limited to 128, the same as for PodSchedulingSpec.PotentialNodes. This may get increased in the future, but not reduced.
	UnsuitableNodes []string `json:"unsuitableNodes,omitempty"`
}
