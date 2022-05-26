// Code generated by go-swagger; DO NOT EDIT.

package v2beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	apimachinery_pkg_api_resource "github.com/kubewarden/k8s-objects/apimachinery/pkg/api/resource"
	apimachinery_pkg_apis_meta_v1 "github.com/kubewarden/k8s-objects/apimachinery/pkg/apis/meta/v1"
)

// ExternalMetricStatus ExternalMetricStatus indicates the current value of a global metric not associated with any Kubernetes object.
//
// swagger:model ExternalMetricStatus
type ExternalMetricStatus struct {

	// currentAverageValue is the current value of metric averaged over autoscaled pods.
	CurrentAverageValue apimachinery_pkg_api_resource.Quantity `json:"currentAverageValue,omitempty"`

	// currentValue is the current value of the metric (as a quantity)
	// Required: true
	CurrentValue *apimachinery_pkg_api_resource.Quantity `json:"currentValue"`

	// metricName is the name of a metric used for autoscaling in metric system.
	// Required: true
	MetricName *string `json:"metricName"`

	// metricSelector is used to identify a specific time series within a given metric.
	MetricSelector apimachinery_pkg_apis_meta_v1.LabelSelector `json:"metricSelector,omitempty"`
}
