// Package dota2 implements an HTTP client for the (undocumented) Dota 2 Web API.
//
// For more information, refer to the main "geyser" package documentation.
package dota2

import (
	"github.com/13k/geyser/v2"
)

const (
	// HostURL is the base URL for the Dota 2 Web API client.
	HostURL = "https://www.dota2.com/webapi"
)

// Client is a thin wrapper around a base `geyser.Client` that automatically sets the base HostURL.
type Client struct {
	geyser.Client
}

var _ geyser.Client = (*Client)(nil)

// New creates a new Dota2 API client.
func New(options ...geyser.ClientOption) (*Client, error) {
	options = append(options, geyser.WithHostURL(HostURL))

	baseClient, err := geyser.New(options...)

	if err != nil {
		return nil, err
	}

	client := &Client{Client: baseClient}

	return client, nil
}
