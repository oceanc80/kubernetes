// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	"context"

	userv1 "github.com/openshift/api/user/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeGroups implements GroupInterface
type FakeGroups struct {
	Fake *FakeUserV1
}

var groupsResource = schema.GroupVersionResource{Group: "user.openshift.io", Version: "v1", Resource: "groups"}

var groupsKind = schema.GroupVersionKind{Group: "user.openshift.io", Version: "v1", Kind: "Group"}

// Get takes name of the group, and returns the corresponding group object, and an error if there is any.
func (c *FakeGroups) Get(ctx context.Context, name string, options v1.GetOptions) (result *userv1.Group, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootGetAction(groupsResource, name), &userv1.Group{})
	if obj == nil {
		return nil, err
	}
	return obj.(*userv1.Group), err
}

// List takes label and field selectors, and returns the list of Groups that match those selectors.
func (c *FakeGroups) List(ctx context.Context, opts v1.ListOptions) (result *userv1.GroupList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootListAction(groupsResource, groupsKind, opts), &userv1.GroupList{})
	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &userv1.GroupList{ListMeta: obj.(*userv1.GroupList).ListMeta}
	for _, item := range obj.(*userv1.GroupList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested groups.
func (c *FakeGroups) Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewRootWatchAction(groupsResource, opts))
}

// Create takes the representation of a group and creates it.  Returns the server's representation of the group, and an error, if there is any.
func (c *FakeGroups) Create(ctx context.Context, group *userv1.Group, opts v1.CreateOptions) (result *userv1.Group, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootCreateAction(groupsResource, group), &userv1.Group{})
	if obj == nil {
		return nil, err
	}
	return obj.(*userv1.Group), err
}

// Update takes the representation of a group and updates it. Returns the server's representation of the group, and an error, if there is any.
func (c *FakeGroups) Update(ctx context.Context, group *userv1.Group, opts v1.UpdateOptions) (result *userv1.Group, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootUpdateAction(groupsResource, group), &userv1.Group{})
	if obj == nil {
		return nil, err
	}
	return obj.(*userv1.Group), err
}

// Delete takes name of the group and deletes it. Returns an error if one occurs.
func (c *FakeGroups) Delete(ctx context.Context, name string, opts v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewRootDeleteAction(groupsResource, name), &userv1.Group{})
	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeGroups) DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error {
	action := testing.NewRootDeleteCollectionAction(groupsResource, listOpts)

	_, err := c.Fake.Invokes(action, &userv1.GroupList{})
	return err
}

// Patch applies the patch and returns the patched group.
func (c *FakeGroups) Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *userv1.Group, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewRootPatchSubresourceAction(groupsResource, name, pt, data, subresources...), &userv1.Group{})
	if obj == nil {
		return nil, err
	}
	return obj.(*userv1.Group), err
}
