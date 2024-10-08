// Copyright 2020 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// *** DISCLAIMER ***
// Config Connector's go-client for CRDs is currently in ALPHA, which means
// that future versions of the go-client may include breaking changes.
// Please try it out and give us feedback!

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/compute/v1beta1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeComputeRegionNetworkEndpointGroups implements ComputeRegionNetworkEndpointGroupInterface
type FakeComputeRegionNetworkEndpointGroups struct {
	Fake *FakeComputeV1beta1
	ns   string
}

var computeregionnetworkendpointgroupsResource = v1beta1.SchemeGroupVersion.WithResource("computeregionnetworkendpointgroups")

var computeregionnetworkendpointgroupsKind = v1beta1.SchemeGroupVersion.WithKind("ComputeRegionNetworkEndpointGroup")

// Get takes name of the computeRegionNetworkEndpointGroup, and returns the corresponding computeRegionNetworkEndpointGroup object, and an error if there is any.
func (c *FakeComputeRegionNetworkEndpointGroups) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.ComputeRegionNetworkEndpointGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(computeregionnetworkendpointgroupsResource, c.ns, name), &v1beta1.ComputeRegionNetworkEndpointGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ComputeRegionNetworkEndpointGroup), err
}

// List takes label and field selectors, and returns the list of ComputeRegionNetworkEndpointGroups that match those selectors.
func (c *FakeComputeRegionNetworkEndpointGroups) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.ComputeRegionNetworkEndpointGroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(computeregionnetworkendpointgroupsResource, computeregionnetworkendpointgroupsKind, c.ns, opts), &v1beta1.ComputeRegionNetworkEndpointGroupList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.ComputeRegionNetworkEndpointGroupList{ListMeta: obj.(*v1beta1.ComputeRegionNetworkEndpointGroupList).ListMeta}
	for _, item := range obj.(*v1beta1.ComputeRegionNetworkEndpointGroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested computeRegionNetworkEndpointGroups.
func (c *FakeComputeRegionNetworkEndpointGroups) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(computeregionnetworkendpointgroupsResource, c.ns, opts))

}

// Create takes the representation of a computeRegionNetworkEndpointGroup and creates it.  Returns the server's representation of the computeRegionNetworkEndpointGroup, and an error, if there is any.
func (c *FakeComputeRegionNetworkEndpointGroups) Create(ctx context.Context, computeRegionNetworkEndpointGroup *v1beta1.ComputeRegionNetworkEndpointGroup, opts v1.CreateOptions) (result *v1beta1.ComputeRegionNetworkEndpointGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(computeregionnetworkendpointgroupsResource, c.ns, computeRegionNetworkEndpointGroup), &v1beta1.ComputeRegionNetworkEndpointGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ComputeRegionNetworkEndpointGroup), err
}

// Update takes the representation of a computeRegionNetworkEndpointGroup and updates it. Returns the server's representation of the computeRegionNetworkEndpointGroup, and an error, if there is any.
func (c *FakeComputeRegionNetworkEndpointGroups) Update(ctx context.Context, computeRegionNetworkEndpointGroup *v1beta1.ComputeRegionNetworkEndpointGroup, opts v1.UpdateOptions) (result *v1beta1.ComputeRegionNetworkEndpointGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(computeregionnetworkendpointgroupsResource, c.ns, computeRegionNetworkEndpointGroup), &v1beta1.ComputeRegionNetworkEndpointGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ComputeRegionNetworkEndpointGroup), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeComputeRegionNetworkEndpointGroups) UpdateStatus(ctx context.Context, computeRegionNetworkEndpointGroup *v1beta1.ComputeRegionNetworkEndpointGroup, opts v1.UpdateOptions) (*v1beta1.ComputeRegionNetworkEndpointGroup, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(computeregionnetworkendpointgroupsResource, "status", c.ns, computeRegionNetworkEndpointGroup), &v1beta1.ComputeRegionNetworkEndpointGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ComputeRegionNetworkEndpointGroup), err
}

// Delete takes name of the computeRegionNetworkEndpointGroup and deletes it. Returns an error if one occurs.
func (c *FakeComputeRegionNetworkEndpointGroups) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteActionWithOptions(computeregionnetworkendpointgroupsResource, c.ns, name, opts), &v1beta1.ComputeRegionNetworkEndpointGroup{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeComputeRegionNetworkEndpointGroups) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(computeregionnetworkendpointgroupsResource, c.ns, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.ComputeRegionNetworkEndpointGroupList{})
	return err
}

// Patch applies the patch and returns the patched computeRegionNetworkEndpointGroup.
func (c *FakeComputeRegionNetworkEndpointGroups) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ComputeRegionNetworkEndpointGroup, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(computeregionnetworkendpointgroupsResource, c.ns, name, pt, data, subresources...), &v1beta1.ComputeRegionNetworkEndpointGroup{})

	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.ComputeRegionNetworkEndpointGroup), err
}
