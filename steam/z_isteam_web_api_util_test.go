// Code generated by geyser. DO NOT EDIT.

package steam_test

import (
	"github.com/13k/geyser"
	"github.com/13k/geyser/steam"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewISteamWebAPIUtil(t *testing.T) {
	client := &steam.Client{}
	iface, err := steam.NewISteamWebAPIUtil(client)

	require.NoError(t, err)
	require.NotNil(t, iface)

	assert.Same(t, client, iface.Client)
	assert.NotNil(t, iface.Interface)
}

func TestISteamWebAPIUtil_GetServerInfo(t *testing.T) {
	var iface *steam.ISteamWebAPIUtil
	var err error
	var req *geyser.Request

	client := &steam.Client{}

	iface, err = steam.NewISteamWebAPIUtil(client)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetServerInfo()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetServerInfo", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestISteamWebAPIUtil_GetSupportedAPIList(t *testing.T) {
	var iface *steam.ISteamWebAPIUtil
	var err error
	var req *geyser.Request

	client := &steam.Client{}

	iface, err = steam.NewISteamWebAPIUtil(client)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetSupportedAPIList()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetSupportedAPIList", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}