// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	apimachinery_pkg_apis_meta_v1 "github.com/kubewarden/k8s-objects/apimachinery/pkg/apis/meta/v1"
)

// PodDisruptionBudgetStatus PodDisruptionBudgetStatus represents information about the status of a PodDisruptionBudget. Status may trail the actual state of a system.
//
// swagger:model PodDisruptionBudgetStatus
type PodDisruptionBudgetStatus struct {

	// Conditions contain conditions for PDB. The disruption controller sets the DisruptionAllowed condition. The following are known values for the reason field (additional reasons could be added in the future): - SyncFailed: The controller encountered an error and wasn't able to compute
	//               the number of allowed disruptions. Therefore no disruptions are
	//               allowed and the status of the condition will be False.
	// - InsufficientPods: The number of pods are either at or below the number
	//                     required by the PodDisruptionBudget. No disruptions are
	//                     allowed and the status of the condition will be False.
	// - SufficientPods: There are more pods than required by the PodDisruptionBudget.
	//                   The condition will be True, and the number of allowed
	//                   disruptions are provided by the disruptionsAllowed property.
	Conditions []apimachinery_pkg_apis_meta_v1.Condition `json:"conditions"`

	// current number of healthy pods
	// Required: true
	CurrentHealthy *int32 `json:"currentHealthy"`

	// minimum desired number of healthy pods
	// Required: true
	DesiredHealthy *int32 `json:"desiredHealthy"`

	// DisruptedPods contains information about pods whose eviction was processed by the API server eviction subresource handler but has not yet been observed by the PodDisruptionBudget controller. A pod will be in this map from the time when the API server processed the eviction request to the time when the pod is seen by PDB controller as having been marked for deletion (or after a timeout). The key in the map is the name of the pod and the value is the time when the API server processed the eviction request. If the deletion didn't occur and a pod is still there it will be removed from the list automatically by PodDisruptionBudget controller after some time. If everything goes smooth this map should be empty for the most of the time. Large number of entries in the map may indicate problems with pod deletions.
	DisruptedPods map[string]apimachinery_pkg_apis_meta_v1.Time `json:"disruptedPods,omitempty"`

	// Number of pod disruptions that are currently allowed.
	// Required: true
	DisruptionsAllowed *int32 `json:"disruptionsAllowed"`

	// total number of pods counted by this disruption budget
	// Required: true
	ExpectedPods *int32 `json:"expectedPods"`

	// Most recent generation observed when updating this PDB status. DisruptionsAllowed and other status information is valid only if observedGeneration equals to PDB's object generation.
	ObservedGeneration int64 `json:"observedGeneration,omitempty"`
}