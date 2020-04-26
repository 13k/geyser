// Code generated by geyser. DO NOT EDIT.

package steam_test

import (
	"github.com/13k/geyser"
	"github.com/13k/geyser/schema"
	"github.com/13k/geyser/steam"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewIEconItems(t *testing.T) {
	client := &steam.Client{}
	appIDs := []uint32{440, 570, 620, 730, 205790, 221540, 238460, 583950}

	for _, appID := range appIDs {
		iface, err := steam.NewIEconItems(client, appID)

		require.NoError(t, err)
		require.NotNil(t, iface)

		assert.Same(t, client, iface.Client)
		assert.NotNil(t, iface.Interface)
	}
}

func TestIEconItems_GetEquippedPlayerItems(t *testing.T) {
	var iface *steam.IEconItems
	var err error
	var req *geyser.Request
	var ok bool

	client := &steam.Client{}

	iface, err = steam.NewIEconItems(client, 440)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetEquippedPlayerItems()

	require.Error(t, err)

	_, ok = err.(*schema.InterfaceMethodNotFoundError)

	assert.Truef(t, ok, "invalid error type %T", err)

	iface, err = steam.NewIEconItems(client, 570)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetEquippedPlayerItems()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetEquippedPlayerItems", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	iface, err = steam.NewIEconItems(client, 620)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetEquippedPlayerItems()

	require.Error(t, err)

	_, ok = err.(*schema.InterfaceMethodNotFoundError)

	assert.Truef(t, ok, "invalid error type %T", err)

	iface, err = steam.NewIEconItems(client, 730)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetEquippedPlayerItems()

	require.Error(t, err)

	_, ok = err.(*schema.InterfaceMethodNotFoundError)

	assert.Truef(t, ok, "invalid error type %T", err)

	iface, err = steam.NewIEconItems(client, 205790)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetEquippedPlayerItems()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetEquippedPlayerItems", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	iface, err = steam.NewIEconItems(client, 221540)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetEquippedPlayerItems()

	require.Error(t, err)

	_, ok = err.(*schema.InterfaceMethodNotFoundError)

	assert.Truef(t, ok, "invalid error type %T", err)

	iface, err = steam.NewIEconItems(client, 238460)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetEquippedPlayerItems()

	require.Error(t, err)

	_, ok = err.(*schema.InterfaceMethodNotFoundError)

	assert.Truef(t, ok, "invalid error type %T", err)

	iface, err = steam.NewIEconItems(client, 583950)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetEquippedPlayerItems()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetEquippedPlayerItems", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestIEconItems_GetPlayerItems(t *testing.T) {
	var iface *steam.IEconItems
	var err error
	var req *geyser.Request
	var ok bool

	client := &steam.Client{}

	iface, err = steam.NewIEconItems(client, 440)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetPlayerItems()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetPlayerItems", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	iface, err = steam.NewIEconItems(client, 570)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetPlayerItems()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetPlayerItems", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	iface, err = steam.NewIEconItems(client, 620)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetPlayerItems()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetPlayerItems", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	iface, err = steam.NewIEconItems(client, 730)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetPlayerItems()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetPlayerItems", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	iface, err = steam.NewIEconItems(client, 205790)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetPlayerItems()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetPlayerItems", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	iface, err = steam.NewIEconItems(client, 221540)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetPlayerItems()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetPlayerItems", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	iface, err = steam.NewIEconItems(client, 238460)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetPlayerItems()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetPlayerItems", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	iface, err = steam.NewIEconItems(client, 583950)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetPlayerItems()

	require.Error(t, err)

	_, ok = err.(*schema.InterfaceMethodNotFoundError)

	assert.Truef(t, ok, "invalid error type %T", err)
}

func TestIEconItems_GetSchema(t *testing.T) {
	var iface *steam.IEconItems
	var err error
	var req *geyser.Request
	var ok bool

	client := &steam.Client{}

	iface, err = steam.NewIEconItems(client, 440)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetSchema(1)

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetSchema", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	iface, err = steam.NewIEconItems(client, 440)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetSchema(2)

	require.Error(t, err)

	_, ok = err.(*schema.InterfaceMethodNotFoundError)

	assert.Truef(t, ok, "invalid error type %T", err)

	iface, err = steam.NewIEconItems(client, 570)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetSchema(1)

	require.Error(t, err)

	_, ok = err.(*schema.InterfaceMethodNotFoundError)

	assert.Truef(t, ok, "invalid error type %T", err)

	iface, err = steam.NewIEconItems(client, 570)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetSchema(2)

	require.Error(t, err)

	_, ok = err.(*schema.InterfaceMethodNotFoundError)

	assert.Truef(t, ok, "invalid error type %T", err)

	iface, err = steam.NewIEconItems(client, 620)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetSchema(1)

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetSchema", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	iface, err = steam.NewIEconItems(client, 620)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetSchema(2)

	require.Error(t, err)

	_, ok = err.(*schema.InterfaceMethodNotFoundError)

	assert.Truef(t, ok, "invalid error type %T", err)

	iface, err = steam.NewIEconItems(client, 730)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetSchema(1)

	require.Error(t, err)

	_, ok = err.(*schema.InterfaceMethodNotFoundError)

	assert.Truef(t, ok, "invalid error type %T", err)

	iface, err = steam.NewIEconItems(client, 730)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetSchema(2)

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetSchema", req.Method.Name)
		assert.Equal(t, 2, req.Method.Version)
	}

	iface, err = steam.NewIEconItems(client, 205790)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetSchema(1)

	require.Error(t, err)

	_, ok = err.(*schema.InterfaceMethodNotFoundError)

	assert.Truef(t, ok, "invalid error type %T", err)

	iface, err = steam.NewIEconItems(client, 205790)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetSchema(2)

	require.Error(t, err)

	_, ok = err.(*schema.InterfaceMethodNotFoundError)

	assert.Truef(t, ok, "invalid error type %T", err)

	iface, err = steam.NewIEconItems(client, 221540)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetSchema(1)

	require.Error(t, err)

	_, ok = err.(*schema.InterfaceMethodNotFoundError)

	assert.Truef(t, ok, "invalid error type %T", err)

	iface, err = steam.NewIEconItems(client, 221540)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetSchema(2)

	require.Error(t, err)

	_, ok = err.(*schema.InterfaceMethodNotFoundError)

	assert.Truef(t, ok, "invalid error type %T", err)

	iface, err = steam.NewIEconItems(client, 238460)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetSchema(1)

	require.Error(t, err)

	_, ok = err.(*schema.InterfaceMethodNotFoundError)

	assert.Truef(t, ok, "invalid error type %T", err)

	iface, err = steam.NewIEconItems(client, 238460)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetSchema(2)

	require.Error(t, err)

	_, ok = err.(*schema.InterfaceMethodNotFoundError)

	assert.Truef(t, ok, "invalid error type %T", err)

	iface, err = steam.NewIEconItems(client, 583950)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetSchema(1)

	require.Error(t, err)

	_, ok = err.(*schema.InterfaceMethodNotFoundError)

	assert.Truef(t, ok, "invalid error type %T", err)

	iface, err = steam.NewIEconItems(client, 583950)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetSchema(2)

	require.Error(t, err)

	_, ok = err.(*schema.InterfaceMethodNotFoundError)

	assert.Truef(t, ok, "invalid error type %T", err)
}

