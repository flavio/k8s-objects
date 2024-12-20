// Code generated by go-swagger; DO NOT EDIT.

package v1

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

// HorizontalPodAutoscalerSpec specification of a horizontal pod autoscaler.
//
// swagger:model HorizontalPodAutoscalerSpec
type HorizontalPodAutoscalerSpec struct {

	// maxReplicas is the upper limit for the number of pods that can be set by the autoscaler; cannot be smaller than MinReplicas.
	// Required: true
	MaxReplicas *int32 `json:"maxReplicas"`

	// minReplicas is the lower limit for the number of replicas to which the autoscaler can scale down.  It defaults to 1 pod.  minReplicas is allowed to be 0 if the alpha feature gate HPAScaleToZero is enabled and at least one Object or External metric is configured.  Scaling is active as long as at least one metric value is available.
	MinReplicas int32 `json:"minReplicas,omitempty"`

	// reference to scaled resource; horizontal pod autoscaler will learn the current resource consumption and will set the desired number of pods by using its Scale subresource.
	// Required: true
	ScaleTargetRef *CrossVersionObjectReference `json:"scaleTargetRef"`

	// targetCPUUtilizationPercentage is the target average CPU utilization (represented as a percentage of requested CPU) over all the pods; if not specified the default autoscaling policy will be used.
	TargetCPUUtilizationPercentage int32 `json:"targetCPUUtilizationPercentage,omitempty"`
}
