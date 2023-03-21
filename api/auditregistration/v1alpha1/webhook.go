// Code generated by go-swagger; DO NOT EDIT.

package v1alpha1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// Webhook Webhook holds the configuration of the webhook
//
// swagger:model Webhook
type Webhook struct {

	// ClientConfig holds the connection parameters for the webhook required
	// Required: true
	ClientConfig *WebhookClientConfig `json:"clientConfig"`

	// Throttle holds the options for throttling the webhook
	Throttle *WebhookThrottleConfig `json:"throttle,omitempty"`
}