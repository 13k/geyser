// Code generated by geyser. DO NOT EDIT.
// API interface: IStoreService.

package steam

import (
	"github.com/13k/geyser/v2"
	"github.com/13k/geyser/v2/schema"
	"net/http"
)

// SchemaIStoreService stores the Interfaces for interface IStoreService.
var SchemaIStoreService = schema.MustNewInterfaces(
	&schema.Interface{
		Methods: schema.MustNewMethods(
			&schema.Method{
				HTTPMethod: http.MethodGet,
				Name:       "GetAppList",
				Params: schema.NewMethodParams(
					&schema.MethodParam{
						Description: "Access key",
						Name:        "key",
						Optional:    false,
						Type:        "string",
					},
					&schema.MethodParam{
						Description: "Return only items that have been modified since this date.",
						Name:        "if_modified_since",
						Optional:    true,
						Type:        "uint32",
					},
					&schema.MethodParam{
						Description: "Return only items that have a description in this language.",
						Name:        "have_description_language",
						Optional:    true,
						Type:        "string",
					},
					&schema.MethodParam{
						Description: "Include games (defaults to enabled)",
						Name:        "include_games",
						Optional:    true,
						Type:        "bool",
					},
					&schema.MethodParam{
						Description: "Include DLC",
						Name:        "include_dlc",
						Optional:    true,
						Type:        "bool",
					},
					&schema.MethodParam{
						Description: "Include software items",
						Name:        "include_software",
						Optional:    true,
						Type:        "bool",
					},
					&schema.MethodParam{
						Description: "Include videos and series",
						Name:        "include_videos",
						Optional:    true,
						Type:        "bool",
					},
					&schema.MethodParam{
						Description: "Include hardware",
						Name:        "include_hardware",
						Optional:    true,
						Type:        "bool",
					},
					&schema.MethodParam{
						Description: "For continuations, this is the last appid returned from the previous call.",
						Name:        "last_appid",
						Optional:    true,
						Type:        "uint32",
					},
					&schema.MethodParam{
						Description: "Number of results to return at a time.  Default 10k, max 50k.",
						Name:        "max_results",
						Optional:    true,
						Type:        "uint32",
					},
				),
				Undocumented: false,
				Version:      1,
			},
		),
		Name:         "IStoreService",
		Undocumented: false,
	},
)

// IStoreService represents interface IStoreService.
type IStoreService struct {
	Client    *Client
	Interface *schema.Interface
}

// NewIStoreService creates a new IStoreService interface.
func NewIStoreService(c *Client) (*IStoreService, error) {
	si, err := SchemaIStoreService.Get(schema.InterfaceKey{Name: "IStoreService"})

	if err != nil {
		return nil, err
	}

	s := &IStoreService{
		Client:    c,
		Interface: si,
	}

	return s, nil
}

// IStoreService creates a new IStoreService interface.
func (c *Client) IStoreService() (*IStoreService, error) {
	return NewIStoreService(c)
}

/*
GetAppList creates a Request for interface method GetAppList.

Parameters

  * key [string] (required): Access key
  * if_modified_since [uint32]: Return only items that have been modified since this date.
  * have_description_language [string]: Return only items that have a description in this language.
  * include_games [bool]: Include games (defaults to enabled)
  * include_dlc [bool]: Include DLC
  * include_software [bool]: Include software items
  * include_videos [bool]: Include videos and series
  * include_hardware [bool]: Include hardware
  * last_appid [uint32]: For continuations, this is the last appid returned from the previous call.
  * max_results [uint32]: Number of results to return at a time.  Default 10k, max 50k.
*/
func (i *IStoreService) GetAppList() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "GetAppList",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := geyser.NewRequest(i.Interface, sm)

	return req, nil
}
