// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// SelectableField SelectableField specifies the JSON path of a field that may be used with field selectors.
//
// swagger:model SelectableField
type SelectableField struct {

	// jsonPath is a simple JSON path which is evaluated against each custom resource to produce a field selector value. Only JSON paths without the array notation are allowed. Must point to a field of type string, boolean or integer. Types with enum values and strings with formats are allowed. If jsonPath refers to absent field in a resource, the jsonPath evaluates to an empty string. Must not point to metdata fields. Required.
	// Required: true
	JSONPath *string `json:"jsonPath"`
}
