// Code generated by geyser. DO NOT EDIT.
// API interface: IDOTA2Fantasy.

package steam

import (
	"github.com/13k/geyser"
	"github.com/13k/geyser/schema"
	"net/http"
)

// SchemaIDOTA2Fantasy stores the Interfaces for interface IDOTA2Fantasy.
var SchemaIDOTA2Fantasy = schema.MustNewInterfaces(
	&schema.Interface{
		Methods: schema.MustNewMethods(
			&schema.Method{
				HTTPMethod: http.MethodGet,
				Name:       "GetFantasyPlayerStats",
				Params: schema.NewMethodParams(
					&schema.MethodParam{
						Description: "The fantasy league ID",
						Name:        "FantasyLeagueID",
						Optional:    false,
						Type:        "uint32",
					},
					&schema.MethodParam{
						Description: "An optional filter for minimum timestamp",
						Name:        "StartTime",
						Optional:    true,
						Type:        "uint32",
					},
					&schema.MethodParam{
						Description: "An optional filter for maximum timestamp",
						Name:        "EndTime",
						Optional:    true,
						Type:        "uint32",
					},
					&schema.MethodParam{
						Description: "An optional filter for a specific match",
						Name:        "MatchID",
						Optional:    true,
						Type:        "uint64",
					},
					&schema.MethodParam{
						Description: "An optional filter for a specific series",
						Name:        "SeriesID",
						Optional:    true,
						Type:        "uint32",
					},
					&schema.MethodParam{
						Description: "An optional filter for a specific player",
						Name:        "PlayerAccountID",
						Optional:    true,
						Type:        "uint32",
					},
				),
				Undocumented: false,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod: http.MethodGet,
				Name:       "GetPlayerOfficialInfo",
				Params: schema.NewMethodParams(
					&schema.MethodParam{
						Description: "The account ID to look up",
						Name:        "accountid",
						Optional:    false,
						Type:        "uint32",
					},
				),
				Undocumented: false,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod:   http.MethodGet,
				Name:         "GetProPlayerList",
				Params:       schema.NewMethodParams(),
				Undocumented: false,
				Version:      1,
			},
		),
		Name:         "IDOTA2Fantasy_205790",
		Undocumented: false,
	},
)

// IDOTA2Fantasy represents interface IDOTA2Fantasy.
//
// Supported AppIDs: 205790.
type IDOTA2Fantasy struct {
	Client    *Client
	Interface *schema.Interface
}

// NewIDOTA2Fantasy creates a new IDOTA2Fantasy interface.
//
// Supported AppIDs: 205790.
func NewIDOTA2Fantasy(c *Client, appID uint32) (*IDOTA2Fantasy, error) {
	si, err := SchemaIDOTA2Fantasy.Get(schema.InterfaceKey{
		AppID: appID,
		Name:  "IDOTA2Fantasy",
	})

	if err != nil {
		return nil, err
	}

	s := &IDOTA2Fantasy{
		Client:    c,
		Interface: si,
	}

	return s, nil
}

// IDOTA2Fantasy creates a new IDOTA2Fantasy interface.
//
// Supported AppIDs: 205790.
func (c *Client) IDOTA2Fantasy(appID uint32) (*IDOTA2Fantasy, error) {
	return NewIDOTA2Fantasy(c, appID)
}

/*
GetFantasyPlayerStats creates a Request for interface method GetFantasyPlayerStats.

Parameters

  * FantasyLeagueID [uint32] (required): The fantasy league ID
  * StartTime [uint32]: An optional filter for minimum timestamp
  * EndTime [uint32]: An optional filter for maximum timestamp
  * MatchID [uint64]: An optional filter for a specific match
  * SeriesID [uint32]: An optional filter for a specific series
  * PlayerAccountID [uint32]: An optional filter for a specific player
*/
func (i *IDOTA2Fantasy) GetFantasyPlayerStats() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "GetFantasyPlayerStats",
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

/*
GetPlayerOfficialInfo creates a Request for interface method GetPlayerOfficialInfo.

Parameters

  * accountid [uint32] (required): The account ID to look up
*/
func (i *IDOTA2Fantasy) GetPlayerOfficialInfo() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "GetPlayerOfficialInfo",
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

// GetProPlayerList creates a Request for interface method GetProPlayerList.
func (i *IDOTA2Fantasy) GetProPlayerList() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "GetProPlayerList",
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
