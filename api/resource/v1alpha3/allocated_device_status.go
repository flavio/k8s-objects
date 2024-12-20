// Code generated by go-swagger; DO NOT EDIT.

package v1alpha3

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	jsonext "encoding/json"

	apimachinery_pkg_apis_meta_v1 "github.com/kubewarden/k8s-objects/apimachinery/pkg/apis/meta/v1"
)

// AllocatedDeviceStatus AllocatedDeviceStatus contains the status of an allocated device, if the driver chooses to report it. This may include driver-specific information.
//
// swagger:model AllocatedDeviceStatus
type AllocatedDeviceStatus struct {

	// Conditions contains the latest observation of the device's state. If the device has been configured according to the class and claim config references, the `Ready` condition should be True.
	Conditions []*apimachinery_pkg_apis_meta_v1.Condition `json:"conditions,omitempty"`

	// Data contains arbitrary driver-specific data.
	//
	// The length of the raw data must be smaller or equal to 10 Ki.
	Data jsonext.RawMessage `json:"data,omitempty"`

	// Device references one device instance via its name in the driver's resource pool. It must be a DNS label.
	// Required: true
	Device *string `json:"device"`

	// Driver specifies the name of the DRA driver whose kubelet plugin should be invoked to process the allocation once the claim is needed on a node.
	//
	// Must be a DNS subdomain and should end with a DNS domain owned by the vendor of the driver.
	// Required: true
	Driver *string `json:"driver"`

	// NetworkData contains network-related information specific to the device.
	NetworkData *NetworkDeviceData `json:"networkData,omitempty"`

	// This name together with the driver name and the device name field identify which device was allocated (`<driver name>/<pool name>/<device name>`).
	//
	// Must not be longer than 253 characters and may contain one or more DNS sub-domains separated by slashes.
	// Required: true
	Pool *string `json:"pool"`
}
