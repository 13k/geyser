// Package steam implements an HTTP client for the Steam Web API.
//
// For more information, refer to the main "geyser" package documentation.
package steam

import (
	"github.com/13k/geyser"
)

const (
	HostURL = "https://api.steampowered.com"
)

type Client struct {
	*geyser.Client
}

// New creates a new Steam API client.
func New(options ...geyser.ClientOption) (*Client, error) {
	options = append(options, geyser.WithBaseURL(HostURL))

	baseClient, err := geyser.New(options...)

	if err != nil {
		return nil, err
	}

	client := &Client{baseClient}

	return client, nil
}
