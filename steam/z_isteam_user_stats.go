// Code generated by geyser. DO NOT EDIT.
// API interface: ISteamUserStats.

package steam

import (
	"github.com/13k/geyser/v2"
	"github.com/13k/geyser/v2/schema"
	"net/http"
)

// SchemaISteamUserStats stores the Interfaces for interface ISteamUserStats.
var SchemaISteamUserStats = schema.MustNewInterfaces(
	&schema.Interface{
		Methods: schema.MustNewMethods(
			&schema.Method{
				HTTPMethod: http.MethodGet,
				Name:       "GetGlobalAchievementPercentagesForApp",
				Params: schema.NewMethodParams(
					&schema.MethodParam{
						Description: "GameID to retrieve the achievement percentages for",
						Name:        "gameid",
						Optional:    false,
						Type:        "uint64",
					},
				),
				Undocumented: false,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod: http.MethodGet,
				Name:       "GetGlobalAchievementPercentagesForApp",
				Params: schema.NewMethodParams(
					&schema.MethodParam{
						Description: "GameID to retrieve the achievement percentages for",
						Name:        "gameid",
						Optional:    false,
						Type:        "uint64",
					},
				),
				Undocumented: false,
				Version:      2,
			},
			&schema.Method{
				HTTPMethod: http.MethodGet,
				Name:       "GetGlobalStatsForGame",
				Params: schema.NewMethodParams(
					&schema.MethodParam{
						Description: "AppID that we're getting global stats for",
						Name:        "appid",
						Optional:    false,
						Type:        "uint32",
					},
					&schema.MethodParam{
						Description: "Number of stats get data for",
						Name:        "count",
						Optional:    false,
						Type:        "uint32",
					},
					&schema.MethodParam{
						Description: "Names of stat to get data for",
						Name:        "name[0]",
						Optional:    false,
						Type:        "string",
					},
					&schema.MethodParam{
						Description: "Start date for daily totals (unix epoch timestamp)",
						Name:        "startdate",
						Optional:    true,
						Type:        "uint32",
					},
					&schema.MethodParam{
						Description: "End date for daily totals (unix epoch timestamp)",
						Name:        "enddate",
						Optional:    true,
						Type:        "uint32",
					},
				),
				Undocumented: false,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod: http.MethodGet,
				Name:       "GetNumberOfCurrentPlayers",
				Params: schema.NewMethodParams(
					&schema.MethodParam{
						Description: "AppID that we're getting user count for",
						Name:        "appid",
						Optional:    false,
						Type:        "uint32",
					},
				),
				Undocumented: false,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod: http.MethodGet,
				Name:       "GetPlayerAchievements",
				Params: schema.NewMethodParams(
					&schema.MethodParam{
						Description: "access key",
						Name:        "key",
						Optional:    false,
						Type:        "string",
					},
					&schema.MethodParam{
						Description: "SteamID of user",
						Name:        "steamid",
						Optional:    false,
						Type:        "uint64",
					},
					&schema.MethodParam{
						Description: "AppID to get achievements for",
						Name:        "appid",
						Optional:    false,
						Type:        "uint32",
					},
					&schema.MethodParam{
						Description: "Language to return strings for",
						Name:        "l",
						Optional:    true,
						Type:        "string",
					},
				),
				Undocumented: false,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod: http.MethodGet,
				Name:       "GetSchemaForGame",
				Params: schema.NewMethodParams(
					&schema.MethodParam{
						Description: "access key",
						Name:        "key",
						Optional:    false,
						Type:        "string",
					},
					&schema.MethodParam{
						Description: "appid of game",
						Name:        "appid",
						Optional:    false,
						Type:        "uint32",
					},
					&schema.MethodParam{
						Description: "localized langauge to return (english, french, etc.)",
						Name:        "l",
						Optional:    true,
						Type:        "string",
					},
				),
				Undocumented: false,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod: http.MethodGet,
				Name:       "GetSchemaForGame",
				Params: schema.NewMethodParams(
					&schema.MethodParam{
						Description: "access key",
						Name:        "key",
						Optional:    false,
						Type:        "string",
					},
					&schema.MethodParam{
						Description: "appid of game",
						Name:        "appid",
						Optional:    false,
						Type:        "uint32",
					},
					&schema.MethodParam{
						Description: "localized language to return (english, french, etc.)",
						Name:        "l",
						Optional:    true,
						Type:        "string",
					},
				),
				Undocumented: false,
				Version:      2,
			},
			&schema.Method{
				HTTPMethod: http.MethodGet,
				Name:       "GetUserStatsForGame",
				Params: schema.NewMethodParams(
					&schema.MethodParam{
						Description: "access key",
						Name:        "key",
						Optional:    false,
						Type:        "string",
					},
					&schema.MethodParam{
						Description: "SteamID of user",
						Name:        "steamid",
						Optional:    false,
						Type:        "uint64",
					},
					&schema.MethodParam{
						Description: "appid of game",
						Name:        "appid",
						Optional:    false,
						Type:        "uint32",
					},
				),
				Undocumented: false,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod: http.MethodGet,
				Name:       "GetUserStatsForGame",
				Params: schema.NewMethodParams(
					&schema.MethodParam{
						Description: "access key",
						Name:        "key",
						Optional:    false,
						Type:        "string",
					},
					&schema.MethodParam{
						Description: "SteamID of user",
						Name:        "steamid",
						Optional:    false,
						Type:        "uint64",
					},
					&schema.MethodParam{
						Description: "appid of game",
						Name:        "appid",
						Optional:    false,
						Type:        "uint32",
					},
				),
				Undocumented: false,
				Version:      2,
			},
		),
		Name:         "ISteamUserStats",
		Undocumented: false,
	},
)

