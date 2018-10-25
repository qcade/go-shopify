package goshopify

import (
	"fmt"
	"time"
)

const checkoutsBasePath = "admin/checkouts"

// CheckoutService is an interface for interfacing with the checkouts endpoints
// of the Shopify API.
// See: https://help.shopify.com/api/reference/checkout
type CheckoutService interface {
	Get(int, interface{}) (*Checkout, error)
}

// CheckoutServiceOp handles communication with the checkout related methods of
// the Shopify API.
type CheckoutServiceOp struct {
	client *Client
}

// Checkout represents a Shopify checkout
type Checkout struct {
	ID                  int                `json:"id,omitempty"`
	CustomerID          int                `json:"customer_id,omitempty"`
	UserID              int                `json:"user_id,omitempty"`
	Token               string             `json:"token,omitempty"`
	Currency            string             `json:"currency,omitempty"`
	Email               string             `json:"email,omitempty"`
	Name                string             `json:"name,omitempty"`
	Note                string             `json:"note,omitempty"`
	PaymentDue          string             `json:"payment_due,omitempty"`
	SubtotalPrice       string             `json:"subtotal_price,omitempty"`
	TotalTax            string             `json:"total_tax,omitempty"`
	TotalPrice          string             `json:"total_price,omitempty"`
	TotalLineItemsPrice string             `json:"total_line_items_price,omitempty"`
	DeviceID            int                `json:"device_id,omitempty"`
	LineItems           []CheckoutLineItem `json:"line_items,omitempty"`
	CreatedAt           *time.Time         `json:"created_at,omitempty"`
	UpdatedAt           *time.Time         `json:"updated_at,omitempty"`
	CompletedAt         *time.Time         `json:"completed_at,omitempty"`
}

// CheckoutLineItem ...
type CheckoutLineItem struct {
	ID                    string     `json:"id,omitempty"`
	Key                   string     `json:"key,omitempty"`
	ProductID             int64      `json:"product_id,omitempty"`
	VariantID             int64      `json:"variant_id,omitempty"`
	Quantity              int64      `json:"quantity,omitempty"`
	Grams                 int64      `json:"grams,omitempty"`
	GiftCard              bool       `json:"gift_card,omitempty"`
	RequiresShipping      bool       `json:"requires_shipping,omitempty"`
	Taxable               bool       `json:"taxable,omitempty"`
	PreTaxPrice           float64    `json:"pre_tax_price,omitempty"`
	CompareAtPrice        string     `json:"compare_at_price,omitempty"`
	DestinationLocationID int64      `json:"destination_location_id,omitempty"`
	DiscountedPrice       string     `json:"discounted_price,omitempty"`
	Discounts             string     `json:"discounts,omitempty"`
	FulfillmentService    string     `json:"fulfillment_service,omitempty"`
	LinePrice             string     `json:"line_price,omitempty"`
	OriginLocationID      int64      `json:"origin_location_id,omitempty"`
	OriginalLinePrice     string     `json:"original_line_price,omitempty"`
	OriginalPrice         string     `json:"original_price,omitempty"`
	Price                 string     `json:"price,omitempty"`
	Sku                   string     `json:"sku,omitempty"`
	TaxCode               string     `json:"tax_code,omitempty"`
	Title                 string     `json:"title,omitempty"`
	TotalDiscount         string     `json:"total_discount,omitempty"`
	VariantTitle          string     `json:"variant_title,omitempty"`
	Vendor                string     `json:"vendor,omitempty"`
	CreatedAt             *time.Time `json:"created_at,omitempty"`
	UpdatedAt             *time.Time `json:"updated_at,omitempty"`
}

// CheckoutResource represents the result from the checkouts/X.json endpoint
type CheckoutResource struct {
	Checkout *Checkout `json:"checkout"`
}

// Get customer
func (s *CheckoutServiceOp) Get(checkoutID int, options interface{}) (*Checkout, error) {
	path := fmt.Sprintf("%s/%v.json", checkoutsBasePath, checkoutID)
	resource := new(CheckoutResource)
	err := s.client.Get(path, resource, options)
	return resource.Checkout, err
}
