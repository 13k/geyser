// Code generated by geyser. DO NOT EDIT.
// API interface: IDOTA2MatchStats.

package dota2

import (
	"github.com/13k/geyser"
	"github.com/13k/geyser/schema"
	"net/http"
)

// SchemaIDOTA2MatchStats stores the Interfaces for interface IDOTA2MatchStats.
var SchemaIDOTA2MatchStats = schema.MustNewInterfaces(
	&schema.Interface{
		Methods: schema.MustNewMethods(
			&schema.Method{
				HTTPMethod:   http.MethodGet,
				Name:         "GetRealtimeStats",
				Params:       schema.NewMethodParams(),
				Undocumented: true,
				Version:      1,
			},
		),
		Name:         "IDOTA2MatchStats",
		Undocumented: true,
	},
)

// IDOTA2MatchStats represents interface IDOTA2MatchStats.
//
// This is an undocumented interface.
type IDOTA2MatchStats struct {
	Client    *Client
	Interface *schema.Interface
}

// NewIDOTA2MatchStats creates a new IDOTA2MatchStats interface.
func NewIDOTA2MatchStats(c *Client) (*IDOTA2MatchStats, error) {
	si, err := SchemaIDOTA2MatchStats.Get(schema.InterfaceKey{Name: "IDOTA2MatchStats"})

	if err != nil {
		return nil, err
	}

	s := &IDOTA2MatchStats{
		Client:    c,
		Interface: si,
	}

	return s, nil
}

// IDOTA2MatchStats creates a new IDOTA2MatchStats interface.
func (c *Client) IDOTA2MatchStats() (*IDOTA2MatchStats, error) {
	return NewIDOTA2MatchStats(c)
}

/*
GetRealtimeStats creates a Request for interface method GetRealtimeStats.

This is an undocumented method.
*/
func (i *IDOTA2MatchStats) GetRealtimeStats() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "GetRealtimeStats",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := &geyser.Request{
		Client:    i.Client,
		Interface: i.Interface,
		Method:    sm,
	}

	return req, nil
}