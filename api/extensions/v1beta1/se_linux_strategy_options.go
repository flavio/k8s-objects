// Code generated by go-swagger; DO NOT EDIT.

package v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	api_core_v1 "github.com/kubewarden/k8s-objects/api/core/v1"
)

// SELinuxStrategyOptions SELinuxStrategyOptions defines the strategy type and any options used to create the strategy. Deprecated: use SELinuxStrategyOptions from policy API Group instead.
//
// swagger:model SELinuxStrategyOptions
type SELinuxStrategyOptions struct {

	// rule is the strategy that will dictate the allowable labels that may be set.
	// Required: true
	Rule *string `json:"rule"`

	// seLinuxOptions required to run as; required for MustRunAs More info: https://kubernetes.io/docs/tasks/configure-pod-container/security-context/
	SELinuxOptions *api_core_v1.SELinuxOptions `json:"seLinuxOptions,omitempty"`
}