// Code generated by go-swagger; DO NOT EDIT.

package v1alpha2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	apimachinery_pkg_apis_meta_v1 "github.com/kubewarden/k8s-objects/apimachinery/pkg/apis/meta/v1"
)

// ResourceSlice ResourceSlice provides information about available resources on individual nodes.
//
// swagger:model ResourceSlice
type ResourceSlice struct {

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	APIVersion string `json:"apiVersion,omitempty"`

	// DriverName identifies the DRA driver providing the capacity information. A field selector can be used to list only ResourceSlice objects with a certain driver name.
	// Required: true
	DriverName *string `json:"driverName"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind string `json:"kind,omitempty"`

	// Standard object metadata
	Metadata *apimachinery_pkg_apis_meta_v1.ObjectMeta `json:"metadata,omitempty"`

	// NamedResources describes available resources using the named resources model.
	NamedResources *NamedResourcesResources `json:"namedResources,omitempty"`

	// NodeName identifies the node which provides the resources if they are local to a node.
	//
	// A field selector can be used to list only ResourceSlice objects with a certain node name.
	NodeName string `json:"nodeName,omitempty"`
}
