// Code generated by go-swagger; DO NOT EDIT.

package v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// ServiceReference ServiceReference holds a reference to Service.legacy.k8s.io
//
// swagger:model ServiceReference
type ServiceReference struct {

	// `name` is the name of the service. Required
	// Required: true
	Name *string `json:"name"`

	// `namespace` is the namespace of the service. Required
	// Required: true
	Namespace *string `json:"namespace"`

	// `path` is an optional URL path which will be sent in any request to this service.
	Path string `json:"path,omitempty"`
}
