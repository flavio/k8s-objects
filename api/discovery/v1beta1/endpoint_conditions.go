// Code generated by go-swagger; DO NOT EDIT.

package v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// EndpointConditions EndpointConditions represents the current condition of an endpoint.
//
// swagger:model EndpointConditions
type EndpointConditions struct {

	// ready indicates that this endpoint is prepared to receive traffic, according to whatever system is managing the endpoint. A nil value indicates an unknown state. In most cases consumers should interpret this unknown state as ready.
	Ready bool `json:"ready,omitempty"`
}
