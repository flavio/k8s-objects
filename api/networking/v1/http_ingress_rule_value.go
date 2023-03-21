// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// HTTPIngressRuleValue HTTPIngressRuleValue is a list of http selectors pointing to backends. In the example: http://<host>/<path>?<searchpart> -> backend where where parts of the url correspond to RFC 3986, this resource will be used to match against everything after the last '/' and before the first '?' or '#'.
//
// swagger:model HTTPIngressRuleValue
type HTTPIngressRuleValue struct {

	// A collection of paths that map requests to backends.
	// Required: true
	Paths []*HTTPIngressPath `json:"paths"`
}
