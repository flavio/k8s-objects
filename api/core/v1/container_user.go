// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// ContainerUser ContainerUser represents user identity information
//
// swagger:model ContainerUser
type ContainerUser struct {

	// Linux holds user identity information initially attached to the first process of the containers in Linux. Note that the actual running identity can be changed if the process has enough privilege to do so.
	Linux *LinuxContainerUser `json:"linux,omitempty"`
}