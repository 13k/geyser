// Code generated by geyser. DO NOT EDIT.
// API interface: ISteamUserOAuth.

package steam

import (
	"github.com/13k/geyser/v2"
	"github.com/13k/geyser/v2/schema"
	"net/http"
)

// SchemaISteamUserOAuth stores the Interfaces for interface ISteamUserOAuth.
var SchemaISteamUserOAuth = schema.MustNewInterfaces(
	&schema.Interface{
		Methods: schema.MustNewMethods(
			&schema.Method{
				HTTPMethod: http.MethodGet,
				Name:       "GetTokenDetails",
				Params: schema.NewMethodParams(
					&schema.MethodParam{
						Description: "OAuth2 token for which to return details",
						Name:        "access_token",
						Optional:    false,
						Type:        "string",
					},
				),
				Undocumented: false,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod:   http.MethodGet,
				Name:         "GetUserSummaries",
				Params:       schema.NewMethodParams(),
				Undocumented: true,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod:   http.MethodGet,
				Name:         "GetGroupSummaries",
				Params:       schema.NewMethodParams(),
				Undocumented: true,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod:   http.MethodGet,
				Name:         "GetGroupList",
				Params:       schema.NewMethodParams(),
				Undocumented: true,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod:   http.MethodGet,
				Name:         "GetFriendList",
				Params:       schema.NewMethodParams(),
				Undocumented: true,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod:   http.MethodGet,
				Name:         "Search",
				Params:       schema.NewMethodParams(),
				Undocumented: true,
				Version:      1,
			},
		),
		Name:         "ISteamUserOAuth",
		Undocumented: false,
	},
)

// ISteamUserOAuth represents interface ISteamUserOAuth.
type ISteamUserOAuth struct {
	Client    *Client
	Interface *schema.Interface
}

// NewISteamUserOAuth creates a new ISteamUserOAuth interface.
func NewISteamUserOAuth(c *Client) (*ISteamUserOAuth, error) {
	si, err := SchemaISteamUserOAuth.Get(schema.InterfaceKey{Name: "ISteamUserOAuth"})

	if err != nil {
		return nil, err
	}

	s := &ISteamUserOAuth{
		Client:    c,
		Interface: si,
	}

	return s, nil
}

// ISteamUserOAuth creates a new ISteamUserOAuth interface.
func (c *Client) ISteamUserOAuth() (*ISteamUserOAuth, error) {
	return NewISteamUserOAuth(c)
}

/*
GetFriendList creates a Request for interface method GetFriendList.

This is an undocumented method.
*/
func (i *ISteamUserOAuth) GetFriendList() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "GetFriendList",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := geyser.NewRequest(i.Interface, sm)

	return req, nil
}

/*
GetGroupList creates a Request for interface method GetGroupList.

This is an undocumented method.
*/
func (i *ISteamUserOAuth) GetGroupList() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "GetGroupList",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := geyser.NewRequest(i.Interface, sm)

	return req, nil
}

/*
GetGroupSummaries creates a Request for interface method GetGroupSummaries.

This is an undocumented method.
*/
func (i *ISteamUserOAuth) GetGroupSummaries() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "GetGroupSummaries",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := geyser.NewRequest(i.Interface, sm)

	return req, nil
}

/*
GetTokenDetails creates a Request for interface method GetTokenDetails.

Parameters

  * access_token [string] (required): OAuth2 token for which to return details
*/
func (i *ISteamUserOAuth) GetTokenDetails() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "GetTokenDetails",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := geyser.NewRequest(i.Interface, sm)

	return req, nil
}

/*
GetUserSummaries creates a Request for interface method GetUserSummaries.

This is an undocumented method.
*/
func (i *ISteamUserOAuth) GetUserSummaries() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "GetUserSummaries",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := geyser.NewRequest(i.Interface, sm)

	return req, nil
}

/*
Search creates a Request for interface method Search.

This is an undocumented method.
*/
func (i *ISteamUserOAuth) Search() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "Search",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := geyser.NewRequest(i.Interface, sm)

	return req, nil
}