func TestIEconItems_GetSchemaItems(t *testing.T) {
	var iface *steam.IEconItems
	var err error
	var req *geyser.Request
	var ok bool

	client := &steam.Client{}

	iface, err = steam.NewIEconItems(client, 440)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetSchemaItems()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetSchemaItems", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	iface, err = steam.NewIEconItems(client, 570)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetSchemaItems()

	require.Error(t, err)

	_, ok = err.(*schema.InterfaceMethodNotFoundError)

	assert.Truef(t, ok, "invalid error type %T", err)

	iface, err = steam.NewIEconItems(client, 620)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetSchemaItems()

	require.Error(t, err)

	_, ok = err.(*schema.InterfaceMethodNotFoundError)

	assert.Truef(t, ok, "invalid error type %T", err)

	iface, err = steam.NewIEconItems(client, 730)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetSchemaItems()

	require.Error(t, err)

	_, ok = err.(*schema.InterfaceMethodNotFoundError)

	assert.Truef(t, ok, "invalid error type %T", err)

	iface, err = steam.NewIEconItems(client, 205790)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetSchemaItems()

	require.Error(t, err)

	_, ok = err.(*schema.InterfaceMethodNotFoundError)

	assert.Truef(t, ok, "invalid error type %T", err)

	iface, err = steam.NewIEconItems(client, 221540)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetSchemaItems()

	require.Error(t, err)

	_, ok = err.(*schema.InterfaceMethodNotFoundError)

	assert.Truef(t, ok, "invalid error type %T", err)

	iface, err = steam.NewIEconItems(client, 238460)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetSchemaItems()

	require.Error(t, err)

	_, ok = err.(*schema.InterfaceMethodNotFoundError)

	assert.Truef(t, ok, "invalid error type %T", err)

	iface, err = steam.NewIEconItems(client, 583950)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetSchemaItems()

	require.Error(t, err)

	_, ok = err.(*schema.InterfaceMethodNotFoundError)

	assert.Truef(t, ok, "invalid error type %T", err)
}

