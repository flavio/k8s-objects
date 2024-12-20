// Code generated by go-swagger; DO NOT EDIT.

package v1alpha3

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	apimachinery_pkg_api_resource "github.com/kubewarden/k8s-objects/apimachinery/pkg/api/resource"
)

// BasicDevice BasicDevice defines one device instance.
//
// swagger:model BasicDevice
type BasicDevice struct {

	// Attributes defines the set of attributes for this device. The name of each attribute must be unique in that set.
	//
	// The maximum number of attributes and capacities combined is 32.
	Attributes map[string]*DeviceAttribute `json:"attributes,omitempty"`

	// Capacity defines the set of capacities for this device. The name of each capacity must be unique in that set.
	//
	// The maximum number of attributes and capacities combined is 32.
	Capacity map[string]*apimachinery_pkg_api_resource.Quantity `json:"capacity,omitempty"`
}
