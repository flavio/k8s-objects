// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// CephFSVolumeSource Represents a Ceph Filesystem mount that lasts the lifetime of a pod Cephfs volumes do not support ownership management or SELinux relabeling.
//
// swagger:model CephFSVolumeSource
type CephFSVolumeSource struct {

	// Required: Monitors is a collection of Ceph monitors More info: https://releases.k8s.io/HEAD/examples/volumes/cephfs/README.md#how-to-use-it
	// Required: true
	Monitors []string `json:"monitors"`

	// Optional: Used as the mounted root, rather than the full Ceph tree, default is /
	Path string `json:"path,omitempty"`

	// Optional: Defaults to false (read/write). ReadOnly here will force the ReadOnly setting in VolumeMounts. More info: https://releases.k8s.io/HEAD/examples/volumes/cephfs/README.md#how-to-use-it
	ReadOnly bool `json:"readOnly,omitempty"`

	// Optional: SecretFile is the path to key ring for User, default is /etc/ceph/user.secret More info: https://releases.k8s.io/HEAD/examples/volumes/cephfs/README.md#how-to-use-it
	SecretFile string `json:"secretFile,omitempty"`

	// Optional: SecretRef is reference to the authentication secret for User, default is empty. More info: https://releases.k8s.io/HEAD/examples/volumes/cephfs/README.md#how-to-use-it
	SecretRef *LocalObjectReference `json:"secretRef,omitempty"`

	// Optional: User is the rados user name, default is admin More info: https://releases.k8s.io/HEAD/examples/volumes/cephfs/README.md#how-to-use-it
	User string `json:"user,omitempty"`
}
