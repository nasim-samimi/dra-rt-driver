/*
 * Copyright 2025 The Kubernetes Authors.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	v1alpha1 "github.com/nasim-samimi/dra-rt-driver/api/example.com/resource/rt/v1alpha1"
	scheme "github.com/nasim-samimi/dra-rt-driver/pkg/example.com/resource/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// RtClaimParametersGetter has a method to return a RtClaimParametersInterface.
// A group's client should implement this interface.
type RtClaimParametersGetter interface {
	RtClaimParameters(namespace string) RtClaimParametersInterface
}

// RtClaimParametersInterface has methods to work with RtClaimParameters resources.
type RtClaimParametersInterface interface {
	Create(ctx context.Context, rtClaimParameters *v1alpha1.RtClaimParameters, opts v1.CreateOptions) (*v1alpha1.RtClaimParameters, error)
	Update(ctx context.Context, rtClaimParameters *v1alpha1.RtClaimParameters, opts v1.UpdateOptions) (*v1alpha1.RtClaimParameters, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.RtClaimParameters, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.RtClaimParametersList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.RtClaimParameters, err error)
	RtClaimParametersExpansion
}

// rtClaimParameters implements RtClaimParametersInterface
type rtClaimParameters struct {
	client rest.Interface
	ns     string
}

// newRtClaimParameters returns a RtClaimParameters
func newRtClaimParameters(c *RtV1alpha1Client, namespace string) *rtClaimParameters {
	return &rtClaimParameters{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the rtClaimParameters, and returns the corresponding rtClaimParameters object, and an error if there is any.
func (c *rtClaimParameters) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.RtClaimParameters, err error) {
	result = &v1alpha1.RtClaimParameters{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("rtclaimparameters").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of RtClaimParameters that match those selectors.
func (c *rtClaimParameters) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.RtClaimParametersList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.RtClaimParametersList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("rtclaimparameters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested rtClaimParameters.
func (c *rtClaimParameters) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("rtclaimparameters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a rtClaimParameters and creates it.  Returns the server's representation of the rtClaimParameters, and an error, if there is any.
func (c *rtClaimParameters) Create(ctx context.Context, rtClaimParameters *v1alpha1.RtClaimParameters, opts v1.CreateOptions) (result *v1alpha1.RtClaimParameters, err error) {
	result = &v1alpha1.RtClaimParameters{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("rtclaimparameters").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(rtClaimParameters).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a rtClaimParameters and updates it. Returns the server's representation of the rtClaimParameters, and an error, if there is any.
func (c *rtClaimParameters) Update(ctx context.Context, rtClaimParameters *v1alpha1.RtClaimParameters, opts v1.UpdateOptions) (result *v1alpha1.RtClaimParameters, err error) {
	result = &v1alpha1.RtClaimParameters{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("rtclaimparameters").
		Name(rtClaimParameters.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(rtClaimParameters).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the rtClaimParameters and deletes it. Returns an error if one occurs.
func (c *rtClaimParameters) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("rtclaimparameters").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *rtClaimParameters) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("rtclaimparameters").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched rtClaimParameters.
func (c *rtClaimParameters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.RtClaimParameters, err error) {
	result = &v1alpha1.RtClaimParameters{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("rtclaimparameters").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
