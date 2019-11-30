// Code generated by github.com/13k/geyser/apigen. DO NOT EDIT.
// API interface: IDOTA2Fantasy

package geyser

import "net/http"

// SchemaDOTA2Fantasy stores the SchemaInterfaces for interface IDOTA2Fantasy
var SchemaDOTA2Fantasy = MustNewSchemaInterfaces(
	&SchemaInterface{
		Methods: NewSchemaMethods(
			&SchemaMethod{
				HTTPMethod: http.MethodGet,
				Name:       "GetFantasyPlayerStats",
				Params: NewSchemaMethodParams(
					&SchemaMethodParam{
						Description: "The fantasy league ID",
						Name:        "FantasyLeagueID",
						Optional:    false,
						Type:        "uint32",
					},
					&SchemaMethodParam{
						Description: "An optional filter for minimum timestamp",
						Name:        "StartTime",
						Optional:    true,
						Type:        "uint32",
					},
					&SchemaMethodParam{
						Description: "An optional filter for maximum timestamp",
						Name:        "EndTime",
						Optional:    true,
						Type:        "uint32",
					},
					&SchemaMethodParam{
						Description: "An optional filter for a specific match",
						Name:        "MatchID",
						Optional:    true,
						Type:        "uint64",
					},
					&SchemaMethodParam{
						Description: "An optional filter for a specific series",
						Name:        "SeriesID",
						Optional:    true,
						Type:        "uint32",
					},
					&SchemaMethodParam{
						Description: "An optional filter for a specific player",
						Name:        "PlayerAccountID",
						Optional:    true,
						Type:        "uint32",
					},
				),
				Version: 1,
			},
			&SchemaMethod{
				HTTPMethod: http.MethodGet,
				Name:       "GetPlayerOfficialInfo",
				Params: NewSchemaMethodParams(
					&SchemaMethodParam{
						Description: "The account ID to look up",
						Name:        "accountid",
						Optional:    false,
						Type:        "uint32",
					},
				),
				Version: 1,
			},
			&SchemaMethod{
				HTTPMethod: http.MethodGet,
				Name:       "GetProPlayerList",
				Params:     NewSchemaMethodParams(),
				Version:    1,
			},
		),
		Name: "IDOTA2Fantasy_205790",
	},
)

// DOTA2Fantasy represents interface IDOTA2Fantasy
// Supported AppIDs: [205790]
type DOTA2Fantasy struct {
	Client    *Client
	Interface *SchemaInterface
}

// NewDOTA2Fantasy creates a new DOTA2Fantasy interface
func NewDOTA2Fantasy(c *Client, appID uint32) (*DOTA2Fantasy, error) {
	si, err := SchemaDOTA2Fantasy.Get("IDOTA2Fantasy", appID)

	if err != nil {
		return nil, err
	}

	s := &DOTA2Fantasy{
		Client:    c,
		Interface: si,
	}

	return s, nil
}

// DOTA2Fantasy creates a new DOTA2Fantasy interface
func (c *Client) DOTA2Fantasy(appID uint32) (*DOTA2Fantasy, error) {
	return NewDOTA2Fantasy(c, appID)
}

// GetFantasyPlayerStats creates a Request for interface method GetFantasyPlayerStats
func (i *DOTA2Fantasy) GetFantasyPlayerStats() (*Request, error) {
	sm, err := i.Interface.Methods.Get("GetFantasyPlayerStats", 1)

	if err != nil {
		return nil, err
	}

	req := &Request{
		Client:    i.Client,
		Interface: i.Interface,
		Method:    sm,
		Result:    &DOTA2FantasyGetFantasyPlayerStats{},
	}

	return req, nil
}

// GetPlayerOfficialInfo creates a Request for interface method GetPlayerOfficialInfo
func (i *DOTA2Fantasy) GetPlayerOfficialInfo() (*Request, error) {
	sm, err := i.Interface.Methods.Get("GetPlayerOfficialInfo", 1)

	if err != nil {
		return nil, err
	}

	req := &Request{
		Client:    i.Client,
		Interface: i.Interface,
		Method:    sm,
		Result:    &DOTA2FantasyGetPlayerOfficialInfo{},
	}

	return req, nil
}

// GetProPlayerList creates a Request for interface method GetProPlayerList
func (i *DOTA2Fantasy) GetProPlayerList() (*Request, error) {
	sm, err := i.Interface.Methods.Get("GetProPlayerList", 1)

	if err != nil {
		return nil, err
	}

	req := &Request{
		Client:    i.Client,
		Interface: i.Interface,
		Method:    sm,
		Result:    &DOTA2FantasyGetProPlayerList{},
	}

	return req, nil
}
