// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/mailru/easyjson"
)

// JSONSchemaProps JSONSchemaProps is a JSON-Schema following Specification Draft 4 (http://json-schema.org/).
//
// swagger:model JSONSchemaProps
type JSONSchemaProps struct {

	// dollar ref
	DollarRef string `json:"$ref,omitempty"`

	// dollar schema
	DollarSchema string `json:"$schema,omitempty"`

	// additional items
	AdditionalItems easyjson.RawMessage `json:"additionalItems,omitempty"`

	// additional properties
	AdditionalProperties easyjson.RawMessage `json:"additionalProperties,omitempty"`

	// all of
	AllOf []*JSONSchemaProps `json:"allOf,omitempty"`

	// any of
	AnyOf []*JSONSchemaProps `json:"anyOf,omitempty"`

	// default is a default value for undefined object fields. Defaulting is a beta feature under the CustomResourceDefaulting feature gate. Defaulting requires spec.preserveUnknownFields to be false.
	Default easyjson.RawMessage `json:"default,omitempty"`

	// definitions
	Definitions map[string]*JSONSchemaProps `json:"definitions,omitempty"`

	// dependencies
	Dependencies map[string]easyjson.RawMessage `json:"dependencies,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// enum
	Enum []easyjson.RawMessage `json:"enum,omitempty"`

	// example
	Example easyjson.RawMessage `json:"example,omitempty"`

	// exclusive maximum
	ExclusiveMaximum bool `json:"exclusiveMaximum,omitempty"`

	// exclusive minimum
	ExclusiveMinimum bool `json:"exclusiveMinimum,omitempty"`

	// external docs
	ExternalDocs *ExternalDocumentation `json:"externalDocs,omitempty"`

	// format
	Format string `json:"format,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// items
	Items easyjson.RawMessage `json:"items,omitempty"`

	// max items
	MaxItems int64 `json:"maxItems,omitempty"`

	// max length
	MaxLength int64 `json:"maxLength,omitempty"`

	// max properties
	MaxProperties int64 `json:"maxProperties,omitempty"`

	// maximum
	Maximum float64 `json:"maximum,omitempty"`

	// min items
	MinItems int64 `json:"minItems,omitempty"`

	// min length
	MinLength int64 `json:"minLength,omitempty"`

	// min properties
	MinProperties int64 `json:"minProperties,omitempty"`

	// minimum
	Minimum float64 `json:"minimum,omitempty"`

	// multiple of
	MultipleOf float64 `json:"multipleOf,omitempty"`

	// not
	Not *JSONSchemaProps `json:"not,omitempty"`

	// nullable
	Nullable bool `json:"nullable,omitempty"`

	// one of
	OneOf []*JSONSchemaProps `json:"oneOf,omitempty"`

	// pattern
	Pattern string `json:"pattern,omitempty"`

	// pattern properties
	PatternProperties map[string]*JSONSchemaProps `json:"patternProperties,omitempty"`

	// properties
	Properties map[string]*JSONSchemaProps `json:"properties,omitempty"`

	// required
	Required []string `json:"required,omitempty"`

	// title
	Title string `json:"title,omitempty"`

	// type
	Type string `json:"type,omitempty"`

	// unique items
	UniqueItems bool `json:"uniqueItems,omitempty"`

	// x-kubernetes-embedded-resource defines that the value is an embedded Kubernetes runtime.Object, with TypeMeta and ObjectMeta. The type must be object. It is allowed to further restrict the embedded object. kind, apiVersion and metadata are validated automatically. x-kubernetes-preserve-unknown-fields is allowed to be true, but does not have to be if the object is fully specified (up to kind, apiVersion, metadata).
	XKubernetesEmbeddedResource bool `json:"x-kubernetes-embedded-resource,omitempty"`

	// x-kubernetes-int-or-string specifies that this value is either an integer or a string. If this is true, an empty type is allowed and type as child of anyOf is permitted if following one of the following patterns:
	//
	// 1) anyOf:
	//    - type: integer
	//    - type: string
	// 2) allOf:
	//    - anyOf:
	//      - type: integer
	//      - type: string
	//    - ... zero or more
	XKubernetesIntOrString bool `json:"x-kubernetes-int-or-string,omitempty"`

	// x-kubernetes-list-map-keys annotates an array with the x-kubernetes-list-type `map` by specifying the keys used as the index of the map.
	//
	// This tag MUST only be used on lists that have the "x-kubernetes-list-type" extension set to "map". Also, the values specified for this attribute must be a scalar typed field of the child structure (no nesting is supported).
	XKubernetesListMapKeys []string `json:"x-kubernetes-list-map-keys,omitempty"`

	// x-kubernetes-list-type annotates an array to further describe its topology. This extension must only be used on lists and may have 3 possible values:
	//
	// 1) `atomic`: the list is treated as a single entity, like a scalar.
	//      Atomic lists will be entirely replaced when updated. This extension
	//      may be used on any type of list (struct, scalar, ...).
	// 2) `set`:
	//      Sets are lists that must not have multiple items with the same value. Each
	//      value must be a scalar (or another atomic type).
	// 3) `map`:
	//      These lists are like maps in that their elements have a non-index key
	//      used to identify them. Order is preserved upon merge. The map tag
	//      must only be used on a list with elements of type object.
	// Defaults to atomic for arrays.
	XKubernetesListType string `json:"x-kubernetes-list-type,omitempty"`

	// x-kubernetes-preserve-unknown-fields stops the API server decoding step from pruning fields which are not specified in the validation schema. This affects fields recursively, but switches back to normal pruning behaviour if nested properties or additionalProperties are specified in the schema. This can either be true or undefined. False is forbidden.
	XKubernetesPreserveUnknownFields bool `json:"x-kubernetes-preserve-unknown-fields,omitempty"`
}
