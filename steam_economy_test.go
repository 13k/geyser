// Code generated by github.com/13k/geyser/apigen. DO NOT EDIT.

package geyser_test

import (
	"github.com/13k/geyser"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewSteamEconomy(t *testing.T) {
	client := &geyser.Client{}
	i, err := geyser.NewSteamEconomy(client)

	require.NoError(t, err)
	require.NotNil(t, i)

	assert.Same(t, client, i.Client)
	assert.NotNil(t, i.Interface)
}

func TestSteamEconomy_GetAssetClassInfo(t *testing.T) {
	var i *geyser.SteamEconomy
	var err error
	var req *geyser.Request
	var result *geyser.SteamEconomyGetAssetClassInfo
	var ok bool

	client := &geyser.Client{}

	i, err = geyser.NewSteamEconomy(client)

	require.NoError(t, err)
	require.NotNil(t, i)

	req, err = i.GetAssetClassInfo()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, i.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetAssetClassInfo", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	result, ok = req.Result.(*geyser.SteamEconomyGetAssetClassInfo)

	if assert.Truef(t, ok, "invalid result type %T", result) {
		assert.NotNil(t, result)
	}
}

func TestSteamEconomy_GetAssetPrices(t *testing.T) {
	var i *geyser.SteamEconomy
	var err error
	var req *geyser.Request
	var result *geyser.SteamEconomyGetAssetPrices
	var ok bool

	client := &geyser.Client{}

	i, err = geyser.NewSteamEconomy(client)

	require.NoError(t, err)
	require.NotNil(t, i)

	req, err = i.GetAssetPrices()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, i.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetAssetPrices", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	result, ok = req.Result.(*geyser.SteamEconomyGetAssetPrices)

	if assert.Truef(t, ok, "invalid result type %T", result) {
		assert.NotNil(t, result)
	}
}