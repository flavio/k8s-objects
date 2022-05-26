// Code generated by go-swagger; DO NOT EDIT.

package v2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	apimachinery_pkg_api_resource "github.com/kubewarden/k8s-objects/apimachinery/pkg/api/resource"
)

// MetricTarget MetricTarget defines the target value, average value, or average utilization of a specific metric
//
// swagger:model MetricTarget
type MetricTarget struct {

	// averageUtilization is the target value of the average of the resource metric across all relevant pods, represented as a percentage of the requested value of the resource for the pods. Currently only valid for Resource metric source type
	AverageUtilization int32 `json:"averageUtilization,omitempty"`

	// averageValue is the target value of the average of the metric across all relevant pods (as a quantity)
	AverageValue apimachinery_pkg_api_resource.Quantity `json:"averageValue,omitempty"`

	// type represents whether the metric type is Utilization, Value, or AverageValue
	// Required: true
	Type *string `json:"type"`

	// value is the target value of the metric (as a quantity).
	Value apimachinery_pkg_api_resource.Quantity `json:"value,omitempty"`
}