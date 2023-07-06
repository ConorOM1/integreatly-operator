// Package v1alpha1 contains API Schema definitions for the dhcp v1alpha1 API group
// +kubebuilder:object:generate=true
// +groupName=addons.managed.openshift.io
package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

const group = "addons.managed.openshift.io"

var (
	// GroupVersion is group version used to register these objects
	GroupVersion = schema.GroupVersion{Group: group, Version: "v1alpha1"}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	SchemeBuilder runtime.SchemeBuilder

	// AddToScheme adds the types in this group-version to the given scheme.
	AddToScheme = SchemeBuilder.AddToScheme
)

func register(objs ...runtime.Object) {
	SchemeBuilder.Register(func(scheme *runtime.Scheme) error {
		scheme.AddKnownTypes(GroupVersion, objs...)
		metav1.AddToGroupVersion(scheme, GroupVersion)
		return nil
	})
}
