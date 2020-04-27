// Code generated by geyser. DO NOT EDIT.
// API interface: ISteamWebUserPresenceOAuth.

package steam

import (
	"github.com/13k/geyser"
	"github.com/13k/geyser/schema"
	"net/http"
)

// SchemaISteamWebUserPresenceOAuth stores the Interfaces for interface ISteamWebUserPresenceOAuth.
var SchemaISteamWebUserPresenceOAuth = schema.MustNewInterfaces(
	&schema.Interface{
		Methods: schema.MustNewMethods(
			&schema.Method{
				HTTPMethod: http.MethodPost,
				Name:       "PollStatus",
				Params: schema.NewMethodParams(
					&schema.MethodParam{
						Description: "Steam ID of the user",
						Name:        "steamid",
						Optional:    false,
						Type:        "string",
					},
					&schema.MethodParam{
						Description: "UMQ Session ID",
						Name:        "umqid",
						Optional:    false,
						Type:        "uint64",
					},
					&schema.MethodParam{
						Description: "Message that was last known to the user",
						Name:        "message",
						Optional:    false,
						Type:        "uint32",
					},
					&schema.MethodParam{
						Description: "Caller-specific poll id",
						Name:        "pollid",
						Optional:    true,
						Type:        "uint32",
					},
					&schema.MethodParam{
						Description: "Long-poll timeout in seconds",
						Name:        "sectimeout",
						Optional:    true,
						Type:        "uint32",
					},
					&schema.MethodParam{
						Description: "How many seconds is client considering itself idle, e.g. screen is off",
						Name:        "secidletime",
						Optional:    true,
						Type:        "uint32",
					},
					&schema.MethodParam{
						Description: "Boolean, 0 (default): return steamid_from in output, 1: return accountid_from",
						Name:        "use_accountids",
						Optional:    true,
						Type:        "uint32",
					},
				),
				Undocumented: false,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod:   http.MethodGet,
				Name:         "Logon",
				Params:       schema.NewMethodParams(),
				Undocumented: true,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod:   http.MethodGet,
				Name:         "Logoff",
				Params:       schema.NewMethodParams(),
				Undocumented: true,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod:   http.MethodGet,
				Name:         "Message",
				Params:       schema.NewMethodParams(),
				Undocumented: true,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod:   http.MethodGet,
				Name:         "DeviceInfo",
				Params:       schema.NewMethodParams(),
				Undocumented: true,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod:   http.MethodGet,
				Name:         "Poll",
				Params:       schema.NewMethodParams(),
				Undocumented: true,
				Version:      1,
			},
		),
		Name:         "ISteamWebUserPresenceOAuth",
		Undocumented: false,
	},
)

// ISteamWebUserPresenceOAuth represents interface ISteamWebUserPresenceOAuth.
type ISteamWebUserPresenceOAuth struct {
	Client    *Client
	Interface *schema.Interface
}

// NewISteamWebUserPresenceOAuth creates a new ISteamWebUserPresenceOAuth interface.
func NewISteamWebUserPresenceOAuth(c *Client) (*ISteamWebUserPresenceOAuth, error) {
	si, err := SchemaISteamWebUserPresenceOAuth.Get(schema.InterfaceKey{Name: "ISteamWebUserPresenceOAuth"})

	if err != nil {
		return nil, err
	}

	s := &ISteamWebUserPresenceOAuth{
		Client:    c,
		Interface: si,
	}

	return s, nil
}

// ISteamWebUserPresenceOAuth creates a new ISteamWebUserPresenceOAuth interface.
func (c *Client) ISteamWebUserPresenceOAuth() (*ISteamWebUserPresenceOAuth, error) {
	return NewISteamWebUserPresenceOAuth(c)
}

/*
DeviceInfo creates a Request for interface method DeviceInfo.

This is an undocumented method.
*/
func (i *ISteamWebUserPresenceOAuth) DeviceInfo() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "DeviceInfo",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := geyser.NewRequest(i.Interface, sm)

	return req, nil
}

/*
Logoff creates a Request for interface method Logoff.

This is an undocumented method.
*/
func (i *ISteamWebUserPresenceOAuth) Logoff() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "Logoff",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := geyser.NewRequest(i.Interface, sm)

	return req, nil
}

/*
Logon creates a Request for interface method Logon.

This is an undocumented method.
*/
func (i *ISteamWebUserPresenceOAuth) Logon() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "Logon",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := geyser.NewRequest(i.Interface, sm)

	return req, nil
}

/*
Message creates a Request for interface method Message.

This is an undocumented method.
*/
func (i *ISteamWebUserPresenceOAuth) Message() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "Message",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := geyser.NewRequest(i.Interface, sm)

	return req, nil
}

/*
Poll creates a Request for interface method Poll.

This is an undocumented method.
*/
func (i *ISteamWebUserPresenceOAuth) Poll() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "Poll",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := geyser.NewRequest(i.Interface, sm)

	return req, nil
}

/*
PollStatus creates a Request for interface method PollStatus.

Parameters

  * steamid [string] (required): Steam ID of the user
  * umqid [uint64] (required): UMQ Session ID
  * message [uint32] (required): Message that was last known to the user
  * pollid [uint32]: Caller-specific poll id
  * sectimeout [uint32]: Long-poll timeout in seconds
  * secidletime [uint32]: How many seconds is client considering itself idle, e.g. screen is off
  * use_accountids [uint32]: Boolean, 0 (default): return steamid_from in output, 1: return accountid_from
*/
func (i *ISteamWebUserPresenceOAuth) PollStatus() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "PollStatus",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := geyser.NewRequest(i.Interface, sm)

	return req, nil
}
