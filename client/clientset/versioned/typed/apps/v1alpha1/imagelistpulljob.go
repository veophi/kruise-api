/*
Copyright 2022 The Kruise Authors.

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

	v1alpha1 "github.com/openkruise/kruise-api/apps/v1alpha1"
	scheme "github.com/openkruise/kruise-api/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ImageListPullJobsGetter has a method to return a ImageListPullJobInterface.
// A group's client should implement this interface.
type ImageListPullJobsGetter interface {
	ImageListPullJobs(namespace string) ImageListPullJobInterface
}

// ImageListPullJobInterface has methods to work with ImageListPullJob resources.
type ImageListPullJobInterface interface {
	Create(ctx context.Context, imageListPullJob *v1alpha1.ImageListPullJob, opts v1.CreateOptions) (*v1alpha1.ImageListPullJob, error)
	Update(ctx context.Context, imageListPullJob *v1alpha1.ImageListPullJob, opts v1.UpdateOptions) (*v1alpha1.ImageListPullJob, error)
	UpdateStatus(ctx context.Context, imageListPullJob *v1alpha1.ImageListPullJob, opts v1.UpdateOptions) (*v1alpha1.ImageListPullJob, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.ImageListPullJob, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.ImageListPullJobList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ImageListPullJob, err error)
	ImageListPullJobExpansion
}

// imageListPullJobs implements ImageListPullJobInterface
type imageListPullJobs struct {
	client rest.Interface
	ns     string
}

// newImageListPullJobs returns a ImageListPullJobs
func newImageListPullJobs(c *AppsV1alpha1Client, namespace string) *imageListPullJobs {
	return &imageListPullJobs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the imageListPullJob, and returns the corresponding imageListPullJob object, and an error if there is any.
func (c *imageListPullJobs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.ImageListPullJob, err error) {
	result = &v1alpha1.ImageListPullJob{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("imagelistpulljobs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ImageListPullJobs that match those selectors.
func (c *imageListPullJobs) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.ImageListPullJobList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ImageListPullJobList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("imagelistpulljobs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested imageListPullJobs.
func (c *imageListPullJobs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("imagelistpulljobs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a imageListPullJob and creates it.  Returns the server's representation of the imageListPullJob, and an error, if there is any.
func (c *imageListPullJobs) Create(ctx context.Context, imageListPullJob *v1alpha1.ImageListPullJob, opts v1.CreateOptions) (result *v1alpha1.ImageListPullJob, err error) {
	result = &v1alpha1.ImageListPullJob{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("imagelistpulljobs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(imageListPullJob).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a imageListPullJob and updates it. Returns the server's representation of the imageListPullJob, and an error, if there is any.
func (c *imageListPullJobs) Update(ctx context.Context, imageListPullJob *v1alpha1.ImageListPullJob, opts v1.UpdateOptions) (result *v1alpha1.ImageListPullJob, err error) {
	result = &v1alpha1.ImageListPullJob{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("imagelistpulljobs").
		Name(imageListPullJob.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(imageListPullJob).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *imageListPullJobs) UpdateStatus(ctx context.Context, imageListPullJob *v1alpha1.ImageListPullJob, opts v1.UpdateOptions) (result *v1alpha1.ImageListPullJob, err error) {
	result = &v1alpha1.ImageListPullJob{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("imagelistpulljobs").
		Name(imageListPullJob.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(imageListPullJob).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the imageListPullJob and deletes it. Returns an error if one occurs.
func (c *imageListPullJobs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("imagelistpulljobs").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *imageListPullJobs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("imagelistpulljobs").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched imageListPullJob.
func (c *imageListPullJobs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.ImageListPullJob, err error) {
	result = &v1alpha1.ImageListPullJob{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("imagelistpulljobs").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
