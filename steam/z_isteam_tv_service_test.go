// Code generated by geyser. DO NOT EDIT.

package steam_test

import (
	"github.com/13k/geyser"
	"github.com/13k/geyser/steam"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewISteamTVService(t *testing.T) {
	client, err := steam.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err := steam.NewISteamTVService(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	assert.Same(t, client, ci.Client)
	assert.NotNil(t, ci.Interface)
}

func TestISteamTVService_AddChatBan(t *testing.T) {
	var ci *steam.ISteamTVService
	var err error
	var req *geyser.Request

	client, err := steam.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err = steam.NewISteamTVService(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	req, err = ci.AddChatBan()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, ci.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "AddChatBan", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestISteamTVService_AddChatModerator(t *testing.T) {
	var ci *steam.ISteamTVService
	var err error
	var req *geyser.Request

	client, err := steam.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err = steam.NewISteamTVService(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	req, err = ci.AddChatModerator()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, ci.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "AddChatModerator", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestISteamTVService_AddWordBan(t *testing.T) {
	var ci *steam.ISteamTVService
	var err error
	var req *geyser.Request

	client, err := steam.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err = steam.NewISteamTVService(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	req, err = ci.AddWordBan()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, ci.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "AddWordBan", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestISteamTVService_CreateBroadcastChannel(t *testing.T) {
	var ci *steam.ISteamTVService
	var err error
	var req *geyser.Request

	client, err := steam.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err = steam.NewISteamTVService(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	req, err = ci.CreateBroadcastChannel()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, ci.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "CreateBroadcastChannel", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestISteamTVService_FollowBroadcastChannel(t *testing.T) {
	var ci *steam.ISteamTVService
	var err error
	var req *geyser.Request

	client, err := steam.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err = steam.NewISteamTVService(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	req, err = ci.FollowBroadcastChannel()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, ci.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "FollowBroadcastChannel", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestISteamTVService_GetBroadcastChannelBroadcasters(t *testing.T) {
	var ci *steam.ISteamTVService
	var err error
	var req *geyser.Request

	client, err := steam.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err = steam.NewISteamTVService(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	req, err = ci.GetBroadcastChannelBroadcasters()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, ci.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetBroadcastChannelBroadcasters", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestISteamTVService_GetBroadcastChannelClips(t *testing.T) {
	var ci *steam.ISteamTVService
	var err error
	var req *geyser.Request

	client, err := steam.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err = steam.NewISteamTVService(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	req, err = ci.GetBroadcastChannelClips()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, ci.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetBroadcastChannelClips", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestISteamTVService_GetBroadcastChannelID(t *testing.T) {
	var ci *steam.ISteamTVService
	var err error
	var req *geyser.Request

	client, err := steam.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err = steam.NewISteamTVService(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	req, err = ci.GetBroadcastChannelID()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, ci.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetBroadcastChannelID", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestISteamTVService_GetBroadcastChannelImages(t *testing.T) {
	var ci *steam.ISteamTVService
	var err error
	var req *geyser.Request

	client, err := steam.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err = steam.NewISteamTVService(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	req, err = ci.GetBroadcastChannelImages()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, ci.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetBroadcastChannelImages", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestISteamTVService_GetBroadcastChannelInteraction(t *testing.T) {
	var ci *steam.ISteamTVService
	var err error
	var req *geyser.Request

	client, err := steam.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err = steam.NewISteamTVService(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	req, err = ci.GetBroadcastChannelInteraction()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, ci.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetBroadcastChannelInteraction", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestISteamTVService_GetBroadcastChannelLinks(t *testing.T) {
	var ci *steam.ISteamTVService
	var err error
	var req *geyser.Request

	client, err := steam.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err = steam.NewISteamTVService(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	req, err = ci.GetBroadcastChannelLinks()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, ci.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetBroadcastChannelLinks", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestISteamTVService_GetBroadcastChannelProfile(t *testing.T) {
	var ci *steam.ISteamTVService
	var err error
	var req *geyser.Request

	client, err := steam.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err = steam.NewISteamTVService(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	req, err = ci.GetBroadcastChannelProfile()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, ci.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetBroadcastChannelProfile", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestISteamTVService_GetBroadcastChannelStatus(t *testing.T) {
	var ci *steam.ISteamTVService
	var err error
	var req *geyser.Request

	client, err := steam.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err = steam.NewISteamTVService(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	req, err = ci.GetBroadcastChannelStatus()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, ci.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetBroadcastChannelStatus", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestISteamTVService_GetChannels(t *testing.T) {
	var ci *steam.ISteamTVService
	var err error
	var req *geyser.Request

	client, err := steam.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err = steam.NewISteamTVService(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	req, err = ci.GetChannels()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, ci.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetChannels", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestISteamTVService_GetChatBans(t *testing.T) {
	var ci *steam.ISteamTVService
	var err error
	var req *geyser.Request

	client, err := steam.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err = steam.NewISteamTVService(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	req, err = ci.GetChatBans()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, ci.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetChatBans", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestISteamTVService_GetChatModerators(t *testing.T) {
	var ci *steam.ISteamTVService
	var err error
	var req *geyser.Request

	client, err := steam.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err = steam.NewISteamTVService(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	req, err = ci.GetChatModerators()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, ci.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetChatModerators", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestISteamTVService_GetFeatured(t *testing.T) {
	var ci *steam.ISteamTVService
	var err error
	var req *geyser.Request

	client, err := steam.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err = steam.NewISteamTVService(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	req, err = ci.GetFeatured()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, ci.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetFeatured", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestISteamTVService_GetFollowedChannels(t *testing.T) {
	var ci *steam.ISteamTVService
	var err error
	var req *geyser.Request

	client, err := steam.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err = steam.NewISteamTVService(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	req, err = ci.GetFollowedChannels()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, ci.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetFollowedChannels", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestISteamTVService_GetGames(t *testing.T) {
	var ci *steam.ISteamTVService
	var err error
	var req *geyser.Request

	client, err := steam.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err = steam.NewISteamTVService(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	req, err = ci.GetGames()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, ci.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetGames", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestISteamTVService_GetHomePageContents(t *testing.T) {
	var ci *steam.ISteamTVService
	var err error
	var req *geyser.Request

	client, err := steam.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err = steam.NewISteamTVService(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	req, err = ci.GetHomePageContents()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, ci.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetHomePageContents", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestISteamTVService_GetMyBroadcastChannels(t *testing.T) {
	var ci *steam.ISteamTVService
	var err error
	var req *geyser.Request

	client, err := steam.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err = steam.NewISteamTVService(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	req, err = ci.GetMyBroadcastChannels()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, ci.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetMyBroadcastChannels", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestISteamTVService_GetSteamTVUserSettings(t *testing.T) {
	var ci *steam.ISteamTVService
	var err error
	var req *geyser.Request

	client, err := steam.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err = steam.NewISteamTVService(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	req, err = ci.GetSteamTVUserSettings()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, ci.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetSteamTVUserSettings", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestISteamTVService_GetSubscribedChannels(t *testing.T) {
	var ci *steam.ISteamTVService
	var err error
	var req *geyser.Request

	client, err := steam.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err = steam.NewISteamTVService(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	req, err = ci.GetSubscribedChannels()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, ci.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetSubscribedChannels", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestISteamTVService_GetWordBans(t *testing.T) {
	var ci *steam.ISteamTVService
	var err error
	var req *geyser.Request

	client, err := steam.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err = steam.NewISteamTVService(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	req, err = ci.GetWordBans()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, ci.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetWordBans", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestISteamTVService_JoinChat(t *testing.T) {
	var ci *steam.ISteamTVService
	var err error
	var req *geyser.Request

	client, err := steam.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err = steam.NewISteamTVService(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	req, err = ci.JoinChat()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, ci.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "JoinChat", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestISteamTVService_ReportBroadcastChannel(t *testing.T) {
	var ci *steam.ISteamTVService
	var err error
	var req *geyser.Request

	client, err := steam.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err = steam.NewISteamTVService(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	req, err = ci.ReportBroadcastChannel()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, ci.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "ReportBroadcastChannel", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestISteamTVService_Search(t *testing.T) {
	var ci *steam.ISteamTVService
	var err error
	var req *geyser.Request

	client, err := steam.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err = steam.NewISteamTVService(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	req, err = ci.Search()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, ci.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "Search", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestISteamTVService_SetBroadcastChannelImage(t *testing.T) {
	var ci *steam.ISteamTVService
	var err error
	var req *geyser.Request

	client, err := steam.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err = steam.NewISteamTVService(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	req, err = ci.SetBroadcastChannelImage()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, ci.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "SetBroadcastChannelImage", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestISteamTVService_SetBroadcastChannelLinkRegions(t *testing.T) {
	var ci *steam.ISteamTVService
	var err error
	var req *geyser.Request

	client, err := steam.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err = steam.NewISteamTVService(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	req, err = ci.SetBroadcastChannelLinkRegions()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, ci.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "SetBroadcastChannelLinkRegions", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestISteamTVService_SetBroadcastChannelProfile(t *testing.T) {
	var ci *steam.ISteamTVService
	var err error
	var req *geyser.Request

	client, err := steam.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err = steam.NewISteamTVService(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	req, err = ci.SetBroadcastChannelProfile()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, ci.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "SetBroadcastChannelProfile", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestISteamTVService_SetSteamTVUserSettings(t *testing.T) {
	var ci *steam.ISteamTVService
	var err error
	var req *geyser.Request

	client, err := steam.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err = steam.NewISteamTVService(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	req, err = ci.SetSteamTVUserSettings()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, ci.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "SetSteamTVUserSettings", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestISteamTVService_SubscribeBroadcastChannel(t *testing.T) {
	var ci *steam.ISteamTVService
	var err error
	var req *geyser.Request

	client, err := steam.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err = steam.NewISteamTVService(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	req, err = ci.SubscribeBroadcastChannel()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, ci.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "SubscribeBroadcastChannel", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}
