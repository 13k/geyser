// Code generated by geyser. DO NOT EDIT.

package steam_test

import (
	"github.com/13k/geyser/v2"
	"github.com/13k/geyser/v2/steam"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewISteamWebAPIUtil(t *testing.T) {
	client, err := steam.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err := steam.NewISteamWebAPIUtil(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	assert.Same(t, client, ci.Client)
	assert.NotNil(t, ci.Interface)
}

func TestISteamWebAPIUtil_GetServerInfo(t *testing.T) {
	var ci *steam.ISteamWebAPIUtil
	var err error
	var req *geyser.Request

	client, err := steam.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err = steam.NewISteamWebAPIUtil(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	req, err = ci.GetServerInfo()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, ci.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetServerInfo", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestISteamWebAPIUtil_GetSupportedAPIList(t *testing.T) {
	var ci *steam.ISteamWebAPIUtil
	var err error
	var req *geyser.Request

	client, err := steam.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err = steam.NewISteamWebAPIUtil(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	req, err = ci.GetSupportedAPIList()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, ci.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetSupportedAPIList", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}
