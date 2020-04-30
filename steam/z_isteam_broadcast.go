// Code generated by geyser. DO NOT EDIT.
// API interface: ISteamBroadcast.

package steam

import (
	"github.com/13k/geyser/v2"
	"github.com/13k/geyser/v2/schema"
	"net/http"
)

// SchemaISteamBroadcast stores the Interfaces for interface ISteamBroadcast.
var SchemaISteamBroadcast = schema.MustNewInterfaces(
	&schema.Interface{
		Methods: schema.MustNewMethods(
			&schema.Method{
				HTTPMethod: http.MethodGet,
				Name:       "ViewerHeartbeat",
				Params: schema.NewMethodParams(
					&schema.MethodParam{
						Description: "Steam ID of the broadcaster",
						Name:        "steamid",
						Optional:    false,
						Type:        "uint64",
					},
					&schema.MethodParam{
						Description: "Broadcast Session ID",
						Name:        "sessionid",
						Optional:    false,
						Type:        "uint64",
					},
					&schema.MethodParam{
						Description: "Viewer token",
						Name:        "token",
						Optional:    false,
						Type:        "uint64",
					},
					&schema.MethodParam{
						Description: "video stream representation watching",
						Name:        "stream",
						Optional:    true,
						Type:        "int32",
					},
				),
				Undocumented: false,
				Version:      1,
			},
		),
		Name:         "ISteamBroadcast",
		Undocumented: false,
	},
)

// ISteamBroadcast represents interface ISteamBroadcast.
type ISteamBroadcast struct {
	Client    *Client
	Interface *schema.Interface
}

// NewISteamBroadcast creates a new ISteamBroadcast interface.
func NewISteamBroadcast(c *Client) (*ISteamBroadcast, error) {
	si, err := SchemaISteamBroadcast.Get(schema.InterfaceKey{Name: "ISteamBroadcast"})

	if err != nil {
		return nil, err
	}

	s := &ISteamBroadcast{
		Client:    c,
		Interface: si,
	}

	return s, nil
}

// ISteamBroadcast creates a new ISteamBroadcast interface.
func (c *Client) ISteamBroadcast() (*ISteamBroadcast, error) {
	return NewISteamBroadcast(c)
}

/*
ViewerHeartbeat creates a Request for interface method ViewerHeartbeat.

Parameters

  * steamid [uint64] (required): Steam ID of the broadcaster
  * sessionid [uint64] (required): Broadcast Session ID
  * token [uint64] (required): Viewer token
  * stream [int32]: video stream representation watching
*/
func (i *ISteamBroadcast) ViewerHeartbeat() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "ViewerHeartbeat",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := geyser.NewRequest(i.Interface, sm)

	return req, nil
}
