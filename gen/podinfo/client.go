// Code generated by goa v3.0.3, DO NOT EDIT.
//
// podinfo client
//
// Command:
// $ goa gen github.com/centric-lt/k8s-101/design

package podinfo

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "podinfo" service client.
type Client struct {
	GetEndpoint goa.Endpoint
}

// NewClient initializes a "podinfo" service client given the endpoints.
func NewClient(get goa.Endpoint) *Client {
	return &Client{
		GetEndpoint: get,
	}
}

// Get calls the "get" endpoint of the "podinfo" service.
func (c *Client) Get(ctx context.Context) (res *Podinforesult, err error) {
	var ires interface{}
	ires, err = c.GetEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.(*Podinforesult), nil
}
