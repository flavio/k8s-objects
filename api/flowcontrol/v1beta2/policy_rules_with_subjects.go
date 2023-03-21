// Code generated by go-swagger; DO NOT EDIT.

package v1beta2

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// PolicyRulesWithSubjects PolicyRulesWithSubjects prescribes a test that applies to a request to an apiserver. The test considers the subject making the request, the verb being requested, and the resource to be acted upon. This PolicyRulesWithSubjects matches a request if and only if both (a) at least one member of subjects matches the request and (b) at least one member of resourceRules or nonResourceRules matches the request.
//
// swagger:model PolicyRulesWithSubjects
type PolicyRulesWithSubjects struct {

	// `nonResourceRules` is a list of NonResourcePolicyRules that identify matching requests according to their verb and the target non-resource URL.
	NonResourceRules []*NonResourcePolicyRule `json:"nonResourceRules,omitempty"`

	// `resourceRules` is a slice of ResourcePolicyRules that identify matching requests according to their verb and the target resource. At least one of `resourceRules` and `nonResourceRules` has to be non-empty.
	ResourceRules []*ResourcePolicyRule `json:"resourceRules,omitempty"`

	// subjects is the list of normal user, serviceaccount, or group that this rule cares about. There must be at least one member in this slice. A slice that includes both the system:authenticated and system:unauthenticated user groups matches every request. Required.
	// Required: true
	Subjects []*Subject `json:"subjects"`
}
