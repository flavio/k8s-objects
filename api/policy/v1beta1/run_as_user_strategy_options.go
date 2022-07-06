// Code generated by go-swagger; DO NOT EDIT.

package v1beta1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// RunAsUserStrategyOptions RunAsUserStrategyOptions defines the strategy type and any options used to create the strategy.
//
// swagger:model RunAsUserStrategyOptions
type RunAsUserStrategyOptions struct {

	// ranges are the allowed ranges of uids that may be used. If you would like to force a single uid then supply a single range with the same start and end. Required for MustRunAs.
	Ranges []*IDRange `json:"ranges,omitempty"`

	// rule is the strategy that will dictate the allowable RunAsUser values that may be set.
	// Required: true
	Rule *string `json:"rule"`
}
