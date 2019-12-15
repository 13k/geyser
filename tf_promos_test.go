// Code generated by github.com/13k/geyser/apigen. DO NOT EDIT.

package geyser_test

import (
	"github.com/13k/geyser"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewTFPromos(t *testing.T) {
	client := &geyser.Client{}
	appIDs := []uint32{440, 570, 620, 205790}

	for _, appID := range appIDs {
		i, err := geyser.NewTFPromos(client, appID)

		require.NoError(t, err)
		require.NotNil(t, i)

		assert.Same(t, client, i.Client)
		assert.NotNil(t, i.Interface)
	}
}

func TestTFPromos_GetItemID(t *testing.T) {
	var i *geyser.TFPromos
	var err error
	var req *geyser.Request
	var result *geyser.TFPromosGetItemID
	var ok bool

	client := &geyser.Client{}

	i, err = geyser.NewTFPromos(client, 440)

	require.NoError(t, err)
	require.NotNil(t, i)

	req, err = i.GetItemID()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, i.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetItemID", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	result, ok = req.Result.(*geyser.TFPromosGetItemID)

	if assert.Truef(t, ok, "invalid result type %T", result) {
		assert.NotNil(t, result)
	}

	i, err = geyser.NewTFPromos(client, 570)

	require.NoError(t, err)
	require.NotNil(t, i)

	req, err = i.GetItemID()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, i.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetItemID", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	result, ok = req.Result.(*geyser.TFPromosGetItemID)

	if assert.Truef(t, ok, "invalid result type %T", result) {
		assert.NotNil(t, result)
	}

	i, err = geyser.NewTFPromos(client, 620)

	require.NoError(t, err)
	require.NotNil(t, i)

	req, err = i.GetItemID()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, i.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetItemID", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	result, ok = req.Result.(*geyser.TFPromosGetItemID)

	if assert.Truef(t, ok, "invalid result type %T", result) {
		assert.NotNil(t, result)
	}

	i, err = geyser.NewTFPromos(client, 205790)

	require.NoError(t, err)
	require.NotNil(t, i)

	req, err = i.GetItemID()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, i.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetItemID", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	result, ok = req.Result.(*geyser.TFPromosGetItemID)

	if assert.Truef(t, ok, "invalid result type %T", result) {
		assert.NotNil(t, result)
	}
}

func TestTFPromos_GrantItem(t *testing.T) {
	var i *geyser.TFPromos
	var err error
	var req *geyser.Request
	var result *geyser.TFPromosGrantItem
	var ok bool

	client := &geyser.Client{}

	i, err = geyser.NewTFPromos(client, 440)

	require.NoError(t, err)
	require.NotNil(t, i)

	req, err = i.GrantItem()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, i.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GrantItem", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	result, ok = req.Result.(*geyser.TFPromosGrantItem)

	if assert.Truef(t, ok, "invalid result type %T", result) {
		assert.NotNil(t, result)
	}

	i, err = geyser.NewTFPromos(client, 570)

	require.NoError(t, err)
	require.NotNil(t, i)

	req, err = i.GrantItem()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, i.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GrantItem", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	result, ok = req.Result.(*geyser.TFPromosGrantItem)

	if assert.Truef(t, ok, "invalid result type %T", result) {
		assert.NotNil(t, result)
	}

	i, err = geyser.NewTFPromos(client, 620)

	require.NoError(t, err)
	require.NotNil(t, i)

	req, err = i.GrantItem()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, i.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GrantItem", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	result, ok = req.Result.(*geyser.TFPromosGrantItem)

	if assert.Truef(t, ok, "invalid result type %T", result) {
		assert.NotNil(t, result)
	}

	i, err = geyser.NewTFPromos(client, 205790)

	require.NoError(t, err)
	require.NotNil(t, i)

	req, err = i.GrantItem()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, i.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GrantItem", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	result, ok = req.Result.(*geyser.TFPromosGrantItem)

	if assert.Truef(t, ok, "invalid result type %T", result) {
		assert.NotNil(t, result)
	}
}
