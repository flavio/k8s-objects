// Code generated by go-swagger; DO NOT EDIT.

package v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// NetworkDeviceData NetworkDeviceData provides network-related details for the allocated device. This information may be filled by drivers or other components to configure or identify the device within a network context.
//
// swagger:model NetworkDeviceData
type NetworkDeviceData struct {

	// HardwareAddress represents the hardware address (e.g. MAC Address) of the device's network interface.
	//
	// Must not be longer than 128 characters.
	HardwareAddress string `json:"hardwareAddress,omitempty"`

	// InterfaceName specifies the name of the network interface associated with the allocated device. This might be the name of a physical or virtual network interface being configured in the pod.
	//
	// Must not be longer than 256 characters.
	InterfaceName string `json:"interfaceName,omitempty"`

	// IPs lists the network addresses assigned to the device's network interface. This can include both IPv4 and IPv6 addresses. The IPs are in the CIDR notation, which includes both the address and the associated subnet mask. e.g.: "192.0.2.5/24" for IPv4 and "2001:db8::5/64" for IPv6.
	Ips []string `json:"ips,omitempty"`
}
