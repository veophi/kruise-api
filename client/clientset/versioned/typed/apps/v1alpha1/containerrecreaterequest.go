/*
Copyright 2020 The Kruise Authors.

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
	"time"

	v1alpha1 "github.com/openkruise/kruise-api/apps/v1alpha1"
	scheme "github.com/openkruise/kruise-api/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ContainerRecreateRequestsGetter has a method to return a ContainerRecreateRequestInterface.
// A group's client should implement this interface.
type ContainerRecreateRequestsGetter interface {
	ContainerRecreateRequests(namespace string) ContainerRecreateRequestInterface
}

// ContainerRecreateRequestInterface has methods to work with ContainerRecreateRequest resources.
type ContainerRecreateRequestInterface interface {
	Create(*v1alpha1.ContainerRecreateRequest) (*v1alpha1.ContainerRecreateRequest, error)
	Update(*v1alpha1.ContainerRecreateRequest) (*v1alpha1.ContainerRecreateRequest, error)
	UpdateStatus(*v1alpha1.ContainerRecreateRequest) (*v1alpha1.ContainerRecreateRequest, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.ContainerRecreateRequest, error)
	List(opts v1.ListOptions) (*v1alpha1.ContainerRecreateRequestList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ContainerRecreateRequest, err error)
	ContainerRecreateRequestExpansion
}

// containerRecreateRequests implements ContainerRecreateRequestInterface
type containerRecreateRequests struct {
	client rest.Interface
	ns     string
}

// newContainerRecreateRequests returns a ContainerRecreateRequests
func newContainerRecreateRequests(c *AppsV1alpha1Client, namespace string) *containerRecreateRequests {
	return &containerRecreateRequests{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the containerRecreateRequest, and returns the corresponding containerRecreateRequest object, and an error if there is any.
func (c *containerRecreateRequests) Get(name string, options v1.GetOptions) (result *v1alpha1.ContainerRecreateRequest, err error) {
	result = &v1alpha1.ContainerRecreateRequest{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("containerrecreaterequests").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ContainerRecreateRequests that match those selectors.
func (c *containerRecreateRequests) List(opts v1.ListOptions) (result *v1alpha1.ContainerRecreateRequestList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.ContainerRecreateRequestList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("containerrecreaterequests").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested containerRecreateRequests.
func (c *containerRecreateRequests) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("containerrecreaterequests").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a containerRecreateRequest and creates it.  Returns the server's representation of the containerRecreateRequest, and an error, if there is any.
func (c *containerRecreateRequests) Create(containerRecreateRequest *v1alpha1.ContainerRecreateRequest) (result *v1alpha1.ContainerRecreateRequest, err error) {
	result = &v1alpha1.ContainerRecreateRequest{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("containerrecreaterequests").
		Body(containerRecreateRequest).
		Do().
		Into(result)
	return
}

// Update takes the representation of a containerRecreateRequest and updates it. Returns the server's representation of the containerRecreateRequest, and an error, if there is any.
func (c *containerRecreateRequests) Update(containerRecreateRequest *v1alpha1.ContainerRecreateRequest) (result *v1alpha1.ContainerRecreateRequest, err error) {
	result = &v1alpha1.ContainerRecreateRequest{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("containerrecreaterequests").
		Name(containerRecreateRequest.Name).
		Body(containerRecreateRequest).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *containerRecreateRequests) UpdateStatus(containerRecreateRequest *v1alpha1.ContainerRecreateRequest) (result *v1alpha1.ContainerRecreateRequest, err error) {
	result = &v1alpha1.ContainerRecreateRequest{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("containerrecreaterequests").
		Name(containerRecreateRequest.Name).
		SubResource("status").
		Body(containerRecreateRequest).
		Do().
		Into(result)
	return
}

// Delete takes name of the containerRecreateRequest and deletes it. Returns an error if one occurs.
func (c *containerRecreateRequests) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("containerrecreaterequests").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *containerRecreateRequests) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("containerrecreaterequests").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched containerRecreateRequest.
func (c *containerRecreateRequests) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.ContainerRecreateRequest, err error) {
	result = &v1alpha1.ContainerRecreateRequest{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("containerrecreaterequests").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
