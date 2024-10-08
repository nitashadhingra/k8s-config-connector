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

package v1beta1

import (
	"context"
	"time"

	v1beta1 "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/apis/compute/v1beta1"
	scheme "github.com/GoogleCloudPlatform/k8s-config-connector/pkg/clients/generated/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// ComputeTargetTCPProxiesGetter has a method to return a ComputeTargetTCPProxyInterface.
// A group's client should implement this interface.
type ComputeTargetTCPProxiesGetter interface {
	ComputeTargetTCPProxies(namespace string) ComputeTargetTCPProxyInterface
}

// ComputeTargetTCPProxyInterface has methods to work with ComputeTargetTCPProxy resources.
type ComputeTargetTCPProxyInterface interface {
	Create(ctx context.Context, computeTargetTCPProxy *v1beta1.ComputeTargetTCPProxy, opts v1.CreateOptions) (*v1beta1.ComputeTargetTCPProxy, error)
	Update(ctx context.Context, computeTargetTCPProxy *v1beta1.ComputeTargetTCPProxy, opts v1.UpdateOptions) (*v1beta1.ComputeTargetTCPProxy, error)
	UpdateStatus(ctx context.Context, computeTargetTCPProxy *v1beta1.ComputeTargetTCPProxy, opts v1.UpdateOptions) (*v1beta1.ComputeTargetTCPProxy, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1beta1.ComputeTargetTCPProxy, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1beta1.ComputeTargetTCPProxyList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ComputeTargetTCPProxy, err error)
	ComputeTargetTCPProxyExpansion
}

// computeTargetTCPProxies implements ComputeTargetTCPProxyInterface
type computeTargetTCPProxies struct {
	client rest.Interface
	ns     string
}

// newComputeTargetTCPProxies returns a ComputeTargetTCPProxies
func newComputeTargetTCPProxies(c *ComputeV1beta1Client, namespace string) *computeTargetTCPProxies {
	return &computeTargetTCPProxies{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the computeTargetTCPProxy, and returns the corresponding computeTargetTCPProxy object, and an error if there is any.
func (c *computeTargetTCPProxies) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.ComputeTargetTCPProxy, err error) {
	result = &v1beta1.ComputeTargetTCPProxy{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("computetargettcpproxies").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of ComputeTargetTCPProxies that match those selectors.
func (c *computeTargetTCPProxies) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.ComputeTargetTCPProxyList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1beta1.ComputeTargetTCPProxyList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("computetargettcpproxies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested computeTargetTCPProxies.
func (c *computeTargetTCPProxies) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("computetargettcpproxies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a computeTargetTCPProxy and creates it.  Returns the server's representation of the computeTargetTCPProxy, and an error, if there is any.
func (c *computeTargetTCPProxies) Create(ctx context.Context, computeTargetTCPProxy *v1beta1.ComputeTargetTCPProxy, opts v1.CreateOptions) (result *v1beta1.ComputeTargetTCPProxy, err error) {
	result = &v1beta1.ComputeTargetTCPProxy{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("computetargettcpproxies").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(computeTargetTCPProxy).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a computeTargetTCPProxy and updates it. Returns the server's representation of the computeTargetTCPProxy, and an error, if there is any.
func (c *computeTargetTCPProxies) Update(ctx context.Context, computeTargetTCPProxy *v1beta1.ComputeTargetTCPProxy, opts v1.UpdateOptions) (result *v1beta1.ComputeTargetTCPProxy, err error) {
	result = &v1beta1.ComputeTargetTCPProxy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("computetargettcpproxies").
		Name(computeTargetTCPProxy.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(computeTargetTCPProxy).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *computeTargetTCPProxies) UpdateStatus(ctx context.Context, computeTargetTCPProxy *v1beta1.ComputeTargetTCPProxy, opts v1.UpdateOptions) (result *v1beta1.ComputeTargetTCPProxy, err error) {
	result = &v1beta1.ComputeTargetTCPProxy{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("computetargettcpproxies").
		Name(computeTargetTCPProxy.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(computeTargetTCPProxy).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the computeTargetTCPProxy and deletes it. Returns an error if one occurs.
func (c *computeTargetTCPProxies) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("computetargettcpproxies").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *computeTargetTCPProxies) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("computetargettcpproxies").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched computeTargetTCPProxy.
func (c *computeTargetTCPProxies) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.ComputeTargetTCPProxy, err error) {
	result = &v1beta1.ComputeTargetTCPProxy{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("computetargettcpproxies").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
