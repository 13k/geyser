// Code generated by geyser. DO NOT EDIT.
// API interface: IEconDOTA2.

package steam

import (
	"github.com/13k/geyser"
	"github.com/13k/geyser/schema"
	"net/http"
)

// SchemaIEconDOTA2 stores the Interfaces for interface IEconDOTA2.
var SchemaIEconDOTA2 = schema.MustNewInterfaces(
	&schema.Interface{
		Methods: schema.MustNewMethods(
			&schema.Method{
				HTTPMethod: http.MethodGet,
				Name:       "GetEventStatsForAccount",
				Params: schema.NewMethodParams(
					&schema.MethodParam{
						Description: "The League ID of the compendium you're looking for.",
						Name:        "eventid",
						Optional:    false,
						Type:        "uint32",
					},
					&schema.MethodParam{
						Description: "The account ID to look up.",
						Name:        "accountid",
						Optional:    false,
						Type:        "uint32",
					},
					&schema.MethodParam{
						Description: "The language to provide hero names in.",
						Name:        "language",
						Optional:    true,
						Type:        "string",
					},
				),
				Undocumented: false,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod: http.MethodGet,
				Name:       "GetGameItems",
				Params: schema.NewMethodParams(
					&schema.MethodParam{
						Description: "The language to provide item names in.",
						Name:        "language",
						Optional:    true,
						Type:        "string",
					},
				),
				Undocumented: false,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod: http.MethodGet,
				Name:       "GetHeroes",
				Params: schema.NewMethodParams(
					&schema.MethodParam{
						Description: "The language to provide hero names in.",
						Name:        "language",
						Optional:    true,
						Type:        "string",
					},
					&schema.MethodParam{
						Description: "Return a list of itemized heroes only.",
						Name:        "itemizedonly",
						Optional:    true,
						Type:        "bool",
					},
				),
				Undocumented: false,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod: http.MethodGet,
				Name:       "GetItemCreators",
				Params: schema.NewMethodParams(
					&schema.MethodParam{
						Description: "The item definition to get creator information for.",
						Name:        "itemdef",
						Optional:    false,
						Type:        "uint32",
					},
				),
				Undocumented: false,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod: http.MethodGet,
				Name:       "GetRarities",
				Params: schema.NewMethodParams(
					&schema.MethodParam{
						Description: "The language to provide rarity names in.",
						Name:        "language",
						Optional:    true,
						Type:        "string",
					},
				),
				Undocumented: false,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod: http.MethodGet,
				Name:       "GetTournamentPrizePool",
				Params: schema.NewMethodParams(
					&schema.MethodParam{
						Description: "The ID of the league to get the prize pool of",
						Name:        "leagueid",
						Optional:    true,
						Type:        "uint32",
					},
				),
				Undocumented: false,
				Version:      1,
			},
		),
		Name:         "IEconDOTA2_570",
		Undocumented: false,
	},
	&schema.Interface{
		Methods: schema.MustNewMethods(
			&schema.Method{
				HTTPMethod: http.MethodGet,
				Name:       "GetEventStatsForAccount",
				Params: schema.NewMethodParams(
					&schema.MethodParam{
						Description: "The League ID of the compendium you're looking for.",
						Name:        "eventid",
						Optional:    false,
						Type:        "uint32",
					},
					&schema.MethodParam{
						Description: "The account ID to look up.",
						Name:        "accountid",
						Optional:    false,
						Type:        "uint32",
					},
					&schema.MethodParam{
						Description: "The language to provide hero names in.",
						Name:        "language",
						Optional:    true,
						Type:        "string",
					},
				),
				Undocumented: false,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod: http.MethodGet,
				Name:       "GetGameItems",
				Params: schema.NewMethodParams(
					&schema.MethodParam{
						Description: "The language to provide item names in.",
						Name:        "language",
						Optional:    true,
						Type:        "string",
					},
				),
				Undocumented: false,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod: http.MethodGet,
				Name:       "GetHeroes",
				Params: schema.NewMethodParams(
					&schema.MethodParam{
						Description: "The language to provide hero names in.",
						Name:        "language",
						Optional:    true,
						Type:        "string",
					},
					&schema.MethodParam{
						Description: "Return a list of itemized heroes only.",
						Name:        "itemizedonly",
						Optional:    true,
						Type:        "bool",
					},
				),
				Undocumented: false,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod: http.MethodGet,
				Name:       "GetItemIconPath",
				Params: schema.NewMethodParams(
					&schema.MethodParam{
						Description: "The item icon name to get the CDN path of",
						Name:        "iconname",
						Optional:    false,
						Type:        "string",
					},
					&schema.MethodParam{
						Description: "The type of image you want. 0 = normal, 1 = large, 2 = ingame",
						Name:        "icontype",
						Optional:    true,
						Type:        "uint32",
					},
				),
				Undocumented: false,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod: http.MethodGet,
				Name:       "GetRarities",
				Params: schema.NewMethodParams(
					&schema.MethodParam{
						Description: "The language to provide rarity names in.",
						Name:        "language",
						Optional:    true,
						Type:        "string",
					},
				),
				Undocumented: false,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod: http.MethodGet,
				Name:       "GetTournamentPrizePool",
				Params: schema.NewMethodParams(
					&schema.MethodParam{
						Description: "The ID of the league to get the prize pool of",
						Name:        "leagueid",
						Optional:    true,
						Type:        "uint32",
					},
				),
				Undocumented: false,
				Version:      1,
			},
		),
		Name:         "IEconDOTA2_205790",
		Undocumented: false,
	},
)

