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

// FakeResourceDistributions implements ResourceDistributionInterface
type FakeResourceDistributions struct {
	Fake *FakeAppsV1alpha1
}

var resourcedistributionsResource = schema.GroupVersionResource{Group: "apps.kruise.io", Version: "v1alpha1", Resource: "resourcedistributions"}

var resourcedistributionsKind = schema.GroupVersionKind{Group: "apps.kruise.io", Version: "v1alpha1", Kind: "ResourceDistribution"}

// Get takes name of the resourceDistribution, and returns the corresponding resourceDistribution object, and an error if there is any.
func (c *FakeResourceDistributions) Get(name string, options v1.GetOptions) (result *v1alpha1.ResourceDistribution, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(resourcedistributionsResource, name), &v1alpha1.ResourceDistribution{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResourceDistribution), err
}

// List takes label and field selectors, and returns the list of ResourceDistributions that match those selectors.
func (c *FakeResourceDistributions) List(opts v1.ListOptions) (result *v1alpha1.ResourceDistributionList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(resourcedistributionsResource, resourcedistributionsKind, opts), &v1alpha1.ResourceDistributionList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.ResourceDistributionList{ListMeta: obj.(*v1alpha1.ResourceDistributionList).ListMeta}
	for _, item := range obj.(*v1alpha1.ResourceDistributionList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested resourceDistributions.
func (c *FakeResourceDistributions) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(resourcedistributionsResource, opts))
}

// Create takes the representation of a resourceDistribution and creates it.  Returns the server's representation of the resourceDistribution, and an error, if there is any.
func (c *FakeResourceDistributions) Create(resourceDistribution *v1alpha1.ResourceDistribution) (result *v1alpha1.ResourceDistribution, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(resourcedistributionsResource, resourceDistribution), &v1alpha1.ResourceDistribution{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResourceDistribution), err
}

// Update takes the representation of a resourceDistribution and updates it. Returns the server's representation of the resourceDistribution, and an error, if there is any.
func (c *FakeResourceDistributions) Update(resourceDistribution *v1alpha1.ResourceDistribution) (result *v1alpha1.ResourceDistribution, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(resourcedistributionsResource, resourceDistribution), &v1alpha1.ResourceDistribution{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResourceDistribution), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeResourceDistributions) UpdateStatus(resourceDistribution *v1alpha1.ResourceDistribution) (*v1alpha1.ResourceDistribution, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(resourcedistributionsResource, "status", resourceDistribution), &v1alpha1.ResourceDistribution{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResourceDistribution), err
}

// Delete takes name of the resourceDistribution and deletes it. Returns an error if one occurs.
func (c *FakeResourceDistributions) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(resourcedistributionsResource, name), &v1alpha1.ResourceDistribution{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeResourceDistributions) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(resourcedistributionsResource, listOptions)

	_, err := c.Fake.Invokes(action, &v1alpha1.ResourceDistributionList{})
	return err
}

// Patch applies the patch and returns the patched resourceDistribution.
func (c *FakeResourceDistributions) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ResourceDistribution, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(resourcedistributionsResource, name, pt, data, subresources...), &v1alpha1.ResourceDistribution{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.ResourceDistribution), err
}
