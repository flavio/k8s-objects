// Code generated by GroupVersionResource generator for getting GVK data. DO NOT EDIT.

package v1alpha1

import "github.com/kubewarden/k8s-objects/apimachinery/pkg/runtime/schema"

// GroupName is the group name use in this package
const GroupName = "internal.apiserver.k8s.io"

// SchemeGroupVersion is group version used to register these objects
var SchemeGroupVersion = schema.GroupVersion{Group: GroupName, Version: "v1alpha1"}

// Resource takes an unqualified resource and returns a Group qualified GroupResource
func Resource(resource string) schema.GroupResource {
    return SchemeGroupVersion.WithResource(resource).GroupResource()
}
