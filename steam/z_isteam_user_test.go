// Code generated by geyser. DO NOT EDIT.

package steam_test

import (
	"github.com/13k/geyser/v2"
	"github.com/13k/geyser/v2/steam"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewISteamUser(t *testing.T) {
	client, err := steam.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err := steam.NewISteamUser(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	assert.Same(t, client, ci.Client)
	assert.NotNil(t, ci.Interface)
}

func TestISteamUser_GetFriendList(t *testing.T) {
	var ci *steam.ISteamUser
	var err error
	var req *geyser.Request

	client, err := steam.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err = steam.NewISteamUser(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	req, err = ci.GetFriendList()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, ci.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetFriendList", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestISteamUser_GetPlayerBans(t *testing.T) {
	var ci *steam.ISteamUser
	var err error
	var req *geyser.Request

	client, err := steam.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err = steam.NewISteamUser(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	req, err = ci.GetPlayerBans()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, ci.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetPlayerBans", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestISteamUser_GetPlayerSummaries(t *testing.T) {
	var ci *steam.ISteamUser
	var err error
	var req *geyser.Request

	client, err := steam.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err = steam.NewISteamUser(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	req, err = ci.GetPlayerSummaries(1)

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, ci.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetPlayerSummaries", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	ci, err = steam.NewISteamUser(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	req, err = ci.GetPlayerSummaries(2)

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, ci.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetPlayerSummaries", req.Method.Name)
		assert.Equal(t, 2, req.Method.Version)
	}
}

func TestISteamUser_GetUserGroupList(t *testing.T) {
	var ci *steam.ISteamUser
	var err error
	var req *geyser.Request

	client, err := steam.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err = steam.NewISteamUser(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	req, err = ci.GetUserGroupList()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, ci.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetUserGroupList", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestISteamUser_ResolveVanityURL(t *testing.T) {
	var ci *steam.ISteamUser
	var err error
	var req *geyser.Request

	client, err := steam.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err = steam.NewISteamUser(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	req, err = ci.ResolveVanityURL()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, ci.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "ResolveVanityURL", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}
