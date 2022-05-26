// Code generated by go-swagger; DO NOT EDIT.

package v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// IDRange IDRange provides a min/max of an allowed range of IDs.
//
// swagger:model IDRange
type IDRange struct {

	// max is the end of the range, inclusive.
	// Required: true
	Max *int64 `json:"max"`

	// min is the start of the range, inclusive.
	// Required: true
	Min *int64 `json:"min"`
}
