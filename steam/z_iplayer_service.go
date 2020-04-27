// Code generated by geyser. DO NOT EDIT.
// API interface: IPlayerService.

package steam

import (
	"github.com/13k/geyser"
	"github.com/13k/geyser/schema"
	"net/http"
)

// SchemaIPlayerService stores the Interfaces for interface IPlayerService.
var SchemaIPlayerService = schema.MustNewInterfaces(
	&schema.Interface{
		Methods: schema.MustNewMethods(
			&schema.Method{
				HTTPMethod: http.MethodPost,
				Name:       "RecordOfflinePlaytime",
				Params: schema.NewMethodParams(
					&schema.MethodParam{
						Description: "",
						Name:        "steamid",
						Optional:    false,
						Type:        "uint64",
					},
					&schema.MethodParam{
						Description: "",
						Name:        "ticket",
						Optional:    false,
						Type:        "string",
					},
					&schema.MethodParam{
						Description: "",
						Name:        "play_sessions",
						Optional:    false,
						Type:        "{message}",
					},
				),
				Undocumented: false,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod: http.MethodGet,
				Name:       "GetRecentlyPlayedGames",
				Params: schema.NewMethodParams(
					&schema.MethodParam{
						Description: "Access key",
						Name:        "key",
						Optional:    false,
						Type:        "string",
					},
					&schema.MethodParam{
						Description: "The player we're asking about",
						Name:        "steamid",
						Optional:    false,
						Type:        "uint64",
					},
					&schema.MethodParam{
						Description: "The number of games to return (0/unset: all)",
						Name:        "count",
						Optional:    false,
						Type:        "uint32",
					},
				),
				Undocumented: false,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod: http.MethodGet,
				Name:       "GetOwnedGames",
				Params: schema.NewMethodParams(
					&schema.MethodParam{
						Description: "Access key",
						Name:        "key",
						Optional:    false,
						Type:        "string",
					},
					&schema.MethodParam{
						Description: "The player we're asking about",
						Name:        "steamid",
						Optional:    false,
						Type:        "uint64",
					},
					&schema.MethodParam{
						Description: "true if we want additional details (name, icon) about each game",
						Name:        "include_appinfo",
						Optional:    false,
						Type:        "bool",
					},
					&schema.MethodParam{
						Description: "Free games are excluded by default.  If this is set, free games the user has played will be returned.",
						Name:        "include_played_free_games",
						Optional:    false,
						Type:        "bool",
					},
					&schema.MethodParam{
						Description: "if set, restricts result set to the passed in apps",
						Name:        "appids_filter",
						Optional:    false,
						Type:        "uint32",
					},
				),
				Undocumented: false,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod: http.MethodGet,
				Name:       "GetSteamLevel",
				Params: schema.NewMethodParams(
					&schema.MethodParam{
						Description: "Access key",
						Name:        "key",
						Optional:    false,
						Type:        "string",
					},
					&schema.MethodParam{
						Description: "The player we're asking about",
						Name:        "steamid",
						Optional:    false,
						Type:        "uint64",
					},
				),
				Undocumented: false,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod: http.MethodGet,
				Name:       "GetBadges",
				Params: schema.NewMethodParams(
					&schema.MethodParam{
						Description: "Access key",
						Name:        "key",
						Optional:    false,
						Type:        "string",
					},
					&schema.MethodParam{
						Description: "The player we're asking about",
						Name:        "steamid",
						Optional:    false,
						Type:        "uint64",
					},
				),
				Undocumented: false,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod: http.MethodGet,
				Name:       "GetCommunityBadgeProgress",
				Params: schema.NewMethodParams(
					&schema.MethodParam{
						Description: "Access key",
						Name:        "key",
						Optional:    false,
						Type:        "string",
					},
					&schema.MethodParam{
						Description: "The player we're asking about",
						Name:        "steamid",
						Optional:    false,
						Type:        "uint64",
					},
					&schema.MethodParam{
						Description: "The badge we're asking about",
						Name:        "badgeid",
						Optional:    false,
						Type:        "int32",
					},
				),
				Undocumented: false,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod: http.MethodGet,
				Name:       "IsPlayingSharedGame",
				Params: schema.NewMethodParams(
					&schema.MethodParam{
						Description: "Access key",
						Name:        "key",
						Optional:    false,
						Type:        "string",
					},
					&schema.MethodParam{
						Description: "The player we're asking about",
						Name:        "steamid",
						Optional:    false,
						Type:        "uint64",
					},
					&schema.MethodParam{
						Description: "The game player is currently playing",
						Name:        "appid_playing",
						Optional:    false,
						Type:        "uint32",
					},
				),
				Undocumented: false,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod:   http.MethodGet,
				Name:         "GetSteamLevelDistribution",
				Params:       schema.NewMethodParams(),
				Undocumented: true,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod:   http.MethodGet,
				Name:         "GetNicknameList",
				Params:       schema.NewMethodParams(),
				Undocumented: true,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod:   http.MethodGet,
				Name:         "AddFriend",
				Params:       schema.NewMethodParams(),
				Undocumented: true,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod:   http.MethodGet,
				Name:         "RemoveFriend",
				Params:       schema.NewMethodParams(),
				Undocumented: true,
				Version:      1,
			},
			&schema.Method{
				HTTPMethod:   http.MethodGet,
				Name:         "IgnoreFriend",
				Params:       schema.NewMethodParams(),
				Undocumented: true,
				Version:      1,
			},
		),
		Name:         "IPlayerService",
		Undocumented: false,
	},
)

