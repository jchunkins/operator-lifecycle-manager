/*
Copyright Red Hat, Inc.

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

package v2

import (
	"context"
	"time"

	v2 "github.com/operator-framework/api/pkg/operators/v2"
	scheme "github.com/operator-framework/operator-lifecycle-manager/pkg/api/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// OperatorConditionsGetter has a method to return a OperatorConditionInterface.
// A group's client should implement this interface.
type OperatorConditionsGetter interface {
	OperatorConditions(namespace string) OperatorConditionInterface
}

// OperatorConditionInterface has methods to work with OperatorCondition resources.
type OperatorConditionInterface interface {
	Create(ctx context.Context, operatorCondition *v2.OperatorCondition, opts v1.CreateOptions) (*v2.OperatorCondition, error)
	Update(ctx context.Context, operatorCondition *v2.OperatorCondition, opts v1.UpdateOptions) (*v2.OperatorCondition, error)
	UpdateStatus(ctx context.Context, operatorCondition *v2.OperatorCondition, opts v1.UpdateOptions) (*v2.OperatorCondition, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*v2.OperatorCondition, error)
	List(ctx context.Context, opts v1.ListOptions) (*v2.OperatorConditionList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2.OperatorCondition, err error)
	OperatorConditionExpansion
}

// operatorConditions implements OperatorConditionInterface
type operatorConditions struct {
	client rest.Interface
	ns     string
}

// newOperatorConditions returns a OperatorConditions
func newOperatorConditions(c *OperatorsV2Client, namespace string) *operatorConditions {
	return &operatorConditions{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the operatorCondition, and returns the corresponding operatorCondition object, and an error if there is any.
func (c *operatorConditions) Get(ctx context.Context, name string, options v1.GetOptions) (result *v2.OperatorCondition, err error) {
	result = &v2.OperatorCondition{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("operatorconditions").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of OperatorConditions that match those selectors.
func (c *operatorConditions) List(ctx context.Context, opts v1.ListOptions) (result *v2.OperatorConditionList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v2.OperatorConditionList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("operatorconditions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested operatorConditions.
func (c *operatorConditions) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("operatorconditions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a operatorCondition and creates it.  Returns the server's representation of the operatorCondition, and an error, if there is any.
func (c *operatorConditions) Create(ctx context.Context, operatorCondition *v2.OperatorCondition, opts v1.CreateOptions) (result *v2.OperatorCondition, err error) {
	result = &v2.OperatorCondition{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("operatorconditions").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(operatorCondition).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a operatorCondition and updates it. Returns the server's representation of the operatorCondition, and an error, if there is any.
func (c *operatorConditions) Update(ctx context.Context, operatorCondition *v2.OperatorCondition, opts v1.UpdateOptions) (result *v2.OperatorCondition, err error) {
	result = &v2.OperatorCondition{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("operatorconditions").
		Name(operatorCondition.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(operatorCondition).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *operatorConditions) UpdateStatus(ctx context.Context, operatorCondition *v2.OperatorCondition, opts v1.UpdateOptions) (result *v2.OperatorCondition, err error) {
	result = &v2.OperatorCondition{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("operatorconditions").
		Name(operatorCondition.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(operatorCondition).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the operatorCondition and deletes it. Returns an error if one occurs.
func (c *operatorConditions) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("operatorconditions").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *operatorConditions) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("operatorconditions").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched operatorCondition.
func (c *operatorConditions) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v2.OperatorCondition, err error) {
	result = &v2.OperatorCondition{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("operatorconditions").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}