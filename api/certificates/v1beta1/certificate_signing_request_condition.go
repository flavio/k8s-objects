// Code generated by go-swagger; DO NOT EDIT.

package v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	apimachinery_pkg_apis_meta_v1 "github.com/kubewarden/k8s-objects/apimachinery/pkg/apis/meta/v1"
)

// CertificateSigningRequestCondition certificate signing request condition
//
// swagger:model CertificateSigningRequestCondition
type CertificateSigningRequestCondition struct {

	// timestamp for the last update to this condition
	LastUpdateTime *apimachinery_pkg_apis_meta_v1.Time `json:"lastUpdateTime,omitempty"`

	// human readable message with details about the request state
	Message string `json:"message,omitempty"`

	// brief reason for the request state
	Reason string `json:"reason,omitempty"`

	// request approval state, currently Approved or Denied.
	// Required: true
	Type *string `json:"type"`
}