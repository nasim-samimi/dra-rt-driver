/*
 * Copyright 2024 The Kubernetes Authors.
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

package fake

import (
	"context"

	v1alpha1 "github.com/nasim-samimi/dra-rt-driver/api/example.com/resource/rt/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeRtClaimParameters implements RtClaimParametersInterface
type FakeRtClaimParameters struct {
	Fake *FakeRtV1alpha1
	ns   string
}

var rtclaimparametersResource = v1alpha1.SchemeGroupVersion.WithResource("rtclaimparameters")

var rtclaimparametersKind = v1alpha1.SchemeGroupVersion.WithKind("RtClaimParameters")

// Get takes name of the rtClaimParameters, and returns the corresponding rtClaimParameters object, and an error if there is any.
func (c *FakeRtClaimParameters) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.RtClaimParameters, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(rtclaimparametersResource, c.ns, name), &v1alpha1.RtClaimParameters{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RtClaimParameters), err
}

// List takes label and field selectors, and returns the list of RtClaimParameters that match those selectors.
func (c *FakeRtClaimParameters) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.RtClaimParametersList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(rtclaimparametersResource, rtclaimparametersKind, c.ns, opts), &v1alpha1.RtClaimParametersList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1alpha1.RtClaimParametersList{ListMeta: obj.(*v1alpha1.RtClaimParametersList).ListMeta}
	for _, item := range obj.(*v1alpha1.RtClaimParametersList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested rtClaimParameters.
func (c *FakeRtClaimParameters) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(rtclaimparametersResource, c.ns, opts))

}

// Create takes the representation of a rtClaimParameters and creates it.  Returns the server's representation of the rtClaimParameters, and an error, if there is any.
func (c *FakeRtClaimParameters) Create(ctx context.Context, rtClaimParameters *v1alpha1.RtClaimParameters, opts v1.CreateOptions) (result *v1alpha1.RtClaimParameters, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(rtclaimparametersResource, c.ns, rtClaimParameters), &v1alpha1.RtClaimParameters{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RtClaimParameters), err
}

// Update takes the representation of a rtClaimParameters and updates it. Returns the server's representation of the rtClaimParameters, and an error, if there is any.
func (c *FakeRtClaimParameters) Update(ctx context.Context, rtClaimParameters *v1alpha1.RtClaimParameters, opts v1.UpdateOptions) (result *v1alpha1.RtClaimParameters, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(rtclaimparametersResource, c.ns, rtClaimParameters), &v1alpha1.RtClaimParameters{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RtClaimParameters), err
}

// Delete takes name of the rtClaimParameters and deletes it. Returns an error if one occurs.
func (c *FakeRtClaimParameters) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(rtclaimparametersResource, c.ns, name, opts), &v1alpha1.RtClaimParameters{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeRtClaimParameters) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(rtclaimparametersResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1alpha1.RtClaimParametersList{})
	return err
}

// Patch applies the patch and returns the patched rtClaimParameters.
func (c *FakeRtClaimParameters) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.RtClaimParameters, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(rtclaimparametersResource, c.ns, name, pt, data, subresources...), &v1alpha1.RtClaimParameters{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1alpha1.RtClaimParameters), err
}
