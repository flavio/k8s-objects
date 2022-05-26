// Code generated by go-swagger; DO NOT EDIT.

package v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	api_core_v1 "github.com/kubewarden/k8s-objects/api/core/v1"
	apimachinery_pkg_util_intstr "github.com/kubewarden/k8s-objects/apimachinery/pkg/util/intstr"
)

// IngressBackend IngressBackend describes all endpoints for a given service and port.
//
// swagger:model IngressBackend
type IngressBackend struct {

	// Resource is an ObjectRef to another Kubernetes resource in the namespace of the Ingress object. If resource is specified, serviceName and servicePort must not be specified.
	Resource api_core_v1.TypedLocalObjectReference `json:"resource,omitempty"`

	// Specifies the name of the referenced service.
	ServiceName string `json:"serviceName,omitempty"`

	// Specifies the port of the referenced service.
	ServicePort apimachinery_pkg_util_intstr.IntOrString `json:"servicePort,omitempty"`
}
