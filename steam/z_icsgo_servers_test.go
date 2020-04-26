// Code generated by geyser. DO NOT EDIT.

package steam_test

import (
	"github.com/13k/geyser"
	"github.com/13k/geyser/steam"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewICSGOServers(t *testing.T) {
	client := &steam.Client{}
	appIDs := []uint32{730}

	for _, appID := range appIDs {
		iface, err := steam.NewICSGOServers(client, appID)

		require.NoError(t, err)
		require.NotNil(t, iface)

		assert.Same(t, client, iface.Client)
		assert.NotNil(t, iface.Interface)
	}
}

func TestICSGOServers_GetGameMapsPlaytime(t *testing.T) {
	var iface *steam.ICSGOServers
	var err error
	var req *geyser.Request

	client := &steam.Client{}

	iface, err = steam.NewICSGOServers(client, 730)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetGameMapsPlaytime()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetGameMapsPlaytime", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}

func TestICSGOServers_GetGameServersStatus(t *testing.T) {
	var iface *steam.ICSGOServers
	var err error
	var req *geyser.Request

	client := &steam.Client{}

	iface, err = steam.NewICSGOServers(client, 730)

	require.NoError(t, err)
	require.NotNil(t, iface)

	req, err = iface.GetGameServersStatus()

	require.NoError(t, err)
	require.NotNil(t, req)

	assert.Same(t, client, req.Client)
	assert.Same(t, iface.Interface, req.Interface)

	if assert.NotNil(t, req.Method) {
		assert.Equal(t, "GetGameServersStatus", req.Method.Name)
		assert.Equal(t, 1, req.Method.Version)
	}
}