// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// ParamKind ParamKind is a tuple of Group Kind and Version.
//
// swagger:model ParamKind
type ParamKind struct {

	// APIVersion is the API group version the resources belong to. In format of "group/version". Required.
	APIVersion string `json:"apiVersion,omitempty"`

	// Kind is the API kind the resources belong to. Required.
	Kind string `json:"kind,omitempty"`
}
