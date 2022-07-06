// Code generated by go-swagger; DO NOT EDIT.

package v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	api_core_v1 "github.com/kubewarden/k8s-objects/api/core/v1"
	apimachinery_pkg_apis_meta_v1 "github.com/kubewarden/k8s-objects/apimachinery/pkg/apis/meta/v1"
)

// Event Event is a report of an event somewhere in the cluster. It generally denotes some state change in the system.
//
// swagger:model Event
type Event struct {

	// action is what action was taken/failed regarding to the regarding object. It is machine-readable. This field can have at most 128 characters.
	Action string `json:"action,omitempty"`

	// APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources
	APIVersion string `json:"apiVersion,omitempty"`

	// deprecatedCount is the deprecated field assuring backward compatibility with core.v1 Event type.
	DeprecatedCount int32 `json:"deprecatedCount,omitempty"`

	// deprecatedFirstTimestamp is the deprecated field assuring backward compatibility with core.v1 Event type.
	DeprecatedFirstTimestamp *apimachinery_pkg_apis_meta_v1.Time `json:"deprecatedFirstTimestamp,omitempty"`

	// deprecatedLastTimestamp is the deprecated field assuring backward compatibility with core.v1 Event type.
	DeprecatedLastTimestamp *apimachinery_pkg_apis_meta_v1.Time `json:"deprecatedLastTimestamp,omitempty"`

	// deprecatedSource is the deprecated field assuring backward compatibility with core.v1 Event type.
	DeprecatedSource *api_core_v1.EventSource `json:"deprecatedSource,omitempty"`

	// eventTime is the time when this Event was first observed. It is required.
	// Required: true
	EventTime *apimachinery_pkg_apis_meta_v1.MicroTime `json:"eventTime"`

	// Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds
	Kind string `json:"kind,omitempty"`

	// metadata
	Metadata *apimachinery_pkg_apis_meta_v1.ObjectMeta `json:"metadata,omitempty"`

	// note is a human-readable description of the status of this operation. Maximal length of the note is 1kB, but libraries should be prepared to handle values up to 64kB.
	Note string `json:"note,omitempty"`

	// reason is why the action was taken. It is human-readable. This field can have at most 128 characters.
	Reason string `json:"reason,omitempty"`

	// regarding contains the object this Event is about. In most cases it's an Object reporting controller implements, e.g. ReplicaSetController implements ReplicaSets and this event is emitted because it acts on some changes in a ReplicaSet object.
	Regarding *api_core_v1.ObjectReference `json:"regarding,omitempty"`

	// related is the optional secondary object for more complex actions. E.g. when regarding object triggers a creation or deletion of related object.
	Related *api_core_v1.ObjectReference `json:"related,omitempty"`

	// reportingController is the name of the controller that emitted this Event, e.g. `kubernetes.io/kubelet`. This field cannot be empty for new Events.
	ReportingController string `json:"reportingController,omitempty"`

	// reportingInstance is the ID of the controller instance, e.g. `kubelet-xyzf`. This field cannot be empty for new Events and it can have at most 128 characters.
	ReportingInstance string `json:"reportingInstance,omitempty"`

	// series is data about the Event series this event represents or nil if it's a singleton Event.
	Series *EventSeries `json:"series,omitempty"`

	// type is the type of this event (Normal, Warning), new types could be added in the future. It is machine-readable.
	Type string `json:"type,omitempty"`
}
