package v1alpha2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// GroupName specifies the group name used to register the objects.
const GroupName = "gateway.nginx.org"

// SchemeGroupVersion is group version used to register these objects.
var SchemeGroupVersion = schema.GroupVersion{Group: GroupName, Version: "v1alpha2"}

// Resource takes an unqualified resource and returns a Group qualified GroupResource.
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

var (
	// SchemeBuilder collects functions that add things to a scheme. It's to allow
	// code to compile without explicitly referencing generated types. You should
	// declare one in each package that will have generated deep copy or conversion
	// functions.
	SchemeBuilder = runtime.NewSchemeBuilder(addKnownTypes)

	// AddToScheme applies all the stored functions to the scheme. A non-nil error
	// indicates that one function failed and the attempt was abandoned.
	AddToScheme = SchemeBuilder.AddToScheme
)

// Adds the list of known types to Scheme.
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&NginxProxy{},
		&NginxProxyList{},
		&ObservabilityPolicy{},
		&ObservabilityPolicyList{},
	)
	// AddToGroupVersion allows the serialization of client types like ListOptions.
	metav1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
