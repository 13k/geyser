// Code generated by geyser. DO NOT EDIT.
// API interface: ISteamDirectory.

package steam

import (
	"github.com/13k/geyser"
	"github.com/13k/geyser/schema"
	"net/http"
)

// SchemaISteamDirectory stores the Interfaces for interface ISteamDirectory.
var SchemaISteamDirectory = schema.MustNewInterfaces(
	&schema.Interface{
		Methods: schema.MustNewMethods(
			&schema.Method{
				HTTPMethod: http.MethodGet,
				Name:       "GetCMList",
				Params: schema.NewMethodParams(
					&schema.MethodParam{
						Description: "Client's Steam cell ID",
						Name:        "cellid",
						Optional:    false,
						Type:        "uint32",
					},
					&schema.MethodParam{
						Description: "Max number of servers to return",
						Name:        "maxcount",
						Optional:    true,
						Type:        "uint32",
					},
				),
				Undocumented: false,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod: http.MethodGet,
				Name:       "GetCSList",
				Params: schema.NewMethodParams(
					&schema.MethodParam{
						Description: "Client's Steam cell ID",
						Name:        "cellid",
						Optional:    false,
						Type:        "uint32",
					},
					&schema.MethodParam{
						Description: "Max number of servers to return",
						Name:        "maxcount",
						Optional:    true,
						Type:        "uint32",
					},
				),
				Undocumented: false,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod:   http.MethodGet,
				Name:         "GetSteamPipeDomains",
				Params:       schema.NewMethodParams(),
				Undocumented: false,
				Version:      1,
			},
		),
		Name:         "ISteamDirectory",
		Undocumented: false,
	},
)

// ISteamDirectory represents interface ISteamDirectory.
type ISteamDirectory struct {
	Client    *Client
	Interface *schema.Interface
}

// NewISteamDirectory creates a new ISteamDirectory interface.
func NewISteamDirectory(c *Client) (*ISteamDirectory, error) {
	si, err := SchemaISteamDirectory.Get(schema.InterfaceKey{Name: "ISteamDirectory"})

	if err != nil {
		return nil, err
	}

	s := &ISteamDirectory{
		Client:    c,
		Interface: si,
	}

	return s, nil
}

// ISteamDirectory creates a new ISteamDirectory interface.
func (c *Client) ISteamDirectory() (*ISteamDirectory, error) {
	return NewISteamDirectory(c)
}

/*
GetCMList creates a Request for interface method GetCMList.

Parameters

  * cellid [uint32] (required): Client's Steam cell ID
  * maxcount [uint32]: Max number of servers to return
*/
func (i *ISteamDirectory) GetCMList() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "GetCMList",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := geyser.NewRequest(i.Interface, sm)

	return req, nil
}

/*
GetCSList creates a Request for interface method GetCSList.

Parameters

  * cellid [uint32] (required): Client's Steam cell ID
  * maxcount [uint32]: Max number of servers to return
*/
func (i *ISteamDirectory) GetCSList() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "GetCSList",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := geyser.NewRequest(i.Interface, sm)

	return req, nil
}

// GetSteamPipeDomains creates a Request for interface method GetSteamPipeDomains.
func (i *ISteamDirectory) GetSteamPipeDomains() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "GetSteamPipeDomains",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := geyser.NewRequest(i.Interface, sm)

	return req, nil
}
