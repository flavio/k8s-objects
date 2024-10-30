// Code generated by go-swagger; DO NOT EDIT.

package v1alpha3

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// PodSchedulingContextSpec PodSchedulingContextSpec describes where resources for the Pod are needed.
//
// swagger:model PodSchedulingContextSpec
type PodSchedulingContextSpec struct {

	// PotentialNodes lists nodes where the Pod might be able to run.
	//
	// The size of this field is limited to 128. This is large enough for many clusters. Larger clusters may need more attempts to find a node that suits all pending resources. This may get increased in the future, but not reduced.
	PotentialNodes []string `json:"potentialNodes,omitempty"`

	// SelectedNode is the node for which allocation of ResourceClaims that are referenced by the Pod and that use "WaitForFirstConsumer" allocation is to be attempted.
	SelectedNode string `json:"selectedNode,omitempty"`
}
