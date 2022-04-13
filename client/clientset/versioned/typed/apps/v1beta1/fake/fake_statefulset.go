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
	"context"

	v1beta1 "github.com/veophi/kruise-api/apps/v1beta1"
	autoscalingv1 "k8s.io/api/autoscaling/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeStatefulSets implements StatefulSetInterface
type FakeStatefulSets struct {
	Fake *FakeAppsV1beta1
	ns   string
}

var statefulsetsResource = schema.GroupVersionResource{Group: "apps.kruise.io", Version: "v1beta1", Resource: "statefulsets"}

var statefulsetsKind = schema.GroupVersionKind{Group: "apps.kruise.io", Version: "v1beta1", Kind: "StatefulSet"}

// Get takes name of the statefulSet, and returns the corresponding statefulSet object, and an error if there is any.
func (c *FakeStatefulSets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.StatefulSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(statefulsetsResource, c.ns, name), &v1beta1.StatefulSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.StatefulSet), err
}

// List takes label and field selectors, and returns the list of StatefulSets that match those selectors.
func (c *FakeStatefulSets) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.StatefulSetList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(statefulsetsResource, statefulsetsKind, c.ns, opts), &v1beta1.StatefulSetList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.StatefulSetList{ListMeta: obj.(*v1beta1.StatefulSetList).ListMeta}
	for _, item := range obj.(*v1beta1.StatefulSetList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested statefulSets.
func (c *FakeStatefulSets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(statefulsetsResource, c.ns, opts))

}

// Create takes the representation of a statefulSet and creates it.  Returns the server's representation of the statefulSet, and an error, if there is any.
func (c *FakeStatefulSets) Create(ctx context.Context, statefulSet *v1beta1.StatefulSet, opts v1.CreateOptions) (result *v1beta1.StatefulSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(statefulsetsResource, c.ns, statefulSet), &v1beta1.StatefulSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.StatefulSet), err
}

// Update takes the representation of a statefulSet and updates it. Returns the server's representation of the statefulSet, and an error, if there is any.
func (c *FakeStatefulSets) Update(ctx context.Context, statefulSet *v1beta1.StatefulSet, opts v1.UpdateOptions) (result *v1beta1.StatefulSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(statefulsetsResource, c.ns, statefulSet), &v1beta1.StatefulSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.StatefulSet), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeStatefulSets) UpdateStatus(ctx context.Context, statefulSet *v1beta1.StatefulSet, opts v1.UpdateOptions) (*v1beta1.StatefulSet, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(statefulsetsResource, "status", c.ns, statefulSet), &v1beta1.StatefulSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.StatefulSet), err
}

// Delete takes name of the statefulSet and deletes it. Returns an error if one occurs.
func (c *FakeStatefulSets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(statefulsetsResource, c.ns, name), &v1beta1.StatefulSet{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeStatefulSets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(statefulsetsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.StatefulSetList{})
	return err
}

// Patch applies the patch and returns the patched statefulSet.
func (c *FakeStatefulSets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.StatefulSet, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(statefulsetsResource, c.ns, name, pt, data, subresources...), &v1beta1.StatefulSet{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.StatefulSet), err
}

// GetScale takes name of the statefulSet, and returns the corresponding scale object, and an error if there is any.
func (c *FakeStatefulSets) GetScale(ctx context.Context, statefulSetName string, options v1.GetOptions) (result *autoscalingv1.Scale, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetSubresourceAction(statefulsetsResource, c.ns, "scale", statefulSetName), &autoscalingv1.Scale{})

	if obj == nil {
		return nil, err
	}
	return obj.(*autoscalingv1.Scale), err
}

// UpdateScale takes the representation of a scale and updates it. Returns the server's representation of the scale, and an error, if there is any.
func (c *FakeStatefulSets) UpdateScale(ctx context.Context, statefulSetName string, scale *autoscalingv1.Scale, opts v1.UpdateOptions) (result *autoscalingv1.Scale, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(statefulsetsResource, "scale", c.ns, scale), &autoscalingv1.Scale{})

	if obj == nil {
		return nil, err
	}
	return obj.(*autoscalingv1.Scale), err
}
