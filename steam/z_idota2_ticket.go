// Code generated by geyser. DO NOT EDIT.
// API interface: IDOTA2Ticket.

package steam

import (
	"github.com/13k/geyser"
	"github.com/13k/geyser/schema"
	"net/http"
)

// SchemaIDOTA2Ticket stores the Interfaces for interface IDOTA2Ticket.
var SchemaIDOTA2Ticket = schema.MustNewInterfaces(
	&schema.Interface{
		Methods: schema.MustNewMethods(
			&schema.Method{
				HTTPMethod: http.MethodGet,
				Name:       "ClaimBadgeReward",
				Params: schema.NewMethodParams(
					&schema.MethodParam{
						Description: "The Badge ID",
						Name:        "BadgeID",
						Optional:    false,
						Type:        "string",
					},
					&schema.MethodParam{
						Description: "Valid Badge Type 1",
						Name:        "ValidBadgeType1",
						Optional:    false,
						Type:        "uint32",
					},
					&schema.MethodParam{
						Description: "Valid Badge Type 2",
						Name:        "ValidBadgeType2",
						Optional:    false,
						Type:        "uint32",
					},
					&schema.MethodParam{
						Description: "Valid Badge Type 3",
						Name:        "ValidBadgeType3",
						Optional:    false,
						Type:        "uint32",
					},
				),
				Undocumented: false,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod: http.MethodGet,
				Name:       "GetSteamIDForBadgeID",
				Params: schema.NewMethodParams(
					&schema.MethodParam{
						Description: "The badge ID",
						Name:        "BadgeID",
						Optional:    false,
						Type:        "string",
					},
				),
				Undocumented: false,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod: http.MethodPost,
				Name:       "SetSteamAccountPurchased",
				Params: schema.NewMethodParams(
					&schema.MethodParam{
						Description: "The 64-bit Steam ID",
						Name:        "steamid",
						Optional:    false,
						Type:        "uint64",
					},
					&schema.MethodParam{
						Description: "Badge Type",
						Name:        "BadgeType",
						Optional:    false,
						Type:        "uint32",
					},
				),
				Undocumented: false,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod: http.MethodGet,
				Name:       "SteamAccountValidForBadgeType",
				Params: schema.NewMethodParams(
					&schema.MethodParam{
						Description: "The 64-bit Steam ID",
						Name:        "steamid",
						Optional:    false,
						Type:        "uint64",
					},
					&schema.MethodParam{
						Description: "Valid Badge Type 1",
						Name:        "ValidBadgeType1",
						Optional:    false,
						Type:        "uint32",
					},
					&schema.MethodParam{
						Description: "Valid Badge Type 2",
						Name:        "ValidBadgeType2",
						Optional:    false,
						Type:        "uint32",
					},
					&schema.MethodParam{
						Description: "Valid Badge Type 3",
						Name:        "ValidBadgeType3",
						Optional:    false,
						Type:        "uint32",
					},
				),
				Undocumented: false,
				Version:      1,
			},
		),
		Name:         "IDOTA2Ticket_570",
		Undocumented: false,
	},
	&schema.Interface{
		Methods: schema.MustNewMethods(
			&schema.Method{
				HTTPMethod: http.MethodGet,
				Name:       "ClaimBadgeReward",
				Params: schema.NewMethodParams(
					&schema.MethodParam{
						Description: "The Badge ID",
						Name:        "BadgeID",
						Optional:    false,
						Type:        "string",
					},
					&schema.MethodParam{
						Description: "Valid Badge Type 1",
						Name:        "ValidBadgeType1",
						Optional:    false,
						Type:        "uint32",
					},
					&schema.MethodParam{
						Description: "Valid Badge Type 2",
						Name:        "ValidBadgeType2",
						Optional:    false,
						Type:        "uint32",
					},
					&schema.MethodParam{
						Description: "Valid Badge Type 3",
						Name:        "ValidBadgeType3",
						Optional:    false,
						Type:        "uint32",
					},
				),
				Undocumented: false,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod: http.MethodGet,
				Name:       "GetSteamIDForBadgeID",
				Params: schema.NewMethodParams(
					&schema.MethodParam{
						Description: "The badge ID",
						Name:        "BadgeID",
						Optional:    false,
						Type:        "string",
					},
				),
				Undocumented: false,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod: http.MethodPost,
				Name:       "SetSteamAccountPurchased",
				Params: schema.NewMethodParams(
					&schema.MethodParam{
						Description: "The 64-bit Steam ID",
						Name:        "steamid",
						Optional:    false,
						Type:        "uint64",
					},
					&schema.MethodParam{
						Description: "Badge Type",
						Name:        "BadgeType",
						Optional:    false,
						Type:        "uint32",
					},
				),
				Undocumented: false,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod: http.MethodGet,
				Name:       "SteamAccountValidForBadgeType",
				Params: schema.NewMethodParams(
					&schema.MethodParam{
						Description: "The 64-bit Steam ID",
						Name:        "steamid",
						Optional:    false,
						Type:        "uint64",
					},
					&schema.MethodParam{
						Description: "Valid Badge Type 1",
						Name:        "ValidBadgeType1",
						Optional:    false,
						Type:        "uint32",
					},
					&schema.MethodParam{
						Description: "Valid Badge Type 2",
						Name:        "ValidBadgeType2",
						Optional:    false,
						Type:        "uint32",
					},
					&schema.MethodParam{
						Description: "Valid Badge Type 3",
						Name:        "ValidBadgeType3",
						Optional:    false,
						Type:        "uint32",
					},
				),
				Undocumented: false,
				Version:      1,
			},
		),
		Name:         "IDOTA2Ticket_205790",
		Undocumented: false,
	},
)

