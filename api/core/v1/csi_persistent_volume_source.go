// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// CSIPersistentVolumeSource Represents storage that is managed by an external CSI volume driver
//
// swagger:model CSIPersistentVolumeSource
type CSIPersistentVolumeSource struct {

	// controllerExpandSecretRef is a reference to the secret object containing sensitive information to pass to the CSI driver to complete the CSI ControllerExpandVolume call. This field is optional, and may be empty if no secret is required. If the secret object contains more than one secret, all secrets are passed.
	ControllerExpandSecretRef *SecretReference `json:"controllerExpandSecretRef,omitempty"`

	// controllerPublishSecretRef is a reference to the secret object containing sensitive information to pass to the CSI driver to complete the CSI ControllerPublishVolume and ControllerUnpublishVolume calls. This field is optional, and may be empty if no secret is required. If the secret object contains more than one secret, all secrets are passed.
	ControllerPublishSecretRef *SecretReference `json:"controllerPublishSecretRef,omitempty"`

	// driver is the name of the driver to use for this volume. Required.
	// Required: true
	Driver *string `json:"driver"`

	// fsType to mount. Must be a filesystem type supported by the host operating system. Ex. "ext4", "xfs", "ntfs".
	FSType string `json:"fsType,omitempty"`

	// nodeExpandSecretRef is a reference to the secret object containing sensitive information to pass to the CSI driver to complete the CSI NodeExpandVolume call. This field is optional, may be omitted if no secret is required. If the secret object contains more than one secret, all secrets are passed.
	NodeExpandSecretRef *SecretReference `json:"nodeExpandSecretRef,omitempty"`

	// nodePublishSecretRef is a reference to the secret object containing sensitive information to pass to the CSI driver to complete the CSI NodePublishVolume and NodeUnpublishVolume calls. This field is optional, and may be empty if no secret is required. If the secret object contains more than one secret, all secrets are passed.
	NodePublishSecretRef *SecretReference `json:"nodePublishSecretRef,omitempty"`

	// nodeStageSecretRef is a reference to the secret object containing sensitive information to pass to the CSI driver to complete the CSI NodeStageVolume and NodeStageVolume and NodeUnstageVolume calls. This field is optional, and may be empty if no secret is required. If the secret object contains more than one secret, all secrets are passed.
	NodeStageSecretRef *SecretReference `json:"nodeStageSecretRef,omitempty"`

	// readOnly value to pass to ControllerPublishVolumeRequest. Defaults to false (read/write).
	ReadOnly bool `json:"readOnly,omitempty"`

	// volumeAttributes of the volume to publish.
	VolumeAttributes map[string]string `json:"volumeAttributes,omitempty"`

	// volumeHandle is the unique volume name returned by the CSI volume plugin’s CreateVolume to refer to the volume on all subsequent calls. Required.
	// Required: true
	VolumeHandle *string `json:"volumeHandle"`
}
