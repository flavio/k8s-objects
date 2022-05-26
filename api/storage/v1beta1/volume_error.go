// Code generated by go-swagger; DO NOT EDIT.

package v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	apimachinery_pkg_apis_meta_v1 "github.com/kubewarden/k8s-objects/apimachinery/pkg/apis/meta/v1"
)

// VolumeError VolumeError captures an error encountered during a volume operation.
//
// swagger:model VolumeError
type VolumeError struct {

	// String detailing the error encountered during Attach or Detach operation. This string may be logged, so it should not contain sensitive information.
	Message string `json:"message,omitempty"`

	// Time the error was encountered.
	Time apimachinery_pkg_apis_meta_v1.Time `json:"time,omitempty"`
}