func TestIEconItems_GetSchemaOverview(t *testing.T) {
	var iface *steam.IEconItems
	var err error
	var req *geyser.Request
	var ok bool

	client := &steam.Client{}

	iface, err = steam.NewIEconItems(client, 440)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetSchemaOverview()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetSchemaOverview", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	iface, err = steam.NewIEconItems(client, 570)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetSchemaOverview()

	require.Error(t, err)

	_, ok = err.(*schema.InterfaceMethodNotFoundError)

	assert.Truef(t, ok, "invalid error type %T", err)

	iface, err = steam.NewIEconItems(client, 620)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetSchemaOverview()

	require.Error(t, err)

	_, ok = err.(*schema.InterfaceMethodNotFoundError)

	assert.Truef(t, ok, "invalid error type %T", err)

	iface, err = steam.NewIEconItems(client, 730)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetSchemaOverview()

	require.Error(t, err)

	_, ok = err.(*schema.InterfaceMethodNotFoundError)

	assert.Truef(t, ok, "invalid error type %T", err)

	iface, err = steam.NewIEconItems(client, 205790)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetSchemaOverview()

	require.Error(t, err)

	_, ok = err.(*schema.InterfaceMethodNotFoundError)

	assert.Truef(t, ok, "invalid error type %T", err)

	iface, err = steam.NewIEconItems(client, 221540)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetSchemaOverview()

	require.Error(t, err)

	_, ok = err.(*schema.InterfaceMethodNotFoundError)

	assert.Truef(t, ok, "invalid error type %T", err)

	iface, err = steam.NewIEconItems(client, 238460)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetSchemaOverview()

	require.Error(t, err)

	_, ok = err.(*schema.InterfaceMethodNotFoundError)

	assert.Truef(t, ok, "invalid error type %T", err)

	iface, err = steam.NewIEconItems(client, 583950)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetSchemaOverview()

	require.Error(t, err)

	_, ok = err.(*schema.InterfaceMethodNotFoundError)

	assert.Truef(t, ok, "invalid error type %T", err)
}

func TestIEconItems_GetSchemaURL(t *testing.T) {
	var iface *steam.IEconItems
	var err error
	var req *geyser.Request
	var ok bool

	client := &steam.Client{}

	iface, err = steam.NewIEconItems(client, 440)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetSchemaURL(1)

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetSchemaURL", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	iface, err = steam.NewIEconItems(client, 440)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetSchemaURL(2)

	require.Error(t, err)

	_, ok = err.(*schema.InterfaceMethodNotFoundError)

	assert.Truef(t, ok, "invalid error type %T", err)

	iface, err = steam.NewIEconItems(client, 570)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetSchemaURL(1)

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetSchemaURL", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	iface, err = steam.NewIEconItems(client, 570)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetSchemaURL(2)

	require.Error(t, err)

	_, ok = err.(*schema.InterfaceMethodNotFoundError)

	assert.Truef(t, ok, "invalid error type %T", err)

	iface, err = steam.NewIEconItems(client, 620)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetSchemaURL(1)

	require.Error(t, err)

	_, ok = err.(*schema.InterfaceMethodNotFoundError)

	assert.Truef(t, ok, "invalid error type %T", err)

	iface, err = steam.NewIEconItems(client, 620)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetSchemaURL(2)

	require.Error(t, err)

	_, ok = err.(*schema.InterfaceMethodNotFoundError)

	assert.Truef(t, ok, "invalid error type %T", err)

	iface, err = steam.NewIEconItems(client, 730)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetSchemaURL(1)

	require.Error(t, err)

	_, ok = err.(*schema.InterfaceMethodNotFoundError)

	assert.Truef(t, ok, "invalid error type %T", err)

	iface, err = steam.NewIEconItems(client, 730)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetSchemaURL(2)

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetSchemaURL", req.Method.Name)
		assert.Equal(t, 2, req.Method.Version)
	}

	iface, err = steam.NewIEconItems(client, 205790)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetSchemaURL(1)

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetSchemaURL", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	iface, err = steam.NewIEconItems(client, 205790)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetSchemaURL(2)

	require.Error(t, err)

	_, ok = err.(*schema.InterfaceMethodNotFoundError)

	assert.Truef(t, ok, "invalid error type %T", err)

	iface, err = steam.NewIEconItems(client, 221540)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetSchemaURL(1)

	require.Error(t, err)

	_, ok = err.(*schema.InterfaceMethodNotFoundError)

	assert.Truef(t, ok, "invalid error type %T", err)

	iface, err = steam.NewIEconItems(client, 221540)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetSchemaURL(2)

	require.Error(t, err)

	_, ok = err.(*schema.InterfaceMethodNotFoundError)

	assert.Truef(t, ok, "invalid error type %T", err)

	iface, err = steam.NewIEconItems(client, 238460)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetSchemaURL(1)

	require.Error(t, err)

	_, ok = err.(*schema.InterfaceMethodNotFoundError)

	assert.Truef(t, ok, "invalid error type %T", err)

	iface, err = steam.NewIEconItems(client, 238460)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetSchemaURL(2)

	require.Error(t, err)

	_, ok = err.(*schema.InterfaceMethodNotFoundError)

	assert.Truef(t, ok, "invalid error type %T", err)

	iface, err = steam.NewIEconItems(client, 583950)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetSchemaURL(1)

	require.Error(t, err)

	_, ok = err.(*schema.InterfaceMethodNotFoundError)

	assert.Truef(t, ok, "invalid error type %T", err)

	iface, err = steam.NewIEconItems(client, 583950)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetSchemaURL(2)

	require.Error(t, err)

	_, ok = err.(*schema.InterfaceMethodNotFoundError)

	assert.Truef(t, ok, "invalid error type %T", err)
}