// IDOTA2Ticket represents interface IDOTA2Ticket.
//
// Supported AppIDs: 570, 205790.
type IDOTA2Ticket struct {
	Client    *Client
	Interface *schema.Interface
}

// NewIDOTA2Ticket creates a new IDOTA2Ticket interface.
//
// Supported AppIDs: 570, 205790.
func NewIDOTA2Ticket(c *Client, appID uint32) (*IDOTA2Ticket, error) {
	si, err := SchemaIDOTA2Ticket.Get(schema.InterfaceKey{
		AppID: appID,
		Name:  "IDOTA2Ticket",
	})

	if err != nil {
		return nil, err
	}

	s := &IDOTA2Ticket{
		Client:    c,
		Interface: si,
	}

	return s, nil
}

// IDOTA2Ticket creates a new IDOTA2Ticket interface.
//
// Supported AppIDs: 570, 205790.
func (c *Client) IDOTA2Ticket(appID uint32) (*IDOTA2Ticket, error) {
	return NewIDOTA2Ticket(c, appID)
}

/*
ClaimBadgeReward creates a Request for interface method ClaimBadgeReward.

Parameters

  * BadgeID [string] (required): The Badge ID
  * ValidBadgeType1 [uint32] (required): Valid Badge Type 1
  * ValidBadgeType2 [uint32] (required): Valid Badge Type 2
  * ValidBadgeType3 [uint32] (required): Valid Badge Type 3
*/
func (i *IDOTA2Ticket) ClaimBadgeReward() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "ClaimBadgeReward",
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
GetSteamIDForBadgeID creates a Request for interface method GetSteamIDForBadgeID.

Parameters

  * BadgeID [string] (required): The badge ID
*/
func (i *IDOTA2Ticket) GetSteamIDForBadgeID() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "GetSteamIDForBadgeID",
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
SetSteamAccountPurchased creates a Request for interface method SetSteamAccountPurchased.

Parameters

  * steamid [uint64] (required): The 64-bit Steam ID
  * BadgeType [uint32] (required): Badge Type
*/
func (i *IDOTA2Ticket) SetSteamAccountPurchased() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "SetSteamAccountPurchased",
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
SteamAccountValidForBadgeType creates a Request for interface method SteamAccountValidForBadgeType.

Parameters

  * steamid [uint64] (required): The 64-bit Steam ID
  * ValidBadgeType1 [uint32] (required): Valid Badge Type 1
  * ValidBadgeType2 [uint32] (required): Valid Badge Type 2
  * ValidBadgeType3 [uint32] (required): Valid Badge Type 3
*/
func (i *IDOTA2Ticket) SteamAccountValidForBadgeType() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "SteamAccountValidForBadgeType",
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