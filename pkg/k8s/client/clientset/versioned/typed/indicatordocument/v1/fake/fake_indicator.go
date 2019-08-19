// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	indicatordocumentv1 "github.com/pivotal/monitoring-indicator-protocol/pkg/k8s/apis/indicatordocument/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	labels "k8s.io/apimachinery/pkg/labels"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	testing "k8s.io/client-go/testing"
)

// FakeIndicators implements IndicatorInterface
type FakeIndicators struct {
	Fake *FakeIndicatorprotocolV1
	ns   string
}

var indicatorsResource = schema.GroupVersionResource{Group: "indicatorprotocol.io", Version: "v1", Resource: "indicators"}

var indicatorsKind = schema.GroupVersionKind{Group: "indicatorprotocol.io", Version: "v1", Kind: "Indicator"}

// Get takes name of the indicator, and returns the corresponding indicator object, and an error if there is any.
func (c *FakeIndicators) Get(name string, options v1.GetOptions) (result *indicatordocumentv1.Indicator, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewGetAction(indicatorsResource, c.ns, name), &indicatordocumentv1.Indicator{})

	if obj == nil {
		return nil, err
	}
	return obj.(*indicatordocumentv1.Indicator), err
}

// List takes label and field selectors, and returns the list of Indicators that match those selectors.
func (c *FakeIndicators) List(opts v1.ListOptions) (result *indicatordocumentv1.IndicatorList, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewListAction(indicatorsResource, indicatorsKind, c.ns, opts), &indicatordocumentv1.IndicatorList{})

	if obj == nil {
		return nil, err
	}

	label, _, _ := testing.ExtractFromListOptions(opts)
	if label == nil {
		label = labels.Everything()
	}
	list := &indicatordocumentv1.IndicatorList{ListMeta: obj.(*indicatordocumentv1.IndicatorList).ListMeta}
	for _, item := range obj.(*indicatordocumentv1.IndicatorList).Items {
		if label.Matches(labels.Set(item.Labels)) {
			list.Items = append(list.Items, item)
		}
	}
	return list, err
}

// Watch returns a watch.Interface that watches the requested indicators.
func (c *FakeIndicators) Watch(opts v1.ListOptions) (watch.Interface, error) {
	return c.Fake.
		InvokesWatch(testing.NewWatchAction(indicatorsResource, c.ns, opts))

}

// Create takes the representation of a indicator and creates it.  Returns the server's representation of the indicator, and an error, if there is any.
func (c *FakeIndicators) Create(indicator *indicatordocumentv1.Indicator) (result *indicatordocumentv1.Indicator, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewCreateAction(indicatorsResource, c.ns, indicator), &indicatordocumentv1.Indicator{})

	if obj == nil {
		return nil, err
	}
	return obj.(*indicatordocumentv1.Indicator), err
}

// Update takes the representation of a indicator and updates it. Returns the server's representation of the indicator, and an error, if there is any.
func (c *FakeIndicators) Update(indicator *indicatordocumentv1.Indicator) (result *indicatordocumentv1.Indicator, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateAction(indicatorsResource, c.ns, indicator), &indicatordocumentv1.Indicator{})

	if obj == nil {
		return nil, err
	}
	return obj.(*indicatordocumentv1.Indicator), err
}

// UpdateStatus was generated because the type contains a Status member.
// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
func (c *FakeIndicators) UpdateStatus(indicator *indicatordocumentv1.Indicator) (*indicatordocumentv1.Indicator, error) {
	obj, err := c.Fake.
		Invokes(testing.NewUpdateSubresourceAction(indicatorsResource, "status", c.ns, indicator), &indicatordocumentv1.Indicator{})

	if obj == nil {
		return nil, err
	}
	return obj.(*indicatordocumentv1.Indicator), err
}

// Delete takes name of the indicator and deletes it. Returns an error if one occurs.
func (c *FakeIndicators) Delete(name string, options *v1.DeleteOptions) error {
	_, err := c.Fake.
		Invokes(testing.NewDeleteAction(indicatorsResource, c.ns, name), &indicatordocumentv1.Indicator{})

	return err
}

// DeleteCollection deletes a collection of objects.
func (c *FakeIndicators) DeleteCollection(options *v1.DeleteOptions, listOptions v1.ListOptions) error {
	action := testing.NewDeleteCollectionAction(indicatorsResource, c.ns, listOptions)

	_, err := c.Fake.Invokes(action, &indicatordocumentv1.IndicatorList{})
	return err
}

// Patch applies the patch and returns the patched indicator.
func (c *FakeIndicators) Patch(name string, pt types.PatchType, data []byte, subresources ...string) (result *indicatordocumentv1.Indicator, err error) {
	obj, err := c.Fake.
		Invokes(testing.NewPatchSubresourceAction(indicatorsResource, c.ns, name, pt, data, subresources...), &indicatordocumentv1.Indicator{})

	if obj == nil {
		return nil, err
	}
	return obj.(*indicatordocumentv1.Indicator), err
}
