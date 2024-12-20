// Code generated by go-swagger; DO NOT EDIT.

package v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	apimachinery_pkg_api_resource "github.com/kubewarden/k8s-objects/apimachinery/pkg/api/resource"
)

// DeviceCapacity DeviceCapacity describes a quantity associated with a device.
//
// swagger:model DeviceCapacity
type DeviceCapacity struct {

	// Value defines how much of a certain device capacity is available.
	// Required: true
	Value *apimachinery_pkg_api_resource.Quantity `json:"value"`
}
