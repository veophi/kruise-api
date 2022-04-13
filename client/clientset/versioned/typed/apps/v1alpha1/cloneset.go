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

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/veophi/kruise-api/apps/v1alpha1"
	scheme "github.com/veophi/kruise-api/client/clientset/versioned/scheme"
	autoscalingv1 "k8s.io/api/autoscaling/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// CloneSetsGetter has a method to return a CloneSetInterface.
// A group's client should implement this interface.
type CloneSetsGetter interface {
	CloneSets(namespace string) CloneSetInterface
}

// CloneSetInterface has methods to work with CloneSet resources.
type CloneSetInterface interface {
	Create(ctx context.Context, cloneSet *v1alpha1.CloneSet, opts v1.CreateOptions) (*v1alpha1.CloneSet, error)
	Update(ctx context.Context, cloneSet *v1alpha1.CloneSet, opts v1.UpdateOptions) (*v1alpha1.CloneSet, error)
	UpdateStatus(ctx context.Context, cloneSet *v1alpha1.CloneSet, opts v1.UpdateOptions) (*v1alpha1.CloneSet, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.CloneSet, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.CloneSetList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CloneSet, err error)
	GetScale(ctx context.Context, cloneSetName string, options v1.GetOptions) (*autoscalingv1.Scale, error)
	UpdateScale(ctx context.Context, cloneSetName string, scale *autoscalingv1.Scale, opts v1.UpdateOptions) (*autoscalingv1.Scale, error)

	CloneSetExpansion
}

// cloneSets implements CloneSetInterface
type cloneSets struct {
	client rest.Interface
	ns     string
}

// newCloneSets returns a CloneSets
func newCloneSets(c *AppsV1alpha1Client, namespace string) *cloneSets {
	return &cloneSets{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the cloneSet, and returns the corresponding cloneSet object, and an error if there is any.
func (c *cloneSets) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.CloneSet, err error) {
	result = &v1alpha1.CloneSet{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("clonesets").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of CloneSets that match those selectors.
func (c *cloneSets) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.CloneSetList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.CloneSetList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("clonesets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested cloneSets.
func (c *cloneSets) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("clonesets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a cloneSet and creates it.  Returns the server's representation of the cloneSet, and an error, if there is any.
func (c *cloneSets) Create(ctx context.Context, cloneSet *v1alpha1.CloneSet, opts v1.CreateOptions) (result *v1alpha1.CloneSet, err error) {
	result = &v1alpha1.CloneSet{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("clonesets").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(cloneSet).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a cloneSet and updates it. Returns the server's representation of the cloneSet, and an error, if there is any.
func (c *cloneSets) Update(ctx context.Context, cloneSet *v1alpha1.CloneSet, opts v1.UpdateOptions) (result *v1alpha1.CloneSet, err error) {
	result = &v1alpha1.CloneSet{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("clonesets").
		Name(cloneSet.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(cloneSet).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *cloneSets) UpdateStatus(ctx context.Context, cloneSet *v1alpha1.CloneSet, opts v1.UpdateOptions) (result *v1alpha1.CloneSet, err error) {
	result = &v1alpha1.CloneSet{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("clonesets").
		Name(cloneSet.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(cloneSet).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the cloneSet and deletes it. Returns an error if one occurs.
func (c *cloneSets) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("clonesets").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *cloneSets) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("clonesets").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched cloneSet.
func (c *cloneSets) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.CloneSet, err error) {
	result = &v1alpha1.CloneSet{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("clonesets").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}

// GetScale takes name of the cloneSet, and returns the corresponding autoscalingv1.Scale object, and an error if there is any.
func (c *cloneSets) GetScale(ctx context.Context, cloneSetName string, options v1.GetOptions) (result *autoscalingv1.Scale, err error) {
	result = &autoscalingv1.Scale{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("clonesets").
		Name(cloneSetName).
		SubResource("scale").
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// UpdateScale takes the top resource name and the representation of a scale and updates it. Returns the server's representation of the scale, and an error, if there is any.
func (c *cloneSets) UpdateScale(ctx context.Context, cloneSetName string, scale *autoscalingv1.Scale, opts v1.UpdateOptions) (result *autoscalingv1.Scale, err error) {
	result = &autoscalingv1.Scale{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("clonesets").
		Name(cloneSetName).
		SubResource("scale").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(scale).
		Do(ctx).
		Into(result)
	return
}
