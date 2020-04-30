// Package steam implements an HTTP client for the Steam Web API.
//
// For more information, refer to the main "geyser" package documentation.
package steam

import (
	"github.com/13k/geyser/v2"
)

const (
	HostURL = "https://api.steampowered.com"
)

type Client struct {
	geyser.Client
}

var _ geyser.Client = (*Client)(nil)

// New creates a new Steam API client.
func New(options ...geyser.ClientOption) (*Client, error) {
	options = append(options, geyser.WithHostURL(HostURL))

	baseClient, err := geyser.New(options...)

	if err != nil {
		return nil, err
	}

	client := &Client{Client: baseClient}

	return client, nil
}
