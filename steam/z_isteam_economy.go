// Code generated by geyser. DO NOT EDIT.
// API interface: ISteamEconomy.

package steam

import (
	"github.com/13k/geyser"
	"github.com/13k/geyser/schema"
	"net/http"
)

// SchemaISteamEconomy stores the Interfaces for interface ISteamEconomy.
var SchemaISteamEconomy = schema.MustNewInterfaces(
	&schema.Interface{
		Methods: schema.MustNewMethods(
			&schema.Method{
				HTTPMethod: http.MethodGet,
				Name:       "GetAssetClassInfo",
				Params: schema.NewMethodParams(
					&schema.MethodParam{
						Description: "Must be a steam economy app.",
						Name:        "appid",
						Optional:    false,
						Type:        "uint32",
					},
					&schema.MethodParam{
						Description: "The user's local language",
						Name:        "language",
						Optional:    true,
						Type:        "string",
					},
					&schema.MethodParam{
						Description: "Number of classes requested. Must be at least one.",
						Name:        "class_count",
						Optional:    false,
						Type:        "uint32",
					},
					&schema.MethodParam{
						Description: "Class ID of the nth class.",
						Name:        "classid0",
						Optional:    false,
						Type:        "uint64",
					},
					&schema.MethodParam{
						Description: "Instance ID of the nth class.",
						Name:        "instanceid0",
						Optional:    true,
						Type:        "uint64",
					},
				),
				Undocumented: false,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod: http.MethodGet,
				Name:       "GetAssetPrices",
				Params: schema.NewMethodParams(
					&schema.MethodParam{
						Description: "Must be a steam economy app.",
						Name:        "appid",
						Optional:    false,
						Type:        "uint32",
					},
					&schema.MethodParam{
						Description: "The currency to filter for",
						Name:        "currency",
						Optional:    true,
						Type:        "string",
					},
					&schema.MethodParam{
						Description: "The user's local language",
						Name:        "language",
						Optional:    true,
						Type:        "string",
					},
				),
				Undocumented: false,
				Version:      1,
			},
		),
		Name:         "ISteamEconomy",
		Undocumented: false,
	},
)

// ISteamEconomy represents interface ISteamEconomy.
type ISteamEconomy struct {
	Client    *Client
	Interface *schema.Interface
}

// NewISteamEconomy creates a new ISteamEconomy interface.
func NewISteamEconomy(c *Client) (*ISteamEconomy, error) {
	si, err := SchemaISteamEconomy.Get(schema.InterfaceKey{Name: "ISteamEconomy"})

	if err != nil {
		return nil, err
	}

	s := &ISteamEconomy{
		Client:    c,
		Interface: si,
	}

	return s, nil
}

// ISteamEconomy creates a new ISteamEconomy interface.
func (c *Client) ISteamEconomy() (*ISteamEconomy, error) {
	return NewISteamEconomy(c)
}

/*
GetAssetClassInfo creates a Request for interface method GetAssetClassInfo.

Parameters

  * appid [uint32] (required): Must be a steam economy app.
  * language [string]: The user's local language
  * class_count [uint32] (required): Number of classes requested. Must be at least one.
  * classid0 [uint64] (required): Class ID of the nth class.
  * instanceid0 [uint64]: Instance ID of the nth class.
*/
func (i *ISteamEconomy) GetAssetClassInfo() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "GetAssetClassInfo",
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
GetAssetPrices creates a Request for interface method GetAssetPrices.

Parameters

  * appid [uint32] (required): Must be a steam economy app.
  * currency [string]: The currency to filter for
  * language [string]: The user's local language
*/
func (i *ISteamEconomy) GetAssetPrices() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "GetAssetPrices",
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