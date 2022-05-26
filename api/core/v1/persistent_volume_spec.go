// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	apimachinery_pkg_api_resource "github.com/kubewarden/k8s-objects/apimachinery/pkg/api/resource"
)

// PersistentVolumeSpec PersistentVolumeSpec is the specification of a persistent volume.
//
// swagger:model PersistentVolumeSpec
type PersistentVolumeSpec struct {

	// AccessModes contains all ways the volume can be mounted. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#access-modes
	AccessModes []string `json:"accessModes"`

	// AWSElasticBlockStore represents an AWS Disk resource that is attached to a kubelet's host machine and then exposed to the pod. More info: https://kubernetes.io/docs/concepts/storage/volumes#awselasticblockstore
	AwsElasticBlockStore *AWSElasticBlockStoreVolumeSource `json:"awsElasticBlockStore,omitempty"`

	// AzureDisk represents an Azure Data Disk mount on the host and bind mount to the pod.
	AzureDisk *AzureDiskVolumeSource `json:"azureDisk,omitempty"`

	// AzureFile represents an Azure File Service mount on the host and bind mount to the pod.
	AzureFile *AzureFilePersistentVolumeSource `json:"azureFile,omitempty"`

	// A description of the persistent volume's resources and capacity. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#capacity
	Capacity map[string]apimachinery_pkg_api_resource.Quantity `json:"capacity,omitempty"`

	// CephFS represents a Ceph FS mount on the host that shares a pod's lifetime
	Cephfs *CephFSPersistentVolumeSource `json:"cephfs,omitempty"`

	// Cinder represents a cinder volume attached and mounted on kubelets host machine. More info: https://examples.k8s.io/mysql-cinder-pd/README.md
	Cinder *CinderPersistentVolumeSource `json:"cinder,omitempty"`

	// ClaimRef is part of a bi-directional binding between PersistentVolume and PersistentVolumeClaim. Expected to be non-nil when bound. claim.VolumeName is the authoritative bind between PV and PVC. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#binding
	ClaimRef *ObjectReference `json:"claimRef,omitempty"`

	// CSI represents storage that is handled by an external CSI driver (Beta feature).
	Csi *CSIPersistentVolumeSource `json:"csi,omitempty"`

	// FC represents a Fibre Channel resource that is attached to a kubelet's host machine and then exposed to the pod.
	Fc *FCVolumeSource `json:"fc,omitempty"`

	// FlexVolume represents a generic volume resource that is provisioned/attached using an exec based plugin.
	FlexVolume *FlexPersistentVolumeSource `json:"flexVolume,omitempty"`

	// Flocker represents a Flocker volume attached to a kubelet's host machine and exposed to the pod for its usage. This depends on the Flocker control service being running
	Flocker *FlockerVolumeSource `json:"flocker,omitempty"`

	// GCEPersistentDisk represents a GCE Disk resource that is attached to a kubelet's host machine and then exposed to the pod. Provisioned by an admin. More info: https://kubernetes.io/docs/concepts/storage/volumes#gcepersistentdisk
	GcePersistentDisk *GCEPersistentDiskVolumeSource `json:"gcePersistentDisk,omitempty"`

	// Glusterfs represents a Glusterfs volume that is attached to a host and exposed to the pod. Provisioned by an admin. More info: https://examples.k8s.io/volumes/glusterfs/README.md
	Glusterfs *GlusterfsPersistentVolumeSource `json:"glusterfs,omitempty"`

	// HostPath represents a directory on the host. Provisioned by a developer or tester. This is useful for single-node development and testing only! On-host storage is not supported in any way and WILL NOT WORK in a multi-node cluster. More info: https://kubernetes.io/docs/concepts/storage/volumes#hostpath
	HostPath *HostPathVolumeSource `json:"hostPath,omitempty"`

	// ISCSI represents an ISCSI Disk resource that is attached to a kubelet's host machine and then exposed to the pod. Provisioned by an admin.
	Iscsi *ISCSIPersistentVolumeSource `json:"iscsi,omitempty"`

	// Local represents directly-attached storage with node affinity
	Local *LocalVolumeSource `json:"local,omitempty"`

	// A list of mount options, e.g. ["ro", "soft"]. Not validated - mount will simply fail if one is invalid. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes/#mount-options
	MountOptions []string `json:"mountOptions"`

	// NFS represents an NFS mount on the host. Provisioned by an admin. More info: https://kubernetes.io/docs/concepts/storage/volumes#nfs
	Nfs *NFSVolumeSource `json:"nfs,omitempty"`

	// NodeAffinity defines constraints that limit what nodes this volume can be accessed from. This field influences the scheduling of pods that use this volume.
	NodeAffinity *VolumeNodeAffinity `json:"nodeAffinity,omitempty"`

	// What happens to a persistent volume when released from its claim. Valid options are Retain (default for manually created PersistentVolumes), Delete (default for dynamically provisioned PersistentVolumes), and Recycle (deprecated). Recycle must be supported by the volume plugin underlying this PersistentVolume. More info: https://kubernetes.io/docs/concepts/storage/persistent-volumes#reclaiming
	PersistentVolumeReclaimPolicy string `json:"persistentVolumeReclaimPolicy,omitempty"`

	// PhotonPersistentDisk represents a PhotonController persistent disk attached and mounted on kubelets host machine
	PhotonPersistentDisk *PhotonPersistentDiskVolumeSource `json:"photonPersistentDisk,omitempty"`

	// PortworxVolume represents a portworx volume attached and mounted on kubelets host machine
	PortworxVolume *PortworxVolumeSource `json:"portworxVolume,omitempty"`

	// Quobyte represents a Quobyte mount on the host that shares a pod's lifetime
	Quobyte *QuobyteVolumeSource `json:"quobyte,omitempty"`

	// RBD represents a Rados Block Device mount on the host that shares a pod's lifetime. More info: https://examples.k8s.io/volumes/rbd/README.md
	Rbd *RBDPersistentVolumeSource `json:"rbd,omitempty"`

	// ScaleIO represents a ScaleIO persistent volume attached and mounted on Kubernetes nodes.
	ScaleIO *ScaleIOPersistentVolumeSource `json:"scaleIO,omitempty"`

	// Name of StorageClass to which this persistent volume belongs. Empty value means that this volume does not belong to any StorageClass.
	StorageClassName string `json:"storageClassName,omitempty"`

	// StorageOS represents a StorageOS volume that is attached to the kubelet's host machine and mounted into the pod More info: https://examples.k8s.io/volumes/storageos/README.md
	Storageos *StorageOSPersistentVolumeSource `json:"storageos,omitempty"`

	// volumeMode defines if a volume is intended to be used with a formatted filesystem or to remain in raw block state. Value of Filesystem is implied when not included in spec. This is a beta feature.
	VolumeMode string `json:"volumeMode,omitempty"`

	// VsphereVolume represents a vSphere volume attached and mounted on kubelets host machine
	VsphereVolume *VsphereVirtualDiskVolumeSource `json:"vsphereVolume,omitempty"`
}
