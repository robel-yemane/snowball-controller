/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	snowballresourcev1 "github.com/robel-yemane/snowball-controller/pkg/apis/snowballresource/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeSnowballResources implements SnowballResourceInterface
type FakeSnowballResources struct {
	Fake *FakeSnowballV1
	ns   string
}

var snowballresourcesResource = schema.GroupVersionResource{Group: "snowball.com", Version: "v1", Resource: "snowballresources"}

var snowballresourcesKind = schema.GroupVersionKind{Group: "snowball.com", Version: "v1", Kind: "SnowballResource"}

// Get takes name of the snowballResource, and returns the corresponding snowballResource object, and an error if there is any.
func (c *FakeSnowballResources) Get(name string, options v1.GetOptions) (result *snowballresourcev1.SnowballResource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(snowballresourcesResource, c.ns, name), &snowballresourcev1.SnowballResource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*snowballresourcev1.SnowballResource), err
}

// List takes label and field selectors, and returns the list of SnowballResources that match those selectors.
func (c *FakeSnowballResources) List(opts v1.ListOptions) (result *snowballresourcev1.SnowballResourceList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(snowballresourcesResource, snowballresourcesKind, c.ns, opts), &snowballresourcev1.SnowballResourceList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &snowballresourcev1.SnowballResourceList{ListMeta: obj.(*snowballresourcev1.SnowballResourceList).ListMeta}
	for _, item := range obj.(*snowballresourcev1.SnowballResourceList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested snowballResources.
func (c *FakeSnowballResources) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(snowballresourcesResource, c.ns, opts))

}

// Create takes the representation of a snowballResource and creates it.  Returns the server's representation of the snowballResource, and an error, if there is any.
func (c *FakeSnowballResources) Create(snowballResource *snowballresourcev1.SnowballResource) (result *snowballresourcev1.SnowballResource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(snowballresourcesResource, c.ns, snowballResource), &snowballresourcev1.SnowballResource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*snowballresourcev1.SnowballResource), err
}

// Update takes the representation of a snowballResource and updates it. Returns the server's representation of the snowballResource, and an error, if there is any.
func (c *FakeSnowballResources) Update(snowballResource *snowballresourcev1.SnowballResource) (result *snowballresourcev1.SnowballResource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(snowballresourcesResource, c.ns, snowballResource), &snowballresourcev1.SnowballResource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*snowballresourcev1.SnowballResource), err
}

// Delete takes name of the snowballResource and deletes it. Returns an error if one occurs.
func (c *FakeSnowballResources) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(snowballresourcesResource, c.ns, name), &snowballresourcev1.SnowballResource{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeSnowballResources) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(snowballresourcesResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &snowballresourcev1.SnowballResourceList{})
	return err
}

// Patch applies the patch and returns the patched snowballResource.
func (c *FakeSnowballResources) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *snowballresourcev1.SnowballResource, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(snowballresourcesResource, c.ns, name, pt, data, subresources...), &snowballresourcev1.SnowballResource{})

	if obj == nil {
		return nil, err
	}
	return obj.(*snowballresourcev1.SnowballResource), err
}