// ISteamUserStats represents interface ISteamUserStats.
type ISteamUserStats struct {
	Client    *Client
	Interface *schema.Interface
}

// NewISteamUserStats creates a new ISteamUserStats interface.
func NewISteamUserStats(c *Client) (*ISteamUserStats, error) {
	si, err := SchemaISteamUserStats.Get(schema.InterfaceKey{Name: "ISteamUserStats"})

	if err != nil {
		return nil, err
	}

	s := &ISteamUserStats{
		Client:    c,
		Interface: si,
	}

	return s, nil
}

// ISteamUserStats creates a new ISteamUserStats interface.
func (c *Client) ISteamUserStats() (*ISteamUserStats, error) {
	return NewISteamUserStats(c)
}

/*
GetGlobalAchievementPercentagesForApp creates a Request for interface method GetGlobalAchievementPercentagesForApp.

Supported versions: 1, 2.

Parameters (v1)

  * gameid [uint64] (required): GameID to retrieve the achievement percentages for

Parameters (v2)

  * gameid [uint64] (required): GameID to retrieve the achievement percentages for
*/
func (i *ISteamUserStats) GetGlobalAchievementPercentagesForApp(version int) (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "GetGlobalAchievementPercentagesForApp",
		Version: version,
	})

	if err != nil {
		return nil, err
	}

	req := geyser.NewRequest(i.Interface, sm)

	return req, nil
}

/*
GetGlobalStatsForGame creates a Request for interface method GetGlobalStatsForGame.

Parameters

  * appid [uint32] (required): AppID that we're getting global stats for
  * count [uint32] (required): Number of stats get data for
  * name[0] [string] (required): Names of stat to get data for
  * startdate [uint32]: Start date for daily totals (unix epoch timestamp)
  * enddate [uint32]: End date for daily totals (unix epoch timestamp)
*/
func (i *ISteamUserStats) GetGlobalStatsForGame() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "GetGlobalStatsForGame",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := geyser.NewRequest(i.Interface, sm)

	return req, nil
}

/*
GetNumberOfCurrentPlayers creates a Request for interface method GetNumberOfCurrentPlayers.

Parameters

  * appid [uint32] (required): AppID that we're getting user count for
*/
func (i *ISteamUserStats) GetNumberOfCurrentPlayers() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "GetNumberOfCurrentPlayers",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := geyser.NewRequest(i.Interface, sm)

	return req, nil
}

/*
GetPlayerAchievements creates a Request for interface method GetPlayerAchievements.

Parameters

  * key [string] (required): access key
  * steamid [uint64] (required): SteamID of user
  * appid [uint32] (required): AppID to get achievements for
  * l [string]: Language to return strings for
*/
func (i *ISteamUserStats) GetPlayerAchievements() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "GetPlayerAchievements",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := geyser.NewRequest(i.Interface, sm)

	return req, nil
}

/*
GetSchemaForGame creates a Request for interface method GetSchemaForGame.

Supported versions: 1, 2.

Parameters (v1)

  * key [string] (required): access key
  * appid [uint32] (required): appid of game
  * l [string]: localized langauge to return (english, french, etc.)

Parameters (v2)

  * key [string] (required): access key
  * appid [uint32] (required): appid of game
  * l [string]: localized language to return (english, french, etc.)
*/
func (i *ISteamUserStats) GetSchemaForGame(version int) (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "GetSchemaForGame",
		Version: version,
	})

	if err != nil {
		return nil, err
	}

	req := geyser.NewRequest(i.Interface, sm)

	return req, nil
}

/*
GetUserStatsForGame creates a Request for interface method GetUserStatsForGame.

Supported versions: 1, 2.

Parameters (v1)

  * key [string] (required): access key
  * steamid [uint64] (required): SteamID of user
  * appid [uint32] (required): appid of game

Parameters (v2)

  * key [string] (required): access key
  * steamid [uint64] (required): SteamID of user
  * appid [uint32] (required): appid of game
*/
func (i *ISteamUserStats) GetUserStatsForGame(version int) (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "GetUserStatsForGame",
		Version: version,
	})

	if err != nil {
		return nil, err
	}

	req := geyser.NewRequest(i.Interface, sm)

	return req, nil
}
