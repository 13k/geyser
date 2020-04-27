// Code generated by geyser. DO NOT EDIT.
// API interface: ICSGOPlayers.

package steam

import (
	"github.com/13k/geyser"
	"github.com/13k/geyser/schema"
	"net/http"
)

// SchemaICSGOPlayers stores the Interfaces for interface ICSGOPlayers.
var SchemaICSGOPlayers = schema.MustNewInterfaces(
	&schema.Interface{
		Methods: schema.MustNewMethods(
			&schema.Method{
				HTTPMethod: http.MethodGet,
				Name:       "GetNextMatchSharingCode",
				Params: schema.NewMethodParams(
					&schema.MethodParam{
						Description: "The SteamID of the user",
						Name:        "steamid",
						Optional:    false,
						Type:        "uint64",
					},
					&schema.MethodParam{
						Description: "Authentication obtained from the SteamID",
						Name:        "steamidkey",
						Optional:    false,
						Type:        "string",
					},
					&schema.MethodParam{
						Description: "Previously known match sharing code obtained from the SteamID",
						Name:        "knowncode",
						Optional:    false,
						Type:        "string",
					},
				),
				Undocumented: false,
				Version:      1,
			},
		),
		Name:         "ICSGOPlayers_730",
		Undocumented: false,
	},
)

// ICSGOPlayers represents interface ICSGOPlayers.
//
// Supported AppIDs: 730.
type ICSGOPlayers struct {
	Client    *Client
	Interface *schema.Interface
}

// NewICSGOPlayers creates a new ICSGOPlayers interface.
//
// Supported AppIDs: 730.
func NewICSGOPlayers(c *Client, appID uint32) (*ICSGOPlayers, error) {
	si, err := SchemaICSGOPlayers.Get(schema.InterfaceKey{
		AppID: appID,
		Name:  "ICSGOPlayers",
	})

	if err != nil {
		return nil, err
	}

	s := &ICSGOPlayers{
		Client:    c,
		Interface: si,
	}

	return s, nil
}

// ICSGOPlayers creates a new ICSGOPlayers interface.
//
// Supported AppIDs: 730.
func (c *Client) ICSGOPlayers(appID uint32) (*ICSGOPlayers, error) {
	return NewICSGOPlayers(c, appID)
}

/*
GetNextMatchSharingCode creates a Request for interface method GetNextMatchSharingCode.

Parameters

  * steamid [uint64] (required): The SteamID of the user
  * steamidkey [string] (required): Authentication obtained from the SteamID
  * knowncode [string] (required): Previously known match sharing code obtained from the SteamID
*/
func (i *ICSGOPlayers) GetNextMatchSharingCode() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "GetNextMatchSharingCode",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := geyser.NewRequest(i.Interface, sm)

	return req, nil
}
