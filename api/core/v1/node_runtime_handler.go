// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// NodeRuntimeHandler NodeRuntimeHandler is a set of runtime handler information.
//
// swagger:model NodeRuntimeHandler
type NodeRuntimeHandler struct {

	// Supported features.
	Features *NodeRuntimeHandlerFeatures `json:"features,omitempty"`

	// Runtime handler name. Empty for the default runtime handler.
	Name string `json:"name,omitempty"`
}