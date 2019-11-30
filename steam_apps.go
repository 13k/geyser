// Code generated by github.com/13k/geyser/apigen. DO NOT EDIT.
// API interface: ISteamApps

package geyser

import "net/http"

// SchemaSteamApps stores the SchemaInterfaces for interface ISteamApps
var SchemaSteamApps = MustNewSchemaInterfaces(
	&SchemaInterface{
		Methods: NewSchemaMethods(
			&SchemaMethod{
				HTTPMethod: http.MethodGet,
				Name:       "GetAppList",
				Params:     NewSchemaMethodParams(),
				Version:    1,
			},
			&SchemaMethod{
				HTTPMethod: http.MethodGet,
				Name:       "GetAppList",
				Params:     NewSchemaMethodParams(),
				Version:    2,
			},
			&SchemaMethod{
				HTTPMethod: http.MethodGet,
				Name:       "GetSDRConfig",
				Params: NewSchemaMethodParams(
					&SchemaMethodParam{
						Description: "AppID of game",
						Name:        "appid",
						Optional:    false,
						Type:        "uint32",
					},
					&SchemaMethodParam{
						Description: "Partner type",
						Name:        "partner",
						Optional:    true,
						Type:        "string",
					},
				),
				Version: 1,
			},
			&SchemaMethod{
				HTTPMethod: http.MethodGet,
				Name:       "GetServersAtAddress",
				Params: NewSchemaMethodParams(
					&SchemaMethodParam{
						Description: "IP or IP:queryport to list",
						Name:        "addr",
						Optional:    false,
						Type:        "string",
					},
				),
				Version: 1,
			},
			&SchemaMethod{
				HTTPMethod: http.MethodGet,
				Name:       "UpToDateCheck",
				Params: NewSchemaMethodParams(
					&SchemaMethodParam{
						Description: "AppID of game",
						Name:        "appid",
						Optional:    false,
						Type:        "uint32",
					},
					&SchemaMethodParam{
						Description: "The installed version of the game",
						Name:        "version",
						Optional:    false,
						Type:        "uint32",
					},
				),
				Version: 1,
			},
		),
		Name: "ISteamApps",
	},
)

// SteamApps represents interface ISteamApps
type SteamApps struct {
	Client    *Client
	Interface *SchemaInterface
}

// NewSteamApps creates a new SteamApps interface
func NewSteamApps(c *Client) (*SteamApps, error) {
	si, err := SchemaSteamApps.Get("ISteamApps", 0)

	if err != nil {
		return nil, err
	}

	s := &SteamApps{
		Client:    c,
		Interface: si,
	}

	return s, nil
}

// SteamApps creates a new SteamApps interface
func (c *Client) SteamApps() (*SteamApps, error) {
	return NewSteamApps(c)
}

// GetAppList creates a Request for interface method GetAppList
// Supported versions: [2 1]
func (i *SteamApps) GetAppList(version int) (*Request, error) {
	sm, err := i.Interface.Methods.Get("GetAppList", version)

	if err != nil {
		return nil, err
	}

	req := &Request{
		Client:    i.Client,
		Interface: i.Interface,
		Method:    sm,
		Result:    &SteamAppsGetAppList{},
	}

	return req, nil
}

// GetSDRConfig creates a Request for interface method GetSDRConfig
func (i *SteamApps) GetSDRConfig() (*Request, error) {
	sm, err := i.Interface.Methods.Get("GetSDRConfig", 1)

	if err != nil {
		return nil, err
	}

	req := &Request{
		Client:    i.Client,
		Interface: i.Interface,
		Method:    sm,
		Result:    &SteamAppsGetSDRConfig{},
	}

	return req, nil
}

// GetServersAtAddress creates a Request for interface method GetServersAtAddress
func (i *SteamApps) GetServersAtAddress() (*Request, error) {
	sm, err := i.Interface.Methods.Get("GetServersAtAddress", 1)

	if err != nil {
		return nil, err
	}

	req := &Request{
		Client:    i.Client,
		Interface: i.Interface,
		Method:    sm,
		Result:    &SteamAppsGetServersAtAddress{},
	}

	return req, nil
}

// UpToDateCheck creates a Request for interface method UpToDateCheck
func (i *SteamApps) UpToDateCheck() (*Request, error) {
	sm, err := i.Interface.Methods.Get("UpToDateCheck", 1)

	if err != nil {
		return nil, err
	}

	req := &Request{
		Client:    i.Client,
		Interface: i.Interface,
		Method:    sm,
		Result:    &SteamAppsUpToDateCheck{},
	}

	return req, nil
}
