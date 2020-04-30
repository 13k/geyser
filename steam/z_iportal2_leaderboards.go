// Code generated by geyser. DO NOT EDIT.
// API interface: IPortal2Leaderboards.

package steam

import (
	"github.com/13k/geyser/v2"
	"github.com/13k/geyser/v2/schema"
	"net/http"
)

// SchemaIPortal2Leaderboards stores the Interfaces for interface IPortal2Leaderboards.
var SchemaIPortal2Leaderboards = schema.MustNewInterfaces(
	&schema.Interface{
		Methods: schema.MustNewMethods(
			&schema.Method{
				HTTPMethod: http.MethodGet,
				Name:       "GetBucketizedData",
				Params: schema.NewMethodParams(
					&schema.MethodParam{
						Description: "The leaderboard name to fetch data for.",
						Name:        "leaderboardName",
						Optional:    false,
						Type:        "string",
					},
				),
				Undocumented: false,
				Version:      1,
			},
		),
		Name:         "IPortal2Leaderboards_620",
		Undocumented: false,
	},
)

// IPortal2Leaderboards represents interface IPortal2Leaderboards.
//
// Supported AppIDs: 620.
type IPortal2Leaderboards struct {
	Client    *Client
	Interface *schema.Interface
}

// NewIPortal2Leaderboards creates a new IPortal2Leaderboards interface.
//
// Supported AppIDs: 620.
func NewIPortal2Leaderboards(c *Client, appID uint32) (*IPortal2Leaderboards, error) {
	si, err := SchemaIPortal2Leaderboards.Get(schema.InterfaceKey{
		AppID: appID,
		Name:  "IPortal2Leaderboards",
	})

	if err != nil {
		return nil, err
	}

	s := &IPortal2Leaderboards{
		Client:    c,
		Interface: si,
	}

	return s, nil
}

// IPortal2Leaderboards creates a new IPortal2Leaderboards interface.
//
// Supported AppIDs: 620.
func (c *Client) IPortal2Leaderboards(appID uint32) (*IPortal2Leaderboards, error) {
	return NewIPortal2Leaderboards(c, appID)
}

/*
GetBucketizedData creates a Request for interface method GetBucketizedData.

Parameters

  * leaderboardName [string] (required): The leaderboard name to fetch data for.
*/
func (i *IPortal2Leaderboards) GetBucketizedData() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "GetBucketizedData",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := geyser.NewRequest(i.Interface, sm)

	return req, nil
}
