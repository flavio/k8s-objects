// Code generated by go-swagger; DO NOT EDIT.

package v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// EndpointPort EndpointPort represents a Port used by an EndpointSlice
//
// swagger:model EndpointPort
type EndpointPort struct {

	// The application protocol for this port. This field follows standard Kubernetes label syntax. Un-prefixed names are reserved for IANA standard service names (as per RFC-6335 and http://www.iana.org/assignments/service-names). Non-standard protocols should use prefixed names. Default is empty string.
	AppProtocol string `json:"appProtocol,omitempty"`

	// The name of this port. All ports in an EndpointSlice must have a unique name. If the EndpointSlice is dervied from a Kubernetes service, this corresponds to the Service.ports[].name. Name must either be an empty string or pass DNS_LABEL validation: * must be no more than 63 characters long. * must consist of lower case alphanumeric characters or '-'. * must start and end with an alphanumeric character. Default is empty string.
	Name string `json:"name,omitempty"`

	// The port number of the endpoint. If this is not specified, ports are not restricted and must be interpreted in the context of the specific consumer.
	Port int32 `json:"port,omitempty"`

	// The IP protocol for this port. Must be UDP, TCP, or SCTP. Default is TCP.
	Protocol string `json:"protocol,omitempty"`
}
