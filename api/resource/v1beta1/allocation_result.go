// Code generated by go-swagger; DO NOT EDIT.

package v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	api_core_v1 "github.com/kubewarden/k8s-objects/api/core/v1"
)

// AllocationResult AllocationResult contains attributes of an allocated resource.
//
// swagger:model AllocationResult
type AllocationResult struct {

	// Devices is the result of allocating devices.
	Devices *DeviceAllocationResult `json:"devices,omitempty"`

	// NodeSelector defines where the allocated resources are available. If unset, they are available everywhere.
	NodeSelector *api_core_v1.NodeSelector `json:"nodeSelector,omitempty"`
}