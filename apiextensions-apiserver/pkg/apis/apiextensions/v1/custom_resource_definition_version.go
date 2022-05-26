// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// CustomResourceDefinitionVersion CustomResourceDefinitionVersion describes a version for CRD.
//
// swagger:model CustomResourceDefinitionVersion
type CustomResourceDefinitionVersion struct {

	// additionalPrinterColumns specifies additional columns returned in Table output. See https://kubernetes.io/docs/reference/using-api/api-concepts/#receiving-resources-as-tables for details. If no columns are specified, a single column displaying the age of the custom resource is used.
	AdditionalPrinterColumns []*CustomResourceColumnDefinition `json:"additionalPrinterColumns"`

	// name is the version name, e.g. “v1”, “v2beta1”, etc. The custom resources are served under this version at `/apis/<group>/<version>/...` if `served` is true.
	// Required: true
	Name *string `json:"name"`

	// schema describes the schema used for validation, pruning, and defaulting of this version of the custom resource.
	Schema *CustomResourceValidation `json:"schema,omitempty"`

	// served is a flag enabling/disabling this version from being served via REST APIs
	// Required: true
	Served *bool `json:"served"`

	// storage indicates this version should be used when persisting custom resources to storage. There must be exactly one version with storage=true.
	// Required: true
	Storage *bool `json:"storage"`

	// subresources specify what subresources this version of the defined custom resource have.
	Subresources *CustomResourceSubresources `json:"subresources,omitempty"`
}
