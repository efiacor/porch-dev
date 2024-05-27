// Copyright 2023-2024 The kpt and Nephio Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by client-gen. DO NOT EDIT.

package v1alpha1

import (
	"context"
	"time"

	scheme "github.com/nephio-project/porch/api/generated/clientset/versioned/scheme"
	v1alpha1 "github.com/nephio-project/porch/api/porch/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// PackageRevisionResourcesGetter has a method to return a PackageRevisionResourcesInterface.
// A group's client should implement this interface.
type PackageRevisionResourcesGetter interface {
	PackageRevisionResources(namespace string) PackageRevisionResourcesInterface
}

// PackageRevisionResourcesInterface has methods to work with PackageRevisionResources resources.
type PackageRevisionResourcesInterface interface {
	Create(ctx context.Context, packageRevisionResources *v1alpha1.PackageRevisionResources, opts v1.CreateOptions) (*v1alpha1.PackageRevisionResources, error)
	Update(ctx context.Context, packageRevisionResources *v1alpha1.PackageRevisionResources, opts v1.UpdateOptions) (*v1alpha1.PackageRevisionResources, error)
	UpdateStatus(ctx context.Context, packageRevisionResources *v1alpha1.PackageRevisionResources, opts v1.UpdateOptions) (*v1alpha1.PackageRevisionResources, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v1alpha1.PackageRevisionResources, error)
	List(ctx context.Context, opts v1.ListOptions) (*v1alpha1.PackageRevisionResourcesList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PackageRevisionResources, err error)
	PackageRevisionResourcesExpansion
}

// packageRevisionResources implements PackageRevisionResourcesInterface
type packageRevisionResources struct {
	client rest.Interface
	ns     string
}

// newPackageRevisionResources returns a PackageRevisionResources
func newPackageRevisionResources(c *PorchV1alpha1Client, namespace string) *packageRevisionResources {
	return &packageRevisionResources{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the packageRevisionResources, and returns the corresponding packageRevisionResources object, and an error if there is any.
func (c *packageRevisionResources) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1alpha1.PackageRevisionResources, err error) {
	result = &v1alpha1.PackageRevisionResources{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("packagerevisionresources").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of PackageRevisionResources that match those selectors.
func (c *packageRevisionResources) List(ctx context.Context, opts v1.ListOptions) (result *v1alpha1.PackageRevisionResourcesList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.PackageRevisionResourcesList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("packagerevisionresources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested packageRevisionResources.
func (c *packageRevisionResources) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("packagerevisionresources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a packageRevisionResources and creates it.  Returns the server's representation of the packageRevisionResources, and an error, if there is any.
func (c *packageRevisionResources) Create(ctx context.Context, packageRevisionResources *v1alpha1.PackageRevisionResources, opts v1.CreateOptions) (result *v1alpha1.PackageRevisionResources, err error) {
	result = &v1alpha1.PackageRevisionResources{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("packagerevisionresources").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(packageRevisionResources).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a packageRevisionResources and updates it. Returns the server's representation of the packageRevisionResources, and an error, if there is any.
func (c *packageRevisionResources) Update(ctx context.Context, packageRevisionResources *v1alpha1.PackageRevisionResources, opts v1.UpdateOptions) (result *v1alpha1.PackageRevisionResources, err error) {
	result = &v1alpha1.PackageRevisionResources{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("packagerevisionresources").
		Name(packageRevisionResources.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(packageRevisionResources).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *packageRevisionResources) UpdateStatus(ctx context.Context, packageRevisionResources *v1alpha1.PackageRevisionResources, opts v1.UpdateOptions) (result *v1alpha1.PackageRevisionResources, err error) {
	result = &v1alpha1.PackageRevisionResources{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("packagerevisionresources").
		Name(packageRevisionResources.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(packageRevisionResources).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the packageRevisionResources and deletes it. Returns an error if one occurs.
func (c *packageRevisionResources) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("packagerevisionresources").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *packageRevisionResources) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("packagerevisionresources").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched packageRevisionResources.
func (c *packageRevisionResources) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1alpha1.PackageRevisionResources, err error) {
	result = &v1alpha1.PackageRevisionResources{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("packagerevisionresources").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
