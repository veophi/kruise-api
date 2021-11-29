/*
Copyright 2021 The Kruise Authors.

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
	v1alpha1 "github.com/openkruise/kruise-api/apps/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeNodeImages implements NodeImageInterface
type FakeNodeImages struct {
	Fake *FakeAppsV1alpha1
}

var nodeimagesResource = schema.GroupVersionResource{Group: "apps.kruise.io", Version: "v1alpha1", Resource: "nodeimages"}

var nodeimagesKind = schema.GroupVersionKind{Group: "apps.kruise.io", Version: "v1alpha1", Kind: "NodeImage"}

// Get takes name of the nodeImage, and returns the corresponding nodeImage object, and an error if there is any.
func (c *FakeNodeImages) Get(name string, options v1.GetOptions) (result *v1alpha1.NodeImage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(nodeimagesResource, name), &v1alpha1.NodeImage{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodeImage), err
}

// List takes label and field selectors, and returns the list of NodeImages that match those selectors.
func (c *FakeNodeImages) List(opts v1.ListOptions) (result *v1alpha1.NodeImageList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(nodeimagesResource, nodeimagesKind, opts), &v1alpha1.NodeImageList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.NodeImageList{ListMeta: obj.(*v1alpha1.NodeImageList).ListMeta}
	for _, item := range obj.(*v1alpha1.NodeImageList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested nodeImages.
func (c *FakeNodeImages) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(nodeimagesResource, opts))
}

// Create takes the representation of a nodeImage and creates it.  Returns the server's representation of the nodeImage, and an error, if there is any.
func (c *FakeNodeImages) Create(nodeImage *v1alpha1.NodeImage) (result *v1alpha1.NodeImage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(nodeimagesResource, nodeImage), &v1alpha1.NodeImage{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodeImage), err
}

// Update takes the representation of a nodeImage and updates it. Returns the server's representation of the nodeImage, and an error, if there is any.
func (c *FakeNodeImages) Update(nodeImage *v1alpha1.NodeImage) (result *v1alpha1.NodeImage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(nodeimagesResource, nodeImage), &v1alpha1.NodeImage{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodeImage), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeNodeImages) UpdateStatus(nodeImage *v1alpha1.NodeImage) (*v1alpha1.NodeImage, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(nodeimagesResource, "status", nodeImage), &v1alpha1.NodeImage{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodeImage), err
}

// Delete takes name of the nodeImage and deletes it. Returns an error if one occurs.
func (c *FakeNodeImages) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(nodeimagesResource, name), &v1alpha1.NodeImage{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeNodeImages) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(nodeimagesResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.NodeImageList{})
	return err
}

// Patch applies the patch and returns the patched nodeImage.
func (c *FakeNodeImages) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.NodeImage, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(nodeimagesResource, name, pt, data, subresources...), &v1alpha1.NodeImage{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.NodeImage), err
}
