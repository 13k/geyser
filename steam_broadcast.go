// Code generated by github.com/13k/geyser/apigen. DO NOT EDIT.
// API interface: ISteamBroadcast

package geyser

import "net/http"

// SchemaSteamBroadcast stores the SchemaInterfaces for interface ISteamBroadcast
var SchemaSteamBroadcast = MustNewSchemaInterfaces(
	&SchemaInterface{
		Methods: NewSchemaMethods(
			&SchemaMethod{
				HTTPMethod: http.MethodGet,
				Name:       "ViewerHeartbeat",
				Params: NewSchemaMethodParams(
					&SchemaMethodParam{
						Description: "Steam ID of the broadcaster",
						Name:        "steamid",
						Optional:    false,
						Type:        "uint64",
					},
					&SchemaMethodParam{
						Description: "Broadcast Session ID",
						Name:        "sessionid",
						Optional:    false,
						Type:        "uint64",
					},
					&SchemaMethodParam{
						Description: "Viewer token",
						Name:        "token",
						Optional:    false,
						Type:        "uint64",
					},
					&SchemaMethodParam{
						Description: "video stream representation watching",
						Name:        "stream",
						Optional:    true,
						Type:        "int32",
					},
				),
				Version: 1,
			},
		),
		Name: "ISteamBroadcast",
	},
)

// SteamBroadcast represents interface ISteamBroadcast
type SteamBroadcast struct {
	Client    *Client
	Interface *SchemaInterface
}

// NewSteamBroadcast creates a new SteamBroadcast interface
func NewSteamBroadcast(c *Client) (*SteamBroadcast, error) {
	si, err := SchemaSteamBroadcast.Get("ISteamBroadcast", 0)

	if err != nil {
		return nil, err
	}

	s := &SteamBroadcast{
		Client:    c,
		Interface: si,
	}

	return s, nil
}

// SteamBroadcast creates a new SteamBroadcast interface
func (c *Client) SteamBroadcast() (*SteamBroadcast, error) {
	return NewSteamBroadcast(c)
}

// ViewerHeartbeat creates a Request for interface method ViewerHeartbeat
func (i *SteamBroadcast) ViewerHeartbeat() (*Request, error) {
	sm, err := i.Interface.Methods.Get("ViewerHeartbeat", 1)

	if err != nil {
		return nil, err
	}

	req := &Request{
		Client:    i.Client,
		Interface: i.Interface,
		Method:    sm,
		Result:    &SteamBroadcastViewerHeartbeat{},
	}

	return req, nil
}
