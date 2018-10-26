package goshopify

import (
	"fmt"
)

// Collect ...
type Collect struct {
	ID           int64 `json:"id,omitempty"`
	ProductID    int64 `json:"product_id,omitempty"`
	CollectionID int64 `json:"collection_id,omitempty"`
}

const collectsBasePath = "admin/collects"

// CollectService is an interface for interfacing with the collects endpoints
// of the Shopify API.
// See: https://help.shopify.com/api/reference/collect
type CollectService interface {
	Get(int, interface{}) (*Collect, error)
	List(interface{}) ([]Collect, error)
}

// CollectServiceOp handles communication with the collect related methods of
// the Shopify API.
type CollectServiceOp struct {
	client *Client
}

// CollectResource represents the result from the collects/X.json endpoint
type CollectResource struct {
	Collect *Collect `json:"collect"`
}

// CollectsResource ...
type CollectsResource struct {
	Collects []Collect `json:"collects"`
}

// Get collect
func (s *CollectServiceOp) Get(collectID int, options interface{}) (*Collect, error) {
	path := fmt.Sprintf("%s/%v.json", collectsBasePath, collectID)

	resource := new(CollectResource)
	err := s.client.Get(path, resource, options)
	return resource.Collect, err
}

// List collects
func (s *CollectServiceOp) List(options interface{}) ([]Collect, error) {
	path := fmt.Sprintf("%s.json", collectsBasePath)
	resource := new(CollectsResource)
	err := s.client.Get(path, resource, options)
	return resource.Collects, err
}
