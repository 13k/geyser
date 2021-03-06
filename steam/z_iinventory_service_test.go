// Code generated by geyser. DO NOT EDIT.

package steam_test

import (
	"github.com/13k/geyser/v2"
	"github.com/13k/geyser/v2/steam"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewIInventoryService(t *testing.T) {
	client, err := steam.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err := steam.NewIInventoryService(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	assert.Same(t, client, ci.Client)
	assert.NotNil(t, ci.Interface)
}

func TestIInventoryService_AddPromoItem(t *testing.T) {
	var ci *steam.IInventoryService
	var err error
	var req *geyser.Request

	client, err := steam.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err = steam.NewIInventoryService(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	req, err = ci.AddPromoItem()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, ci.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "AddPromoItem", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestIInventoryService_CombineItemStacks(t *testing.T) {
	var ci *steam.IInventoryService
	var err error
	var req *geyser.Request

	client, err := steam.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err = steam.NewIInventoryService(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	req, err = ci.CombineItemStacks()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, ci.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "CombineItemStacks", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestIInventoryService_ConsumeItem(t *testing.T) {
	var ci *steam.IInventoryService
	var err error
	var req *geyser.Request

	client, err := steam.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err = steam.NewIInventoryService(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	req, err = ci.ConsumeItem()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, ci.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "ConsumeItem", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestIInventoryService_ExchangeItem(t *testing.T) {
	var ci *steam.IInventoryService
	var err error
	var req *geyser.Request

	client, err := steam.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err = steam.NewIInventoryService(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	req, err = ci.ExchangeItem()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, ci.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "ExchangeItem", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestIInventoryService_GetInventory(t *testing.T) {
	var ci *steam.IInventoryService
	var err error
	var req *geyser.Request

	client, err := steam.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err = steam.NewIInventoryService(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	req, err = ci.GetInventory()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, ci.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetInventory", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestIInventoryService_GetItemDefMeta(t *testing.T) {
	var ci *steam.IInventoryService
	var err error
	var req *geyser.Request

	client, err := steam.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err = steam.NewIInventoryService(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	req, err = ci.GetItemDefMeta()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, ci.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetItemDefMeta", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestIInventoryService_GetItemDefs(t *testing.T) {
	var ci *steam.IInventoryService
	var err error
	var req *geyser.Request

	client, err := steam.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err = steam.NewIInventoryService(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	req, err = ci.GetItemDefs()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, ci.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetItemDefs", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestIInventoryService_GetPriceSheet(t *testing.T) {
	var ci *steam.IInventoryService
	var err error
	var req *geyser.Request

	client, err := steam.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err = steam.NewIInventoryService(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	req, err = ci.GetPriceSheet()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, ci.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetPriceSheet", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestIInventoryService_SplitItemStack(t *testing.T) {
	var ci *steam.IInventoryService
	var err error
	var req *geyser.Request

	client, err := steam.New()

	require.NoError(t, err)
	require.NotNil(t, client)

	ci, err = steam.NewIInventoryService(client)

	require.NoError(t, err)
	require.NotNil(t, ci)

	req, err = ci.SplitItemStack()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, ci.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "SplitItemStack", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}
