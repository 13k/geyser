// Code generated by geyser. DO NOT EDIT.

package steam_test

import (
	"github.com/13k/geyser"
	"github.com/13k/geyser/steam"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewISteamUserOAuth(t *testing.T) {
	client := &steam.Client{}
	iface, err := steam.NewISteamUserOAuth(client)

	require.NoError(t, err)
	require.NotNil(t, iface)

	assert.Same(t, client, iface.Client)
	assert.NotNil(t, iface.Interface)
}

func TestISteamUserOAuth_GetFriendList(t *testing.T) {
	var iface *steam.ISteamUserOAuth
	var err error
	var req *geyser.Request

	client := &steam.Client{}

	iface, err = steam.NewISteamUserOAuth(client)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetFriendList()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetFriendList", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestISteamUserOAuth_GetGroupList(t *testing.T) {
	var iface *steam.ISteamUserOAuth
	var err error
	var req *geyser.Request

	client := &steam.Client{}

	iface, err = steam.NewISteamUserOAuth(client)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetGroupList()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetGroupList", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestISteamUserOAuth_GetGroupSummaries(t *testing.T) {
	var iface *steam.ISteamUserOAuth
	var err error
	var req *geyser.Request

	client := &steam.Client{}

	iface, err = steam.NewISteamUserOAuth(client)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetGroupSummaries()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetGroupSummaries", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestISteamUserOAuth_GetTokenDetails(t *testing.T) {
	var iface *steam.ISteamUserOAuth
	var err error
	var req *geyser.Request

	client := &steam.Client{}

	iface, err = steam.NewISteamUserOAuth(client)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetTokenDetails()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetTokenDetails", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestISteamUserOAuth_GetUserSummaries(t *testing.T) {
	var iface *steam.ISteamUserOAuth
	var err error
	var req *geyser.Request

	client := &steam.Client{}

	iface, err = steam.NewISteamUserOAuth(client)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetUserSummaries()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetUserSummaries", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestISteamUserOAuth_Search(t *testing.T) {
	var iface *steam.ISteamUserOAuth
	var err error
	var req *geyser.Request

	client := &steam.Client{}

	iface, err = steam.NewISteamUserOAuth(client)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.Search()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "Search", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}