// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// ServiceAccountSubject ServiceAccountSubject holds detailed information for service-account-kind subject.
//
// swagger:model ServiceAccountSubject
type ServiceAccountSubject struct {

	// `name` is the name of matching ServiceAccount objects, or "*" to match regardless of name. Required.
	// Required: true
	Name *string `json:"name"`

	// `namespace` is the namespace of matching ServiceAccount objects. Required.
	// Required: true
	Namespace *string `json:"namespace"`
}
