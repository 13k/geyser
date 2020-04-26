// Code generated by geyser. DO NOT EDIT.

package steam_test

import (
	"github.com/13k/geyser"
	"github.com/13k/geyser/steam"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewISteamApps(t *testing.T) {
	client := &steam.Client{}
	iface, err := steam.NewISteamApps(client)

	require.NoError(t, err)
	require.NotNil(t, iface)

	assert.Same(t, client, iface.Client)
	assert.NotNil(t, iface.Interface)
}

func TestISteamApps_GetAppList(t *testing.T) {
	var iface *steam.ISteamApps
	var err error
	var req *geyser.Request

	client := &steam.Client{}

	iface, err = steam.NewISteamApps(client)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetAppList(1)

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetAppList", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	iface, err = steam.NewISteamApps(client)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetAppList(2)

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetAppList", req.Method.Name)
		assert.Equal(t, 2, req.Method.Version)
	}
}

func TestISteamApps_GetSDRConfig(t *testing.T) {
	var iface *steam.ISteamApps
	var err error
	var req *geyser.Request

	client := &steam.Client{}

	iface, err = steam.NewISteamApps(client)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetSDRConfig()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetSDRConfig", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestISteamApps_GetServersAtAddress(t *testing.T) {
	var iface *steam.ISteamApps
	var err error
	var req *geyser.Request

	client := &steam.Client{}

	iface, err = steam.NewISteamApps(client)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetServersAtAddress()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetServersAtAddress", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestISteamApps_UpToDateCheck(t *testing.T) {
	var iface *steam.ISteamApps
	var err error
	var req *geyser.Request

	client := &steam.Client{}

	iface, err = steam.NewISteamApps(client)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.UpToDateCheck()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "UpToDateCheck", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}