// IEconDOTA2 represents interface IEconDOTA2.
//
// Supported AppIDs: 570, 205790.
type IEconDOTA2 struct {
	Client    *Client
	Interface *schema.Interface
}

// NewIEconDOTA2 creates a new IEconDOTA2 interface.
//
// Supported AppIDs: 570, 205790.
func NewIEconDOTA2(c *Client, appID uint32) (*IEconDOTA2, error) {
	si, err := SchemaIEconDOTA2.Get(schema.InterfaceKey{
		AppID: appID,
		Name:  "IEconDOTA2",
	})

	if err != nil {
		return nil, err
	}

	s := &IEconDOTA2{
		Client:    c,
		Interface: si,
	}

	return s, nil
}

// IEconDOTA2 creates a new IEconDOTA2 interface.
//
// Supported AppIDs: 570, 205790.
func (c *Client) IEconDOTA2(appID uint32) (*IEconDOTA2, error) {
	return NewIEconDOTA2(c, appID)
}

/*
GetEventStatsForAccount creates a Request for interface method GetEventStatsForAccount.

Parameters

  * eventid [uint32] (required): The League ID of the compendium you're looking for.
  * accountid [uint32] (required): The account ID to look up.
  * language [string]: The language to provide hero names in.
*/
func (i *IEconDOTA2) GetEventStatsForAccount() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "GetEventStatsForAccount",
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
GetGameItems creates a Request for interface method GetGameItems.

Parameters

  * language [string]: The language to provide item names in.
*/
func (i *IEconDOTA2) GetGameItems() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "GetGameItems",
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
GetHeroes creates a Request for interface method GetHeroes.

Parameters

  * language [string]: The language to provide hero names in.
  * itemizedonly [bool]: Return a list of itemized heroes only.
*/
func (i *IEconDOTA2) GetHeroes() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "GetHeroes",
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
GetItemCreators creates a Request for interface method GetItemCreators.

Parameters

  * itemdef [uint32] (required): The item definition to get creator information for.
*/
func (i *IEconDOTA2) GetItemCreators() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "GetItemCreators",
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
GetItemIconPath creates a Request for interface method GetItemIconPath.

Parameters

  * iconname [string] (required): The item icon name to get the CDN path of
  * icontype [uint32]: The type of image you want. 0 = normal, 1 = large, 2 = ingame
*/
func (i *IEconDOTA2) GetItemIconPath() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "GetItemIconPath",
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
GetRarities creates a Request for interface method GetRarities.

Parameters

  * language [string]: The language to provide rarity names in.
*/
func (i *IEconDOTA2) GetRarities() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "GetRarities",
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
GetTournamentPrizePool creates a Request for interface method GetTournamentPrizePool.

Parameters

  * leagueid [uint32]: The ID of the league to get the prize pool of
*/
func (i *IEconDOTA2) GetTournamentPrizePool() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "GetTournamentPrizePool",
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