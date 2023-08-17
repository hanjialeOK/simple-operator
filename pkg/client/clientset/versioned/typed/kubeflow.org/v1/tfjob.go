/*
Copyright 2023 hanjialeok.

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

package v1

import (
	"context"
	"time"

	v1 "github.com/hanjialeok/simple-operator/pkg/apis/kubeflow.org/v1"
	scheme "github.com/hanjialeok/simple-operator/pkg/client/clientset/versioned/scheme"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
)

// TFJobsGetter has a method to return a TFJobInterface.
// A group's client should implement this interface.
type TFJobsGetter interface {
	TFJobs(namespace string) TFJobInterface
}

// TFJobInterface has methods to work with TFJob resources.
type TFJobInterface interface {
	Create(ctx context.Context, tFJob *v1.TFJob, opts metav1.CreateOptions) (*v1.TFJob, error)
	Update(ctx context.Context, tFJob *v1.TFJob, opts metav1.UpdateOptions) (*v1.TFJob, error)
	UpdateStatus(ctx context.Context, tFJob *v1.TFJob, opts metav1.UpdateOptions) (*v1.TFJob, error)
	Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error
	Get(ctx context.Context, name string, opts metav1.GetOptions) (*v1.TFJob, error)
	List(ctx context.Context, opts metav1.ListOptions) (*v1.TFJobList, error)
	Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.TFJob, err error)
	TFJobExpansion
}

// tFJobs implements TFJobInterface
type tFJobs struct {
	client rest.Interface
	ns     string
}

// newTFJobs returns a TFJobs
func newTFJobs(c *KubeflowV1Client, namespace string) *tFJobs {
	return &tFJobs{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the tFJob, and returns the corresponding tFJob object, and an error if there is any.
func (c *tFJobs) Get(ctx context.Context, name string, options metav1.GetOptions) (result *v1.TFJob, err error) {
	result = &v1.TFJob{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("tfjobs").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do(ctx).
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of TFJobs that match those selectors.
func (c *tFJobs) List(ctx context.Context, opts metav1.ListOptions) (result *v1.TFJobList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1.TFJobList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("tfjobs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do(ctx).
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested tFJobs.
func (c *tFJobs) Watch(ctx context.Context, opts metav1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("tfjobs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch(ctx)
}

// Create takes the representation of a tFJob and creates it.  Returns the server's representation of the tFJob, and an error, if there is any.
func (c *tFJobs) Create(ctx context.Context, tFJob *v1.TFJob, opts metav1.CreateOptions) (result *v1.TFJob, err error) {
	result = &v1.TFJob{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("tfjobs").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(tFJob).
		Do(ctx).
		Into(result)
	return
}

// Update takes the representation of a tFJob and updates it. Returns the server's representation of the tFJob, and an error, if there is any.
func (c *tFJobs) Update(ctx context.Context, tFJob *v1.TFJob, opts metav1.UpdateOptions) (result *v1.TFJob, err error) {
	result = &v1.TFJob{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("tfjobs").
		Name(tFJob.Name).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(tFJob).
		Do(ctx).
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *tFJobs) UpdateStatus(ctx context.Context, tFJob *v1.TFJob, opts metav1.UpdateOptions) (result *v1.TFJob, err error) {
	result = &v1.TFJob{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("tfjobs").
		Name(tFJob.Name).
		SubResource("status").
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(tFJob).
		Do(ctx).
		Into(result)
	return
}

// Delete takes name of the tFJob and deletes it. Returns an error if one occurs.
func (c *tFJobs) Delete(ctx context.Context, name string, opts metav1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("tfjobs").
		Name(name).
		Body(&opts).
		Do(ctx).
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *tFJobs) DeleteCollection(ctx context.Context, opts metav1.DeleteOptions, listOpts metav1.ListOptions) error {
	var timeout time.Duration
	if listOpts.TimeoutSeconds != nil {
		timeout = time.Duration(*listOpts.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("tfjobs").
		VersionedParams(&listOpts, scheme.ParameterCodec).
		Timeout(timeout).
		Body(&opts).
		Do(ctx).
		Error()
}

// Patch applies the patch and returns the patched tFJob.
func (c *tFJobs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts metav1.PatchOptions, subresources ...string) (result *v1.TFJob, err error) {
	result = &v1.TFJob{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("tfjobs").
		Name(name).
		SubResource(subresources...).
		VersionedParams(&opts, scheme.ParameterCodec).
		Body(data).
		Do(ctx).
		Into(result)
	return
}
