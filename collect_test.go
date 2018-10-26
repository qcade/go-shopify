package goshopify

import (
	"testing"

	httpmock "gopkg.in/jarcoal/httpmock.v1"
)

func TestCollectGet(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", "https://fooshop.myshopify.com/admin/collects/123456.json",
		httpmock.NewBytesResponder(200, loadFixture("collect.json")))

	collect, err := client.Collect.Get(123456, nil)
	if err != nil {
		t.Error(err)
	}

	// Test a few fields
	cases := []struct {
		field    string
		expected interface{}
		actual   interface{}
	}{
		{"ProductID", int64(921728736), collect.ProductID},
		{"CollectionID", int64(841564295), collect.CollectionID},
	}

	for _, c := range cases {
		if c.expected != c.actual {
			t.Errorf("Collect.%v returned %v, expected %v", c.field, c.actual, c.expected)
		}
	}
}

func TestCollectList(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", "https://fooshop.myshopify.com/admin/collects.json",
		httpmock.NewBytesResponder(200, loadFixture("collects.json")))

	collects, err := client.Collect.List(nil)
	if err != nil {
		t.Error(err)
	}
	if len(collects) != 1 {
		t.Fail()
	}

	collect := collects[0]

	// Test a few fields
	cases := []struct {
		field    string
		expected interface{}
		actual   interface{}
	}{
		{"ProductID", int64(921728736), collect.ProductID},
		{"CollectionID", int64(841564295), collect.CollectionID},
	}

	for _, c := range cases {
		if c.expected != c.actual {
			t.Errorf("Collect.%v returned %v, expected %v", c.field, c.actual, c.expected)
		}
	}
}
