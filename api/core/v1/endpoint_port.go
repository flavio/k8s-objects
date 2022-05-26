// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// EndpointPort EndpointPort is a tuple that describes a single port.
//
// swagger:model EndpointPort
type EndpointPort struct {

	// The application protocol for this port. This field follows standard Kubernetes label syntax. Un-prefixed names are reserved for IANA standard service names (as per RFC-6335 and http://www.iana.org/assignments/service-names). Non-standard protocols should use prefixed names such as mycompany.com/my-custom-protocol. Field can be enabled with ServiceAppProtocol feature gate.
	AppProtocol string `json:"appProtocol,omitempty"`

	// The name of this port.  This must match the 'name' field in the corresponding ServicePort. Must be a DNS_LABEL. Optional only if one port is defined.
	Name string `json:"name,omitempty"`

	// The port number of the endpoint.
	// Required: true
	Port *int32 `json:"port"`

	// The IP protocol for this port. Must be UDP, TCP, or SCTP. Default is TCP.
	Protocol string `json:"protocol,omitempty"`
}
