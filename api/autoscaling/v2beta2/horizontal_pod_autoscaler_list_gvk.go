// Code generated by GroupVersionResource generator for getting GVK data. DO NOT EDIT.

package v2beta2

import "github.com/kubewarden/k8s-objects/apimachinery/pkg/runtime/schema"

func (v *HorizontalPodAutoscalerList) GroupVersionKind() schema.GroupVersionKind {
    kind := v.Kind
    apiVersion := v.APIVersion
    if kind == "" {
        kind = "HorizontalPodAutoscalerList"
    }
    if apiVersion == "" {
        apiVersion = SchemeGroupVersion.String()
    }

    return schema.FromAPIVersionAndKind(apiVersion, kind)
}
