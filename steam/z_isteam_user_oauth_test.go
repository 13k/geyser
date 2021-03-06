// Code generated by geyser. DO NOT EDIT.

package steam_test

import (
	"github.com/13k/geyser/v2"
	"github.com/13k/geyser/v2/steam"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewISteamUserOAuth(t *testing.T) {
	client, err := steam.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err := steam.NewISteamUserOAuth(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	assert.Same(t, client, ci.Client)
	assert.NotNil(t, ci.Interface)
}

func TestISteamUserOAuth_GetFriendList(t *testing.T) {
	var ci *steam.ISteamUserOAuth
	var err error
	var req *geyser.Request

	client, err := steam.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err = steam.NewISteamUserOAuth(client)

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

func TestISteamUserOAuth_GetGroupList(t *testing.T) {
	var ci *steam.ISteamUserOAuth
	var err error
	var req *geyser.Request

	client, err := steam.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err = steam.NewISteamUserOAuth(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	req, err = ci.GetGroupList()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, ci.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetGroupList", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestISteamUserOAuth_GetGroupSummaries(t *testing.T) {
	var ci *steam.ISteamUserOAuth
	var err error
	var req *geyser.Request

	client, err := steam.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err = steam.NewISteamUserOAuth(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	req, err = ci.GetGroupSummaries()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, ci.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetGroupSummaries", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestISteamUserOAuth_GetTokenDetails(t *testing.T) {
	var ci *steam.ISteamUserOAuth
	var err error
	var req *geyser.Request

	client, err := steam.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err = steam.NewISteamUserOAuth(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	req, err = ci.GetTokenDetails()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, ci.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetTokenDetails", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestISteamUserOAuth_GetUserSummaries(t *testing.T) {
	var ci *steam.ISteamUserOAuth
	var err error
	var req *geyser.Request

	client, err := steam.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err = steam.NewISteamUserOAuth(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	req, err = ci.GetUserSummaries()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, ci.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetUserSummaries", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestISteamUserOAuth_Search(t *testing.T) {
	var ci *steam.ISteamUserOAuth
	var err error
	var req *geyser.Request

	client, err := steam.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err = steam.NewISteamUserOAuth(client)

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
