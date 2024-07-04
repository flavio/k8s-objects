// Code generated by go-swagger; DO NOT EDIT.

package v1alpha2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	apimachinery_pkg_api_resource "github.com/kubewarden/k8s-objects/apimachinery/pkg/api/resource"
)

// NamedResourcesAttribute NamedResourcesAttribute is a combination of an attribute name and its value.
//
// swagger:model NamedResourcesAttribute
type NamedResourcesAttribute struct {

	// BoolValue is a true/false value.
	Bool bool `json:"bool,omitempty"`

	// IntValue is a 64-bit integer.
	Int int64 `json:"int,omitempty"`

	// IntSliceValue is an array of 64-bit integers.
	IntSlice *NamedResourcesIntSlice `json:"intSlice,omitempty"`

	// Name is unique identifier among all resource instances managed by the driver on the node. It must be a DNS subdomain.
	// Required: true
	Name *string `json:"name"`

	// QuantityValue is a quantity.
	Quantity *apimachinery_pkg_api_resource.Quantity `json:"quantity,omitempty"`

	// StringValue is a string.
	String string `json:"string,omitempty"`

	// StringSliceValue is an array of strings.
	StringSlice *NamedResourcesStringSlice `json:"stringSlice,omitempty"`

	// VersionValue is a semantic version according to semver.org spec 2.0.0.
	Version string `json:"version,omitempty"`
}
