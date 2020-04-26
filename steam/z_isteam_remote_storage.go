// Code generated by geyser. DO NOT EDIT.
// API interface: ISteamRemoteStorage.

package steam

import (
	"github.com/13k/geyser"
	"github.com/13k/geyser/schema"
	"net/http"
)

// SchemaISteamRemoteStorage stores the Interfaces for interface ISteamRemoteStorage.
var SchemaISteamRemoteStorage = schema.MustNewInterfaces(
	&schema.Interface{
		Methods: schema.MustNewMethods(
			&schema.Method{
				HTTPMethod: http.MethodPost,
				Name:       "GetCollectionDetails",
				Params: schema.NewMethodParams(
					&schema.MethodParam{
						Description: "Number of collections being requested",
						Name:        "collectioncount",
						Optional:    false,
						Type:        "uint32",
					},
					&schema.MethodParam{
						Description: "collection ids to get the details for",
						Name:        "publishedfileids[0]",
						Optional:    false,
						Type:        "uint64",
					},
				),
				Undocumented: false,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod: http.MethodPost,
				Name:       "GetPublishedFileDetails",
				Params: schema.NewMethodParams(
					&schema.MethodParam{
						Description: "Number of items being requested",
						Name:        "itemcount",
						Optional:    false,
						Type:        "uint32",
					},
					&schema.MethodParam{
						Description: "published file id to look up",
						Name:        "publishedfileids[0]",
						Optional:    false,
						Type:        "uint64",
					},
				),
				Undocumented: false,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod: http.MethodGet,
				Name:       "GetUGCFileDetails",
				Params: schema.NewMethodParams(
					&schema.MethodParam{
						Description: "If specified, only returns details if the file is owned by the SteamID specified",
						Name:        "steamid",
						Optional:    true,
						Type:        "uint64",
					},
					&schema.MethodParam{
						Description: "ID of UGC file to get info for",
						Name:        "ugcid",
						Optional:    false,
						Type:        "uint64",
					},
					&schema.MethodParam{
						Description: "appID of product",
						Name:        "appid",
						Optional:    false,
						Type:        "uint32",
					},
				),
				Undocumented: false,
				Version:      1,
			},
		),
		Name:         "ISteamRemoteStorage",
		Undocumented: false,
	},
)

// ISteamRemoteStorage represents interface ISteamRemoteStorage.
type ISteamRemoteStorage struct {
	Client    *Client
	Interface *schema.Interface
}

// NewISteamRemoteStorage creates a new ISteamRemoteStorage interface.
func NewISteamRemoteStorage(c *Client) (*ISteamRemoteStorage, error) {
	si, err := SchemaISteamRemoteStorage.Get(schema.InterfaceKey{Name: "ISteamRemoteStorage"})

	if err != nil {
		return nil, err
	}

	s := &ISteamRemoteStorage{
		Client:    c,
		Interface: si,
	}

	return s, nil
}

// ISteamRemoteStorage creates a new ISteamRemoteStorage interface.
func (c *Client) ISteamRemoteStorage() (*ISteamRemoteStorage, error) {
	return NewISteamRemoteStorage(c)
}

/*
GetCollectionDetails creates a Request for interface method GetCollectionDetails.

Parameters

  * collectioncount [uint32] (required): Number of collections being requested
  * publishedfileids[0] [uint64] (required): collection ids to get the details for
*/
func (i *ISteamRemoteStorage) GetCollectionDetails() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "GetCollectionDetails",
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
GetPublishedFileDetails creates a Request for interface method GetPublishedFileDetails.

Parameters

  * itemcount [uint32] (required): Number of items being requested
  * publishedfileids[0] [uint64] (required): published file id to look up
*/
func (i *ISteamRemoteStorage) GetPublishedFileDetails() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "GetPublishedFileDetails",
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
GetUGCFileDetails creates a Request for interface method GetUGCFileDetails.

Parameters

  * steamid [uint64]: If specified, only returns details if the file is owned by the SteamID specified
  * ugcid [uint64] (required): ID of UGC file to get info for
  * appid [uint32] (required): appID of product
*/
func (i *ISteamRemoteStorage) GetUGCFileDetails() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "GetUGCFileDetails",
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