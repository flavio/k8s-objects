// Code generated by GroupVersionResource generator for getting GVK data. DO NOT EDIT.

package v1beta1

import "github.com/kubewarden/k8s-objects/apimachinery/pkg/runtime/schema"

func (v *NetworkPolicy) GroupVersionKind() schema.GroupVersionKind {
    kind := v.Kind
    apiVersion := v.APIVersion
    if kind == "" {
        kind = "NetworkPolicy"
    }
    if apiVersion == "" {
        apiVersion = SchemeGroupVersion.String()
    }

    return schema.FromAPIVersionAndKind(apiVersion, kind)
}