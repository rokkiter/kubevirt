/*
Copyright 2022 The KubeVirt Authors.

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

package fake

import (
	"context"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
	v1beta1 "kubevirt.io/containerized-data-importer-api/pkg/apis/core/v1beta1"
)

// FakeCDIs implements CDIInterface
type FakeCDIs struct {
	Fake *FakeCdiV1beta1
}

var cdisResource = schema.GroupVersionResource{Group: "cdi.kubevirt.io", Version: "v1beta1", Resource: "cdis"}

var cdisKind = schema.GroupVersionKind{Group: "cdi.kubevirt.io", Version: "v1beta1", Kind: "CDI"}

// Get takes name of the cDI, and returns the corresponding cDI object, and an error if there is any.
func (c *FakeCDIs) Get(ctx context.Context, name string, options v1.GetOptions) (result *v1beta1.CDI, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(cdisResource, name), &v1beta1.CDI{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.CDI), err
}

// List takes label and field selectors, and returns the list of CDIs that match those selectors.
func (c *FakeCDIs) List(ctx context.Context, opts v1.ListOptions) (result *v1beta1.CDIList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(cdisResource, cdisKind, opts), &v1beta1.CDIList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &v1beta1.CDIList{ListMeta: obj.(*v1beta1.CDIList).ListMeta}
	for _, item := range obj.(*v1beta1.CDIList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested cDIs.
func (c *FakeCDIs) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(cdisResource, opts))
}

// Create takes the representation of a cDI and creates it.  Returns the server's representation of the cDI, and an error, if there is any.
func (c *FakeCDIs) Create(ctx context.Context, cDI *v1beta1.CDI, opts v1.CreateOptions) (result *v1beta1.CDI, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(cdisResource, cDI), &v1beta1.CDI{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.CDI), err
}

// Update takes the representation of a cDI and updates it. Returns the server's representation of the cDI, and an error, if there is any.
func (c *FakeCDIs) Update(ctx context.Context, cDI *v1beta1.CDI, opts v1.UpdateOptions) (result *v1beta1.CDI, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(cdisResource, cDI), &v1beta1.CDI{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.CDI), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeCDIs) UpdateStatus(ctx context.Context, cDI *v1beta1.CDI, opts v1.UpdateOptions) (*v1beta1.CDI, error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateSubresourceAction(cdisResource, "status", cDI), &v1beta1.CDI{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.CDI), err
}

// Delete takes name of the cDI and deletes it. Returns an error if one occurs.
func (c *FakeCDIs) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(cdisResource, name), &v1beta1.CDI{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeCDIs) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(cdisResource, listOpts)

	_, err := c.Fake.Invokes(action, &v1beta1.CDIList{})
	return err
}

// Patch applies the patch and returns the patched cDI.
func (c *FakeCDIs) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *v1beta1.CDI, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(cdisResource, name, pt, data, subresources...), &v1beta1.CDI{})
	if obj == nil {
		return nil, err
	}
	return obj.(*v1beta1.CDI), err
}
