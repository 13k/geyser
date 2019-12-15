// Code generated by github.com/13k/geyser/apigen. DO NOT EDIT.

package geyser_test

import (
	"github.com/13k/geyser"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewSteamNews(t *testing.T) {
	client := &geyser.Client{}
	i, err := geyser.NewSteamNews(client)

	require.NoError(t, err)
	require.NotNil(t, i)

	assert.Same(t, client, i.Client)
	assert.NotNil(t, i.Interface)
}

func TestSteamNews_GetNewsForApp(t *testing.T) {
	var i *geyser.SteamNews
	var err error
	var req *geyser.Request
	var result *geyser.SteamNewsGetNewsForApp
	var ok bool

	client := &geyser.Client{}

	i, err = geyser.NewSteamNews(client)

	require.NoError(t, err)
	require.NotNil(t, i)

	req, err = i.GetNewsForApp(1)

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, i.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetNewsForApp", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	result, ok = req.Result.(*geyser.SteamNewsGetNewsForApp)

	if assert.Truef(t, ok, "invalid result type %T", result) {
		assert.NotNil(t, result)
	}

	i, err = geyser.NewSteamNews(client)

	require.NoError(t, err)
	require.NotNil(t, i)

	req, err = i.GetNewsForApp(2)

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, i.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetNewsForApp", req.Method.Name)
		assert.Equal(t, 2, req.Method.Version)
	}

	result, ok = req.Result.(*geyser.SteamNewsGetNewsForApp)

	if assert.Truef(t, ok, "invalid result type %T", result) {
		assert.NotNil(t, result)
	}
}