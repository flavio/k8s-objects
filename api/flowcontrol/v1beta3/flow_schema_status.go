// Code generated by go-swagger; DO NOT EDIT.

package v1beta3

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// FlowSchemaStatus FlowSchemaStatus represents the current state of a FlowSchema.
//
// swagger:model FlowSchemaStatus
type FlowSchemaStatus struct {

	// `conditions` is a list of the current states of FlowSchema.
	Conditions []*FlowSchemaCondition `json:"conditions,omitempty"`
}
