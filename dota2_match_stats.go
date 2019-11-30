// Code generated by github.com/13k/geyser/apigen. DO NOT EDIT.
// API interface: IDOTA2MatchStats

package geyser

import "net/http"

// SchemaDOTA2MatchStats stores the SchemaInterfaces for interface IDOTA2MatchStats
var SchemaDOTA2MatchStats = MustNewSchemaInterfaces(
	&SchemaInterface{
		Methods: NewSchemaMethods(
			&SchemaMethod{
				HTTPMethod: http.MethodGet,
				Name:       "GetRealtimeStats",
				Params: NewSchemaMethodParams(
					&SchemaMethodParam{
						Description: "",
						Name:        "server_steam_id",
						Optional:    false,
						Type:        "uint64",
					},
				),
				Version: 1,
			},
		),
		Name: "IDOTA2MatchStats_205790",
	},
	&SchemaInterface{
		Methods: NewSchemaMethods(
			&SchemaMethod{
				HTTPMethod: http.MethodGet,
				Name:       "GetRealtimeStats",
				Params: NewSchemaMethodParams(
					&SchemaMethodParam{
						Description: "",
						Name:        "server_steam_id",
						Optional:    false,
						Type:        "uint64",
					},
				),
				Version: 1,
			},
		),
		Name: "IDOTA2MatchStats_570",
	},
)

// DOTA2MatchStats represents interface IDOTA2MatchStats
// Supported AppIDs: [205790 570]
type DOTA2MatchStats struct {
	Client    *Client
	Interface *SchemaInterface
}

// NewDOTA2MatchStats creates a new DOTA2MatchStats interface
func NewDOTA2MatchStats(c *Client, appID uint32) (*DOTA2MatchStats, error) {
	si, err := SchemaDOTA2MatchStats.Get("IDOTA2MatchStats", appID)

	if err != nil {
		return nil, err
	}

	s := &DOTA2MatchStats{
		Client:    c,
		Interface: si,
	}

	return s, nil
}

// DOTA2MatchStats creates a new DOTA2MatchStats interface
func (c *Client) DOTA2MatchStats(appID uint32) (*DOTA2MatchStats, error) {
	return NewDOTA2MatchStats(c, appID)
}

// GetRealtimeStats creates a Request for interface method GetRealtimeStats
func (i *DOTA2MatchStats) GetRealtimeStats() (*Request, error) {
	sm, err := i.Interface.Methods.Get("GetRealtimeStats", 1)

	if err != nil {
		return nil, err
	}

	req := &Request{
		Client:    i.Client,
		Interface: i.Interface,
		Method:    sm,
		Result:    &DOTA2MatchStatsGetRealtimeStats{},
	}

	return req, nil
}
