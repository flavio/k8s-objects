// Code generated by go-swagger; DO NOT EDIT.

package v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/mailru/easyjson"
)

// CustomResourceSubresources CustomResourceSubresources defines the status and scale subresources for CustomResources.
//
// swagger:model CustomResourceSubresources
type CustomResourceSubresources struct {

	// Scale denotes the scale subresource for CustomResources
	Scale *CustomResourceSubresourceScale `json:"scale,omitempty"`

	// Status denotes the status subresource for CustomResources
	Status easyjson.RawMessage `json:"status,omitempty"`
}