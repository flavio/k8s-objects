// Code generated by go-swagger; DO NOT EDIT.

package v1alpha2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// ResourceHandle ResourceHandle holds opaque resource data for processing by a specific kubelet plugin.
//
// swagger:model ResourceHandle
type ResourceHandle struct {

	// Data contains the opaque data associated with this ResourceHandle. It is set by the controller component of the resource driver whose name matches the DriverName set in the ResourceClaimStatus this ResourceHandle is embedded in. It is set at allocation time and is intended for processing by the kubelet plugin whose name matches the DriverName set in this ResourceHandle.
	//
	// The maximum size of this field is 16KiB. This may get increased in the future, but not reduced.
	Data string `json:"data,omitempty"`

	// DriverName specifies the name of the resource driver whose kubelet plugin should be invoked to process this ResourceHandle's data once it lands on a node. This may differ from the DriverName set in ResourceClaimStatus this ResourceHandle is embedded in.
	DriverName string `json:"driverName,omitempty"`

	// If StructuredData is set, then it needs to be used instead of Data.
	StructuredData *StructuredResourceHandle `json:"structuredData,omitempty"`
}
