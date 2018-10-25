package goshopify

import (
	"testing"
	"time"

	httpmock "gopkg.in/jarcoal/httpmock.v1"
)

func TestCheckoutGet(t *testing.T) {
	setup()
	defer teardown()

	httpmock.RegisterResponder("GET", "https://fooshop.myshopify.com/admin/checkouts/123456.json",
		httpmock.NewBytesResponder(200, loadFixture("checkout.json")))

	checkout, err := client.Checkout.Get(123456, nil)
	if err != nil {
		t.Errorf("Checkout.Get returned error: %v", err)
	}

	_ = checkout

	// Check that dates are parsed
	location, _ := time.LoadLocation("America/Montreal")
	d := time.Date(2012, time.October, 12, 7, 5, 27, 0, location)
	if !d.Equal(*checkout.CreatedAt) {
		t.Errorf("Checkout.CreatedAt returned %+v, expected %+v", checkout.CreatedAt, d)
	}

	// Test a few fields
	cases := []struct {
		field    string
		expected interface{}
		actual   interface{}
	}{
		{"Token", "exuw7apwoycchjuwtiqg8nytfhphr62a", checkout.Token},
		{"Currency", "USD", checkout.Currency},
		{"Email", "bob.norman@hostmail.com", checkout.Email},
		{"Name", "#862052962", checkout.Name},
		{"PaymentDue", "379.69", checkout.PaymentDue},
		{"Note", "", checkout.Note},
	}

	for _, c := range cases {
		if c.expected != c.actual {
			t.Errorf("Checkout.%v returned %v, expected %v", c.field, c.actual, c.expected)
		}
	}
}
