// Code generated by go-swagger; DO NOT EDIT.

package v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// VolumeAttachmentSource VolumeAttachmentSource represents a volume that should be attached. Right now only PersistenVolumes can be attached via external attacher, in future we may allow also inline volumes in pods. Exactly one member can be set.
//
// swagger:model VolumeAttachmentSource
type VolumeAttachmentSource struct {

	// Name of the persistent volume to attach.
	PersistentVolumeName string `json:"persistentVolumeName,omitempty"`
}
