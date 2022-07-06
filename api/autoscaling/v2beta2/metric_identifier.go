// Code generated by go-swagger; DO NOT EDIT.

package v2beta2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	apimachinery_pkg_apis_meta_v1 "github.com/kubewarden/k8s-objects/apimachinery/pkg/apis/meta/v1"
)

// MetricIdentifier MetricIdentifier defines the name and optionally selector for a metric
//
// swagger:model MetricIdentifier
type MetricIdentifier struct {

	// name is the name of the given metric
	// Required: true
	Name *string `json:"name"`

	// selector is the string-encoded form of a standard kubernetes label selector for the given metric When set, it is passed as an additional parameter to the metrics server for more specific metrics scoping. When unset, just the metricName will be used to gather metrics.
	Selector *apimachinery_pkg_apis_meta_v1.LabelSelector `json:"selector,omitempty"`
}
