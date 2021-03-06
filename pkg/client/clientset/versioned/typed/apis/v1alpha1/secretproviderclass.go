/*
Copyright 2020 The Kubernetes Authors.

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

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	rest "k8s.io/client-go/rest"
	v1alpha1 "sigs.k8s.io/secrets-store-csi-driver/apis/v1alpha1"
	scheme "sigs.k8s.io/secrets-store-csi-driver/pkg/client/clientset/versioned/scheme"
)

// SecretProviderClassesGetter has a method to return a SecretProviderClassInterface.
// A group's client should implement this interface.
type SecretProviderClassesGetter interface {
	SecretProviderClasses(namespace string) SecretProviderClassInterface
}

// SecretProviderClassInterface has methods to work with SecretProviderClass resources.
type SecretProviderClassInterface interface {
	Create(*v1alpha1.SecretProviderClass) (*v1alpha1.SecretProviderClass, error)
	Update(*v1alpha1.SecretProviderClass) (*v1alpha1.SecretProviderClass, error)
	UpdateStatus(*v1alpha1.SecretProviderClass) (*v1alpha1.SecretProviderClass, error)
	Delete(name string, options *v1.DeleteOptions) error
	DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error
	Get(name string, options v1.GetOptions) (*v1alpha1.SecretProviderClass, error)
	List(opts v1.ListOptions) (*v1alpha1.SecretProviderClassList, error)
	Watch(opts v1.ListOptions) (watch.Interface, error)
	Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SecretProviderClass, err error)
	SecretProviderClassExpansion
}

// secretProviderClasses implements SecretProviderClassInterface
type secretProviderClasses struct {
	client rest.Interface
	ns     string
}

// newSecretProviderClasses returns a SecretProviderClasses
func newSecretProviderClasses(c *SecretsstoreV1alpha1Client, namespace string) *secretProviderClasses {
	return &secretProviderClasses{
		client: c.RESTClient(),
		ns:     namespace,
	}
}

// Get takes name of the secretProviderClass, and returns the corresponding secretProviderClass object, and an error if there is any.
func (c *secretProviderClasses) Get(name string, options v1.GetOptions) (result *v1alpha1.SecretProviderClass, err error) {
	result = &v1alpha1.SecretProviderClass{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("secretproviderclasses").
		Name(name).
		VersionedParams(&options, scheme.ParameterCodec).
		Do().
		Into(result)
	return
}

// List takes label and field selectors, and returns the list of SecretProviderClasses that match those selectors.
func (c *secretProviderClasses) List(opts v1.ListOptions) (result *v1alpha1.SecretProviderClassList, err error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	result = &v1alpha1.SecretProviderClassList{}
	err = c.client.Get().
		Namespace(c.ns).
		Resource("secretproviderclasses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Do().
		Into(result)
	return
}

// Watch returns a watch.Interface that watches the requested secretProviderClasses.
func (c *secretProviderClasses) Watch(opts v1.ListOptions) (watch.Interface, error) {
	var timeout time.Duration
	if opts.TimeoutSeconds != nil {
		timeout = time.Duration(*opts.TimeoutSeconds) * time.Second
	}
	opts.Watch = true
	return c.client.Get().
		Namespace(c.ns).
		Resource("secretproviderclasses").
		VersionedParams(&opts, scheme.ParameterCodec).
		Timeout(timeout).
		Watch()
}

// Create takes the representation of a secretProviderClass and creates it.  Returns the server's representation of the secretProviderClass, and an error, if there is any.
func (c *secretProviderClasses) Create(secretProviderClass *v1alpha1.SecretProviderClass) (result *v1alpha1.SecretProviderClass, err error) {
	result = &v1alpha1.SecretProviderClass{}
	err = c.client.Post().
		Namespace(c.ns).
		Resource("secretproviderclasses").
		Body(secretProviderClass).
		Do().
		Into(result)
	return
}

// Update takes the representation of a secretProviderClass and updates it. Returns the server's representation of the secretProviderClass, and an error, if there is any.
func (c *secretProviderClasses) Update(secretProviderClass *v1alpha1.SecretProviderClass) (result *v1alpha1.SecretProviderClass, err error) {
	result = &v1alpha1.SecretProviderClass{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("secretproviderclasses").
		Name(secretProviderClass.Name).
		Body(secretProviderClass).
		Do().
		Into(result)
	return
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().

func (c *secretProviderClasses) UpdateStatus(secretProviderClass *v1alpha1.SecretProviderClass) (result *v1alpha1.SecretProviderClass, err error) {
	result = &v1alpha1.SecretProviderClass{}
	err = c.client.Put().
		Namespace(c.ns).
		Resource("secretproviderclasses").
		Name(secretProviderClass.Name).
		SubResource("status").
		Body(secretProviderClass).
		Do().
		Into(result)
	return
}

// Delete takes name of the secretProviderClass and deletes it. Returns an error if one occurs.
func (c *secretProviderClasses) Delete(name string, options *v1.DeleteOptions) error {
	return c.client.Delete().
		Namespace(c.ns).
		Resource("secretproviderclasses").
		Name(name).
		Body(options).
		Do().
		Error()
}

// DeleteCollection deletes a collection of objects.
func (c *secretProviderClasses) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	var timeout time.Duration
	if listOptions.TimeoutSeconds != nil {
		timeout = time.Duration(*listOptions.TimeoutSeconds) * time.Second
	}
	return c.client.Delete().
		Namespace(c.ns).
		Resource("secretproviderclasses").
		VersionedParams(&listOptions, scheme.ParameterCodec).
		Timeout(timeout).
		Body(options).
		Do().
		Error()
}

// Patch applies the patch and returns the patched secretProviderClass.
func (c *secretProviderClasses) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *v1alpha1.SecretProviderClass, err error) {
	result = &v1alpha1.SecretProviderClass{}
	err = c.client.Patch(pt).
		Namespace(c.ns).
		Resource("secretproviderclasses").
		SubResource(subresources...).
		Name(name).
		Body(data).
		Do().
		Into(result)
	return
}
