// Package dota2 implements an HTTP client for the (undocumented) Dota 2 Web API.
//
// For more information, refer to the main "geyser" package documentation.
package dota2

import (
	"github.com/13k/geyser"
)

const (
	HostURL = "https://www.dota2.com/webapi"
)

type Client struct {
	*geyser.Client
}

// New creates a new Dota2 API client.
func New(options ...geyser.ClientOption) (*Client, error) {
	options = append(options, geyser.WithBaseURL(HostURL))

	baseClient, err := geyser.New(options...)

	if err != nil {
		return nil, err
	}

	client := &Client{baseClient}

	return client, nil
}
