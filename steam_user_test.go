// Code generated by github.com/13k/geyser/apigen. DO NOT EDIT.

package geyser_test

import (
	"github.com/13k/geyser"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewSteamUser(t *testing.T) {
	client := &geyser.Client{}
	i, err := geyser.NewSteamUser(client)

	require.NoError(t, err)
	require.NotNil(t, i)

	assert.Same(t, client, i.Client)
	assert.NotNil(t, i.Interface)
}

func TestSteamUser_GetFriendList(t *testing.T) {
	var i *geyser.SteamUser
	var err error
	var req *geyser.Request
	var result *geyser.SteamUserGetFriendList
	var ok bool

	client := &geyser.Client{}

	i, err = geyser.NewSteamUser(client)

	require.NoError(t, err)
	require.NotNil(t, i)

	req, err = i.GetFriendList()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, i.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetFriendList", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	result, ok = req.Result.(*geyser.SteamUserGetFriendList)

	if assert.Truef(t, ok, "invalid result type %T", result) {
		assert.NotNil(t, result)
	}
}

func TestSteamUser_GetPlayerBans(t *testing.T) {
	var i *geyser.SteamUser
	var err error
	var req *geyser.Request
	var result *geyser.SteamUserGetPlayerBans
	var ok bool

	client := &geyser.Client{}

	i, err = geyser.NewSteamUser(client)

	require.NoError(t, err)
	require.NotNil(t, i)

	req, err = i.GetPlayerBans()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, i.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetPlayerBans", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	result, ok = req.Result.(*geyser.SteamUserGetPlayerBans)

	if assert.Truef(t, ok, "invalid result type %T", result) {
		assert.NotNil(t, result)
	}
}

func TestSteamUser_GetPlayerSummaries(t *testing.T) {
	var i *geyser.SteamUser
	var err error
	var req *geyser.Request
	var result *geyser.SteamUserGetPlayerSummaries
	var ok bool

	client := &geyser.Client{}

	i, err = geyser.NewSteamUser(client)

	require.NoError(t, err)
	require.NotNil(t, i)

	req, err = i.GetPlayerSummaries(1)

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, i.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetPlayerSummaries", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	result, ok = req.Result.(*geyser.SteamUserGetPlayerSummaries)

	if assert.Truef(t, ok, "invalid result type %T", result) {
		assert.NotNil(t, result)
	}

	i, err = geyser.NewSteamUser(client)

	require.NoError(t, err)
	require.NotNil(t, i)

	req, err = i.GetPlayerSummaries(2)

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, i.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetPlayerSummaries", req.Method.Name)
		assert.Equal(t, 2, req.Method.Version)
	}

	result, ok = req.Result.(*geyser.SteamUserGetPlayerSummaries)

	if assert.Truef(t, ok, "invalid result type %T", result) {
		assert.NotNil(t, result)
	}
}

func TestSteamUser_GetUserGroupList(t *testing.T) {
	var i *geyser.SteamUser
	var err error
	var req *geyser.Request
	var result *geyser.SteamUserGetUserGroupList
	var ok bool

	client := &geyser.Client{}

	i, err = geyser.NewSteamUser(client)

	require.NoError(t, err)
	require.NotNil(t, i)

	req, err = i.GetUserGroupList()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, i.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetUserGroupList", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	result, ok = req.Result.(*geyser.SteamUserGetUserGroupList)

	if assert.Truef(t, ok, "invalid result type %T", result) {
		assert.NotNil(t, result)
	}
}

func TestSteamUser_ResolveVanityURL(t *testing.T) {
	var i *geyser.SteamUser
	var err error
	var req *geyser.Request
	var result *geyser.SteamUserResolveVanityURL
	var ok bool

	client := &geyser.Client{}

	i, err = geyser.NewSteamUser(client)

	require.NoError(t, err)
	require.NotNil(t, i)

	req, err = i.ResolveVanityURL()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, i.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "ResolveVanityURL", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	result, ok = req.Result.(*geyser.SteamUserResolveVanityURL)

	if assert.Truef(t, ok, "invalid result type %T", result) {
		assert.NotNil(t, result)
	}
}