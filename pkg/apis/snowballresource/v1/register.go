package v1

import (
	"github.com/robel-yemane/snowball-controller/pkg/apis/snowballresource"
	meta_v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// GroupVersion is the identifier for the API which includes
// the name of the group and the version of the API

var SchemeGroupVersion = schema.GroupVersion{
	Group:   snowballresource.GroupName,
	Version: "v1",
}

// Kind takes an unqualified kind and returns back a Group qualified GroupKind
func Kind(kind string) schema.GroupKind {
	return SchemeGroupVersion.WithKind(kind).GroupKind()
}

// Resource takes an unqualified resource and returns a Group qualified GroupResource
func Resource(resource string) schema.GroupResource {
	return SchemeGroupVersion.WithResource(resource).GroupResource()
}

// create a SchemeBuilder which uses functions to add types to
// the scheme

var AddToScheme = runtime.NewSchemeBuilder(addKnownTypes).AddToScheme


// addKnownTypes adds our types to the API scheme by registering
// SnowballResource and SnowballResourceList
func addKnownTypes(scheme *runtime.Scheme) error {
	scheme.AddKnownTypes(SchemeGroupVersion,
		&SnowballResource{},
		&SnowballResourceList{},
	)

	// register the thpe in the scheme
	meta_v1.AddToGroupVersion(scheme, SchemeGroupVersion)
	return nil
}