func TestIEconItems_GetStoreMetaData(t *testing.T) {
	var iface *steam.IEconItems
	var err error
	var req *geyser.Request
	var ok bool

	client := &steam.Client{}

	iface, err = steam.NewIEconItems(client, 440)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetStoreMetaData()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetStoreMetaData", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	iface, err = steam.NewIEconItems(client, 570)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetStoreMetaData()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetStoreMetaData", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	iface, err = steam.NewIEconItems(client, 620)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetStoreMetaData()

	require.Error(t, err)

	_, ok = err.(*schema.InterfaceMethodNotFoundError)

	assert.Truef(t, ok, "invalid error type %T", err)

	iface, err = steam.NewIEconItems(client, 730)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetStoreMetaData()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetStoreMetaData", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	iface, err = steam.NewIEconItems(client, 205790)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetStoreMetaData()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetStoreMetaData", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	iface, err = steam.NewIEconItems(client, 221540)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetStoreMetaData()

	require.Error(t, err)

	_, ok = err.(*schema.InterfaceMethodNotFoundError)

	assert.Truef(t, ok, "invalid error type %T", err)

	iface, err = steam.NewIEconItems(client, 238460)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetStoreMetaData()

	require.Error(t, err)

	_, ok = err.(*schema.InterfaceMethodNotFoundError)

	assert.Truef(t, ok, "invalid error type %T", err)

	iface, err = steam.NewIEconItems(client, 583950)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetStoreMetaData()

	require.Error(t, err)

	_, ok = err.(*schema.InterfaceMethodNotFoundError)

	assert.Truef(t, ok, "invalid error type %T", err)
}

func TestIEconItems_GetStoreStatus(t *testing.T) {
	var iface *steam.IEconItems
	var err error
	var req *geyser.Request
	var ok bool

	client := &steam.Client{}

	iface, err = steam.NewIEconItems(client, 440)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetStoreStatus()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetStoreStatus", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}

	iface, err = steam.NewIEconItems(client, 570)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetStoreStatus()

	require.Error(t, err)

	_, ok = err.(*schema.InterfaceMethodNotFoundError)

	assert.Truef(t, ok, "invalid error type %T", err)

	iface, err = steam.NewIEconItems(client, 620)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetStoreStatus()

	require.Error(t, err)

	_, ok = err.(*schema.InterfaceMethodNotFoundError)

	assert.Truef(t, ok, "invalid error type %T", err)

	iface, err = steam.NewIEconItems(client, 730)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetStoreStatus()

	require.Error(t, err)

	_, ok = err.(*schema.InterfaceMethodNotFoundError)

	assert.Truef(t, ok, "invalid error type %T", err)

	iface, err = steam.NewIEconItems(client, 205790)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetStoreStatus()

	require.Error(t, err)

	_, ok = err.(*schema.InterfaceMethodNotFoundError)

	assert.Truef(t, ok, "invalid error type %T", err)

	iface, err = steam.NewIEconItems(client, 221540)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetStoreStatus()

	require.Error(t, err)

	_, ok = err.(*schema.InterfaceMethodNotFoundError)

	assert.Truef(t, ok, "invalid error type %T", err)

	iface, err = steam.NewIEconItems(client, 238460)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetStoreStatus()

	require.Error(t, err)

	_, ok = err.(*schema.InterfaceMethodNotFoundError)

	assert.Truef(t, ok, "invalid error type %T", err)

	iface, err = steam.NewIEconItems(client, 583950)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetStoreStatus()

	require.Error(t, err)

	_, ok = err.(*schema.InterfaceMethodNotFoundError)

	assert.Truef(t, ok, "invalid error type %T", err)
}