// IPlayerService represents interface IPlayerService.
type IPlayerService struct {
	Client    *Client
	Interface *schema.Interface
}

// NewIPlayerService creates a new IPlayerService interface.
func NewIPlayerService(c *Client) (*IPlayerService, error) {
	si, err := SchemaIPlayerService.Get(schema.InterfaceKey{Name: "IPlayerService"})

	if err != nil {
		return nil, err
	}

	s := &IPlayerService{
		Client:    c,
		Interface: si,
	}

	return s, nil
}

// IPlayerService creates a new IPlayerService interface.
func (c *Client) IPlayerService() (*IPlayerService, error) {
	return NewIPlayerService(c)
}

/*
AddFriend creates a Request for interface method AddFriend.

This is an undocumented method.
*/
func (i *IPlayerService) AddFriend() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "AddFriend",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := geyser.NewRequest(i.Interface, sm)

	return req, nil
}

/*
GetBadges creates a Request for interface method GetBadges.

Parameters

  * key [string] (required): Access key
  * steamid [uint64] (required): The player we're asking about
*/
func (i *IPlayerService) GetBadges() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "GetBadges",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := geyser.NewRequest(i.Interface, sm)

	return req, nil
}

/*
GetCommunityBadgeProgress creates a Request for interface method GetCommunityBadgeProgress.

Parameters

  * key [string] (required): Access key
  * steamid [uint64] (required): The player we're asking about
  * badgeid [int32] (required): The badge we're asking about
*/
func (i *IPlayerService) GetCommunityBadgeProgress() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "GetCommunityBadgeProgress",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := geyser.NewRequest(i.Interface, sm)

	return req, nil
}

/*
GetNicknameList creates a Request for interface method GetNicknameList.

This is an undocumented method.
*/
func (i *IPlayerService) GetNicknameList() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "GetNicknameList",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := geyser.NewRequest(i.Interface, sm)

	return req, nil
}

/*
GetOwnedGames creates a Request for interface method GetOwnedGames.

Parameters

  * key [string] (required): Access key
  * steamid [uint64] (required): The player we're asking about
  * include_appinfo [bool] (required): true if we want additional details (name, icon) about each game
  * include_played_free_games [bool] (required): Free games are excluded by default.  If this is set, free games the user has played will be returned.
  * appids_filter [uint32] (required): if set, restricts result set to the passed in apps
*/
func (i *IPlayerService) GetOwnedGames() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "GetOwnedGames",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := geyser.NewRequest(i.Interface, sm)

	return req, nil
}

/*
GetRecentlyPlayedGames creates a Request for interface method GetRecentlyPlayedGames.

Parameters

  * key [string] (required): Access key
  * steamid [uint64] (required): The player we're asking about
  * count [uint32] (required): The number of games to return (0/unset: all)
*/
func (i *IPlayerService) GetRecentlyPlayedGames() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "GetRecentlyPlayedGames",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := geyser.NewRequest(i.Interface, sm)

	return req, nil
}

/*
GetSteamLevel creates a Request for interface method GetSteamLevel.

Parameters

  * key [string] (required): Access key
  * steamid [uint64] (required): The player we're asking about
*/
func (i *IPlayerService) GetSteamLevel() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "GetSteamLevel",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := geyser.NewRequest(i.Interface, sm)

	return req, nil
}

/*
GetSteamLevelDistribution creates a Request for interface method GetSteamLevelDistribution.

This is an undocumented method.
*/
func (i *IPlayerService) GetSteamLevelDistribution() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "GetSteamLevelDistribution",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := geyser.NewRequest(i.Interface, sm)

	return req, nil
}

/*
IgnoreFriend creates a Request for interface method IgnoreFriend.

This is an undocumented method.
*/
func (i *IPlayerService) IgnoreFriend() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "IgnoreFriend",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := geyser.NewRequest(i.Interface, sm)

	return req, nil
}

/*
IsPlayingSharedGame creates a Request for interface method IsPlayingSharedGame.

Parameters

  * key [string] (required): Access key
  * steamid [uint64] (required): The player we're asking about
  * appid_playing [uint32] (required): The game player is currently playing
*/
func (i *IPlayerService) IsPlayingSharedGame() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "IsPlayingSharedGame",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := geyser.NewRequest(i.Interface, sm)

	return req, nil
}

/*
RecordOfflinePlaytime creates a Request for interface method RecordOfflinePlaytime.

Parameters

  * steamid [uint64] (required)
  * ticket [string] (required)
  * play_sessions [{message}] (required)
*/
func (i *IPlayerService) RecordOfflinePlaytime() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "RecordOfflinePlaytime",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := geyser.NewRequest(i.Interface, sm)

	return req, nil
}

/*
RemoveFriend creates a Request for interface method RemoveFriend.

This is an undocumented method.
*/
func (i *IPlayerService) RemoveFriend() (*geyser.Request, error) {
	sm, err := i.Interface.Methods.Get(schema.MethodKey{
		Name:    "RemoveFriend",
		Version: 1,
	})

	if err != nil {
		return nil, err
	}

	req := geyser.NewRequest(i.Interface, sm)

	return req, nil
}
