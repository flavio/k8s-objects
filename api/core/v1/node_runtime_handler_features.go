// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// NodeRuntimeHandlerFeatures NodeRuntimeHandlerFeatures is a set of features implemented by the runtime handler.
//
// swagger:model NodeRuntimeHandlerFeatures
type NodeRuntimeHandlerFeatures struct {

	// RecursiveReadOnlyMounts is set to true if the runtime handler supports RecursiveReadOnlyMounts.
	RecursiveReadOnlyMounts bool `json:"recursiveReadOnlyMounts,omitempty"`

	// UserNamespaces is set to true if the runtime handler supports UserNamespaces, including for volumes.
	UserNamespaces bool `json:"userNamespaces,omitempty"`
}