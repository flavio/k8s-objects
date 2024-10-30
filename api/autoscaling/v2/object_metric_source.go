// Code generated by go-swagger; DO NOT EDIT.

package v2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// ObjectMetricSource ObjectMetricSource indicates how to scale on a metric describing a kubernetes object (for example, hits-per-second on an Ingress object).
//
// swagger:model ObjectMetricSource
type ObjectMetricSource struct {

	// describedObject specifies the descriptions of a object,such as kind,name apiVersion
	// Required: true
	DescribedObject *CrossVersionObjectReference `json:"describedObject"`

	// metric identifies the target metric by name and selector
	// Required: true
	Metric *MetricIdentifier `json:"metric"`

	// target specifies the target value for the given metric
	// Required: true
	Target *MetricTarget `json:"target"`
}
