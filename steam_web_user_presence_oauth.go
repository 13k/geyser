// Code generated by github.com/13k/geyser/apigen. DO NOT EDIT.
// API interface: ISteamWebUserPresenceOAuth.

package geyser

import "net/http"

// SchemaSteamWebUserPresenceOAuth stores the SchemaInterfaces for interface ISteamWebUserPresenceOAuth.
var SchemaSteamWebUserPresenceOAuth = MustNewSchemaInterfaces(
	&SchemaInterface{
		Methods: MustNewSchemaMethods(
			&SchemaMethod{
				HTTPMethod: http.MethodPost,
				Name:       "PollStatus",
				Params: NewSchemaMethodParams(
					&SchemaMethodParam{
						Description: "Steam ID of the user",
						Name:        "steamid",
						Optional:    false,
						Type:        "string",
					},
					&SchemaMethodParam{
						Description: "UMQ Session ID",
						Name:        "umqid",
						Optional:    false,
						Type:        "uint64",
					},
					&SchemaMethodParam{
						Description: "Message that was last known to the user",
						Name:        "message",
						Optional:    false,
						Type:        "uint32",
					},
					&SchemaMethodParam{
						Description: "Caller-specific poll id",
						Name:        "pollid",
						Optional:    true,
						Type:        "uint32",
					},
					&SchemaMethodParam{
						Description: "Long-poll timeout in seconds",
						Name:        "sectimeout",
						Optional:    true,
						Type:        "uint32",
					},
					&SchemaMethodParam{
						Description: "How many seconds is client considering itself idle, e.g. screen is off",
						Name:        "secidletime",
						Optional:    true,
						Type:        "uint32",
					},
					&SchemaMethodParam{
						Description: "Boolean, 0 (default): return steamid_from in output, 1: return accountid_from",
						Name:        "use_accountids",
						Optional:    true,
						Type:        "uint32",
					},
				),
				Undocumented: false,
				Version:      1,
			},
			&SchemaMethod{
				HTTPMethod:   http.MethodGet,
				Name:         "Logon",
				Params:       NewSchemaMethodParams(),
				Undocumented: true,
				Version:      1,
			},
			&SchemaMethod{
				HTTPMethod:   http.MethodGet,
				Name:         "Logoff",
				Params:       NewSchemaMethodParams(),
				Undocumented: true,
				Version:      1,
			},
			&SchemaMethod{
				HTTPMethod:   http.MethodGet,
				Name:         "Message",
				Params:       NewSchemaMethodParams(),
				Undocumented: true,
				Version:      1,
			},
			&SchemaMethod{
				HTTPMethod:   http.MethodGet,
				Name:         "DeviceInfo",
				Params:       NewSchemaMethodParams(),
				Undocumented: true,
				Version:      1,
			},
			&SchemaMethod{
				HTTPMethod:   http.MethodGet,
				Name:         "Poll",
				Params:       NewSchemaMethodParams(),
				Undocumented: true,
				Version:      1,
			},
		),
		Name:         "ISteamWebUserPresenceOAuth",
		Undocumented: false,
	},
)

// SteamWebUserPresenceOAuth represents interface ISteamWebUserPresenceOAuth.
type SteamWebUserPresenceOAuth struct {
	Client    *Client
	Interface *SchemaInterface
}

// NewSteamWebUserPresenceOAuth creates a new SteamWebUserPresenceOAuth interface.
func NewSteamWebUserPresenceOAuth(c *Client) (*SteamWebUserPresenceOAuth, error) {
	si, err := SchemaSteamWebUserPresenceOAuth.Get(SchemaInterfaceKey{Name: "ISteamWebUserPresenceOAuth"})

	if err != nil {
		return nil, err
	}

	s := &SteamWebUserPresenceOAuth{
		Client:    c,
		Interface: si,
	}

	return s, nil
}

// SteamWebUserPresenceOAuth creates a new SteamWebUserPresenceOAuth interface.
func (c *Client) SteamWebUserPresenceOAuth() (*SteamWebUserPresenceOAuth, error) {
	return NewSteamWebUserPresenceOAuth(c)
}

/*
DeviceInfo creates a Request for interface method DeviceInfo.

This is an undocumented method.
*/
func (i *SteamWebUserPresenceOAuth) DeviceInfo() (*Request, error) {
	sm, err := i.Interface.Methods.Get(SchemaMethodKey{
		Name:    "DeviceInfo",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := &Request{
		Client:    i.Client,
		Interface: i.Interface,
		Method:    sm,
		Result:    &SteamWebUserPresenceOAuthDeviceInfo{},
	}

	return req, nil
}

/*
Logoff creates a Request for interface method Logoff.

This is an undocumented method.
*/
func (i *SteamWebUserPresenceOAuth) Logoff() (*Request, error) {
	sm, err := i.Interface.Methods.Get(SchemaMethodKey{
		Name:    "Logoff",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := &Request{
		Client:    i.Client,
		Interface: i.Interface,
		Method:    sm,
		Result:    &SteamWebUserPresenceOAuthLogoff{},
	}

	return req, nil
}

/*
Logon creates a Request for interface method Logon.

This is an undocumented method.
*/
func (i *SteamWebUserPresenceOAuth) Logon() (*Request, error) {
	sm, err := i.Interface.Methods.Get(SchemaMethodKey{
		Name:    "Logon",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := &Request{
		Client:    i.Client,
		Interface: i.Interface,
		Method:    sm,
		Result:    &SteamWebUserPresenceOAuthLogon{},
	}

	return req, nil
}

/*
Message creates a Request for interface method Message.

This is an undocumented method.
*/
func (i *SteamWebUserPresenceOAuth) Message() (*Request, error) {
	sm, err := i.Interface.Methods.Get(SchemaMethodKey{
		Name:    "Message",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := &Request{
		Client:    i.Client,
		Interface: i.Interface,
		Method:    sm,
		Result:    &SteamWebUserPresenceOAuthMessage{},
	}

	return req, nil
}

/*
Poll creates a Request for interface method Poll.

This is an undocumented method.
*/
func (i *SteamWebUserPresenceOAuth) Poll() (*Request, error) {
	sm, err := i.Interface.Methods.Get(SchemaMethodKey{
		Name:    "Poll",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := &Request{
		Client:    i.Client,
		Interface: i.Interface,
		Method:    sm,
		Result:    &SteamWebUserPresenceOAuthPoll{},
	}

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
func (i *SteamWebUserPresenceOAuth) PollStatus() (*Request, error) {
	sm, err := i.Interface.Methods.Get(SchemaMethodKey{
		Name:    "PollStatus",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := &Request{
		Client:    i.Client,
		Interface: i.Interface,
		Method:    sm,
		Result:    &SteamWebUserPresenceOAuthPollStatus{},
	}

	return req, nil
}
