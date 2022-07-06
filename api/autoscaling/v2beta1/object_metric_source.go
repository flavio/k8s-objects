// Code generated by go-swagger; DO NOT EDIT.

package v2beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	apimachinery_pkg_api_resource "github.com/kubewarden/k8s-objects/apimachinery/pkg/api/resource"
	apimachinery_pkg_apis_meta_v1 "github.com/kubewarden/k8s-objects/apimachinery/pkg/apis/meta/v1"
)

// ObjectMetricSource ObjectMetricSource indicates how to scale on a metric describing a kubernetes object (for example, hits-per-second on an Ingress object).
//
// swagger:model ObjectMetricSource
type ObjectMetricSource struct {

	// averageValue is the target value of the average of the metric across all relevant pods (as a quantity)
	AverageValue *apimachinery_pkg_api_resource.Quantity `json:"averageValue,omitempty"`

	// metricName is the name of the metric in question.
	// Required: true
	MetricName *string `json:"metricName"`

	// selector is the string-encoded form of a standard kubernetes label selector for the given metric When set, it is passed as an additional parameter to the metrics server for more specific metrics scoping When unset, just the metricName will be used to gather metrics.
	Selector *apimachinery_pkg_apis_meta_v1.LabelSelector `json:"selector,omitempty"`

	// target is the described Kubernetes object.
	// Required: true
	Target *CrossVersionObjectReference `json:"target"`

	// targetValue is the target value of the metric (as a quantity).
	// Required: true
	TargetValue *apimachinery_pkg_api_resource.Quantity `json:"targetValue"`
}
