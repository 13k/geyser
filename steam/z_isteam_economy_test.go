// Code generated by geyser. DO NOT EDIT.

package steam_test

import (
	"github.com/13k/geyser"
	"github.com/13k/geyser/steam"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewISteamEconomy(t *testing.T) {
	client, err := steam.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err := steam.NewISteamEconomy(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	assert.Same(t, client, ci.Client)
	assert.NotNil(t, ci.Interface)
}

func TestISteamEconomy_GetAssetClassInfo(t *testing.T) {
	var ci *steam.ISteamEconomy
	var err error
	var req *geyser.Request

	client, err := steam.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err = steam.NewISteamEconomy(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	req, err = ci.GetAssetClassInfo()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, ci.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetAssetClassInfo", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestISteamEconomy_GetAssetPrices(t *testing.T) {
	var ci *steam.ISteamEconomy
	var err error
	var req *geyser.Request

	client, err := steam.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err = steam.NewISteamEconomy(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	req, err = ci.GetAssetPrices()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, ci.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetAssetPrices", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}
