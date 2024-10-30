// Code generated by go-swagger; DO NOT EDIT.

package v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	apimachinery_pkg_apis_meta_v1 "github.com/kubewarden/k8s-objects/apimachinery/pkg/apis/meta/v1"
)

// MigrationCondition Describes the state of a migration at a certain point.
//
// swagger:model MigrationCondition
type MigrationCondition struct {

	// The last time this condition was updated.
	LastUpdateTime *apimachinery_pkg_apis_meta_v1.Time `json:"lastUpdateTime,omitempty"`

	// A human readable message indicating details about the transition.
	Message string `json:"message,omitempty"`

	// The reason for the condition's last transition.
	Reason string `json:"reason,omitempty"`

	// Status of the condition, one of True, False, Unknown.
	// Required: true
	Status *string `json:"status"`

	// Type of the condition.
	// Required: true
	Type *string `json:"type"`
}
