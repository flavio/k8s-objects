// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// LoadBalancerIngress LoadBalancerIngress represents the status of a load-balancer ingress point: traffic intended for the service should be sent to an ingress point.
//
// swagger:model LoadBalancerIngress
type LoadBalancerIngress struct {

	// Hostname is set for load-balancer ingress points that are DNS based (typically AWS load-balancers)
	Hostname string `json:"hostname,omitempty"`

	// IP is set for load-balancer ingress points that are IP based (typically GCE or OpenStack load-balancers)
	IP string `json:"ip,omitempty"`
}
